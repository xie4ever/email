package email

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

const (
	fromStr     = "From: %s<%s>\r\n"
	toStr       = "To: %s\r\n"
	subjectStr  = "Subject: =?UTF-8?B?%s?= \r\n"
	contentType = "Content-Type: text/html; charset=UTF-8\r\n\r\n"
)

// Send ...
func (s *Sender) Send(e *email) error {
	if e == nil {
		return nil
	}

	if err := s.check(e); err != nil {
		return err
	}

	var toAddressList []string
	e.toAddressList = uniqueStringList(e.toAddressList...)
	for _, toAddress := range e.toAddressList {
		if err := s.isValidAddress(toAddress); err != nil {
			if s.ignoreError {
				continue
			} else {
				return err
			}
		}
		toAddressList = append(toAddressList, toAddress)
	}

	message := fmt.Sprintf(fromStr, s.config.username, s.config.address)
	message += "%s"
	message += fmt.Sprintf(subjectStr, base64.StdEncoding.EncodeToString([]byte(e.subject)))
	message += contentType
	message += e.content

	auth := smtp.PlainAuth("", s.config.address, s.config.password, s.config.host)

	for _, toAddress := range toAddressList {
		var (
			toWho      = fmt.Sprintf(toStr, toAddress)
			tmpMessage = fmt.Sprintf(message, toWho)
			tmpAddr    = []string{toAddress}
		)
		if err := smtp.SendMail(s.hostPortString(), auth, s.address, tmpAddr, []byte(tmpMessage)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Sender) check(e *email) error {
	if len(e.subject) == 0 {
		return errSubjectIsEmpty
	}
	if len(e.content) == 0 {
		return errContentIsEmpty
	}
	if len(e.toAddressList) == 0 {
		return errToAddressListIsEmpty
	}
	if len(s.host) == 0 {
		return errHostNotFound
	}
	if len(s.port) == 0 {
		return errPortNotFound
	}
	if len(s.address) == 0 {
		return errAddressNotFound
	}
	if len(s.username) == 0 {
		return errUsernameNotFound
	}
	if len(s.password) == 0 {
		return errPasswordNotFound
	}
	return nil
}

func (s *Sender) hostPortString() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}
