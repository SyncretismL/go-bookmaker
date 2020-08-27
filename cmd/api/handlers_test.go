package main

import (
	"bookmaker/internal/config"
	"bookmaker/pkg/logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestReady(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9000/api/v1/ready", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	log, err := logger.NewLogger(config.Log{})

	checkError(err, t)

	handler := &Handler{
		logger: log,
		db:     db,
	}

	http.HandlerFunc(handler.ready).
		ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
}
