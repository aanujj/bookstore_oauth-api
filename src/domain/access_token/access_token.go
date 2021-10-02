package access_token

import (
	"strings"
	"time"

	"github.com/bookstore_oauth-api/src/utils/errors"
)

type AccessToken struct {
	Token    string `json:"access_token"`
	UserID   int64  `json:"user_id"`
	ClientID int64  `json:"client_id"` // which app user is using andorid or web [set a time limit as per user]
	Expires  int64  `json:"exipres"`
}

//validate access token

func (act *AccessToken) Validate() *errors.RestErr {
	act.Token = strings.TrimSpace(act.Token)
	if len(act.Token) == 0 {
		return errors.StatusBadRequestError("invalid access token id ")
	}
	if act.UserID <= 0 {
		return errors.StatusBadRequestError("invalid user id ")
	}
	if act.ClientID <= 0 {
		return errors.StatusBadRequestError("invalid client id ")
	}
	if act.Expires <= 0 {
		return errors.StatusBadRequestError("invalid expiration time ! ")
	}
	return nil
}

//we are giving user access token valid for 24 hours
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
}

//false means it is not expired yet , if we made true is expired
func (act *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(act.Expires, 0)
	return now.After(expirationTime)
}
