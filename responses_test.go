package spec3

import (
	"reflect"
	"testing"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

func TestResponseMap_Get(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForResponseMap()}, args{"default"}, &Response{Description: "default response"}},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForResponseMap()}, args{"201"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseMap_GetOK(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *Response
		wantOK       bool
	}{
		{"Should fetch the item when existent key is passed", fields{buildOrderMapForResponseMap()}, args{"200"}, &Response{Description: "OK"}, true},
		{"Should return nil when non-existent key is passed", fields{buildOrderMapForResponseMap()}, args{"401"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			got, got1 := s.GetOK(tt.args.key)
			if !reflect.DeepEqual(got, tt.wantResponse) {
				t.Errorf("responseMap.GetOK() got = %v, want %v", got, tt.wantResponse)
			}
			if got1 != tt.wantOK {
				t.Errorf("responseMap.GetOK() got1 = %v, want %v", got1, tt.wantOK)
			}
		})
	}
}

func TestResponseMap_Set(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		key string
		val *Response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOK bool
	}{
		{"Should set value when non-existent response code is passed", fields{buildOrderMapForResponseMap()}, args{"201", &Response{Description: "Created"}}, true},
		{"Should fail when existent response code is passed", fields{buildOrderMapForResponseMap()}, args{"200", &Response{Description: "another OK"}}, false},
		{"Should fail when empty key is passed", fields{buildOrderMapForResponseMap()}, args{"", &Response{Description: "description of item #empty"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			if got := s.Set(tt.args.key, tt.args.val); got != tt.wantOK {
				t.Fatalf("responseMap.Set() = %v, want %v", got, tt.wantOK)
			}

			if tt.wantOK {
				gotVal, gotOK := s.GetOK(tt.args.key)
				if !gotOK {
					t.Fatalf("responseMap.GetOK().OK = %v, want %v", gotOK, true)
				}
				if !reflect.DeepEqual(gotVal, tt.args.val) {
					t.Fatalf("responseMap.GetOK().val = %v, want %v", gotVal, tt.args.val)
				}
			}
		})
	}
}

func TestResponseMap_ForEach(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		fn func(string, *Response) error
	}
	type foundResponse struct {
		response *Response
		found    bool
	}
	tests := []struct {
		name             string
		fields           fields
		wantValInForEach map[string]*foundResponse
		wantErr          error
	}{
		{
			"Should iterate 4 items for responseMap fixture",
			fields{buildOrderMapForResponseMap()},
			map[string]*foundResponse{
				"default": &foundResponse{&Response{Description: "default response"}, false},
				"200":     &foundResponse{&Response{Description: "OK"}, false},
				"404":     &foundResponse{&Response{Description: "Not found"}, false},
				"500":     &foundResponse{&Response{Description: "Internal Error"}, false},
			},
			nil,
		},
		{
			"Should return empty array when there are no values in responseMap",
			fields{},
			map[string]*foundResponse{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			err := s.ForEach(func(key string, gotResponse *Response) error {
				if wantVal, ok := tt.wantValInForEach[key]; ok {
					if !reflect.DeepEqual(wantVal.response, gotResponse) {
						t.Fatalf("responseMap.ForEach() for key = %s val = %v, want = %v", key, gotResponse, wantVal.response)
					}
					wantVal.found = true
				} else {
					t.Fatalf("responseMap.ForEach() for key = %s val = %v, want = %v", key, gotResponse, wantVal)
				}
				return nil
			})

			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Fatalf("responseMap.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr == nil {
				for key2, val2 := range tt.wantValInForEach {
					if !val2.found {
						t.Fatalf("responseMap.ForEach() key = %s not found during foreach()", key2)
					}
				}
			}
		})
	}
}

func TestResponseMap_Keys(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{"Should return 4 items for responseMap fixture", fields{buildOrderMapForResponseMap()}, []string{"default", "200", "404", "500"}},
		{"Should return empty array when there are no values in responseMap", fields{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			got := s.Keys()
			if len(got) != 0 || len(tt.wantKeys) != 0 {
				if !reflect.DeepEqual(got, tt.wantKeys) {
					t.Errorf("responseMap.Keys() = %v, want %v", got, tt.wantKeys)
				}
			}
		})
	}
}

func TestResponseMap_MarshalJSON(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := responseMap{
				data: tt.fields.data,
			}
			got, err := s.MarshalJSON()
			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Errorf("responseMap.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responseMap.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseMap_MarshalEasyJSON(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		w *jwriter.Writer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := responseMap{
				data: tt.fields.data,
			}
			s.MarshalEasyJSON(tt.args.w)
		})
	}
}

func TestResponseMap_UnmarshalJSON(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			err := s.UnmarshalJSON(tt.args.data)
			if err == nil && tt.wantErr == nil {
				// nothing to assert here
			} else if err != tt.wantErr {
				t.Errorf("responseMap.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseMap_UnmarshalEasyJSON(t *testing.T) {
	type fields struct {
		data OrderedMap
	}
	type args struct {
		l *jlexer.Lexer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &responseMap{
				data: tt.fields.data,
			}
			s.UnmarshalEasyJSON(tt.args.l)
		})
	}
}

func buildEmptyOrderMapForResponseMap() OrderedMap {
	return OrderedMap{
		filter: matchResponseCode,
	}
}

func buildOrderMapForResponseMap() OrderedMap {
	return OrderedMap{
		data: map[string]interface{}{
			"default": &Response{Description: "default response"},
			"200":     &Response{Description: "OK"},
			"404":     &Response{Description: "Not found"},
			"500":     &Response{Description: "Internal Error"},
		},
		keys: []string{
			"default",
			"200",
			"404",
			"500",
		},
		filter: matchResponseCode,
	}
}

func buildResponseMapFixture() responseMap {
	m := responseMap{
		data: buildOrderMapForResponseMap(),
	}

	return m
}
