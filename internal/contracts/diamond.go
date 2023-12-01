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

// IDiamondCutFacetFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondMetaData contains all meta data concerning the Diamond contract.
var DiamondMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_contractOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_diamondCutFacet\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DiamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDiamondCutFacet.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamondCutFacet.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InitializationFunctionReverted\",\"inputs\":[{\"name\":\"_initializationContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]",
	Bin: "0x60806040526040516113b23803806113b283398101604081905261002291610f61565b61002b82610136565b604080516001808252818301909252600091816020015b604080516060808201835260008083526020830152918101919091528152602001906001900390816100425750506040805160018082528183019092529192506000919060208083019080368337019050509050631f931c1c60e01b816000815181106100b1576100b1610f94565b6001600160e01b031990921660209283029190910182015260408051606081019091526001600160a01b038516815290810160008152602001828152508260008151811061010157610101610f94565b602002602001018190525061012d826000604051806020016040528060008152506101b960201b60201c565b505050506111fe565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c132080546001600160a01b031981166001600160a01b03848116918217909355604051600080516020611306833981519152939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60005b83518110156103965760008482815181106101d9576101d9610f94565b6020026020010151602001519050600060028111156101fa576101fa610faa565b81600281111561020c5761020c610faa565b036102605761025b85838151811061022657610226610f94565b60200260200101516000015186848151811061024457610244610f94565b6020026020010151604001516103e160201b60201c565b610383565b600181600281111561027457610274610faa565b036102c35761025b85838151811061028e5761028e610f94565b6020026020010151600001518684815181106102ac576102ac610f94565b60200260200101516040015161065960201b60201c565b60028160028111156102d7576102d7610faa565b036103265761025b8583815181106102f1576102f1610f94565b60200260200101516000015186848151811061030f5761030f610f94565b6020026020010151604001516108da60201b60201c565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b60648201526084015b60405180910390fd5b508061038e81610fd6565b9150506101bc565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb6738383836040516103ca9392919061103f565b60405180910390a16103dc8282610a31565b505050565b60008151116104345760405162461bcd60e51b815260206004820152602b602482015260008051602061139283398151915260448201526a1858d95d081d1bc818dd5d60aa1b606482015260840161037a565b6000805160206113068339815191526001600160a01b03831661049c5760405162461bcd60e51b815260206004820152602c602482015260008051602061134e83398151915260448201526b65206164647265737328302960a01b606482015260840161037a565b6001600160a01b0383166000908152600182016020526040812054906001600160601b03821690036104d2576104d28285610af7565b60005b83518110156106525760008482815181106104f2576104f2610f94565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156105985760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f60448201527f6e207468617420616c7265616479206578697374730000000000000000000000606482015260840161037a565b6001600160e01b0319821660008181526020878152604080832080546001600160a01b03908116600160a01b6001600160601b038c16021782558c168085526001808c0185529285208054938401815585528385206008840401805463ffffffff60079095166004026101000a948502191660e08a901c94909402939093179092559390925287905281546001600160a01b0319161790558361063a8161113f565b9450505050808061064a90610fd6565b9150506104d5565b5050505050565b60008151116106ac5760405162461bcd60e51b815260206004820152602b602482015260008051602061139283398151915260448201526a1858d95d081d1bc818dd5d60aa1b606482015260840161037a565b6000805160206113068339815191526001600160a01b0383166107145760405162461bcd60e51b815260206004820152602c602482015260008051602061134e83398151915260448201526b65206164647265737328302960a01b606482015260840161037a565b6001600160a01b0383166000908152600182016020526040812054906001600160601b038216900361074a5761074a8285610af7565b60005b835181101561065257600084828151811061076a5761076a610f94565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b0390811690871681036108155760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e0000000000000000606482015260840161037a565b610820858284610b61565b6001600160e01b0319821660008181526020878152604080832080546001600160a01b03908116600160a01b6001600160601b038c16021782558c168085526001808c0185529285208054938401815585528385206008840401805463ffffffff60079095166004026101000a948502191660e08a901c94909402939093179092559390925287905281546001600160a01b031916179055836108c28161113f565b945050505080806108d290610fd6565b91505061074d565b600081511161092d5760405162461bcd60e51b815260206004820152602b602482015260008051602061139283398151915260448201526a1858d95d081d1bc818dd5d60aa1b606482015260840161037a565b6000805160206113068339815191526001600160a01b038316156109b95760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f7665206661636574206164647260448201527f657373206d757374206265206164647265737328302900000000000000000000606482015260840161037a565b60005b8251811015610a2b5760008382815181106109d9576109d9610f94565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b0316610a16848284610b61565b50508080610a2390610fd6565b9150506109bc565b50505050565b6001600160a01b038216610a43575050565b610a658260405180606001604052806028815260200161132660289139610f24565b600080836001600160a01b031683604051610a80919061116d565b600060405180830381855af49150503d8060008114610abb576040519150601f19603f3d011682016040523d82523d6000602084013e610ac0565b606091505b509150915081610a2b57805115610ada5780518082602001fd5b838360405163192105d760e01b815260040161037a929190611189565b610b198160405180606001604052806024815260200161136e60249139610f24565b6002820180546001600160a01b0390921660008181526001948501602090815260408220860185905594840183559182529290200180546001600160a01b0319169091179055565b6001600160a01b038216610bdd5760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e2774206578697374000000000000000000606482015260840161037a565b306001600160a01b03831603610c4c5760405162461bcd60e51b815260206004820152602e60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526d3a30b1363290333ab731ba34b7b760911b606482015260840161037a565b6001600160e01b03198116600090815260208481526040808320546001600160a01b0386168452600180880190935290832054600160a01b9091046001600160601b03169291610c9b916111b5565b9050808214610d8d576001600160a01b03841660009081526001860160205260408120805483908110610cd057610cd0610f94565b600091825260208083206008830401546001600160a01b038916845260018a019091526040909220805460079092166004026101000a90920460e01b925082919085908110610d2157610d21610f94565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c929092029390931790556001600160e01b03199290921682528690526040902080546001600160a01b0316600160a01b6001600160601b038516021790555b6001600160a01b03841660009081526001860160205260409020805480610db657610db66111ce565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b03198516825286905260408120819055819003610652576002850154600090610e19906001906111b5565b6001600160a01b0386166000908152600180890160205260409091200154909150808214610ec8576000876002018381548110610e5857610e58610f94565b6000918252602090912001546002890180546001600160a01b039092169250829184908110610e8957610e89610f94565b600091825260208083209190910180546001600160a01b0319166001600160a01b03948516179055929091168152600189810190925260409020018190555b86600201805480610edb57610edb6111ce565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0388168252600189810190915260408220015550505050505050565b813b8181610a2b5760405162461bcd60e51b815260040161037a91906111e4565b80516001600160a01b0381168114610f5c57600080fd5b919050565b60008060408385031215610f7457600080fd5b610f7d83610f45565b9150610f8b60208401610f45565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610fe857610fe8610fc0565b5060010190565b60005b8381101561100a578181015183820152602001610ff2565b50506000910152565b6000815180845261102b816020860160208601610fef565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b8481101561110f57898403607f19018652815180516001600160a01b031685528381015189860190600381106110ae57634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b808310156110fa5783516001600160e01b03191682529286019260019290920191908601906110d0565b50978501979550505090820190600101611068565b50506001600160a01b038a169088015286810360408801526111318189611013565b9a9950505050505050505050565b60006001600160601b038281166002600160601b0319810161116357611163610fc0565b6001019392505050565b6000825161117f818460208701610fef565b9190910192915050565b6001600160a01b03831681526040602082018190526000906111ad90830184611013565b949350505050565b818103818111156111c8576111c8610fc0565b92915050565b634e487b7160e01b600052603160045260246000fd5b6020815260006111f76020830184611013565b9392505050565b60fa8061120c6000396000f3fe608060405236600a57005b600080356001600160e01b03191681527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c602081905260409091205481906001600160a01b03168060a15760405162461bcd60e51b815260206004820181905260248201527f4469616d6f6e643a2046756e6374696f6e20646f6573206e6f74206578697374604482015260640160405180910390fd5b3660008037600080366000845af43d6000803e80801560bf573d6000f35b3d6000fdfea26469706673582212202b0217970ffee65bdba80c7aaf9be77e145fdfb76c5649aed40caa1b96456b7a64736f6c63430008150033c8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a204164642066616365742063616e277420624c69624469616d6f6e644375743a204e657720666163657420686173206e6f20636f64654c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e2066",
}

// DiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondMetaData.ABI instead.
var DiamondABI = DiamondMetaData.ABI

// DiamondBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondMetaData.Bin instead.
var DiamondBin = DiamondMetaData.Bin

// DeployDiamond deploys a new Ethereum contract, binding an instance of Diamond to it.
func DeployDiamond(auth *bind.TransactOpts, backend bind.ContractBackend, _contractOwner common.Address, _diamondCutFacet common.Address) (common.Address, *types.Transaction, *Diamond, error) {
	parsed, err := DiamondMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondBin), backend, _contractOwner, _diamondCutFacet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// Diamond is an auto generated Go binding around an Ethereum contract.
type Diamond struct {
	DiamondCaller     // Read-only binding to the contract
	DiamondTransactor // Write-only binding to the contract
	DiamondFilterer   // Log filterer for contract events
}

// DiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondSession struct {
	Contract     *Diamond          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCallerSession struct {
	Contract *DiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondTransactorSession struct {
	Contract     *DiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondRaw struct {
	Contract *Diamond // Generic contract binding to access the raw methods on
}

// DiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCallerRaw struct {
	Contract *DiamondCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondTransactorRaw struct {
	Contract *DiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamond creates a new instance of Diamond, bound to a specific deployed contract.
func NewDiamond(address common.Address, backend bind.ContractBackend) (*Diamond, error) {
	contract, err := bindDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// NewDiamondCaller creates a new read-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondCaller(address common.Address, caller bind.ContractCaller) (*DiamondCaller, error) {
	contract, err := bindDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCaller{contract: contract}, nil
}

// NewDiamondTransactor creates a new write-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondTransactor, error) {
	contract, err := bindDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondTransactor{contract: contract}, nil
}

// NewDiamondFilterer creates a new log filterer instance of Diamond, bound to a specific deployed contract.
func NewDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondFilterer, error) {
	contract, err := bindDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondFilterer{contract: contract}, nil
}

// bindDiamond binds a generic wrapper to an already deployed contract.
func bindDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.DiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondSession) Receive() (*types.Transaction, error) {
	return _Diamond.Contract.Receive(&_Diamond.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondTransactorSession) Receive() (*types.Transaction, error) {
	return _Diamond.Contract.Receive(&_Diamond.TransactOpts)
}

// DiamondDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the Diamond contract.
type DiamondDiamondCutIterator struct {
	Event *DiamondDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondDiamondCut)
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
		it.Event = new(DiamondDiamondCut)
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
func (it *DiamondDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondDiamondCut represents a DiamondCut event raised by the Diamond contract.
type DiamondDiamondCut struct {
	DiamondCut []IDiamondCutFacetFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Diamond *DiamondFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondDiamondCutIterator, error) {

	logs, sub, err := _Diamond.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondDiamondCutIterator{contract: _Diamond.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Diamond *DiamondFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondDiamondCut) (event.Subscription, error) {

	logs, sub, err := _Diamond.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondDiamondCut)
				if err := _Diamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Diamond *DiamondFilterer) ParseDiamondCut(log types.Log) (*DiamondDiamondCut, error) {
	event := new(DiamondDiamondCut)
	if err := _Diamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Diamond contract.
type DiamondOwnershipTransferredIterator struct {
	Event *DiamondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DiamondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondOwnershipTransferred)
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
		it.Event = new(DiamondOwnershipTransferred)
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
func (it *DiamondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondOwnershipTransferred represents a OwnershipTransferred event raised by the Diamond contract.
type DiamondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Diamond *DiamondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DiamondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Diamond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnershipTransferredIterator{contract: _Diamond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Diamond *DiamondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DiamondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Diamond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondOwnershipTransferred)
				if err := _Diamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Diamond *DiamondFilterer) ParseOwnershipTransferred(log types.Log) (*DiamondOwnershipTransferred, error) {
	event := new(DiamondOwnershipTransferred)
	if err := _Diamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
