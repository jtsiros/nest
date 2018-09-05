package config

// TokenURL provides the default authentication URL for OAuth API.
const TokenURL = "https://api.home.nest.com/oauth2/access_token"

// APIURL provides the default Nest REST API endpoint.
const APIURL = "https://developer-api.nest.com"

// AuthURL for generating authorization request
const AuthURL = "https://home.nest.com/login/oauth2"

// Config represents configuration for Nest authentication API
type Config struct {
	ClientID string
	Secret   string
	APIURL   string
}
