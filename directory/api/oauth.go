package api

import (
	"fmt"
	"net/http"
	"strings"
)

//GetOrganization extracts the organization from the user:memberof:organization scope passed in a JTW token
func getOrganization(r *http.Request) (string, error) {
	org := ""
	scopes := GetScopes(r)
	for _, scope := range scopes {
		if strings.HasPrefix(scope, "user:memberof:") {
			org = strings.TrimPrefix(scope, "user:memberof:")
			break
		}
	}

	if org == "" {
		return "", fmt.Errorf("No scope user:memberof: present in the JWT token")
	}

	return org, nil
}
