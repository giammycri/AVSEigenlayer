// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworldl2

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

// HelloWorldL2SumVerificationRequest is an auto generated low-level Go binding around an user-defined struct.
type HelloWorldL2SumVerificationRequest struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Timestamp     *big.Int
	Verified      bool
	IsCorrect     bool
}

// HelloWorldL2MetaData contains all meta data concerning the HelloWorldL2 contract.
var HelloWorldL2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskMailbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_avsAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeSumVerification\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executorOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHelloWorldL2.SumVerificationRequest\",\"components\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verified\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTaskVerified\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestSumVerification\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verified\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskMailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifySumOnChain\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"SumVerificationCompleted\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"isCorrect\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"actualResult\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SumVerificationRequested\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"taskId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"a\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"claimedResult\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610a29380380610a2983398101604081905261002e9161011b565b6001600160a01b0383166100895760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964205461736b4d61696c626f782061646472657373000000000060448201526064015b60405180910390fd5b6001600160a01b0382166100df5760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964204156532061646472657373000000000000000000000000006044820152606401610080565b6001600160a01b03928316608052911660a05263ffffffff1660c052610167565b80516001600160a01b0381168114610116575f5ffd5b919050565b5f5f5f6060848603121561012d575f5ffd5b61013684610100565b925061014460208501610100565b9150604084015163ffffffff8116811461015c575f5ffd5b809150509250925092565b60805160a05160c0516108836101a65f395f81816101c9015261047701525f8181610205015261044d01525f818161024401526104c801526108835ff3fe608060405234801561000f575f5ffd5b506004361061009b575f3560e01c80639d866985116100635780639d8669851461012b578063c428b075146101c4578063da324c1314610200578063f42a9e131461023f578063fb1e61ca14610266575f5ffd5b80631480eefd1461009f5780631751637d146100d95780632ee28d8a146100ee578063586717301461010f5780636ef6a57714610118575b5f5ffd5b6100c46100ad36600461070a565b5f9081526020819052604090206005015460ff1690565b60405190151581526020015b60405180910390f35b6100ec6100e7366004610721565b6102d7565b005b6101016100fc366004610753565b6103f3565b6040519081526020016100d0565b61010160015481565b6100c4610126366004610753565b61063b565b61018361013936600461070a565b5f602081905290815260409020805460018201546002830154600384015460048501546005909501549394929391926001600160a01b039091169160ff8082169161010090041687565b604080519788526020880196909652948601939093526001600160a01b0390911660608501526080840152151560a0830152151560c082015260e0016100d0565b6101eb7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100d0565b6102277f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100d0565b6102277f000000000000000000000000000000000000000000000000000000000000000081565b61027961027436600461070a565b610650565b6040516100d091908151815260208083015190820152604080830151908201526060808301516001600160a01b0316908201526080808301519082015260a08083015115159082015260c09182015115159181019190915260e00190565b5f82815260208190526040812060048101549091036103315760405162461bcd60e51b815260206004820152601160248201527014995c5d595cdd081b9bdd08199bdd5b99607a1b60448201526064015b60405180910390fd5b600581015460ff16156103795760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481d995c9a599a595960821b6044820152606401610328565b600581018054600161ffff199091166101008515150217811790915581015481545f916103a591610790565b90505f847ff989d1a1a6bec05f85fc1007c18a198dba77239ccd89846d5f20d270d2f1c25085846040516103e59291909115158252602082015260400190565b60405180910390a350505050565b600180545f9182919082610406836107a9565b90915550604080516020810183905290810187905260608101869052608081018590529091505f9060a00160408051601f1981840301815282820182526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116845263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208581019190915283516060810185525f81529081018590528084018390529251631fb66f5d60e01b81529194507f00000000000000000000000000000000000000000000000000000000000000001690631fb66f5d906104fd9084906004016107c1565b6020604051808303815f875af1158015610519573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061053d9190610836565b6040805160e0810182528a815260208082018b81528284018b815233606080860182815242608088019081525f60a0890181815260c08a018281528c8352828a52918b902099518a55965160018a01559451600289015590516003880180546001600160a01b0319166001600160a01b03909216919091179055516004870155925160059095018054925161ffff1990931695151561ff00191695909517610100921515929092029190911790935583518d81529182018c90529281018a905292975091869188917f5adbe79f92223c99195cea3bc9fe23b34f0ab961000394479a95f8f782ec8e8e910160405180910390a4505050509392505050565b5f816106478486610790565b14949350505050565b6106966040518060e001604052805f81526020015f81526020015f81526020015f6001600160a01b031681526020015f81526020015f151581526020015f151581525090565b505f9081526020818152604091829020825160e0810184528154815260018201549281019290925260028101549282019290925260038201546001600160a01b031660608201526004820154608082015260059091015460ff808216151560a084015261010090910416151560c082015290565b5f6020828403121561071a575f5ffd5b5035919050565b5f5f60408385031215610732575f5ffd5b8235915060208301358015158114610748575f5ffd5b809150509250929050565b5f5f5f60608486031215610765575f5ffd5b505081359360208301359350604090920135919050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156107a3576107a361077c565b92915050565b5f600182016107ba576107ba61077c565b5060010190565b6020815260018060a01b0382511660208201525f602083015160018060a01b03815116604084015263ffffffff602082015116606084015250604083015160808084015280518060a0850152806020830160c086015e5f60c0828601015260c0601f19601f8301168501019250505092915050565b5f60208284031215610846575f5ffd5b505191905056fea26469706673582212203ab42b94aee547a09f3cd5e15d33ca117f2d11c1474c2085683f34940134af2e64736f6c634300081b0033",
}

// HelloWorldL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloWorldL2MetaData.ABI instead.
var HelloWorldL2ABI = HelloWorldL2MetaData.ABI

// HelloWorldL2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloWorldL2MetaData.Bin instead.
var HelloWorldL2Bin = HelloWorldL2MetaData.Bin

// DeployHelloWorldL2 deploys a new Ethereum contract, binding an instance of HelloWorldL2 to it.
func DeployHelloWorldL2(auth *bind.TransactOpts, backend bind.ContractBackend, _taskMailbox common.Address, _avsAddress common.Address, _executorOperatorSetId uint32) (common.Address, *types.Transaction, *HelloWorldL2, error) {
	parsed, err := HelloWorldL2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloWorldL2Bin), backend, _taskMailbox, _avsAddress, _executorOperatorSetId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HelloWorldL2{HelloWorldL2Caller: HelloWorldL2Caller{contract: contract}, HelloWorldL2Transactor: HelloWorldL2Transactor{contract: contract}, HelloWorldL2Filterer: HelloWorldL2Filterer{contract: contract}}, nil
}

// HelloWorldL2 is an auto generated Go binding around an Ethereum contract.
type HelloWorldL2 struct {
	HelloWorldL2Caller     // Read-only binding to the contract
	HelloWorldL2Transactor // Write-only binding to the contract
	HelloWorldL2Filterer   // Log filterer for contract events
}

// HelloWorldL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type HelloWorldL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloWorldL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloWorldL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloWorldL2Session struct {
	Contract     *HelloWorldL2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloWorldL2CallerSession struct {
	Contract *HelloWorldL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HelloWorldL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloWorldL2TransactorSession struct {
	Contract     *HelloWorldL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HelloWorldL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type HelloWorldL2Raw struct {
	Contract *HelloWorldL2 // Generic contract binding to access the raw methods on
}

// HelloWorldL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloWorldL2CallerRaw struct {
	Contract *HelloWorldL2Caller // Generic read-only contract binding to access the raw methods on
}

// HelloWorldL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloWorldL2TransactorRaw struct {
	Contract *HelloWorldL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloWorldL2 creates a new instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2(address common.Address, backend bind.ContractBackend) (*HelloWorldL2, error) {
	contract, err := bindHelloWorldL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2{HelloWorldL2Caller: HelloWorldL2Caller{contract: contract}, HelloWorldL2Transactor: HelloWorldL2Transactor{contract: contract}, HelloWorldL2Filterer: HelloWorldL2Filterer{contract: contract}}, nil
}

// NewHelloWorldL2Caller creates a new read-only instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2Caller(address common.Address, caller bind.ContractCaller) (*HelloWorldL2Caller, error) {
	contract, err := bindHelloWorldL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2Caller{contract: contract}, nil
}

// NewHelloWorldL2Transactor creates a new write-only instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2Transactor(address common.Address, transactor bind.ContractTransactor) (*HelloWorldL2Transactor, error) {
	contract, err := bindHelloWorldL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2Transactor{contract: contract}, nil
}

// NewHelloWorldL2Filterer creates a new log filterer instance of HelloWorldL2, bound to a specific deployed contract.
func NewHelloWorldL2Filterer(address common.Address, filterer bind.ContractFilterer) (*HelloWorldL2Filterer, error) {
	contract, err := bindHelloWorldL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2Filterer{contract: contract}, nil
}

// bindHelloWorldL2 binds a generic wrapper to an already deployed contract.
func bindHelloWorldL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloWorldL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorldL2 *HelloWorldL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorldL2.Contract.HelloWorldL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorldL2 *HelloWorldL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.HelloWorldL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorldL2 *HelloWorldL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.HelloWorldL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorldL2 *HelloWorldL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorldL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorldL2 *HelloWorldL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorldL2 *HelloWorldL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.contract.Transact(opts, method, params...)
}

// AvsAddress is a free data retrieval call binding the contract method 0xda324c13.
//
// Solidity: function avsAddress() view returns(address)
func (_HelloWorldL2 *HelloWorldL2Caller) AvsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "avsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsAddress is a free data retrieval call binding the contract method 0xda324c13.
//
// Solidity: function avsAddress() view returns(address)
func (_HelloWorldL2 *HelloWorldL2Session) AvsAddress() (common.Address, error) {
	return _HelloWorldL2.Contract.AvsAddress(&_HelloWorldL2.CallOpts)
}

// AvsAddress is a free data retrieval call binding the contract method 0xda324c13.
//
// Solidity: function avsAddress() view returns(address)
func (_HelloWorldL2 *HelloWorldL2CallerSession) AvsAddress() (common.Address, error) {
	return _HelloWorldL2.Contract.AvsAddress(&_HelloWorldL2.CallOpts)
}

// ExecutorOperatorSetId is a free data retrieval call binding the contract method 0xc428b075.
//
// Solidity: function executorOperatorSetId() view returns(uint32)
func (_HelloWorldL2 *HelloWorldL2Caller) ExecutorOperatorSetId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "executorOperatorSetId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ExecutorOperatorSetId is a free data retrieval call binding the contract method 0xc428b075.
//
// Solidity: function executorOperatorSetId() view returns(uint32)
func (_HelloWorldL2 *HelloWorldL2Session) ExecutorOperatorSetId() (uint32, error) {
	return _HelloWorldL2.Contract.ExecutorOperatorSetId(&_HelloWorldL2.CallOpts)
}

// ExecutorOperatorSetId is a free data retrieval call binding the contract method 0xc428b075.
//
// Solidity: function executorOperatorSetId() view returns(uint32)
func (_HelloWorldL2 *HelloWorldL2CallerSession) ExecutorOperatorSetId() (uint32, error) {
	return _HelloWorldL2.Contract.ExecutorOperatorSetId(&_HelloWorldL2.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 taskHash) view returns((uint256,uint256,uint256,address,uint256,bool,bool))
func (_HelloWorldL2 *HelloWorldL2Caller) GetRequest(opts *bind.CallOpts, taskHash [32]byte) (HelloWorldL2SumVerificationRequest, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "getRequest", taskHash)

	if err != nil {
		return *new(HelloWorldL2SumVerificationRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(HelloWorldL2SumVerificationRequest)).(*HelloWorldL2SumVerificationRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 taskHash) view returns((uint256,uint256,uint256,address,uint256,bool,bool))
func (_HelloWorldL2 *HelloWorldL2Session) GetRequest(taskHash [32]byte) (HelloWorldL2SumVerificationRequest, error) {
	return _HelloWorldL2.Contract.GetRequest(&_HelloWorldL2.CallOpts, taskHash)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 taskHash) view returns((uint256,uint256,uint256,address,uint256,bool,bool))
func (_HelloWorldL2 *HelloWorldL2CallerSession) GetRequest(taskHash [32]byte) (HelloWorldL2SumVerificationRequest, error) {
	return _HelloWorldL2.Contract.GetRequest(&_HelloWorldL2.CallOpts, taskHash)
}

// IsTaskVerified is a free data retrieval call binding the contract method 0x1480eefd.
//
// Solidity: function isTaskVerified(bytes32 taskHash) view returns(bool)
func (_HelloWorldL2 *HelloWorldL2Caller) IsTaskVerified(opts *bind.CallOpts, taskHash [32]byte) (bool, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "isTaskVerified", taskHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTaskVerified is a free data retrieval call binding the contract method 0x1480eefd.
//
// Solidity: function isTaskVerified(bytes32 taskHash) view returns(bool)
func (_HelloWorldL2 *HelloWorldL2Session) IsTaskVerified(taskHash [32]byte) (bool, error) {
	return _HelloWorldL2.Contract.IsTaskVerified(&_HelloWorldL2.CallOpts, taskHash)
}

// IsTaskVerified is a free data retrieval call binding the contract method 0x1480eefd.
//
// Solidity: function isTaskVerified(bytes32 taskHash) view returns(bool)
func (_HelloWorldL2 *HelloWorldL2CallerSession) IsTaskVerified(taskHash [32]byte) (bool, error) {
	return _HelloWorldL2.Contract.IsTaskVerified(&_HelloWorldL2.CallOpts, taskHash)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 a, uint256 b, uint256 claimedResult, address requester, uint256 timestamp, bool verified, bool isCorrect)
func (_HelloWorldL2 *HelloWorldL2Caller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Timestamp     *big.Int
	Verified      bool
	IsCorrect     bool
}, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		A             *big.Int
		B             *big.Int
		ClaimedResult *big.Int
		Requester     common.Address
		Timestamp     *big.Int
		Verified      bool
		IsCorrect     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimedResult = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Verified = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.IsCorrect = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 a, uint256 b, uint256 claimedResult, address requester, uint256 timestamp, bool verified, bool isCorrect)
func (_HelloWorldL2 *HelloWorldL2Session) Requests(arg0 [32]byte) (struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Timestamp     *big.Int
	Verified      bool
	IsCorrect     bool
}, error) {
	return _HelloWorldL2.Contract.Requests(&_HelloWorldL2.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint256 a, uint256 b, uint256 claimedResult, address requester, uint256 timestamp, bool verified, bool isCorrect)
func (_HelloWorldL2 *HelloWorldL2CallerSession) Requests(arg0 [32]byte) (struct {
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Requester     common.Address
	Timestamp     *big.Int
	Verified      bool
	IsCorrect     bool
}, error) {
	return _HelloWorldL2.Contract.Requests(&_HelloWorldL2.CallOpts, arg0)
}

// TaskCounter is a free data retrieval call binding the contract method 0x58671730.
//
// Solidity: function taskCounter() view returns(uint256)
func (_HelloWorldL2 *HelloWorldL2Caller) TaskCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "taskCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCounter is a free data retrieval call binding the contract method 0x58671730.
//
// Solidity: function taskCounter() view returns(uint256)
func (_HelloWorldL2 *HelloWorldL2Session) TaskCounter() (*big.Int, error) {
	return _HelloWorldL2.Contract.TaskCounter(&_HelloWorldL2.CallOpts)
}

// TaskCounter is a free data retrieval call binding the contract method 0x58671730.
//
// Solidity: function taskCounter() view returns(uint256)
func (_HelloWorldL2 *HelloWorldL2CallerSession) TaskCounter() (*big.Int, error) {
	return _HelloWorldL2.Contract.TaskCounter(&_HelloWorldL2.CallOpts)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_HelloWorldL2 *HelloWorldL2Caller) TaskMailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "taskMailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_HelloWorldL2 *HelloWorldL2Session) TaskMailbox() (common.Address, error) {
	return _HelloWorldL2.Contract.TaskMailbox(&_HelloWorldL2.CallOpts)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_HelloWorldL2 *HelloWorldL2CallerSession) TaskMailbox() (common.Address, error) {
	return _HelloWorldL2.Contract.TaskMailbox(&_HelloWorldL2.CallOpts)
}

// VerifySumOnChain is a free data retrieval call binding the contract method 0x6ef6a577.
//
// Solidity: function verifySumOnChain(uint256 a, uint256 b, uint256 result) pure returns(bool)
func (_HelloWorldL2 *HelloWorldL2Caller) VerifySumOnChain(opts *bind.CallOpts, a *big.Int, b *big.Int, result *big.Int) (bool, error) {
	var out []interface{}
	err := _HelloWorldL2.contract.Call(opts, &out, "verifySumOnChain", a, b, result)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySumOnChain is a free data retrieval call binding the contract method 0x6ef6a577.
//
// Solidity: function verifySumOnChain(uint256 a, uint256 b, uint256 result) pure returns(bool)
func (_HelloWorldL2 *HelloWorldL2Session) VerifySumOnChain(a *big.Int, b *big.Int, result *big.Int) (bool, error) {
	return _HelloWorldL2.Contract.VerifySumOnChain(&_HelloWorldL2.CallOpts, a, b, result)
}

// VerifySumOnChain is a free data retrieval call binding the contract method 0x6ef6a577.
//
// Solidity: function verifySumOnChain(uint256 a, uint256 b, uint256 result) pure returns(bool)
func (_HelloWorldL2 *HelloWorldL2CallerSession) VerifySumOnChain(a *big.Int, b *big.Int, result *big.Int) (bool, error) {
	return _HelloWorldL2.Contract.VerifySumOnChain(&_HelloWorldL2.CallOpts, a, b, result)
}

// CompleteSumVerification is a paid mutator transaction binding the contract method 0x1751637d.
//
// Solidity: function completeSumVerification(bytes32 taskHash, bool isCorrect) returns()
func (_HelloWorldL2 *HelloWorldL2Transactor) CompleteSumVerification(opts *bind.TransactOpts, taskHash [32]byte, isCorrect bool) (*types.Transaction, error) {
	return _HelloWorldL2.contract.Transact(opts, "completeSumVerification", taskHash, isCorrect)
}

// CompleteSumVerification is a paid mutator transaction binding the contract method 0x1751637d.
//
// Solidity: function completeSumVerification(bytes32 taskHash, bool isCorrect) returns()
func (_HelloWorldL2 *HelloWorldL2Session) CompleteSumVerification(taskHash [32]byte, isCorrect bool) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.CompleteSumVerification(&_HelloWorldL2.TransactOpts, taskHash, isCorrect)
}

// CompleteSumVerification is a paid mutator transaction binding the contract method 0x1751637d.
//
// Solidity: function completeSumVerification(bytes32 taskHash, bool isCorrect) returns()
func (_HelloWorldL2 *HelloWorldL2TransactorSession) CompleteSumVerification(taskHash [32]byte, isCorrect bool) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.CompleteSumVerification(&_HelloWorldL2.TransactOpts, taskHash, isCorrect)
}

// RequestSumVerification is a paid mutator transaction binding the contract method 0x2ee28d8a.
//
// Solidity: function requestSumVerification(uint256 a, uint256 b, uint256 claimedResult) returns(bytes32 taskHash)
func (_HelloWorldL2 *HelloWorldL2Transactor) RequestSumVerification(opts *bind.TransactOpts, a *big.Int, b *big.Int, claimedResult *big.Int) (*types.Transaction, error) {
	return _HelloWorldL2.contract.Transact(opts, "requestSumVerification", a, b, claimedResult)
}

// RequestSumVerification is a paid mutator transaction binding the contract method 0x2ee28d8a.
//
// Solidity: function requestSumVerification(uint256 a, uint256 b, uint256 claimedResult) returns(bytes32 taskHash)
func (_HelloWorldL2 *HelloWorldL2Session) RequestSumVerification(a *big.Int, b *big.Int, claimedResult *big.Int) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.RequestSumVerification(&_HelloWorldL2.TransactOpts, a, b, claimedResult)
}

// RequestSumVerification is a paid mutator transaction binding the contract method 0x2ee28d8a.
//
// Solidity: function requestSumVerification(uint256 a, uint256 b, uint256 claimedResult) returns(bytes32 taskHash)
func (_HelloWorldL2 *HelloWorldL2TransactorSession) RequestSumVerification(a *big.Int, b *big.Int, claimedResult *big.Int) (*types.Transaction, error) {
	return _HelloWorldL2.Contract.RequestSumVerification(&_HelloWorldL2.TransactOpts, a, b, claimedResult)
}

// HelloWorldL2SumVerificationCompletedIterator is returned from FilterSumVerificationCompleted and is used to iterate over the raw logs and unpacked data for SumVerificationCompleted events raised by the HelloWorldL2 contract.
type HelloWorldL2SumVerificationCompletedIterator struct {
	Event *HelloWorldL2SumVerificationCompleted // Event containing the contract specifics and raw log

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
func (it *HelloWorldL2SumVerificationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldL2SumVerificationCompleted)
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
		it.Event = new(HelloWorldL2SumVerificationCompleted)
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
func (it *HelloWorldL2SumVerificationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldL2SumVerificationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldL2SumVerificationCompleted represents a SumVerificationCompleted event raised by the HelloWorldL2 contract.
type HelloWorldL2SumVerificationCompleted struct {
	TaskHash     [32]byte
	TaskId       *big.Int
	IsCorrect    bool
	ActualResult *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSumVerificationCompleted is a free log retrieval operation binding the contract event 0xf989d1a1a6bec05f85fc1007c18a198dba77239ccd89846d5f20d270d2f1c250.
//
// Solidity: event SumVerificationCompleted(bytes32 indexed taskHash, uint256 indexed taskId, bool isCorrect, uint256 actualResult)
func (_HelloWorldL2 *HelloWorldL2Filterer) FilterSumVerificationCompleted(opts *bind.FilterOpts, taskHash [][32]byte, taskId []*big.Int) (*HelloWorldL2SumVerificationCompletedIterator, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _HelloWorldL2.contract.FilterLogs(opts, "SumVerificationCompleted", taskHashRule, taskIdRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2SumVerificationCompletedIterator{contract: _HelloWorldL2.contract, event: "SumVerificationCompleted", logs: logs, sub: sub}, nil
}

// WatchSumVerificationCompleted is a free log subscription operation binding the contract event 0xf989d1a1a6bec05f85fc1007c18a198dba77239ccd89846d5f20d270d2f1c250.
//
// Solidity: event SumVerificationCompleted(bytes32 indexed taskHash, uint256 indexed taskId, bool isCorrect, uint256 actualResult)
func (_HelloWorldL2 *HelloWorldL2Filterer) WatchSumVerificationCompleted(opts *bind.WatchOpts, sink chan<- *HelloWorldL2SumVerificationCompleted, taskHash [][32]byte, taskId []*big.Int) (event.Subscription, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _HelloWorldL2.contract.WatchLogs(opts, "SumVerificationCompleted", taskHashRule, taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldL2SumVerificationCompleted)
				if err := _HelloWorldL2.contract.UnpackLog(event, "SumVerificationCompleted", log); err != nil {
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

// ParseSumVerificationCompleted is a log parse operation binding the contract event 0xf989d1a1a6bec05f85fc1007c18a198dba77239ccd89846d5f20d270d2f1c250.
//
// Solidity: event SumVerificationCompleted(bytes32 indexed taskHash, uint256 indexed taskId, bool isCorrect, uint256 actualResult)
func (_HelloWorldL2 *HelloWorldL2Filterer) ParseSumVerificationCompleted(log types.Log) (*HelloWorldL2SumVerificationCompleted, error) {
	event := new(HelloWorldL2SumVerificationCompleted)
	if err := _HelloWorldL2.contract.UnpackLog(event, "SumVerificationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldL2SumVerificationRequestedIterator is returned from FilterSumVerificationRequested and is used to iterate over the raw logs and unpacked data for SumVerificationRequested events raised by the HelloWorldL2 contract.
type HelloWorldL2SumVerificationRequestedIterator struct {
	Event *HelloWorldL2SumVerificationRequested // Event containing the contract specifics and raw log

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
func (it *HelloWorldL2SumVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldL2SumVerificationRequested)
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
		it.Event = new(HelloWorldL2SumVerificationRequested)
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
func (it *HelloWorldL2SumVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldL2SumVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldL2SumVerificationRequested represents a SumVerificationRequested event raised by the HelloWorldL2 contract.
type HelloWorldL2SumVerificationRequested struct {
	TaskHash      [32]byte
	TaskId        *big.Int
	Requester     common.Address
	A             *big.Int
	B             *big.Int
	ClaimedResult *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSumVerificationRequested is a free log retrieval operation binding the contract event 0x5adbe79f92223c99195cea3bc9fe23b34f0ab961000394479a95f8f782ec8e8e.
//
// Solidity: event SumVerificationRequested(bytes32 indexed taskHash, uint256 indexed taskId, address indexed requester, uint256 a, uint256 b, uint256 claimedResult)
func (_HelloWorldL2 *HelloWorldL2Filterer) FilterSumVerificationRequested(opts *bind.FilterOpts, taskHash [][32]byte, taskId []*big.Int, requester []common.Address) (*HelloWorldL2SumVerificationRequestedIterator, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _HelloWorldL2.contract.FilterLogs(opts, "SumVerificationRequested", taskHashRule, taskIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldL2SumVerificationRequestedIterator{contract: _HelloWorldL2.contract, event: "SumVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchSumVerificationRequested is a free log subscription operation binding the contract event 0x5adbe79f92223c99195cea3bc9fe23b34f0ab961000394479a95f8f782ec8e8e.
//
// Solidity: event SumVerificationRequested(bytes32 indexed taskHash, uint256 indexed taskId, address indexed requester, uint256 a, uint256 b, uint256 claimedResult)
func (_HelloWorldL2 *HelloWorldL2Filterer) WatchSumVerificationRequested(opts *bind.WatchOpts, sink chan<- *HelloWorldL2SumVerificationRequested, taskHash [][32]byte, taskId []*big.Int, requester []common.Address) (event.Subscription, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _HelloWorldL2.contract.WatchLogs(opts, "SumVerificationRequested", taskHashRule, taskIdRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldL2SumVerificationRequested)
				if err := _HelloWorldL2.contract.UnpackLog(event, "SumVerificationRequested", log); err != nil {
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

// ParseSumVerificationRequested is a log parse operation binding the contract event 0x5adbe79f92223c99195cea3bc9fe23b34f0ab961000394479a95f8f782ec8e8e.
//
// Solidity: event SumVerificationRequested(bytes32 indexed taskHash, uint256 indexed taskId, address indexed requester, uint256 a, uint256 b, uint256 claimedResult)
func (_HelloWorldL2 *HelloWorldL2Filterer) ParseSumVerificationRequested(log types.Log) (*HelloWorldL2SumVerificationRequested, error) {
	event := new(HelloWorldL2SumVerificationRequested)
	if err := _HelloWorldL2.contract.UnpackLog(event, "SumVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
