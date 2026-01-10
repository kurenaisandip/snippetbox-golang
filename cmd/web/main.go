package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// })) // this will add the file and line number of the log call

	// Initialize a new instance of our application struct, containing the
	// dependencies (for now, just the structured logger).
	app := &application{
		logger: logger,
	}

	// log.Printf("starting server on %s", *addr)
	// logger.Info("starting server", "addr", *addr)
	logger.Info("starting server", slog.Any("addr", ":4000"))

	err := http.ListenAndServe(*addr, app.routes())
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
