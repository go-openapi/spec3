package spec3

import (
	"fmt"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// Paths holds the relative paths to the individual endpoints and their operations.
// The path is appended to the URL from the Server Object in order to construct the full URL.
// The Paths MAY be empty, due to ACL constraints.
type Paths struct {
	data OrderedMap
}

func (p *Paths) ForEach(fn func(string, []string) error) error {
	for _, k := range p.data.Keys() {
		scopes, ok := p.data.Get(k).([]string)
		if !ok {
			return fmt.Errorf("security requirement scopes not a []string but %T", p.data.Get(k))
		}
		if err := fn(k, scopes); err != nil {
			return err
		}
	}
	return nil
}

// Keys gets list of all the keys
func (p *Paths) Keys() []string {
	return p.data.Keys()
}

// MarshalJSON supports json.Marshaler interface
func (p Paths) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	encodeSortedMap(&w, p.data)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (p Paths) MarshalEasyJSON(w *jwriter.Writer) {
	encodeSortedMap(w, p.data)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (p *Paths) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	decodeSortedMap(&r, &p.data)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (p *Paths) UnmarshalEasyJSON(l *jlexer.Lexer) {
	decodeSortedMap(l, &p.data)
}

func (p *Paths) Get(path string) *PathItem {
	v, ok := p.data.GetOK(path)
	if !ok {
		return nil
	}

	pi, ok := v.(*PathItem)
	if !ok {
		return nil
	}
	return pi
}

func (p *Paths) GetOK(path string) (*PathItem, bool) {
	v, ok := p.data.GetOK(path)
	if !ok {
		return nil, false
	}

	pi, ok := v.(*PathItem)
	return pi, ok
}

func (p *Paths) Set(path string, item *PathItem) bool {
	return p.data.Set(path, item)
}
