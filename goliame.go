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
	To     []string
	Passwd string
	Host   string
	Port   string
	Body   []byte
}

// New builds a new msg instance
func New(from, passwd, host, port string, to []string) *Msg {
	return &Msg{
		From:   from,
		To:     to,
		Passwd: passwd,
		Host:   host,
		Port:   port,
	}
}

// Send message, passing email parameters
func (m *Msg) Send() (bool, error) {
	auth := smtp.PlainAuth("", m.From, m.Passwd, m.Host)

	// send email with real body
	if m.Body != nil {
		uri := fmt.Sprintf("%s:%s", m.Host, m.Port)
		body := fmt.Sprintf("Subject: CERT MONITOR\n%v \n\n", m.Body)
		err := smtp.SendMail(uri, auth, m.From, m.To, []byte(body))

		if err != nil {
			return false, err
		}

		return true, nil
	}

	// If body is nil, you cannot send email: we need to read file first
	err := errors.New("no email body")
	return false, err
}

// LoadFile fetch file content in order to build email body
func (m *Msg) LoadFile(fpath string) error {
	// build file content object
	fcontent, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	// assigne to field's struct
	m.Body = fcontent

	return nil
}
