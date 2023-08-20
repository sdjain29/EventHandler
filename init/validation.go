package init

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func SetupValidation() {
	Validate = validator.New()
}
