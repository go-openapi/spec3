package spec3

// XML a metadata object that allows for more fine-tuned XML model definitions.
type XML struct {
	VendorExtensible

	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Prefix    string `json:"prefix"`
	Attribute bool   `json:"attribute"`
	Wrapped   bool   `json:"wrapped"`
}
