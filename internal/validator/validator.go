package validator

import (
	"regexp"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// NOTE: contains all the errors in a hashmap
type Validator struct {
	Errors map[string]string
}

// constructor
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid returns true if the errors map is doesn't contain any entries
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// adds a new error to the current validator
func (v *Validator) AddError(key, msg string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = msg
	}
}

// Check() adds a the error to the validator if validation check is !ok
func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

// generic PermittedValue() returns true if the given value is in the provided list
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}

	return false
}

// returns true if the provided string matches the provided the regular expression
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// generic returns true if all the values are Unique in the whole array/slice
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
