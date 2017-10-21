package spec3

import (
	"encoding/json"
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
