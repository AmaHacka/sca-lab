// Прокси погоды — учебный сервис ЛР №4 (вариант 04).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
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

// handleExtract достаёт поле из JSON по пути из параметра path.
func handleExtract(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	value := gjson.GetBytes(body, r.URL.Query().Get("path"))
	writeJSON(w, map[string]string{"value": value.String()})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/extract", handleExtract)

	var handler http.Handler = mux
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("weather: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
