package spec3

// Parameter describes a single operation parameter.
// A unique parameter is defined by a combination of a name and location.
type Parameter struct {
	VendorExtensible
	Reference

	Name            string               `json:"name"`
	In              string               `json:"in"`
	Description     string               `json:"description"`
	Required        bool                 `json:"required"`
	Deprecated      bool                 `json:"deprecated"`
	AllowEmptyValue bool                 `json:"allowEmptyValue"`
	Style           string               `json:"style"`
	Explode         bool                 `json:"explode"`
	AllowReserved   bool                 `json:"allowReserved"`
	Schema          Schema               `json:"schema"`
	Example         interface{}          `json:"example"`
	Examples        map[string]Example   `json:"examples"`
	Contents        map[string]MediaType `json:"contents"`
}

type Parameters struct {
	data OrderedMap
}

func NewParameters() Parameters {
	return Parameters{
		data: OrderedMap{
			filter: MatchNonEmptyKeys,
		},
	}
}

// Get gets the security requirement by key
func (s *Parameters) Get(key string) *Parameter {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*Parameter)
}

// GetOK checks if the key exists in the security requirement
func (s *Parameters) GetOK(key string) (*Parameter, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*Parameter)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *Parameters) Set(key string, val *Parameter) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *Parameters) ForEach(fn func(string, *Parameter) error) error {
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
func (s *Parameters) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *Parameters) Implement Marshal & Unmarshal -> JSON, YAML
