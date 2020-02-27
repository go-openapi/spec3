package spec3

// Server object representing a Server.
// easyjson:json
type Server struct {
	VendorExtensible

	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerVariable object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	VendorExtensible

	Enum        []string
	Default     string
	Description string
}
