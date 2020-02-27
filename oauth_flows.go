package spec3

// OAuthFlows allows configuration of the supported OAuth Flows.
type OAuthFlows struct {
	VendorExtensible

	Implicit          OAuthFlow `json:"implicit,omitempty"`
	Password          OAuthFlow `json:"password,omitempty"`
	ClientCredentials OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitempty"`
}

// OAuthFlow configuration details for a supported OAuth Flow
type OAuthFlow struct {
	VendorExtensible

	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}
