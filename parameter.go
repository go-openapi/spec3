package spec3

// Parameter describes a single operation parameter.
// A unique parameter is defined by a combination of a name and location.
type Parameter struct {
	VendorExtensible
	Reference

	Name            string            `json:"name"`
	In              string            `json:"in"`
	Description     string            `json:"description"`
	Required        bool              `json:"required"`
	Deprecated      bool              `json:"deprecated"`
	AllowEmptyValue bool              `json:"allowEmptyValue"`
	Style           string            `json:"style"`
	Explode         bool              `json:"explode"`
	AllowReserved   bool              `json:"allowReserved"`
	Schema          Schema            `json:"schema"`
	Example         interface{}       `json:"example"`
	Examples        OrderedExamples   `json:"examples"`
	Contents        OrderedMediaTypes `json:"contents"`
}

type OrderedParameters struct {
	data OrderedMap
}

func NewOrderedParameters() OrderedParameters {
	return OrderedParameters{
		data: OrderedMap{
			filter: MatchNonEmptyKeys,
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedParameters) Get(key string) *Parameter {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Parameter)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedParameters) GetOK(key string) (*Parameter, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Parameter)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedParameters) Set(key string, val *Parameter) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedParameters) ForEach(fn func(string, *Parameter) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*Parameter)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedParameters) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedParameters) Implement Marshal & Unmarshal -> JSON, YAML
