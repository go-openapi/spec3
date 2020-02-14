package spec3

// ExternalDocumentation allows referencing an external resource for extended documentation.
// easyjson:json
type ExternalDocumentation struct {
	VendorExtensible

	Description string
	URL         string
}
