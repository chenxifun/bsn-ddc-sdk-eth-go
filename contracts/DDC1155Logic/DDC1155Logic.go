// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DDC1155Logic

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

// DDC1155LogicABI is the input ABI used to generate the binding from.
const DDC1155LogicABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"EnterBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ExitBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ddcURI\",\"type\":\"string\"}],\"name\":\"SetURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ddcIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ddcIds\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ddcIds\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ddcURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ddcIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ddcURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"ddcURIs\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeMintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorityProxyAddress\",\"type\":\"address\"}],\"name\":\"setAuthorityProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"chargeProxyAddress\",\"type\":\"address\"}],\"name\":\"setChargeProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DDC1155Logic is an auto generated Go binding around an Ethereum contract.
type DDC1155Logic struct {
	DDC1155LogicCaller     // Read-only binding to the contract
	DDC1155LogicTransactor // Write-only binding to the contract
	DDC1155LogicFilterer   // Log filterer for contract events
}

// DDC1155LogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type DDC1155LogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC1155LogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DDC1155LogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC1155LogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DDC1155LogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC1155LogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DDC1155LogicSession struct {
	Contract     *DDC1155Logic     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDC1155LogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DDC1155LogicCallerSession struct {
	Contract *DDC1155LogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DDC1155LogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DDC1155LogicTransactorSession struct {
	Contract     *DDC1155LogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DDC1155LogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type DDC1155LogicRaw struct {
	Contract *DDC1155Logic // Generic contract binding to access the raw methods on
}

// DDC1155LogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DDC1155LogicCallerRaw struct {
	Contract *DDC1155LogicCaller // Generic read-only contract binding to access the raw methods on
}

// DDC1155LogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DDC1155LogicTransactorRaw struct {
	Contract *DDC1155LogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDDC1155Logic creates a new instance of DDC1155Logic, bound to a specific deployed contract.
func NewDDC1155Logic(address common.Address, backend bind.ContractBackend) (*DDC1155Logic, error) {
	contract, err := bindDDC1155Logic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DDC1155Logic{DDC1155LogicCaller: DDC1155LogicCaller{contract: contract}, DDC1155LogicTransactor: DDC1155LogicTransactor{contract: contract}, DDC1155LogicFilterer: DDC1155LogicFilterer{contract: contract}}, nil
}

// NewDDC1155LogicCaller creates a new read-only instance of DDC1155Logic, bound to a specific deployed contract.
func NewDDC1155LogicCaller(address common.Address, caller bind.ContractCaller) (*DDC1155LogicCaller, error) {
	contract, err := bindDDC1155Logic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicCaller{contract: contract}, nil
}

// NewDDC1155LogicTransactor creates a new write-only instance of DDC1155Logic, bound to a specific deployed contract.
func NewDDC1155LogicTransactor(address common.Address, transactor bind.ContractTransactor) (*DDC1155LogicTransactor, error) {
	contract, err := bindDDC1155Logic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicTransactor{contract: contract}, nil
}

// NewDDC1155LogicFilterer creates a new log filterer instance of DDC1155Logic, bound to a specific deployed contract.
func NewDDC1155LogicFilterer(address common.Address, filterer bind.ContractFilterer) (*DDC1155LogicFilterer, error) {
	contract, err := bindDDC1155Logic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicFilterer{contract: contract}, nil
}

// bindDDC1155Logic binds a generic wrapper to an already deployed contract.
func bindDDC1155Logic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DDC1155LogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDC1155Logic *DDC1155LogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DDC1155Logic.Contract.DDC1155LogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDC1155Logic *DDC1155LogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.DDC1155LogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDC1155Logic *DDC1155LogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.DDC1155LogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDC1155Logic *DDC1155LogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DDC1155Logic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDC1155Logic *DDC1155LogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDC1155Logic *DDC1155LogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 ddcId) view returns(uint256 balance)
func (_DDC1155Logic *DDC1155LogicCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, ddcId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DDC1155Logic.contract.Call(opts, out, "balanceOf", owner, ddcId)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 ddcId) view returns(uint256 balance)
func (_DDC1155Logic *DDC1155LogicSession) BalanceOf(owner common.Address, ddcId *big.Int) (*big.Int, error) {
	return _DDC1155Logic.Contract.BalanceOf(&_DDC1155Logic.CallOpts, owner, ddcId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 ddcId) view returns(uint256 balance)
func (_DDC1155Logic *DDC1155LogicCallerSession) BalanceOf(owner common.Address, ddcId *big.Int) (*big.Int, error) {
	return _DDC1155Logic.Contract.BalanceOf(&_DDC1155Logic.CallOpts, owner, ddcId)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ddcIds) view returns(uint256[])
func (_DDC1155Logic *DDC1155LogicCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ddcIds []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DDC1155Logic.contract.Call(opts, out, "balanceOfBatch", owners, ddcIds)
	return *ret0, err
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ddcIds) view returns(uint256[])
func (_DDC1155Logic *DDC1155LogicSession) BalanceOfBatch(owners []common.Address, ddcIds []*big.Int) ([]*big.Int, error) {
	return _DDC1155Logic.Contract.BalanceOfBatch(&_DDC1155Logic.CallOpts, owners, ddcIds)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ddcIds) view returns(uint256[])
func (_DDC1155Logic *DDC1155LogicCallerSession) BalanceOfBatch(owners []common.Address, ddcIds []*big.Int) ([]*big.Int, error) {
	return _DDC1155Logic.Contract.BalanceOfBatch(&_DDC1155Logic.CallOpts, owners, ddcIds)
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC1155Logic *DDC1155LogicCaller) DdcURI(opts *bind.CallOpts, ddcId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DDC1155Logic.contract.Call(opts, out, "ddcURI", ddcId)
	return *ret0, err
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC1155Logic *DDC1155LogicSession) DdcURI(ddcId *big.Int) (string, error) {
	return _DDC1155Logic.Contract.DdcURI(&_DDC1155Logic.CallOpts, ddcId)
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC1155Logic *DDC1155LogicCallerSession) DdcURI(ddcId *big.Int) (string, error) {
	return _DDC1155Logic.Contract.DdcURI(&_DDC1155Logic.CallOpts, ddcId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC1155Logic *DDC1155LogicCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DDC1155Logic.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC1155Logic *DDC1155LogicSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DDC1155Logic.Contract.IsApprovedForAll(&_DDC1155Logic.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC1155Logic *DDC1155LogicCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DDC1155Logic.Contract.IsApprovedForAll(&_DDC1155Logic.CallOpts, owner, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC1155Logic *DDC1155LogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DDC1155Logic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC1155Logic *DDC1155LogicSession) Owner() (common.Address, error) {
	return _DDC1155Logic.Contract.Owner(&_DDC1155Logic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC1155Logic *DDC1155LogicCallerSession) Owner() (common.Address, error) {
	return _DDC1155Logic.Contract.Owner(&_DDC1155Logic.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC1155Logic *DDC1155LogicCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DDC1155Logic.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC1155Logic *DDC1155LogicSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DDC1155Logic.Contract.SupportsInterface(&_DDC1155Logic.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC1155Logic *DDC1155LogicCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DDC1155Logic.Contract.SupportsInterface(&_DDC1155Logic.CallOpts, interfaceId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address owner, uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) Burn(opts *bind.TransactOpts, owner common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "burn", owner, ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address owner, uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicSession) Burn(owner common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.Burn(&_DDC1155Logic.TransactOpts, owner, ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address owner, uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) Burn(owner common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.Burn(&_DDC1155Logic.TransactOpts, owner, ddcId)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xb2dc5dc3.
//
// Solidity: function burnBatch(address owner, uint256[] ddcIds) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) BurnBatch(opts *bind.TransactOpts, owner common.Address, ddcIds []*big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "burnBatch", owner, ddcIds)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xb2dc5dc3.
//
// Solidity: function burnBatch(address owner, uint256[] ddcIds) returns()
func (_DDC1155Logic *DDC1155LogicSession) BurnBatch(owner common.Address, ddcIds []*big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.BurnBatch(&_DDC1155Logic.TransactOpts, owner, ddcIds)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xb2dc5dc3.
//
// Solidity: function burnBatch(address owner, uint256[] ddcIds) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) BurnBatch(owner common.Address, ddcIds []*big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.BurnBatch(&_DDC1155Logic.TransactOpts, owner, ddcIds)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) Freeze(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "freeze", ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicSession) Freeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.Freeze(&_DDC1155Logic.TransactOpts, ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) Freeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.Freeze(&_DDC1155Logic.TransactOpts, ddcId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC1155Logic *DDC1155LogicTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC1155Logic *DDC1155LogicSession) Initialize() (*types.Transaction, error) {
	return _DDC1155Logic.Contract.Initialize(&_DDC1155Logic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) Initialize() (*types.Transaction, error) {
	return _DDC1155Logic.Contract.Initialize(&_DDC1155Logic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC1155Logic *DDC1155LogicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC1155Logic *DDC1155LogicSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDC1155Logic.Contract.RenounceOwnership(&_DDC1155Logic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDC1155Logic.Contract.RenounceOwnership(&_DDC1155Logic.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ddcIds, uint256[] amounts, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcIds []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "safeBatchTransferFrom", from, to, ddcIds, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ddcIds, uint256[] amounts, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicSession) SafeBatchTransferFrom(from common.Address, to common.Address, ddcIds []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeBatchTransferFrom(&_DDC1155Logic.TransactOpts, from, to, ddcIds, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ddcIds, uint256[] amounts, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ddcIds []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeBatchTransferFrom(&_DDC1155Logic.TransactOpts, from, to, ddcIds, amounts, data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xb55bc617.
//
// Solidity: function safeMint(address to, uint256 amount, string _ddcURI, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, amount *big.Int, _ddcURI string, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "safeMint", to, amount, _ddcURI, data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xb55bc617.
//
// Solidity: function safeMint(address to, uint256 amount, string _ddcURI, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicSession) SafeMint(to common.Address, amount *big.Int, _ddcURI string, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeMint(&_DDC1155Logic.TransactOpts, to, amount, _ddcURI, data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xb55bc617.
//
// Solidity: function safeMint(address to, uint256 amount, string _ddcURI, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SafeMint(to common.Address, amount *big.Int, _ddcURI string, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeMint(&_DDC1155Logic.TransactOpts, to, amount, _ddcURI, data)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0x63570355.
//
// Solidity: function safeMintBatch(address to, uint256[] amounts, string[] ddcURIs, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SafeMintBatch(opts *bind.TransactOpts, to common.Address, amounts []*big.Int, ddcURIs []string, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "safeMintBatch", to, amounts, ddcURIs, data)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0x63570355.
//
// Solidity: function safeMintBatch(address to, uint256[] amounts, string[] ddcURIs, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicSession) SafeMintBatch(to common.Address, amounts []*big.Int, ddcURIs []string, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeMintBatch(&_DDC1155Logic.TransactOpts, to, amounts, ddcURIs, data)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0x63570355.
//
// Solidity: function safeMintBatch(address to, uint256[] amounts, string[] ddcURIs, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SafeMintBatch(to common.Address, amounts []*big.Int, ddcURIs []string, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeMintBatch(&_DDC1155Logic.TransactOpts, to, amounts, ddcURIs, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, uint256 amount, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "safeTransferFrom", from, to, ddcId, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, uint256 amount, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicSession) SafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeTransferFrom(&_DDC1155Logic.TransactOpts, from, to, ddcId, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, uint256 amount, bytes data) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SafeTransferFrom(&_DDC1155Logic.TransactOpts, from, to, ddcId, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC1155Logic *DDC1155LogicSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetApprovalForAll(&_DDC1155Logic.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetApprovalForAll(&_DDC1155Logic.TransactOpts, operator, approved)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SetAuthorityProxyAddress(opts *bind.TransactOpts, authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "setAuthorityProxyAddress", authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC1155Logic *DDC1155LogicSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetAuthorityProxyAddress(&_DDC1155Logic.TransactOpts, authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetAuthorityProxyAddress(&_DDC1155Logic.TransactOpts, authorityProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SetChargeProxyAddress(opts *bind.TransactOpts, chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "setChargeProxyAddress", chargeProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC1155Logic *DDC1155LogicSession) SetChargeProxyAddress(chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetChargeProxyAddress(&_DDC1155Logic.TransactOpts, chargeProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SetChargeProxyAddress(chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetChargeProxyAddress(&_DDC1155Logic.TransactOpts, chargeProxyAddress)
}

// SetURI is a paid mutator transaction binding the contract method 0x685e8247.
//
// Solidity: function setURI(address owner, uint256 ddcId, string ddcURI_) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) SetURI(opts *bind.TransactOpts, owner common.Address, ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "setURI", owner, ddcId, ddcURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x685e8247.
//
// Solidity: function setURI(address owner, uint256 ddcId, string ddcURI_) returns()
func (_DDC1155Logic *DDC1155LogicSession) SetURI(owner common.Address, ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetURI(&_DDC1155Logic.TransactOpts, owner, ddcId, ddcURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x685e8247.
//
// Solidity: function setURI(address owner, uint256 ddcId, string ddcURI_) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) SetURI(owner common.Address, ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.SetURI(&_DDC1155Logic.TransactOpts, owner, ddcId, ddcURI_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC1155Logic *DDC1155LogicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.TransferOwnership(&_DDC1155Logic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.TransferOwnership(&_DDC1155Logic.TransactOpts, newOwner)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) UnFreeze(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "unFreeze", ddcId)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicSession) UnFreeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.UnFreeze(&_DDC1155Logic.TransactOpts, ddcId)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) UnFreeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.UnFreeze(&_DDC1155Logic.TransactOpts, ddcId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC1155Logic *DDC1155LogicTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC1155Logic *DDC1155LogicSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.UpgradeTo(&_DDC1155Logic.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.UpgradeTo(&_DDC1155Logic.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC1155Logic *DDC1155LogicTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC1155Logic *DDC1155LogicSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.UpgradeToAndCall(&_DDC1155Logic.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC1155Logic *DDC1155LogicTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC1155Logic.Contract.UpgradeToAndCall(&_DDC1155Logic.TransactOpts, newImplementation, data)
}

// DDC1155LogicAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the DDC1155Logic contract.
type DDC1155LogicAdminChangedIterator struct {
	Event *DDC1155LogicAdminChanged // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicAdminChanged)
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
		it.Event = new(DDC1155LogicAdminChanged)
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
func (it *DDC1155LogicAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicAdminChanged represents a AdminChanged event raised by the DDC1155Logic contract.
type DDC1155LogicAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DDC1155LogicAdminChangedIterator, error) {

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicAdminChangedIterator{contract: _DDC1155Logic.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DDC1155LogicAdminChanged) (event.Subscription, error) {

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicAdminChanged)
				if err := _DDC1155Logic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseAdminChanged(log types.Log) (*DDC1155LogicAdminChanged, error) {
	event := new(DDC1155LogicAdminChanged)
	if err := _DDC1155Logic.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DDC1155Logic contract.
type DDC1155LogicApprovalForAllIterator struct {
	Event *DDC1155LogicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicApprovalForAll)
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
		it.Event = new(DDC1155LogicApprovalForAll)
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
func (it *DDC1155LogicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicApprovalForAll represents a ApprovalForAll event raised by the DDC1155Logic contract.
type DDC1155LogicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DDC1155LogicApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicApprovalForAllIterator{contract: _DDC1155Logic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DDC1155LogicApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicApprovalForAll)
				if err := _DDC1155Logic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseApprovalForAll(log types.Log) (*DDC1155LogicApprovalForAll, error) {
	event := new(DDC1155LogicApprovalForAll)
	if err := _DDC1155Logic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the DDC1155Logic contract.
type DDC1155LogicBeaconUpgradedIterator struct {
	Event *DDC1155LogicBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicBeaconUpgraded)
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
		it.Event = new(DDC1155LogicBeaconUpgraded)
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
func (it *DDC1155LogicBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicBeaconUpgraded represents a BeaconUpgraded event raised by the DDC1155Logic contract.
type DDC1155LogicBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*DDC1155LogicBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicBeaconUpgradedIterator{contract: _DDC1155Logic.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *DDC1155LogicBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicBeaconUpgraded)
				if err := _DDC1155Logic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseBeaconUpgraded(log types.Log) (*DDC1155LogicBeaconUpgraded, error) {
	event := new(DDC1155LogicBeaconUpgraded)
	if err := _DDC1155Logic.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicEnterBlacklistIterator is returned from FilterEnterBlacklist and is used to iterate over the raw logs and unpacked data for EnterBlacklist events raised by the DDC1155Logic contract.
type DDC1155LogicEnterBlacklistIterator struct {
	Event *DDC1155LogicEnterBlacklist // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicEnterBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicEnterBlacklist)
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
		it.Event = new(DDC1155LogicEnterBlacklist)
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
func (it *DDC1155LogicEnterBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicEnterBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicEnterBlacklist represents a EnterBlacklist event raised by the DDC1155Logic contract.
type DDC1155LogicEnterBlacklist struct {
	Sender common.Address
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterBlacklist is a free log retrieval operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterEnterBlacklist(opts *bind.FilterOpts, sender []common.Address) (*DDC1155LogicEnterBlacklistIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "EnterBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicEnterBlacklistIterator{contract: _DDC1155Logic.contract, event: "EnterBlacklist", logs: logs, sub: sub}, nil
}

// WatchEnterBlacklist is a free log subscription operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchEnterBlacklist(opts *bind.WatchOpts, sink chan<- *DDC1155LogicEnterBlacklist, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "EnterBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicEnterBlacklist)
				if err := _DDC1155Logic.contract.UnpackLog(event, "EnterBlacklist", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseEnterBlacklist(log types.Log) (*DDC1155LogicEnterBlacklist, error) {
	event := new(DDC1155LogicEnterBlacklist)
	if err := _DDC1155Logic.contract.UnpackLog(event, "EnterBlacklist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicExitBlacklistIterator is returned from FilterExitBlacklist and is used to iterate over the raw logs and unpacked data for ExitBlacklist events raised by the DDC1155Logic contract.
type DDC1155LogicExitBlacklistIterator struct {
	Event *DDC1155LogicExitBlacklist // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicExitBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicExitBlacklist)
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
		it.Event = new(DDC1155LogicExitBlacklist)
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
func (it *DDC1155LogicExitBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicExitBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicExitBlacklist represents a ExitBlacklist event raised by the DDC1155Logic contract.
type DDC1155LogicExitBlacklist struct {
	Sender common.Address
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitBlacklist is a free log retrieval operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterExitBlacklist(opts *bind.FilterOpts, sender []common.Address) (*DDC1155LogicExitBlacklistIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "ExitBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicExitBlacklistIterator{contract: _DDC1155Logic.contract, event: "ExitBlacklist", logs: logs, sub: sub}, nil
}

// WatchExitBlacklist is a free log subscription operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchExitBlacklist(opts *bind.WatchOpts, sink chan<- *DDC1155LogicExitBlacklist, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "ExitBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicExitBlacklist)
				if err := _DDC1155Logic.contract.UnpackLog(event, "ExitBlacklist", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseExitBlacklist(log types.Log) (*DDC1155LogicExitBlacklist, error) {
	event := new(DDC1155LogicExitBlacklist)
	if err := _DDC1155Logic.contract.UnpackLog(event, "ExitBlacklist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DDC1155Logic contract.
type DDC1155LogicOwnershipTransferredIterator struct {
	Event *DDC1155LogicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicOwnershipTransferred)
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
		it.Event = new(DDC1155LogicOwnershipTransferred)
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
func (it *DDC1155LogicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicOwnershipTransferred represents a OwnershipTransferred event raised by the DDC1155Logic contract.
type DDC1155LogicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DDC1155LogicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicOwnershipTransferredIterator{contract: _DDC1155Logic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DDC1155LogicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicOwnershipTransferred)
				if err := _DDC1155Logic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseOwnershipTransferred(log types.Log) (*DDC1155LogicOwnershipTransferred, error) {
	event := new(DDC1155LogicOwnershipTransferred)
	if err := _DDC1155Logic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicSetURIIterator is returned from FilterSetURI and is used to iterate over the raw logs and unpacked data for SetURI events raised by the DDC1155Logic contract.
type DDC1155LogicSetURIIterator struct {
	Event *DDC1155LogicSetURI // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicSetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicSetURI)
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
		it.Event = new(DDC1155LogicSetURI)
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
func (it *DDC1155LogicSetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicSetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicSetURI represents a SetURI event raised by the DDC1155Logic contract.
type DDC1155LogicSetURI struct {
	DdcId  *big.Int
	DdcURI string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetURI is a free log retrieval operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterSetURI(opts *bind.FilterOpts, ddcId []*big.Int) (*DDC1155LogicSetURIIterator, error) {

	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "SetURI", ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicSetURIIterator{contract: _DDC1155Logic.contract, event: "SetURI", logs: logs, sub: sub}, nil
}

// WatchSetURI is a free log subscription operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchSetURI(opts *bind.WatchOpts, sink chan<- *DDC1155LogicSetURI, ddcId []*big.Int) (event.Subscription, error) {

	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "SetURI", ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicSetURI)
				if err := _DDC1155Logic.contract.UnpackLog(event, "SetURI", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseSetURI(log types.Log) (*DDC1155LogicSetURI, error) {
	event := new(DDC1155LogicSetURI)
	if err := _DDC1155Logic.contract.UnpackLog(event, "SetURI", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the DDC1155Logic contract.
type DDC1155LogicTransferBatchIterator struct {
	Event *DDC1155LogicTransferBatch // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicTransferBatch)
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
		it.Event = new(DDC1155LogicTransferBatch)
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
func (it *DDC1155LogicTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicTransferBatch represents a TransferBatch event raised by the DDC1155Logic contract.
type DDC1155LogicTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	DdcIds   []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds, uint256[] amounts)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DDC1155LogicTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicTransferBatchIterator{contract: _DDC1155Logic.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds, uint256[] amounts)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *DDC1155LogicTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicTransferBatch)
				if err := _DDC1155Logic.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds, uint256[] amounts)
func (_DDC1155Logic *DDC1155LogicFilterer) ParseTransferBatch(log types.Log) (*DDC1155LogicTransferBatch, error) {
	event := new(DDC1155LogicTransferBatch)
	if err := _DDC1155Logic.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the DDC1155Logic contract.
type DDC1155LogicTransferSingleIterator struct {
	Event *DDC1155LogicTransferSingle // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicTransferSingle)
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
		it.Event = new(DDC1155LogicTransferSingle)
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
func (it *DDC1155LogicTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicTransferSingle represents a TransferSingle event raised by the DDC1155Logic contract.
type DDC1155LogicTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	DdcId    *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 ddcId, uint256 amount)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DDC1155LogicTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicTransferSingleIterator{contract: _DDC1155Logic.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 ddcId, uint256 amount)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *DDC1155LogicTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicTransferSingle)
				if err := _DDC1155Logic.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 ddcId, uint256 amount)
func (_DDC1155Logic *DDC1155LogicFilterer) ParseTransferSingle(log types.Log) (*DDC1155LogicTransferSingle, error) {
	event := new(DDC1155LogicTransferSingle)
	if err := _DDC1155Logic.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DDC1155LogicUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the DDC1155Logic contract.
type DDC1155LogicUpgradedIterator struct {
	Event *DDC1155LogicUpgraded // Event containing the contract specifics and raw log

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
func (it *DDC1155LogicUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC1155LogicUpgraded)
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
		it.Event = new(DDC1155LogicUpgraded)
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
func (it *DDC1155LogicUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC1155LogicUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC1155LogicUpgraded represents a Upgraded event raised by the DDC1155Logic contract.
type DDC1155LogicUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC1155Logic *DDC1155LogicFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DDC1155LogicUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DDC1155Logic.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DDC1155LogicUpgradedIterator{contract: _DDC1155Logic.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC1155Logic *DDC1155LogicFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DDC1155LogicUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DDC1155Logic.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC1155LogicUpgraded)
				if err := _DDC1155Logic.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_DDC1155Logic *DDC1155LogicFilterer) ParseUpgraded(log types.Log) (*DDC1155LogicUpgraded, error) {
	event := new(DDC1155LogicUpgraded)
	if err := _DDC1155Logic.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}
