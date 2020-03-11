package spec3

// MediaType provides schema and examples for the media type identified by its key.
type MediaType struct {
	VendorExtensible

	Schema   Schema           `json:"schema,omitempty"`
	Example  interface{}      `json:"example,omitempty"`
	Examples OrderedExamples  `json:"examples,omitempty"`
	Encoding OrderedEncodings `json:"encoding,omitempty"`
}

// OrderedMediaTypes is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedMediaTypes struct {
	data OrderedMap
}

// NewOrderedMediaTypes creates a new instance of OrderedMediaTypes with correct filter
func NewOrderedMediaTypes() OrderedMediaTypes {
	return OrderedMediaTypes{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedMediaTypes) Get(key string) *MediaType {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*MediaType)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedMediaTypes) GetOK(key string) (*MediaType, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*MediaType)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedMediaTypes) Set(key string, val *MediaType) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedMediaTypes) ForEach(fn func(string, *MediaType) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*MediaType)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedMediaTypes) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedMediaTypes) Implement Marshal & Unmarshal -> JSON, YAML
