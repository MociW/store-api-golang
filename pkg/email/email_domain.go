package email

type EmailData struct {
	Name string
	OTP  string
}

type EmailService interface {
	Send(to string, subject string, body EmailData) error
}
