package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedMediaTypes_Get(t *testing.T) {
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
		want   *MediaType
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"skipParam"}, &MediaType{Example: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedMediaTypes{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedMediaTypes.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedMediaTypes_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantMediaType *MediaType
		wantOK        bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"limitParam"}, &MediaType{Example: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedMediaTypes{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantMediaType) {
				t.Errorf("OrderedMediaTypes.GetOK() got = %v, want %v", got, tt.wantMediaType)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedMediaTypes.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedMediaTypes_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *MediaType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"getParam", &MediaType{Example: "Getting OrderedMediaTypes"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"limitParam", &MediaType{Example: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedMediaTypes()}, args{"", &MediaType{Example: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedMediaTypes{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedMediaTypes.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedMediaTypes.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedMediaTypes.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedMediaTypes_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *MediaType) error
	}
	type foundMediaType struct {
		parameter *MediaType
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundMediaType
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedMediaTypes fixture",
			fields{buildOrderMapForOrderedMediaTypes()},
			map[string]*foundMediaType{
				"skipParam":  &foundMediaType{&MediaType{Example: "default parameter"}, false},
				"limitParam": &foundMediaType{&MediaType{Example: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedMediaTypes",
			fields{},
			map[string]*foundMediaType{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedMediaTypes{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotMediaType *MediaType) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotMediaType) {
						t.Fatalf("OrderedMediaTypes.ForEach() for key = %s val = %v, want = %v", key, gotMediaType, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedMediaTypes.ForEach() for key = %s val = %v, want = %v", key, gotMediaType, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedMediaTypes.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedMediaTypes.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedMediaTypes_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedMediaTypes fixture", fields{buildOrderMapForOrderedMediaTypes()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedMediaTypes", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedMediaTypes{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedMediaTypes.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedMediaTypes() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedMediaTypes() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &MediaType{Example: "default parameter"},
			"limitParam": &MediaType{Example: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedMediaTypesFixture() OrderedMediaTypes {
	m := OrderedMediaTypes{
		data: buildOrderMapForOrderedMediaTypes(),
	}

	return m
}
