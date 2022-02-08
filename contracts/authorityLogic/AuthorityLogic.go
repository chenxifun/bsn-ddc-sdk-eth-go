// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package authorityLogic

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

// AuthorityLogicABI is the input ABI used to generate the binding from.
const AuthorityLogicABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIAuthority.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"AddFunction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"CrossPlatformApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIAuthority.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"DelFunction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIAuthority.State\",\"name\":\"platformState\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumIAuthority.State\",\"name\":\"operatorState\",\"type\":\"uint8\"}],\"name\":\"UpdateAccountState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"accountName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountDID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"leaderDID\",\"type\":\"string\"}],\"name\":\"addAccountByOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIAuthority.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"addFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"accountName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountDID\",\"type\":\"string\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumIAuthority.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"checkAvailableAndRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"crossPlatformApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"crossPlatformCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIAuthority.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"delFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"enumIAuthority.Role\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"enumIAuthority.State\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"enumIAuthority.State\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIAuthority.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getFunctions\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"hasFunctionPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"acc2\",\"type\":\"address\"}],\"name\":\"onePlatformCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumIAuthority.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"changePlatformState\",\"type\":\"bool\"}],\"name\":\"updateAccountState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// AuthorityLogic is an auto generated Go binding around an Ethereum contract.
type AuthorityLogic struct {
	AuthorityLogicCaller     // Read-only binding to the contract
	AuthorityLogicTransactor // Write-only binding to the contract
	AuthorityLogicFilterer   // Log filterer for contract events
}

// AuthorityLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorityLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorityLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorityLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthorityLogicSession struct {
	Contract     *AuthorityLogic   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorityLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorityLogicCallerSession struct {
	Contract *AuthorityLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AuthorityLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorityLogicTransactorSession struct {
	Contract     *AuthorityLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AuthorityLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorityLogicRaw struct {
	Contract *AuthorityLogic // Generic contract binding to access the raw methods on
}

// AuthorityLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorityLogicCallerRaw struct {
	Contract *AuthorityLogicCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorityLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorityLogicTransactorRaw struct {
	Contract *AuthorityLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthorityLogic creates a new instance of AuthorityLogic, bound to a specific deployed contract.
func NewAuthorityLogic(address common.Address, backend bind.ContractBackend) (*AuthorityLogic, error) {
	contract, err := bindAuthorityLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogic{AuthorityLogicCaller: AuthorityLogicCaller{contract: contract}, AuthorityLogicTransactor: AuthorityLogicTransactor{contract: contract}, AuthorityLogicFilterer: AuthorityLogicFilterer{contract: contract}}, nil
}

// NewAuthorityLogicCaller creates a new read-only instance of AuthorityLogic, bound to a specific deployed contract.
func NewAuthorityLogicCaller(address common.Address, caller bind.ContractCaller) (*AuthorityLogicCaller, error) {
	contract, err := bindAuthorityLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicCaller{contract: contract}, nil
}

// NewAuthorityLogicTransactor creates a new write-only instance of AuthorityLogic, bound to a specific deployed contract.
func NewAuthorityLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorityLogicTransactor, error) {
	contract, err := bindAuthorityLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicTransactor{contract: contract}, nil
}

// NewAuthorityLogicFilterer creates a new log filterer instance of AuthorityLogic, bound to a specific deployed contract.
func NewAuthorityLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorityLogicFilterer, error) {
	contract, err := bindAuthorityLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicFilterer{contract: contract}, nil
}

// bindAuthorityLogic binds a generic wrapper to an already deployed contract.
func bindAuthorityLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthorityLogic *AuthorityLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthorityLogic.Contract.AuthorityLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthorityLogic *AuthorityLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AuthorityLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthorityLogic *AuthorityLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AuthorityLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthorityLogic *AuthorityLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthorityLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthorityLogic *AuthorityLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthorityLogic *AuthorityLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.contract.Transact(opts, method, params...)
}

// AccountAvailable is a free data retrieval call binding the contract method 0x95c2a8d9.
//
// Solidity: function accountAvailable(address account) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCaller) AccountAvailable(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "accountAvailable", account)
	return *ret0, err
}

// AccountAvailable is a free data retrieval call binding the contract method 0x95c2a8d9.
//
// Solidity: function accountAvailable(address account) view returns(bool)
func (_AuthorityLogic *AuthorityLogicSession) AccountAvailable(account common.Address) (bool, error) {
	return _AuthorityLogic.Contract.AccountAvailable(&_AuthorityLogic.CallOpts, account)
}

// AccountAvailable is a free data retrieval call binding the contract method 0x95c2a8d9.
//
// Solidity: function accountAvailable(address account) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCallerSession) AccountAvailable(account common.Address) (bool, error) {
	return _AuthorityLogic.Contract.AccountAvailable(&_AuthorityLogic.CallOpts, account)
}

// CheckAvailableAndRole is a free data retrieval call binding the contract method 0xed5cad64.
//
// Solidity: function checkAvailableAndRole(address account, uint8 role) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCaller) CheckAvailableAndRole(opts *bind.CallOpts, account common.Address, role uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "checkAvailableAndRole", account, role)
	return *ret0, err
}

// CheckAvailableAndRole is a free data retrieval call binding the contract method 0xed5cad64.
//
// Solidity: function checkAvailableAndRole(address account, uint8 role) view returns(bool)
func (_AuthorityLogic *AuthorityLogicSession) CheckAvailableAndRole(account common.Address, role uint8) (bool, error) {
	return _AuthorityLogic.Contract.CheckAvailableAndRole(&_AuthorityLogic.CallOpts, account, role)
}

// CheckAvailableAndRole is a free data retrieval call binding the contract method 0xed5cad64.
//
// Solidity: function checkAvailableAndRole(address account, uint8 role) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCallerSession) CheckAvailableAndRole(account common.Address, role uint8) (bool, error) {
	return _AuthorityLogic.Contract.CheckAvailableAndRole(&_AuthorityLogic.CallOpts, account, role)
}

// CrossPlatformCheck is a free data retrieval call binding the contract method 0x01843caa.
//
// Solidity: function crossPlatformCheck(address from, address to) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCaller) CrossPlatformCheck(opts *bind.CallOpts, from common.Address, to common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "crossPlatformCheck", from, to)
	return *ret0, err
}

// CrossPlatformCheck is a free data retrieval call binding the contract method 0x01843caa.
//
// Solidity: function crossPlatformCheck(address from, address to) view returns(bool)
func (_AuthorityLogic *AuthorityLogicSession) CrossPlatformCheck(from common.Address, to common.Address) (bool, error) {
	return _AuthorityLogic.Contract.CrossPlatformCheck(&_AuthorityLogic.CallOpts, from, to)
}

// CrossPlatformCheck is a free data retrieval call binding the contract method 0x01843caa.
//
// Solidity: function crossPlatformCheck(address from, address to) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCallerSession) CrossPlatformCheck(from common.Address, to common.Address) (bool, error) {
	return _AuthorityLogic.Contract.CrossPlatformCheck(&_AuthorityLogic.CallOpts, from, to)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns(string, string, uint8, string, uint8, uint8, string)
func (_AuthorityLogic *AuthorityLogicCaller) GetAccount(opts *bind.CallOpts, account common.Address) (string, string, uint8, string, uint8, uint8, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(uint8)
		ret3 = new(string)
		ret4 = new(uint8)
		ret5 = new(uint8)
		ret6 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _AuthorityLogic.contract.Call(opts, out, "getAccount", account)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns(string, string, uint8, string, uint8, uint8, string)
func (_AuthorityLogic *AuthorityLogicSession) GetAccount(account common.Address) (string, string, uint8, string, uint8, uint8, string, error) {
	return _AuthorityLogic.Contract.GetAccount(&_AuthorityLogic.CallOpts, account)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns(string, string, uint8, string, uint8, uint8, string)
func (_AuthorityLogic *AuthorityLogicCallerSession) GetAccount(account common.Address) (string, string, uint8, string, uint8, uint8, string, error) {
	return _AuthorityLogic.Contract.GetAccount(&_AuthorityLogic.CallOpts, account)
}

// GetFunctions is a free data retrieval call binding the contract method 0xb50f1a5e.
//
// Solidity: function getFunctions(uint8 role, address contractAddress) view returns(bytes4[])
func (_AuthorityLogic *AuthorityLogicCaller) GetFunctions(opts *bind.CallOpts, role uint8, contractAddress common.Address) ([][4]byte, error) {
	var (
		ret0 = new([][4]byte)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "getFunctions", role, contractAddress)
	return *ret0, err
}

// GetFunctions is a free data retrieval call binding the contract method 0xb50f1a5e.
//
// Solidity: function getFunctions(uint8 role, address contractAddress) view returns(bytes4[])
func (_AuthorityLogic *AuthorityLogicSession) GetFunctions(role uint8, contractAddress common.Address) ([][4]byte, error) {
	return _AuthorityLogic.Contract.GetFunctions(&_AuthorityLogic.CallOpts, role, contractAddress)
}

// GetFunctions is a free data retrieval call binding the contract method 0xb50f1a5e.
//
// Solidity: function getFunctions(uint8 role, address contractAddress) view returns(bytes4[])
func (_AuthorityLogic *AuthorityLogicCallerSession) GetFunctions(role uint8, contractAddress common.Address) ([][4]byte, error) {
	return _AuthorityLogic.Contract.GetFunctions(&_AuthorityLogic.CallOpts, role, contractAddress)
}

// HasFunctionPermission is a free data retrieval call binding the contract method 0x470cf0e0.
//
// Solidity: function hasFunctionPermission(address account, address contractAddress, bytes4 sig) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCaller) HasFunctionPermission(opts *bind.CallOpts, account common.Address, contractAddress common.Address, sig [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "hasFunctionPermission", account, contractAddress, sig)
	return *ret0, err
}

// HasFunctionPermission is a free data retrieval call binding the contract method 0x470cf0e0.
//
// Solidity: function hasFunctionPermission(address account, address contractAddress, bytes4 sig) view returns(bool)
func (_AuthorityLogic *AuthorityLogicSession) HasFunctionPermission(account common.Address, contractAddress common.Address, sig [4]byte) (bool, error) {
	return _AuthorityLogic.Contract.HasFunctionPermission(&_AuthorityLogic.CallOpts, account, contractAddress, sig)
}

// HasFunctionPermission is a free data retrieval call binding the contract method 0x470cf0e0.
//
// Solidity: function hasFunctionPermission(address account, address contractAddress, bytes4 sig) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCallerSession) HasFunctionPermission(account common.Address, contractAddress common.Address, sig [4]byte) (bool, error) {
	return _AuthorityLogic.Contract.HasFunctionPermission(&_AuthorityLogic.CallOpts, account, contractAddress, sig)
}

// OnePlatformCheck is a free data retrieval call binding the contract method 0x3a723d17.
//
// Solidity: function onePlatformCheck(address acc1, address acc2) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCaller) OnePlatformCheck(opts *bind.CallOpts, acc1 common.Address, acc2 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "onePlatformCheck", acc1, acc2)
	return *ret0, err
}

// OnePlatformCheck is a free data retrieval call binding the contract method 0x3a723d17.
//
// Solidity: function onePlatformCheck(address acc1, address acc2) view returns(bool)
func (_AuthorityLogic *AuthorityLogicSession) OnePlatformCheck(acc1 common.Address, acc2 common.Address) (bool, error) {
	return _AuthorityLogic.Contract.OnePlatformCheck(&_AuthorityLogic.CallOpts, acc1, acc2)
}

// OnePlatformCheck is a free data retrieval call binding the contract method 0x3a723d17.
//
// Solidity: function onePlatformCheck(address acc1, address acc2) view returns(bool)
func (_AuthorityLogic *AuthorityLogicCallerSession) OnePlatformCheck(acc1 common.Address, acc2 common.Address) (bool, error) {
	return _AuthorityLogic.Contract.OnePlatformCheck(&_AuthorityLogic.CallOpts, acc1, acc2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuthorityLogic *AuthorityLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuthorityLogic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuthorityLogic *AuthorityLogicSession) Owner() (common.Address, error) {
	return _AuthorityLogic.Contract.Owner(&_AuthorityLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuthorityLogic *AuthorityLogicCallerSession) Owner() (common.Address, error) {
	return _AuthorityLogic.Contract.Owner(&_AuthorityLogic.CallOpts)
}

// AddAccountByOperator is a paid mutator transaction binding the contract method 0xe3f00c3a.
//
// Solidity: function addAccountByOperator(address account, string accountName, string accountDID, string leaderDID) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) AddAccountByOperator(opts *bind.TransactOpts, account common.Address, accountName string, accountDID string, leaderDID string) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "addAccountByOperator", account, accountName, accountDID, leaderDID)
}

// AddAccountByOperator is a paid mutator transaction binding the contract method 0xe3f00c3a.
//
// Solidity: function addAccountByOperator(address account, string accountName, string accountDID, string leaderDID) returns()
func (_AuthorityLogic *AuthorityLogicSession) AddAccountByOperator(account common.Address, accountName string, accountDID string, leaderDID string) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AddAccountByOperator(&_AuthorityLogic.TransactOpts, account, accountName, accountDID, leaderDID)
}

// AddAccountByOperator is a paid mutator transaction binding the contract method 0xe3f00c3a.
//
// Solidity: function addAccountByOperator(address account, string accountName, string accountDID, string leaderDID) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) AddAccountByOperator(account common.Address, accountName string, accountDID string, leaderDID string) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AddAccountByOperator(&_AuthorityLogic.TransactOpts, account, accountName, accountDID, leaderDID)
}

// AddFunction is a paid mutator transaction binding the contract method 0xfaf3c53f.
//
// Solidity: function addFunction(uint8 role, address contractAddress, bytes4 sig) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) AddFunction(opts *bind.TransactOpts, role uint8, contractAddress common.Address, sig [4]byte) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "addFunction", role, contractAddress, sig)
}

// AddFunction is a paid mutator transaction binding the contract method 0xfaf3c53f.
//
// Solidity: function addFunction(uint8 role, address contractAddress, bytes4 sig) returns()
func (_AuthorityLogic *AuthorityLogicSession) AddFunction(role uint8, contractAddress common.Address, sig [4]byte) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AddFunction(&_AuthorityLogic.TransactOpts, role, contractAddress, sig)
}

// AddFunction is a paid mutator transaction binding the contract method 0xfaf3c53f.
//
// Solidity: function addFunction(uint8 role, address contractAddress, bytes4 sig) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) AddFunction(role uint8, contractAddress common.Address, sig [4]byte) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AddFunction(&_AuthorityLogic.TransactOpts, role, contractAddress, sig)
}

// AddOperator is a paid mutator transaction binding the contract method 0xe1576897.
//
// Solidity: function addOperator(address operator, string accountName, string accountDID) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address, accountName string, accountDID string) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "addOperator", operator, accountName, accountDID)
}

// AddOperator is a paid mutator transaction binding the contract method 0xe1576897.
//
// Solidity: function addOperator(address operator, string accountName, string accountDID) returns()
func (_AuthorityLogic *AuthorityLogicSession) AddOperator(operator common.Address, accountName string, accountDID string) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AddOperator(&_AuthorityLogic.TransactOpts, operator, accountName, accountDID)
}

// AddOperator is a paid mutator transaction binding the contract method 0xe1576897.
//
// Solidity: function addOperator(address operator, string accountName, string accountDID) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) AddOperator(operator common.Address, accountName string, accountDID string) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.AddOperator(&_AuthorityLogic.TransactOpts, operator, accountName, accountDID)
}

// CrossPlatformApproval is a paid mutator transaction binding the contract method 0x05e847fb.
//
// Solidity: function crossPlatformApproval(address from, address to, bool approved) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) CrossPlatformApproval(opts *bind.TransactOpts, from common.Address, to common.Address, approved bool) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "crossPlatformApproval", from, to, approved)
}

// CrossPlatformApproval is a paid mutator transaction binding the contract method 0x05e847fb.
//
// Solidity: function crossPlatformApproval(address from, address to, bool approved) returns()
func (_AuthorityLogic *AuthorityLogicSession) CrossPlatformApproval(from common.Address, to common.Address, approved bool) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.CrossPlatformApproval(&_AuthorityLogic.TransactOpts, from, to, approved)
}

// CrossPlatformApproval is a paid mutator transaction binding the contract method 0x05e847fb.
//
// Solidity: function crossPlatformApproval(address from, address to, bool approved) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) CrossPlatformApproval(from common.Address, to common.Address, approved bool) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.CrossPlatformApproval(&_AuthorityLogic.TransactOpts, from, to, approved)
}

// DelFunction is a paid mutator transaction binding the contract method 0xa6ff8f0c.
//
// Solidity: function delFunction(uint8 role, address contractAddress, bytes4 sig) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) DelFunction(opts *bind.TransactOpts, role uint8, contractAddress common.Address, sig [4]byte) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "delFunction", role, contractAddress, sig)
}

// DelFunction is a paid mutator transaction binding the contract method 0xa6ff8f0c.
//
// Solidity: function delFunction(uint8 role, address contractAddress, bytes4 sig) returns()
func (_AuthorityLogic *AuthorityLogicSession) DelFunction(role uint8, contractAddress common.Address, sig [4]byte) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.DelFunction(&_AuthorityLogic.TransactOpts, role, contractAddress, sig)
}

// DelFunction is a paid mutator transaction binding the contract method 0xa6ff8f0c.
//
// Solidity: function delFunction(uint8 role, address contractAddress, bytes4 sig) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) DelFunction(role uint8, contractAddress common.Address, sig [4]byte) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.DelFunction(&_AuthorityLogic.TransactOpts, role, contractAddress, sig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AuthorityLogic *AuthorityLogicTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AuthorityLogic *AuthorityLogicSession) Initialize() (*types.Transaction, error) {
	return _AuthorityLogic.Contract.Initialize(&_AuthorityLogic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) Initialize() (*types.Transaction, error) {
	return _AuthorityLogic.Contract.Initialize(&_AuthorityLogic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthorityLogic *AuthorityLogicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthorityLogic *AuthorityLogicSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuthorityLogic.Contract.RenounceOwnership(&_AuthorityLogic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuthorityLogic.Contract.RenounceOwnership(&_AuthorityLogic.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthorityLogic *AuthorityLogicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.TransferOwnership(&_AuthorityLogic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.TransferOwnership(&_AuthorityLogic.TransactOpts, newOwner)
}

// UpdateAccountState is a paid mutator transaction binding the contract method 0xbac7ad52.
//
// Solidity: function updateAccountState(address account, uint8 state, bool changePlatformState) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) UpdateAccountState(opts *bind.TransactOpts, account common.Address, state uint8, changePlatformState bool) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "updateAccountState", account, state, changePlatformState)
}

// UpdateAccountState is a paid mutator transaction binding the contract method 0xbac7ad52.
//
// Solidity: function updateAccountState(address account, uint8 state, bool changePlatformState) returns()
func (_AuthorityLogic *AuthorityLogicSession) UpdateAccountState(account common.Address, state uint8, changePlatformState bool) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.UpdateAccountState(&_AuthorityLogic.TransactOpts, account, state, changePlatformState)
}

// UpdateAccountState is a paid mutator transaction binding the contract method 0xbac7ad52.
//
// Solidity: function updateAccountState(address account, uint8 state, bool changePlatformState) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) UpdateAccountState(account common.Address, state uint8, changePlatformState bool) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.UpdateAccountState(&_AuthorityLogic.TransactOpts, account, state, changePlatformState)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AuthorityLogic *AuthorityLogicTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AuthorityLogic *AuthorityLogicSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.UpgradeTo(&_AuthorityLogic.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.UpgradeTo(&_AuthorityLogic.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuthorityLogic *AuthorityLogicTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuthorityLogic.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuthorityLogic *AuthorityLogicSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.UpgradeToAndCall(&_AuthorityLogic.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuthorityLogic *AuthorityLogicTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuthorityLogic.Contract.UpgradeToAndCall(&_AuthorityLogic.TransactOpts, newImplementation, data)
}

// AuthorityLogicAddAccountIterator is returned from FilterAddAccount and is used to iterate over the raw logs and unpacked data for AddAccount events raised by the AuthorityLogic contract.
type AuthorityLogicAddAccountIterator struct {
	Event *AuthorityLogicAddAccount // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicAddAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicAddAccount)
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
		it.Event = new(AuthorityLogicAddAccount)
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
func (it *AuthorityLogicAddAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicAddAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicAddAccount represents a AddAccount event raised by the AuthorityLogic contract.
type AuthorityLogicAddAccount struct {
	Caller  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddAccount is a free log retrieval operation binding the contract event 0xd246aa772c5574778b374ed1202c22bac9c87a04fd3624439c2642fa6ca62171.
//
// Solidity: event AddAccount(address indexed caller, address indexed account)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterAddAccount(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*AuthorityLogicAddAccountIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "AddAccount", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicAddAccountIterator{contract: _AuthorityLogic.contract, event: "AddAccount", logs: logs, sub: sub}, nil
}

// WatchAddAccount is a free log subscription operation binding the contract event 0xd246aa772c5574778b374ed1202c22bac9c87a04fd3624439c2642fa6ca62171.
//
// Solidity: event AddAccount(address indexed caller, address indexed account)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchAddAccount(opts *bind.WatchOpts, sink chan<- *AuthorityLogicAddAccount, caller []common.Address, account []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "AddAccount", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicAddAccount)
				if err := _AuthorityLogic.contract.UnpackLog(event, "AddAccount", log); err != nil {
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

// ParseAddAccount is a log parse operation binding the contract event 0xd246aa772c5574778b374ed1202c22bac9c87a04fd3624439c2642fa6ca62171.
//
// Solidity: event AddAccount(address indexed caller, address indexed account)
func (_AuthorityLogic *AuthorityLogicFilterer) ParseAddAccount(log types.Log) (*AuthorityLogicAddAccount, error) {
	event := new(AuthorityLogicAddAccount)
	if err := _AuthorityLogic.contract.UnpackLog(event, "AddAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicAddFunctionIterator is returned from FilterAddFunction and is used to iterate over the raw logs and unpacked data for AddFunction events raised by the AuthorityLogic contract.
type AuthorityLogicAddFunctionIterator struct {
	Event *AuthorityLogicAddFunction // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicAddFunctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicAddFunction)
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
		it.Event = new(AuthorityLogicAddFunction)
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
func (it *AuthorityLogicAddFunctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicAddFunctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicAddFunction represents a AddFunction event raised by the AuthorityLogic contract.
type AuthorityLogicAddFunction struct {
	Operator        common.Address
	Role            uint8
	ContractAddress common.Address
	Sig             [4]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddFunction is a free log retrieval operation binding the contract event 0xd2925d5c000f95a189be5c6c90c0dc14f33c077900241eb3e7c0a287d914e650.
//
// Solidity: event AddFunction(address indexed operator, uint8 indexed role, address contractAddress, bytes4 sig)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterAddFunction(opts *bind.FilterOpts, operator []common.Address, role []uint8) (*AuthorityLogicAddFunctionIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "AddFunction", operatorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicAddFunctionIterator{contract: _AuthorityLogic.contract, event: "AddFunction", logs: logs, sub: sub}, nil
}

// WatchAddFunction is a free log subscription operation binding the contract event 0xd2925d5c000f95a189be5c6c90c0dc14f33c077900241eb3e7c0a287d914e650.
//
// Solidity: event AddFunction(address indexed operator, uint8 indexed role, address contractAddress, bytes4 sig)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchAddFunction(opts *bind.WatchOpts, sink chan<- *AuthorityLogicAddFunction, operator []common.Address, role []uint8) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "AddFunction", operatorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicAddFunction)
				if err := _AuthorityLogic.contract.UnpackLog(event, "AddFunction", log); err != nil {
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

// ParseAddFunction is a log parse operation binding the contract event 0xd2925d5c000f95a189be5c6c90c0dc14f33c077900241eb3e7c0a287d914e650.
//
// Solidity: event AddFunction(address indexed operator, uint8 indexed role, address contractAddress, bytes4 sig)
func (_AuthorityLogic *AuthorityLogicFilterer) ParseAddFunction(log types.Log) (*AuthorityLogicAddFunction, error) {
	event := new(AuthorityLogicAddFunction)
	if err := _AuthorityLogic.contract.UnpackLog(event, "AddFunction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AuthorityLogic contract.
type AuthorityLogicAdminChangedIterator struct {
	Event *AuthorityLogicAdminChanged // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicAdminChanged)
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
		it.Event = new(AuthorityLogicAdminChanged)
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
func (it *AuthorityLogicAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicAdminChanged represents a AdminChanged event raised by the AuthorityLogic contract.
type AuthorityLogicAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AuthorityLogicAdminChangedIterator, error) {

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicAdminChangedIterator{contract: _AuthorityLogic.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AuthorityLogicAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicAdminChanged)
				if err := _AuthorityLogic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AuthorityLogic *AuthorityLogicFilterer) ParseAdminChanged(log types.Log) (*AuthorityLogicAdminChanged, error) {
	event := new(AuthorityLogicAdminChanged)
	if err := _AuthorityLogic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AuthorityLogic contract.
type AuthorityLogicBeaconUpgradedIterator struct {
	Event *AuthorityLogicBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicBeaconUpgraded)
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
		it.Event = new(AuthorityLogicBeaconUpgraded)
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
func (it *AuthorityLogicBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicBeaconUpgraded represents a BeaconUpgraded event raised by the AuthorityLogic contract.
type AuthorityLogicBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AuthorityLogicBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicBeaconUpgradedIterator{contract: _AuthorityLogic.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AuthorityLogicBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicBeaconUpgraded)
				if err := _AuthorityLogic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AuthorityLogic *AuthorityLogicFilterer) ParseBeaconUpgraded(log types.Log) (*AuthorityLogicBeaconUpgraded, error) {
	event := new(AuthorityLogicBeaconUpgraded)
	if err := _AuthorityLogic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicCrossPlatformApprovalIterator is returned from FilterCrossPlatformApproval and is used to iterate over the raw logs and unpacked data for CrossPlatformApproval events raised by the AuthorityLogic contract.
type AuthorityLogicCrossPlatformApprovalIterator struct {
	Event *AuthorityLogicCrossPlatformApproval // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicCrossPlatformApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicCrossPlatformApproval)
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
		it.Event = new(AuthorityLogicCrossPlatformApproval)
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
func (it *AuthorityLogicCrossPlatformApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicCrossPlatformApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicCrossPlatformApproval represents a CrossPlatformApproval event raised by the AuthorityLogic contract.
type AuthorityLogicCrossPlatformApproval struct {
	From     common.Address
	To       common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCrossPlatformApproval is a free log retrieval operation binding the contract event 0x944da3aaddcf895695572377f73f4b10c3822d1ea336c5d59f364f3e5e1b1218.
//
// Solidity: event CrossPlatformApproval(address indexed from, address indexed to, bool approved)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterCrossPlatformApproval(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AuthorityLogicCrossPlatformApprovalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "CrossPlatformApproval", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicCrossPlatformApprovalIterator{contract: _AuthorityLogic.contract, event: "CrossPlatformApproval", logs: logs, sub: sub}, nil
}

// WatchCrossPlatformApproval is a free log subscription operation binding the contract event 0x944da3aaddcf895695572377f73f4b10c3822d1ea336c5d59f364f3e5e1b1218.
//
// Solidity: event CrossPlatformApproval(address indexed from, address indexed to, bool approved)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchCrossPlatformApproval(opts *bind.WatchOpts, sink chan<- *AuthorityLogicCrossPlatformApproval, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "CrossPlatformApproval", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicCrossPlatformApproval)
				if err := _AuthorityLogic.contract.UnpackLog(event, "CrossPlatformApproval", log); err != nil {
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

// ParseCrossPlatformApproval is a log parse operation binding the contract event 0x944da3aaddcf895695572377f73f4b10c3822d1ea336c5d59f364f3e5e1b1218.
//
// Solidity: event CrossPlatformApproval(address indexed from, address indexed to, bool approved)
func (_AuthorityLogic *AuthorityLogicFilterer) ParseCrossPlatformApproval(log types.Log) (*AuthorityLogicCrossPlatformApproval, error) {
	event := new(AuthorityLogicCrossPlatformApproval)
	if err := _AuthorityLogic.contract.UnpackLog(event, "CrossPlatformApproval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicDelFunctionIterator is returned from FilterDelFunction and is used to iterate over the raw logs and unpacked data for DelFunction events raised by the AuthorityLogic contract.
type AuthorityLogicDelFunctionIterator struct {
	Event *AuthorityLogicDelFunction // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicDelFunctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicDelFunction)
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
		it.Event = new(AuthorityLogicDelFunction)
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
func (it *AuthorityLogicDelFunctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicDelFunctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicDelFunction represents a DelFunction event raised by the AuthorityLogic contract.
type AuthorityLogicDelFunction struct {
	Operator        common.Address
	Role            uint8
	ContractAddress common.Address
	Sig             [4]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelFunction is a free log retrieval operation binding the contract event 0xfd7321e06212c7e20cb40ffe695d160cb669329e947bbcab975e41e985153417.
//
// Solidity: event DelFunction(address indexed operator, uint8 indexed role, address contractAddress, bytes4 sig)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterDelFunction(opts *bind.FilterOpts, operator []common.Address, role []uint8) (*AuthorityLogicDelFunctionIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "DelFunction", operatorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicDelFunctionIterator{contract: _AuthorityLogic.contract, event: "DelFunction", logs: logs, sub: sub}, nil
}

// WatchDelFunction is a free log subscription operation binding the contract event 0xfd7321e06212c7e20cb40ffe695d160cb669329e947bbcab975e41e985153417.
//
// Solidity: event DelFunction(address indexed operator, uint8 indexed role, address contractAddress, bytes4 sig)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchDelFunction(opts *bind.WatchOpts, sink chan<- *AuthorityLogicDelFunction, operator []common.Address, role []uint8) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "DelFunction", operatorRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicDelFunction)
				if err := _AuthorityLogic.contract.UnpackLog(event, "DelFunction", log); err != nil {
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

// ParseDelFunction is a log parse operation binding the contract event 0xfd7321e06212c7e20cb40ffe695d160cb669329e947bbcab975e41e985153417.
//
// Solidity: event DelFunction(address indexed operator, uint8 indexed role, address contractAddress, bytes4 sig)
func (_AuthorityLogic *AuthorityLogicFilterer) ParseDelFunction(log types.Log) (*AuthorityLogicDelFunction, error) {
	event := new(AuthorityLogicDelFunction)
	if err := _AuthorityLogic.contract.UnpackLog(event, "DelFunction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuthorityLogic contract.
type AuthorityLogicOwnershipTransferredIterator struct {
	Event *AuthorityLogicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicOwnershipTransferred)
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
		it.Event = new(AuthorityLogicOwnershipTransferred)
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
func (it *AuthorityLogicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicOwnershipTransferred represents a OwnershipTransferred event raised by the AuthorityLogic contract.
type AuthorityLogicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuthorityLogicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicOwnershipTransferredIterator{contract: _AuthorityLogic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuthorityLogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicOwnershipTransferred)
				if err := _AuthorityLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuthorityLogic *AuthorityLogicFilterer) ParseOwnershipTransferred(log types.Log) (*AuthorityLogicOwnershipTransferred, error) {
	event := new(AuthorityLogicOwnershipTransferred)
	if err := _AuthorityLogic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicUpdateAccountStateIterator is returned from FilterUpdateAccountState and is used to iterate over the raw logs and unpacked data for UpdateAccountState events raised by the AuthorityLogic contract.
type AuthorityLogicUpdateAccountStateIterator struct {
	Event *AuthorityLogicUpdateAccountState // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicUpdateAccountStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicUpdateAccountState)
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
		it.Event = new(AuthorityLogicUpdateAccountState)
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
func (it *AuthorityLogicUpdateAccountStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicUpdateAccountStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicUpdateAccountState represents a UpdateAccountState event raised by the AuthorityLogic contract.
type AuthorityLogicUpdateAccountState struct {
	Account       common.Address
	PlatformState uint8
	OperatorState uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccountState is a free log retrieval operation binding the contract event 0x3a99d1f905eb66fd0850dd6e2156668b6ae1f16f41cd49dd7c6d4fb5147f784f.
//
// Solidity: event UpdateAccountState(address indexed account, uint8 platformState, uint8 operatorState)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterUpdateAccountState(opts *bind.FilterOpts, account []common.Address) (*AuthorityLogicUpdateAccountStateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "UpdateAccountState", accountRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicUpdateAccountStateIterator{contract: _AuthorityLogic.contract, event: "UpdateAccountState", logs: logs, sub: sub}, nil
}

// WatchUpdateAccountState is a free log subscription operation binding the contract event 0x3a99d1f905eb66fd0850dd6e2156668b6ae1f16f41cd49dd7c6d4fb5147f784f.
//
// Solidity: event UpdateAccountState(address indexed account, uint8 platformState, uint8 operatorState)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchUpdateAccountState(opts *bind.WatchOpts, sink chan<- *AuthorityLogicUpdateAccountState, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "UpdateAccountState", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicUpdateAccountState)
				if err := _AuthorityLogic.contract.UnpackLog(event, "UpdateAccountState", log); err != nil {
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

// ParseUpdateAccountState is a log parse operation binding the contract event 0x3a99d1f905eb66fd0850dd6e2156668b6ae1f16f41cd49dd7c6d4fb5147f784f.
//
// Solidity: event UpdateAccountState(address indexed account, uint8 platformState, uint8 operatorState)
func (_AuthorityLogic *AuthorityLogicFilterer) ParseUpdateAccountState(log types.Log) (*AuthorityLogicUpdateAccountState, error) {
	event := new(AuthorityLogicUpdateAccountState)
	if err := _AuthorityLogic.contract.UnpackLog(event, "UpdateAccountState", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorityLogicUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AuthorityLogic contract.
type AuthorityLogicUpgradedIterator struct {
	Event *AuthorityLogicUpgraded // Event containing the contract specifics and raw log

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
func (it *AuthorityLogicUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityLogicUpgraded)
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
		it.Event = new(AuthorityLogicUpgraded)
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
func (it *AuthorityLogicUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityLogicUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityLogicUpgraded represents a Upgraded event raised by the AuthorityLogic contract.
type AuthorityLogicUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuthorityLogic *AuthorityLogicFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuthorityLogicUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuthorityLogic.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityLogicUpgradedIterator{contract: _AuthorityLogic.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuthorityLogic *AuthorityLogicFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuthorityLogicUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuthorityLogic.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityLogicUpgraded)
				if err := _AuthorityLogic.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AuthorityLogic *AuthorityLogicFilterer) ParseUpgraded(log types.Log) (*AuthorityLogicUpgraded, error) {
	event := new(AuthorityLogicUpgraded)
	if err := _AuthorityLogic.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}
