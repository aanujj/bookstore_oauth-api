package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessToken(t *testing.T) {
	act := GetNewAccessToken()

	// if act != {
	// 	fmt.Println("Brand new access token cannot be nil ")
	// }

	assert.Equal(t, "", act.Token, "Brand new access token cannot be nil ")
	assert.Equal(t, false, act.IsExpired(), "brand new access token should not be expired")
	assert.Equal(t, 0, int(act.UserID), "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	act := AccessToken{} // here we have created empty access _token

	//current time +24 hours [act.expires]
	act.Expires = time.Now().UTC().Add(24 * time.Hour).Unix()
	assert.Equal(t, false, act.IsExpired(), "Empty access Token must be expired")

}
