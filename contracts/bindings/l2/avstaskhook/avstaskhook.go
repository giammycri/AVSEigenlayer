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

// AVSTaskHookMetaData contains all meta data concerning the AVSTaskHook contract.
var AVSTaskHookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"hookTask\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TaskHooked\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061015b8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c806327f5cab41461002d575b5f5ffd5b61004061003b366004610089565b610042565b005b336001600160a01b03167fb0738a695d9bb0308414924b57a7b842be09926f59a0ab342c7ca67d210fefdb838360405161007d9291906100f7565b60405180910390a25050565b5f5f6020838503121561009a575f5ffd5b823567ffffffffffffffff8111156100b0575f5ffd5b8301601f810185136100c0575f5ffd5b803567ffffffffffffffff8111156100d6575f5ffd5b8560208284010111156100e7575f5ffd5b6020919091019590945092505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea26469706673582212206e082b0ce6acb5f12d08832d0f48fb788400f5dde2fc077000c826cc2d2f4f7364736f6c634300081b0033",
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

// HookTask is a paid mutator transaction binding the contract method 0x27f5cab4.
//
// Solidity: function hookTask(bytes data) returns()
func (_AVSTaskHook *AVSTaskHookTransactor) HookTask(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _AVSTaskHook.contract.Transact(opts, "hookTask", data)
}

// HookTask is a paid mutator transaction binding the contract method 0x27f5cab4.
//
// Solidity: function hookTask(bytes data) returns()
func (_AVSTaskHook *AVSTaskHookSession) HookTask(data []byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HookTask(&_AVSTaskHook.TransactOpts, data)
}

// HookTask is a paid mutator transaction binding the contract method 0x27f5cab4.
//
// Solidity: function hookTask(bytes data) returns()
func (_AVSTaskHook *AVSTaskHookTransactorSession) HookTask(data []byte) (*types.Transaction, error) {
	return _AVSTaskHook.Contract.HookTask(&_AVSTaskHook.TransactOpts, data)
}

// AVSTaskHookTaskHookedIterator is returned from FilterTaskHooked and is used to iterate over the raw logs and unpacked data for TaskHooked events raised by the AVSTaskHook contract.
type AVSTaskHookTaskHookedIterator struct {
	Event *AVSTaskHookTaskHooked // Event containing the contract specifics and raw log

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
func (it *AVSTaskHookTaskHookedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSTaskHookTaskHooked)
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
		it.Event = new(AVSTaskHookTaskHooked)
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
func (it *AVSTaskHookTaskHookedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSTaskHookTaskHookedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSTaskHookTaskHooked represents a TaskHooked event raised by the AVSTaskHook contract.
type AVSTaskHookTaskHooked struct {
	Sender common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskHooked is a free log retrieval operation binding the contract event 0xb0738a695d9bb0308414924b57a7b842be09926f59a0ab342c7ca67d210fefdb.
//
// Solidity: event TaskHooked(address indexed sender, bytes data)
func (_AVSTaskHook *AVSTaskHookFilterer) FilterTaskHooked(opts *bind.FilterOpts, sender []common.Address) (*AVSTaskHookTaskHookedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AVSTaskHook.contract.FilterLogs(opts, "TaskHooked", senderRule)
	if err != nil {
		return nil, err
	}
	return &AVSTaskHookTaskHookedIterator{contract: _AVSTaskHook.contract, event: "TaskHooked", logs: logs, sub: sub}, nil
}

// WatchTaskHooked is a free log subscription operation binding the contract event 0xb0738a695d9bb0308414924b57a7b842be09926f59a0ab342c7ca67d210fefdb.
//
// Solidity: event TaskHooked(address indexed sender, bytes data)
func (_AVSTaskHook *AVSTaskHookFilterer) WatchTaskHooked(opts *bind.WatchOpts, sink chan<- *AVSTaskHookTaskHooked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AVSTaskHook.contract.WatchLogs(opts, "TaskHooked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSTaskHookTaskHooked)
				if err := _AVSTaskHook.contract.UnpackLog(event, "TaskHooked", log); err != nil {
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

// ParseTaskHooked is a log parse operation binding the contract event 0xb0738a695d9bb0308414924b57a7b842be09926f59a0ab342c7ca67d210fefdb.
//
// Solidity: event TaskHooked(address indexed sender, bytes data)
func (_AVSTaskHook *AVSTaskHookFilterer) ParseTaskHooked(log types.Log) (*AVSTaskHookTaskHooked, error) {
	event := new(AVSTaskHookTaskHooked)
	if err := _AVSTaskHook.contract.UnpackLog(event, "TaskHooked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
