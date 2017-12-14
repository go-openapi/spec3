package spec3

//easyjson:json
type Info struct {
	VendorExtensible

	Title          string
	Description    string
	TermsOfService string
	Contact        *Contact
	License        *License
	Version        string
}
