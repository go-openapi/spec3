package spec3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// MapEntry represents a key value pair
type MapEntry struct {
	Key   string
	Value interface{}
}

// Filter for deciding which keys make it into the sorted map
type Filter func(string) bool

// MatchAll keys, is used as default filter
func MatchAll(_ string) bool { return true }

// MatchNonEmptyKeys keys, is used to allow only non empty strings
func MatchNonEmptyKeys(key string) bool { return key != "" }

// MatchExtension is used as filter for vendor extensions
func MatchExtension(key string) bool { return strings.HasPrefix(key, "x-") }

// Normalizer is used to normalize keys when writing to a map
type Normalizer func(string) string

// LowerCaseKeys lowercases keys when looking up in the map
var LowerCaseKeys = strings.ToLower

// NOPNormalizer passes the key through, used as default
func NOPNormalizer(s string) string { return s }

// OrderedMap is a map that preserves insertion order
type OrderedMap struct {
	filter    Filter
	normalize Normalizer
	data      map[string]interface{}
	keys      []string
}

// Len of the known keys
func (s *OrderedMap) Len() int {
	return len(s.keys)
}

// GetOK get a value for the specified key, the boolean result indicates if the value was found or not
func (s *OrderedMap) GetOK(key string) (interface{}, bool) {
	v, ok := s.data[s.normalizeKey(key)]
	return v, ok
}

// Get get a value for the specified key
func (s *OrderedMap) Get(key string) interface{} {
	val, ok := s.data[s.normalizeKey(key)]
	if !ok {
		return nil
	}
	return val
}

func (s *OrderedMap) normalizeKey(key string) string {
	if s.normalize == nil {
		s.normalize = NOPNormalizer
	}

	return s.normalize(key)
}

func (s *OrderedMap) allows(key string) bool {
	if s.filter == nil {
		s.filter = MatchAll
	}
	return s.filter(key)
}

// Set a value in the map
func (s *OrderedMap) Set(key string, value interface{}) bool {
	key = s.normalizeKey(key)
	if !s.allows(key) {
		return false
	}

	if s.data == nil {
		s.data = make(map[string]interface{})
	}
	_, ok := s.data[key]
	s.data[key] = value
	if !ok {
		s.keys = append(s.keys, key)
	}
	return !ok
}

// Delete a value from the map
func (s *OrderedMap) Delete(k string) bool {
	key := s.normalizeKey(k)
	if !s.allows(key) {
		return false
	}

	_, ok := s.data[key]
	if !ok {
		return false
	}

	delete(s.data, key)
	for i, k := range s.keys {
		if k == key {
			s.keys = append(s.keys[:i], s.keys[i+1:]...)
		}
	}

	if len(s.keys) == 0 {
		s.data = nil
		s.keys = nil
	}
	return ok
}

// Keys in the order of addition to the map
func (s *OrderedMap) Keys() []string {
	return s.keys[:]
}

// ForEach executes a function for each value in the map
func (s *OrderedMap) ForEach(fn func(string, interface{}) error) error {
	for _, k := range s.Keys() {
		val := s.Get(k)
		if err := fn(k, val); err != nil {
			return err
		}
	}
	return nil
}

// Values in the order of addition to the map
func (s *OrderedMap) Values() []interface{} {
	values := make([]interface{}, len(s.keys))
	for i, k := range s.keys {
		values[i] = s.data[k]
	}
	return values
}

// Entries in the order of addition to the map
func (s *OrderedMap) Entries() []MapEntry {
	values := make([]MapEntry, len(s.keys))
	for i, k := range s.keys {
		values[i] = MapEntry{Key: k, Value: s.data[k]}
	}
	return values
}

func (s OrderedMap) String() string {
	if s.data == nil {
		return ""
	}

	var b bytes.Buffer
	b.WriteByte('{')
	b.WriteByte(' ')
	first := true
	for _, k := range s.keys {
		if !first {
			b.WriteRune(',')
			b.WriteRune(' ')
		}
		first = false
		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(fmt.Sprintf("%#v", s.data[k]))
	}
	if !first {
		b.WriteByte(' ')
	}
	b.WriteByte('}')
	return b.String()
}

// MarshalJSON supports json.Marshaler interface
func (s OrderedMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	encodeSortedMap(&w, s)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (s OrderedMap) MarshalEasyJSON(w *jwriter.Writer) {
	encodeSortedMap(w, s)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (s *OrderedMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	decodeSortedMap(&r, s)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (s *OrderedMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	decodeSortedMap(l, s)
}

func encodeSortedMap(out *jwriter.Writer, in OrderedMap) {
	if in.data == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
		return
	}

	out.RawByte('{')
	first := true
	for _, k := range in.keys {
		_ = first
		if !first {
			out.RawByte(',')
		}
		first = false
		out.String(k)
		out.RawByte(':')
		value := in.data[k]
		if m, ok := value.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := value.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(value))
		}
	}

	out.RawByte('}')
}

func decodeSortedMap(in *jlexer.Lexer, out *OrderedMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := string(in.String())
		in.WantColon()
		out.Set(key, in.Interface())
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

// OrderedStrings is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedStrings struct {
	data OrderedMap
}

// NewOrderedStrings creates a new instance of OrderedStrings with correct filter
func NewOrderedStrings() OrderedStrings {
	return OrderedStrings{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedStrings) Get(key string) *string {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*string)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedStrings) GetOK(key string) (*string, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*string)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedStrings) Set(key string, val *string) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedStrings) ForEach(fn func(string, *string) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*string)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedStrings) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedStrings) Implement Marshal & Unmarshal -> JSON, YAML
