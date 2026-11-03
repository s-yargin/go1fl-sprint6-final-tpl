package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	indexFile = "../index.html"
	formField = "myFile"
)

// Index sends the upload form to the client.
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	html, err := os.ReadFile(indexFile)
	if err != nil {
		http.Error(w, "cannot read upload form", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err = w.Write(html); err != nil {
		return
	}
}

// Upload reads a submitted file, converts its contents and saves the result.
func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile(formField)
	if err != nil {
		http.Error(w, "cannot get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read uploaded file", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "cannot convert file contents", http.StatusInternalServerError)
		return
	}

	name := outputName(header.Filename)
	if err := os.WriteFile(name, []byte(result), 0644); err != nil {
		http.Error(w, "cannot save converted file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err = fmt.Fprint(w, result); err != nil {
		return
	}
}

func outputName(uploadedName string) string {
	// UTC.String contains colons, which are invalid in Windows file names.
	timestamp := strings.ReplaceAll(time.Now().UTC().String(), ":", "-")
	return timestamp + filepath.Ext(uploadedName)
}
