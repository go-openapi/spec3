package spec3

import (
	"encoding/json"
	"strings"
)

// Extensions vendor specific extensions
type Extensions map[string]interface{}

// Add adds a value to these extensions
func (e Extensions) Add(key string, value interface{}) {
	realKey := strings.ToLower(key)
	e[realKey] = value
}

// GetString gets a string value from the extensions
func (e Extensions) GetString(key string) (string, bool) {
	if v, ok := e[strings.ToLower(key)]; ok {
		str, ok := v.(string)
		return str, ok
	}
	return "", false
}

// GetBool gets a string value from the extensions
func (e Extensions) GetBool(key string) (bool, bool) {
	if v, ok := e[strings.ToLower(key)]; ok {
		str, ok := v.(bool)
		return str, ok
	}
	return false, false
}

// GetStringSlice gets a string value from the extensions
func (e Extensions) GetStringSlice(key string) ([]string, bool) {
	if v, ok := e[strings.ToLower(key)]; ok {
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

// VendorExtensible composition block.
type VendorExtensible struct {
	Extensions Extensions
}

// AddExtension adds an extension to this extensible object
func (v *VendorExtensible) AddExtension(key string, value interface{}) {
	if value == nil {
		return
	}
	if v.Extensions == nil {
		v.Extensions = make(map[string]interface{})
	}
	v.Extensions.Add(key, value)
}

// MarshalJSON marshals the extensions to json
func (v VendorExtensible) MarshalJSON() ([]byte, error) {
	toser := make(map[string]interface{})
	for k, v := range v.Extensions {
		lk := strings.ToLower(k)
		if strings.HasPrefix(lk, "x-") {
			toser[k] = v
		}
	}
	return json.Marshal(toser)
}

// UnmarshalJSON for this extensible object
func (v *VendorExtensible) UnmarshalJSON(data []byte) error {
	var d map[string]interface{}
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	for k, vv := range d {
		lk := strings.ToLower(k)
		if strings.HasPrefix(lk, "x-") {
			if v.Extensions == nil {
				v.Extensions = map[string]interface{}{}
			}
			v.Extensions[k] = vv
		}
	}
	return nil
}
