// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package runner

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
)

// ExtractorMetaData contains all meta data concerning the Extractor contract.
var ExtractorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deductTransferFee\",\"type\":\"bool\"}],\"name\":\"reflectionFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rAmount\",\"type\":\"uint256\"}],\"name\":\"tokenFromReflection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokenTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExtractorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtractorMetaData.ABI instead.
var ExtractorABI = ExtractorMetaData.ABI

// Extractor is an auto generated Go binding around an Ethereum contract.
type Extractor struct {
	ExtractorCaller     // Read-only binding to the contract
	ExtractorTransactor // Write-only binding to the contract
	ExtractorFilterer   // Log filterer for contract events
}

// ExtractorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtractorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtractorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtractorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtractorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtractorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtractorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtractorSession struct {
	Contract     *Extractor        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtractorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtractorCallerSession struct {
	Contract *ExtractorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ExtractorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtractorTransactorSession struct {
	Contract     *ExtractorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExtractorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtractorRaw struct {
	Contract *Extractor // Generic contract binding to access the raw methods on
}

// ExtractorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtractorCallerRaw struct {
	Contract *ExtractorCaller // Generic read-only contract binding to access the raw methods on
}

// ExtractorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtractorTransactorRaw struct {
	Contract *ExtractorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtractor creates a new instance of Extractor, bound to a specific deployed contract.
func NewExtractor(address common.Address, backend bind.ContractBackend) (*Extractor, error) {
	contract, err := bindExtractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Extractor{ExtractorCaller: ExtractorCaller{contract: contract}, ExtractorTransactor: ExtractorTransactor{contract: contract}, ExtractorFilterer: ExtractorFilterer{contract: contract}}, nil
}

// NewExtractorCaller creates a new read-only instance of Extractor, bound to a specific deployed contract.
func NewExtractorCaller(address common.Address, caller bind.ContractCaller) (*ExtractorCaller, error) {
	contract, err := bindExtractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtractorCaller{contract: contract}, nil
}

// NewExtractorTransactor creates a new write-only instance of Extractor, bound to a specific deployed contract.
func NewExtractorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtractorTransactor, error) {
	contract, err := bindExtractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtractorTransactor{contract: contract}, nil
}

// NewExtractorFilterer creates a new log filterer instance of Extractor, bound to a specific deployed contract.
func NewExtractorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtractorFilterer, error) {
	contract, err := bindExtractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtractorFilterer{contract: contract}, nil
}

// bindExtractor binds a generic wrapper to an already deployed contract.
func bindExtractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtractorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extractor *ExtractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extractor.Contract.ExtractorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extractor *ExtractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extractor.Contract.ExtractorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extractor *ExtractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extractor.Contract.ExtractorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extractor *ExtractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extractor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extractor *ExtractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extractor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extractor *ExtractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extractor.Contract.contract.Transact(opts, method, params...)
}

// ReflectionFromToken is a free data retrieval call binding the contract method 0x4549b039.
//
// Solidity: function reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
func (_Extractor *ExtractorCaller) ReflectionFromToken(opts *bind.CallOpts, tAmount *big.Int, deductTransferFee bool) (*big.Int, error) {
	var out []interface{}
	err := _Extractor.contract.Call(opts, &out, "reflectionFromToken", tAmount, deductTransferFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReflectionFromToken is a free data retrieval call binding the contract method 0x4549b039.
//
// Solidity: function reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
func (_Extractor *ExtractorSession) ReflectionFromToken(tAmount *big.Int, deductTransferFee bool) (*big.Int, error) {
	return _Extractor.Contract.ReflectionFromToken(&_Extractor.CallOpts, tAmount, deductTransferFee)
}

// ReflectionFromToken is a free data retrieval call binding the contract method 0x4549b039.
//
// Solidity: function reflectionFromToken(uint256 tAmount, bool deductTransferFee) view returns(uint256)
func (_Extractor *ExtractorCallerSession) ReflectionFromToken(tAmount *big.Int, deductTransferFee bool) (*big.Int, error) {
	return _Extractor.Contract.ReflectionFromToken(&_Extractor.CallOpts, tAmount, deductTransferFee)
}

// TokenFromReflection is a free data retrieval call binding the contract method 0x2d838119.
//
// Solidity: function tokenFromReflection(uint256 rAmount) view returns(uint256)
func (_Extractor *ExtractorCaller) TokenFromReflection(opts *bind.CallOpts, rAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extractor.contract.Call(opts, &out, "tokenFromReflection", rAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenFromReflection is a free data retrieval call binding the contract method 0x2d838119.
//
// Solidity: function tokenFromReflection(uint256 rAmount) view returns(uint256)
func (_Extractor *ExtractorSession) TokenFromReflection(rAmount *big.Int) (*big.Int, error) {
	return _Extractor.Contract.TokenFromReflection(&_Extractor.CallOpts, rAmount)
}

// TokenFromReflection is a free data retrieval call binding the contract method 0x2d838119.
//
// Solidity: function tokenFromReflection(uint256 rAmount) view returns(uint256)
func (_Extractor *ExtractorCallerSession) TokenFromReflection(rAmount *big.Int) (*big.Int, error) {
	return _Extractor.Contract.TokenFromReflection(&_Extractor.CallOpts, rAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Extractor *ExtractorTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Extractor *ExtractorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.Withdraw(&_Extractor.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.Withdraw(&_Extractor.TransactOpts, amount)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Extractor *ExtractorTransactor) Withdraw0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdraw0")
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Extractor *ExtractorSession) Withdraw0() (*types.Transaction, error) {
	return _Extractor.Contract.Withdraw0(&_Extractor.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Extractor *ExtractorTransactorSession) Withdraw0() (*types.Transaction, error) {
	return _Extractor.Contract.Withdraw0(&_Extractor.TransactOpts)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactor) Withdraw1(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdraw1", to, amount)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_Extractor *ExtractorSession) Withdraw1(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.Withdraw1(&_Extractor.TransactOpts, to, amount)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) Withdraw1(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.Withdraw1(&_Extractor.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Extractor *ExtractorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawETH(&_Extractor.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawETH(&_Extractor.TransactOpts, to, amount)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Extractor *ExtractorTransactor) WithdrawETH0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawETH0")
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Extractor *ExtractorSession) WithdrawETH0() (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawETH0(&_Extractor.TransactOpts)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Extractor *ExtractorTransactorSession) WithdrawETH0() (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawETH0(&_Extractor.TransactOpts)
}

// WithdrawETH1 is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Extractor *ExtractorTransactor) WithdrawETH1(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawETH1", amount)
}

// WithdrawETH1 is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Extractor *ExtractorSession) WithdrawETH1(amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawETH1(&_Extractor.TransactOpts, amount)
}

// WithdrawETH1 is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) WithdrawETH1(amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawETH1(&_Extractor.TransactOpts, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawTo", to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_Extractor *ExtractorSession) WithdrawTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawTo(&_Extractor.TransactOpts, to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) WithdrawTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawTo(&_Extractor.TransactOpts, to, amount)
}

// WithdrawTo0 is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactor) WithdrawTo0(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawTo0", token, to, amount)
}

// WithdrawTo0 is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorSession) WithdrawTo0(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawTo0(&_Extractor.TransactOpts, token, to, amount)
}

// WithdrawTo0 is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) WithdrawTo0(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawTo0(&_Extractor.TransactOpts, token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawToken", token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawToken(&_Extractor.TransactOpts, token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawToken(&_Extractor.TransactOpts, token, to, amount)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0x223c217b.
//
// Solidity: function withdrawTokenTo(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactor) WithdrawTokenTo(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.contract.Transact(opts, "withdrawTokenTo", token, to, amount)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0x223c217b.
//
// Solidity: function withdrawTokenTo(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorSession) WithdrawTokenTo(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawTokenTo(&_Extractor.TransactOpts, token, to, amount)
}

// WithdrawTokenTo is a paid mutator transaction binding the contract method 0x223c217b.
//
// Solidity: function withdrawTokenTo(address token, address to, uint256 amount) returns()
func (_Extractor *ExtractorTransactorSession) WithdrawTokenTo(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Extractor.Contract.WithdrawTokenTo(&_Extractor.TransactOpts, token, to, amount)
}
