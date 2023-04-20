package authorization

import (
	"net/http"
	"strings"
)

func ExtractRoleFromRequest(r *http.Request) string {
	token := r.Header.Get("Role")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
