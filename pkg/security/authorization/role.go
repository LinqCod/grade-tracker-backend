package authorization

import (
	"net/http"
)

func ExtractRoleFromRequest(r *http.Request) string {
	return r.Header.Get("Role")
}
