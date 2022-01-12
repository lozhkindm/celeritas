package celeritas

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	RootPath string
	Routes   *chi.Mux
	config   config
}

type config struct {
	port     string
	renderer string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middlewares"},
	}

	if err := c.init(pathConfig); err != nil {
		return err
	}
	if err := c.checkDotEnv(rootPath); err != nil {
		return err
	}
	if err := godotenv.Load(fmt.Sprintf("%s/.env", rootPath)); err != nil {
		return err
	}

	infoLog, errorLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath
	c.config = config{port: os.Getenv("PORT"), renderer: os.Getenv("RENDERER")}
	c.Routes = c.routes().(*chi.Mux)

	return nil
}

func (c *Celeritas) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      c.routes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
		IdleTimeout:  30 * time.Second,
		ErrorLog:     c.ErrorLog,
	}
	c.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	if err := srv.ListenAndServe(); err != nil {
		c.ErrorLog.Fatal(err)
	}
}

func (c *Celeritas) init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		if err := c.CreateDirIfNotExists(fmt.Sprintf("%s/%s", root, path)); err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	if err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path)); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
