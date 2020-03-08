package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedRequestBodies_Get(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RequestBody
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"skipParam"}, &RequestBody{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedRequestBodies{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedRequestBodies.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedRequestBodies_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantRequestBody *RequestBody
		wantOK          bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"limitParam"}, &RequestBody{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedRequestBodies{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantRequestBody) {
				t.Errorf("OrderedRequestBodies.GetOK() got = %v, want %v", got, tt.wantRequestBody)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedRequestBodies.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedRequestBodies_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *RequestBody
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"getParam", &RequestBody{Description: "Getting OrderedRequestBodies"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"limitParam", &RequestBody{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedRequestBodies()}, args{"", &RequestBody{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedRequestBodies{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedRequestBodies.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedRequestBodies.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedRequestBodies.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedRequestBodies_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *RequestBody) error
	}
	type foundRequestBody struct {
		parameter *RequestBody
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundRequestBody
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedRequestBodies fixture",
			fields{buildOrderMapForOrderedRequestBodies()},
			map[string]*foundRequestBody{
				"skipParam":  &foundRequestBody{&RequestBody{Description: "default parameter"}, false},
				"limitParam": &foundRequestBody{&RequestBody{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedRequestBodies",
			fields{},
			map[string]*foundRequestBody{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedRequestBodies{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotRequestBody *RequestBody) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotRequestBody) {
						t.Fatalf("OrderedRequestBodies.ForEach() for key = %s val = %v, want = %v", key, gotRequestBody, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedRequestBodies.ForEach() for key = %s val = %v, want = %v", key, gotRequestBody, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedRequestBodies.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedRequestBodies.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedRequestBodies_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedRequestBodies fixture", fields{buildOrderMapForOrderedRequestBodies()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedRequestBodies", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedRequestBodies{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedRequestBodies.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedRequestBodies() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedRequestBodies() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &RequestBody{Description: "default parameter"},
			"limitParam": &RequestBody{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedRequestBodiesFixture() OrderedRequestBodies {
	m := OrderedRequestBodies{
		data: buildOrderMapForOrderedRequestBodies(),
	}

	return m
}
