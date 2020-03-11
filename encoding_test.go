package spec3

import (
	"reflect"
	"testing"
)

func TestOrderedEncodings_Get(t *testing.T) {
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
		want   *Encoding
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedEncodings()}, args{"historyMetadata"}, &Encoding{ContentType: "application/xml; charset=utf-8"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedEncodings()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedEncodings{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedEncodings.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedEncodings_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantEncoding *Encoding
		wantOK       bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedEncodings()}, args{"profileImage"}, &Encoding{ContentType: "image/png, image/jpeg"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedEncodings()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedEncodings{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantEncoding) {
				t.Errorf("OrderedEncodings.GetOK() got = %v, want %v", got, tt.wantEncoding)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedEncodings.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedEncodings_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Encoding
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedEncodings()}, args{"getParam", &Encoding{ContentType: "Getting OrderedEncodings"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedEncodings()}, args{"profileImage", &Encoding{ContentType: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedEncodings()}, args{"", &Encoding{ContentType: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedEncodings{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedEncodings.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedEncodings.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedEncodings.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedEncodings_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Encoding) error
	}
	type foundEncoding struct {
		parameter *Encoding
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundEncoding
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedEncodings fixture",
			fields{buildOrderMapForOrderedEncodings()},
			map[string]*foundEncoding{
				"historyMetadata": &foundEncoding{&Encoding{ContentType: "application/xml; charset=utf-8"}, false},
				"profileImage":    &foundEncoding{&Encoding{ContentType: "image/png, image/jpeg"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedEncodings",
			fields{},
			map[string]*foundEncoding{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedEncodings{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotEncoding *Encoding) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotEncoding) {
						t.Fatalf("OrderedEncodings.ForEach() for key = %s val = %v, want = %v", key, gotEncoding, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedEncodings.ForEach() for key = %s val = %v, want = %v", key, gotEncoding, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedEncodings.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedEncodings.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedEncodings_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedEncodings fixture", fields{buildOrderMapForOrderedEncodings()}, []string{"historyMetadata", "profileImage"}},
		{"Should return empty array when there are no values in OrderedEncodings", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedEncodings{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedEncodings.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedEncodings() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedEncodings() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"historyMetadata": &Encoding{ContentType: "application/xml; charset=utf-8"},
			"profileImage":    &Encoding{ContentType: "image/png, image/jpeg"},
		},
		keys: []string{
			"historyMetadata",
			"profileImage",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedEncodingsFixture() OrderedEncodings {
	m := OrderedEncodings{
		data: buildOrderMapForOrderedEncodings(),
	}

	return m
}
