package spec3

// OpenAPI is the root document object of the OpenAPI document
type OpenAPI struct {
	VendorExtensible

	Components   *Components             `json:"components,omitempty"`
	ExternalDocs []ExternalDocumentation `json:"externalDocs,omitempty"`
	Info         *Info                   `json:"info,omitempty"`
	OpenAPI      string                  `json:"openapi"`
	Paths        *PathItem               `json:"paths"`
	Security     []SecurityRequirement   `json:"security,omitempty"`
	Servers      []Server                `json:"servers,omitempty"`
	Tags         []Tag                   `json:"tags,omitempty"`
}
