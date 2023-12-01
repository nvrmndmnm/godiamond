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

// OwnershipFacetMetaData contains all meta data concerning the OwnershipFacet contract.
var OwnershipFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50610248806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610074565b6040516001600160a01b03909116815260200160405180910390f35b61007261006d3660046101e2565b6100ac565b005b60006100a77fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c1320546001600160a01b031690565b905090565b6100b46100c0565b6100bd8161014d565b50565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c600401546001600160a01b0316331461014b5760405162461bcd60e51b815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201526132b960f11b606482015260840160405180910390fd5b565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c132080546001600160a01b031981166001600160a01b038481169182179093556040517fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6000602082840312156101f457600080fd5b81356001600160a01b038116811461020b57600080fd5b939250505056fea2646970667358221220cb3178293de4aa15462dd3799ca53e55faadfd5cbc6acc3670a502739d4a95d264736f6c63430008150033",
}

// OwnershipFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnershipFacetMetaData.ABI instead.
var OwnershipFacetABI = OwnershipFacetMetaData.ABI

// OwnershipFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnershipFacetMetaData.Bin instead.
var OwnershipFacetBin = OwnershipFacetMetaData.Bin

// DeployOwnershipFacet deploys a new Ethereum contract, binding an instance of OwnershipFacet to it.
func DeployOwnershipFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnershipFacet, error) {
	parsed, err := OwnershipFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnershipFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnershipFacet{OwnershipFacetCaller: OwnershipFacetCaller{contract: contract}, OwnershipFacetTransactor: OwnershipFacetTransactor{contract: contract}, OwnershipFacetFilterer: OwnershipFacetFilterer{contract: contract}}, nil
}

// OwnershipFacet is an auto generated Go binding around an Ethereum contract.
type OwnershipFacet struct {
	OwnershipFacetCaller     // Read-only binding to the contract
	OwnershipFacetTransactor // Write-only binding to the contract
	OwnershipFacetFilterer   // Log filterer for contract events
}

// OwnershipFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnershipFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnershipFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnershipFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnershipFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnershipFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnershipFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnershipFacetSession struct {
	Contract     *OwnershipFacet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnershipFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnershipFacetCallerSession struct {
	Contract *OwnershipFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OwnershipFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnershipFacetTransactorSession struct {
	Contract     *OwnershipFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OwnershipFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnershipFacetRaw struct {
	Contract *OwnershipFacet // Generic contract binding to access the raw methods on
}

// OwnershipFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnershipFacetCallerRaw struct {
	Contract *OwnershipFacetCaller // Generic read-only contract binding to access the raw methods on
}

// OwnershipFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnershipFacetTransactorRaw struct {
	Contract *OwnershipFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnershipFacet creates a new instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacet(address common.Address, backend bind.ContractBackend) (*OwnershipFacet, error) {
	contract, err := bindOwnershipFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacet{OwnershipFacetCaller: OwnershipFacetCaller{contract: contract}, OwnershipFacetTransactor: OwnershipFacetTransactor{contract: contract}, OwnershipFacetFilterer: OwnershipFacetFilterer{contract: contract}}, nil
}

// NewOwnershipFacetCaller creates a new read-only instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacetCaller(address common.Address, caller bind.ContractCaller) (*OwnershipFacetCaller, error) {
	contract, err := bindOwnershipFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetCaller{contract: contract}, nil
}

// NewOwnershipFacetTransactor creates a new write-only instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnershipFacetTransactor, error) {
	contract, err := bindOwnershipFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetTransactor{contract: contract}, nil
}

// NewOwnershipFacetFilterer creates a new log filterer instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnershipFacetFilterer, error) {
	contract, err := bindOwnershipFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetFilterer{contract: contract}, nil
}

// bindOwnershipFacet binds a generic wrapper to an already deployed contract.
func bindOwnershipFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnershipFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnershipFacet *OwnershipFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnershipFacet.Contract.OwnershipFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnershipFacet *OwnershipFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.OwnershipFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnershipFacet *OwnershipFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.OwnershipFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnershipFacet *OwnershipFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnershipFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnershipFacet *OwnershipFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnershipFacet *OwnershipFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_OwnershipFacet *OwnershipFacetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnershipFacet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_OwnershipFacet *OwnershipFacetSession) Owner() (common.Address, error) {
	return _OwnershipFacet.Contract.Owner(&_OwnershipFacet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_OwnershipFacet *OwnershipFacetCallerSession) Owner() (common.Address, error) {
	return _OwnershipFacet.Contract.Owner(&_OwnershipFacet.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_OwnershipFacet *OwnershipFacetTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _OwnershipFacet.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_OwnershipFacet *OwnershipFacetSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.TransferOwnership(&_OwnershipFacet.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_OwnershipFacet *OwnershipFacetTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.TransferOwnership(&_OwnershipFacet.TransactOpts, _newOwner)
}

// OwnershipFacetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnershipFacet contract.
type OwnershipFacetOwnershipTransferredIterator struct {
	Event *OwnershipFacetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnershipFacetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnershipFacetOwnershipTransferred)
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
		it.Event = new(OwnershipFacetOwnershipTransferred)
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
func (it *OwnershipFacetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnershipFacetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnershipFacetOwnershipTransferred represents a OwnershipTransferred event raised by the OwnershipFacet contract.
type OwnershipFacetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnershipFacetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnershipFacet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetOwnershipTransferredIterator{contract: _OwnershipFacet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnershipFacetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnershipFacet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnershipFacetOwnershipTransferred)
				if err := _OwnershipFacet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OwnershipFacet *OwnershipFacetFilterer) ParseOwnershipTransferred(log types.Log) (*OwnershipFacetOwnershipTransferred, error) {
	event := new(OwnershipFacetOwnershipTransferred)
	if err := _OwnershipFacet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnershipFacetOwnershipTransferred0Iterator is returned from FilterOwnershipTransferred0 and is used to iterate over the raw logs and unpacked data for OwnershipTransferred0 events raised by the OwnershipFacet contract.
type OwnershipFacetOwnershipTransferred0Iterator struct {
	Event *OwnershipFacetOwnershipTransferred0 // Event containing the contract specifics and raw log

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
func (it *OwnershipFacetOwnershipTransferred0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnershipFacetOwnershipTransferred0)
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
		it.Event = new(OwnershipFacetOwnershipTransferred0)
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
func (it *OwnershipFacetOwnershipTransferred0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnershipFacetOwnershipTransferred0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnershipFacetOwnershipTransferred0 represents a OwnershipTransferred0 event raised by the OwnershipFacet contract.
type OwnershipFacetOwnershipTransferred0 struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred0 is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) FilterOwnershipTransferred0(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnershipFacetOwnershipTransferred0Iterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnershipFacet.contract.FilterLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetOwnershipTransferred0Iterator{contract: _OwnershipFacet.contract, event: "OwnershipTransferred0", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred0 is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) WatchOwnershipTransferred0(opts *bind.WatchOpts, sink chan<- *OwnershipFacetOwnershipTransferred0, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnershipFacet.contract.WatchLogs(opts, "OwnershipTransferred0", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnershipFacetOwnershipTransferred0)
				if err := _OwnershipFacet.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
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

// ParseOwnershipTransferred0 is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) ParseOwnershipTransferred0(log types.Log) (*OwnershipFacetOwnershipTransferred0, error) {
	event := new(OwnershipFacetOwnershipTransferred0)
	if err := _OwnershipFacet.contract.UnpackLog(event, "OwnershipTransferred0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
