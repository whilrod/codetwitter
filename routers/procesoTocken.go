package routers

import (
	"errors"

	"src/codetwitter/bd"
	"src/codetwitter/models"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Materdeldesarrollo_GrupoDeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
