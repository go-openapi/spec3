package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedHeaders_Get(t *testing.T) {
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
		want   *Header
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedHeaders()}, args{"skipParam"}, &Header{Parameter{Description: "default parameter"}}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedHeaders()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedHeaders{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedHeaders.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedHeaders_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHeader *Header
		wantOK     bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedHeaders()}, args{"limitParam"}, &Header{Parameter{Description: "OK"}}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedHeaders()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedHeaders{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantHeader) {
				t.Errorf("OrderedHeaders.GetOK() got = %v, want %v", got, tt.wantHeader)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedHeaders.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedHeaders_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Header
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedHeaders()}, args{"getParam", &Header{Parameter{Description: "Getting OrderedHeaders"}}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedHeaders()}, args{"limitParam", &Header{Parameter{Description: "another OK"}}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedHeaders()}, args{"", &Header{Parameter{Description: "description of item #empty"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedHeaders{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedHeaders.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedHeaders.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedHeaders.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedHeaders_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Header) error
	}
	type foundHeader struct {
		parameter *Header
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundHeader
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedHeaders fixture",
			fields{buildOrderMapForOrderedHeaders()},
			map[string]*foundHeader{
				"skipParam":  &foundHeader{&Header{Parameter{Description: "default parameter"}}, false},
				"limitParam": &foundHeader{&Header{Parameter{Description: "OK"}}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedHeaders",
			fields{},
			map[string]*foundHeader{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedHeaders{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotHeader *Header) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotHeader) {
						t.Fatalf("OrderedHeaders.ForEach() for key = %s val = %v, want = %v", key, gotHeader, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedHeaders.ForEach() for key = %s val = %v, want = %v", key, gotHeader, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedHeaders.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedHeaders.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedHeaders_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedHeaders fixture", fields{buildOrderMapForOrderedHeaders()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedHeaders", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedHeaders{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedHeaders.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedHeaders() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedHeaders() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &Header{Parameter{Description: "default parameter"}},
			"limitParam": &Header{Parameter{Description: "OK"}},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedHeadersFixture() OrderedHeaders {
	m := OrderedHeaders{
		data: buildOrderMapForOrderedHeaders(),
	}

	return m
}
