package spec3

import (
	"fmt"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

type SecurityRequirement struct {
	data OrderedMap
}

func (s *SecurityRequirement) Get(key string) []string {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil
	}
	sr, ok := v.([]string)
	if !ok {
		return nil
	}
	return sr
}

func (s *SecurityRequirement) GetOK(key string) ([]string, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.([]string)
	return sr, ok
}

func (s *SecurityRequirement) Set(key string, scopes ...string) bool {
	return s.data.Set(key, scopes)
}

func (s *SecurityRequirement) ForEach(fn func(string, []string) error) error {
	for _, k := range s.data.Keys() {
		scopes, ok := s.data.Get(k).([]string)
		if !ok {
			return fmt.Errorf("security requirement scopes not a []string but %T", s.data.Get(k))
		}
		if err := fn(k, scopes); err != nil {
			return err
		}
	}
	return nil
}

func (s *SecurityRequirement) Keys() []string {
	return s.data.Keys()
}

// MarshalJSON supports json.Marshaler interface
func (s SecurityRequirement) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	encodeSortedMap(&w, s.data)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (s SecurityRequirement) MarshalEasyJSON(w *jwriter.Writer) {
	encodeSortedMap(w, s.data)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (s *SecurityRequirement) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	decodeSortedMap(&r, &s.data)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (s *SecurityRequirement) UnmarshalEasyJSON(l *jlexer.Lexer) {
	decodeSortedMap(l, &s.data)
}
