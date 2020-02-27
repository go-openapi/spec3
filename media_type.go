package spec3

// MediaType provides schema and examples for the media type identified by its key.
type MediaType struct {
	VendorExtensible

	Schema   Schema              `json:"schema,omitempty"`
	Example  interface{}         `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}
