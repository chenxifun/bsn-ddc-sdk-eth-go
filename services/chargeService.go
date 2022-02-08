package services

import (
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/chargeLogic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

var (
	Sigs_721  = []string{"0xd0def521", "0xf6dda936", "0x23b872dd", "0xb88d4fde", "0x42966c68"}
	Sigs_1155 = []string{"0xb55bc617", "0x63570355", "0xf242432a", "0x2eb2c2d6", "0x9dc29fac", "0xb2dc5dc3"}
)

type ChargeService struct {
	session *chargeLogic.ChargeLogicSession
}

func (s *ChargeService) Recharge(to string, value string) (txId string, err error) {

	v, ok := new(big.Int).SetString(value, 10)

	if !ok {
		return "", errors.Errorf("convert value [%s] feild ", value)
	}

	tx, err := s.session.Recharge(common.HexToAddress(to), v)

	if err != nil {
		return "", errors.WithMessage(err, "call charge contract Recharge has error")
	}
	txId = tx.Hash().String()

	return
}

func (s *ChargeService) BalanceOf(account string) string {

	value, _ := s.session.BalanceOf(common.HexToAddress(account))

	return value.String()
}

func (s *ChargeService) QueryFee(ddcAddr string, sig [4]byte) uint32 {

	amount, _ := s.session.QueryFee(common.HexToAddress(ddcAddr), sig)

	return amount
}
