package services

import (
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/DDC1155Logic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

type DDC1155Service struct {
	session *DDC1155Logic.DDC1155LogicSession
}

func (s *DDC1155Service) SafeMint(to string, amount *big.Int, uri string, data []byte) (txId string, err error) {

	tx, err := s.session.SafeMint(common.HexToAddress(to), amount, uri, data)
	if err != nil {
		return "", errors.WithMessage(err, "call ddc 1155 contract SafeMint has error")
	}

	txId = tx.Hash().String()
	return
}

func (s *DDC1155Service) SafeMintBatch(to string, amount []*big.Int, uri []string, data []byte) (txId string, err error) {

	tx, err := s.session.SafeMintBatch(common.HexToAddress(to), amount, uri, data)
	if err != nil {
		return "", errors.WithMessage(err, "call ddc 1155 contract SafeMintBatch has error")
	}

	txId = tx.Hash().String()
	return
}
