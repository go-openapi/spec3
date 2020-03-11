package spec3

// Example
type Example struct {
	VendorExtensible
	Reference

	Summary       string      `json:"summary"`
	Description   string      `json:"description"`
	Value         interface{} `json:"value"`
	ExternalValue string      `json:"externalValue"`
}

// OrderedExamples is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedExamples struct {
	data OrderedMap
}

// NewOrderedExamples creates a new instance of OrderedExamples with correct filter
func NewOrderedExamples() OrderedExamples {
	return OrderedExamples{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedExamples) Get(key string) *Example {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Example)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedExamples) GetOK(key string) (*Example, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Example)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedExamples) Set(key string, val *Example) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedExamples) ForEach(fn func(string, *Example) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Example)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedExamples) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedExamples) Implement Marshal & Unmarshal -> JSON, YAML
