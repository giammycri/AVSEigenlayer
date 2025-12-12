// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avstaskhook

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// TaskParams is an auto generated low-level Go binding around an user-defined struct.
type TaskParams struct {
	RefundCollector     common.Address
	ExecutorOperatorSet OperatorSet
	Payload             []byte
}

// AVSTaskHookMetaData contains all meta data concerning the AVSTaskHook contract.
var AVSTaskHookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateTaskFee\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structTaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleCallback\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handlePostTaskCreation\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handlePostTaskResultSubmission\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTaskResultCorrect\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCompleted\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskResults\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatePostTaskExecution\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validatePreTaskCreation\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structTaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validatePreTaskResultSubmission\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executorCert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"TaskCallbackReceived\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b506108e88061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c806362fee0371161006e57806362fee037146101395780636adf155c1461014c578063ba33565d1461015f578063bc20a46714610175578063e06cd27e14610188578063e082467e146101b8575f5ffd5b806309c5c450146100aa578063154f2c71146100bd57806330995529146100f45780633a0e94f61461010757806351fe309814610127575b5f5ffd5b6100bb6100b83660046104e5565b50565b005b6100df6100cb3660046104e5565b60016020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6100df6101023660046104e5565b6101c6565b61011a6101153660046104e5565b6102ea565b6040516100eb91906104fc565b6100bb610135366004610562565b5050565b61011a6101473660046104e5565b610381565b6100bb61015a3660046105f2565b505050565b6100bb61016d36600461063a565b505050505050565b6100bb6101833660046105f2565b610471565b61019b6101963660046106c3565b505f90565b6040516bffffffffffffffffffffffff90911681526020016100eb565b6100bb6101353660046106fd565b5f8181526001602052604081205460ff1661021d5760405162461bcd60e51b815260206004820152601260248201527115185cdac81b9bdd0818dbdb5c1b195d195960721b60448201526064015b60405180910390fd5b5f828152602081905260408120805461023590610725565b80601f016020809104026020016040519081016040528092919081815260200182805461026190610725565b80156102ac5780601f10610283576101008083540402835291602001916102ac565b820191905f5260205f20905b81548152906001019060200180831161028f57829003601f168201915b5050505050905060208151106102e25780601f815181106102cf576102cf610757565b60209101015160f81c6001149392505050565b505f92915050565b5f602081905290815260409020805461030290610725565b80601f016020809104026020016040519081016040528092919081815260200182805461032e90610725565b80156103795780601f1061035057610100808354040283529160200191610379565b820191905f5260205f20905b81548152906001019060200180831161035c57829003601f168201915b505050505081565b5f8181526001602052604090205460609060ff166103d65760405162461bcd60e51b815260206004820152601260248201527115185cdac81b9bdd0818dbdb5c1b195d195960721b6044820152606401610214565b5f82815260208190526040902080546103ee90610725565b80601f016020809104026020016040519081016040528092919081815260200182805461041a90610725565b80156104655780601f1061043c57610100808354040283529160200191610465565b820191905f5260205f20905b81548152906001019060200180831161044857829003601f168201915b50505050509050919050565b5f8381526020819052604090206104898284836107ca565b505f83815260016020819052604091829020805460ff191690911790555183907f8a76a391e5963534dd00fe295c579d9add687337d9e909b9eab5eb4196a0e888906104d89085908590610884565b60405180910390a2505050565b5f602082840312156104f5575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610547575f5ffd5b919050565b5f6080828403121561055c575f5ffd5b50919050565b5f5f60408385031215610573575f5ffd5b61057c83610531565b9150602083013567ffffffffffffffff811115610597575f5ffd5b6105a38582860161054c565b9150509250929050565b5f5f83601f8401126105bd575f5ffd5b50813567ffffffffffffffff8111156105d4575f5ffd5b6020830191508360208285010111156105eb575f5ffd5b9250929050565b5f5f5f60408486031215610604575f5ffd5b83359250602084013567ffffffffffffffff811115610621575f5ffd5b61062d868287016105ad565b9497909650939450505050565b5f5f5f5f5f5f6080878903121561064f575f5ffd5b61065887610531565b955060208701359450604087013567ffffffffffffffff81111561067a575f5ffd5b61068689828a016105ad565b909550935050606087013567ffffffffffffffff8111156106a5575f5ffd5b6106b189828a016105ad565b979a9699509497509295939492505050565b5f602082840312156106d3575f5ffd5b813567ffffffffffffffff8111156106e9575f5ffd5b6106f58482850161054c565b949350505050565b5f5f6040838503121561070e575f5ffd5b61071783610531565b946020939093013593505050565b600181811c9082168061073957607f821691505b60208210810361055c57634e487b7160e01b5f52602260045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b601f82111561015a57805f5260205f20601f840160051c810160208510156107a45750805b601f840160051c820191505b818110156107c3575f81556001016107b0565b5050505050565b67ffffffffffffffff8311156107e2576107e261076b565b6107f6836107f08354610725565b8361077f565b5f601f841160018114610827575f85156108105750838201355b5f19600387901b1c1916600186901b1783556107c3565b5f83815260208120601f198716915b828110156108565786850135825560209485019460019092019101610836565b5086821015610872575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea26469706673582212207737a04959e1378888db7fe15e6580cfc4b6167b437070b491c1cb6613d8585f64736f6c634300081b0033",
}

// AVSTaskHookABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSTaskHookMetaData.ABI instead.
var AVSTaskHookABI = AVSTaskHookMetaData.ABI

// AVSTaskHookBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AVSTaskHookMetaData.Bin instead.
var AVSTaskHookBin = AVSTaskHookMetaData.Bin

// DeployAVSTaskHook deploys a new Ethereum contract, binding an instance of AVSTaskHook to it.
func DeployAVSTaskHook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AVSTaskHook, error) {
	parsed, err := AVSTaskHookMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AVSTaskHookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AVSTaskHook{AVSTaskHookCaller: AVSTaskHookCaller{contract: contract}, AVSTaskHookTransactor: AVSTaskHookTransactor{contract: contract}, AVSTaskHookFilterer: AVSTaskHookFilterer{contract: contract}}, nil
}

// AVSTaskHook is an auto generated Go binding around an Ethereum contract.
type AVSTaskHook struct {
	AVSTaskHookCaller     // Read-only binding to the contract
	AVSTaskHookTransactor // Write-only binding to the contract
	AVSTaskHookFilterer   // Log filterer for contract events
}

// AVSTaskHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSTaskHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSTaskHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSTaskHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSTaskHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSTaskHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSTaskHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSTaskHookSession struct {
	Contract     *AVSTaskHook      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSTaskHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSTaskHookCallerSession struct {
	Contract *AVSTaskHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AVSTaskHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSTaskHookTransactorSession struct {
	Contract     *AVSTaskHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AVSTaskHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSTaskHookRaw struct {
	Contract *AVSTaskHook // Generic contract binding to access the raw methods on
}

// AVSTaskHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSTaskHookCallerRaw struct {
	Contract *AVSTaskHookCaller // Generic read-only contract binding to access the raw methods on
}

// AVSTaskHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSTaskHookTransactorRaw struct {
	Contract *AVSTaskHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSTaskHook creates a new instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHook(address common.Address, backend bind.ContractBackend) (*AVSTaskHook, error) {
	contract, err := bindAVSTaskHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHook{AVSTaskHookCaller: AVSTaskHookCaller{contract: contract}, AVSTaskHookTransactor: AVSTaskHookTransactor{contract: contract}, AVSTaskHookFilterer: AVSTaskHookFilterer{contract: contract}}, nil
}

// NewAVSTaskHookCaller creates a new read-only instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHookCaller(address common.Address, caller bind.ContractCaller) (*AVSTaskHookCaller, error) {
	contract, err := bindAVSTaskHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookCaller{contract: contract}, nil
}

// NewAVSTaskHookTransactor creates a new write-only instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHookTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSTaskHookTransactor, error) {
	contract, err := bindAVSTaskHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookTransactor{contract: contract}, nil
}

// NewAVSTaskHookFilterer creates a new log filterer instance of AVSTaskHook, bound to a specific deployed contract.
func NewAVSTaskHookFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSTaskHookFilterer, error) {
	contract, err := bindAVSTaskHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookFilterer{contract: contract}, nil
}

// bindAVSTaskHook binds a generic wrapper to an already deployed contract.
func bindAVSTaskHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSTaskHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSTaskHook *AVSTaskHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSTaskHook.Contract.AVSTaskHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSTaskHook *AVSTaskHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.AVSTaskHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSTaskHook *AVSTaskHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.AVSTaskHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSTaskHook *AVSTaskHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSTaskHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSTaskHook *AVSTaskHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSTaskHook *AVSTaskHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.contract.Transact(opts, method, params...)
}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) taskParams) pure returns(uint96 fee)
func (_AVSTaskHook *AVSTaskHookCaller) CalculateTaskFee(opts *bind.CallOpts, taskParams TaskParams) (*big.Int, error) {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "calculateTaskFee", taskParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) taskParams) pure returns(uint96 fee)
func (_AVSTaskHook *AVSTaskHookSession) CalculateTaskFee(taskParams TaskParams) (*big.Int, error) {
	return _AVSTaskHook.Contract.CalculateTaskFee(&_AVSTaskHook.CallOpts, taskParams)
}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) taskParams) pure returns(uint96 fee)
func (_AVSTaskHook *AVSTaskHookCallerSession) CalculateTaskFee(taskParams TaskParams) (*big.Int, error) {
	return _AVSTaskHook.Contract.CalculateTaskFee(&_AVSTaskHook.CallOpts, taskParams)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_AVSTaskHook *AVSTaskHookCaller) GetTaskResult(opts *bind.CallOpts, taskHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "getTaskResult", taskHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_AVSTaskHook *AVSTaskHookSession) GetTaskResult(taskHash [32]byte) ([]byte, error) {
	return _AVSTaskHook.Contract.GetTaskResult(&_AVSTaskHook.CallOpts, taskHash)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x62fee037.
//
// Solidity: function getTaskResult(bytes32 taskHash) view returns(bytes)
func (_AVSTaskHook *AVSTaskHookCallerSession) GetTaskResult(taskHash [32]byte) ([]byte, error) {
	return _AVSTaskHook.Contract.GetTaskResult(&_AVSTaskHook.CallOpts, taskHash)
}

// IsTaskResultCorrect is a free data retrieval call binding the contract method 0x30995529.
//
// Solidity: function isTaskResultCorrect(bytes32 taskHash) view returns(bool)
func (_AVSTaskHook *AVSTaskHookCaller) IsTaskResultCorrect(opts *bind.CallOpts, taskHash [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "isTaskResultCorrect", taskHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTaskResultCorrect is a free data retrieval call binding the contract method 0x30995529.
//
// Solidity: function isTaskResultCorrect(bytes32 taskHash) view returns(bool)
func (_AVSTaskHook *AVSTaskHookSession) IsTaskResultCorrect(taskHash [32]byte) (bool, error) {
	return _AVSTaskHook.Contract.IsTaskResultCorrect(&_AVSTaskHook.CallOpts, taskHash)
}

// IsTaskResultCorrect is a free data retrieval call binding the contract method 0x30995529.
//
// Solidity: function isTaskResultCorrect(bytes32 taskHash) view returns(bool)
func (_AVSTaskHook *AVSTaskHookCallerSession) IsTaskResultCorrect(taskHash [32]byte) (bool, error) {
	return _AVSTaskHook.Contract.IsTaskResultCorrect(&_AVSTaskHook.CallOpts, taskHash)
}

// TaskCompleted is a free data retrieval call binding the contract method 0x154f2c71.
//
// Solidity: function taskCompleted(bytes32 ) view returns(bool)
func (_AVSTaskHook *AVSTaskHookCaller) TaskCompleted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "taskCompleted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskCompleted is a free data retrieval call binding the contract method 0x154f2c71.
//
// Solidity: function taskCompleted(bytes32 ) view returns(bool)
func (_AVSTaskHook *AVSTaskHookSession) TaskCompleted(arg0 [32]byte) (bool, error) {
	return _AVSTaskHook.Contract.TaskCompleted(&_AVSTaskHook.CallOpts, arg0)
}

// TaskCompleted is a free data retrieval call binding the contract method 0x154f2c71.
//
// Solidity: function taskCompleted(bytes32 ) view returns(bool)
func (_AVSTaskHook *AVSTaskHookCallerSession) TaskCompleted(arg0 [32]byte) (bool, error) {
	return _AVSTaskHook.Contract.TaskCompleted(&_AVSTaskHook.CallOpts, arg0)
}

// TaskResults is a free data retrieval call binding the contract method 0x3a0e94f6.
//
// Solidity: function taskResults(bytes32 ) view returns(bytes)
func (_AVSTaskHook *AVSTaskHookCaller) TaskResults(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "taskResults", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TaskResults is a free data retrieval call binding the contract method 0x3a0e94f6.
//
// Solidity: function taskResults(bytes32 ) view returns(bytes)
func (_AVSTaskHook *AVSTaskHookSession) TaskResults(arg0 [32]byte) ([]byte, error) {
	return _AVSTaskHook.Contract.TaskResults(&_AVSTaskHook.CallOpts, arg0)
}

// TaskResults is a free data retrieval call binding the contract method 0x3a0e94f6.
//
// Solidity: function taskResults(bytes32 ) view returns(bytes)
func (_AVSTaskHook *AVSTaskHookCallerSession) TaskResults(arg0 [32]byte) ([]byte, error) {
	return _AVSTaskHook.Contract.TaskResults(&_AVSTaskHook.CallOpts, arg0)
}

// ValidatePostTaskExecution is a free data retrieval call binding the contract method 0x6adf155c.
//
// Solidity: function validatePostTaskExecution(bytes32 taskHash, bytes result) pure returns()
func (_AVSTaskHook *AVSTaskHookCaller) ValidatePostTaskExecution(opts *bind.CallOpts, taskHash [32]byte, result []byte) error {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "validatePostTaskExecution", taskHash, result)

	if err != nil {
		return err
	}

	return err

}

// ValidatePostTaskExecution is a free data retrieval call binding the contract method 0x6adf155c.
//
// Solidity: function validatePostTaskExecution(bytes32 taskHash, bytes result) pure returns()
func (_AVSTaskHook *AVSTaskHookSession) ValidatePostTaskExecution(taskHash [32]byte, result []byte) error {
	return _AVSTaskHook.Contract.ValidatePostTaskExecution(&_AVSTaskHook.CallOpts, taskHash, result)
}

// ValidatePostTaskExecution is a free data retrieval call binding the contract method 0x6adf155c.
//
// Solidity: function validatePostTaskExecution(bytes32 taskHash, bytes result) pure returns()
func (_AVSTaskHook *AVSTaskHookCallerSession) ValidatePostTaskExecution(taskHash [32]byte, result []byte) error {
	return _AVSTaskHook.Contract.ValidatePostTaskExecution(&_AVSTaskHook.CallOpts, taskHash, result)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address creator, (address,(address,uint32),bytes) taskParams) pure returns()
func (_AVSTaskHook *AVSTaskHookCaller) ValidatePreTaskCreation(opts *bind.CallOpts, creator common.Address, taskParams TaskParams) error {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "validatePreTaskCreation", creator, taskParams)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address creator, (address,(address,uint32),bytes) taskParams) pure returns()
func (_AVSTaskHook *AVSTaskHookSession) ValidatePreTaskCreation(creator common.Address, taskParams TaskParams) error {
	return _AVSTaskHook.Contract.ValidatePreTaskCreation(&_AVSTaskHook.CallOpts, creator, taskParams)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address creator, (address,(address,uint32),bytes) taskParams) pure returns()
func (_AVSTaskHook *AVSTaskHookCallerSession) ValidatePreTaskCreation(creator common.Address, taskParams TaskParams) error {
	return _AVSTaskHook.Contract.ValidatePreTaskCreation(&_AVSTaskHook.CallOpts, creator, taskParams)
}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address sender, bytes32 taskHash, bytes executorCert, bytes result) pure returns()
func (_AVSTaskHook *AVSTaskHookCaller) ValidatePreTaskResultSubmission(opts *bind.CallOpts, sender common.Address, taskHash [32]byte, executorCert []byte, result []byte) error {
	var out []interface{}
	err := _AVSTaskHook.contract.Call(opts, &out, "validatePreTaskResultSubmission", sender, taskHash, executorCert, result)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address sender, bytes32 taskHash, bytes executorCert, bytes result) pure returns()
func (_AVSTaskHook *AVSTaskHookSession) ValidatePreTaskResultSubmission(sender common.Address, taskHash [32]byte, executorCert []byte, result []byte) error {
	return _AVSTaskHook.Contract.ValidatePreTaskResultSubmission(&_AVSTaskHook.CallOpts, sender, taskHash, executorCert, result)
}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address sender, bytes32 taskHash, bytes executorCert, bytes result) pure returns()
func (_AVSTaskHook *AVSTaskHookCallerSession) ValidatePreTaskResultSubmission(sender common.Address, taskHash [32]byte, executorCert []byte, result []byte) error {
	return _AVSTaskHook.Contract.ValidatePreTaskResultSubmission(&_AVSTaskHook.CallOpts, sender, taskHash, executorCert, result)
}

// HandleCallback is a paid mutator transaction binding the contract method 0xbc20a467.
//
// Solidity: function handleCallback(bytes32 taskHash, bytes result) returns()
func (_AVSTaskHook *AVSTaskHookTransactor) HandleCallback(opts *bind.TransactOpts, taskHash [32]byte, result []byte) (*types.Transaction, error) {
	return _AVSTaskHook.contract.Transact(opts, "handleCallback", taskHash, result)
}

// HandleCallback is a paid mutator transaction binding the contract method 0xbc20a467.
//
// Solidity: function handleCallback(bytes32 taskHash, bytes result) returns()
func (_AVSTaskHook *AVSTaskHookSession) HandleCallback(taskHash [32]byte, result []byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandleCallback(&_AVSTaskHook.TransactOpts, taskHash, result)
}

// HandleCallback is a paid mutator transaction binding the contract method 0xbc20a467.
//
// Solidity: function handleCallback(bytes32 taskHash, bytes result) returns()
func (_AVSTaskHook *AVSTaskHookTransactorSession) HandleCallback(taskHash [32]byte, result []byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandleCallback(&_AVSTaskHook.TransactOpts, taskHash, result)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 taskHash) returns()
func (_AVSTaskHook *AVSTaskHookTransactor) HandlePostTaskCreation(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.contract.Transact(opts, "handlePostTaskCreation", taskHash)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 taskHash) returns()
func (_AVSTaskHook *AVSTaskHookSession) HandlePostTaskCreation(taskHash [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskCreation(&_AVSTaskHook.TransactOpts, taskHash)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 taskHash) returns()
func (_AVSTaskHook *AVSTaskHookTransactorSession) HandlePostTaskCreation(taskHash [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskCreation(&_AVSTaskHook.TransactOpts, taskHash)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xe082467e.
//
// Solidity: function handlePostTaskResultSubmission(address sender, bytes32 taskHash) returns()
func (_AVSTaskHook *AVSTaskHookTransactor) HandlePostTaskResultSubmission(opts *bind.TransactOpts, sender common.Address, taskHash [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.contract.Transact(opts, "handlePostTaskResultSubmission", sender, taskHash)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xe082467e.
//
// Solidity: function handlePostTaskResultSubmission(address sender, bytes32 taskHash) returns()
func (_AVSTaskHook *AVSTaskHookSession) HandlePostTaskResultSubmission(sender common.Address, taskHash [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskResultSubmission(&_AVSTaskHook.TransactOpts, sender, taskHash)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xe082467e.
//
// Solidity: function handlePostTaskResultSubmission(address sender, bytes32 taskHash) returns()
func (_AVSTaskHook *AVSTaskHookTransactorSession) HandlePostTaskResultSubmission(sender common.Address, taskHash [32]byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HandlePostTaskResultSubmission(&_AVSTaskHook.TransactOpts, sender, taskHash)
}

// AVSTaskHookTaskCallbackReceivedIterator is returned from FilterTaskCallbackReceived and is used to iterate over the raw logs and unpacked data for TaskCallbackReceived events raised by the AVSTaskHook contract.
type AVSTaskHookTaskCallbackReceivedIterator struct {
	Event *AVSTaskHookTaskCallbackReceived // Event containing the contract specifics and raw log

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
func (it *AVSTaskHookTaskCallbackReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSTaskHookTaskCallbackReceived)
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
		it.Event = new(AVSTaskHookTaskCallbackReceived)
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
func (it *AVSTaskHookTaskCallbackReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSTaskHookTaskCallbackReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSTaskHookTaskCallbackReceived represents a TaskCallbackReceived event raised by the AVSTaskHook contract.
type AVSTaskHookTaskCallbackReceived struct {
	TaskHash [32]byte
	Result   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskCallbackReceived is a free log retrieval operation binding the contract event 0x8a76a391e5963534dd00fe295c579d9add687337d9e909b9eab5eb4196a0e888.
//
// Solidity: event TaskCallbackReceived(bytes32 indexed taskHash, bytes result)
func (_AVSTaskHook *AVSTaskHookFilterer) FilterTaskCallbackReceived(opts *bind.FilterOpts, taskHash [][32]byte) (*AVSTaskHookTaskCallbackReceivedIterator, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _AVSTaskHook.contract.FilterLogs(opts, "TaskCallbackReceived", taskHashRule)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookTaskCallbackReceivedIterator{contract: _AVSTaskHook.contract, event: "TaskCallbackReceived", logs: logs, sub: sub}, nil
}

// WatchTaskCallbackReceived is a free log subscription operation binding the contract event 0x8a76a391e5963534dd00fe295c579d9add687337d9e909b9eab5eb4196a0e888.
//
// Solidity: event TaskCallbackReceived(bytes32 indexed taskHash, bytes result)
func (_AVSTaskHook *AVSTaskHookFilterer) WatchTaskCallbackReceived(opts *bind.WatchOpts, sink chan<- *AVSTaskHookTaskCallbackReceived, taskHash [][32]byte) (event.Subscription, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _AVSTaskHook.contract.WatchLogs(opts, "TaskCallbackReceived", taskHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSTaskHookTaskCallbackReceived)
				if err := _AVSTaskHook.contract.UnpackLog(event, "TaskCallbackReceived", log); err != nil {
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

// ParseTaskCallbackReceived is a log parse operation binding the contract event 0x8a76a391e5963534dd00fe295c579d9add687337d9e909b9eab5eb4196a0e888.
//
// Solidity: event TaskCallbackReceived(bytes32 indexed taskHash, bytes result)
func (_AVSTaskHook *AVSTaskHookFilterer) ParseTaskCallbackReceived(log types.Log) (*AVSTaskHookTaskCallbackReceived, error) {
	event := new(AVSTaskHookTaskCallbackReceived)
	if err := _AVSTaskHook.contract.UnpackLog(event, "TaskCallbackReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
