package spec3

// Header follows the structure of the Parameter Object
type Header struct {
	Parameter
}

// OrderedHeaders is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedHeaders struct {
	data OrderedMap
}

// NewOrderedHeaders creates a new instance of OrderedHeaders with correct filter
func NewOrderedHeaders() OrderedHeaders {
	return OrderedHeaders{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedHeaders) Get(key string) *Header {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Header)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedHeaders) GetOK(key string) (*Header, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Header)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedHeaders) Set(key string, val *Header) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedHeaders) ForEach(fn func(string, *Header) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Header)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedHeaders) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedHeaders) Implement Marshal & Unmarshal -> JSON, YAML
