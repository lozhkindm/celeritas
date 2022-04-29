package mailer

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	resource *dockertest.Resource
	pool     *dockertest.Pool
	mailer   = Mail{
		Domain:       "localhost",
		TemplatesDir: "./testdata/mails",
		Host:         "localhost",
		Port:         11025,
		Encryption:   "none",
		FromAddress:  "test@test.test",
		FromName:     "from_test",
		Jobs:         make(chan Message, 1),
		Results:      make(chan Result, 1),
	}
	dummyMsg = Message{
		From:        "ignat@senkin.com",
		FromName:    "Ignat",
		To:          "senka@ignatov.com",
		Subject:     "Hi, Senka!",
		Template:    "test",
		Attachments: []string{"./testdata/mails/test.html.tmpl"},
	}
)

func TestMain(m *testing.M) {
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	pool = p
	opts := dockertest.RunOptions{
		Repository:   "mailhog/mailhog",
		Tag:          "latest",
		ExposedPorts: []string{"1025", "8025"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"1025": {
				{
					HostIP:   "0.0.0.0",
					HostPort: "11025",
				},
			},
			"8025": {
				{
					HostIP:   "0.0.0.0",
					HostPort: "18025",
				},
			},
		},
	}

	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}

	time.Sleep(2 * time.Second)
	go mailer.ListenForMail()
	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)
}
