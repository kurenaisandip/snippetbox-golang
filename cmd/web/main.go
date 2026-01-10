package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// })) // this will add the file and line number of the log call

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// log.Printf("starting server on %s", *addr)
	// logger.Info("starting server", "addr", *addr)
	logger.Info("starting server", slog.Any("addr", ":4000"))

	err := http.ListenAndServe(*addr, mux)
	// err := http.ListenAndServe(":4000", mux)
	// log.Fatal(err)
	logger.Error(err.Error())
	os.Exit(1)
}
