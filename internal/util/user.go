package util

import (
	"errors"
	"strconv"
)

func GetGenderFromId(id string) (string, error) {
	idInt, err := strconv.Atoi(id[6:8])
	if err != nil {
		return "", err
	}
	if idInt > 0 && idInt < 32 {
		return "m", nil
	} else if idInt > 40 && idInt < 72 {
		return "f", nil
	} else {
		return "", errors.New("invalid gender code")
	}
}

func GetDateOfBirthFromId(id string) (string, error) {
	gender, err := GetGenderFromId(id)
	if err != nil {
		return "", err
	}

	var dateOfBirth int

	if gender == "m" {
		dateOfBirth, err = strconv.Atoi(id[6:8])
		if err != nil {
			return "", err
		}
	} else {
		dateOfBirth, err = strconv.Atoi(id[6:8])
		if err != nil {
			return "", err
		}
		dateOfBirth -= 40
	}

	var monthOfBirth string = id[8:10]
	var yearOfBirth string = id[10:12]

	stringDateOfBirth := strconv.Itoa(dateOfBirth)
	if len(stringDateOfBirth) == 1 {
		stringDateOfBirth = "0" + stringDateOfBirth
	}

	if len(monthOfBirth) == 1 {
		monthOfBirth = "0" + monthOfBirth
	}

	// if yearOfBirth[0] == '0' || yearOfBirth[0] == '1' {
	// 	yearOfBirth = "20" + yearOfBirth
	// } else {
	// 	yearOfBirth = "19" + yearOfBirth
	// }

	return yearOfBirth + "-" + monthOfBirth + "-" + stringDateOfBirth, nil
}

func GetYearOfBirth(date string) string {
	return date[0:4]
}

func GetMonthOfBirth(date string) string {
	return date[5:7]
}

func GetDateOfBirth(date string) string {
	return date[8:10]
}
