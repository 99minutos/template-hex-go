package server

type ValidationError struct {
	Field     string `json:"field"`
	Error     string `json:"error"`
	FieldName string `json:"field_name"`
}
