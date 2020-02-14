package spec3

// Contact contains information for the exposed API.
//easyjson:json
type Contact struct {
	VendorExtensible

	Name  string
	URL   string
	Email string
}
