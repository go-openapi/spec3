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

type Filter func(string) bool

func MatchAll(_ string) bool         { return true }
func MatchExtension(key string) bool { return strings.HasPrefix(key, "x-") }

type Normalizer func(string) string

var LowerCaseKeys = strings.ToLower

// SortedMap is a map that preserves insertion order
type SortedMap struct {
	data map[string]interface{}
	keys []string
}

// Len of the known keys
func (s *SortedMap) Len() int {
	return len(s.keys)
}

// GetOK get a value for the specified key, the boolean result indicates if the value was found or not
func (s *SortedMap) GetOK(key string) (interface{}, bool) {
	v, ok := s.data[key]
	return v, ok
}

// Get get a value for the specified key
func (s *SortedMap) Get(key string) interface{} {
	return s.data[key]
}

// Set a value in the map
func (s *SortedMap) Set(key string, value interface{}) bool {
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
func (s *SortedMap) Delete(key string) bool {
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
func (s *SortedMap) Keys() []string {
	return s.keys[:]
}

// Values in the order of addition to the map
func (s *SortedMap) Values() []interface{} {
	values := make([]interface{}, len(s.keys))
	for i, k := range s.keys {
		values[i] = s.data[k]
	}
	return values
}

// Entries in the order of addition to the map
func (s *SortedMap) Entries() []MapEntry {
	values := make([]MapEntry, len(s.keys))
	for i, k := range s.keys {
		values[i] = MapEntry{Key: k, Value: s.data[k]}
	}
	return values
}

func (s SortedMap) String() string {
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
func (s SortedMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	encodeSortedMap(&w, s)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (s SortedMap) MarshalEasyJSON(w *jwriter.Writer) {
	encodeSortedMap(w, s)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (s *SortedMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	decodeSortedMap(&r, s)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (s *SortedMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	decodeSortedMap(l, s)
}

func encodeSortedMap(out *jwriter.Writer, in SortedMap) {
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

func decodeSortedMap(in *jlexer.Lexer, out *SortedMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	if !in.IsDelim('}') {
		out.data = make(map[string]interface{})
	} else {
		out.data = nil
	}
	for !in.IsDelim('}') {
		key := string(in.String())
		in.WantColon()
		out.data[key] = in.Interface()
		out.keys = append(out.keys, key)
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
