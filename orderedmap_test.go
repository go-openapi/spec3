package spec3

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/brianvoe/gofakeit"
)

func randomData(ln int) (OrderedMap, []string) {
	d := OrderedMap{
		data: make(map[string]interface{}, ln),
	}
	var keys []string
	known := make(map[string]struct{}, ln)
	for i := 0; i < ln; i++ {
		key := randomWord(known)
		d.Set(key, key+"Value")
		keys = append(keys, key)
	}

	return d, keys
}

func randomWord(known map[string]struct{}) string {
	ok := true
	var key string
	for ok {
		key = gofakeit.Password(true, true, true, false, false, 15)
		_, ok = known[key]
	}
	known[key] = struct{}{}
	return key
}

func TestSortedMap_Keys(t *testing.T) {
	data, keys := randomData(100)
	require.Len(t, keys, 100)
	for i, key := range data.Keys() {
		require.Equal(t, keys[i], key)
	}

	// check that sort is not lexical
	sort.Strings(keys)
	var cnt int
	for i, key := range data.Keys() {
		if keys[i] == key {
			cnt++
		}
	}
	assert.NotEqual(t, cnt, 100)
}

func TestSortedMap_Values(t *testing.T) {
	data, keys := randomData(100)

	for i, key := range data.Values() {
		require.Equal(t, keys[i]+"Value", key)
	}
}

func TestSortedMap_Entries(t *testing.T) {
	data, keys := randomData(100)

	for i, key := range data.Entries() {
		require.Equal(t, keys[i], key.Key)
		require.Equal(t, keys[i]+"Value", key.Value)
	}
}

func TestSortedMap_GetSet(t *testing.T) {
	var data OrderedMap
	data.Set("first", "value1")
	data.Set("second", "value2")
	data.Set("first", "value3")
	assert.Equal(t, 2, data.Len())
	assert.Equal(t, "value3", data.Get("first"))
	assert.Equal(t, "value2", data.Get("second"))

	v, ok := data.GetOK("second")
	assert.Equal(t, "value2", v)
	assert.True(t, ok)

	_, ok = data.GetOK("notthere")
	assert.False(t, ok)
}

func TestSortedMap_Delete(t *testing.T) {
	var data OrderedMap
	data.Set("first", "value1")
	data.Set("second", "value2")
	data.Set("third", "value3")
	assert.Equal(t, 3, data.Len())

	data.Delete("first")
	assert.Equal(t, 2, data.Len())
	data.Delete("second")
	assert.Equal(t, 1, data.Len())
	data.Delete("third")
	assert.Equal(t, 0, data.Len())

	v, ok := data.GetOK("third")
	assert.False(t, ok)
	assert.Nil(t, v)
	assert.Nil(t, data.Get("third"))
}

func TestSortedMap_String(t *testing.T) {
	var data OrderedMap
	data.Set("first", "value1")
	data.Set("second", "value2")
	data.Set("third", "value3")
	assert.Equal(t, `{ first: "value1", second: "value2", third: "value3" }`, data.String())

	var d OrderedMap
	assert.Equal(t, "", d.String())

	var d1 OrderedMap
	d1.Set("first", "value1")
	assert.Equal(t, `{ first: "value1" }`, d1.String())
}

func TestSortedMap_ToJSON(t *testing.T) {
	var nested OrderedMap
	nested.Set("akey", "value")
	nested.Set("other", "some")

	var data OrderedMap
	data.Set("first", "value1")
	data.Set("second", 2)
	data.Set("third", nested)

	b, err := json.Marshal(data)
	if assert.NoError(t, err) {
		assert.Equal(t, `{"first":"value1","second":2,"third":{"akey":"value","other":"some"}}`, string(b))
	}
}

func TestSortedMap_FromJSON(t *testing.T) {
	var data OrderedMap

	js := `{"first":"value1","second":2,"third":{"akey":"value","other":"some"}}`
	err := json.Unmarshal([]byte(js), &data)
	if assert.NoError(t, err) {
		assert.Equal(t, "value1", data.Get("first"))
		assert.Equal(t, float64(2), data.Get("second"))
		assert.Equal(t, map[string]interface{}{"akey": "value", "other": "some"}, data.Get("third"))
	}
}

func TestOrderedStrings_Get(t *testing.T) {
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
		want   *string
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedStrings()}, args{"skipParam"}, ToPtrString("default parameter")},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedStrings()}, args{"getParam"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedStrings{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedStrings.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedStrings_GetOK(t *testing.T) {
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
		wantString *string
		wantOK     bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForOrderedStrings()}, args{"limitParam"}, ToPtrString("OK"), true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForOrderedStrings()}, args{"getParam"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedStrings{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantString) {
				t.Errorf("OrderedStrings.GetOK() got = %v, want %v", got, tt.wantString)
			}
			if got1 != tt.wantOK {
				t.Errorf("OrderedStrings.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestOrderedStrings_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent parameter code is passed", fields{buildOrderMapForOrderedStrings()}, args{"getParam", ToPtrString("Getting OrderedStrings")}, true},
		{"Should fail when existent parameter code is passed", fields{buildOrderMapForOrderedStrings()}, args{"limitParam", ToPtrString("another OK")}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForOrderedStrings()}, args{"", ToPtrString("description of item #empty")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedStrings{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("OrderedStrings.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("OrderedStrings.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("OrderedStrings.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestOrderedStrings_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *string) error
	}
	type foundString struct {
		parameter *string
		found     bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundString
		wantErr          error
	}{
		{
			"Should iterate 4 items for OrderedStrings fixture",
			fields{buildOrderMapForOrderedStrings()},
			map[string]*foundString{
				"skipParam":  &foundString{ToPtrString("default parameter"), false},
				"limitParam": &foundString{ToPtrString("OK"), false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in OrderedStrings",
			fields{},
			map[string]*foundString{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedStrings{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotString *string) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.parameter, gotString) {
						t.Fatalf("OrderedStrings.ForEach() for key = %s val = %v, want = %v", key, gotString, wantVal.parameter)
					}
					wantVal.found = true
				} else {
					t.Fatalf("OrderedStrings.ForEach() for key = %s val = %v, want = %v", key, gotString, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("OrderedStrings.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("OrderedStrings.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestOrderedStrings_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 2 items for OrderedStrings fixture", fields{buildOrderMapForOrderedStrings()}, []string{"skipParam", "limitParam"}},
		{"Should return empty array when there are no values in OrderedStrings", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderedStrings{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("OrderedStrings.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func buildEmptyOrderMapForOrderedStrings() OrderedMap {
	return OrderedMap{
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderMapForOrderedStrings() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"skipParam":  ToPtrString("default parameter"),
			"limitParam": ToPtrString("OK"),
		},
		keys: []string{
			"skipParam",
			"limitParam",
		},
		filter: MatchNonEmptyKeys,
	}
}

func buildOrderedStringsFixture() OrderedStrings {
	m := OrderedStrings{
		data: buildOrderMapForOrderedStrings(),
	}

	return m
}

func ToPtrString(val string) *string {
	return &val
}
