// Сервис заметок — учебный сервис ЛР №4 (вариант 01).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/text/language"
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

// handleLang определяет предпочитаемые языки клиента.
func handleLang(w http.ResponseWriter, r *http.Request) {
	tags, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	langs := make([]string, 0, len(tags))
	for _, t := range tags {
		langs = append(langs, t.String())
	}
	writeJSON(w, map[string][]string{"languages": langs})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", handleHealth)
	r.HandleFunc("/ws", handleWS)
	r.HandleFunc("/lang", handleLang)

	var handler http.Handler = r
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("notes: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
