package spec3

// Discriminator When request bodies or response payloads may be one of a number of different schemas, a discriminator object can be used to aid in serialization, deserialization, and validation.
// The discriminator is a specific object in a schema which is used to inform the consumer of the specification of an alternative schema based on the value associated with it.
type Discriminator struct {
	PropertyName string         `json:"propertyName"`
	Mapping      OrderedStrings `json:"mapping"`
}
