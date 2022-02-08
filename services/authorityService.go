package services

import (
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/authorityLogic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func NewAuthorityService(session *authorityLogic.AuthorityLogicSession) *AuthorityService {

	return &AuthorityService{
		session: session,
	}

}

type AuthorityService struct {
	session *authorityLogic.AuthorityLogicSession
}

func (s *AuthorityService) AddAccountByOperator(account, accName, AccDID, leaderDid string) (txId string, err error) {

	tx, err := s.session.AddAccountByOperator(common.HexToAddress(account), accName, AccDID, leaderDid)
	if err != nil {
		return "", errors.WithMessage(err, "call authority contract AddAccountByOperator has error")
	}
	txId = tx.Hash().String()

	return
}
