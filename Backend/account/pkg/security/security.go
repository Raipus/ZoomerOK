package security

type SecurityInterface interface {
	SendConfirmEmail(username, confirmationLink string) []byte
	SendChangePassword(username, resetLink string) []byte
}
