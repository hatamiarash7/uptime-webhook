package requests

import "github.com/go-playground/validator/v10"

// GetErrorMessages returns a map of field name and error message
func GetErrorMessages(errors validator.ValidationErrors) map[string]string {
	// TODO: validation messages should be improved
	var errorsMessages = make(map[string]string)
	for i := range errors {
		errorsMessages[errors[i].Field()] = errors[i].Error()
	}
	return errorsMessages
}
