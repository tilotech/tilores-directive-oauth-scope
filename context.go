package directive

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// contextKey is used for accessing the context variables.
//
// This is required to prevent conflicts with other context variables.
type contextKey string

var scopesContext = contextKey("scopes")

// ScopesFromContext returns the current scopes.
//
// Panics if no scopes were set previously.
func scopesFromContext(ctx context.Context) []string {
	return ctx.Value(scopesContext).([]string)
}

// ContextWithScopes creates a context that contains the scopes.
func ContextWithScopes(ctx context.Context, request *events.APIGatewayProxyRequest) (context.Context, error) {
	scopes, err := scopesFromAuthorizer(request)
	if err != nil {
		return ctx, fmt.Errorf("could not resolve scopes: %v", err)
	}
	return context.WithValue(ctx, scopesContext, scopes), nil
}

func scopesFromAuthorizer(request *events.APIGatewayProxyRequest) ([]string, error) {
	claims, ok := request.RequestContext.Authorizer["claims"]
	if !ok {
		return nil, fmt.Errorf("claims not found")
	}
	cm, ok := claims.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("claims is not a map")
	}
	scopes, ok := cm["scope"]
	if !ok {
		return nil, fmt.Errorf("scope not found")
	}
	scopesString, ok := scopes.(string)
	if !ok {
		return nil, fmt.Errorf("scope is not a string")
	}
	return strings.Split(scopesString, " "), nil
}
