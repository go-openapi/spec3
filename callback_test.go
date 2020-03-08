package spec3

import (
	"reflect"
	"testing"
)

func TestCallback_Get(t *testing.T) {
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
		want   *PathItem
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForCallback()}, args{"skipParam"}, &PathItem{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForCallback()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Callback{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Callback.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCallback_GetOK(t *testing.T) {
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
		wantPathItem *PathItem
		wantOK       bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForCallback()}, args{"limitParam"}, &PathItem{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForCallback()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Callback{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantPathItem) {
				t.Errorf("Callback.GetOK() got = %v, want %v", got, tt.wantPathItem)
			}
			if got1 != tt.wantOK {
				t.Errorf("Callback.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestCallback_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *PathItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForCallback()}, args{"getParam", &PathItem{Description: "Getting Callback"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForCallback()}, args{"limitParam", &PathItem{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForCallback()}, args{"", &PathItem{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Callback{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("Callback.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("Callback.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("Callback.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestCallback_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *PathItem) error
	}
	type foundPathItem struct {
		parameter *PathItem
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundPathItem
		wantErr          error
	}{
		{
			"Should iterate 4 items for Callback fixture",
			fields{buildOrderMapForCallback()},
			map[string]*foundPathItem{
				"skipParam":  &foundPathItem{&PathItem{Description: "default parameter"}, false},
				"limitParam": &foundPathItem{&PathItem{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in Callback",
			fields{},
			map[string]*foundPathItem{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Callback{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotPathItem *PathItem) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotPathItem) {
						t.Fatalf("Callback.ForEach() for key = %s val = %v, want = %v", key, gotPathItem, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("Callback.ForEach() for key = %s val = %v, want = %v", key, gotPathItem, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("Callback.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("Callback.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestCallback_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for Callback fixture", fields{buildOrderMapForCallback()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in Callback", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Callback{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("Callback.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForCallback() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForCallback() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &PathItem{Description: "default parameter"},
			"limitParam": &PathItem{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildCallbackFixture() Callback {
	m := Callback{
		data: buildOrderMapForCallback(),
	}

	return m
}

// Test cases for the OrderedCallbacks

/*

// FIXME: fix the test cases

func TestOrderedCallbacks_Get(t *testing.T) {
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
		want   *Callback
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"skipParam"}, &Callback{Description: "default parameter"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedCallbacks{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedCallbacks.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedCallbacks_GetOK(t *testing.T) {
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
		wantCallback *Callback
		wantOK       bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"limitParam"}, &Callback{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedCallbacks{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantCallback) {
				t.Errorf("OrderedCallbacks.GetOK() got = %v, want %v", got, tt.wantCallback)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedCallbacks.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedCallbacks_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Callback
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"getParam", &Callback{Description: "Getting OrderedCallbacks"}}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"limitParam", &Callback{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedCallbacks()}, args{"", &Callback{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedCallbacks{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedCallbacks.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedCallbacks.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedCallbacks.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedCallbacks_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Callback) error
	}
	type foundCallback struct {
		parameter *Callback
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundCallback
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedCallbacks fixture",
			fields{buildOrderMapForOrderedCallbacks()},
			map[string]*foundCallback{
				"skipParam":  &foundCallback{&Callback{Description: "default parameter"}, false},
				"limitParam": &foundCallback{&Callback{Description: "OK"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedCallbacks",
			fields{},
			map[string]*foundCallback{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedCallbacks{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotCallback *Callback) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotCallback) {
						t.Fatalf("OrderedCallbacks.ForEach() for key = %s val = %v, want = %v", key, gotCallback, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedCallbacks.ForEach() for key = %s val = %v, want = %v", key, gotCallback, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedCallbacks.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedCallbacks.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedCallbacks_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedCallbacks fixture", fields{buildOrderMapForOrderedCallbacks()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedCallbacks", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedCallbacks{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedCallbacks.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedCallbacks() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedCallbacks() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  &Callback{Description: "default parameter"},
			"limitParam": &Callback{Description: "OK"},
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedCallbacksFixture() OrderedCallbacks {
	m := OrderedCallbacks{
		data: buildOrderMapForOrderedCallbacks(),
	}

	return m
}
*/
