package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedLinks_Get(t *testing.T) {
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
		want   *Link
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedLinks()}, args{"skipParam"}, &Link{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedLinks()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedLinks{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedLinks.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedLinks_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLink *Link
		wantOK   bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedLinks()}, args{"limitParam"}, &Link{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedLinks()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedLinks{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantLink) {
				t.Errorf("OrderedLinks.GetOK() got = %v, want %v", got, tt.wantLink)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedLinks.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedLinks_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Link
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedLinks()}, args{"getParam", &Link{Description: "Getting OrderedLinks"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedLinks()}, args{"limitParam", &Link{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedLinks()}, args{"", &Link{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedLinks{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedLinks.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedLinks.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedLinks.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedLinks_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Link) error
	}
	type foundLink struct {
		parameter *Link
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundLink
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedLinks fixture",
			fields{buildOrderMapForOrderedLinks()},
			map[string]*foundLink{
				"skipParam":  &foundLink{&Link{Description: "default parameter"}, false},
				"limitParam": &foundLink{&Link{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedLinks",
			fields{},
			map[string]*foundLink{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedLinks{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotLink *Link) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotLink) {
						t.Fatalf("OrderedLinks.ForEach() for key = %s val = %v, want = %v", key, gotLink, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedLinks.ForEach() for key = %s val = %v, want = %v", key, gotLink, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedLinks.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedLinks.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedLinks_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedLinks fixture", fields{buildOrderMapForOrderedLinks()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedLinks", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedLinks{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedLinks.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedLinks() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedLinks() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &Link{Description: "default parameter"},
			"limitParam": &Link{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedLinksFixture() OrderedLinks {
	m := OrderedLinks{
		data: buildOrderMapForOrderedLinks(),
	}

	return m
}
