package strcase

// Camel returns a camel-cased version of the input string.
// It currently only supports conversions from pascal case.
func Camel(in string) string {
	return camel(in)
}

// Pascal returns a pascal-cased version of the input string.
// It currently only supports conversions from camel case.
func Pascal(in string) string {
	return pascal(in)
}
