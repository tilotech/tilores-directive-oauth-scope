package directive

import (
	"context"
)

// HasPrivilege checks the JWT token inside the provided context.Context whether it has
// the provided privilege in its oAuth scopes.
func HasPrivilege(ctx context.Context, privilege string) bool {
	scopes := scopesFromContext(ctx)
	for _, scope := range scopes {
		if scope == privilege {
			return true
		}
	}
	return false
}
