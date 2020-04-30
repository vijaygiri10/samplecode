package helpers

import (
	"regexp"
	"time"
)

func ValidateEmail(email interface{}) bool {
	switch email.(type) {
	case string:
		return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email.(string))
	case []string:
		emails := email.([]string)
		for _, email := range emails {
			if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
				return false
			}
		}
	default:
		return false
	}
	return true
}

func ValidateDateTime(datetime string) bool {
	Re := regexp.MustCompile(`(19|20)[0-9][0-9]-(0[0-9]|1[0-2])-(0[1-9]|([12][0-9]|3[01]))T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]`)
	return Re.MatchString(datetime)
}

func ValidateTime(from string, to string) bool {
	fromTime, _ := time.Parse(time.RFC3339, from+"+00:00")
	toTime, _ := time.Parse(time.RFC3339, to+"+00:00")

	return (fromTime.Unix() < toTime.Unix())
}

func ValidateTimeZone(timezone string) bool {
	_, err := time.LoadLocation(timezone)
	return (err == nil)
}
