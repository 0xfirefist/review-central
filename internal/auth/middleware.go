package auth

import (
	"context"
	"net/http"
)

// var userCtxKey = &contextKey{"user"}

// type contextKey struct {
// 	name string
// }

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookies := r.Cookies()

			token := ""
			for _, cookie := range cookies {
				if cookie.Name == "user" {
					token = cookie.Value
				}
			}

			// Allow unauthenticated users in
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), "token", token)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) string {
	raw := ctx.Value("token").(string)
	return raw
}
