package spec3

// SecurityScheme defines a security scheme that can be used by the operations.
// Supported schemes are HTTP authentication, an API key (either as a header, a cookie parameter or as a query parameter), OAuth2's common flows (implicit, password, application and access code) as defined in RFC6749, and OpenID Connect Discovery.
type SecurityScheme struct {
	VendorExtensible
	Reference

	Type             string    `json:"type,omitempty"`
	Description      string    `json:"description,omitempty"`
	Name             string    `json:"name,omitempty"`
	In               string    `json:"in,omitempty"`
	Scheme           string    `json:"scheme,omitempty"`
	BearerFormat     string    `json:"bearerFormat,omitempty"`
	Flows            OAuthFlow `json:"flows,omitempty"`
	OpenIDConnectURL string    `json:"openIdConnectUrl,omitempty"`
}

// OrderedSecuritySchemes is a map between a variable name and its value. The value is used for substitution in the server's URL template.
type OrderedSecuritySchemes struct {
	data OrderedMap
}

// NewOrderedSecuritySchemes creates a new instance of OrderedSecuritySchemes with correct filter
func NewOrderedSecuritySchemes() OrderedSecuritySchemes {
	return OrderedSecuritySchemes{
		data: OrderedMap{
			filter: MatchNonEmptyKeys, // TODO: check if keys are some regex or just any non empty string
		},
	}
}

// Get gets the security requirement by key
func (s *OrderedSecuritySchemes) Get(key string) *SecurityScheme {
	v := s.data.Get(key)
	if v == nil {
		return nil
	}
	return v.(*SecurityScheme)
}

// GetOK checks if the key exists in the security requirement
func (s *OrderedSecuritySchemes) GetOK(key string) (*SecurityScheme, bool) {
	v, ok := s.data.GetOK(key)
	if !ok {
		return nil, ok
	}

	sr, ok := v.(*SecurityScheme)
	return sr, ok
}

// Set sets the value to the security requirement
func (s *OrderedSecuritySchemes) Set(key string, val *SecurityScheme) bool {
	return s.data.Set(key, val)
}

// ForEach executes the function for each security requirement
func (s *OrderedSecuritySchemes) ForEach(fn func(string, *SecurityScheme) error) error {
	s.data.ForEach(func(key string, val interface{}) error {
		response, _ := val.(*SecurityScheme)
		if err := fn(key, response); err != nil {
			return err
		}
		return nil
	})
	return nil
}

// Keys gets the list of keys
func (s *OrderedSecuritySchemes) Keys() []string {
	return s.data.Keys()
}

// TODO: (s *OrderedSecuritySchemes) Implement Marshal & Unmarshal -> JSON, YAML
