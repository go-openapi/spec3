package spec3

//easyjson:json
type Tag struct {
	VendorExtensible

	Name         string
	Description  string
	ExternalDocs *ExternalDocumentation
}
