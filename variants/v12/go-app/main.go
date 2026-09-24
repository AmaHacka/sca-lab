// Файловый обменник — учебный сервис ЛР №4 (вариант 12).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/ulikunitz/xz"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// writeJSON отправляет ответ в формате JSON.
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

// handleCompress сжимает тело запроса в формат xz.
func handleCompress(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	zw, err := xz.NewWriter(&buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(zw, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	zw.Close()
	w.Header().Set("Content-Type", "application/x-xz")
	w.Write(buf.Bytes())
}

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/health", handleHealth)
	r.HandleFunc("/compress", handleCompress)

	var handler http.Handler = r
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("fileshare: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
