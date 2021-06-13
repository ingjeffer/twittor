package middleware

import (
	"net/http"

	"github.com/ingjeffer/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
