package spec3

// PathItem describes the operations available on a single path.
// A Path Item MAY be empty, due to ACL constraints.
// The path itself is still exposed to the documentation viewer but they will not know which operations and parameters are available.
type PathItem struct {
	Reference

	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Get         Operation   `json:"get"`
	Put         Operation   `json:"put"`
	Post        Operation   `json:"post"`
	Delete      Operation   `json:"delete"`
	Options     Operation   `json:"options"`
	Head        Operation   `json:"head"`
	Patch       Operation   `json:"patch"`
	Trace       Operation   `json:"trace"`
	Servers     []Server    `json:"servers"`
	Parameters  []Parameter `json:"parameters"`
}
