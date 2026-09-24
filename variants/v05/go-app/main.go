// Трекер задач — учебный сервис ЛР №4 (вариант 05).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/html"
)

// writeJSON отправляет ответ в формате JSON.
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

// handleTitle извлекает заголовок из присланной HTML-страницы.
func handleTitle(w http.ResponseWriter, r *http.Request) {
	doc, err := html.Parse(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var title string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = strings.TrimSpace(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	writeJSON(w, map[string]string{"title": title})
}

func main() {
	app := fiber.New()
	app.All("/health", adaptor.HTTPHandlerFunc(handleHealth))
	app.All("/metrics", adaptor.HTTPHandlerFunc(promhttp.Handler().ServeHTTP))
	app.All("/title", adaptor.HTTPHandlerFunc(handleTitle))

	log.Println("tracker: listening on :8080")
	log.Fatal(app.Listen(":8080"))
}
