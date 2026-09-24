// Агрегатор RSS — учебный сервис ЛР №4 (вариант 11).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
	"golang.org/x/net/html"
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

// handleExtract достаёт поле из JSON по пути из параметра path.
func handleExtract(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	value := gjson.GetBytes(body, r.URL.Query().Get("path"))
	writeJSON(w, map[string]string{"value": value.String()})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/title", handleTitle)
	mux.HandleFunc("/lang", handleLang)
	mux.HandleFunc("/extract", handleExtract)

	var handler http.Handler = mux
	log.Println("rss: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
