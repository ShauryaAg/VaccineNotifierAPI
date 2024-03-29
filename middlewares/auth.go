package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"cov-api/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			decoded, err := utils.ParseToken(jwtToken)
			if err != nil || decoded == nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else {
				r.Header.Set("decoded", (*decoded)["userId"].(string))
				next.ServeHTTP(w, r)
			}
		}
	})
}
