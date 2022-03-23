package app

import (
	"context"
	"fmt"
	"gogo/models"
	u "gogo/utils"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/new", "/api/user/login"} // Doğrulama istemeyen endpointler
		requestPath := r.URL.Path                               // mevcut istek yolu

		// Gelen isteğin doğrulama isteyip istemediği kontrol edilir
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // Header'dan token alınır

		if tokenHeader == "" { // Token yoksa "403 Unauthorized" hatası dönülür
			response = u.Message(false, "Token gönderilmelidir!")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") // Token'ın "Bearer {token} / Token {token}" formatında gelip gelmediği kontrol edilir
		if len(splitted) != 2 {
			response = u.Message(false, "Hatalı ya da geçersiz token!")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] // Token'ın doğrulama yapmamıza yarayan kısmı alınır
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { // Token hatalı ise 403 hatası dönülür
			response = u.Message(false, "Token hatalı!")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { // Token geçersiz ise 403 hatası dönülür
			response = u.Message(false, "Token geçersiz!")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// Doğrula başarılı ise işleme devam edilir
		fmt.Sprintf("Kullanıcı %", tk.Username) // Kullanıcı adı console'a basılır
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
