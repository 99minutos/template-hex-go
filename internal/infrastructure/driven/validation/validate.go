package validation

import (
	"fmt"
	"reflect"
	"service/internal/domain/server"
	"service/internal/helpers/shortcodes"
	"service/internal/infrastructure/driven/fiber_server"
	"sync"

	"github.com/go-playground/validator/v10"
)

type SingletonValidator struct {
	validator *validator.Validate
}

var (
	instance *SingletonValidator
	once     sync.Once
)

func GetInstance() *SingletonValidator {
	once.Do(func() {
		instance = &SingletonValidator{
			validator: validator.New(),
		}
	})
	return instance
}

func (v *SingletonValidator) Validate(data interface{}) *fiber_server.ErrorDispatcher {
	errManager := fiber_server.NewErrorDispatcher()
	err := v.validator.Struct(data)
	if err == nil {
		return nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		t := reflect.TypeOf(data)
		if !shortcodes.IsStruct(data) {
			t = t.Elem()
		}
		field, exists := t.FieldByName(err.Field())

		var validationError server.ValidationError
		if exists {
			jsonTag := field.Tag.Get("json")
			validationError = server.ValidationError{
				Field:     err.Field(),
				Error:     getValidationMessage(err),
				FieldName: jsonTag,
			}
		} else {
			validationError = server.ValidationError{
				Field: err.Field(),
				Error: getValidationMessage(err),
			}
		}

		errManager.AddError(validationError)
	}
	return errManager
}

func getValidationMessage(err validator.FieldError) string {
	switch tag := err.Tag(); tag {
	case "required":
		return fmt.Sprintf("This field is required!")
	case "oneof":
		return fmt.Sprintf("This field should be one of this values: (%s)", shortcodes.ReplaceSpacesForDelimiter(err.Param(), shortcodes.CommaSpace))
	default:
		return fmt.Sprintf("This field don´t match the rule: %s.", tag)
	}
}
