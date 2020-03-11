package spec3

// Encoding definition applied to a single schema property.
type Encoding struct {
	VendorExtensible

	ContentType   string         `json:"contentType,omitempty"`
	Headers       OrderedHeaders `json:"headers,omitempty"`
	Style         string         `json:"style,omitempty"`
	Explode       bool           `json:"explode,omitempty"`
	AllowReserved bool           `json:"allowReserved,omitempty"`
}

// OrderedEncodings is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedEncodings struct {
	data OrderedMap
}

// NewOrderedEncodings creates a new instance of OrderedEncodings with correct filter
func NewOrderedEncodings() OrderedEncodings {
	return OrderedEncodings{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedEncodings) Get(key string) *Encoding {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Encoding)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedEncodings) GetOK(key string) (*Encoding, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Encoding)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedEncodings) Set(key string, val *Encoding) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedEncodings) ForEach(fn func(string, *Encoding) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Encoding)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedEncodings) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedEncodings) Implement Marshal & Unmarshal -> JSON, YAML
