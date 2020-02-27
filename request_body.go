package spec3

// RequestBody describes a single request body.
type RequestBody struct {
	VendorExtensible
	Reference

	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	Required    string `json:"required,omitempty"`
}
