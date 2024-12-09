package utils

import "github.com/go-playground/validator/v10"

func ValidateBodyErrMsg(err error) map[string]string {
	errMsg := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMsg[err.Field()] = "The " + err.Field() + " field is required"
		case "email":
			errMsg[err.Field()] = "The " + err.Field() + " field must be a valid email address"
		case "gte":
			errMsg[err.Field()] = "The " + err.Field() + " field must be greater than or equal to " + err.Param()
		case "lte":
			errMsg[err.Field()] = "The " + err.Field() + " field must be less than or equal to " + err.Param()
		case "len":
			errMsg[err.Field()] = "The " + err.Field() + " field must be " + err.Param() + " characters long"
		case "max":
			errMsg[err.Field()] = "The " + err.Field() + " field must be less than or equal to " + err.Param()
		case "min":
			errMsg[err.Field()] = "The " + err.Field() + " field must be greater than or equal to " + err.Param()
		default:
			errMsg[err.Field()] = "The " + err.Field() + " field is invalid"
		}
	}

	return errMsg
}
