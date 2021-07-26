package goliame

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/smtp"
)

// Msg is a struct to hold email information
type Msg struct {
	From   string
	To     string
	Passwd string
	Host   string
	Port   string
	Body   []byte
}

// Send message, passing email parameters
func (m *Msg) Send() (bool, error) {
	auth := smtp.PlainAuth("", m.From, m.Passwd, m.Host)
	uri := fmt.Sprintf("%s:%s", m.Host, m.Port)

	if m.Body != nil {
		err := smtp.SendMail(uri, auth, m.From, []string{m.To}, []byte(m.Body))

		if err != nil {
			return false, err
		}
		return true, nil
	}

	// If body is nil, you cannot send email: we need to read file first
	err := errors.New("no email body")
	return false, err
}

// ReadFile fetch file content in order to build email body
func (m *Msg) ReadFile(filename string) error {
	// build file content object
	fcontent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// assigne to field's struct
	m.Body = fcontent

	return nil
}
