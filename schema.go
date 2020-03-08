package spec3

// Schema object allows the definition of input and output data types. These types can be objects, but also primitives and arrays. This object is an extended subset of the JSON Schema Specification Wright Draft 00.
type Schema struct {
	VendorExtensible
	Reference
}

// OrderedSchemas is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedSchemas struct {
	data OrderedMap
}

// NewOrderedSchemas creates a new instance of OrderedSchemas with correct filter
func NewOrderedSchemas() OrderedSchemas {
	return OrderedSchemas{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedSchemas) Get(key string) *Schema {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Schema)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedSchemas) GetOK(key string) (*Schema, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Schema)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedSchemas) Set(key string, val *Schema) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedSchemas) ForEach(fn func(string, *Schema) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Schema)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedSchemas) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedSchemas) Implement Marshal & Unmarshal -> JSON, YAML
