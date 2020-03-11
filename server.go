package spec3

// Server object representing a Server.
// easyjson:json
type Server struct {
	VendorExtensible

	URL         string
	Description string
	Variables   ServerVariables
}

// ServerVariable object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	VendorExtensible

	Enum        []string
	Default     string
	Description string
}

// ServerVariables is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type ServerVariables struct {
	data OrderedMap
}

// NewServerVariables creates a new instance of ServerVariables with correct filter
func NewServerVariables() ServerVariables {
	return ServerVariables{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *ServerVariables) Get(key string) *ServerVariable {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*ServerVariable)
}

// GetOK checks if the key exists in the security requirement
func (s *ServerVariables) GetOK(key string) (*ServerVariable, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*ServerVariable)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *ServerVariables) Set(key string, val *ServerVariable) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *ServerVariables) ForEach(fn func(string, *ServerVariable) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*ServerVariable)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *ServerVariables) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *ServerVariables) Implement Marshal & Unmarshal -> JSON, YAML
