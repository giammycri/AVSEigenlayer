// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package taskavsregistrar

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

// TaskAVSRegistrarMetaData contains all meta data concerning the TaskAVSRegistrar contract.
var TaskAVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRegisteredOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorAllowlisted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredOperatorsList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeOperatorFromAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSInitialized\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAllowlisted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516111a53803806111a583398101604081905261002e9161022d565b6001600160a01b0383166100895760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420416c6c6f636174696f6e4d616e616765720000000000000060448201526064015b60405180910390fd5b6001600160a01b0382166100df5760405162461bcd60e51b815260206004820152601460248201527f496e76616c6964204b65795265676973747261720000000000000000000000006044820152606401610080565b6001600160a01b0381166101355760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964205065726d697373696f6e436f6e74726f6c6c6572000000006044820152606401610080565b6001600160a01b0380841660805282811660a052811660c05261015661015e565b505050610277565b5f54610100900460ff16156101c55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610080565b5f5460ff90811614610214575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461022a575f5ffd5b50565b5f5f5f6060848603121561023f575f5ffd5b835161024a81610216565b602085015190935061025b81610216565b604085015190925061026c81610216565b809150509250925092565b60805160a05160c051610f046102a15f395f6101a701525f61016801525f6102950152610f045ff3fe608060405234801561000f575f5ffd5b5060043610610106575f3560e01c8063715018a61161009e578063c63fd5021161006e578063c63fd5021461027d578063ca8aa7c714610290578063da324c13146102b7578063f2fde38b146102ca578063fffc9182146102dd575f5ffd5b8063715018a6146102175780638b8738a01461021f5780638da5cb5b1461024a578063b52657871461025b575f5ffd5b80633ec45c7e116100d95780633ec45c7e146101635780634657e26a146101a2578063485cc955146101c95780636b1906f8146101dc575f5ffd5b806317aef8751461010a5780631fc372781461011f578063241217fa1461013d578063303ca95614610150575b5f5ffd5b61011d610118366004610bdc565b6102f0565b005b610127610305565b6040516101349190610bfc565b60405180910390f35b61011d61014b366004610bdc565b610479565b61011d61015e366004610c8f565b61048c565b61018a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610134565b61018a7f000000000000000000000000000000000000000000000000000000000000000081565b61011d6101d7366004610cec565b61057f565b6102076101ea366004610bdc565b6001600160a01b03165f9081526066602052604090205460ff1690565b6040519015158152602001610134565b61011d61077a565b61020761022d366004610bdc565b6001600160a01b03165f9081526067602052604090205460ff1690565b6033546001600160a01b031661018a565b610207610269366004610bdc565b6065546001600160a01b0391821691161490565b61011d61028b366004610d31565b61078d565b61018a7f000000000000000000000000000000000000000000000000000000000000000081565b60655461018a906001600160a01b031681565b61011d6102d8366004610bdc565b610974565b61018a6102eb366004610e34565b6109ea565b6102f8610a12565b610302815f610a6c565b50565b60605f805b60685481101561036c5760665f6068838154811061032a5761032a610e4b565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615610364578161036081610e5f565b9250505b60010161030a565b505f8167ffffffffffffffff81111561038757610387610d1d565b6040519080825280602002602001820160405280156103b0578160200160208202803683370190505b5090505f805b6068548110156104705760665f606883815481106103d6576103d6610e4b565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615610468576068818154811061041557610415610e4b565b905f5260205f20015f9054906101000a90046001600160a01b031683838151811061044257610442610e4b565b6001600160a01b03909216602092830291909101909101528161046481610e5f565b9250505b6001016103b6565b50909392505050565b610481610a12565b610302816001610a6c565b6065546001600160a01b038481169116146104dc5760405162461bcd60e51b815260206004820152600b60248201526a496e76616c69642041565360a81b60448201526064015b60405180910390fd5b6001600160a01b0384165f9081526066602052604090205460ff166105345760405162461bcd60e51b815260206004820152600e60248201526d139bdd081c9959da5cdd195c995960921b60448201526064016104d3565b6001600160a01b0384165f81815260666020526040808220805460ff19169055517f6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d9190a250505050565b5f54610100900460ff161580801561059d57505f54600160ff909116105b806105b65750303b1580156105b657505f5460ff166001145b6106195760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104d3565b5f805460ff19166001179055801561063a575f805461ff0019166101001790555b6001600160a01b0383166106865760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420415653206164647265737360681b60448201526064016104d3565b6001600160a01b0382166106d45760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b60448201526064016104d3565b6106dc610b13565b6106e582610b41565b606580546001600160a01b0319166001600160a01b03858116918217909255604051918416917f6f590f00594e422d605859c867818ebccb321e2bdc9dd35c5edb1e0a067d0601905f90a38015610775575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610782610a12565b61078b5f610b41565b565b6065546001600160a01b038581169116146107d85760405162461bcd60e51b815260206004820152600b60248201526a496e76616c69642041565360a81b60448201526064016104d3565b6001600160a01b0385166108215760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b21037b832b930ba37b960811b60448201526064016104d3565b6001600160a01b0385165f9081526066602052604090205460ff161561087e5760405162461bcd60e51b8152602060048201526012602482015271105b1c9958591e481c9959da5cdd195c995960721b60448201526064016104d3565b6001600160a01b0385165f9081526067602052604090205460ff166108e55760405162461bcd60e51b815260206004820152601860248201527f4f70657261746f72206e6f7420616c6c6f776c6973746564000000000000000060448201526064016104d3565b6001600160a01b0385165f81815260666020526040808220805460ff1916600190811790915560688054918201815583527fa2153420d844928b4421650203c77babc8b33d7f2e7b450e2966db0c220977530180546001600160a01b03191684179055517f4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e59190a25050505050565b61097c610a12565b6001600160a01b0381166109e15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104d3565b61030281610b41565b606881815481106109f9575f80fd5b5f918252602090912001546001600160a01b0316905081565b6033546001600160a01b0316331461078b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d3565b6001600160a01b038216610ab55760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b21037b832b930ba37b960811b60448201526064016104d3565b6001600160a01b0382165f81815260676020908152604091829020805460ff191685151590811790915591519182527fbc9494536acaf3f702bbec70ff790bb3930056ae9e781e8077097b9de89a8e18910160405180910390a25050565b5f54610100900460ff16610b395760405162461bcd60e51b81526004016104d390610e83565b61078b610b92565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610bb85760405162461bcd60e51b81526004016104d390610e83565b61078b33610b41565b80356001600160a01b0381168114610bd7575f5ffd5b919050565b5f60208284031215610bec575f5ffd5b610bf582610bc1565b9392505050565b602080825282518282018190525f918401906040840190835b81811015610c3c5783516001600160a01b0316835260209384019390920191600101610c15565b509095945050505050565b5f5f83601f840112610c57575f5ffd5b50813567ffffffffffffffff811115610c6e575f5ffd5b6020830191508360208260051b8501011115610c88575f5ffd5b9250929050565b5f5f5f5f60608587031215610ca2575f5ffd5b610cab85610bc1565b9350610cb960208601610bc1565b9250604085013567ffffffffffffffff811115610cd4575f5ffd5b610ce087828801610c47565b95989497509550505050565b5f5f60408385031215610cfd575f5ffd5b610d0683610bc1565b9150610d1460208401610bc1565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f60808688031215610d45575f5ffd5b610d4e86610bc1565b9450610d5c60208701610bc1565b9350604086013567ffffffffffffffff811115610d77575f5ffd5b610d8388828901610c47565b909450925050606086013567ffffffffffffffff811115610da2575f5ffd5b8601601f81018813610db2575f5ffd5b803567ffffffffffffffff811115610dcc57610dcc610d1d565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610dfb57610dfb610d1d565b6040528181528282016020018a1015610e12575f5ffd5b816020840160208301375f602083830101528093505050509295509295909350565b5f60208284031215610e44575f5ffd5b5035919050565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610e7c57634e487b7160e01b5f52601160045260245ffd5b5060010190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea26469706673582212204c6898b555c189b55317a6ddb395fb9cdf74b4494ecd349e062ba94f4e08bbd464736f6c634300081b0033",
}

// TaskAVSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskAVSRegistrarMetaData.ABI instead.
var TaskAVSRegistrarABI = TaskAVSRegistrarMetaData.ABI

// TaskAVSRegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskAVSRegistrarMetaData.Bin instead.
var TaskAVSRegistrarBin = TaskAVSRegistrarMetaData.Bin

// DeployTaskAVSRegistrar deploys a new Ethereum contract, binding an instance of TaskAVSRegistrar to it.
func DeployTaskAVSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, _allocationManager common.Address, _keyRegistrar common.Address, _permissionController common.Address) (common.Address, *types.Transaction, *TaskAVSRegistrar, error) {
	parsed, err := TaskAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskAVSRegistrarBin), backend, _allocationManager, _keyRegistrar, _permissionController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TaskAVSRegistrar{TaskAVSRegistrarCaller: TaskAVSRegistrarCaller{contract: contract}, TaskAVSRegistrarTransactor: TaskAVSRegistrarTransactor{contract: contract}, TaskAVSRegistrarFilterer: TaskAVSRegistrarFilterer{contract: contract}}, nil
}

// TaskAVSRegistrar is an auto generated Go binding around an Ethereum contract.
type TaskAVSRegistrar struct {
	TaskAVSRegistrarCaller     // Read-only binding to the contract
	TaskAVSRegistrarTransactor // Write-only binding to the contract
	TaskAVSRegistrarFilterer   // Log filterer for contract events
}

// TaskAVSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskAVSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskAVSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskAVSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskAVSRegistrarSession struct {
	Contract     *TaskAVSRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskAVSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskAVSRegistrarCallerSession struct {
	Contract *TaskAVSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TaskAVSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskAVSRegistrarTransactorSession struct {
	Contract     *TaskAVSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TaskAVSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskAVSRegistrarRaw struct {
	Contract *TaskAVSRegistrar // Generic contract binding to access the raw methods on
}

// TaskAVSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskAVSRegistrarCallerRaw struct {
	Contract *TaskAVSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// TaskAVSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskAVSRegistrarTransactorRaw struct {
	Contract *TaskAVSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskAVSRegistrar creates a new instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrar(address common.Address, backend bind.ContractBackend) (*TaskAVSRegistrar, error) {
	contract, err := bindTaskAVSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrar{TaskAVSRegistrarCaller: TaskAVSRegistrarCaller{contract: contract}, TaskAVSRegistrarTransactor: TaskAVSRegistrarTransactor{contract: contract}, TaskAVSRegistrarFilterer: TaskAVSRegistrarFilterer{contract: contract}}, nil
}

// NewTaskAVSRegistrarCaller creates a new read-only instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*TaskAVSRegistrarCaller, error) {
	contract, err := bindTaskAVSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarCaller{contract: contract}, nil
}

// NewTaskAVSRegistrarTransactor creates a new write-only instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskAVSRegistrarTransactor, error) {
	contract, err := bindTaskAVSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarTransactor{contract: contract}, nil
}

// NewTaskAVSRegistrarFilterer creates a new log filterer instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskAVSRegistrarFilterer, error) {
	contract, err := bindTaskAVSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarFilterer{contract: contract}, nil
}

// bindTaskAVSRegistrar binds a generic wrapper to an already deployed contract.
func bindTaskAVSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskAVSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AllocationManager() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AllocationManager(&_TaskAVSRegistrar.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) AllocationManager() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AllocationManager(&_TaskAVSRegistrar.CallOpts)
}

// AvsAddress is a free data retrieval call binding the contract method 0xda324c13.
//
// Solidity: function avsAddress() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) AvsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "avsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsAddress is a free data retrieval call binding the contract method 0xda324c13.
//
// Solidity: function avsAddress() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AvsAddress() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AvsAddress(&_TaskAVSRegistrar.CallOpts)
}

// AvsAddress is a free data retrieval call binding the contract method 0xda324c13.
//
// Solidity: function avsAddress() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) AvsAddress() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AvsAddress(&_TaskAVSRegistrar.CallOpts)
}

// GetRegisteredOperators is a free data retrieval call binding the contract method 0x1fc37278.
//
// Solidity: function getRegisteredOperators() view returns(address[])
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetRegisteredOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getRegisteredOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredOperators is a free data retrieval call binding the contract method 0x1fc37278.
//
// Solidity: function getRegisteredOperators() view returns(address[])
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetRegisteredOperators() ([]common.Address, error) {
	return _TaskAVSRegistrar.Contract.GetRegisteredOperators(&_TaskAVSRegistrar.CallOpts)
}

// GetRegisteredOperators is a free data retrieval call binding the contract method 0x1fc37278.
//
// Solidity: function getRegisteredOperators() view returns(address[])
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetRegisteredOperators() ([]common.Address, error) {
	return _TaskAVSRegistrar.Contract.GetRegisteredOperators(&_TaskAVSRegistrar.CallOpts)
}

// IsOperatorAllowlisted is a free data retrieval call binding the contract method 0x8b8738a0.
//
// Solidity: function isOperatorAllowlisted(address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) IsOperatorAllowlisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "isOperatorAllowlisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorAllowlisted is a free data retrieval call binding the contract method 0x8b8738a0.
//
// Solidity: function isOperatorAllowlisted(address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) IsOperatorAllowlisted(operator common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.IsOperatorAllowlisted(&_TaskAVSRegistrar.CallOpts, operator)
}

// IsOperatorAllowlisted is a free data retrieval call binding the contract method 0x8b8738a0.
//
// Solidity: function isOperatorAllowlisted(address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) IsOperatorAllowlisted(operator common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.IsOperatorAllowlisted(&_TaskAVSRegistrar.CallOpts, operator)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "isOperatorRegistered", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.IsOperatorRegistered(&_TaskAVSRegistrar.CallOpts, operator)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.IsOperatorRegistered(&_TaskAVSRegistrar.CallOpts, operator)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) KeyRegistrar() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.KeyRegistrar(&_TaskAVSRegistrar.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) KeyRegistrar() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.KeyRegistrar(&_TaskAVSRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) Owner() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Owner(&_TaskAVSRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) Owner() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Owner(&_TaskAVSRegistrar.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) PermissionController() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.PermissionController(&_TaskAVSRegistrar.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) PermissionController() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.PermissionController(&_TaskAVSRegistrar.CallOpts)
}

// RegisteredOperatorsList is a free data retrieval call binding the contract method 0xfffc9182.
//
// Solidity: function registeredOperatorsList(uint256 ) view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) RegisteredOperatorsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "registeredOperatorsList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredOperatorsList is a free data retrieval call binding the contract method 0xfffc9182.
//
// Solidity: function registeredOperatorsList(uint256 ) view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RegisteredOperatorsList(arg0 *big.Int) (common.Address, error) {
	return _TaskAVSRegistrar.Contract.RegisteredOperatorsList(&_TaskAVSRegistrar.CallOpts, arg0)
}

// RegisteredOperatorsList is a free data retrieval call binding the contract method 0xfffc9182.
//
// Solidity: function registeredOperatorsList(uint256 ) view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) RegisteredOperatorsList(arg0 *big.Int) (common.Address, error) {
	return _TaskAVSRegistrar.Contract.RegisteredOperatorsList(&_TaskAVSRegistrar.CallOpts, arg0)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) SupportsAVS(opts *bind.CallOpts, avs common.Address) (bool, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "supportsAVS", avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) SupportsAVS(avs common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.SupportsAVS(&_TaskAVSRegistrar.CallOpts, avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) SupportsAVS(avs common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.SupportsAVS(&_TaskAVSRegistrar.CallOpts, avs)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x241217fa.
//
// Solidity: function addOperatorToAllowlist(address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) AddOperatorToAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "addOperatorToAllowlist", operator)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x241217fa.
//
// Solidity: function addOperatorToAllowlist(address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AddOperatorToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.AddOperatorToAllowlist(&_TaskAVSRegistrar.TransactOpts, operator)
}

// AddOperatorToAllowlist is a paid mutator transaction binding the contract method 0x241217fa.
//
// Solidity: function addOperatorToAllowlist(address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) AddOperatorToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.AddOperatorToAllowlist(&_TaskAVSRegistrar.TransactOpts, operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "deregisterOperator", operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.DeregisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.DeregisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _avs, address _owner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) Initialize(opts *bind.TransactOpts, _avs common.Address, _owner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "initialize", _avs, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _avs, address _owner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) Initialize(_avs common.Address, _owner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.Initialize(&_TaskAVSRegistrar.TransactOpts, _avs, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _avs, address _owner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) Initialize(_avs common.Address, _owner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.Initialize(&_TaskAVSRegistrar.TransactOpts, _avs, _owner)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes operatorSignature) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "registerOperator", operator, avs, operatorSetIds, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes operatorSignature) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RegisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes operatorSignature) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, operatorSignature []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RegisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds, operatorSignature)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x17aef875.
//
// Solidity: function removeOperatorFromAllowlist(address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RemoveOperatorFromAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "removeOperatorFromAllowlist", operator)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x17aef875.
//
// Solidity: function removeOperatorFromAllowlist(address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RemoveOperatorFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RemoveOperatorFromAllowlist(&_TaskAVSRegistrar.TransactOpts, operator)
}

// RemoveOperatorFromAllowlist is a paid mutator transaction binding the contract method 0x17aef875.
//
// Solidity: function removeOperatorFromAllowlist(address operator) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RemoveOperatorFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RemoveOperatorFromAllowlist(&_TaskAVSRegistrar.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RenounceOwnership(&_TaskAVSRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RenounceOwnership(&_TaskAVSRegistrar.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TransferOwnership(&_TaskAVSRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TransferOwnership(&_TaskAVSRegistrar.TransactOpts, newOwner)
}

// TaskAVSRegistrarAVSInitializedIterator is returned from FilterAVSInitialized and is used to iterate over the raw logs and unpacked data for AVSInitialized events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarAVSInitializedIterator struct {
	Event *TaskAVSRegistrarAVSInitialized // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarAVSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarAVSInitialized)
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
		it.Event = new(TaskAVSRegistrarAVSInitialized)
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
func (it *TaskAVSRegistrarAVSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarAVSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarAVSInitialized represents a AVSInitialized event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarAVSInitialized struct {
	Avs   common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAVSInitialized is a free log retrieval operation binding the contract event 0x6f590f00594e422d605859c867818ebccb321e2bdc9dd35c5edb1e0a067d0601.
//
// Solidity: event AVSInitialized(address indexed avs, address indexed owner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterAVSInitialized(opts *bind.FilterOpts, avs []common.Address, owner []common.Address) (*TaskAVSRegistrarAVSInitializedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "AVSInitialized", avsRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarAVSInitializedIterator{contract: _TaskAVSRegistrar.contract, event: "AVSInitialized", logs: logs, sub: sub}, nil
}

// WatchAVSInitialized is a free log subscription operation binding the contract event 0x6f590f00594e422d605859c867818ebccb321e2bdc9dd35c5edb1e0a067d0601.
//
// Solidity: event AVSInitialized(address indexed avs, address indexed owner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchAVSInitialized(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarAVSInitialized, avs []common.Address, owner []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "AVSInitialized", avsRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarAVSInitialized)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "AVSInitialized", log); err != nil {
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

// ParseAVSInitialized is a log parse operation binding the contract event 0x6f590f00594e422d605859c867818ebccb321e2bdc9dd35c5edb1e0a067d0601.
//
// Solidity: event AVSInitialized(address indexed avs, address indexed owner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseAVSInitialized(log types.Log) (*TaskAVSRegistrarAVSInitialized, error) {
	event := new(TaskAVSRegistrarAVSInitialized)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "AVSInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarInitializedIterator struct {
	Event *TaskAVSRegistrarInitialized // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarInitialized)
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
		it.Event = new(TaskAVSRegistrarInitialized)
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
func (it *TaskAVSRegistrarInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarInitialized represents a Initialized event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskAVSRegistrarInitializedIterator, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarInitializedIterator{contract: _TaskAVSRegistrar.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarInitialized) (event.Subscription, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarInitialized)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseInitialized(log types.Log) (*TaskAVSRegistrarInitialized, error) {
	event := new(TaskAVSRegistrarInitialized)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorAllowlistedIterator is returned from FilterOperatorAllowlisted and is used to iterate over the raw logs and unpacked data for OperatorAllowlisted events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorAllowlistedIterator struct {
	Event *TaskAVSRegistrarOperatorAllowlisted // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorAllowlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorAllowlisted)
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
		it.Event = new(TaskAVSRegistrarOperatorAllowlisted)
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
func (it *TaskAVSRegistrarOperatorAllowlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorAllowlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorAllowlisted represents a OperatorAllowlisted event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorAllowlisted struct {
	Operator common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAllowlisted is a free log retrieval operation binding the contract event 0xbc9494536acaf3f702bbec70ff790bb3930056ae9e781e8077097b9de89a8e18.
//
// Solidity: event OperatorAllowlisted(address indexed operator, bool allowed)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorAllowlisted(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorAllowlistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorAllowlisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorAllowlistedIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorAllowlisted", logs: logs, sub: sub}, nil
}

// WatchOperatorAllowlisted is a free log subscription operation binding the contract event 0xbc9494536acaf3f702bbec70ff790bb3930056ae9e781e8077097b9de89a8e18.
//
// Solidity: event OperatorAllowlisted(address indexed operator, bool allowed)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorAllowlisted(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorAllowlisted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorAllowlisted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorAllowlisted)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorAllowlisted", log); err != nil {
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

// ParseOperatorAllowlisted is a log parse operation binding the contract event 0xbc9494536acaf3f702bbec70ff790bb3930056ae9e781e8077097b9de89a8e18.
//
// Solidity: event OperatorAllowlisted(address indexed operator, bool allowed)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorAllowlisted(log types.Log) (*TaskAVSRegistrarOperatorAllowlisted, error) {
	event := new(TaskAVSRegistrarOperatorAllowlisted)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorAllowlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorDeregisteredIterator struct {
	Event *TaskAVSRegistrarOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorDeregistered)
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
		it.Event = new(TaskAVSRegistrarOperatorDeregistered)
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
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorDeregistered represents a OperatorDeregistered event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorDeregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorDeregisteredIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorDeregistered)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorDeregistered(log types.Log) (*TaskAVSRegistrarOperatorDeregistered, error) {
	event := new(TaskAVSRegistrarOperatorDeregistered)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRegisteredIterator struct {
	Event *TaskAVSRegistrarOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorRegistered)
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
		it.Event = new(TaskAVSRegistrarOperatorRegistered)
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
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorRegistered represents a OperatorRegistered event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorRegisteredIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorRegistered)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorRegistered(log types.Log) (*TaskAVSRegistrarOperatorRegistered, error) {
	event := new(TaskAVSRegistrarOperatorRegistered)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOwnershipTransferredIterator struct {
	Event *TaskAVSRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOwnershipTransferred)
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
		it.Event = new(TaskAVSRegistrarOwnershipTransferred)
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
func (it *TaskAVSRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaskAVSRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOwnershipTransferredIterator{contract: _TaskAVSRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOwnershipTransferred)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*TaskAVSRegistrarOwnershipTransferred, error) {
	event := new(TaskAVSRegistrarOwnershipTransferred)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
