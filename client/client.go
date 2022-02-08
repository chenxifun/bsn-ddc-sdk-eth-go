package client

import (
	"context"
	"crypto/ecdsa"
	"github.com/chenxifun/bsn-ddc-sdk-eth/config"
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/DDC1155Logic"
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/DDC721Logic"
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/authorityLogic"
	"github.com/chenxifun/bsn-ddc-sdk-eth/contracts/chargeLogic"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

func NewDDCClient(conf *config.Config, hexKey string) (*DDCClient, error) {
	key, err := NewHexKey(hexKey)
	if err != nil {
		return nil, err
	}
	return NewDDCClientByKey(conf, key)
}

func NewDDCClientByKey(conf *config.Config, key *ecdsa.PrivateKey) (*DDCClient, error) {

	ethCli, err := ethclient.Dial(conf.NodeUrl)
	if err != nil {
		return nil, errors.WithMessage(err, "creates eth client has error")
	}

	chainId, err := ethCli.ChainID(context.Background())
	if err != nil {
		return nil, errors.WithMessage(err, "query eth chain Id has error")
	}

	transactOpts := NewKeyedTransactor(key, chainId, conf.GasPrice, conf.GasLimit)

	cli := &DDCClient{
		client:       ethCli,
		conf:         conf,
		transactOpts: transactOpts,
		callOpts:     bind.CallOpts{From: transactOpts.From},
	}

	return cli, nil
}

func NewHexKey(hexKey string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := hexutil.Decode(hexKey)
	if err != nil {
		return nil, errors.WithMessagef(err, "decode hex key [%s] has error", hexKey)
	}

	pk, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, errors.WithMessage(err, "creates a private key has error")
	}

	return pk, nil
}

type DDCClient struct {
	client *ethclient.Client
	conf   *config.Config

	transactOpts *bind.TransactOpts
	callOpts     bind.CallOpts

	//chainId      *big.Int

	authoritySession *authorityLogic.AuthorityLogicSession
	chargeSession    *chargeLogic.ChargeLogicSession
	ddc721Session    *DDC721Logic.DDC721LogicSession
	ddc1155Session   *DDC1155Logic.DDC1155LogicSession
}

func (cli *DDCClient) ETHClient() *ethclient.Client {
	return cli.client
}

func (cli *DDCClient) NewAuthoritysession() (*authorityLogic.AuthorityLogicSession, error) {

	addr := common.HexToAddress(cli.conf.AuthorityAddress)
	contract, err := authorityLogic.NewAuthorityLogic(addr, cli.client)
	if err != nil {
		return nil, errors.WithMessagef(err, "new authority contract [%s] has error", cli.conf.AuthorityAddress)
	}

	return &authorityLogic.AuthorityLogicSession{contract, cli.callOpts, *cli.transactOpts}, nil
}

func (cli *DDCClient) NewChargeSession() (*chargeLogic.ChargeLogicSession, error) {

	addr := common.HexToAddress(cli.conf.ChargeAddress)
	session, err := chargeLogic.NewChargeLogic(addr, cli.client)
	if err != nil {
		return nil, errors.WithMessagef(err, "new charge contract [%s] has error", cli.conf.ChargeAddress)
	}
	return &chargeLogic.ChargeLogicSession{session, cli.callOpts, *cli.transactOpts}, err
}

func (cli *DDCClient) NewDDC721Session() (*DDC721Logic.DDC721LogicSession, error) {

	addr := common.HexToAddress(cli.conf.Ddc721Address)
	session, err := DDC721Logic.NewDDC721Logic(addr, cli.client)
	if err != nil {
		return nil, errors.WithMessagef(err, "new ddc 1155 contract [%s] has error", cli.conf.Ddc721Address)
	}
	return &DDC721Logic.DDC721LogicSession{session, cli.callOpts, *cli.transactOpts}, err
}

func (cli *DDCClient) NewDDC1155Session() (*DDC1155Logic.DDC1155LogicSession, error) {

	addr := common.HexToAddress(cli.conf.Ddc1155Address)
	session, err := DDC1155Logic.NewDDC1155Logic(addr, cli.client)
	if err != nil {
		return nil, errors.WithMessagef(err, "new ddc 1155 contract [%s] has error", cli.conf.Ddc1155Address)
	}
	return &DDC1155Logic.DDC1155LogicSession{session, cli.callOpts, *cli.transactOpts}, err
}
