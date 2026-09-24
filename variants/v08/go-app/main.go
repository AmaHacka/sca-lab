// Генератор отчётов — учебный сервис ЛР №4 (вариант 08).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mholt/archiver/v3"
	"github.com/microcosm-cc/bluemonday"
)

// writeJSON отправляет ответ в формате JSON.
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

// handleArchiveFormat определяет формат архива по имени файла.
func handleArchiveFormat(w http.ResponseWriter, r *http.Request) {
	f, err := archiver.ByExtension(r.URL.Query().Get("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]string{"format": fmt.Sprintf("%T", f)})
}

var policy = bluemonday.UGCPolicy()

// handleSanitize очищает пользовательский HTML перед вставкой в отчёт.
func handleSanitize(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(policy.SanitizeBytes(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/archive/format", handleArchiveFormat)
	mux.HandleFunc("/sanitize", handleSanitize)

	var handler http.Handler = mux
	log.Println("reports: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
