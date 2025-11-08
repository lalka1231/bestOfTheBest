package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	content := string(data)
	result, err := service.AutoDetectAndConvert(content)
	if err != nil {
		http.Error(w, "Error converting content", http.StatusInternalServerError)
		return
	}

	filename := "result_" + time.Now().UTC().Format("20060102_150405") + ".txt"
	err = os.WriteFile(filename, []byte(result), 0755)
	if err != nil {
		http.Error(w, "Error saving result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
