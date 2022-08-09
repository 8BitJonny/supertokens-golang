package tpmodels

import (
	"github.com/supertokens/supertokens-golang/supertokens"
)

type TypeProvider struct {
	ID string

	GetAuthorisationRedirectURL    func(clientID *string, redirectURI string, userContext supertokens.UserContext) (TypeAuthorisationRedirect, error)
	ExchangeAuthCodeForOAuthTokens func(clientID *string, redirectInfo TypeRedirectURIInfo, userContext supertokens.UserContext) (TypeOAuthTokens, error) // For apple, add userInfo from callbackInfo to oAuthTOkens
	GetUserInfo                    func(clientID *string, oAuthTokens TypeOAuthTokens, userContext supertokens.UserContext) (TypeUserInfo, error)
}

type TypeRedirectURIQueryParams map[string]interface{}
type TypeOAuthTokens map[string]interface{}

type TypeAuthorisationRedirect struct {
	URLWithQueryParams string
	PKCECodeVerifier   *string
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

type TypeCodeChallenge struct {
	CodeChallenge       string
	CodeChallengeMethod string
}

type TypeSignInUpInput struct {
	// Either of the below
	RedirectURIInfo *TypeRedirectURIInfo
	OAuthTokens     *TypeOAuthTokens
}

type TypeRedirectURIInfo struct {
	RedirectURI            string
	RedirectURIQueryParams *TypeRedirectURIQueryParams // This is separate because of apple
	PKCECodeVerifier       *string                     // Optional, if PKCE enabled
}

type TypeResponsesFromProvider struct {
	OAuthTokens *TypeOAuthTokens
	UserInfo    *map[string]interface{}
}
