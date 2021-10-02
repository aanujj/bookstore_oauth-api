package db

import (
	"github.com/bookstore_oauth-api/src/client/cassandra"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken       = "Select access_token , user_id, client_id ,expires From access_tokens where access_token = ?;"
	queryCreateAccessToken    = "Insert Into access_tokens(access_token , user_id, client_id ,expires) VALUES (?,?,?,?); "
	queryUpdateExpirationTime = "Update access_tokens set expires=? where access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	CreateAT(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct{}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	//implement get access token from cassandra db

	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.StatusInternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.Token, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		return nil, errors.StatusInternalServerError(err.Error())
	}

	return &result, nil
}

func (dbr *dbRepository) CreateAT(newAT access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(queryCreateAccessToken, newAT.Token, newAT.UserID, newAT.ClientID, newAT.Expires).Exec(); err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	return nil
}
func (dbr *dbRepository) UpdateExpirationTime(newAT access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(queryUpdateExpirationTime, newAT.Expires, newAT.Token).Exec(); err != nil {
		return errors.StatusInternalServerError(err.Error())
	}
	return nil
}
