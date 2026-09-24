// Каталог книг — учебный сервис ЛР №4 (вариант 09).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	yaml "gopkg.in/yaml.v3"
)

// writeJSON отправляет ответ в формате JSON.
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

// handleConfig разбирает YAML-конфигурацию из тела запроса.
func handleConfig(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var cfg map[string]interface{}
	if err := yaml.Unmarshal(body, &cfg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]int{"keys": len(cfg)})
}

func main() {
	r := gin.Default()
	r.Any("/health", gin.WrapF(handleHealth))
	r.Any("/config", gin.WrapF(handleConfig))

	var handler http.Handler = r
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("books: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
