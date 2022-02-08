// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chargeLogic

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChargeLogicABI is the input ABI used to generate the binding from.
const ChargeLogicABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"}],\"name\":\"DelDDC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"DelFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accAddr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"}],\"name\":\"delDDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"delFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"queryFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"selfRecharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorityProxyAddress\",\"type\":\"address\"}],\"name\":\"setAuthorityProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ChargeLogic is an auto generated Go binding around an Ethereum contract.
type ChargeLogic struct {
	ChargeLogicCaller     // Read-only binding to the contract
	ChargeLogicTransactor // Write-only binding to the contract
	ChargeLogicFilterer   // Log filterer for contract events
}

// ChargeLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChargeLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChargeLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChargeLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChargeLogicSession struct {
	Contract     *ChargeLogic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChargeLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChargeLogicCallerSession struct {
	Contract *ChargeLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChargeLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChargeLogicTransactorSession struct {
	Contract     *ChargeLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChargeLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChargeLogicRaw struct {
	Contract *ChargeLogic // Generic contract binding to access the raw methods on
}

// ChargeLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChargeLogicCallerRaw struct {
	Contract *ChargeLogicCaller // Generic read-only contract binding to access the raw methods on
}

// ChargeLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChargeLogicTransactorRaw struct {
	Contract *ChargeLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChargeLogic creates a new instance of ChargeLogic, bound to a specific deployed contract.
func NewChargeLogic(address common.Address, backend bind.ContractBackend) (*ChargeLogic, error) {
	contract, err := bindChargeLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChargeLogic{ChargeLogicCaller: ChargeLogicCaller{contract: contract}, ChargeLogicTransactor: ChargeLogicTransactor{contract: contract}, ChargeLogicFilterer: ChargeLogicFilterer{contract: contract}}, nil
}

// NewChargeLogicCaller creates a new read-only instance of ChargeLogic, bound to a specific deployed contract.
func NewChargeLogicCaller(address common.Address, caller bind.ContractCaller) (*ChargeLogicCaller, error) {
	contract, err := bindChargeLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicCaller{contract: contract}, nil
}

// NewChargeLogicTransactor creates a new write-only instance of ChargeLogic, bound to a specific deployed contract.
func NewChargeLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*ChargeLogicTransactor, error) {
	contract, err := bindChargeLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicTransactor{contract: contract}, nil
}

// NewChargeLogicFilterer creates a new log filterer instance of ChargeLogic, bound to a specific deployed contract.
func NewChargeLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*ChargeLogicFilterer, error) {
	contract, err := bindChargeLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicFilterer{contract: contract}, nil
}

// bindChargeLogic binds a generic wrapper to an already deployed contract.
func bindChargeLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChargeLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChargeLogic *ChargeLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChargeLogic.Contract.ChargeLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChargeLogic *ChargeLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargeLogic.Contract.ChargeLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChargeLogic *ChargeLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChargeLogic.Contract.ChargeLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChargeLogic *ChargeLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChargeLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChargeLogic *ChargeLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargeLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChargeLogic *ChargeLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChargeLogic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_ChargeLogic *ChargeLogicCaller) BalanceOf(opts *bind.CallOpts, accAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChargeLogic.contract.Call(opts, out, "balanceOf", accAddr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_ChargeLogic *ChargeLogicSession) BalanceOf(accAddr common.Address) (*big.Int, error) {
	return _ChargeLogic.Contract.BalanceOf(&_ChargeLogic.CallOpts, accAddr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_ChargeLogic *ChargeLogicCallerSession) BalanceOf(accAddr common.Address) (*big.Int, error) {
	return _ChargeLogic.Contract.BalanceOf(&_ChargeLogic.CallOpts, accAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChargeLogic *ChargeLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChargeLogic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChargeLogic *ChargeLogicSession) Owner() (common.Address, error) {
	return _ChargeLogic.Contract.Owner(&_ChargeLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChargeLogic *ChargeLogicCallerSession) Owner() (common.Address, error) {
	return _ChargeLogic.Contract.Owner(&_ChargeLogic.CallOpts)
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_ChargeLogic *ChargeLogicCaller) QueryFee(opts *bind.CallOpts, ddcAddr common.Address, sig [4]byte) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ChargeLogic.contract.Call(opts, out, "queryFee", ddcAddr, sig)
	return *ret0, err
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_ChargeLogic *ChargeLogicSession) QueryFee(ddcAddr common.Address, sig [4]byte) (uint32, error) {
	return _ChargeLogic.Contract.QueryFee(&_ChargeLogic.CallOpts, ddcAddr, sig)
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_ChargeLogic *ChargeLogicCallerSession) QueryFee(ddcAddr common.Address, sig [4]byte) (uint32, error) {
	return _ChargeLogic.Contract.QueryFee(&_ChargeLogic.CallOpts, ddcAddr, sig)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChargeLogic *ChargeLogicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChargeLogic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChargeLogic *ChargeLogicSession) TotalSupply() (*big.Int, error) {
	return _ChargeLogic.Contract.TotalSupply(&_ChargeLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChargeLogic *ChargeLogicCallerSession) TotalSupply() (*big.Int, error) {
	return _ChargeLogic.Contract.TotalSupply(&_ChargeLogic.CallOpts)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_ChargeLogic *ChargeLogicTransactor) DelDDC(opts *bind.TransactOpts, ddcAddr common.Address) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "delDDC", ddcAddr)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_ChargeLogic *ChargeLogicSession) DelDDC(ddcAddr common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.DelDDC(&_ChargeLogic.TransactOpts, ddcAddr)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) DelDDC(ddcAddr common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.DelDDC(&_ChargeLogic.TransactOpts, ddcAddr)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_ChargeLogic *ChargeLogicTransactor) DelFee(opts *bind.TransactOpts, ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "delFee", ddcAddr, sig)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_ChargeLogic *ChargeLogicSession) DelFee(ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _ChargeLogic.Contract.DelFee(&_ChargeLogic.TransactOpts, ddcAddr, sig)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) DelFee(ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _ChargeLogic.Contract.DelFee(&_ChargeLogic.TransactOpts, ddcAddr, sig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChargeLogic *ChargeLogicTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChargeLogic *ChargeLogicSession) Initialize() (*types.Transaction, error) {
	return _ChargeLogic.Contract.Initialize(&_ChargeLogic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChargeLogic *ChargeLogicTransactorSession) Initialize() (*types.Transaction, error) {
	return _ChargeLogic.Contract.Initialize(&_ChargeLogic.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_ChargeLogic *ChargeLogicTransactor) Pay(opts *bind.TransactOpts, payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "pay", payer, sig, ddcId)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_ChargeLogic *ChargeLogicSession) Pay(payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.Pay(&_ChargeLogic.TransactOpts, payer, sig, ddcId)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) Pay(payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.Pay(&_ChargeLogic.TransactOpts, payer, sig, ddcId)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_ChargeLogic *ChargeLogicTransactor) Recharge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "recharge", to, amount)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_ChargeLogic *ChargeLogicSession) Recharge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.Recharge(&_ChargeLogic.TransactOpts, to, amount)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) Recharge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.Recharge(&_ChargeLogic.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChargeLogic *ChargeLogicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChargeLogic *ChargeLogicSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChargeLogic.Contract.RenounceOwnership(&_ChargeLogic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChargeLogic *ChargeLogicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChargeLogic.Contract.RenounceOwnership(&_ChargeLogic.TransactOpts)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_ChargeLogic *ChargeLogicTransactor) SelfRecharge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "selfRecharge", amount)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_ChargeLogic *ChargeLogicSession) SelfRecharge(amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.SelfRecharge(&_ChargeLogic.TransactOpts, amount)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) SelfRecharge(amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.SelfRecharge(&_ChargeLogic.TransactOpts, amount)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_ChargeLogic *ChargeLogicTransactor) SetAuthorityProxyAddress(opts *bind.TransactOpts, authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "setAuthorityProxyAddress", authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_ChargeLogic *ChargeLogicSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.SetAuthorityProxyAddress(&_ChargeLogic.TransactOpts, authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.SetAuthorityProxyAddress(&_ChargeLogic.TransactOpts, authorityProxyAddress)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_ChargeLogic *ChargeLogicTransactor) SetFee(opts *bind.TransactOpts, ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "setFee", ddcAddr, sig, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_ChargeLogic *ChargeLogicSession) SetFee(ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _ChargeLogic.Contract.SetFee(&_ChargeLogic.TransactOpts, ddcAddr, sig, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) SetFee(ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _ChargeLogic.Contract.SetFee(&_ChargeLogic.TransactOpts, ddcAddr, sig, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_ChargeLogic *ChargeLogicTransactor) Settlement(opts *bind.TransactOpts, ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "settlement", ddcAddr, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_ChargeLogic *ChargeLogicSession) Settlement(ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.Settlement(&_ChargeLogic.TransactOpts, ddcAddr, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) Settlement(ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ChargeLogic.Contract.Settlement(&_ChargeLogic.TransactOpts, ddcAddr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChargeLogic *ChargeLogicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChargeLogic *ChargeLogicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.TransferOwnership(&_ChargeLogic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.TransferOwnership(&_ChargeLogic.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ChargeLogic *ChargeLogicTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ChargeLogic *ChargeLogicSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.UpgradeTo(&_ChargeLogic.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ChargeLogic *ChargeLogicTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ChargeLogic.Contract.UpgradeTo(&_ChargeLogic.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChargeLogic *ChargeLogicTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChargeLogic.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChargeLogic *ChargeLogicSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChargeLogic.Contract.UpgradeToAndCall(&_ChargeLogic.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChargeLogic *ChargeLogicTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChargeLogic.Contract.UpgradeToAndCall(&_ChargeLogic.TransactOpts, newImplementation, data)
}

// ChargeLogicAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ChargeLogic contract.
type ChargeLogicAdminChangedIterator struct {
	Event *ChargeLogicAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicAdminChanged represents a AdminChanged event raised by the ChargeLogic contract.
type ChargeLogicAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ChargeLogic *ChargeLogicFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ChargeLogicAdminChangedIterator, error) {

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ChargeLogicAdminChangedIterator{contract: _ChargeLogic.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ChargeLogic *ChargeLogicFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ChargeLogicAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicAdminChanged)
				if err := _ChargeLogic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ChargeLogic *ChargeLogicFilterer) ParseAdminChanged(log types.Log) (*ChargeLogicAdminChanged, error) {
	event := new(ChargeLogicAdminChanged)
	if err := _ChargeLogic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ChargeLogic contract.
type ChargeLogicBeaconUpgradedIterator struct {
	Event *ChargeLogicBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicBeaconUpgraded represents a BeaconUpgraded event raised by the ChargeLogic contract.
type ChargeLogicBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ChargeLogic *ChargeLogicFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ChargeLogicBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicBeaconUpgradedIterator{contract: _ChargeLogic.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ChargeLogic *ChargeLogicFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ChargeLogicBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicBeaconUpgraded)
				if err := _ChargeLogic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ChargeLogic *ChargeLogicFilterer) ParseBeaconUpgraded(log types.Log) (*ChargeLogicBeaconUpgraded, error) {
	event := new(ChargeLogicBeaconUpgraded)
	if err := _ChargeLogic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicDelDDCIterator is returned from FilterDelDDC and is used to iterate over the raw logs and unpacked data for DelDDC events raised by the ChargeLogic contract.
type ChargeLogicDelDDCIterator struct {
	Event *ChargeLogicDelDDC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicDelDDCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicDelDDC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicDelDDC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicDelDDCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicDelDDCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicDelDDC represents a DelDDC event raised by the ChargeLogic contract.
type ChargeLogicDelDDC struct {
	DdcAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelDDC is a free log retrieval operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_ChargeLogic *ChargeLogicFilterer) FilterDelDDC(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeLogicDelDDCIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "DelDDC", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicDelDDCIterator{contract: _ChargeLogic.contract, event: "DelDDC", logs: logs, sub: sub}, nil
}

// WatchDelDDC is a free log subscription operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_ChargeLogic *ChargeLogicFilterer) WatchDelDDC(opts *bind.WatchOpts, sink chan<- *ChargeLogicDelDDC, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "DelDDC", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicDelDDC)
				if err := _ChargeLogic.contract.UnpackLog(event, "DelDDC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelDDC is a log parse operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_ChargeLogic *ChargeLogicFilterer) ParseDelDDC(log types.Log) (*ChargeLogicDelDDC, error) {
	event := new(ChargeLogicDelDDC)
	if err := _ChargeLogic.contract.UnpackLog(event, "DelDDC", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicDelFeeIterator is returned from FilterDelFee and is used to iterate over the raw logs and unpacked data for DelFee events raised by the ChargeLogic contract.
type ChargeLogicDelFeeIterator struct {
	Event *ChargeLogicDelFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicDelFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicDelFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicDelFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicDelFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicDelFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicDelFee represents a DelFee event raised by the ChargeLogic contract.
type ChargeLogicDelFee struct {
	DdcAddr common.Address
	Sig     [4]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelFee is a free log retrieval operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_ChargeLogic *ChargeLogicFilterer) FilterDelFee(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeLogicDelFeeIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "DelFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicDelFeeIterator{contract: _ChargeLogic.contract, event: "DelFee", logs: logs, sub: sub}, nil
}

// WatchDelFee is a free log subscription operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_ChargeLogic *ChargeLogicFilterer) WatchDelFee(opts *bind.WatchOpts, sink chan<- *ChargeLogicDelFee, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "DelFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicDelFee)
				if err := _ChargeLogic.contract.UnpackLog(event, "DelFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelFee is a log parse operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_ChargeLogic *ChargeLogicFilterer) ParseDelFee(log types.Log) (*ChargeLogicDelFee, error) {
	event := new(ChargeLogicDelFee)
	if err := _ChargeLogic.contract.UnpackLog(event, "DelFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChargeLogic contract.
type ChargeLogicOwnershipTransferredIterator struct {
	Event *ChargeLogicOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicOwnershipTransferred represents a OwnershipTransferred event raised by the ChargeLogic contract.
type ChargeLogicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChargeLogic *ChargeLogicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChargeLogicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicOwnershipTransferredIterator{contract: _ChargeLogic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChargeLogic *ChargeLogicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChargeLogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicOwnershipTransferred)
				if err := _ChargeLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChargeLogic *ChargeLogicFilterer) ParseOwnershipTransferred(log types.Log) (*ChargeLogicOwnershipTransferred, error) {
	event := new(ChargeLogicOwnershipTransferred)
	if err := _ChargeLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the ChargeLogic contract.
type ChargeLogicPayIterator struct {
	Event *ChargeLogicPay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicPay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicPay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicPay represents a Pay event raised by the ChargeLogic contract.
type ChargeLogicPay struct {
	Payer  common.Address
	Payee  common.Address
	Sig    [4]byte
	Amount uint32
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_ChargeLogic *ChargeLogicFilterer) FilterPay(opts *bind.FilterOpts, payer []common.Address, payee []common.Address) (*ChargeLogicPayIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "Pay", payerRule, payeeRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicPayIterator{contract: _ChargeLogic.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_ChargeLogic *ChargeLogicFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *ChargeLogicPay, payer []common.Address, payee []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "Pay", payerRule, payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicPay)
				if err := _ChargeLogic.contract.UnpackLog(event, "Pay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePay is a log parse operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_ChargeLogic *ChargeLogicFilterer) ParsePay(log types.Log) (*ChargeLogicPay, error) {
	event := new(ChargeLogicPay)
	if err := _ChargeLogic.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the ChargeLogic contract.
type ChargeLogicRechargeIterator struct {
	Event *ChargeLogicRecharge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicRecharge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicRecharge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicRecharge represents a Recharge event raised by the ChargeLogic contract.
type ChargeLogicRecharge struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_ChargeLogic *ChargeLogicFilterer) FilterRecharge(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChargeLogicRechargeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "Recharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicRechargeIterator{contract: _ChargeLogic.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_ChargeLogic *ChargeLogicFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *ChargeLogicRecharge, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "Recharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicRecharge)
				if err := _ChargeLogic.contract.UnpackLog(event, "Recharge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecharge is a log parse operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_ChargeLogic *ChargeLogicFilterer) ParseRecharge(log types.Log) (*ChargeLogicRecharge, error) {
	event := new(ChargeLogicRecharge)
	if err := _ChargeLogic.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the ChargeLogic contract.
type ChargeLogicSetFeeIterator struct {
	Event *ChargeLogicSetFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicSetFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicSetFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicSetFee represents a SetFee event raised by the ChargeLogic contract.
type ChargeLogicSetFee struct {
	DdcAddr common.Address
	Sig     [4]byte
	Amount  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_ChargeLogic *ChargeLogicFilterer) FilterSetFee(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeLogicSetFeeIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "SetFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicSetFeeIterator{contract: _ChargeLogic.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_ChargeLogic *ChargeLogicFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *ChargeLogicSetFee, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "SetFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicSetFee)
				if err := _ChargeLogic.contract.UnpackLog(event, "SetFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFee is a log parse operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_ChargeLogic *ChargeLogicFilterer) ParseSetFee(log types.Log) (*ChargeLogicSetFee, error) {
	event := new(ChargeLogicSetFee)
	if err := _ChargeLogic.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the ChargeLogic contract.
type ChargeLogicSettlementIterator struct {
	Event *ChargeLogicSettlement // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicSettlement)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicSettlement)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicSettlement represents a Settlement event raised by the ChargeLogic contract.
type ChargeLogicSettlement struct {
	AccAddr common.Address
	DdcAddr common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_ChargeLogic *ChargeLogicFilterer) FilterSettlement(opts *bind.FilterOpts, accAddr []common.Address, ddcAddr []common.Address) (*ChargeLogicSettlementIterator, error) {

	var accAddrRule []interface{}
	for _, accAddrItem := range accAddr {
		accAddrRule = append(accAddrRule, accAddrItem)
	}
	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "Settlement", accAddrRule, ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicSettlementIterator{contract: _ChargeLogic.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_ChargeLogic *ChargeLogicFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *ChargeLogicSettlement, accAddr []common.Address, ddcAddr []common.Address) (event.Subscription, error) {

	var accAddrRule []interface{}
	for _, accAddrItem := range accAddr {
		accAddrRule = append(accAddrRule, accAddrItem)
	}
	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "Settlement", accAddrRule, ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicSettlement)
				if err := _ChargeLogic.contract.UnpackLog(event, "Settlement", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlement is a log parse operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_ChargeLogic *ChargeLogicFilterer) ParseSettlement(log types.Log) (*ChargeLogicSettlement, error) {
	event := new(ChargeLogicSettlement)
	if err := _ChargeLogic.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChargeLogicUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ChargeLogic contract.
type ChargeLogicUpgradedIterator struct {
	Event *ChargeLogicUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeLogicUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeLogicUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeLogicUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeLogicUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeLogicUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeLogicUpgraded represents a Upgraded event raised by the ChargeLogic contract.
type ChargeLogicUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChargeLogic *ChargeLogicFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChargeLogicUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChargeLogic.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChargeLogicUpgradedIterator{contract: _ChargeLogic.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChargeLogic *ChargeLogicFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChargeLogicUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChargeLogic.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeLogicUpgraded)
				if err := _ChargeLogic.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChargeLogic *ChargeLogicFilterer) ParseUpgraded(log types.Log) (*ChargeLogicUpgraded, error) {
	event := new(ChargeLogicUpgraded)
	if err := _ChargeLogic.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}
