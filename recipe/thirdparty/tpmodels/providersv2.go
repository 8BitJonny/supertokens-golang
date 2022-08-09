package tpmodels

import "github.com/supertokens/supertokens-golang/supertokens"

type TypeAppleInput struct {
	Config   []AppleConfig
	Override func(provider AppleProvider) AppleProvider
}

type AppleConfig struct {
	ClientID     string
	ClientSecret AppleClientSecret
	Scope        []string
}

type AppleClientSecret struct {
	KeyId      string
	PrivateKey string
	TeamId     string
}

type AppleProvider struct {
	GetConfig func(clientID *string, userContext supertokens.UserContext) (AppleConfig, error)

	GetAuthorisationRedirectURL    func(clientID *string, redirectURI string, userContext supertokens.UserContext) (TypeAuthorisationRedirect, error)
	ExchangeAuthCodeForOAuthTokens func(clientID *string, callbackInfo TypeCallbackInfo, userContext supertokens.UserContext) (TypeOAuthTokens, error)
	GetUserInfo                    func(clientID *string, oAuthTokens TypeOAuthTokens, userContext supertokens.UserContext) (TypeUserInfo, error)
}

type TypeOktaInput struct {
	Config   []OktaConfig
	Override func(provider OktaProvider) OktaProvider
}

type OktaConfig struct {
	ClientID     string
	ClientSecret string
	OktaDomain   string
	Scope        []string
}

type OktaProvider struct {
	GetConfig func(clientID *string, userContext supertokens.UserContext) (OktaConfig, error)

	GetAuthorisationRedirectURL    func(clientID *string, redirectURI string, userContext supertokens.UserContext) (TypeAuthorisationRedirect, error)
	ExchangeAuthCodeForOAuthTokens func(clientID *string, callbackInfo TypeCallbackInfo, userContext supertokens.UserContext) (TypeOAuthTokens, error)
	GetUserInfo                    func(clientID *string, oAuthTokens TypeOAuthTokens, userContext supertokens.UserContext) (TypeUserInfo, error)
}
