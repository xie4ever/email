package test

import (
	"testing"
)

// TestInvalidAddress ...
func TestInvalidAddress(t *testing.T) {
	mail := s.NewEmail().
		SetSubject("test").
		SetContent("content").
		SetToAddressList("503822883")
	if err := s.Send(mail); err != nil {
		t.Fatal(err)
	}
}

// TestEmptyAddress ...
func TestEmptyAddress(t *testing.T) {
	mail := s.NewEmail().
		SetSubject("test").
		SetContent("content")
	if err := s.Send(mail); err != nil {
		t.Fatal(err)
	}
}
