package email

import (
	"github.com/badoux/checkmail"
)

// Error ...
type Error string

// Error ...
func (e Error) Error() string {
	return string(e)
}

// DIY error
const (
	errSubjectIsEmpty       = Error("subject is empty")
	errContentIsEmpty       = Error("content is empty")
	errToAddressListIsEmpty = Error("toAddressList is empty")
	errHostNotFound         = Error("host not found")
	errPortNotFound         = Error("port not found")
	errAddressNotFound      = Error("address not found")
	errUsernameNotFound     = Error("username not found")
	errPasswordNotFound     = Error("password not found")
)

// IsSlightError 是否轻微错误（不建议报警）
func IsSlightError(err error) bool {
	if _, ok := err.(checkmail.SmtpError); ok && err != nil {
		return true
	}
	return false
}

// IsSevereError 是否严重错误（建议报警）
func IsSevereError(err error) bool {
	return !IsSlightError(err)
}
