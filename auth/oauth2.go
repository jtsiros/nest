package auth

import (
	"github.com/jtsiros/nest/config"
	"golang.org/x/oauth2"
)

// NewConfig creates a new oauth2-backed configuration for handling Nest API
// Authentication. This is intended to be used with an oauth2.Client object.
func NewConfig(cfg config.Config) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: cfg.ClientID,
		Scopes:   []string{"authorization_code"},
		Endpoint: oauth2.Endpoint{
			TokenURL: config.TokenURL,
			AuthURL:  config.AuthURL,
		},
	}
	return conf
}

// NewConfigWithToken creates a new oauth2 Token Source based on a static generated token.
// This is helpful when a token does not have an expiration or fetching an authorization code
// is not desirable or required.
func NewConfigWithToken(token string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
}
