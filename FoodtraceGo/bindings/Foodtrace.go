// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// FoodItemFoodItem is an auto generated low-level Go binding around an user-defined struct.
type FoodItemFoodItem struct {
	Id                *big.Int
	Name              string
	Producer          common.Address
	Details           string
	LogisticsRecords  []string
	LogisticsRecorded bool
}

// FoodtraceMetaData contains all meta data concerning the Foodtrace contract.
var FoodtraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"addLogisticsProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"FoodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"LogisticsProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"logisticsRecord\",\"type\":\"string\"}],\"name\":\"LogisticsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_details\",\"type\":\"string\"}],\"name\":\"registerFood\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_foodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_logisticsInfo\",\"type\":\"string\"}],\"name\":\"updateLogistice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foodItemCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"foodItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"logisticsRecorded\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_foodId\",\"type\":\"uint256\"}],\"name\":\"getFoodInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"logisticsRecords\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"logisticsRecorded\",\"type\":\"bool\"}],\"internalType\":\"structFoodItem.FoodItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"logisticsProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logisticsRecordedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FoodtraceABI is the input ABI used to generate the binding from.
// Deprecated: Use FoodtraceMetaData.ABI instead.
var FoodtraceABI = FoodtraceMetaData.ABI

// Foodtrace is an auto generated Go binding around an Ethereum contract.
type Foodtrace struct {
	FoodtraceCaller     // Read-only binding to the contract
	FoodtraceTransactor // Write-only binding to the contract
	FoodtraceFilterer   // Log filterer for contract events
}

// FoodtraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoodtraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodtraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoodtraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodtraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoodtraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodtraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoodtraceSession struct {
	Contract     *Foodtrace        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoodtraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoodtraceCallerSession struct {
	Contract *FoodtraceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FoodtraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoodtraceTransactorSession struct {
	Contract     *FoodtraceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FoodtraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoodtraceRaw struct {
	Contract *Foodtrace // Generic contract binding to access the raw methods on
}

// FoodtraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoodtraceCallerRaw struct {
	Contract *FoodtraceCaller // Generic read-only contract binding to access the raw methods on
}

// FoodtraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoodtraceTransactorRaw struct {
	Contract *FoodtraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoodtrace creates a new instance of Foodtrace, bound to a specific deployed contract.
func NewFoodtrace(address common.Address, backend bind.ContractBackend) (*Foodtrace, error) {
	contract, err := bindFoodtrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foodtrace{FoodtraceCaller: FoodtraceCaller{contract: contract}, FoodtraceTransactor: FoodtraceTransactor{contract: contract}, FoodtraceFilterer: FoodtraceFilterer{contract: contract}}, nil
}

// NewFoodtraceCaller creates a new read-only instance of Foodtrace, bound to a specific deployed contract.
func NewFoodtraceCaller(address common.Address, caller bind.ContractCaller) (*FoodtraceCaller, error) {
	contract, err := bindFoodtrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoodtraceCaller{contract: contract}, nil
}

// NewFoodtraceTransactor creates a new write-only instance of Foodtrace, bound to a specific deployed contract.
func NewFoodtraceTransactor(address common.Address, transactor bind.ContractTransactor) (*FoodtraceTransactor, error) {
	contract, err := bindFoodtrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoodtraceTransactor{contract: contract}, nil
}

// NewFoodtraceFilterer creates a new log filterer instance of Foodtrace, bound to a specific deployed contract.
func NewFoodtraceFilterer(address common.Address, filterer bind.ContractFilterer) (*FoodtraceFilterer, error) {
	contract, err := bindFoodtrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoodtraceFilterer{contract: contract}, nil
}

// bindFoodtrace binds a generic wrapper to an already deployed contract.
func bindFoodtrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FoodtraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foodtrace *FoodtraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foodtrace.Contract.FoodtraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foodtrace *FoodtraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foodtrace.Contract.FoodtraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foodtrace *FoodtraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foodtrace.Contract.FoodtraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foodtrace *FoodtraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foodtrace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foodtrace *FoodtraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foodtrace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foodtrace *FoodtraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foodtrace.Contract.contract.Transact(opts, method, params...)
}

// FoodItemCount is a free data retrieval call binding the contract method 0x872ce6dc.
//
// Solidity: function foodItemCount() view returns(uint256)
func (_Foodtrace *FoodtraceCaller) FoodItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Foodtrace.contract.Call(opts, &out, "foodItemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoodItemCount is a free data retrieval call binding the contract method 0x872ce6dc.
//
// Solidity: function foodItemCount() view returns(uint256)
func (_Foodtrace *FoodtraceSession) FoodItemCount() (*big.Int, error) {
	return _Foodtrace.Contract.FoodItemCount(&_Foodtrace.CallOpts)
}

// FoodItemCount is a free data retrieval call binding the contract method 0x872ce6dc.
//
// Solidity: function foodItemCount() view returns(uint256)
func (_Foodtrace *FoodtraceCallerSession) FoodItemCount() (*big.Int, error) {
	return _Foodtrace.Contract.FoodItemCount(&_Foodtrace.CallOpts)
}

// FoodItems is a free data retrieval call binding the contract method 0x0ac52e65.
//
// Solidity: function foodItems(uint256 ) view returns(uint256 id, string name, address producer, string details, bool logisticsRecorded)
func (_Foodtrace *FoodtraceCaller) FoodItems(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	Producer          common.Address
	Details           string
	LogisticsRecorded bool
}, error) {
	var out []interface{}
	err := _Foodtrace.contract.Call(opts, &out, "foodItems", arg0)

	outstruct := new(struct {
		Id                *big.Int
		Name              string
		Producer          common.Address
		Details           string
		LogisticsRecorded bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Producer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Details = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.LogisticsRecorded = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// FoodItems is a free data retrieval call binding the contract method 0x0ac52e65.
//
// Solidity: function foodItems(uint256 ) view returns(uint256 id, string name, address producer, string details, bool logisticsRecorded)
func (_Foodtrace *FoodtraceSession) FoodItems(arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	Producer          common.Address
	Details           string
	LogisticsRecorded bool
}, error) {
	return _Foodtrace.Contract.FoodItems(&_Foodtrace.CallOpts, arg0)
}

// FoodItems is a free data retrieval call binding the contract method 0x0ac52e65.
//
// Solidity: function foodItems(uint256 ) view returns(uint256 id, string name, address producer, string details, bool logisticsRecorded)
func (_Foodtrace *FoodtraceCallerSession) FoodItems(arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	Producer          common.Address
	Details           string
	LogisticsRecorded bool
}, error) {
	return _Foodtrace.Contract.FoodItems(&_Foodtrace.CallOpts, arg0)
}

// GetFoodInfo is a free data retrieval call binding the contract method 0x0a8eb8bc.
//
// Solidity: function getFoodInfo(uint256 _foodId) view returns((uint256,string,address,string,string[],bool))
func (_Foodtrace *FoodtraceCaller) GetFoodInfo(opts *bind.CallOpts, _foodId *big.Int) (FoodItemFoodItem, error) {
	var out []interface{}
	err := _Foodtrace.contract.Call(opts, &out, "getFoodInfo", _foodId)

	if err != nil {
		return *new(FoodItemFoodItem), err
	}

	out0 := *abi.ConvertType(out[0], new(FoodItemFoodItem)).(*FoodItemFoodItem)

	return out0, err

}

// GetFoodInfo is a free data retrieval call binding the contract method 0x0a8eb8bc.
//
// Solidity: function getFoodInfo(uint256 _foodId) view returns((uint256,string,address,string,string[],bool))
func (_Foodtrace *FoodtraceSession) GetFoodInfo(_foodId *big.Int) (FoodItemFoodItem, error) {
	return _Foodtrace.Contract.GetFoodInfo(&_Foodtrace.CallOpts, _foodId)
}

// GetFoodInfo is a free data retrieval call binding the contract method 0x0a8eb8bc.
//
// Solidity: function getFoodInfo(uint256 _foodId) view returns((uint256,string,address,string,string[],bool))
func (_Foodtrace *FoodtraceCallerSession) GetFoodInfo(_foodId *big.Int) (FoodItemFoodItem, error) {
	return _Foodtrace.Contract.GetFoodInfo(&_Foodtrace.CallOpts, _foodId)
}

// LogisticsProviders is a free data retrieval call binding the contract method 0x71910b86.
//
// Solidity: function logisticsProviders(address ) view returns(bool)
func (_Foodtrace *FoodtraceCaller) LogisticsProviders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Foodtrace.contract.Call(opts, &out, "logisticsProviders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LogisticsProviders is a free data retrieval call binding the contract method 0x71910b86.
//
// Solidity: function logisticsProviders(address ) view returns(bool)
func (_Foodtrace *FoodtraceSession) LogisticsProviders(arg0 common.Address) (bool, error) {
	return _Foodtrace.Contract.LogisticsProviders(&_Foodtrace.CallOpts, arg0)
}

// LogisticsProviders is a free data retrieval call binding the contract method 0x71910b86.
//
// Solidity: function logisticsProviders(address ) view returns(bool)
func (_Foodtrace *FoodtraceCallerSession) LogisticsProviders(arg0 common.Address) (bool, error) {
	return _Foodtrace.Contract.LogisticsProviders(&_Foodtrace.CallOpts, arg0)
}

// LogisticsRecordedCount is a free data retrieval call binding the contract method 0xc7714b60.
//
// Solidity: function logisticsRecordedCount() view returns(uint256)
func (_Foodtrace *FoodtraceCaller) LogisticsRecordedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Foodtrace.contract.Call(opts, &out, "logisticsRecordedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LogisticsRecordedCount is a free data retrieval call binding the contract method 0xc7714b60.
//
// Solidity: function logisticsRecordedCount() view returns(uint256)
func (_Foodtrace *FoodtraceSession) LogisticsRecordedCount() (*big.Int, error) {
	return _Foodtrace.Contract.LogisticsRecordedCount(&_Foodtrace.CallOpts)
}

// LogisticsRecordedCount is a free data retrieval call binding the contract method 0xc7714b60.
//
// Solidity: function logisticsRecordedCount() view returns(uint256)
func (_Foodtrace *FoodtraceCallerSession) LogisticsRecordedCount() (*big.Int, error) {
	return _Foodtrace.Contract.LogisticsRecordedCount(&_Foodtrace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Foodtrace *FoodtraceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Foodtrace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Foodtrace *FoodtraceSession) Owner() (common.Address, error) {
	return _Foodtrace.Contract.Owner(&_Foodtrace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Foodtrace *FoodtraceCallerSession) Owner() (common.Address, error) {
	return _Foodtrace.Contract.Owner(&_Foodtrace.CallOpts)
}

// AddLogisticsProvider is a paid mutator transaction binding the contract method 0x52b91103.
//
// Solidity: function addLogisticsProvider(address _provider) returns()
func (_Foodtrace *FoodtraceTransactor) AddLogisticsProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _Foodtrace.contract.Transact(opts, "addLogisticsProvider", _provider)
}

// AddLogisticsProvider is a paid mutator transaction binding the contract method 0x52b91103.
//
// Solidity: function addLogisticsProvider(address _provider) returns()
func (_Foodtrace *FoodtraceSession) AddLogisticsProvider(_provider common.Address) (*types.Transaction, error) {
	return _Foodtrace.Contract.AddLogisticsProvider(&_Foodtrace.TransactOpts, _provider)
}

// AddLogisticsProvider is a paid mutator transaction binding the contract method 0x52b91103.
//
// Solidity: function addLogisticsProvider(address _provider) returns()
func (_Foodtrace *FoodtraceTransactorSession) AddLogisticsProvider(_provider common.Address) (*types.Transaction, error) {
	return _Foodtrace.Contract.AddLogisticsProvider(&_Foodtrace.TransactOpts, _provider)
}

// RegisterFood is a paid mutator transaction binding the contract method 0x94184622.
//
// Solidity: function registerFood(string _name, string _details) returns()
func (_Foodtrace *FoodtraceTransactor) RegisterFood(opts *bind.TransactOpts, _name string, _details string) (*types.Transaction, error) {
	return _Foodtrace.contract.Transact(opts, "registerFood", _name, _details)
}

// RegisterFood is a paid mutator transaction binding the contract method 0x94184622.
//
// Solidity: function registerFood(string _name, string _details) returns()
func (_Foodtrace *FoodtraceSession) RegisterFood(_name string, _details string) (*types.Transaction, error) {
	return _Foodtrace.Contract.RegisterFood(&_Foodtrace.TransactOpts, _name, _details)
}

// RegisterFood is a paid mutator transaction binding the contract method 0x94184622.
//
// Solidity: function registerFood(string _name, string _details) returns()
func (_Foodtrace *FoodtraceTransactorSession) RegisterFood(_name string, _details string) (*types.Transaction, error) {
	return _Foodtrace.Contract.RegisterFood(&_Foodtrace.TransactOpts, _name, _details)
}

// UpdateLogistice is a paid mutator transaction binding the contract method 0x67422ecc.
//
// Solidity: function updateLogistice(uint256 _foodId, string _logisticsInfo) returns()
func (_Foodtrace *FoodtraceTransactor) UpdateLogistice(opts *bind.TransactOpts, _foodId *big.Int, _logisticsInfo string) (*types.Transaction, error) {
	return _Foodtrace.contract.Transact(opts, "updateLogistice", _foodId, _logisticsInfo)
}

// UpdateLogistice is a paid mutator transaction binding the contract method 0x67422ecc.
//
// Solidity: function updateLogistice(uint256 _foodId, string _logisticsInfo) returns()
func (_Foodtrace *FoodtraceSession) UpdateLogistice(_foodId *big.Int, _logisticsInfo string) (*types.Transaction, error) {
	return _Foodtrace.Contract.UpdateLogistice(&_Foodtrace.TransactOpts, _foodId, _logisticsInfo)
}

// UpdateLogistice is a paid mutator transaction binding the contract method 0x67422ecc.
//
// Solidity: function updateLogistice(uint256 _foodId, string _logisticsInfo) returns()
func (_Foodtrace *FoodtraceTransactorSession) UpdateLogistice(_foodId *big.Int, _logisticsInfo string) (*types.Transaction, error) {
	return _Foodtrace.Contract.UpdateLogistice(&_Foodtrace.TransactOpts, _foodId, _logisticsInfo)
}

// FoodtraceFoodRegisteredIterator is returned from FilterFoodRegistered and is used to iterate over the raw logs and unpacked data for FoodRegistered events raised by the Foodtrace contract.
type FoodtraceFoodRegisteredIterator struct {
	Event *FoodtraceFoodRegistered // Event containing the contract specifics and raw log

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
func (it *FoodtraceFoodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodtraceFoodRegistered)
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
		it.Event = new(FoodtraceFoodRegistered)
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
func (it *FoodtraceFoodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodtraceFoodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodtraceFoodRegistered represents a FoodRegistered event raised by the Foodtrace contract.
type FoodtraceFoodRegistered struct {
	Id       *big.Int
	Name     string
	Producer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFoodRegistered is a free log retrieval operation binding the contract event 0x5bcc0d3c964557be3d068e7be65e7e7e1849efaef771568652a400d5d3293626.
//
// Solidity: event FoodRegistered(uint256 id, string name, address producer)
func (_Foodtrace *FoodtraceFilterer) FilterFoodRegistered(opts *bind.FilterOpts) (*FoodtraceFoodRegisteredIterator, error) {

	logs, sub, err := _Foodtrace.contract.FilterLogs(opts, "FoodRegistered")
	if err != nil {
		return nil, err
	}
	return &FoodtraceFoodRegisteredIterator{contract: _Foodtrace.contract, event: "FoodRegistered", logs: logs, sub: sub}, nil
}

// WatchFoodRegistered is a free log subscription operation binding the contract event 0x5bcc0d3c964557be3d068e7be65e7e7e1849efaef771568652a400d5d3293626.
//
// Solidity: event FoodRegistered(uint256 id, string name, address producer)
func (_Foodtrace *FoodtraceFilterer) WatchFoodRegistered(opts *bind.WatchOpts, sink chan<- *FoodtraceFoodRegistered) (event.Subscription, error) {

	logs, sub, err := _Foodtrace.contract.WatchLogs(opts, "FoodRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodtraceFoodRegistered)
				if err := _Foodtrace.contract.UnpackLog(event, "FoodRegistered", log); err != nil {
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

// ParseFoodRegistered is a log parse operation binding the contract event 0x5bcc0d3c964557be3d068e7be65e7e7e1849efaef771568652a400d5d3293626.
//
// Solidity: event FoodRegistered(uint256 id, string name, address producer)
func (_Foodtrace *FoodtraceFilterer) ParseFoodRegistered(log types.Log) (*FoodtraceFoodRegistered, error) {
	event := new(FoodtraceFoodRegistered)
	if err := _Foodtrace.contract.UnpackLog(event, "FoodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoodtraceLogisticsProviderAddedIterator is returned from FilterLogisticsProviderAdded and is used to iterate over the raw logs and unpacked data for LogisticsProviderAdded events raised by the Foodtrace contract.
type FoodtraceLogisticsProviderAddedIterator struct {
	Event *FoodtraceLogisticsProviderAdded // Event containing the contract specifics and raw log

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
func (it *FoodtraceLogisticsProviderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodtraceLogisticsProviderAdded)
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
		it.Event = new(FoodtraceLogisticsProviderAdded)
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
func (it *FoodtraceLogisticsProviderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodtraceLogisticsProviderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodtraceLogisticsProviderAdded represents a LogisticsProviderAdded event raised by the Foodtrace contract.
type FoodtraceLogisticsProviderAdded struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogisticsProviderAdded is a free log retrieval operation binding the contract event 0x1f9e16a285225789af2788825a63c4efb270c52d15d84d2969be7f239c96061e.
//
// Solidity: event LogisticsProviderAdded(address provider)
func (_Foodtrace *FoodtraceFilterer) FilterLogisticsProviderAdded(opts *bind.FilterOpts) (*FoodtraceLogisticsProviderAddedIterator, error) {

	logs, sub, err := _Foodtrace.contract.FilterLogs(opts, "LogisticsProviderAdded")
	if err != nil {
		return nil, err
	}
	return &FoodtraceLogisticsProviderAddedIterator{contract: _Foodtrace.contract, event: "LogisticsProviderAdded", logs: logs, sub: sub}, nil
}

// WatchLogisticsProviderAdded is a free log subscription operation binding the contract event 0x1f9e16a285225789af2788825a63c4efb270c52d15d84d2969be7f239c96061e.
//
// Solidity: event LogisticsProviderAdded(address provider)
func (_Foodtrace *FoodtraceFilterer) WatchLogisticsProviderAdded(opts *bind.WatchOpts, sink chan<- *FoodtraceLogisticsProviderAdded) (event.Subscription, error) {

	logs, sub, err := _Foodtrace.contract.WatchLogs(opts, "LogisticsProviderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodtraceLogisticsProviderAdded)
				if err := _Foodtrace.contract.UnpackLog(event, "LogisticsProviderAdded", log); err != nil {
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

// ParseLogisticsProviderAdded is a log parse operation binding the contract event 0x1f9e16a285225789af2788825a63c4efb270c52d15d84d2969be7f239c96061e.
//
// Solidity: event LogisticsProviderAdded(address provider)
func (_Foodtrace *FoodtraceFilterer) ParseLogisticsProviderAdded(log types.Log) (*FoodtraceLogisticsProviderAdded, error) {
	event := new(FoodtraceLogisticsProviderAdded)
	if err := _Foodtrace.contract.UnpackLog(event, "LogisticsProviderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoodtraceLogisticsRecordedIterator is returned from FilterLogisticsRecorded and is used to iterate over the raw logs and unpacked data for LogisticsRecorded events raised by the Foodtrace contract.
type FoodtraceLogisticsRecordedIterator struct {
	Event *FoodtraceLogisticsRecorded // Event containing the contract specifics and raw log

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
func (it *FoodtraceLogisticsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodtraceLogisticsRecorded)
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
		it.Event = new(FoodtraceLogisticsRecorded)
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
func (it *FoodtraceLogisticsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodtraceLogisticsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodtraceLogisticsRecorded represents a LogisticsRecorded event raised by the Foodtrace contract.
type FoodtraceLogisticsRecorded struct {
	Id              *big.Int
	LogisticsRecord string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogisticsRecorded is a free log retrieval operation binding the contract event 0x528527ca056a90b130d12f4c6b3732c95f08f0f199704b6677b67a0bd62648a5.
//
// Solidity: event LogisticsRecorded(uint256 id, string logisticsRecord)
func (_Foodtrace *FoodtraceFilterer) FilterLogisticsRecorded(opts *bind.FilterOpts) (*FoodtraceLogisticsRecordedIterator, error) {

	logs, sub, err := _Foodtrace.contract.FilterLogs(opts, "LogisticsRecorded")
	if err != nil {
		return nil, err
	}
	return &FoodtraceLogisticsRecordedIterator{contract: _Foodtrace.contract, event: "LogisticsRecorded", logs: logs, sub: sub}, nil
}

// WatchLogisticsRecorded is a free log subscription operation binding the contract event 0x528527ca056a90b130d12f4c6b3732c95f08f0f199704b6677b67a0bd62648a5.
//
// Solidity: event LogisticsRecorded(uint256 id, string logisticsRecord)
func (_Foodtrace *FoodtraceFilterer) WatchLogisticsRecorded(opts *bind.WatchOpts, sink chan<- *FoodtraceLogisticsRecorded) (event.Subscription, error) {

	logs, sub, err := _Foodtrace.contract.WatchLogs(opts, "LogisticsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodtraceLogisticsRecorded)
				if err := _Foodtrace.contract.UnpackLog(event, "LogisticsRecorded", log); err != nil {
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

// ParseLogisticsRecorded is a log parse operation binding the contract event 0x528527ca056a90b130d12f4c6b3732c95f08f0f199704b6677b67a0bd62648a5.
//
// Solidity: event LogisticsRecorded(uint256 id, string logisticsRecord)
func (_Foodtrace *FoodtraceFilterer) ParseLogisticsRecorded(log types.Log) (*FoodtraceLogisticsRecorded, error) {
	event := new(FoodtraceLogisticsRecorded)
	if err := _Foodtrace.contract.UnpackLog(event, "LogisticsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
