package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(v interface{}) interface{} {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(v)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			jsonField := getJSONField(v, err.Field())
			tag := strings.ToUpper(err.Tag()) // Convert tag to uppercase
			if err.Param() != "" {
				validationErrors[jsonField] = fmt.Sprintf("%s_%s", tag, err.Param())
			} else {
				validationErrors[jsonField] = tag // Use the uppercase tag
			}
		}
		return validationErrors
	}
	return nil
}

// getJSONField retrieves the JSON tag for a given struct field
func getJSONField(structType interface{}, fieldName string) string {
	field, _ := reflect.TypeOf(structType).Elem().FieldByName(fieldName)
	return field.Tag.Get("json")
}
