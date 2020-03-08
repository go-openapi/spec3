package spec3

// Link represents a possible design-time link for a response.
// The presence of a link does not guarantee the caller's ability to successfully invoke it, rather it provides a known relationship and traversal mechanism between responses and other operations.
type Link struct {
	VendorExtensible
	Reference

	OperationRef string                 `json:"operationRef,omitempty"`
	OperationID  string                 `json:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Server       Server                 `json:"server,omitempty"`
}

// OrderedLinks is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedLinks struct {
	data OrderedMap
}

// NewOrderedLinks creates a new instance of OrderedLinks with correct filter
func NewOrderedLinks() OrderedLinks {
	return OrderedLinks{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedLinks) Get(key string) *Link {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Link)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedLinks) GetOK(key string) (*Link, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Link)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedLinks) Set(key string, val *Link) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedLinks) ForEach(fn func(string, *Link) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Link)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedLinks) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedLinks) Implement Marshal & Unmarshal -> JSON, YAML
