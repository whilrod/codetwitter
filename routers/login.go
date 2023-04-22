package routers

import (
	"encoding/json"
	"net/http"
	"src/codetwitter/bd"
	"src/codetwitter/jwt"
	"src/codetwitter/models"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	//w.Header().Add("Authorization", "Bearer")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválido "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña inválido ", 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al generar token correspondiente. "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")

	//http.Error(w, "Aquí pasé", http.StatusAccepted)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(1 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
