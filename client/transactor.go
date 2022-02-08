package client

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strconv"
)

//1900000000
var gasPrice = new(big.Int).SetInt64(1000000000)

//var gasPrice = new(big.Int).SetInt64(1900000000)

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
func NewKeyedTransactor(key *ecdsa.PrivateKey, chainId *big.Int, gasPrice string, gaslimit string) *bind.TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	to := &bind.TransactOpts{
		//GasPrice: gasPrice,
		//GasLimit: 7863392,

		From: keyAddr,
		Signer: func(signer1 types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {

			signer := types.NewEIP155Signer(chainId)
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

	if gasPrice != "" {
		gp, ok := new(big.Int).SetString(gasPrice, 10)
		if ok {
			to.GasPrice = gp
		}
	}

	if gaslimit != "" {
		gl, err := strconv.ParseUint(gaslimit, 10, 64)
		if err != nil {
			to.GasLimit = gl
		}

	}

	return to
}
