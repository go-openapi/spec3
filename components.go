package spec3

// Components holds a set of reusable objects for different aspects of the OAS.
// All objects defined within the components object will have no effect on the API unless they are explicitly referenced from properties outside the components object.
type Components struct {
	VendorExtensible

	Schemas         OrderedSchemas         `json:"schemas,omitempty"`
	Responses       OrderedResponses       `json:"responses,omitempty"`
	Parameters      OrderedParameters      `json:"parameters,omitempty"`
	Examples        OrderedExamples        `json:"examples,omitempty"`
	RequestBodies   OrderedRequestBodies   `json:"requestBodies,omitempty"`
	Headers         OrderedHeaders         `json:"headers,omitempty"`
	SecuritySchemes OrderedSecuritySchemes `json:"securitySchemes,omitempty"`
	Links           OrderedLinks           `json:"links,omitempty"`
	Callbacks       OrderedCallbacks       `json:"callbacks,omitempty"`
}
