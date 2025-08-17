package authvalidation

type PlayerRegister struct {
}

func (r PlayerRegister) ValidatePhoneNumber(phoneNumber string) bool {
	// todo: use regex to validate phone number
	if len(phoneNumber) != 11 {
		return false
	}

	return true
}
