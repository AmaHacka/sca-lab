// Сервис аутентификации — учебный сервис ЛР №4 (вариант 10).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

// handleHash возвращает bcrypt-хэш присланного пароля.
func handleHash(w http.ResponseWriter, r *http.Request) {
	password, _ := io.ReadAll(r.Body)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]string{"hash": string(hash)})
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
	r.HandleFunc("/hash", handleHash)
	r.HandleFunc("/token", handleToken)

	var handler http.Handler = r
	handler = csrf.Protect([]byte("lab4-32-byte-long-csrf-auth-key!"), csrf.Secure(false))(handler)
	handler = h2c.NewHandler(handler, &http2.Server{}) // HTTP/2 без TLS для внутренней сети
	log.Println("auth: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
