package spec3

// Responses is a container for the expected responses of an operation. The container maps a HTTP response code to the expected response.
type Responses struct {
	data map[string]ResponseObject
}

// Response describes a single response from an API Operation, including design-time, static links to operations based on the response.
type Response struct {
	VendorExtensible
	Reference

	Description string               `json:"description"`
	Headers     map[string]Header    `json:"headers"`
	Content     map[string]MediaType `json:"content"`
	Links       map[string]Link      `json:"links"`
}
