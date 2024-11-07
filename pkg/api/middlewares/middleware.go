package middlewares

import (
	"context"
	"github.com/dafuqqqyunglean/music_library/pkg/service/auth"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	UserCtx             = "UserId"
)

type UserAuthMiddleware struct {
	service auth.AuthorizationService
	ctx     utility.AppContext
}

func NewUserAuthMiddleware(service auth.AuthorizationService, ctx utility.AppContext) *UserAuthMiddleware {
	return &UserAuthMiddleware{
		service: service,
		ctx:     ctx,
	}
}

func (m *UserAuthMiddleware) UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			utility.NewErrorResponse(w, m.ctx, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			utility.NewErrorResponse(w, m.ctx, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := m.service.ParseToken(headerParts[1])
		if err != nil {
			utility.NewErrorResponse(w, m.ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), UserCtx, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
