package spec3

// Operation describes a single API operation on a path.
type Operation struct {
	VendorExtensible

	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   []Parameter           `json:"parameters,omitempty"`
	RequestBody  RequestBody           `json:"requestBody,omitempty"`
	Responses    Responses             `json:"responses,omitempty"`
	Callbacks    map[string]Callback   `json:"callbacks,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty"`
	Servers      []Server              `json:"servers,omitempty"`
}
