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

// Test1FacetMetaData contains all meta data concerning the Test1Facet contract.
var Test1FacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"test1Func1\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func10\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func11\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func12\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func13\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func14\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func15\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func16\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func17\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func18\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func19\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func3\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func4\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func5\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func6\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func7\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func8\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"test1Func9\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TestEvent\",\"inputs\":[{\"name\":\"something\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506101cd806100206000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806351b68a4d116100b85780639abf97aa1161007c5780639abf97aa14610164578063b0e8fcc714610164578063cbb835f614610164578063cd0bae0914610164578063cf3bbe1814610164578063db32da151461016457600080fd5b806351b68a4d1461016457806371a99d6f1461016457806377e9d0d61461016457806381b5207d1461016457806387952d221461016457600080fd5b806324c1d5a7116100ff57806324c1d5a714610164578063292c460d146101645780632cb83248146101645780634484b3b91461016457806350eb3f431461016457600080fd5b806301ffc9a71461013c5780630716c2ae14610164578063110460471461016457806319c841ab1461016457806319e3b53314610164575b600080fd5b61015061014a366004610166565b50600090565b604051901515815260200160405180910390f35b005b60006020828403121561017857600080fd5b81356001600160e01b03198116811461019057600080fd5b939250505056fea26469706673582212209e48941288a9fd5638d5c68e9457319fd30c6a65ec80df23e9d5506f396828e364736f6c63430008150033",
}

// Test1FacetABI is the input ABI used to generate the binding from.
// Deprecated: Use Test1FacetMetaData.ABI instead.
var Test1FacetABI = Test1FacetMetaData.ABI

// Test1FacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Test1FacetMetaData.Bin instead.
var Test1FacetBin = Test1FacetMetaData.Bin

// DeployTest1Facet deploys a new Ethereum contract, binding an instance of Test1Facet to it.
func DeployTest1Facet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test1Facet, error) {
	parsed, err := Test1FacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Test1FacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test1Facet{Test1FacetCaller: Test1FacetCaller{contract: contract}, Test1FacetTransactor: Test1FacetTransactor{contract: contract}, Test1FacetFilterer: Test1FacetFilterer{contract: contract}}, nil
}

// Test1Facet is an auto generated Go binding around an Ethereum contract.
type Test1Facet struct {
	Test1FacetCaller     // Read-only binding to the contract
	Test1FacetTransactor // Write-only binding to the contract
	Test1FacetFilterer   // Log filterer for contract events
}

// Test1FacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type Test1FacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test1FacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Test1FacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test1FacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Test1FacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test1FacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Test1FacetSession struct {
	Contract     *Test1Facet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Test1FacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Test1FacetCallerSession struct {
	Contract *Test1FacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Test1FacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Test1FacetTransactorSession struct {
	Contract     *Test1FacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Test1FacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type Test1FacetRaw struct {
	Contract *Test1Facet // Generic contract binding to access the raw methods on
}

// Test1FacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Test1FacetCallerRaw struct {
	Contract *Test1FacetCaller // Generic read-only contract binding to access the raw methods on
}

// Test1FacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Test1FacetTransactorRaw struct {
	Contract *Test1FacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest1Facet creates a new instance of Test1Facet, bound to a specific deployed contract.
func NewTest1Facet(address common.Address, backend bind.ContractBackend) (*Test1Facet, error) {
	contract, err := bindTest1Facet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test1Facet{Test1FacetCaller: Test1FacetCaller{contract: contract}, Test1FacetTransactor: Test1FacetTransactor{contract: contract}, Test1FacetFilterer: Test1FacetFilterer{contract: contract}}, nil
}

// NewTest1FacetCaller creates a new read-only instance of Test1Facet, bound to a specific deployed contract.
func NewTest1FacetCaller(address common.Address, caller bind.ContractCaller) (*Test1FacetCaller, error) {
	contract, err := bindTest1Facet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Test1FacetCaller{contract: contract}, nil
}

// NewTest1FacetTransactor creates a new write-only instance of Test1Facet, bound to a specific deployed contract.
func NewTest1FacetTransactor(address common.Address, transactor bind.ContractTransactor) (*Test1FacetTransactor, error) {
	contract, err := bindTest1Facet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Test1FacetTransactor{contract: contract}, nil
}

// NewTest1FacetFilterer creates a new log filterer instance of Test1Facet, bound to a specific deployed contract.
func NewTest1FacetFilterer(address common.Address, filterer bind.ContractFilterer) (*Test1FacetFilterer, error) {
	contract, err := bindTest1Facet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Test1FacetFilterer{contract: contract}, nil
}

// bindTest1Facet binds a generic wrapper to an already deployed contract.
func bindTest1Facet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Test1FacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test1Facet *Test1FacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test1Facet.Contract.Test1FacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test1Facet *Test1FacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1FacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test1Facet *Test1FacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1FacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test1Facet *Test1FacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test1Facet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test1Facet *Test1FacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test1Facet *Test1FacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test1Facet.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Test1Facet *Test1FacetCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Test1Facet *Test1FacetSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Test1Facet.Contract.SupportsInterface(&_Test1Facet.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_Test1Facet *Test1FacetCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _Test1Facet.Contract.SupportsInterface(&_Test1Facet.CallOpts, _interfaceID)
}

// Test1Func1 is a paid mutator transaction binding the contract method 0x19e3b533.
//
// Solidity: function test1Func1() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func1")
}

// Test1Func1 is a paid mutator transaction binding the contract method 0x19e3b533.
//
// Solidity: function test1Func1() returns()
func (_Test1Facet *Test1FacetSession) Test1Func1() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func1(&_Test1Facet.TransactOpts)
}

// Test1Func1 is a paid mutator transaction binding the contract method 0x19e3b533.
//
// Solidity: function test1Func1() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func1() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func1(&_Test1Facet.TransactOpts)
}

// Test1Func10 is a paid mutator transaction binding the contract method 0x87952d22.
//
// Solidity: function test1Func10() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func10(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func10")
}

// Test1Func10 is a paid mutator transaction binding the contract method 0x87952d22.
//
// Solidity: function test1Func10() returns()
func (_Test1Facet *Test1FacetSession) Test1Func10() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func10(&_Test1Facet.TransactOpts)
}

// Test1Func10 is a paid mutator transaction binding the contract method 0x87952d22.
//
// Solidity: function test1Func10() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func10() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func10(&_Test1Facet.TransactOpts)
}

// Test1Func11 is a paid mutator transaction binding the contract method 0x50eb3f43.
//
// Solidity: function test1Func11() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func11(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func11")
}

// Test1Func11 is a paid mutator transaction binding the contract method 0x50eb3f43.
//
// Solidity: function test1Func11() returns()
func (_Test1Facet *Test1FacetSession) Test1Func11() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func11(&_Test1Facet.TransactOpts)
}

// Test1Func11 is a paid mutator transaction binding the contract method 0x50eb3f43.
//
// Solidity: function test1Func11() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func11() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func11(&_Test1Facet.TransactOpts)
}

// Test1Func12 is a paid mutator transaction binding the contract method 0x81b5207d.
//
// Solidity: function test1Func12() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func12(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func12")
}

// Test1Func12 is a paid mutator transaction binding the contract method 0x81b5207d.
//
// Solidity: function test1Func12() returns()
func (_Test1Facet *Test1FacetSession) Test1Func12() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func12(&_Test1Facet.TransactOpts)
}

// Test1Func12 is a paid mutator transaction binding the contract method 0x81b5207d.
//
// Solidity: function test1Func12() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func12() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func12(&_Test1Facet.TransactOpts)
}

// Test1Func13 is a paid mutator transaction binding the contract method 0x19c841ab.
//
// Solidity: function test1Func13() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func13(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func13")
}

// Test1Func13 is a paid mutator transaction binding the contract method 0x19c841ab.
//
// Solidity: function test1Func13() returns()
func (_Test1Facet *Test1FacetSession) Test1Func13() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func13(&_Test1Facet.TransactOpts)
}

// Test1Func13 is a paid mutator transaction binding the contract method 0x19c841ab.
//
// Solidity: function test1Func13() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func13() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func13(&_Test1Facet.TransactOpts)
}

// Test1Func14 is a paid mutator transaction binding the contract method 0x51b68a4d.
//
// Solidity: function test1Func14() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func14(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func14")
}

// Test1Func14 is a paid mutator transaction binding the contract method 0x51b68a4d.
//
// Solidity: function test1Func14() returns()
func (_Test1Facet *Test1FacetSession) Test1Func14() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func14(&_Test1Facet.TransactOpts)
}

// Test1Func14 is a paid mutator transaction binding the contract method 0x51b68a4d.
//
// Solidity: function test1Func14() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func14() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func14(&_Test1Facet.TransactOpts)
}

// Test1Func15 is a paid mutator transaction binding the contract method 0x2cb83248.
//
// Solidity: function test1Func15() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func15(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func15")
}

// Test1Func15 is a paid mutator transaction binding the contract method 0x2cb83248.
//
// Solidity: function test1Func15() returns()
func (_Test1Facet *Test1FacetSession) Test1Func15() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func15(&_Test1Facet.TransactOpts)
}

// Test1Func15 is a paid mutator transaction binding the contract method 0x2cb83248.
//
// Solidity: function test1Func15() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func15() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func15(&_Test1Facet.TransactOpts)
}

// Test1Func16 is a paid mutator transaction binding the contract method 0x77e9d0d6.
//
// Solidity: function test1Func16() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func16(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func16")
}

// Test1Func16 is a paid mutator transaction binding the contract method 0x77e9d0d6.
//
// Solidity: function test1Func16() returns()
func (_Test1Facet *Test1FacetSession) Test1Func16() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func16(&_Test1Facet.TransactOpts)
}

// Test1Func16 is a paid mutator transaction binding the contract method 0x77e9d0d6.
//
// Solidity: function test1Func16() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func16() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func16(&_Test1Facet.TransactOpts)
}

// Test1Func17 is a paid mutator transaction binding the contract method 0x4484b3b9.
//
// Solidity: function test1Func17() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func17(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func17")
}

// Test1Func17 is a paid mutator transaction binding the contract method 0x4484b3b9.
//
// Solidity: function test1Func17() returns()
func (_Test1Facet *Test1FacetSession) Test1Func17() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func17(&_Test1Facet.TransactOpts)
}

// Test1Func17 is a paid mutator transaction binding the contract method 0x4484b3b9.
//
// Solidity: function test1Func17() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func17() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func17(&_Test1Facet.TransactOpts)
}

// Test1Func18 is a paid mutator transaction binding the contract method 0x9abf97aa.
//
// Solidity: function test1Func18() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func18(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func18")
}

// Test1Func18 is a paid mutator transaction binding the contract method 0x9abf97aa.
//
// Solidity: function test1Func18() returns()
func (_Test1Facet *Test1FacetSession) Test1Func18() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func18(&_Test1Facet.TransactOpts)
}

// Test1Func18 is a paid mutator transaction binding the contract method 0x9abf97aa.
//
// Solidity: function test1Func18() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func18() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func18(&_Test1Facet.TransactOpts)
}

// Test1Func19 is a paid mutator transaction binding the contract method 0x292c460d.
//
// Solidity: function test1Func19() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func19(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func19")
}

// Test1Func19 is a paid mutator transaction binding the contract method 0x292c460d.
//
// Solidity: function test1Func19() returns()
func (_Test1Facet *Test1FacetSession) Test1Func19() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func19(&_Test1Facet.TransactOpts)
}

// Test1Func19 is a paid mutator transaction binding the contract method 0x292c460d.
//
// Solidity: function test1Func19() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func19() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func19(&_Test1Facet.TransactOpts)
}

// Test1Func2 is a paid mutator transaction binding the contract method 0x0716c2ae.
//
// Solidity: function test1Func2() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func2")
}

// Test1Func2 is a paid mutator transaction binding the contract method 0x0716c2ae.
//
// Solidity: function test1Func2() returns()
func (_Test1Facet *Test1FacetSession) Test1Func2() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func2(&_Test1Facet.TransactOpts)
}

// Test1Func2 is a paid mutator transaction binding the contract method 0x0716c2ae.
//
// Solidity: function test1Func2() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func2() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func2(&_Test1Facet.TransactOpts)
}

// Test1Func20 is a paid mutator transaction binding the contract method 0xb0e8fcc7.
//
// Solidity: function test1Func20() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func20")
}

// Test1Func20 is a paid mutator transaction binding the contract method 0xb0e8fcc7.
//
// Solidity: function test1Func20() returns()
func (_Test1Facet *Test1FacetSession) Test1Func20() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func20(&_Test1Facet.TransactOpts)
}

// Test1Func20 is a paid mutator transaction binding the contract method 0xb0e8fcc7.
//
// Solidity: function test1Func20() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func20() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func20(&_Test1Facet.TransactOpts)
}

// Test1Func3 is a paid mutator transaction binding the contract method 0x11046047.
//
// Solidity: function test1Func3() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func3(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func3")
}

// Test1Func3 is a paid mutator transaction binding the contract method 0x11046047.
//
// Solidity: function test1Func3() returns()
func (_Test1Facet *Test1FacetSession) Test1Func3() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func3(&_Test1Facet.TransactOpts)
}

// Test1Func3 is a paid mutator transaction binding the contract method 0x11046047.
//
// Solidity: function test1Func3() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func3() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func3(&_Test1Facet.TransactOpts)
}

// Test1Func4 is a paid mutator transaction binding the contract method 0xcf3bbe18.
//
// Solidity: function test1Func4() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func4(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func4")
}

// Test1Func4 is a paid mutator transaction binding the contract method 0xcf3bbe18.
//
// Solidity: function test1Func4() returns()
func (_Test1Facet *Test1FacetSession) Test1Func4() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func4(&_Test1Facet.TransactOpts)
}

// Test1Func4 is a paid mutator transaction binding the contract method 0xcf3bbe18.
//
// Solidity: function test1Func4() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func4() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func4(&_Test1Facet.TransactOpts)
}

// Test1Func5 is a paid mutator transaction binding the contract method 0x24c1d5a7.
//
// Solidity: function test1Func5() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func5(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func5")
}

// Test1Func5 is a paid mutator transaction binding the contract method 0x24c1d5a7.
//
// Solidity: function test1Func5() returns()
func (_Test1Facet *Test1FacetSession) Test1Func5() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func5(&_Test1Facet.TransactOpts)
}

// Test1Func5 is a paid mutator transaction binding the contract method 0x24c1d5a7.
//
// Solidity: function test1Func5() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func5() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func5(&_Test1Facet.TransactOpts)
}

// Test1Func6 is a paid mutator transaction binding the contract method 0xcbb835f6.
//
// Solidity: function test1Func6() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func6(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func6")
}

// Test1Func6 is a paid mutator transaction binding the contract method 0xcbb835f6.
//
// Solidity: function test1Func6() returns()
func (_Test1Facet *Test1FacetSession) Test1Func6() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func6(&_Test1Facet.TransactOpts)
}

// Test1Func6 is a paid mutator transaction binding the contract method 0xcbb835f6.
//
// Solidity: function test1Func6() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func6() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func6(&_Test1Facet.TransactOpts)
}

// Test1Func7 is a paid mutator transaction binding the contract method 0x71a99d6f.
//
// Solidity: function test1Func7() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func7(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func7")
}

// Test1Func7 is a paid mutator transaction binding the contract method 0x71a99d6f.
//
// Solidity: function test1Func7() returns()
func (_Test1Facet *Test1FacetSession) Test1Func7() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func7(&_Test1Facet.TransactOpts)
}

// Test1Func7 is a paid mutator transaction binding the contract method 0x71a99d6f.
//
// Solidity: function test1Func7() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func7() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func7(&_Test1Facet.TransactOpts)
}

// Test1Func8 is a paid mutator transaction binding the contract method 0xdb32da15.
//
// Solidity: function test1Func8() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func8(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func8")
}

// Test1Func8 is a paid mutator transaction binding the contract method 0xdb32da15.
//
// Solidity: function test1Func8() returns()
func (_Test1Facet *Test1FacetSession) Test1Func8() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func8(&_Test1Facet.TransactOpts)
}

// Test1Func8 is a paid mutator transaction binding the contract method 0xdb32da15.
//
// Solidity: function test1Func8() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func8() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func8(&_Test1Facet.TransactOpts)
}

// Test1Func9 is a paid mutator transaction binding the contract method 0xcd0bae09.
//
// Solidity: function test1Func9() returns()
func (_Test1Facet *Test1FacetTransactor) Test1Func9(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "test1Func9")
}

// Test1Func9 is a paid mutator transaction binding the contract method 0xcd0bae09.
//
// Solidity: function test1Func9() returns()
func (_Test1Facet *Test1FacetSession) Test1Func9() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func9(&_Test1Facet.TransactOpts)
}

// Test1Func9 is a paid mutator transaction binding the contract method 0xcd0bae09.
//
// Solidity: function test1Func9() returns()
func (_Test1Facet *Test1FacetTransactorSession) Test1Func9() (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1Func9(&_Test1Facet.TransactOpts)
}

// Test1FacetTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the Test1Facet contract.
type Test1FacetTestEventIterator struct {
	Event *Test1FacetTestEvent // Event containing the contract specifics and raw log

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
func (it *Test1FacetTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Test1FacetTestEvent)
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
		it.Event = new(Test1FacetTestEvent)
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
func (it *Test1FacetTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Test1FacetTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Test1FacetTestEvent represents a TestEvent event raised by the Test1Facet contract.
type Test1FacetTestEvent struct {
	Something common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0xab77f9000c19702a713e62164a239e3764dde2ba5265c7551f9a49e0d304530d.
//
// Solidity: event TestEvent(address something)
func (_Test1Facet *Test1FacetFilterer) FilterTestEvent(opts *bind.FilterOpts) (*Test1FacetTestEventIterator, error) {

	logs, sub, err := _Test1Facet.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &Test1FacetTestEventIterator{contract: _Test1Facet.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0xab77f9000c19702a713e62164a239e3764dde2ba5265c7551f9a49e0d304530d.
//
// Solidity: event TestEvent(address something)
func (_Test1Facet *Test1FacetFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *Test1FacetTestEvent) (event.Subscription, error) {

	logs, sub, err := _Test1Facet.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Test1FacetTestEvent)
				if err := _Test1Facet.contract.UnpackLog(event, "TestEvent", log); err != nil {
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

// ParseTestEvent is a log parse operation binding the contract event 0xab77f9000c19702a713e62164a239e3764dde2ba5265c7551f9a49e0d304530d.
//
// Solidity: event TestEvent(address something)
func (_Test1Facet *Test1FacetFilterer) ParseTestEvent(log types.Log) (*Test1FacetTestEvent, error) {
	event := new(Test1FacetTestEvent)
	if err := _Test1Facet.contract.UnpackLog(event, "TestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
