// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DDC721Logic

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

// DDC721LogicABI is the input ABI used to generate the binding from.
const DDC721LogicABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"EnterBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ExitBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetNameAndSymbol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ddcURI\",\"type\":\"string\"}],\"name\":\"SetURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ddcURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ddcURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorityProxyAddress\",\"type\":\"address\"}],\"name\":\"setAuthorityProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"chargeProxyAddress\",\"type\":\"address\"}],\"name\":\"setChargeProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setNameAndSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DDC721Logic is an auto generated Go binding around an Ethereum contract.
type DDC721Logic struct {
	DDC721LogicCaller     // Read-only binding to the contract
	DDC721LogicTransactor // Write-only binding to the contract
	DDC721LogicFilterer   // Log filterer for contract events
}

// DDC721LogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type DDC721LogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC721LogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DDC721LogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC721LogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DDC721LogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC721LogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DDC721LogicSession struct {
	Contract     *DDC721Logic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDC721LogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DDC721LogicCallerSession struct {
	Contract *DDC721LogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DDC721LogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DDC721LogicTransactorSession struct {
	Contract     *DDC721LogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DDC721LogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type DDC721LogicRaw struct {
	Contract *DDC721Logic // Generic contract binding to access the raw methods on
}

// DDC721LogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DDC721LogicCallerRaw struct {
	Contract *DDC721LogicCaller // Generic read-only contract binding to access the raw methods on
}

// DDC721LogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DDC721LogicTransactorRaw struct {
	Contract *DDC721LogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDDC721Logic creates a new instance of DDC721Logic, bound to a specific deployed contract.
func NewDDC721Logic(address common.Address, backend bind.ContractBackend) (*DDC721Logic, error) {
	contract, err := bindDDC721Logic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DDC721Logic{DDC721LogicCaller: DDC721LogicCaller{contract: contract}, DDC721LogicTransactor: DDC721LogicTransactor{contract: contract}, DDC721LogicFilterer: DDC721LogicFilterer{contract: contract}}, nil
}

// NewDDC721LogicCaller creates a new read-only instance of DDC721Logic, bound to a specific deployed contract.
func NewDDC721LogicCaller(address common.Address, caller bind.ContractCaller) (*DDC721LogicCaller, error) {
	contract, err := bindDDC721Logic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicCaller{contract: contract}, nil
}

// NewDDC721LogicTransactor creates a new write-only instance of DDC721Logic, bound to a specific deployed contract.
func NewDDC721LogicTransactor(address common.Address, transactor bind.ContractTransactor) (*DDC721LogicTransactor, error) {
	contract, err := bindDDC721Logic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicTransactor{contract: contract}, nil
}

// NewDDC721LogicFilterer creates a new log filterer instance of DDC721Logic, bound to a specific deployed contract.
func NewDDC721LogicFilterer(address common.Address, filterer bind.ContractFilterer) (*DDC721LogicFilterer, error) {
	contract, err := bindDDC721Logic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicFilterer{contract: contract}, nil
}

// bindDDC721Logic binds a generic wrapper to an already deployed contract.
func bindDDC721Logic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DDC721LogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDC721Logic *DDC721LogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DDC721Logic.Contract.DDC721LogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDC721Logic *DDC721LogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721Logic.Contract.DDC721LogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDC721Logic *DDC721LogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDC721Logic.Contract.DDC721LogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDC721Logic *DDC721LogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DDC721Logic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDC721Logic *DDC721LogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721Logic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDC721Logic *DDC721LogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDC721Logic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DDC721Logic *DDC721LogicCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DDC721Logic *DDC721LogicSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DDC721Logic.Contract.BalanceOf(&_DDC721Logic.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DDC721Logic *DDC721LogicCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DDC721Logic.Contract.BalanceOf(&_DDC721Logic.CallOpts, owner)
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC721Logic *DDC721LogicCaller) DdcURI(opts *bind.CallOpts, ddcId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "ddcURI", ddcId)
	return *ret0, err
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC721Logic *DDC721LogicSession) DdcURI(ddcId *big.Int) (string, error) {
	return _DDC721Logic.Contract.DdcURI(&_DDC721Logic.CallOpts, ddcId)
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC721Logic *DDC721LogicCallerSession) DdcURI(ddcId *big.Int) (string, error) {
	return _DDC721Logic.Contract.DdcURI(&_DDC721Logic.CallOpts, ddcId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ddcId) view returns(address)
func (_DDC721Logic *DDC721LogicCaller) GetApproved(opts *bind.CallOpts, ddcId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "getApproved", ddcId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ddcId) view returns(address)
func (_DDC721Logic *DDC721LogicSession) GetApproved(ddcId *big.Int) (common.Address, error) {
	return _DDC721Logic.Contract.GetApproved(&_DDC721Logic.CallOpts, ddcId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ddcId) view returns(address)
func (_DDC721Logic *DDC721LogicCallerSession) GetApproved(ddcId *big.Int) (common.Address, error) {
	return _DDC721Logic.Contract.GetApproved(&_DDC721Logic.CallOpts, ddcId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC721Logic *DDC721LogicCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC721Logic *DDC721LogicSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DDC721Logic.Contract.IsApprovedForAll(&_DDC721Logic.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC721Logic *DDC721LogicCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DDC721Logic.Contract.IsApprovedForAll(&_DDC721Logic.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DDC721Logic *DDC721LogicCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DDC721Logic *DDC721LogicSession) Name() (string, error) {
	return _DDC721Logic.Contract.Name(&_DDC721Logic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DDC721Logic *DDC721LogicCallerSession) Name() (string, error) {
	return _DDC721Logic.Contract.Name(&_DDC721Logic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC721Logic *DDC721LogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC721Logic *DDC721LogicSession) Owner() (common.Address, error) {
	return _DDC721Logic.Contract.Owner(&_DDC721Logic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC721Logic *DDC721LogicCallerSession) Owner() (common.Address, error) {
	return _DDC721Logic.Contract.Owner(&_DDC721Logic.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ddcId) view returns(address)
func (_DDC721Logic *DDC721LogicCaller) OwnerOf(opts *bind.CallOpts, ddcId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "ownerOf", ddcId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ddcId) view returns(address)
func (_DDC721Logic *DDC721LogicSession) OwnerOf(ddcId *big.Int) (common.Address, error) {
	return _DDC721Logic.Contract.OwnerOf(&_DDC721Logic.CallOpts, ddcId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ddcId) view returns(address)
func (_DDC721Logic *DDC721LogicCallerSession) OwnerOf(ddcId *big.Int) (common.Address, error) {
	return _DDC721Logic.Contract.OwnerOf(&_DDC721Logic.CallOpts, ddcId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC721Logic *DDC721LogicCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC721Logic *DDC721LogicSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DDC721Logic.Contract.SupportsInterface(&_DDC721Logic.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC721Logic *DDC721LogicCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DDC721Logic.Contract.SupportsInterface(&_DDC721Logic.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DDC721Logic *DDC721LogicCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DDC721Logic.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DDC721Logic *DDC721LogicSession) Symbol() (string, error) {
	return _DDC721Logic.Contract.Symbol(&_DDC721Logic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DDC721Logic *DDC721LogicCallerSession) Symbol() (string, error) {
	return _DDC721Logic.Contract.Symbol(&_DDC721Logic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactor) Approve(opts *bind.TransactOpts, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "approve", to, ddcId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicSession) Approve(to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Approve(&_DDC721Logic.TransactOpts, to, ddcId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) Approve(to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Approve(&_DDC721Logic.TransactOpts, to, ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactor) Burn(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "burn", ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicSession) Burn(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Burn(&_DDC721Logic.TransactOpts, ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) Burn(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Burn(&_DDC721Logic.TransactOpts, ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactor) Freeze(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "freeze", ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicSession) Freeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Freeze(&_DDC721Logic.TransactOpts, ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) Freeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Freeze(&_DDC721Logic.TransactOpts, ddcId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC721Logic *DDC721LogicTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC721Logic *DDC721LogicSession) Initialize() (*types.Transaction, error) {
	return _DDC721Logic.Contract.Initialize(&_DDC721Logic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC721Logic *DDC721LogicTransactorSession) Initialize() (*types.Transaction, error) {
	return _DDC721Logic.Contract.Initialize(&_DDC721Logic.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string ddcURI_) returns()
func (_DDC721Logic *DDC721LogicTransactor) Mint(opts *bind.TransactOpts, to common.Address, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "mint", to, ddcURI_)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string ddcURI_) returns()
func (_DDC721Logic *DDC721LogicSession) Mint(to common.Address, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Mint(&_DDC721Logic.TransactOpts, to, ddcURI_)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string ddcURI_) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) Mint(to common.Address, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721Logic.Contract.Mint(&_DDC721Logic.TransactOpts, to, ddcURI_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC721Logic *DDC721LogicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC721Logic *DDC721LogicSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDC721Logic.Contract.RenounceOwnership(&_DDC721Logic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC721Logic *DDC721LogicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDC721Logic.Contract.RenounceOwnership(&_DDC721Logic.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0xf6dda936.
//
// Solidity: function safeMint(address to, string _ddcURI, bytes _data) returns()
func (_DDC721Logic *DDC721LogicTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, _ddcURI string, _data []byte) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "safeMint", to, _ddcURI, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xf6dda936.
//
// Solidity: function safeMint(address to, string _ddcURI, bytes _data) returns()
func (_DDC721Logic *DDC721LogicSession) SafeMint(to common.Address, _ddcURI string, _data []byte) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SafeMint(&_DDC721Logic.TransactOpts, to, _ddcURI, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xf6dda936.
//
// Solidity: function safeMint(address to, string _ddcURI, bytes _data) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SafeMint(to common.Address, _ddcURI string, _data []byte) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SafeMint(&_DDC721Logic.TransactOpts, to, _ddcURI, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, bytes data) returns()
func (_DDC721Logic *DDC721LogicTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "safeTransferFrom", from, to, ddcId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, bytes data) returns()
func (_DDC721Logic *DDC721LogicSession) SafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SafeTransferFrom(&_DDC721Logic.TransactOpts, from, to, ddcId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, bytes data) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SafeTransferFrom(&_DDC721Logic.TransactOpts, from, to, ddcId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC721Logic *DDC721LogicTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC721Logic *DDC721LogicSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetApprovalForAll(&_DDC721Logic.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetApprovalForAll(&_DDC721Logic.TransactOpts, operator, approved)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC721Logic *DDC721LogicTransactor) SetAuthorityProxyAddress(opts *bind.TransactOpts, authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "setAuthorityProxyAddress", authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC721Logic *DDC721LogicSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetAuthorityProxyAddress(&_DDC721Logic.TransactOpts, authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetAuthorityProxyAddress(&_DDC721Logic.TransactOpts, authorityProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC721Logic *DDC721LogicTransactor) SetChargeProxyAddress(opts *bind.TransactOpts, chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "setChargeProxyAddress", chargeProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC721Logic *DDC721LogicSession) SetChargeProxyAddress(chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetChargeProxyAddress(&_DDC721Logic.TransactOpts, chargeProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SetChargeProxyAddress(chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetChargeProxyAddress(&_DDC721Logic.TransactOpts, chargeProxyAddress)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string name_, string symbol_) returns()
func (_DDC721Logic *DDC721LogicTransactor) SetNameAndSymbol(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "setNameAndSymbol", name_, symbol_)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string name_, string symbol_) returns()
func (_DDC721Logic *DDC721LogicSession) SetNameAndSymbol(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetNameAndSymbol(&_DDC721Logic.TransactOpts, name_, symbol_)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string name_, string symbol_) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SetNameAndSymbol(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetNameAndSymbol(&_DDC721Logic.TransactOpts, name_, symbol_)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 ddcId, string ddcURI_) returns()
func (_DDC721Logic *DDC721LogicTransactor) SetURI(opts *bind.TransactOpts, ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "setURI", ddcId, ddcURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 ddcId, string ddcURI_) returns()
func (_DDC721Logic *DDC721LogicSession) SetURI(ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetURI(&_DDC721Logic.TransactOpts, ddcId, ddcURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 ddcId, string ddcURI_) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) SetURI(ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721Logic.Contract.SetURI(&_DDC721Logic.TransactOpts, ddcId, ddcURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "transferFrom", from, to, ddcId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicSession) TransferFrom(from common.Address, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.TransferFrom(&_DDC721Logic.TransactOpts, from, to, ddcId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) TransferFrom(from common.Address, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.TransferFrom(&_DDC721Logic.TransactOpts, from, to, ddcId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC721Logic *DDC721LogicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC721Logic *DDC721LogicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.TransferOwnership(&_DDC721Logic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.TransferOwnership(&_DDC721Logic.TransactOpts, newOwner)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactor) UnFreeze(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "unFreeze", ddcId)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicSession) UnFreeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.UnFreeze(&_DDC721Logic.TransactOpts, ddcId)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) UnFreeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721Logic.Contract.UnFreeze(&_DDC721Logic.TransactOpts, ddcId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC721Logic *DDC721LogicTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC721Logic *DDC721LogicSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.UpgradeTo(&_DDC721Logic.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC721Logic *DDC721LogicTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DDC721Logic.Contract.UpgradeTo(&_DDC721Logic.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC721Logic *DDC721LogicTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC721Logic.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC721Logic *DDC721LogicSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC721Logic.Contract.UpgradeToAndCall(&_DDC721Logic.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC721Logic *DDC721LogicTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC721Logic.Contract.UpgradeToAndCall(&_DDC721Logic.TransactOpts, newImplementation, data)
}

// DDC721LogicAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the DDC721Logic contract.
type DDC721LogicAdminChangedIterator struct {
	Event *DDC721LogicAdminChanged // Event containing the contract specifics and raw log

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
func (it *DDC721LogicAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicAdminChanged)
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
		it.Event = new(DDC721LogicAdminChanged)
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
func (it *DDC721LogicAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicAdminChanged represents a AdminChanged event raised by the DDC721Logic contract.
type DDC721LogicAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC721Logic *DDC721LogicFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DDC721LogicAdminChangedIterator, error) {

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DDC721LogicAdminChangedIterator{contract: _DDC721Logic.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC721Logic *DDC721LogicFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DDC721LogicAdminChanged) (event.Subscription, error) {

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicAdminChanged)
				if err := _DDC721Logic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_DDC721Logic *DDC721LogicFilterer) ParseAdminChanged(log types.Log) (*DDC721LogicAdminChanged, error) {
	event := new(DDC721LogicAdminChanged)
	if err := _DDC721Logic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DDC721Logic contract.
type DDC721LogicApprovalIterator struct {
	Event *DDC721LogicApproval // Event containing the contract specifics and raw log

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
func (it *DDC721LogicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicApproval)
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
		it.Event = new(DDC721LogicApproval)
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
func (it *DDC721LogicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicApproval represents a Approval event raised by the DDC721Logic contract.
type DDC721LogicApproval struct {
	Owner    common.Address
	Approved common.Address
	DdcId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed ddcId)
func (_DDC721Logic *DDC721LogicFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, ddcId []*big.Int) (*DDC721LogicApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicApprovalIterator{contract: _DDC721Logic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed ddcId)
func (_DDC721Logic *DDC721LogicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DDC721LogicApproval, owner []common.Address, approved []common.Address, ddcId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicApproval)
				if err := _DDC721Logic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed ddcId)
func (_DDC721Logic *DDC721LogicFilterer) ParseApproval(log types.Log) (*DDC721LogicApproval, error) {
	event := new(DDC721LogicApproval)
	if err := _DDC721Logic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DDC721Logic contract.
type DDC721LogicApprovalForAllIterator struct {
	Event *DDC721LogicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DDC721LogicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicApprovalForAll)
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
		it.Event = new(DDC721LogicApprovalForAll)
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
func (it *DDC721LogicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicApprovalForAll represents a ApprovalForAll event raised by the DDC721Logic contract.
type DDC721LogicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC721Logic *DDC721LogicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DDC721LogicApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicApprovalForAllIterator{contract: _DDC721Logic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC721Logic *DDC721LogicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DDC721LogicApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicApprovalForAll)
				if err := _DDC721Logic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC721Logic *DDC721LogicFilterer) ParseApprovalForAll(log types.Log) (*DDC721LogicApprovalForAll, error) {
	event := new(DDC721LogicApprovalForAll)
	if err := _DDC721Logic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the DDC721Logic contract.
type DDC721LogicBeaconUpgradedIterator struct {
	Event *DDC721LogicBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *DDC721LogicBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicBeaconUpgraded)
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
		it.Event = new(DDC721LogicBeaconUpgraded)
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
func (it *DDC721LogicBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicBeaconUpgraded represents a BeaconUpgraded event raised by the DDC721Logic contract.
type DDC721LogicBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC721Logic *DDC721LogicFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*DDC721LogicBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicBeaconUpgradedIterator{contract: _DDC721Logic.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC721Logic *DDC721LogicFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *DDC721LogicBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicBeaconUpgraded)
				if err := _DDC721Logic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_DDC721Logic *DDC721LogicFilterer) ParseBeaconUpgraded(log types.Log) (*DDC721LogicBeaconUpgraded, error) {
	event := new(DDC721LogicBeaconUpgraded)
	if err := _DDC721Logic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicEnterBlacklistIterator is returned from FilterEnterBlacklist and is used to iterate over the raw logs and unpacked data for EnterBlacklist events raised by the DDC721Logic contract.
type DDC721LogicEnterBlacklistIterator struct {
	Event *DDC721LogicEnterBlacklist // Event containing the contract specifics and raw log

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
func (it *DDC721LogicEnterBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicEnterBlacklist)
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
		it.Event = new(DDC721LogicEnterBlacklist)
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
func (it *DDC721LogicEnterBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicEnterBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicEnterBlacklist represents a EnterBlacklist event raised by the DDC721Logic contract.
type DDC721LogicEnterBlacklist struct {
	Sender common.Address
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterBlacklist is a free log retrieval operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721Logic *DDC721LogicFilterer) FilterEnterBlacklist(opts *bind.FilterOpts, sender []common.Address) (*DDC721LogicEnterBlacklistIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "EnterBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicEnterBlacklistIterator{contract: _DDC721Logic.contract, event: "EnterBlacklist", logs: logs, sub: sub}, nil
}

// WatchEnterBlacklist is a free log subscription operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721Logic *DDC721LogicFilterer) WatchEnterBlacklist(opts *bind.WatchOpts, sink chan<- *DDC721LogicEnterBlacklist, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "EnterBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicEnterBlacklist)
				if err := _DDC721Logic.contract.UnpackLog(event, "EnterBlacklist", log); err != nil {
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

// ParseEnterBlacklist is a log parse operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721Logic *DDC721LogicFilterer) ParseEnterBlacklist(log types.Log) (*DDC721LogicEnterBlacklist, error) {
	event := new(DDC721LogicEnterBlacklist)
	if err := _DDC721Logic.contract.UnpackLog(event, "EnterBlacklist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicExitBlacklistIterator is returned from FilterExitBlacklist and is used to iterate over the raw logs and unpacked data for ExitBlacklist events raised by the DDC721Logic contract.
type DDC721LogicExitBlacklistIterator struct {
	Event *DDC721LogicExitBlacklist // Event containing the contract specifics and raw log

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
func (it *DDC721LogicExitBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicExitBlacklist)
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
		it.Event = new(DDC721LogicExitBlacklist)
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
func (it *DDC721LogicExitBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicExitBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicExitBlacklist represents a ExitBlacklist event raised by the DDC721Logic contract.
type DDC721LogicExitBlacklist struct {
	Sender common.Address
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitBlacklist is a free log retrieval operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721Logic *DDC721LogicFilterer) FilterExitBlacklist(opts *bind.FilterOpts, sender []common.Address) (*DDC721LogicExitBlacklistIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "ExitBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicExitBlacklistIterator{contract: _DDC721Logic.contract, event: "ExitBlacklist", logs: logs, sub: sub}, nil
}

// WatchExitBlacklist is a free log subscription operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721Logic *DDC721LogicFilterer) WatchExitBlacklist(opts *bind.WatchOpts, sink chan<- *DDC721LogicExitBlacklist, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "ExitBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicExitBlacklist)
				if err := _DDC721Logic.contract.UnpackLog(event, "ExitBlacklist", log); err != nil {
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

// ParseExitBlacklist is a log parse operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721Logic *DDC721LogicFilterer) ParseExitBlacklist(log types.Log) (*DDC721LogicExitBlacklist, error) {
	event := new(DDC721LogicExitBlacklist)
	if err := _DDC721Logic.contract.UnpackLog(event, "ExitBlacklist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DDC721Logic contract.
type DDC721LogicOwnershipTransferredIterator struct {
	Event *DDC721LogicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DDC721LogicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicOwnershipTransferred)
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
		it.Event = new(DDC721LogicOwnershipTransferred)
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
func (it *DDC721LogicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicOwnershipTransferred represents a OwnershipTransferred event raised by the DDC721Logic contract.
type DDC721LogicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDC721Logic *DDC721LogicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DDC721LogicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicOwnershipTransferredIterator{contract: _DDC721Logic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDC721Logic *DDC721LogicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DDC721LogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicOwnershipTransferred)
				if err := _DDC721Logic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DDC721Logic *DDC721LogicFilterer) ParseOwnershipTransferred(log types.Log) (*DDC721LogicOwnershipTransferred, error) {
	event := new(DDC721LogicOwnershipTransferred)
	if err := _DDC721Logic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicSetNameAndSymbolIterator is returned from FilterSetNameAndSymbol and is used to iterate over the raw logs and unpacked data for SetNameAndSymbol events raised by the DDC721Logic contract.
type DDC721LogicSetNameAndSymbolIterator struct {
	Event *DDC721LogicSetNameAndSymbol // Event containing the contract specifics and raw log

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
func (it *DDC721LogicSetNameAndSymbolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicSetNameAndSymbol)
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
		it.Event = new(DDC721LogicSetNameAndSymbol)
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
func (it *DDC721LogicSetNameAndSymbolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicSetNameAndSymbolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicSetNameAndSymbol represents a SetNameAndSymbol event raised by the DDC721Logic contract.
type DDC721LogicSetNameAndSymbol struct {
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetNameAndSymbol is a free log retrieval operation binding the contract event 0xd150542982eaa18f20ceb10f524c418899e6a331a81ee4a70daa921f976fab3b.
//
// Solidity: event SetNameAndSymbol(string name, string symbol)
func (_DDC721Logic *DDC721LogicFilterer) FilterSetNameAndSymbol(opts *bind.FilterOpts) (*DDC721LogicSetNameAndSymbolIterator, error) {

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "SetNameAndSymbol")
	if err != nil {
		return nil, err
	}
	return &DDC721LogicSetNameAndSymbolIterator{contract: _DDC721Logic.contract, event: "SetNameAndSymbol", logs: logs, sub: sub}, nil
}

// WatchSetNameAndSymbol is a free log subscription operation binding the contract event 0xd150542982eaa18f20ceb10f524c418899e6a331a81ee4a70daa921f976fab3b.
//
// Solidity: event SetNameAndSymbol(string name, string symbol)
func (_DDC721Logic *DDC721LogicFilterer) WatchSetNameAndSymbol(opts *bind.WatchOpts, sink chan<- *DDC721LogicSetNameAndSymbol) (event.Subscription, error) {

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "SetNameAndSymbol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicSetNameAndSymbol)
				if err := _DDC721Logic.contract.UnpackLog(event, "SetNameAndSymbol", log); err != nil {
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

// ParseSetNameAndSymbol is a log parse operation binding the contract event 0xd150542982eaa18f20ceb10f524c418899e6a331a81ee4a70daa921f976fab3b.
//
// Solidity: event SetNameAndSymbol(string name, string symbol)
func (_DDC721Logic *DDC721LogicFilterer) ParseSetNameAndSymbol(log types.Log) (*DDC721LogicSetNameAndSymbol, error) {
	event := new(DDC721LogicSetNameAndSymbol)
	if err := _DDC721Logic.contract.UnpackLog(event, "SetNameAndSymbol", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicSetURIIterator is returned from FilterSetURI and is used to iterate over the raw logs and unpacked data for SetURI events raised by the DDC721Logic contract.
type DDC721LogicSetURIIterator struct {
	Event *DDC721LogicSetURI // Event containing the contract specifics and raw log

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
func (it *DDC721LogicSetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicSetURI)
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
		it.Event = new(DDC721LogicSetURI)
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
func (it *DDC721LogicSetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicSetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicSetURI represents a SetURI event raised by the DDC721Logic contract.
type DDC721LogicSetURI struct {
	DdcId  *big.Int
	DdcURI string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetURI is a free log retrieval operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC721Logic *DDC721LogicFilterer) FilterSetURI(opts *bind.FilterOpts, ddcId []*big.Int) (*DDC721LogicSetURIIterator, error) {

	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "SetURI", ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicSetURIIterator{contract: _DDC721Logic.contract, event: "SetURI", logs: logs, sub: sub}, nil
}

// WatchSetURI is a free log subscription operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC721Logic *DDC721LogicFilterer) WatchSetURI(opts *bind.WatchOpts, sink chan<- *DDC721LogicSetURI, ddcId []*big.Int) (event.Subscription, error) {

	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "SetURI", ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicSetURI)
				if err := _DDC721Logic.contract.UnpackLog(event, "SetURI", log); err != nil {
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

// ParseSetURI is a log parse operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC721Logic *DDC721LogicFilterer) ParseSetURI(log types.Log) (*DDC721LogicSetURI, error) {
	event := new(DDC721LogicSetURI)
	if err := _DDC721Logic.contract.UnpackLog(event, "SetURI", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DDC721Logic contract.
type DDC721LogicTransferIterator struct {
	Event *DDC721LogicTransfer // Event containing the contract specifics and raw log

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
func (it *DDC721LogicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicTransfer)
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
		it.Event = new(DDC721LogicTransfer)
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
func (it *DDC721LogicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicTransfer represents a Transfer event raised by the DDC721Logic contract.
type DDC721LogicTransfer struct {
	From  common.Address
	To    common.Address
	DdcId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed ddcId)
func (_DDC721Logic *DDC721LogicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, ddcId []*big.Int) (*DDC721LogicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "Transfer", fromRule, toRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicTransferIterator{contract: _DDC721Logic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed ddcId)
func (_DDC721Logic *DDC721LogicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DDC721LogicTransfer, from []common.Address, to []common.Address, ddcId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "Transfer", fromRule, toRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicTransfer)
				if err := _DDC721Logic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed ddcId)
func (_DDC721Logic *DDC721LogicFilterer) ParseTransfer(log types.Log) (*DDC721LogicTransfer, error) {
	event := new(DDC721LogicTransfer)
	if err := _DDC721Logic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC721LogicUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the DDC721Logic contract.
type DDC721LogicUpgradedIterator struct {
	Event *DDC721LogicUpgraded // Event containing the contract specifics and raw log

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
func (it *DDC721LogicUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721LogicUpgraded)
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
		it.Event = new(DDC721LogicUpgraded)
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
func (it *DDC721LogicUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LogicUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721LogicUpgraded represents a Upgraded event raised by the DDC721Logic contract.
type DDC721LogicUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC721Logic *DDC721LogicFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DDC721LogicUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DDC721Logic.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LogicUpgradedIterator{contract: _DDC721Logic.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC721Logic *DDC721LogicFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DDC721LogicUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DDC721Logic.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721LogicUpgraded)
				if err := _DDC721Logic.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_DDC721Logic *DDC721LogicFilterer) ParseUpgraded(log types.Log) (*DDC721LogicUpgraded, error) {
	event := new(DDC721LogicUpgraded)
	if err := _DDC721Logic.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}
