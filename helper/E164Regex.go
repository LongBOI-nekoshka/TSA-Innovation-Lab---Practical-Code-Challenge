package helper

import (
	"regexp"
)

func CheckIfE164OrAussieNum(phoneNumber string) bool {
	matched, _ := regexp.Match(`^(\+?\(61\)|\(\+?61\)|\+?61|\(0[1-9]\)|0[1-9])?( ?-?[0-9]){7,9}$`, []byte(phoneNumber))
	return matched
}
