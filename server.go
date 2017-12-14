package spec3

//easyjson:json
type Server struct {
	VendorExtensible

	URL         string
	Description string
	Variables   map[string]ServerVariable
}

type ServerVariable struct {
	VendorExtensible

	Enum        []string
	Default     string
	Description string
}
