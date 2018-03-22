package validator

import "regexp"

// IsPhoneNumber checks if string is phone number or not
func IsPhoneNumber(s string) bool {
	matched, err := regexp.MatchString("^(+d{1,2}s)?(?d{3})?[s.-]d{3}[s.-]d{4}$", s)
	if err != nil {
		return false
	}
	return matched
}
