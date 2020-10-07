package core

import (
	"github.com/go-playground/validator"
	"regexp"
)

func SkuValidation(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	var matches []string = regex.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}
