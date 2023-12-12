// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package facets

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

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondCutFacetMetaData contains all meta data concerning the DiamondCutFacet contract.
var DiamondCutFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"diamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"internalType\":\"structIDiamondCut.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamondCut.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DiamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamondCut.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DiamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamondCut.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InitializationFunctionReverted\",\"inputs\":[{\"name\":\"_initializationContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061138f806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631f931c1c14610030575b600080fd5b61004361003e366004610d6a565b610045565b005b61004d61009e565b61009761005a8587610eb0565b8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061011a92505050565b5050505050565b6000805160206112ee833981519152600401546001600160a01b031633146101185760405162461bcd60e51b815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201526132b960f11b60648201526084015b60405180910390fd5b565b60005b83518110156102e057600084828151811061013a5761013a610ff4565b60200260200101516020015190506000600281111561015b5761015b61100a565b81600281111561016d5761016d61100a565b036101bb576101b685838151811061018757610187610ff4565b6020026020010151600001518684815181106101a5576101a5610ff4565b60200260200101516040015161032b565b6102cd565b60018160028111156101cf576101cf61100a565b03610218576101b68583815181106101e9576101e9610ff4565b60200260200101516000015186848151811061020757610207610ff4565b6020026020010151604001516104a3565b600281600281111561022c5761022c61100a565b03610275576101b685838151811061024657610246610ff4565b60200260200101516000015186848151811061026457610264610ff4565b602002602001015160400151610633565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b606482015260840161010f565b50806102d881611036565b91505061011d565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb6738383836040516103149392919061109f565b60405180910390a16103268282610751565b505050565b600081511161034c5760405162461bcd60e51b815260040161010f9061119f565b6000805160206112ee8339815191526001600160a01b0383166103815760405162461bcd60e51b815260040161010f906111ea565b6001600160a01b0383166000908152600182016020526040812054906001600160601b03821690036103b7576103b78285610817565b60005b83518110156100975760008482815181106103d7576103d7610ff4565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156104755760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f6044820152746e207468617420616c72656164792065786973747360581b606482015260840161010f565b6104818583868a610881565b8361048b81611236565b9450505050808061049b90611036565b9150506103ba565b60008151116104c45760405162461bcd60e51b815260040161010f9061119f565b6000805160206112ee8339815191526001600160a01b0383166104f95760405162461bcd60e51b815260040161010f906111ea565b6001600160a01b0383166000908152600182016020526040812054906001600160601b038216900361052f5761052f8285610817565b60005b835181101561009757600084828151811061054f5761054f610ff4565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b0390811690871681036105fa5760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e0000000000000000606482015260840161010f565b610605858284610921565b6106118583868a610881565b8361061b81611236565b9450505050808061062b90611036565b915050610532565b60008151116106545760405162461bcd60e51b815260040161010f9061119f565b6000805160206112ee8339815191526001600160a01b038316156106d95760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f76652066616365742061646472604482015275657373206d757374206265206164647265737328302960501b606482015260840161010f565b60005b825181101561074b5760008382815181106106f9576106f9610ff4565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b0316610736848284610921565b5050808061074390611036565b9150506106dc565b50505050565b6001600160a01b038216610763575050565b6107858260405180606001604052806028815260200161130e60289139610ce4565b600080836001600160a01b0316836040516107a0919061125c565b600060405180830381855af49150503d80600081146107db576040519150601f19603f3d011682016040523d82523d6000602084013e6107e0565b606091505b50915091508161074b578051156107fa5780518082602001fd5b838360405163192105d760e01b815260040161010f929190611278565b6108398160405180606001604052806024815260200161133660249139610ce4565b6002820180546001600160a01b0390921660008181526001948501602090815260408220860185905594840183559182529290200180546001600160a01b0319169091179055565b6001600160e01b0319831660008181526020868152604080832080546001600160601b03909716600160a01b026001600160a01b0397881617815594909516808352600180890183529583208054968701815583528183206008870401805460e09890981c60046007909816979097026101000a96870263ffffffff9097021990971695909517909555529290915281546001600160a01b031916179055565b6001600160a01b03821661099d5760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e2774206578697374000000000000000000606482015260840161010f565b306001600160a01b03831603610a0c5760405162461bcd60e51b815260206004820152602e60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526d3a30b1363290333ab731ba34b7b760911b606482015260840161010f565b6001600160e01b03198116600090815260208481526040808320546001600160a01b0386168452600180880190935290832054600160a01b9091046001600160601b03169291610a5b916112a4565b9050808214610b4d576001600160a01b03841660009081526001860160205260408120805483908110610a9057610a90610ff4565b600091825260208083206008830401546001600160a01b038916845260018a019091526040909220805460079092166004026101000a90920460e01b925082919085908110610ae157610ae1610ff4565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c929092029390931790556001600160e01b03199290921682528690526040902080546001600160a01b0316600160a01b6001600160601b038516021790555b6001600160a01b03841660009081526001860160205260409020805480610b7657610b766112bd565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b03198516825286905260408120819055819003610097576002850154600090610bd9906001906112a4565b6001600160a01b0386166000908152600180890160205260409091200154909150808214610c88576000876002018381548110610c1857610c18610ff4565b6000918252602090912001546002890180546001600160a01b039092169250829184908110610c4957610c49610ff4565b600091825260208083209190910180546001600160a01b0319166001600160a01b03948516179055929091168152600189810190925260409020018190555b86600201805480610c9b57610c9b6112bd565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0388168252600189810190915260408220015550505050505050565b813b818161074b5760405162461bcd60e51b815260040161010f91906112d3565b80356001600160a01b0381168114610d1c57600080fd5b919050565b60008083601f840112610d3357600080fd5b50813567ffffffffffffffff811115610d4b57600080fd5b602083019150836020828501011115610d6357600080fd5b9250929050565b600080600080600060608688031215610d8257600080fd5b853567ffffffffffffffff80821115610d9a57600080fd5b818801915088601f830112610dae57600080fd5b813581811115610dbd57600080fd5b8960208260051b8501011115610dd257600080fd5b60208301975080965050610de860208901610d05565b94506040880135915080821115610dfe57600080fd5b50610e0b88828901610d21565b969995985093965092949392505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610e5557610e55610e1c565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e8457610e84610e1c565b604052919050565b600067ffffffffffffffff821115610ea657610ea6610e1c565b5060051b60200190565b6000610ec3610ebe84610e8c565b610e5b565b83815260208082019190600586811b860136811115610ee157600080fd5b865b81811015610fe757803567ffffffffffffffff80821115610f045760008081fd5b818a01915060608236031215610f1a5760008081fd5b610f22610e32565b610f2b83610d05565b81528683013560038110610f3f5760008081fd5b8188015260408381013583811115610f575760008081fd5b939093019236601f850112610f6e57600092508283fd5b83359250610f7e610ebe84610e8c565b83815292871b84018801928881019036851115610f9b5760008081fd5b948901945b84861015610fd05785356001600160e01b031981168114610fc15760008081fd5b82529489019490890190610fa0565b918301919091525088525050948301948301610ee3565b5092979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161104857611048611020565b5060010190565b60005b8381101561106a578181015183820152602001611052565b50506000910152565b6000815180845261108b81602086016020860161104f565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b8481101561116f57898403607f19018652815180516001600160a01b0316855283810151898601906003811061110e57634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b8083101561115a5783516001600160e01b0319168252928601926001929092019190860190611130565b509785019795505050908201906001016110c8565b50506001600160a01b038a169088015286810360408801526111918189611073565b9a9950505050505050505050565b6020808252602b908201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660408201526a1858d95d081d1bc818dd5d60aa1b606082015260800190565b6020808252602c908201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260408201526b65206164647265737328302960a01b606082015260800190565b60006001600160601b0380831681810361125257611252611020565b6001019392505050565b6000825161126e81846020870161104f565b9190910192915050565b6001600160a01b038316815260406020820181905260009061129c90830184611073565b949350505050565b818103818111156112b7576112b7611020565b92915050565b634e487b7160e01b600052603160045260246000fd5b6020815260006112e66020830184611073565b939250505056fec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a204e657720666163657420686173206e6f20636f6465a2646970667358221220639ac78ee32fb154bc61af3577e36b8409ddad97a5fe67bf448538881f0eb0b564736f6c63430008150033",
}

// DiamondCutFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondCutFacetMetaData.ABI instead.
var DiamondCutFacetABI = DiamondCutFacetMetaData.ABI

// DiamondCutFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondCutFacetMetaData.Bin instead.
var DiamondCutFacetBin = DiamondCutFacetMetaData.Bin

// DeployDiamondCutFacet deploys a new Ethereum contract, binding an instance of DiamondCutFacet to it.
func DeployDiamondCutFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondCutFacet, error) {
	parsed, err := DiamondCutFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondCutFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// DiamondCutFacet is an auto generated Go binding around an Ethereum contract.
type DiamondCutFacet struct {
	DiamondCutFacetCaller     // Read-only binding to the contract
	DiamondCutFacetTransactor // Write-only binding to the contract
	DiamondCutFacetFilterer   // Log filterer for contract events
}

// DiamondCutFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCutFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondCutFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondCutFacetSession struct {
	Contract     *DiamondCutFacet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCutFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCutFacetCallerSession struct {
	Contract *DiamondCutFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DiamondCutFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondCutFacetTransactorSession struct {
	Contract     *DiamondCutFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DiamondCutFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondCutFacetRaw struct {
	Contract *DiamondCutFacet // Generic contract binding to access the raw methods on
}

// DiamondCutFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCutFacetCallerRaw struct {
	Contract *DiamondCutFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondCutFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactorRaw struct {
	Contract *DiamondCutFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondCutFacet creates a new instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacet(address common.Address, backend bind.ContractBackend) (*DiamondCutFacet, error) {
	contract, err := bindDiamondCutFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// NewDiamondCutFacetCaller creates a new read-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondCutFacetCaller, error) {
	contract, err := bindDiamondCutFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetCaller{contract: contract}, nil
}

// NewDiamondCutFacetTransactor creates a new write-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondCutFacetTransactor, error) {
	contract, err := bindDiamondCutFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetTransactor{contract: contract}, nil
}

// NewDiamondCutFacetFilterer creates a new log filterer instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondCutFacetFilterer, error) {
	contract, err := bindDiamondCutFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetFilterer{contract: contract}, nil
}

// bindDiamondCutFacet binds a generic wrapper to an already deployed contract.
func bindDiamondCutFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondCutFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.DiamondCutFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCutFacetDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCutIterator struct {
	Event *DiamondCutFacetDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondCutFacetDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutFacetDiamondCut)
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
		it.Event = new(DiamondCutFacetDiamondCut)
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
func (it *DiamondCutFacetDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutFacetDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutFacetDiamondCut represents a DiamondCut event raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondCutFacetDiamondCutIterator, error) {

	logs, sub, err := _DiamondCutFacet.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetDiamondCutIterator{contract: _DiamondCutFacet.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondCutFacetDiamondCut) (event.Subscription, error) {

	logs, sub, err := _DiamondCutFacet.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutFacetDiamondCut)
				if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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
func (_DiamondCutFacet *DiamondCutFacetFilterer) ParseDiamondCut(log types.Log) (*DiamondCutFacetDiamondCut, error) {
	event := new(DiamondCutFacetDiamondCut)
	if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondCutFacetDiamondCut0Iterator is returned from FilterDiamondCut0 and is used to iterate over the raw logs and unpacked data for DiamondCut0 events raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCut0Iterator struct {
	Event *DiamondCutFacetDiamondCut0 // Event containing the contract specifics and raw log

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
func (it *DiamondCutFacetDiamondCut0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutFacetDiamondCut0)
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
		it.Event = new(DiamondCutFacetDiamondCut0)
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
func (it *DiamondCutFacetDiamondCut0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutFacetDiamondCut0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutFacetDiamondCut0 represents a DiamondCut0 event raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCut0 struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut0 is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) FilterDiamondCut0(opts *bind.FilterOpts) (*DiamondCutFacetDiamondCut0Iterator, error) {

	logs, sub, err := _DiamondCutFacet.contract.FilterLogs(opts, "DiamondCut0")
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetDiamondCut0Iterator{contract: _DiamondCutFacet.contract, event: "DiamondCut0", logs: logs, sub: sub}, nil
}

// WatchDiamondCut0 is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) WatchDiamondCut0(opts *bind.WatchOpts, sink chan<- *DiamondCutFacetDiamondCut0) (event.Subscription, error) {

	logs, sub, err := _DiamondCutFacet.contract.WatchLogs(opts, "DiamondCut0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutFacetDiamondCut0)
				if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut0", log); err != nil {
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

// ParseDiamondCut0 is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) ParseDiamondCut0(log types.Log) (*DiamondCutFacetDiamondCut0, error) {
	event := new(DiamondCutFacetDiamondCut0)
	if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
