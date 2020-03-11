package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedExamples_Get(t *testing.T) {
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
		want   *Example
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedExamples()}, args{"skipParam"}, &Example{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedExamples()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedExamples{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedExamples.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedExamples_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantExample *Example
		wantOK      bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedExamples()}, args{"limitParam"}, &Example{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedExamples()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedExamples{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantExample) {
				t.Errorf("OrderedExamples.GetOK() got = %v, want %v", got, tt.wantExample)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedExamples.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedExamples_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Example
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedExamples()}, args{"getParam", &Example{Description: "Getting OrderedExamples"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedExamples()}, args{"limitParam", &Example{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedExamples()}, args{"", &Example{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedExamples{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedExamples.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedExamples.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedExamples.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedExamples_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Example) error
	}
	type foundExample struct {
		parameter *Example
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundExample
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedExamples fixture",
			fields{buildOrderMapForOrderedExamples()},
			map[string]*foundExample{
				"skipParam":  &foundExample{&Example{Description: "default parameter"}, false},
				"limitParam": &foundExample{&Example{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedExamples",
			fields{},
			map[string]*foundExample{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedExamples{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotExample *Example) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotExample) {
						t.Fatalf("OrderedExamples.ForEach() for key = %s val = %v, want = %v", key, gotExample, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedExamples.ForEach() for key = %s val = %v, want = %v", key, gotExample, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedExamples.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedExamples.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedExamples_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedExamples fixture", fields{buildOrderMapForOrderedExamples()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedExamples", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedExamples{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedExamples.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedExamples() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedExamples() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &Example{Description: "default parameter"},
			"limitParam": &Example{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedExamplesFixture() OrderedExamples {
	m := OrderedExamples{
		data: buildOrderMapForOrderedExamples(),
	}

	return m
}
