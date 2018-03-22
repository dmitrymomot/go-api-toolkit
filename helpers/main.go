package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/jmcvetta/randutil"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

// JSONToMap converts json to map
func JSONToMap(j string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	if err := json.Unmarshal([]byte(j), &data); err != nil {
		return nil, err
	}
	return data, nil
}

// StructToMap converts structure to map
func StructToMap(s interface{}) (map[string]interface{}, error) {
	jdata, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return JSONToMap(string(jdata))
}

// MapToStruct converts map to structure
func MapToStruct(data map[string]interface{}, output *interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:  "json",
		Metadata: nil,
		Result:   output,
	})
	if err != nil {
		return err
	}
	if err := decoder.Decode(data); err != nil {
		return err
	}
	return nil
}

// Hash hashes string
func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// HashString hashes string
func HashString(s string) string {
	hash, err := Hash(s)
	if err != nil {
		log.Println(err)
		return ""
	}
	return hash
}

// CheckHash compares hash with string
func CheckHash(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}

// RandString helper
func RandString(n int) string {
	str, _ := randutil.AlphaString(n)
	return str
}

// IsPhoneNumber checks if string is phone number or not
func IsPhoneNumber(s string) bool {
	matched, err := regexp.MatchString("^(+d{1,2}s)?(?d{3})?[s.-]d{3}[s.-]d{4}$", s)
	if err != nil {
		return false
	}
	return matched
}

// ParseDOB parses date of birth from format mm/dd/yyyy
func ParseDOB(dob string) (time.Time, error) {
	birthday, err := time.Parse("01/02/2006", dob)
	if err != nil {
		return time.Time{}, errors.New("Wrong birthday value format")
	}
	if birthday.IsZero() {
		return time.Time{}, errors.New("Wrong birthday value format")
	}
	return birthday, nil
}

// CheckUserAge checks whether user has enough age to use application
func CheckUserAge(minAge int, birthday time.Time) error {
	now := time.Now()
	minBirthday := time.Date(
		now.Year()-minAge,
		now.Month(),
		now.Day(),
		0, 0, 0, 0, time.UTC)
	if birthday.After(minBirthday) == true {
		return errors.New("Sorry, you're too young")
	}
	return nil
}
