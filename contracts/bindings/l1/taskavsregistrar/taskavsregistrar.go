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

// TaskAVSRegistrarAvsConfig is an auto generated low-level Go binding around an user-defined struct.
type TaskAVSRegistrarAvsConfig struct {
	AggregatorOperatorSetId uint32
	ExecutorOperatorSetIds  []uint32
}

// TaskAVSRegistrarMetaData contains all meta data concerning the TaskAVSRegistrar contract.
var TaskAVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorToAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregatorOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executorOperatorSetIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvsConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structTaskAVSRegistrar.AvsConfig\",\"components\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOperatorAllowlisted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"operatorSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredOperatorsList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeOperatorFromAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAvsConfig\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structTaskAVSRegistrar.AvsConfig\",\"components\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSInitialized\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AvsConfigSet\",\"inputs\":[{\"name\":\"aggregatorOperatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"executorOperatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAllowlisted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611d02380380611d0283398101604081905261002e9161022d565b6001600160a01b0383166100895760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420416c6c6f636174696f6e4d616e616765720000000000000060448201526064015b60405180910390fd5b6001600160a01b0382166100df5760405162461bcd60e51b815260206004820152601460248201527f496e76616c6964204b65795265676973747261720000000000000000000000006044820152606401610080565b6001600160a01b0381166101355760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964205065726d697373696f6e436f6e74726f6c6c6572000000006044820152606401610080565b6001600160a01b0380841660805282811660a052811660c05261015661015e565b505050610277565b5f54610100900460ff16156101c55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610080565b5f5460ff90811614610214575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461022a575f5ffd5b50565b5f5f5f6060848603121561023f575f5ffd5b835161024a81610216565b602085015190935061025b81610216565b604085015190925061026c81610216565b809150509250925092565b60805160a05160c051611a616102a15f395f61022601525f6101d201525f61035e0152611a615ff3fe608060405234801561000f575f5ffd5b5060043610610148575f3560e01c80638481931d116100bf578063c63fd50211610079578063c63fd50214610346578063ca8aa7c714610359578063d1f2e81d14610380578063da324c1314610393578063f2fde38b146103a6578063fffc9182146103b9575f5ffd5b80638481931d1461029e5780638b2942d4146102be5780638b8738a0146102d55780638da5cb5b14610300578063adc8ea7414610311578063b526578714610324575f5ffd5b80633ec45c7e116101105780633ec45c7e146101cd57806341f548f01461020c5780634657e26a14610221578063485cc955146102485780636b1906f81461025b578063715018a614610296575f5ffd5b806317aef8751461014c5780631fc3727814610161578063241217fa1461017f578063303ca956146101925780633dc70058146101a5575b5f5ffd5b61015f61015a36600461111d565b6103cc565b005b6101696103e1565b604051610176919061113d565b60405180910390f35b61015f61018d36600461111d565b610554565b61015f6101a03660046111cf565b610567565b6101b86101b336600461122b565b61065a565b60405163ffffffff9091168152602001610176565b6101f47f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610176565b610214610691565b6040516101769190611242565b6101f47f000000000000000000000000000000000000000000000000000000000000000081565b61015f6102563660046112a8565b610739565b61028661026936600461111d565b6001600160a01b03165f9081526067602052604090205460ff1690565b6040519015158152602001610176565b61015f6109d1565b6102b16102ac36600461111d565b6109e4565b60405161017691906112d9565b6065546101b890600160a01b900463ffffffff1681565b6102866102e336600461111d565b6001600160a01b03165f9081526068602052604090205460ff1690565b6033546001600160a01b03166101f4565b61015f61031f3660046113db565b610a8d565b61028661033236600461111d565b6065546001600160a01b0391821691161490565b61015f610354366004611438565b610b63565b6101f47f000000000000000000000000000000000000000000000000000000000000000081565b61015f61038e3660046114e6565b610d51565b6065546101f4906001600160a01b031681565b61015f6103b436600461111d565b610eaa565b6101f46103c736600461122b565b610f20565b6103d4610f48565b6103de815f610fa2565b50565b60605f805b6069548110156104485760675f60698381548110610406576104066115d9565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205460ff1615610440578161043c816115ed565b9250505b6001016103e6565b505f816001600160401b038111156104625761046261130e565b60405190808252806020026020018201604052801561048b578160200160208202803683370190505b5090505f805b60695481101561054b5760675f606983815481106104b1576104b16115d9565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205460ff161561054357606981815481106104f0576104f06115d9565b905f5260205f20015f9054906101000a90046001600160a01b031683838151811061051d5761051d6115d9565b6001600160a01b03909216602092830291909101909101528161053f816115ed565b9250505b600101610491565b50909392505050565b61055c610f48565b6103de816001610fa2565b6065546001600160a01b038481169116146105b75760405162461bcd60e51b815260206004820152600b60248201526a496e76616c69642041565360a81b60448201526064015b60405180910390fd5b6001600160a01b0384165f9081526067602052604090205460ff1661060f5760405162461bcd60e51b815260206004820152600e60248201526d139bdd081c9959da5cdd195c995960921b60448201526064016105ae565b6001600160a01b0384165f81815260676020526040808220805460ff19169055517f6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d9190a250505050565b60668181548110610669575f80fd5b905f5260205f209060089182820401919006600402915054906101000a900463ffffffff1681565b604080518082019091525f815260606020820152606554600160a01b900463ffffffff168152606680546040805160208084028201810190925282815292919083018282801561072957602002820191905f5260205f20905f905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116106ec5790505b5050505050816020018190525090565b5f54610100900460ff161580801561075757505f54600160ff909116105b806107705750303b15801561077057505f5460ff166001145b6107d35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105ae565b5f805460ff1916600117905580156107f4575f805461ff0019166101001790555b6001600160a01b0383166108405760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420415653206164647265737360681b60448201526064016105ae565b6001600160a01b03821661088e5760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b60448201526064016105ae565b61089661101f565b61089f8261104d565b606580546001600160a01b038581166001600160c01b03199092168217909255606680546001810182555f9182527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943546008820401805463ffffffff60079093166004026101000a92830219169091179055604051928516927f6f590f00594e422d605859c867818ebccb321e2bdc9dd35c5edb1e0a067d06019190a37f836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51606560149054906101000a900463ffffffff16606660405161097f929190611611565b60405180910390a180156109cc575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6109d9610f48565b6109e25f61104d565b565b6001600160a01b0381165f908152606a60205260409020805460609190610a0a906117b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610a36906117b8565b8015610a815780601f10610a5857610100808354040283529160200191610a81565b820191905f5260205f20905b815481529060010190602001808311610a6457829003601f168201915b50505050509050919050565b610a95610f48565b6001600160a01b038216610abb5760405162461bcd60e51b81526004016105ae906117f0565b5f815111610afc5760405162461bcd60e51b815260206004820152600e60248201526d125b9d985b1a59081cdbd8dad95d60921b60448201526064016105ae565b6001600160a01b0382165f908152606a60205260409020610b1d8282611865565b50816001600160a01b03167f0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b82604051610b5791906112d9565b60405180910390a25050565b6065546001600160a01b03858116911614610bae5760405162461bcd60e51b815260206004820152600b60248201526a496e76616c69642041565360a81b60448201526064016105ae565b6001600160a01b038516610bd45760405162461bcd60e51b81526004016105ae906117f0565b6001600160a01b0385165f9081526067602052604090205460ff1615610c315760405162461bcd60e51b8152602060048201526012602482015271105b1c9958591e481c9959da5cdd195c995960721b60448201526064016105ae565b6001600160a01b0385165f9081526068602052604090205460ff16610c985760405162461bcd60e51b815260206004820152601860248201527f4f70657261746f72206e6f7420616c6c6f776c6973746564000000000000000060448201526064016105ae565b805115610cc2575f81806020019051810190610cb4919061191f565b9050610cc08682610afc565b505b6001600160a01b0385165f81815260676020526040808220805460ff1916600190811790915560698054918201815583527f7fb4302e8e91f9110a6554c2c0a24601252c2a42c2220ca988efcfe3999143080180546001600160a01b03191684179055517f4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e59190a25050505050565b610d59610f48565b5f81602001515111610dbf5760405162461bcd60e51b815260206004820152602960248201527f4578656375746f72206f70657261746f7220736574204944732063616e6e6f7460448201526820626520656d70747960b81b60648201526084016105ae565b80516065805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055610df060665f6110cd565b5f5b816020015151811015610e6757606682602001518281518110610e1757610e176115d9565b6020908102919091018101518254600181810185555f948552929093206008840401805463ffffffff92831660046007909616959095026101000a94850292909402199093161790915501610df2565b50805160208201516040517f836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef5192610e9f929091611993565b60405180910390a150565b610eb2610f48565b6001600160a01b038116610f175760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105ae565b6103de8161104d565b60698181548110610f2f575f80fd5b5f918252602090912001546001600160a01b0316905081565b6033546001600160a01b031633146109e25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ae565b6001600160a01b038216610fc85760405162461bcd60e51b81526004016105ae906117f0565b6001600160a01b0382165f81815260686020908152604091829020805460ff191685151590811790915591519182527fbc9494536acaf3f702bbec70ff790bb3930056ae9e781e8077097b9de89a8e189101610b57565b5f54610100900460ff166110455760405162461bcd60e51b81526004016105ae906119e0565b6109e261109e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166110c45760405162461bcd60e51b81526004016105ae906119e0565b6109e23361104d565b5080545f825560070160089004905f5260205f20908101906103de91905b808211156110fe575f81556001016110eb565b5090565b80356001600160a01b0381168114611118575f5ffd5b919050565b5f6020828403121561112d575f5ffd5b61113682611102565b9392505050565b602080825282518282018190525f918401906040840190835b8181101561117d5783516001600160a01b0316835260209384019390920191600101611156565b509095945050505050565b5f5f83601f840112611198575f5ffd5b5081356001600160401b038111156111ae575f5ffd5b6020830191508360208260051b85010111156111c8575f5ffd5b9250929050565b5f5f5f5f606085870312156111e2575f5ffd5b6111eb85611102565b93506111f960208601611102565b925060408501356001600160401b03811115611213575f5ffd5b61121f87828801611188565b95989497509550505050565b5f6020828403121561123b575f5ffd5b5035919050565b6020808252825163ffffffff1682820152828101516040808401528051606084018190525f929190910190829060808501905b8083101561129e5763ffffffff8451168252602082019150602084019350600183019250611275565b5095945050505050565b5f5f604083850312156112b9575f5ffd5b6112c283611102565b91506112d060208401611102565b90509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156113445761134461130e565b60405290565b604051601f8201601f191681016001600160401b03811182821017156113725761137261130e565b604052919050565b5f6001600160401b038211156113925761139261130e565b50601f01601f191660200190565b5f6113b26113ad8461137a565b61134a565b90508281528383830111156113c5575f5ffd5b828260208301375f602084830101529392505050565b5f5f604083850312156113ec575f5ffd5b6113f583611102565b915060208301356001600160401b0381111561140f575f5ffd5b8301601f8101851361141f575f5ffd5b61142e858235602084016113a0565b9150509250929050565b5f5f5f5f5f6080868803121561144c575f5ffd5b61145586611102565b945061146360208701611102565b935060408601356001600160401b0381111561147d575f5ffd5b61148988828901611188565b90945092505060608601356001600160401b038111156114a7575f5ffd5b8601601f810188136114b7575f5ffd5b6114c6888235602084016113a0565b9150509295509295909350565b803563ffffffff81168114611118575f5ffd5b5f602082840312156114f6575f5ffd5b81356001600160401b0381111561150b575f5ffd5b82016040818503121561151c575f5ffd5b611524611322565b61152d826114d3565b815260208201356001600160401b03811115611547575f5ffd5b80830192505084601f83011261155b575f5ffd5b81356001600160401b038111156115745761157461130e565b8060051b6115846020820161134a565b9182526020818501810192908101908884111561159f575f5ffd5b6020860195505b838610156115c8576115b7866114d3565b8252602095860195909101906115a6565b602085015250919695505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f6001820161160a57634e487b7160e01b5f52601160045260245ffd5b5060010190565b5f6040820163ffffffff8516835260406020840152808454611637818490815260200190565b5f8781526020812094509092505b816007820110156116bd57835463ffffffff8082168552602082811c821690860152604082811c821690860152606082811c821690860152608082811c82169086015260a082811c82169086015260c082811c9091169085015260e090811c9084015260019093019261010090920191600801611645565b925492818110156116db5763ffffffff841683526020909201916001015b818110156116fb57602084901c63ffffffff168352602092909201916001015b8181101561171a5763ffffffff604085901c1683526020909201916001015b818110156117395763ffffffff606085901c1683526020909201916001015b818110156117585763ffffffff608085901c1683526020909201916001015b818110156117775763ffffffff60a085901c1683526020909201916001015b818110156117965763ffffffff60c085901c1683526020909201916001015b818110156117ac5760e084901c83526020830192505b50909695505050505050565b600181811c908216806117cc57607f821691505b6020821081036117ea57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526010908201526f24b73b30b634b21037b832b930ba37b960811b604082015260600190565b601f8211156109cc57805f5260205f20601f840160051c8101602085101561183f5750805b601f840160051c820191505b8181101561185e575f815560010161184b565b5050505050565b81516001600160401b0381111561187e5761187e61130e565b6118928161188c84546117b8565b8461181a565b6020601f8211600181146118c4575f83156118ad5750848201515b5f19600385901b1c1916600184901b17845561185e565b5f84815260208120601f198516915b828110156118f357878501518255602094850194600190920191016118d3565b508482101561191057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6020828403121561192f575f5ffd5b81516001600160401b03811115611944575f5ffd5b8201601f81018413611954575f5ffd5b80516119626113ad8261137a565b818152856020838501011115611976575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b5f6040820163ffffffff85168352604060208401528084518083526060850191506020860192505f5b818110156117ac57835163ffffffff168352602093840193909201916001016119bc565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea264697066735822122079061cd44b8144c6f51c7d428ac67f3c8b12f764596a2580e679202a7eaded7264736f6c634300081b0033",
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

// AggregatorOperatorSetId is a free data retrieval call binding the contract method 0x8b2942d4.
//
// Solidity: function aggregatorOperatorSetId() view returns(uint32)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) AggregatorOperatorSetId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "aggregatorOperatorSetId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AggregatorOperatorSetId is a free data retrieval call binding the contract method 0x8b2942d4.
//
// Solidity: function aggregatorOperatorSetId() view returns(uint32)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AggregatorOperatorSetId() (uint32, error) {
	return _TaskAVSRegistrar.Contract.AggregatorOperatorSetId(&_TaskAVSRegistrar.CallOpts)
}

// AggregatorOperatorSetId is a free data retrieval call binding the contract method 0x8b2942d4.
//
// Solidity: function aggregatorOperatorSetId() view returns(uint32)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) AggregatorOperatorSetId() (uint32, error) {
	return _TaskAVSRegistrar.Contract.AggregatorOperatorSetId(&_TaskAVSRegistrar.CallOpts)
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

// ExecutorOperatorSetIds is a free data retrieval call binding the contract method 0x3dc70058.
//
// Solidity: function executorOperatorSetIds(uint256 ) view returns(uint32)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) ExecutorOperatorSetIds(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "executorOperatorSetIds", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ExecutorOperatorSetIds is a free data retrieval call binding the contract method 0x3dc70058.
//
// Solidity: function executorOperatorSetIds(uint256 ) view returns(uint32)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) ExecutorOperatorSetIds(arg0 *big.Int) (uint32, error) {
	return _TaskAVSRegistrar.Contract.ExecutorOperatorSetIds(&_TaskAVSRegistrar.CallOpts, arg0)
}

// ExecutorOperatorSetIds is a free data retrieval call binding the contract method 0x3dc70058.
//
// Solidity: function executorOperatorSetIds(uint256 ) view returns(uint32)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) ExecutorOperatorSetIds(arg0 *big.Int) (uint32, error) {
	return _TaskAVSRegistrar.Contract.ExecutorOperatorSetIds(&_TaskAVSRegistrar.CallOpts, arg0)
}

// GetAvsConfig is a free data retrieval call binding the contract method 0x41f548f0.
//
// Solidity: function getAvsConfig() view returns((uint32,uint32[]) config)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetAvsConfig(opts *bind.CallOpts) (TaskAVSRegistrarAvsConfig, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getAvsConfig")

	if err != nil {
		return *new(TaskAVSRegistrarAvsConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskAVSRegistrarAvsConfig)).(*TaskAVSRegistrarAvsConfig)

	return out0, err

}

// GetAvsConfig is a free data retrieval call binding the contract method 0x41f548f0.
//
// Solidity: function getAvsConfig() view returns((uint32,uint32[]) config)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetAvsConfig() (TaskAVSRegistrarAvsConfig, error) {
	return _TaskAVSRegistrar.Contract.GetAvsConfig(&_TaskAVSRegistrar.CallOpts)
}

// GetAvsConfig is a free data retrieval call binding the contract method 0x41f548f0.
//
// Solidity: function getAvsConfig() view returns((uint32,uint32[]) config)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetAvsConfig() (TaskAVSRegistrarAvsConfig, error) {
	return _TaskAVSRegistrar.Contract.GetAvsConfig(&_TaskAVSRegistrar.CallOpts)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetOperatorSocket(opts *bind.CallOpts, operator common.Address) (string, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getOperatorSocket", operator)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _TaskAVSRegistrar.Contract.GetOperatorSocket(&_TaskAVSRegistrar.CallOpts, operator)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _TaskAVSRegistrar.Contract.GetOperatorSocket(&_TaskAVSRegistrar.CallOpts, operator)
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

// SetAvsConfig is a paid mutator transaction binding the contract method 0xd1f2e81d.
//
// Solidity: function setAvsConfig((uint32,uint32[]) config) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) SetAvsConfig(opts *bind.TransactOpts, config TaskAVSRegistrarAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "setAvsConfig", config)
}

// SetAvsConfig is a paid mutator transaction binding the contract method 0xd1f2e81d.
//
// Solidity: function setAvsConfig((uint32,uint32[]) config) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) SetAvsConfig(config TaskAVSRegistrarAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.SetAvsConfig(&_TaskAVSRegistrar.TransactOpts, config)
}

// SetAvsConfig is a paid mutator transaction binding the contract method 0xd1f2e81d.
//
// Solidity: function setAvsConfig((uint32,uint32[]) config) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) SetAvsConfig(config TaskAVSRegistrarAvsConfig) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.SetAvsConfig(&_TaskAVSRegistrar.TransactOpts, config)
}

// SetOperatorSocket is a paid mutator transaction binding the contract method 0xadc8ea74.
//
// Solidity: function setOperatorSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) SetOperatorSocket(opts *bind.TransactOpts, operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "setOperatorSocket", operator, socket)
}

// SetOperatorSocket is a paid mutator transaction binding the contract method 0xadc8ea74.
//
// Solidity: function setOperatorSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) SetOperatorSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.SetOperatorSocket(&_TaskAVSRegistrar.TransactOpts, operator, socket)
}

// SetOperatorSocket is a paid mutator transaction binding the contract method 0xadc8ea74.
//
// Solidity: function setOperatorSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) SetOperatorSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.SetOperatorSocket(&_TaskAVSRegistrar.TransactOpts, operator, socket)
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

// TaskAVSRegistrarAvsConfigSetIterator is returned from FilterAvsConfigSet and is used to iterate over the raw logs and unpacked data for AvsConfigSet events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarAvsConfigSetIterator struct {
	Event *TaskAVSRegistrarAvsConfigSet // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarAvsConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarAvsConfigSet)
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
		it.Event = new(TaskAVSRegistrarAvsConfigSet)
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
func (it *TaskAVSRegistrarAvsConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarAvsConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarAvsConfigSet represents a AvsConfigSet event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarAvsConfigSet struct {
	AggregatorOperatorSetId uint32
	ExecutorOperatorSetIds  []uint32
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAvsConfigSet is a free log retrieval operation binding the contract event 0x836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51.
//
// Solidity: event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterAvsConfigSet(opts *bind.FilterOpts) (*TaskAVSRegistrarAvsConfigSetIterator, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "AvsConfigSet")
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarAvsConfigSetIterator{contract: _TaskAVSRegistrar.contract, event: "AvsConfigSet", logs: logs, sub: sub}, nil
}

// WatchAvsConfigSet is a free log subscription operation binding the contract event 0x836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51.
//
// Solidity: event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchAvsConfigSet(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarAvsConfigSet) (event.Subscription, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "AvsConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarAvsConfigSet)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "AvsConfigSet", log); err != nil {
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

// ParseAvsConfigSet is a log parse operation binding the contract event 0x836f1d33f6d85cfc7b24565d309c6e1486cf56dd3d8267a9651e05b88342ef51.
//
// Solidity: event AvsConfigSet(uint32 aggregatorOperatorSetId, uint32[] executorOperatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseAvsConfigSet(log types.Log) (*TaskAVSRegistrarAvsConfigSet, error) {
	event := new(TaskAVSRegistrarAvsConfigSet)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "AvsConfigSet", log); err != nil {
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

// TaskAVSRegistrarOperatorSocketSetIterator is returned from FilterOperatorSocketSet and is used to iterate over the raw logs and unpacked data for OperatorSocketSet events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorSocketSetIterator struct {
	Event *TaskAVSRegistrarOperatorSocketSet // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorSocketSet)
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
		it.Event = new(TaskAVSRegistrarOperatorSocketSet)
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
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorSocketSet represents a OperatorSocketSet event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorSocketSet struct {
	Operator common.Address
	Socket   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketSet is a free log retrieval operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorSocketSet(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorSocketSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorSocketSetIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorSocketSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketSet is a free log subscription operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorSocketSet(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorSocketSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorSocketSet)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
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

// ParseOperatorSocketSet is a log parse operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorSocketSet(log types.Log) (*TaskAVSRegistrarOperatorSocketSet, error) {
	event := new(TaskAVSRegistrarOperatorSocketSet)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
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
