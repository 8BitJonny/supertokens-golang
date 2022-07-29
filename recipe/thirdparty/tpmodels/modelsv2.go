package tpmodels

import (
	"github.com/supertokens/supertokens-golang/supertokens"
)

type TypeCallbackQueryParams map[string]interface{}
type TypeOAuthTokens map[string]interface{}

type TypeCallbackInfo struct {
	URI         string
	QueryParams TypeCallbackQueryParams
}

type TypeAuthorisationRedirect struct {
	URL         string
	QueryParams map[string]interface{}
}

type TypeEmailInfo struct {
	Email      string `json:"id"`
	IsVerified bool   `json:"isVerified"`
}

type TypeUserInfo struct {
	ThirdPartyUserId        string                 `json:"thirdPartyUserId"`
	EmailInfo               *TypeEmailInfo         `json:"emailInfo"`
	RawResponseFromProvider map[string]interface{} `json:"rawResponseFromProvider"`
}

type TypeProvider struct {
	ID string

	GenerateState                  func() (string, error)
	GetAuthorisationRedirectURL    func(input TypeGetAuthorisationRedirectURLInput, userContext supertokens.UserContext) (TypeAuthorisationRedirect, error)
	ExchangeAuthCodeForOAuthTokens func(input TypeExchangeAuthCodeForOAuthTokensInput, userContext supertokens.UserContext) (TypeOAuthTokens, error)
	GetUserInfo                    func(input TypeGetUserInfoInput, userContext supertokens.UserContext) (TypeUserInfo, error)
}

type TypeGetAuthorisationRedirectURLInput struct {
	ClientID    string
	CallbackURI string
}

type TypeExchangeAuthCodeForOAuthTokensInput struct {
	ClientID     string
	CallbackInfo TypeCallbackInfo
}

type TypeGetUserInfoInput struct {
	ClientID     string
	CallbackInfo TypeCallbackInfo
	OAuthTokens  TypeOAuthTokens
}

type TypeSignInUpInput struct {
	ClientID     string
	State        string
	CallbackInfo *TypeCallbackInfo
	OAuthTokens  *TypeOAuthTokens
}

type TypeResponsesFromProvider struct {
	CallbackInfo *TypeCallbackInfo
	OAuthTokens  *TypeOAuthTokens
	UserInfo     *map[string]interface{}
}
