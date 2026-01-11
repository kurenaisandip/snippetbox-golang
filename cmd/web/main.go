package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String(
		"dsn",
		"sqlserver://sa:1@SANDIP-RAI?database=snippetbox&instanceName=SQLEXPRESS",
		"SQL Server data source name",
	)

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// })) // this will add the file and line number of the log call

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// Initialize a new instance of our application struct, containing the
	// dependencies (for now, just the structured logger).
	app := &application{
		logger: logger,
	}

	// log.Printf("starting server on %s", *addr)
	// logger.Info("starting server", "addr", *addr)
	logger.Info("starting server", slog.Any("addr", ":4000"))

	err = http.ListenAndServe(*addr, app.routes())
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
