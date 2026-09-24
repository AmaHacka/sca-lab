// Чат-бот — учебный сервис ЛР №4 (вариант 07).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{}

// handleWS — эхо-сервер поверх WebSocket.
func handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			return
		}
		if err := c.WriteMessage(mt, msg); err != nil {
			return
		}
	}
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
	r := mux.NewRouter()
	r.HandleFunc("/health", handleHealth)
	r.HandleFunc("/ws", handleWS)
	r.HandleFunc("/token", handleToken)

	var handler http.Handler = r
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("chatbot: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
