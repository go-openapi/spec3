package spec3

import "github.com/go-openapi/jsonreference"

// Reference is a simple object to allow referencing other components in the specification, internally and externally.
type Reference struct {
	Ref jsonreference.Ref
}
