package access_token

import (
	"strings"

	"github.com/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	CreateAT(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

//it is just databaserepo
type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	CreateAT(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

//struct for making methods
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.StatusBadRequestError("invalid access token id ")

	}
	accessToken, err := s.repository.GetById(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, err

}
func (s *service) CreateAT(at AccessToken) *errors.RestErr {
	at.Token = strings.TrimSpace(at.Token)
	if len(at.Token) == 0 {
		return errors.StatusBadRequestError("invalid access token id ")
	}
	err := s.repository.CreateAT(at)
	if err != nil {
		return errors.StatusBadRequestError("error while creating At")
	}

	return nil

}
func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
