package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedSecuritySchemes_Get(t *testing.T) {
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
		want   *SecurityScheme
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"skipParam"}, &SecurityScheme{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSecuritySchemes{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSecuritySchemes.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSecuritySchemes_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantSecurityScheme *SecurityScheme
		wantOK             bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"limitParam"}, &SecurityScheme{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSecuritySchemes{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantSecurityScheme) {
				t.Errorf("OrderedSecuritySchemes.GetOK() got = %v, want %v", got, tt.wantSecurityScheme)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedSecuritySchemes.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedSecuritySchemes_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *SecurityScheme
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"getParam", &SecurityScheme{Description: "Getting OrderedSecuritySchemes"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"limitParam", &SecurityScheme{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedSecuritySchemes()}, args{"", &SecurityScheme{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSecuritySchemes{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedSecuritySchemes.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedSecuritySchemes.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedSecuritySchemes.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedSecuritySchemes_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *SecurityScheme) error
	}
	type foundSecurityScheme struct {
		parameter *SecurityScheme
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundSecurityScheme
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedSecuritySchemes fixture",
			fields{buildOrderMapForOrderedSecuritySchemes()},
			map[string]*foundSecurityScheme{
				"skipParam":  &foundSecurityScheme{&SecurityScheme{Description: "default parameter"}, false},
				"limitParam": &foundSecurityScheme{&SecurityScheme{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedSecuritySchemes",
			fields{},
			map[string]*foundSecurityScheme{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSecuritySchemes{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotSecurityScheme *SecurityScheme) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotSecurityScheme) {
						t.Fatalf("OrderedSecuritySchemes.ForEach() for key = %s val = %v, want = %v", key, gotSecurityScheme, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedSecuritySchemes.ForEach() for key = %s val = %v, want = %v", key, gotSecurityScheme, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedSecuritySchemes.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedSecuritySchemes.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedSecuritySchemes_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedSecuritySchemes fixture", fields{buildOrderMapForOrderedSecuritySchemes()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedSecuritySchemes", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSecuritySchemes{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedSecuritySchemes.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedSecuritySchemes() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedSecuritySchemes() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &SecurityScheme{Description: "default parameter"},
			"limitParam": &SecurityScheme{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedSecuritySchemesFixture() OrderedSecuritySchemes {
	m := OrderedSecuritySchemes{
		data: buildOrderMapForOrderedSecuritySchemes(),
	}

	return m
}
