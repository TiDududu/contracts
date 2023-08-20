// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplemessage

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SimplemessageMetaData contains all meta data concerning the Simplemessage contract.
var SimplemessageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SimplemessageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplemessageMetaData.ABI instead.
var SimplemessageABI = SimplemessageMetaData.ABI

// Simplemessage is an auto generated Go binding around an Ethereum contract.
type Simplemessage struct {
	SimplemessageCaller     // Read-only binding to the contract
	SimplemessageTransactor // Write-only binding to the contract
	SimplemessageFilterer   // Log filterer for contract events
}

// SimplemessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplemessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplemessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplemessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplemessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplemessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplemessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplemessageSession struct {
	Contract     *Simplemessage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplemessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplemessageCallerSession struct {
	Contract *SimplemessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimplemessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplemessageTransactorSession struct {
	Contract     *SimplemessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimplemessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplemessageRaw struct {
	Contract *Simplemessage // Generic contract binding to access the raw methods on
}

// SimplemessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplemessageCallerRaw struct {
	Contract *SimplemessageCaller // Generic read-only contract binding to access the raw methods on
}

// SimplemessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplemessageTransactorRaw struct {
	Contract *SimplemessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplemessage creates a new instance of Simplemessage, bound to a specific deployed contract.
func NewSimplemessage(address common.Address, backend bind.ContractBackend) (*Simplemessage, error) {
	contract, err := bindSimplemessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplemessage{SimplemessageCaller: SimplemessageCaller{contract: contract}, SimplemessageTransactor: SimplemessageTransactor{contract: contract}, SimplemessageFilterer: SimplemessageFilterer{contract: contract}}, nil
}

// NewSimplemessageCaller creates a new read-only instance of Simplemessage, bound to a specific deployed contract.
func NewSimplemessageCaller(address common.Address, caller bind.ContractCaller) (*SimplemessageCaller, error) {
	contract, err := bindSimplemessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplemessageCaller{contract: contract}, nil
}

// NewSimplemessageTransactor creates a new write-only instance of Simplemessage, bound to a specific deployed contract.
func NewSimplemessageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplemessageTransactor, error) {
	contract, err := bindSimplemessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplemessageTransactor{contract: contract}, nil
}

// NewSimplemessageFilterer creates a new log filterer instance of Simplemessage, bound to a specific deployed contract.
func NewSimplemessageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplemessageFilterer, error) {
	contract, err := bindSimplemessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplemessageFilterer{contract: contract}, nil
}

// bindSimplemessage binds a generic wrapper to an already deployed contract.
func bindSimplemessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplemessageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplemessage *SimplemessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplemessage.Contract.SimplemessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplemessage *SimplemessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplemessage.Contract.SimplemessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplemessage *SimplemessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplemessage.Contract.SimplemessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplemessage *SimplemessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplemessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplemessage *SimplemessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplemessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplemessage *SimplemessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplemessage.Contract.contract.Transact(opts, method, params...)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_Simplemessage *SimplemessageCaller) Message(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Simplemessage.contract.Call(opts, &out, "message")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_Simplemessage *SimplemessageSession) Message() (string, error) {
	return _Simplemessage.Contract.Message(&_Simplemessage.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() view returns(string)
func (_Simplemessage *SimplemessageCallerSession) Message() (string, error) {
	return _Simplemessage.Contract.Message(&_Simplemessage.CallOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _message) returns()
func (_Simplemessage *SimplemessageTransactor) SetMessage(opts *bind.TransactOpts, _message string) (*types.Transaction, error) {
	return _Simplemessage.contract.Transact(opts, "setMessage", _message)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _message) returns()
func (_Simplemessage *SimplemessageSession) SetMessage(_message string) (*types.Transaction, error) {
	return _Simplemessage.Contract.SetMessage(&_Simplemessage.TransactOpts, _message)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _message) returns()
func (_Simplemessage *SimplemessageTransactorSession) SetMessage(_message string) (*types.Transaction, error) {
	return _Simplemessage.Contract.SetMessage(&_Simplemessage.TransactOpts, _message)
}
