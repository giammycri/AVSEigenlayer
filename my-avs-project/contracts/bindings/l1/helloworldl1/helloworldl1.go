// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworldl1

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

// HelloWorldL1SumVerificationTask is an auto generated low-level Go binding around an user-defined struct.
type HelloWorldL1SumVerificationTask struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Verified      bool
	IsCorrect     bool
	Timestamp     *big.Int
}

// HelloWorldL1MetaData contains all meta data concerning the HelloWorldL1 contract.
var HelloWorldL1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeSumVerification\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHelloWorldL1.SumVerificationTask\",\"components\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verified\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTaskVerified\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestSumVerification\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tasks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verified\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifySumOnChain\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"SumVerificationCompleted\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"actualResult\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SumVerificationRequested\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"a\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f5ffd5b506100316040518060600160405280602c8152602001610a1a602c9139610036565b6100e7565b61007d8160405160240161004a91906100b2565b60408051601f198184030181529190526020810180516001600160e01b0390811663104c13eb60e21b1790915261008016565b50565b61007d8161009360201b61069f1760201c565b5f6a636f6e736f6c652e6c6f6790505f5f835160208501845afa505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b610926806100f45f395ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c80636ef6a577116100585780636ef6a577146101225780638d9776721461014557806394b7a9bf146101e1578063d8fb6f841461020d575f5ffd5b80631d65e77e1461007e5780632ee28d8a146100f85780635867173014610119575b5f5ffd5b61009161008c3660046107a1565b610222565b6040516100ef91908151815260208083015190820152604080830151908201526060808301516001600160a01b03169082015260808083015115159082015260a08083015115159082015260c0918201519181019190915260e00190565b60405180910390f35b61010b6101063660046107b8565b6102e0565b6040519081526020016100ef565b61010b60015481565b6101356101303660046107b8565b6104d6565b60405190151581526020016100ef565b6101a06101533660046107a1565b5f60208190529081526040902080546001820154600283015460038401546004909401549293919290916001600160a01b0381169160ff600160a01b8304811692600160a81b9004169087565b604080519788526020880196909652948601939093526001600160a01b03909116606085015215156080840152151560a083015260c082015260e0016100ef565b6101356101ef3660046107a1565b5f90815260208190526040902060030154600160a01b900460ff1690565b61022061021b3660046107e1565b6104eb565b005b6102686040518060e001604052805f81526020015f81526020015f81526020015f6001600160a01b031681526020015f151581526020015f151581526020015f81525090565b505f9081526020818152604091829020825160e0810184528154815260018201549281019290925260028101549282019290925260038201546001600160a01b038116606083015260ff600160a01b8204811615156080840152600160a81b90910416151560a082015260049091015460c082015290565b600180545f91826102f083610827565b909155506040805160e08101825286815260208082018781528284018781523360608086018281525f6080880181815260a089018281524260c08b019081528c8452838a52928b902099518a55965160018a0155945160028901559051600388018054955196511515600160a81b0260ff60a81b19971515600160a01b026001600160a81b03199097166001600160a01b039093169290921795909517959095169490941790925591516004909401939093558351898152918201889052928101869052929350909183917f74396ec72e1af7caf32d3950fc17c8d8e32f4efe70ee2ec3c327f4091dab1347910160405180910390a36104246040518060400160405280601b81526020017f53756d20766572696669636174696f6e207265717565737465643a00000000008152506106be565b6104506040518060400160405280600a81526020016910102a30b9b59024a21d60b11b81525082610704565b610476604051806040016040528060048152602001631010209d60e11b81525085610704565b61049c604051806040016040528060048152602001631010211d60e11b81525084610704565b6104cf60405180604001604052806011815260200170101021b630b4b6b2b2102932b9bab63a1d60791b81525083610704565b9392505050565b5f816104e2848661083f565b14949350505050565b5f82815260208190526040902060030154600160a01b900460ff161561054f5760405162461bcd60e51b815260206004820152601560248201527415185cdac8185b1c9958591e481d995c9a599a5959605a1b604482015260640160405180910390fd5b5f828152602081905260408120600381018054841515600160a81b0261ffff60a01b1990911617600160a01b1790556001810154905461058f919061083f565b6040805184151581526020810183905291925084917fac2034b96316771cf452968ad23adf2a71e4c283d311ad2b66de9b6f7f91a7dd910160405180910390a261060d6040518060400160405280601b81526020017f53756d20766572696669636174696f6e20636f6d706c657465643a00000000008152506106be565b6106396040518060400160405280600a81526020016910102a30b9b59024a21d60b11b81525084610704565b6106686040518060400160405280600d81526020016c101024b99021b7b93932b1ba1d60991b8152508361074d565b61069a6040518060400160405280601081526020016f101020b1ba3ab0b6102932b9bab63a1d60811b81525082610704565b505050565b5f6a636f6e736f6c652e6c6f6790505f5f835160208501845afa505050565b610701816040516024016106d29190610886565b60408051601f198184030181529190526020810180516001600160e01b031663104c13eb60e21b17905261078e565b50565b610749828260405160240161071a929190610898565b60408051601f198184030181529190526020810180516001600160e01b0316632d839cb360e21b17905261078e565b5050565b61074982826040516024016107639291906108b9565b60408051601f198184030181529190526020810180516001600160e01b031663c3b5563560e01b1790525b6107018161069f565b61079f6108dc565b565b5f602082840312156107b1575f5ffd5b5035919050565b5f5f5f606084860312156107ca575f5ffd5b505081359360208301359350604090920135919050565b5f5f604083850312156107f2575f5ffd5b8235915060208301358015158114610808575f5ffd5b809150509250929050565b634e487b7160e01b5f52601160045260245ffd5b5f6001820161083857610838610813565b5060010190565b8082018082111561085257610852610813565b92915050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6104cf6020830184610858565b604081525f6108aa6040830185610858565b90508260208301529392505050565b604081525f6108cb6040830185610858565b905082151560208301529392505050565b634e487b7160e01b5f52605160045260245ffdfea26469706673582212206bd1a72e79985ad917f689a17ded29919929639fe4a33c90d18b9e49f20a645164736f6c634300081b003348656c6c6f576f726c644c31206465706c6f796564202d2053756d20566572696669636174696f6e20415653",
}

// HelloWorldL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloWorldL1MetaData.ABI instead.
var HelloWorldL1ABI = HelloWorldL1MetaData.ABI

// HelloWorldL1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloWorldL1MetaData.Bin instead.
var HelloWorldL1Bin = HelloWorldL1MetaData.Bin

// DeployHelloWorldL1 deploys a new Ethereum contract, binding an instance of HelloWorldL1 to it.
func DeployHelloWorldL1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HelloWorldL1, error) {
	parsed, err := HelloWorldL1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloWorldL1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HelloWorldL1{HelloWorldL1Caller: HelloWorldL1Caller{contract: contract}, HelloWorldL1Transactor: HelloWorldL1Transactor{contract: contract}, HelloWorldL1Filterer: HelloWorldL1Filterer{contract: contract}}, nil
}

// HelloWorldL1 is an auto generated Go binding around an Ethereum contract.
type HelloWorldL1 struct {
	HelloWorldL1Caller     // Read-only binding to the contract
	HelloWorldL1Transactor // Write-only binding to the contract
	HelloWorldL1Filterer   // Log filterer for contract events
}

// HelloWorldL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type HelloWorldL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloWorldL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloWorldL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloWorldL1Session struct {
	Contract     *HelloWorldL1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloWorldL1CallerSession struct {
	Contract *HelloWorldL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HelloWorldL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloWorldL1TransactorSession struct {
	Contract     *HelloWorldL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HelloWorldL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type HelloWorldL1Raw struct {
	Contract *HelloWorldL1 // Generic contract binding to access the raw methods on
}

// HelloWorldL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloWorldL1CallerRaw struct {
	Contract *HelloWorldL1Caller // Generic read-only contract binding to access the raw methods on
}

// HelloWorldL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloWorldL1TransactorRaw struct {
	Contract *HelloWorldL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloWorldL1 creates a new instance of HelloWorldL1, bound to a specific deployed contract.
func NewHelloWorldL1(address common.Address, backend bind.ContractBackend) (*HelloWorldL1, error) {
	contract, err := bindHelloWorldL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL1{HelloWorldL1Caller: HelloWorldL1Caller{contract: contract}, HelloWorldL1Transactor: HelloWorldL1Transactor{contract: contract}, HelloWorldL1Filterer: HelloWorldL1Filterer{contract: contract}}, nil
}

// NewHelloWorldL1Caller creates a new read-only instance of HelloWorldL1, bound to a specific deployed contract.
func NewHelloWorldL1Caller(address common.Address, caller bind.ContractCaller) (*HelloWorldL1Caller, error) {
	contract, err := bindHelloWorldL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL1Caller{contract: contract}, nil
}

// NewHelloWorldL1Transactor creates a new write-only instance of HelloWorldL1, bound to a specific deployed contract.
func NewHelloWorldL1Transactor(address common.Address, transactor bind.ContractTransactor) (*HelloWorldL1Transactor, error) {
	contract, err := bindHelloWorldL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL1Transactor{contract: contract}, nil
}

// NewHelloWorldL1Filterer creates a new log filterer instance of HelloWorldL1, bound to a specific deployed contract.
func NewHelloWorldL1Filterer(address common.Address, filterer bind.ContractFilterer) (*HelloWorldL1Filterer, error) {
	contract, err := bindHelloWorldL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL1Filterer{contract: contract}, nil
}

// bindHelloWorldL1 binds a generic wrapper to an already deployed contract.
func bindHelloWorldL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloWorldL1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorldL1 *HelloWorldL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorldL1.Contract.HelloWorldL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorldL1 *HelloWorldL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.HelloWorldL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorldL1 *HelloWorldL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.HelloWorldL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorldL1 *HelloWorldL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorldL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorldL1 *HelloWorldL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorldL1 *HelloWorldL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,uint256,uint256,address,bool,bool,uint256))
func (_HelloWorldL1 *HelloWorldL1Caller) GetTask(opts *bind.CallOpts, taskId *big.Int) (HelloWorldL1SumVerificationTask, error) {
	var out []interface{}
	err := _HelloWorldL1.contract.Call(opts, &out, "getTask", taskId)

	if err != nil {
		return *new(HelloWorldL1SumVerificationTask), err
	}

	out0 := *abi.ConvertType(out[0], new(HelloWorldL1SumVerificationTask)).(*HelloWorldL1SumVerificationTask)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,uint256,uint256,address,bool,bool,uint256))
func (_HelloWorldL1 *HelloWorldL1Session) GetTask(taskId *big.Int) (HelloWorldL1SumVerificationTask, error) {
	return _HelloWorldL1.Contract.GetTask(&_HelloWorldL1.CallOpts, taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 taskId) view returns((uint256,uint256,uint256,address,bool,bool,uint256))
func (_HelloWorldL1 *HelloWorldL1CallerSession) GetTask(taskId *big.Int) (HelloWorldL1SumVerificationTask, error) {
	return _HelloWorldL1.Contract.GetTask(&_HelloWorldL1.CallOpts, taskId)
}

// IsTaskVerified is a free data retrieval call binding the contract method 0x94b7a9bf.
//
// Solidity: function isTaskVerified(uint256 taskId) view returns(bool)
func (_HelloWorldL1 *HelloWorldL1Caller) IsTaskVerified(opts *bind.CallOpts, taskId *big.Int) (bool, error) {
	var out []interface{}
	err := _HelloWorldL1.contract.Call(opts, &out, "isTaskVerified", taskId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTaskVerified is a free data retrieval call binding the contract method 0x94b7a9bf.
//
// Solidity: function isTaskVerified(uint256 taskId) view returns(bool)
func (_HelloWorldL1 *HelloWorldL1Session) IsTaskVerified(taskId *big.Int) (bool, error) {
	return _HelloWorldL1.Contract.IsTaskVerified(&_HelloWorldL1.CallOpts, taskId)
}

// IsTaskVerified is a free data retrieval call binding the contract method 0x94b7a9bf.
//
// Solidity: function isTaskVerified(uint256 taskId) view returns(bool)
func (_HelloWorldL1 *HelloWorldL1CallerSession) IsTaskVerified(taskId *big.Int) (bool, error) {
	return _HelloWorldL1.Contract.IsTaskVerified(&_HelloWorldL1.CallOpts, taskId)
}

// TaskCounter is a free data retrieval call binding the contract method 0x58671730.
//
// Solidity: function taskCounter() view returns(uint256)
func (_HelloWorldL1 *HelloWorldL1Caller) TaskCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HelloWorldL1.contract.Call(opts, &out, "taskCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCounter is a free data retrieval call binding the contract method 0x58671730.
//
// Solidity: function taskCounter() view returns(uint256)
func (_HelloWorldL1 *HelloWorldL1Session) TaskCounter() (*big.Int, error) {
	return _HelloWorldL1.Contract.TaskCounter(&_HelloWorldL1.CallOpts)
}

// TaskCounter is a free data retrieval call binding the contract method 0x58671730.
//
// Solidity: function taskCounter() view returns(uint256)
func (_HelloWorldL1 *HelloWorldL1CallerSession) TaskCounter() (*big.Int, error) {
	return _HelloWorldL1.Contract.TaskCounter(&_HelloWorldL1.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 a, uint256 b, uint256 claimedResult, address requester, bool verified, bool isCorrect, uint256 timestamp)
func (_HelloWorldL1 *HelloWorldL1Caller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Verified      bool
	IsCorrect     bool
	Timestamp     *big.Int
}, error) {
	var out []interface{}
	err := _HelloWorldL1.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		A             *big.Int
		B             *big.Int
		ClaimedResult *big.Int
		Requester     common.Address
		Verified      bool
		IsCorrect     bool
		Timestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimedResult = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Verified = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.IsCorrect = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 a, uint256 b, uint256 claimedResult, address requester, bool verified, bool isCorrect, uint256 timestamp)
func (_HelloWorldL1 *HelloWorldL1Session) Tasks(arg0 *big.Int) (struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Verified      bool
	IsCorrect     bool
	Timestamp     *big.Int
}, error) {
	return _HelloWorldL1.Contract.Tasks(&_HelloWorldL1.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 a, uint256 b, uint256 claimedResult, address requester, bool verified, bool isCorrect, uint256 timestamp)
func (_HelloWorldL1 *HelloWorldL1CallerSession) Tasks(arg0 *big.Int) (struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Verified      bool
	IsCorrect     bool
	Timestamp     *big.Int
}, error) {
	return _HelloWorldL1.Contract.Tasks(&_HelloWorldL1.CallOpts, arg0)
}

// VerifySumOnChain is a free data retrieval call binding the contract method 0x6ef6a577.
//
// Solidity: function verifySumOnChain(uint256 a, uint256 b, uint256 result) pure returns(bool)
func (_HelloWorldL1 *HelloWorldL1Caller) VerifySumOnChain(opts *bind.CallOpts, a *big.Int, b *big.Int, result *big.Int) (bool, error) {
	var out []interface{}
	err := _HelloWorldL1.contract.Call(opts, &out, "verifySumOnChain", a, b, result)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySumOnChain is a free data retrieval call binding the contract method 0x6ef6a577.
//
// Solidity: function verifySumOnChain(uint256 a, uint256 b, uint256 result) pure returns(bool)
func (_HelloWorldL1 *HelloWorldL1Session) VerifySumOnChain(a *big.Int, b *big.Int, result *big.Int) (bool, error) {
	return _HelloWorldL1.Contract.VerifySumOnChain(&_HelloWorldL1.CallOpts, a, b, result)
}

// VerifySumOnChain is a free data retrieval call binding the contract method 0x6ef6a577.
//
// Solidity: function verifySumOnChain(uint256 a, uint256 b, uint256 result) pure returns(bool)
func (_HelloWorldL1 *HelloWorldL1CallerSession) VerifySumOnChain(a *big.Int, b *big.Int, result *big.Int) (bool, error) {
	return _HelloWorldL1.Contract.VerifySumOnChain(&_HelloWorldL1.CallOpts, a, b, result)
}

// CompleteSumVerification is a paid mutator transaction binding the contract method 0xd8fb6f84.
//
// Solidity: function completeSumVerification(uint256 taskId, bool isCorrect) returns()
func (_HelloWorldL1 *HelloWorldL1Transactor) CompleteSumVerification(opts *bind.TransactOpts, taskId *big.Int, isCorrect bool) (*types.Transaction, error) {
	return _HelloWorldL1.contract.Transact(opts, "completeSumVerification", taskId, isCorrect)
}

// CompleteSumVerification is a paid mutator transaction binding the contract method 0xd8fb6f84.
//
// Solidity: function completeSumVerification(uint256 taskId, bool isCorrect) returns()
func (_HelloWorldL1 *HelloWorldL1Session) CompleteSumVerification(taskId *big.Int, isCorrect bool) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.CompleteSumVerification(&_HelloWorldL1.TransactOpts, taskId, isCorrect)
}

// CompleteSumVerification is a paid mutator transaction binding the contract method 0xd8fb6f84.
//
// Solidity: function completeSumVerification(uint256 taskId, bool isCorrect) returns()
func (_HelloWorldL1 *HelloWorldL1TransactorSession) CompleteSumVerification(taskId *big.Int, isCorrect bool) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.CompleteSumVerification(&_HelloWorldL1.TransactOpts, taskId, isCorrect)
}

// RequestSumVerification is a paid mutator transaction binding the contract method 0x2ee28d8a.
//
// Solidity: function requestSumVerification(uint256 a, uint256 b, uint256 claimedResult) returns(uint256 taskId)
func (_HelloWorldL1 *HelloWorldL1Transactor) RequestSumVerification(opts *bind.TransactOpts, a *big.Int, b *big.Int, claimedResult *big.Int) (*types.Transaction, error) {
	return _HelloWorldL1.contract.Transact(opts, "requestSumVerification", a, b, claimedResult)
}

// RequestSumVerification is a paid mutator transaction binding the contract method 0x2ee28d8a.
//
// Solidity: function requestSumVerification(uint256 a, uint256 b, uint256 claimedResult) returns(uint256 taskId)
func (_HelloWorldL1 *HelloWorldL1Session) RequestSumVerification(a *big.Int, b *big.Int, claimedResult *big.Int) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.RequestSumVerification(&_HelloWorldL1.TransactOpts, a, b, claimedResult)
}

// RequestSumVerification is a paid mutator transaction binding the contract method 0x2ee28d8a.
//
// Solidity: function requestSumVerification(uint256 a, uint256 b, uint256 claimedResult) returns(uint256 taskId)
func (_HelloWorldL1 *HelloWorldL1TransactorSession) RequestSumVerification(a *big.Int, b *big.Int, claimedResult *big.Int) (*types.Transaction, error) {
	return _HelloWorldL1.Contract.RequestSumVerification(&_HelloWorldL1.TransactOpts, a, b, claimedResult)
}

// HelloWorldL1SumVerificationCompletedIterator is returned from FilterSumVerificationCompleted and is used to iterate over the raw logs and unpacked data for SumVerificationCompleted events raised by the HelloWorldL1 contract.
type HelloWorldL1SumVerificationCompletedIterator struct {
	Event *HelloWorldL1SumVerificationCompleted // Event containing the contract specifics and raw log

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
func (it *HelloWorldL1SumVerificationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldL1SumVerificationCompleted)
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
		it.Event = new(HelloWorldL1SumVerificationCompleted)
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
func (it *HelloWorldL1SumVerificationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldL1SumVerificationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldL1SumVerificationCompleted represents a SumVerificationCompleted event raised by the HelloWorldL1 contract.
type HelloWorldL1SumVerificationCompleted struct {
	TaskId       *big.Int
	IsCorrect    bool
	ActualResult *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSumVerificationCompleted is a free log retrieval operation binding the contract event 0xac2034b96316771cf452968ad23adf2a71e4c283d311ad2b66de9b6f7f91a7dd.
//
// Solidity: event SumVerificationCompleted(uint256 indexed taskId, bool isCorrect, uint256 actualResult)
func (_HelloWorldL1 *HelloWorldL1Filterer) FilterSumVerificationCompleted(opts *bind.FilterOpts, taskId []*big.Int) (*HelloWorldL1SumVerificationCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _HelloWorldL1.contract.FilterLogs(opts, "SumVerificationCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL1SumVerificationCompletedIterator{contract: _HelloWorldL1.contract, event: "SumVerificationCompleted", logs: logs, sub: sub}, nil
}

// WatchSumVerificationCompleted is a free log subscription operation binding the contract event 0xac2034b96316771cf452968ad23adf2a71e4c283d311ad2b66de9b6f7f91a7dd.
//
// Solidity: event SumVerificationCompleted(uint256 indexed taskId, bool isCorrect, uint256 actualResult)
func (_HelloWorldL1 *HelloWorldL1Filterer) WatchSumVerificationCompleted(opts *bind.WatchOpts, sink chan<- *HelloWorldL1SumVerificationCompleted, taskId []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _HelloWorldL1.contract.WatchLogs(opts, "SumVerificationCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldL1SumVerificationCompleted)
				if err := _HelloWorldL1.contract.UnpackLog(event, "SumVerificationCompleted", log); err != nil {
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

// ParseSumVerificationCompleted is a log parse operation binding the contract event 0xac2034b96316771cf452968ad23adf2a71e4c283d311ad2b66de9b6f7f91a7dd.
//
// Solidity: event SumVerificationCompleted(uint256 indexed taskId, bool isCorrect, uint256 actualResult)
func (_HelloWorldL1 *HelloWorldL1Filterer) ParseSumVerificationCompleted(log types.Log) (*HelloWorldL1SumVerificationCompleted, error) {
	event := new(HelloWorldL1SumVerificationCompleted)
	if err := _HelloWorldL1.contract.UnpackLog(event, "SumVerificationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldL1SumVerificationRequestedIterator is returned from FilterSumVerificationRequested and is used to iterate over the raw logs and unpacked data for SumVerificationRequested events raised by the HelloWorldL1 contract.
type HelloWorldL1SumVerificationRequestedIterator struct {
	Event *HelloWorldL1SumVerificationRequested // Event containing the contract specifics and raw log

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
func (it *HelloWorldL1SumVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldL1SumVerificationRequested)
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
		it.Event = new(HelloWorldL1SumVerificationRequested)
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
func (it *HelloWorldL1SumVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldL1SumVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldL1SumVerificationRequested represents a SumVerificationRequested event raised by the HelloWorldL1 contract.
type HelloWorldL1SumVerificationRequested struct {
	TaskId        *big.Int
	Requester     common.Address
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSumVerificationRequested is a free log retrieval operation binding the contract event 0x74396ec72e1af7caf32d3950fc17c8d8e32f4efe70ee2ec3c327f4091dab1347.
//
// Solidity: event SumVerificationRequested(uint256 indexed taskId, address indexed requester, uint256 a, uint256 b, uint256 claimedResult)
func (_HelloWorldL1 *HelloWorldL1Filterer) FilterSumVerificationRequested(opts *bind.FilterOpts, taskId []*big.Int, requester []common.Address) (*HelloWorldL1SumVerificationRequestedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _HelloWorldL1.contract.FilterLogs(opts, "SumVerificationRequested", taskIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL1SumVerificationRequestedIterator{contract: _HelloWorldL1.contract, event: "SumVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchSumVerificationRequested is a free log subscription operation binding the contract event 0x74396ec72e1af7caf32d3950fc17c8d8e32f4efe70ee2ec3c327f4091dab1347.
//
// Solidity: event SumVerificationRequested(uint256 indexed taskId, address indexed requester, uint256 a, uint256 b, uint256 claimedResult)
func (_HelloWorldL1 *HelloWorldL1Filterer) WatchSumVerificationRequested(opts *bind.WatchOpts, sink chan<- *HelloWorldL1SumVerificationRequested, taskId []*big.Int, requester []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _HelloWorldL1.contract.WatchLogs(opts, "SumVerificationRequested", taskIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldL1SumVerificationRequested)
				if err := _HelloWorldL1.contract.UnpackLog(event, "SumVerificationRequested", log); err != nil {
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

// ParseSumVerificationRequested is a log parse operation binding the contract event 0x74396ec72e1af7caf32d3950fc17c8d8e32f4efe70ee2ec3c327f4091dab1347.
//
// Solidity: event SumVerificationRequested(uint256 indexed taskId, address indexed requester, uint256 a, uint256 b, uint256 claimedResult)
func (_HelloWorldL1 *HelloWorldL1Filterer) ParseSumVerificationRequested(log types.Log) (*HelloWorldL1SumVerificationRequested, error) {
	event := new(HelloWorldL1SumVerificationRequested)
	if err := _HelloWorldL1.contract.UnpackLog(event, "SumVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
