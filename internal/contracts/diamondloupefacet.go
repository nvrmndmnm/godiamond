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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// DiamondLoupeFacetMetaData contains all meta data concerning the DiamondLoupeFacet contract.
var DiamondLoupeFacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061068e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806301ffc9a71461005c57806352ef6b2c146100bd5780637a0ed627146100d2578063adfca15e146100e7578063cdffacc614610107575b600080fd5b6100a861006a366004610469565b6001600160e01b03191660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f602052604090205460ff1690565b60405190151581526020015b60405180910390f35b6100c561015f565b6040516100b4919061049a565b6100da6101d2565b6040516100b4919061052c565b6100fa6100f53660046105a9565b61039d565b6040516100b491906105d2565b610147610115366004610469565b6001600160e01b031916600090815260008051602061063983398151915260205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016100b4565b60606000600080516020610639833981519152600281018054604080516020808402820181019092528281529394508301828280156101c757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116101a9575b505050505091505090565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e54606090600080516020610639833981519152908067ffffffffffffffff811115610220576102206105e5565b60405190808252806020026020018201604052801561026657816020015b60408051808201909152600081526060602082015281526020019060019003908161023e5790505b50925060005b8181101561039757600083600201828154811061028b5761028b6105fb565b9060005260206000200160009054906101000a90046001600160a01b03169050808583815181106102be576102be6105fb565b6020908102919091018101516001600160a01b03928316905290821660009081526001860182526040908190208054825181850281018501909352808352919290919083018282801561035d57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161031f5790505b5050505050858381518110610374576103746105fb565b60200260200101516020018190525050808061038f90610611565b91505061026c565b50505090565b6001600160a01b03811660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60209081526040918290208054835181840281018401909452808452606093600080516020610639833981519152939092919083018282801561045c57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161041e5790505b5050505050915050919050565b60006020828403121561047b57600080fd5b81356001600160e01b03198116811461049357600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156104db5783516001600160a01b0316835292840192918401916001016104b6565b50909695505050505050565b600081518084526020808501945080840160005b838110156105215781516001600160e01b031916875295820195908201906001016104fb565b509495945050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561059b57888303603f19018552815180516001600160a01b03168452870151878401879052610588878501826104e7565b9588019593505090860190600101610553565b509098975050505050505050565b6000602082840312156105bb57600080fd5b81356001600160a01b038116811461049357600080fd5b60208152600061049360208301846104e7565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006001820161063157634e487b7160e01b600052601160045260246000fd5b506001019056fec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131ca26469706673582212205e7a764d33818564c556958bd284931894c9923adf48c7d773b501da2ac00e4364736f6c63430008110033",
}

// DiamondLoupeFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondLoupeFacetMetaData.ABI instead.
var DiamondLoupeFacetABI = DiamondLoupeFacetMetaData.ABI

// DiamondLoupeFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondLoupeFacetMetaData.Bin instead.
var DiamondLoupeFacetBin = DiamondLoupeFacetMetaData.Bin

// DeployDiamondLoupeFacet deploys a new Ethereum contract, binding an instance of DiamondLoupeFacet to it.
func DeployDiamondLoupeFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondLoupeFacet, error) {
	parsed, err := DiamondLoupeFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondLoupeFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// DiamondLoupeFacet is an auto generated Go binding around an Ethereum contract.
type DiamondLoupeFacet struct {
	DiamondLoupeFacetCaller     // Read-only binding to the contract
	DiamondLoupeFacetTransactor // Write-only binding to the contract
	DiamondLoupeFacetFilterer   // Log filterer for contract events
}

// DiamondLoupeFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondLoupeFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondLoupeFacetSession struct {
	Contract     *DiamondLoupeFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondLoupeFacetCallerSession struct {
	Contract *DiamondLoupeFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DiamondLoupeFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondLoupeFacetTransactorSession struct {
	Contract     *DiamondLoupeFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondLoupeFacetRaw struct {
	Contract *DiamondLoupeFacet // Generic contract binding to access the raw methods on
}

// DiamondLoupeFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCallerRaw struct {
	Contract *DiamondLoupeFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondLoupeFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactorRaw struct {
	Contract *DiamondLoupeFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondLoupeFacet creates a new instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacet(address common.Address, backend bind.ContractBackend) (*DiamondLoupeFacet, error) {
	contract, err := bindDiamondLoupeFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// NewDiamondLoupeFacetCaller creates a new read-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondLoupeFacetCaller, error) {
	contract, err := bindDiamondLoupeFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetCaller{contract: contract}, nil
}

// NewDiamondLoupeFacetTransactor creates a new write-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondLoupeFacetTransactor, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetTransactor{contract: contract}, nil
}

// NewDiamondLoupeFacetFilterer creates a new log filterer instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondLoupeFacetFilterer, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetFilterer{contract: contract}, nil
}

// bindDiamondLoupeFacet binds a generic wrapper to an already deployed contract.
func bindDiamondLoupeFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondLoupeFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}
