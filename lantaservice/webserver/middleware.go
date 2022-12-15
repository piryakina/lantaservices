package webserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func SessionMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userSession, err := GetSession(r)
		if err != nil {
			// if session not exist and page is not internal
			log := strings.Contains(r.URL.Path, "login") || strings.Contains(r.URL.Path, "role") || strings.Contains(r.URL.Path, "news") //TODO убрать мидлварю по test
			//sign := strings.Contains(r.URL.Path, "signin")
			if log {
				inner.ServeHTTP(w, r)
			} else {
				JsonResponse(w, StatusResponse{
					Status: false,
					Detail: "вы не авторизованы!",
				}, 401)
			}

			//http.Redirect(w, r, "/", 302)

			return
		}
		ctx := r.Context()
		data, ok := ctx.Value("user").(map[string]interface{})
		if !ok {
			data = make(map[string]interface{})
		}
		data["UserID"] = userSession.ID
		data["Login"] = userSession.Login
		fmt.Println(userSession.Login, userSession.ID) //todo
		ctx = context.WithValue(r.Context(), "user", data)
		r = r.WithContext(ctx)
		inner.ServeHTTP(w, r)
	})
}
