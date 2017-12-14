package spec3

type OpenAPI struct {
	VendorExtensible

	OpenAPI      string                  `json:"openapi"`
	Info         *Info                   `json:"info,omitempty"`
	Servers      []Server                `json:"servers,omitempty"`
	Paths        *PathItem               `json:"paths"`
	Security     []SecurityRequirement   `json:"security,omitempty"`
	Tags         []Tag                   `json:"tags,omitempty"`
	ExternalDocs []ExternalDocumentation `json:"externalDocs,omitempty"`
}
