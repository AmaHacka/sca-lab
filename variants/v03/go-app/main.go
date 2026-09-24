// Сервис аватаров — учебный сервис ЛР №4 (вариант 03).
// Код намеренно простой: предмет работы — зависимости, а не логика.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/image/webp"
)

// writeJSON отправляет ответ в формате JSON.
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

// handleImageInfo возвращает размеры присланного WebP-изображения.
func handleImageInfo(w http.ResponseWriter, r *http.Request) {
	cfg, err := webp.DecodeConfig(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]int{"width": cfg.Width, "height": cfg.Height})
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

func main() {
	e := echo.New()
	e.Any("/health", echo.WrapHandler(http.HandlerFunc(handleHealth)))
	e.Any("/image/info", echo.WrapHandler(http.HandlerFunc(handleImageInfo)))
	e.Any("/hash", echo.WrapHandler(http.HandlerFunc(handleHash)))

	var handler http.Handler = e
	log.Println("avatars: listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
