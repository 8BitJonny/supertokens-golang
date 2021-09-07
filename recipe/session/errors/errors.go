package errors

const (
	UnauthorizedErrorStr       = "UNAUTHORISED"
	TryRefreshTokenErrorStr    = "TRY_REFRESH_TOKEN"
	TokenTheftDetectedErrorStr = "TOKEN_THEFT_DETECTED"
)

// TryRefreshTokenError used for when the refresh API needs to be called
type TryRefreshTokenError struct {
	Msg string
}

func (err TryRefreshTokenError) Error() string {
	return err.Msg
}

// TokenTheftDetectedError used for when token theft has happened for a session
type TokenTheftDetectedError struct {
	Msg     string
	Payload TokenTheftDetectedErrorPayload
}

// TODO: test if this gets proper deserialised
type TokenTheftDetectedErrorPayload struct {
	SessionHandle string `json:"sessionHandle"`
	UserID        string `json:"userId"`
}

func (err TokenTheftDetectedError) Error() string {
	return err.Msg
}

// UnauthorizedError used for when the user has been logged out
type UnauthorizedError struct {
	Msg          string
	ClearCookies *bool
}

func (err UnauthorizedError) Error() string {
	return err.Msg
}
