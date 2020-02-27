package spec3

// Tag adds metadata to a single tag that is used by the Operation Object.
// It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
// easyjson:json
type Tag struct {
	VendorExtensible

	Name         string
	Description  string
	ExternalDocs *ExternalDocumentation
}
