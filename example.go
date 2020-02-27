package spec3

// Example
type Example struct {
	VendorExtensible
	Reference

	Summary       string      `json:"summary"`
	Description   string      `json:"description"`
	Value         interface{} `json:"value"`
	ExternalValue string      `json:"externalValue"`
}
