package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ingjeffer/twittor/bd"
	"github.com/ingjeffer/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet ", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
