package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedSchemas_Get(t *testing.T) {
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
		want   *Schema
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedSchemas()}, args{"skipParam"}, buildSchema("default parameter")},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedSchemas()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSchemas{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedSchemas.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSchemas_GetOK(t *testing.T) {
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
		wantSchema *Schema
		wantOK     bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedSchemas()}, args{"limitParam"}, buildSchema("OK"), true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedSchemas()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSchemas{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantSchema) {
				t.Errorf("OrderedSchemas.GetOK() got = %v, want %v", got, tt.wantSchema)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedSchemas.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedSchemas_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedSchemas()}, args{"getParam", buildSchema("Getting OrderedSchemas")}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedSchemas()}, args{"limitParam", buildSchema("another OK")}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedSchemas()}, args{"", buildSchema("description of item #empty")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSchemas{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedSchemas.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedSchemas.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedSchemas.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedSchemas_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Schema) error
	}
	type foundSchema struct {
		parameter *Schema
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundSchema
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedSchemas fixture",
			fields{buildOrderMapForOrderedSchemas()},
			map[string]*foundSchema{
				"skipParam":  &foundSchema{buildSchema("default parameter"), false},
				"limitParam": &foundSchema{buildSchema("OK"), false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedSchemas",
			fields{},
			map[string]*foundSchema{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSchemas{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotSchema *Schema) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotSchema) {
						t.Fatalf("OrderedSchemas.ForEach() for key = %s val = %v, want = %v", key, gotSchema, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedSchemas.ForEach() for key = %s val = %v, want = %v", key, gotSchema, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedSchemas.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedSchemas.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedSchemas_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedSchemas fixture", fields{buildOrderMapForOrderedSchemas()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedSchemas", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedSchemas{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedSchemas.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildSchema(value string) *Schema {
	extensions := Extensions{}
	extensions.Set("x-some-key", value)
	return &Schema{VendorExtensible{extensions}, Reference{}}
}

func buildEmptyOrderMapForOrderedSchemas() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedSchemas() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  buildSchema("default parameter"),
			"limitParam": buildSchema("OK"),
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedSchemasFixture() OrderedSchemas {
	m := OrderedSchemas{
		data: buildOrderMapForOrderedSchemas(),
	}

	return m
}
