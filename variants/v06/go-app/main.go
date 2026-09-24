// SSH-оркестратор — учебный сервис ЛР №4 (вариант 06).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/ssh"
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

// handleKey разбирает открытый SSH-ключ в формате authorized_keys.
func handleKey(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	key, comment, _, _, err := ssh.ParseAuthorizedKey(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]string{"type": key.Type(), "fingerprint": ssh.FingerprintSHA256(key), "comment": comment})
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
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/keys/parse", handleKey)
	mux.HandleFunc("/config", handleConfig)

	var handler http.Handler = mux
	log.Println("sshorch: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
