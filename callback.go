package spec3

// Callback A map of possible out-of band callbacks related to the parent operation.
// Each value in the map is a Path Item Object that describes a set of requests that may be initiated by the API provider and the expected responses.
// The key value used to identify the callback object is an expression, evaluated at runtime, that identifies a URL to use for the callback operation.
type Callback struct {
	data OrderedMap
}

// NewCallback creates a new instance of Callback with correct filter
func NewCallback() Callback {
	return Callback{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *Callback) Get(key string) *PathItem {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*PathItem)
}

// GetOK checks if the key exists in the security requirement
func (s *Callback) GetOK(key string) (*PathItem, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*PathItem)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *Callback) Set(key string, val *PathItem) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *Callback) ForEach(fn func(string, *PathItem) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*PathItem)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *Callback) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *Callback) Implement Marshal & Unmarshal -> JSON, YAML

// OrderedCallbacks is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedCallbacks struct {
	data OrderedMap
}

// NewOrderedCallbacks creates a new instance of OrderedCallbacks with correct filter
func NewOrderedCallbacks() OrderedCallbacks {
	return OrderedCallbacks{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedCallbacks) Get(key string) *Callback {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Callback)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedCallbacks) GetOK(key string) (*Callback, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Callback)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedCallbacks) Set(key string, val *Callback) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedCallbacks) ForEach(fn func(string, *Callback) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Callback)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedCallbacks) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedCallbacks) Implement Marshal & Unmarshal -> JSON, YAML
