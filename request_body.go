package spec3

// RequestBody describes a single request body.
type RequestBody struct {
	VendorExtensible
	Reference

	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	Required    string `json:"required,omitempty"`
}

// OrderedRequestBodies is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedRequestBodies struct {
	data OrderedMap
}

// NewOrderedRequestBodies creates a new instance of OrderedRequestBodies with correct filter
func NewOrderedRequestBodies() OrderedRequestBodies {
	return OrderedRequestBodies{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedRequestBodies) Get(key string) *RequestBody {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*RequestBody)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedRequestBodies) GetOK(key string) (*RequestBody, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*RequestBody)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedRequestBodies) Set(key string, val *RequestBody) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedRequestBodies) ForEach(fn func(string, *RequestBody) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*RequestBody)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedRequestBodies) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedRequestBodies) Implement Marshal & Unmarshal -> JSON, YAML
