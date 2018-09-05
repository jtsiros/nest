package auth

import (
	"testing"

	"github.com/jtsiros/nest/config"
	"github.com/stretchr/testify/assert"
)

func TestValidOAuthConfig(t *testing.T) {
	appCfg := config.Config{ClientID: "123"}
	cfg := NewConfig(appCfg)
	assert.Equal(t, "123", cfg.ClientID)
	assert.Equal(t, []string{"authorization_code"}, cfg.Scopes)
	assert.NotNil(t, cfg.Endpoint.AuthURL)
	assert.NotNil(t, cfg.Endpoint.TokenURL)
}

func TestStaticToken(t *testing.T) {
	token := "123"
	tok, err := NewConfigWithToken(token).Token()
	assert.Nil(t, err, "error setting static token source")
	assert.Equal(t, "123", tok.AccessToken)
}
