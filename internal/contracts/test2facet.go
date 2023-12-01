// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Test2FacetMetaData contains all meta data concerning the Test2Facet contract.
var Test2FacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"test2Func1\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func10\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func11\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func12\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func13\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func14\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func15\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func16\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func17\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func18\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func19\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func3\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func4\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func5\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func6\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func7\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func8\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test2Func9\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610169806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80638ee8be30116100ad578063d2f0c73e11610071578063d2f0c73e14610131578063e5f687b214610131578063e7de23a414610131578063ea36b55814610131578063ef3f4d781461013157600080fd5b80638ee8be301461013157806391d0396b14610131578063c670641d14610131578063ca5fa5c014610131578063caae8f231461013157600080fd5b80632e463958116100f45780632e463958146101315780635fd6312b146101315780636dc16b0114610131578063792a8e2e14610131578063884280a61461013157600080fd5b806303feeeae146101315780630c103a93146101315780630e4cd7fc14610131578063148843091461013157806317fd06e714610131575b600080fd5b00fea2646970667358221220cb8383a24087045634e5736a3e6ebb3594effa6b1d8856cf25dccbca3249463b64736f6c63430008150033",
}

// Test2FacetABI is the input ABI used to generate the binding from.
// Deprecated: Use Test2FacetMetaData.ABI instead.
var Test2FacetABI = Test2FacetMetaData.ABI

// Test2FacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Test2FacetMetaData.Bin instead.
var Test2FacetBin = Test2FacetMetaData.Bin

// DeployTest2Facet deploys a new Ethereum contract, binding an instance of Test2Facet to it.
func DeployTest2Facet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test2Facet, error) {
	parsed, err := Test2FacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Test2FacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test2Facet{Test2FacetCaller: Test2FacetCaller{contract: contract}, Test2FacetTransactor: Test2FacetTransactor{contract: contract}, Test2FacetFilterer: Test2FacetFilterer{contract: contract}}, nil
}

// Test2Facet is an auto generated Go binding around an Ethereum contract.
type Test2Facet struct {
	Test2FacetCaller     // Read-only binding to the contract
	Test2FacetTransactor // Write-only binding to the contract
	Test2FacetFilterer   // Log filterer for contract events
}

// Test2FacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type Test2FacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2FacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Test2FacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2FacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Test2FacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2FacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Test2FacetSession struct {
	Contract     *Test2Facet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Test2FacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Test2FacetCallerSession struct {
	Contract *Test2FacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Test2FacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Test2FacetTransactorSession struct {
	Contract     *Test2FacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Test2FacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type Test2FacetRaw struct {
	Contract *Test2Facet // Generic contract binding to access the raw methods on
}

// Test2FacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Test2FacetCallerRaw struct {
	Contract *Test2FacetCaller // Generic read-only contract binding to access the raw methods on
}

// Test2FacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Test2FacetTransactorRaw struct {
	Contract *Test2FacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest2Facet creates a new instance of Test2Facet, bound to a specific deployed contract.
func NewTest2Facet(address common.Address, backend bind.ContractBackend) (*Test2Facet, error) {
	contract, err := bindTest2Facet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test2Facet{Test2FacetCaller: Test2FacetCaller{contract: contract}, Test2FacetTransactor: Test2FacetTransactor{contract: contract}, Test2FacetFilterer: Test2FacetFilterer{contract: contract}}, nil
}

// NewTest2FacetCaller creates a new read-only instance of Test2Facet, bound to a specific deployed contract.
func NewTest2FacetCaller(address common.Address, caller bind.ContractCaller) (*Test2FacetCaller, error) {
	contract, err := bindTest2Facet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Test2FacetCaller{contract: contract}, nil
}

// NewTest2FacetTransactor creates a new write-only instance of Test2Facet, bound to a specific deployed contract.
func NewTest2FacetTransactor(address common.Address, transactor bind.ContractTransactor) (*Test2FacetTransactor, error) {
	contract, err := bindTest2Facet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Test2FacetTransactor{contract: contract}, nil
}

// NewTest2FacetFilterer creates a new log filterer instance of Test2Facet, bound to a specific deployed contract.
func NewTest2FacetFilterer(address common.Address, filterer bind.ContractFilterer) (*Test2FacetFilterer, error) {
	contract, err := bindTest2Facet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Test2FacetFilterer{contract: contract}, nil
}

// bindTest2Facet binds a generic wrapper to an already deployed contract.
func bindTest2Facet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Test2FacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test2Facet *Test2FacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test2Facet.Contract.Test2FacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test2Facet *Test2FacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2FacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test2Facet *Test2FacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2FacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test2Facet *Test2FacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test2Facet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test2Facet *Test2FacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test2Facet *Test2FacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test2Facet.Contract.contract.Transact(opts, method, params...)
}

// Test2Func1 is a paid mutator transaction binding the contract method 0xea36b558.
//
// Solidity: function test2Func1() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func1")
}

// Test2Func1 is a paid mutator transaction binding the contract method 0xea36b558.
//
// Solidity: function test2Func1() returns()
func (_Test2Facet *Test2FacetSession) Test2Func1() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func1(&_Test2Facet.TransactOpts)
}

// Test2Func1 is a paid mutator transaction binding the contract method 0xea36b558.
//
// Solidity: function test2Func1() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func1() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func1(&_Test2Facet.TransactOpts)
}

// Test2Func10 is a paid mutator transaction binding the contract method 0x8ee8be30.
//
// Solidity: function test2Func10() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func10(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func10")
}

// Test2Func10 is a paid mutator transaction binding the contract method 0x8ee8be30.
//
// Solidity: function test2Func10() returns()
func (_Test2Facet *Test2FacetSession) Test2Func10() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func10(&_Test2Facet.TransactOpts)
}

// Test2Func10 is a paid mutator transaction binding the contract method 0x8ee8be30.
//
// Solidity: function test2Func10() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func10() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func10(&_Test2Facet.TransactOpts)
}

// Test2Func11 is a paid mutator transaction binding the contract method 0x884280a6.
//
// Solidity: function test2Func11() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func11(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func11")
}

// Test2Func11 is a paid mutator transaction binding the contract method 0x884280a6.
//
// Solidity: function test2Func11() returns()
func (_Test2Facet *Test2FacetSession) Test2Func11() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func11(&_Test2Facet.TransactOpts)
}

// Test2Func11 is a paid mutator transaction binding the contract method 0x884280a6.
//
// Solidity: function test2Func11() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func11() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func11(&_Test2Facet.TransactOpts)
}

// Test2Func12 is a paid mutator transaction binding the contract method 0xca5fa5c0.
//
// Solidity: function test2Func12() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func12(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func12")
}

// Test2Func12 is a paid mutator transaction binding the contract method 0xca5fa5c0.
//
// Solidity: function test2Func12() returns()
func (_Test2Facet *Test2FacetSession) Test2Func12() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func12(&_Test2Facet.TransactOpts)
}

// Test2Func12 is a paid mutator transaction binding the contract method 0xca5fa5c0.
//
// Solidity: function test2Func12() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func12() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func12(&_Test2Facet.TransactOpts)
}

// Test2Func13 is a paid mutator transaction binding the contract method 0x6dc16b01.
//
// Solidity: function test2Func13() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func13(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func13")
}

// Test2Func13 is a paid mutator transaction binding the contract method 0x6dc16b01.
//
// Solidity: function test2Func13() returns()
func (_Test2Facet *Test2FacetSession) Test2Func13() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func13(&_Test2Facet.TransactOpts)
}

// Test2Func13 is a paid mutator transaction binding the contract method 0x6dc16b01.
//
// Solidity: function test2Func13() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func13() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func13(&_Test2Facet.TransactOpts)
}

// Test2Func14 is a paid mutator transaction binding the contract method 0x91d0396b.
//
// Solidity: function test2Func14() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func14(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func14")
}

// Test2Func14 is a paid mutator transaction binding the contract method 0x91d0396b.
//
// Solidity: function test2Func14() returns()
func (_Test2Facet *Test2FacetSession) Test2Func14() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func14(&_Test2Facet.TransactOpts)
}

// Test2Func14 is a paid mutator transaction binding the contract method 0x91d0396b.
//
// Solidity: function test2Func14() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func14() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func14(&_Test2Facet.TransactOpts)
}

// Test2Func15 is a paid mutator transaction binding the contract method 0x03feeeae.
//
// Solidity: function test2Func15() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func15(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func15")
}

// Test2Func15 is a paid mutator transaction binding the contract method 0x03feeeae.
//
// Solidity: function test2Func15() returns()
func (_Test2Facet *Test2FacetSession) Test2Func15() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func15(&_Test2Facet.TransactOpts)
}

// Test2Func15 is a paid mutator transaction binding the contract method 0x03feeeae.
//
// Solidity: function test2Func15() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func15() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func15(&_Test2Facet.TransactOpts)
}

// Test2Func16 is a paid mutator transaction binding the contract method 0x2e463958.
//
// Solidity: function test2Func16() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func16(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func16")
}

// Test2Func16 is a paid mutator transaction binding the contract method 0x2e463958.
//
// Solidity: function test2Func16() returns()
func (_Test2Facet *Test2FacetSession) Test2Func16() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func16(&_Test2Facet.TransactOpts)
}

// Test2Func16 is a paid mutator transaction binding the contract method 0x2e463958.
//
// Solidity: function test2Func16() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func16() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func16(&_Test2Facet.TransactOpts)
}

// Test2Func17 is a paid mutator transaction binding the contract method 0x14884309.
//
// Solidity: function test2Func17() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func17(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func17")
}

// Test2Func17 is a paid mutator transaction binding the contract method 0x14884309.
//
// Solidity: function test2Func17() returns()
func (_Test2Facet *Test2FacetSession) Test2Func17() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func17(&_Test2Facet.TransactOpts)
}

// Test2Func17 is a paid mutator transaction binding the contract method 0x14884309.
//
// Solidity: function test2Func17() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func17() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func17(&_Test2Facet.TransactOpts)
}

// Test2Func18 is a paid mutator transaction binding the contract method 0x0c103a93.
//
// Solidity: function test2Func18() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func18(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func18")
}

// Test2Func18 is a paid mutator transaction binding the contract method 0x0c103a93.
//
// Solidity: function test2Func18() returns()
func (_Test2Facet *Test2FacetSession) Test2Func18() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func18(&_Test2Facet.TransactOpts)
}

// Test2Func18 is a paid mutator transaction binding the contract method 0x0c103a93.
//
// Solidity: function test2Func18() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func18() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func18(&_Test2Facet.TransactOpts)
}

// Test2Func19 is a paid mutator transaction binding the contract method 0x5fd6312b.
//
// Solidity: function test2Func19() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func19(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func19")
}

// Test2Func19 is a paid mutator transaction binding the contract method 0x5fd6312b.
//
// Solidity: function test2Func19() returns()
func (_Test2Facet *Test2FacetSession) Test2Func19() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func19(&_Test2Facet.TransactOpts)
}

// Test2Func19 is a paid mutator transaction binding the contract method 0x5fd6312b.
//
// Solidity: function test2Func19() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func19() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func19(&_Test2Facet.TransactOpts)
}

// Test2Func2 is a paid mutator transaction binding the contract method 0xe7de23a4.
//
// Solidity: function test2Func2() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func2")
}

// Test2Func2 is a paid mutator transaction binding the contract method 0xe7de23a4.
//
// Solidity: function test2Func2() returns()
func (_Test2Facet *Test2FacetSession) Test2Func2() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func2(&_Test2Facet.TransactOpts)
}

// Test2Func2 is a paid mutator transaction binding the contract method 0xe7de23a4.
//
// Solidity: function test2Func2() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func2() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func2(&_Test2Facet.TransactOpts)
}

// Test2Func20 is a paid mutator transaction binding the contract method 0x792a8e2e.
//
// Solidity: function test2Func20() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func20")
}

// Test2Func20 is a paid mutator transaction binding the contract method 0x792a8e2e.
//
// Solidity: function test2Func20() returns()
func (_Test2Facet *Test2FacetSession) Test2Func20() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func20(&_Test2Facet.TransactOpts)
}

// Test2Func20 is a paid mutator transaction binding the contract method 0x792a8e2e.
//
// Solidity: function test2Func20() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func20() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func20(&_Test2Facet.TransactOpts)
}

// Test2Func3 is a paid mutator transaction binding the contract method 0x0e4cd7fc.
//
// Solidity: function test2Func3() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func3(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func3")
}

// Test2Func3 is a paid mutator transaction binding the contract method 0x0e4cd7fc.
//
// Solidity: function test2Func3() returns()
func (_Test2Facet *Test2FacetSession) Test2Func3() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func3(&_Test2Facet.TransactOpts)
}

// Test2Func3 is a paid mutator transaction binding the contract method 0x0e4cd7fc.
//
// Solidity: function test2Func3() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func3() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func3(&_Test2Facet.TransactOpts)
}

// Test2Func4 is a paid mutator transaction binding the contract method 0xc670641d.
//
// Solidity: function test2Func4() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func4(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func4")
}

// Test2Func4 is a paid mutator transaction binding the contract method 0xc670641d.
//
// Solidity: function test2Func4() returns()
func (_Test2Facet *Test2FacetSession) Test2Func4() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func4(&_Test2Facet.TransactOpts)
}

// Test2Func4 is a paid mutator transaction binding the contract method 0xc670641d.
//
// Solidity: function test2Func4() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func4() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func4(&_Test2Facet.TransactOpts)
}

// Test2Func5 is a paid mutator transaction binding the contract method 0xd2f0c73e.
//
// Solidity: function test2Func5() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func5(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func5")
}

// Test2Func5 is a paid mutator transaction binding the contract method 0xd2f0c73e.
//
// Solidity: function test2Func5() returns()
func (_Test2Facet *Test2FacetSession) Test2Func5() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func5(&_Test2Facet.TransactOpts)
}

// Test2Func5 is a paid mutator transaction binding the contract method 0xd2f0c73e.
//
// Solidity: function test2Func5() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func5() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func5(&_Test2Facet.TransactOpts)
}

// Test2Func6 is a paid mutator transaction binding the contract method 0x17fd06e7.
//
// Solidity: function test2Func6() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func6(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func6")
}

// Test2Func6 is a paid mutator transaction binding the contract method 0x17fd06e7.
//
// Solidity: function test2Func6() returns()
func (_Test2Facet *Test2FacetSession) Test2Func6() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func6(&_Test2Facet.TransactOpts)
}

// Test2Func6 is a paid mutator transaction binding the contract method 0x17fd06e7.
//
// Solidity: function test2Func6() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func6() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func6(&_Test2Facet.TransactOpts)
}

// Test2Func7 is a paid mutator transaction binding the contract method 0xef3f4d78.
//
// Solidity: function test2Func7() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func7(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func7")
}

// Test2Func7 is a paid mutator transaction binding the contract method 0xef3f4d78.
//
// Solidity: function test2Func7() returns()
func (_Test2Facet *Test2FacetSession) Test2Func7() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func7(&_Test2Facet.TransactOpts)
}

// Test2Func7 is a paid mutator transaction binding the contract method 0xef3f4d78.
//
// Solidity: function test2Func7() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func7() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func7(&_Test2Facet.TransactOpts)
}

// Test2Func8 is a paid mutator transaction binding the contract method 0xe5f687b2.
//
// Solidity: function test2Func8() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func8(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func8")
}

// Test2Func8 is a paid mutator transaction binding the contract method 0xe5f687b2.
//
// Solidity: function test2Func8() returns()
func (_Test2Facet *Test2FacetSession) Test2Func8() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func8(&_Test2Facet.TransactOpts)
}

// Test2Func8 is a paid mutator transaction binding the contract method 0xe5f687b2.
//
// Solidity: function test2Func8() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func8() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func8(&_Test2Facet.TransactOpts)
}

// Test2Func9 is a paid mutator transaction binding the contract method 0xcaae8f23.
//
// Solidity: function test2Func9() returns()
func (_Test2Facet *Test2FacetTransactor) Test2Func9(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.contract.Transact(opts, "test2Func9")
}

// Test2Func9 is a paid mutator transaction binding the contract method 0xcaae8f23.
//
// Solidity: function test2Func9() returns()
func (_Test2Facet *Test2FacetSession) Test2Func9() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func9(&_Test2Facet.TransactOpts)
}

// Test2Func9 is a paid mutator transaction binding the contract method 0xcaae8f23.
//
// Solidity: function test2Func9() returns()
func (_Test2Facet *Test2FacetTransactorSession) Test2Func9() (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2Func9(&_Test2Facet.TransactOpts)
}
