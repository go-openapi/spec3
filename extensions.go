package spec3

import (
	"encoding/json"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Extensions vendor specific extensions
type Extensions struct {
	data OrderedMap
}

func (e Extensions) Set(key string, value interface{}) bool {
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	return e.data.Set(key, value)
}

// Add adds a value to these extensions
func (e Extensions) Add(key string, value interface{}) {
	e.Set(key, value)
}

func (e Extensions) GetOK(key string) (interface{}, bool) {
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	return e.data.GetOK(key)
}

func (e Extensions) Get(key string) interface{} {
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	return e.data.Get(key)
}

// GetString gets a string value from the extensions
func (e Extensions) GetString(key string) (string, bool) {
	if v, ok := e.GetOK(key); ok {
		str, ok := v.(string)
		return str, ok
	}
	return "", false
}

// GetBool gets a boolean value from the extensions
func (e Extensions) GetBool(key string) (bool, bool) {
	if v, ok := e.GetOK(key); ok {
		str, ok := v.(bool)
		return str, ok
	}
	return false, false
}

// GetInt gets an int value from the extensions
func (e Extensions) GetInt(key string) (int, bool) {
	if v, ok := e.GetOK(key); ok {
		switch res := v.(type) {
		case int:
			return res, ok
		case int8:
			return int(res), ok
		case int16:
			return int(res), ok
		case int32:
			return int(res), ok
		case int64:
			return int(res), ok
		default:
			return 0, ok
		}
	}
	return 0, false
}

// GetInt32 gets an int32 value from the extensions
func (e Extensions) GetInt32(key string) (int32, bool) {
	if v, ok := e.GetOK(key); ok {
		str, ok := v.(int32)
		return str, ok
	}
	return 0, false
}

// GetInt64 gets an int64 value from the extensions
func (e Extensions) GetInt64(key string) (int64, bool) {
	if v, ok := e.GetOK(key); ok {
		str, ok := v.(int64)
		return str, ok
	}
	return 0, false
}

// GetStringSlice gets a string value from the extensions
func (e Extensions) GetStringSlice(key string) ([]string, bool) {
	if v, ok := e.GetOK(key); ok {
		strv, ok := v.([]string) // get out quick
		if ok {
			return strv, ok
		}

		// do the thing
		arr, ok := v.([]interface{})
		if !ok {
			return nil, false
		}
		var strs []string
		for _, iface := range arr {
			str, ok := iface.(string)
			if !ok {
				return nil, false
			}
			strs = append(strs, str)
		}
		return strs, ok
	}
	return nil, false
}

// MarshalJSON supports json.Marshaler interface
func (e Extensions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	encodeSortedMap(&w, e.data)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (e Extensions) MarshalEasyJSON(w *jwriter.Writer) {
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	encodeSortedMap(w, e.data)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (e *Extensions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	decodeSortedMap(&r, &e.data)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (e *Extensions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	e.data.filter = MatchExtension
	e.data.normalize = LowerCaseKeys
	decodeSortedMap(l, &e.data)
}

// VendorExtensible composition block.
type VendorExtensible struct {
	Extensions Extensions
}

// AddExtension adds an extension to this extensible object
func (v *VendorExtensible) AddExtension(key string, value interface{}) {
	if value == nil {
		return
	}
	v.Extensions.Add(key, value)
}

// MarshalJSON marshals the extensions to json
func (v VendorExtensible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Extensions)
}

// UnmarshalJSON for this extensible object
func (v *VendorExtensible) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &v.Extensions); err != nil {
		return err
	}
	return nil
}
