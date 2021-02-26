package email

type email struct {
	subject       string
	content       string
	toAddressList []string
}

// NewEmail ...
func (s *Sender) NewEmail() *email {
	return &email{}
}

// SetSubject ...
func (e *email) SetSubject(subject string) *email {
	e.subject = subject
	return e
}

// SetContent ...
func (e *email) SetContent(content string) *email {
	e.content = content
	return e
}

// SetToAddressList ...
func (e *email) SetToAddressList(toAddressList ...string) *email {
	e.toAddressList = toAddressList
	return e
}
