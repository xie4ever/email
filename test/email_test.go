package test

import (
	"testing"

	"email"
)

const (
	testHost     = "your host"     // 邮件服务host
	testPort     = "your port"     // 邮件服务端口
	testAddress  = "your address"  // 邮件服务发送者邮箱
	testUsername = "your username" // 邮件服务发送者
	testPassword = "your password" // 邮件服务密码
)

var s *email.Sender

func init() {
	s, _ = email.NewSender(
		email.WithHost(testHost),
		email.WithPort(testPort),
		email.WithAddress(testAddress),
		email.WithUsername(testUsername),
		email.WithPassword(testPassword),
	)
}

// TestSend ...
func TestSend(t *testing.T) {
	mail := s.NewEmail().
		SetSubject("test").
		SetContent("content").
		SetToAddressList("503822883@qq.com")
	if err := s.Send(mail); err != nil {
		t.Fatal(err)
	}
}
