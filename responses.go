package spec3

import (
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// Responses is a container for the expected responses of an operation. The container maps a HTTP response code to the expected response.
type Responses struct {
	data responseMap
}

// Response describes a single response from an API Operation, including design-time, static links to operations based on the response.
type Response struct {
	VendorExtensible
	Reference

	Description string               `json:"description"`
	Headers     map[string]Header    `json:"headers"`
	Content     map[string]MediaType `json:"content"`
	Links       map[string]Link      `json:"links"`
}

// responseMap contains map of string with Response
type responseMap struct {
	data OrderedMap
}

func NewResponseMap() responseMap {
	return responseMap{
		data: OrderedMap{
			filter: matchResponseCode,
		},
	}
}

// Get gets the security requirement by key
func (s *responseMap) Get(key string) *Response {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Response)
}

// GetOK checks if the key exists in the security requirement
func (s *responseMap) GetOK(key string) (*Response, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Response)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *responseMap) Set(key string, val *Response) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *responseMap) ForEach(fn func(string, *Response) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Response)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *responseMap) Keys() []string {
	return s.data.Keys()
}

// MarshalJSON supports json.Marshaler interface
func (s responseMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	encodeSortedMap(&w, s.data)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (s responseMap) MarshalEasyJSON(w *jwriter.Writer) {
	encodeSortedMap(w, s.data)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (s *responseMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	decodeSortedMap(&r, &s.data)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (s *responseMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	decodeSortedMap(l, &s.data)
}

// matchResponseCode is used as filter for response codes
func matchResponseCode(key string) bool {
	if key == "default" ||
		key == "100" ||
		key == "101" ||
		key == "200" ||
		key == "201" ||
		key == "202" ||
		key == "203" ||
		key == "204" ||
		key == "205" ||
		key == "206" ||
		key == "300" ||
		key == "301" ||
		key == "302" ||
		key == "303" ||
		key == "304" ||
		key == "305" ||
		key == "307" ||
		key == "400" ||
		key == "401" ||
		key == "402" ||
		key == "403" ||
		key == "404" ||
		key == "405" ||
		key == "406" ||
		key == "407" ||
		key == "408" ||
		key == "409" ||
		key == "410" ||
		key == "411" ||
		key == "412" ||
		key == "413" ||
		key == "414" ||
		key == "415" ||
		key == "416" ||
		key == "417" ||
		key == "426" ||
		key == "500" ||
		key == "501" ||
		key == "502" ||
		key == "503" ||
		key == "504" ||
		key == "505" {
		return true
	}
	return false
}
