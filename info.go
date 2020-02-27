package spec3

// Info provides metadata about the API.
// The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
// easyjson:json
type Info struct {
	VendorExtensible

	Title          string
	Description    string
	TermsOfService string
	Contact        *Contact
	License        *License
	Version        string
}
