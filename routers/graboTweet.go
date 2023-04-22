package routers

import (
	"encoding/json"
	"net/http"
	"src/codetwitter/bd"
	"src/codetwitter/models"
	"time"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet((registro))
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, intente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha podido insertar el Tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
