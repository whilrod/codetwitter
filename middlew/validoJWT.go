package middlew

import (
	"net/http"
	"src/codetwitter/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("Authorization", "Bearer")
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {

			http.Error(w, "Error en Token! "+err.Error(), http.StatusBadRequest)

			return
		}
		next.ServeHTTP(w, r)
	}
}
