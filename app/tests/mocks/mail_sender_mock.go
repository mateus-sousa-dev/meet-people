package mocks

type EmailSenderMock struct {
}

func (e *EmailSenderMock) SendAccountActivationEmail(emailTo string) error {
	return nil
}
