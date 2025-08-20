// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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

// PancakeV2MetaData contains all meta data concerning the PancakeV2 contract.
var PancakeV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_CODE_PAIR_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeV2MetaData.ABI instead.
var PancakeV2ABI = PancakeV2MetaData.ABI

// PancakeV2 is an auto generated Go binding around an Ethereum contract.
type PancakeV2 struct {
	PancakeV2Caller     // Read-only binding to the contract
	PancakeV2Transactor // Write-only binding to the contract
	PancakeV2Filterer   // Log filterer for contract events
}

// PancakeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeV2Session struct {
	Contract     *PancakeV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeV2CallerSession struct {
	Contract *PancakeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PancakeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeV2TransactorSession struct {
	Contract     *PancakeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PancakeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeV2Raw struct {
	Contract *PancakeV2 // Generic contract binding to access the raw methods on
}

// PancakeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeV2CallerRaw struct {
	Contract *PancakeV2Caller // Generic read-only contract binding to access the raw methods on
}

// PancakeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeV2TransactorRaw struct {
	Contract *PancakeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeV2 creates a new instance of PancakeV2, bound to a specific deployed contract.
func NewPancakeV2(address common.Address, backend bind.ContractBackend) (*PancakeV2, error) {
	contract, err := bindPancakeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeV2{PancakeV2Caller: PancakeV2Caller{contract: contract}, PancakeV2Transactor: PancakeV2Transactor{contract: contract}, PancakeV2Filterer: PancakeV2Filterer{contract: contract}}, nil
}

// NewPancakeV2Caller creates a new read-only instance of PancakeV2, bound to a specific deployed contract.
func NewPancakeV2Caller(address common.Address, caller bind.ContractCaller) (*PancakeV2Caller, error) {
	contract, err := bindPancakeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV2Caller{contract: contract}, nil
}

// NewPancakeV2Transactor creates a new write-only instance of PancakeV2, bound to a specific deployed contract.
func NewPancakeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*PancakeV2Transactor, error) {
	contract, err := bindPancakeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV2Transactor{contract: contract}, nil
}

// NewPancakeV2Filterer creates a new log filterer instance of PancakeV2, bound to a specific deployed contract.
func NewPancakeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*PancakeV2Filterer, error) {
	contract, err := bindPancakeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeV2Filterer{contract: contract}, nil
}

// bindPancakeV2 binds a generic wrapper to an already deployed contract.
func bindPancakeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV2 *PancakeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV2.Contract.PancakeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV2 *PancakeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV2.Contract.PancakeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV2 *PancakeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV2.Contract.PancakeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV2 *PancakeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV2 *PancakeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV2 *PancakeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV2.Contract.contract.Transact(opts, method, params...)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PancakeV2 *PancakeV2Caller) INITCODEPAIRHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PancakeV2.contract.Call(opts, &out, "INIT_CODE_PAIR_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PancakeV2 *PancakeV2Session) INITCODEPAIRHASH() ([32]byte, error) {
	return _PancakeV2.Contract.INITCODEPAIRHASH(&_PancakeV2.CallOpts)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PancakeV2 *PancakeV2CallerSession) INITCODEPAIRHASH() ([32]byte, error) {
	return _PancakeV2.Contract.INITCODEPAIRHASH(&_PancakeV2.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PancakeV2 *PancakeV2Caller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PancakeV2.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PancakeV2 *PancakeV2Session) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _PancakeV2.Contract.AllPairs(&_PancakeV2.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PancakeV2 *PancakeV2CallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _PancakeV2.Contract.AllPairs(&_PancakeV2.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PancakeV2 *PancakeV2Caller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeV2.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PancakeV2 *PancakeV2Session) AllPairsLength() (*big.Int, error) {
	return _PancakeV2.Contract.AllPairsLength(&_PancakeV2.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PancakeV2 *PancakeV2CallerSession) AllPairsLength() (*big.Int, error) {
	return _PancakeV2.Contract.AllPairsLength(&_PancakeV2.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PancakeV2 *PancakeV2Caller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV2.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PancakeV2 *PancakeV2Session) FeeTo() (common.Address, error) {
	return _PancakeV2.Contract.FeeTo(&_PancakeV2.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PancakeV2 *PancakeV2CallerSession) FeeTo() (common.Address, error) {
	return _PancakeV2.Contract.FeeTo(&_PancakeV2.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PancakeV2 *PancakeV2Caller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV2.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PancakeV2 *PancakeV2Session) FeeToSetter() (common.Address, error) {
	return _PancakeV2.Contract.FeeToSetter(&_PancakeV2.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PancakeV2 *PancakeV2CallerSession) FeeToSetter() (common.Address, error) {
	return _PancakeV2.Contract.FeeToSetter(&_PancakeV2.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PancakeV2 *PancakeV2Caller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PancakeV2.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PancakeV2 *PancakeV2Session) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PancakeV2.Contract.GetPair(&_PancakeV2.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PancakeV2 *PancakeV2CallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PancakeV2.Contract.GetPair(&_PancakeV2.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PancakeV2 *PancakeV2Transactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PancakeV2.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PancakeV2 *PancakeV2Session) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PancakeV2.Contract.CreatePair(&_PancakeV2.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PancakeV2 *PancakeV2TransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PancakeV2.Contract.CreatePair(&_PancakeV2.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PancakeV2 *PancakeV2Transactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _PancakeV2.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PancakeV2 *PancakeV2Session) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _PancakeV2.Contract.SetFeeTo(&_PancakeV2.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PancakeV2 *PancakeV2TransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _PancakeV2.Contract.SetFeeTo(&_PancakeV2.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PancakeV2 *PancakeV2Transactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _PancakeV2.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PancakeV2 *PancakeV2Session) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _PancakeV2.Contract.SetFeeToSetter(&_PancakeV2.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PancakeV2 *PancakeV2TransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _PancakeV2.Contract.SetFeeToSetter(&_PancakeV2.TransactOpts, _feeToSetter)
}

// PancakeV2PairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the PancakeV2 contract.
type PancakeV2PairCreatedIterator struct {
	Event *PancakeV2PairCreated // Event containing the contract specifics and raw log

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
func (it *PancakeV2PairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV2PairCreated)
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
		it.Event = new(PancakeV2PairCreated)
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
func (it *PancakeV2PairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV2PairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV2PairCreated represents a PairCreated event raised by the PancakeV2 contract.
type PancakeV2PairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PancakeV2 *PancakeV2Filterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*PancakeV2PairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _PancakeV2.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &PancakeV2PairCreatedIterator{contract: _PancakeV2.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PancakeV2 *PancakeV2Filterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *PancakeV2PairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _PancakeV2.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV2PairCreated)
				if err := _PancakeV2.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PancakeV2 *PancakeV2Filterer) ParsePairCreated(log types.Log) (*PancakeV2PairCreated, error) {
	event := new(PancakeV2PairCreated)
	if err := _PancakeV2.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
