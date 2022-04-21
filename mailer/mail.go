package mailer

import (
	"bytes"
	"fmt"
	"html/template"
)

type Mail struct {
	Domain       string
	TemplatesDir string
	Host         string
	Port         string
	Username     string
	Password     string
	Encryption   string
	FromAddress  string
	FromName     string
	Jobs         chan Message
	Results      chan Result
	API          string
	APIKey       string
	ApiUrl       string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Template    string
	Attachments []string
	Data        interface{}
}

type Result struct {
	Success bool
	Error   error
}

func (m *Mail) ListenForMail() {
	for {
		msg := <-m.Jobs
		if err := m.Send(msg); err != nil {
			m.Results <- Result{Success: false, Error: err}
		} else {
			m.Results <- Result{Success: true, Error: nil}
		}
	}
}

func (m *Mail) Send(msg Message) error {
	// TODO: select API or SMTP
	return m.SendSMTPMessage(msg)
}

func (m *Mail) SendSMTPMessage(msg Message) error {
	return nil
}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	tmpl := fmt.Sprintf("%s/%s.html.tmpl", m.TemplatesDir, msg.Template)
	t, err := template.New("email-html").ParseFiles(tmpl)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", msg.Data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func (m *Mail) buildPlainTextMessage(msg Message) (string, error) {
	tmpl := fmt.Sprintf("%s/%s.plain.tmpl", m.TemplatesDir, msg.Template)
	t, err := template.New("email-plain").ParseFiles(tmpl)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "body", msg.Data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
