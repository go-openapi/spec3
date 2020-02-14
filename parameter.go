package spec3

// Parameter describes a single operation parameter.
// A unique parameter is defined by a combination of a name and location.
type Parameter struct {
	VendorExtensible
	Reference

	Name            string               `json:"name"`
	In              string               `json:"in"`
	Description     string               `json:"description"`
	Required        bool                 `json:"required"`
	Deprecated      bool                 `json:"deprecated"`
	AllowEmptyValue bool                 `json:"allowEmptyValue"`
	Style           string               `json:"style"`
	Explode         bool                 `json:"explode"`
	AllowReserved   bool                 `json:"allowReserved"`
	Schema          Schema               `json:"schema"`
	Example         interface{}          `json:"example"`
	Examples        map[string]Example   `json:"examples"`
	Contents        map[string]MediaType `json:"contents"`
}
