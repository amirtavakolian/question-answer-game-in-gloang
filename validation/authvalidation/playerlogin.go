package authvalidation

type PlayerLogin struct {
}

func (r PlayerLogin) ValidatePhoneNumber(phoneNumber string) bool {
	// todo: use regex to validate phone number
	if len(phoneNumber) != 11 {
		return false
	}

	return true
}

func (r PlayerLogin) ValidatePassword(password string) bool {
	// todo: use regex to validate phone number
	if len(password) < 6 {
		return false
	}

	return true
}
