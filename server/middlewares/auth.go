package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/alfasya/imgo/utils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//parse token from request header
		tokenString := r.Header.Get("Authorization")

		//handling empty authorization header
		if tokenString == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		stringParts := strings.Split(tokenString, " ")

		if len(stringParts) == 0 || strings.ToLower(stringParts[0]) != "bearer" {
			http.Error(w, "missing bearer token", http.StatusUnauthorized)
			return
		}

		tokenString = stringParts[1]

		owner, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		//add claims to request context
		ctx := context.WithValue(r.Context(), "owner", owner)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func FileAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		owner, ok := r.Context().Value("owner").(utils.Owner)
// 		if !ok {
// 			http.Error(w, "invalid type", http.StatusBadRequest)
// 			return
// 		}

// 		uuid := r.PathValue("uuid")

// 		if owner.UserUUID != uuid {
// 			http.Error(w, "forbidden", http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
