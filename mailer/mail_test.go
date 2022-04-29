package mailer

import "testing"

func TestMail_SendSMTPMessage(t *testing.T) {
	if err := mailer.SendSMTPMessage(dummyMsg); err != nil {
		t.Errorf("error when sending a message %s", err)
	}
}

func TestMail_SendViaChan(t *testing.T) {
	mailer.Jobs <- dummyMsg
	if res := <-mailer.Results; res.Error != nil {
		t.Errorf("error when sending a message: %s", res.Error)
	}
}

func TestMail_SendViaChan_Invalid(t *testing.T) {
	msg := dummyMsg
	msg.To = "invalid_to"
	mailer.Jobs <- msg
	if res := <-mailer.Results; res.Error == nil {
		t.Error("err is nil when it should not be")
	}
}

func TestMail_SendViaAPI(t *testing.T) {
	msg := dummyMsg
	msg.From = ""
	msg.FromName = ""

	mail := mailer
	mail.API = "unknown"
	mail.APIKey = "abc123"
	mail.APIUrl = "https://localhost.com"

	if err := mail.SendViaAPI(msg, "unknown"); err == nil {
		t.Error("err is nil when it should not be")
	}
}

func TestMail_buildHTMLMessage(t *testing.T) {
	if _, err := mailer.buildHTMLMessage(dummyMsg); err != nil {
		t.Errorf("error when building an html message: %s", err)
	}
}

func TestMail_buildPlainTextMessage(t *testing.T) {
	if _, err := mailer.buildPlainTextMessage(dummyMsg); err != nil {
		t.Errorf("error when building a plain message: %s", err)
	}
}

func TestMail_Send(t *testing.T) {
	if err := mailer.Send(dummyMsg); err != nil {
		t.Errorf("error when sending a message: %s", err)
	}

	mail := mailer
	mail.API = "unknown"
	mail.APIKey = "abc123"
	mail.APIUrl = "https://localhost.com"

	if err := mail.Send(dummyMsg); err == nil {
		t.Error("err is nil when it should not be")
	}
}

func TestMail_selectAPI(t *testing.T) {
	mail := mailer
	mail.API = "unknown"
	if err := mail.selectAPI(dummyMsg); err == nil {
		t.Error("err is nil when it should not be")
	}
}
