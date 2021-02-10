package main

import (
	"clydotron/skt_t/pkg/models/mongo"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template

	dbx     mongo.MongoDBX //replace with interface:
	session *sessions.Session
}

func main() {

	addr := flag.String("addr", ":4040", "HTTP network address")
	secret := flag.String("secret", "s6Nxh+pPbnzHbS*+9Pk8xGWhTzbca@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
		dbx:           mongo.MongoDBX{},
		session:       session,
	}

	err = app.dbx.ConnectToMongo()
	if err != nil {
		errorLog.Fatal(err)
	}

	app.dbx.InitControllers()

	//defer app.dbx.DisconnectFromMongo()

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.createRouterP(),
	}

	infoLog.Println("Starting server on", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
