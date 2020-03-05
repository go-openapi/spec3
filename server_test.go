package spec3

import (
	"reflect"
	"testing"
)

func TestServerVariables_Get(t *testing.T) {
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
		want   *ServerVariable
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForServerVariables()}, args{"skipParam"}, &ServerVariable{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForServerVariables()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerVariables{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerVariables.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerVariables_GetOK(t *testing.T) {
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
		wantServerVariable *ServerVariable
		wantOK             bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForServerVariables()}, args{"limitParam"}, &ServerVariable{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForServerVariables()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerVariables{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantServerVariable) {
				t.Errorf("ServerVariables.GetOK() got = %v, want %v", got, tt.wantServerVariable)
			}
			if got1 != tt.wantOK {
				t.Errorf("ServerVariables.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestServerVariables_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *ServerVariable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForServerVariables()}, args{"getParam", &ServerVariable{Description: "Getting ServerVariables"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForServerVariables()}, args{"limitParam", &ServerVariable{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForServerVariables()}, args{"", &ServerVariable{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerVariables{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("ServerVariables.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("ServerVariables.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("ServerVariables.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestServerVariables_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *ServerVariable) error
	}
	type foundServerVariable struct {
		parameter *ServerVariable
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundServerVariable
		wantErr          error
	}{
		{
			"Should iterate 4 items for ServerVariables fixture",
			fields{buildOrderMapForServerVariables()},
			map[string]*foundServerVariable{
				"skipParam":  &foundServerVariable{&ServerVariable{Description: "default parameter"}, false},
				"limitParam": &foundServerVariable{&ServerVariable{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in ServerVariables",
			fields{},
			map[string]*foundServerVariable{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerVariables{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotServerVariable *ServerVariable) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotServerVariable) {
						t.Fatalf("ServerVariables.ForEach() for key = %s val = %v, want = %v", key, gotServerVariable, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("ServerVariables.ForEach() for key = %s val = %v, want = %v", key, gotServerVariable, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("ServerVariables.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("ServerVariables.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestServerVariables_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for ServerVariables fixture", fields{buildOrderMapForServerVariables()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in ServerVariables", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerVariables{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("ServerVariables.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForServerVariables() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForServerVariables() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &ServerVariable{Description: "default parameter"},
			"limitParam": &ServerVariable{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildServerVariablesFixture() ServerVariables {
	m := ServerVariables{
		data: buildOrderMapForServerVariables(),
	}

	return m
}
