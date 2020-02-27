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
