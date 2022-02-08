package services

import (
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/DDC721Logic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
)

type DDC721Service struct {
	session *DDC721Logic.DDC721LogicSession
	client  *ethclient.Client
}

func (s *DDC721Service) SafeMint(to string, uri string, data []byte) (txId string, err error) {

	tx, err := s.session.SafeMint(common.HexToAddress(to), uri, data)
	if err != nil {
		return "", errors.WithMessage(err, "call ddc 721 contract SafeMint has error")
	}

	txId = tx.Hash().String()
	return
}

func (s *DDC721Service) Mint(to string, uri string) (txId string, err error) {

	tx, err := s.session.Mint(common.HexToAddress(to), uri)
	if err != nil {
		return "", errors.WithMessage(err, "call ddc 721 contract Mint has error")
	}

	txId = tx.Hash().String()
	return
}

func (s *DDC721Service) TransferFrom(from, to string, ddcId *big.Int) (txId string, err error) {

	tx, err := s.session.TransferFrom(common.HexToAddress(from), common.HexToAddress(to), ddcId)
	if err != nil {
		return "", errors.WithMessage(err, "call ddc 721 contract TransferFrom has error")
	}

	txId = tx.Hash().String()
	return
}
