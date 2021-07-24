package api

import (
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/models"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func EmailExists(apiImplementation models.APIImplementation, options models.APIOptions) error {
	if apiImplementation.EmailExistsGET == nil {
		options.OtherHandler(options.Res, options.Req)
		return nil
	}
	email := options.Req.URL.Query().Get("email")
	if email == "" {
		return supertokens.BadInputError{Msg: "Please provide the email as a GET param"}
	}
	result := apiImplementation.EmailExistsGET(email, options)
	supertokens.Send200Response(options.Res, result)
	return nil
}
