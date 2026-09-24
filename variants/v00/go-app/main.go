// Конвертер конфигураций — учебный сервис ЛР №4 (вариант 00).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	yaml "gopkg.in/yaml.v2"
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

var jwtSecret = []byte("lab4-not-a-secret")

// handleToken выдаёт JWT для тестового пользователя.
func handleToken(w http.ResponseWriter, r *http.Request) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "student"})
	s, err := t.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"token": s})
}

func main() {
	r := gin.Default()
	r.Any("/health", gin.WrapF(handleHealth))
	r.Any("/config", gin.WrapF(handleConfig))
	r.Any("/token", gin.WrapF(handleToken))

	var handler http.Handler = r
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("configconv: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
