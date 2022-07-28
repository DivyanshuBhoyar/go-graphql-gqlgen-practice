package auth

import (
	"context"
	"github.com/DivyanshuBhoyar/gqlgen-prac/internal/users"
	"github.com/DivyanshuBhoyar/gqlgen-prac/pkg/jwt"
	"net/http"
	"strconv"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauth users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// validate JWT token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user from token and check if user exists in db
			user := users.User{Username: username}
			id, err := users.GetUserIdByU8sername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			user.ID = strconv.Itoa(id)
			// put it ot the context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			
			// and pass the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}