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

// PancakeV3MetaData contains all meta data concerning the PancakeV3 contract.
var PancakeV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"FeeAmountExtraInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"SetLmPoolDeployer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"WhiteListAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacingExtraInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lmPoolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setFeeAmountExtraInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol0\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol1\",\"type\":\"uint32\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lmPool\",\"type\":\"address\"}],\"name\":\"setLmPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"setLmPoolDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"setWhiteListAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakeV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeV3MetaData.ABI instead.
var PancakeV3ABI = PancakeV3MetaData.ABI

// PancakeV3 is an auto generated Go binding around an Ethereum contract.
type PancakeV3 struct {
	PancakeV3Caller     // Read-only binding to the contract
	PancakeV3Transactor // Write-only binding to the contract
	PancakeV3Filterer   // Log filterer for contract events
}

// PancakeV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeV3Session struct {
	Contract     *PancakeV3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeV3CallerSession struct {
	Contract *PancakeV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PancakeV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeV3TransactorSession struct {
	Contract     *PancakeV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PancakeV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeV3Raw struct {
	Contract *PancakeV3 // Generic contract binding to access the raw methods on
}

// PancakeV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeV3CallerRaw struct {
	Contract *PancakeV3Caller // Generic read-only contract binding to access the raw methods on
}

// PancakeV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeV3TransactorRaw struct {
	Contract *PancakeV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeV3 creates a new instance of PancakeV3, bound to a specific deployed contract.
func NewPancakeV3(address common.Address, backend bind.ContractBackend) (*PancakeV3, error) {
	contract, err := bindPancakeV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeV3{PancakeV3Caller: PancakeV3Caller{contract: contract}, PancakeV3Transactor: PancakeV3Transactor{contract: contract}, PancakeV3Filterer: PancakeV3Filterer{contract: contract}}, nil
}

// NewPancakeV3Caller creates a new read-only instance of PancakeV3, bound to a specific deployed contract.
func NewPancakeV3Caller(address common.Address, caller bind.ContractCaller) (*PancakeV3Caller, error) {
	contract, err := bindPancakeV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV3Caller{contract: contract}, nil
}

// NewPancakeV3Transactor creates a new write-only instance of PancakeV3, bound to a specific deployed contract.
func NewPancakeV3Transactor(address common.Address, transactor bind.ContractTransactor) (*PancakeV3Transactor, error) {
	contract, err := bindPancakeV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV3Transactor{contract: contract}, nil
}

// NewPancakeV3Filterer creates a new log filterer instance of PancakeV3, bound to a specific deployed contract.
func NewPancakeV3Filterer(address common.Address, filterer bind.ContractFilterer) (*PancakeV3Filterer, error) {
	contract, err := bindPancakeV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeV3Filterer{contract: contract}, nil
}

// bindPancakeV3 binds a generic wrapper to an already deployed contract.
func bindPancakeV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV3 *PancakeV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV3.Contract.PancakeV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV3 *PancakeV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV3.Contract.PancakeV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV3 *PancakeV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV3.Contract.PancakeV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV3 *PancakeV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV3 *PancakeV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV3 *PancakeV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV3.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3 *PancakeV3Caller) FeeAmountTickSpacing(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeV3.contract.Call(opts, &out, "feeAmountTickSpacing", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3 *PancakeV3Session) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _PancakeV3.Contract.FeeAmountTickSpacing(&_PancakeV3.CallOpts, arg0)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3 *PancakeV3CallerSession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _PancakeV3.Contract.FeeAmountTickSpacing(&_PancakeV3.CallOpts, arg0)
}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3 *PancakeV3Caller) FeeAmountTickSpacingExtraInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	var out []interface{}
	err := _PancakeV3.contract.Call(opts, &out, "feeAmountTickSpacingExtraInfo", arg0)

	outstruct := new(struct {
		WhitelistRequested bool
		Enabled            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WhitelistRequested = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3 *PancakeV3Session) FeeAmountTickSpacingExtraInfo(arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	return _PancakeV3.Contract.FeeAmountTickSpacingExtraInfo(&_PancakeV3.CallOpts, arg0)
}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3 *PancakeV3CallerSession) FeeAmountTickSpacingExtraInfo(arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	return _PancakeV3.Contract.FeeAmountTickSpacingExtraInfo(&_PancakeV3.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3 *PancakeV3Caller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3 *PancakeV3Session) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _PancakeV3.Contract.GetPool(&_PancakeV3.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3 *PancakeV3CallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _PancakeV3.Contract.GetPool(&_PancakeV3.CallOpts, arg0, arg1, arg2)
}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3 *PancakeV3Caller) LmPoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3.contract.Call(opts, &out, "lmPoolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3 *PancakeV3Session) LmPoolDeployer() (common.Address, error) {
	return _PancakeV3.Contract.LmPoolDeployer(&_PancakeV3.CallOpts)
}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3 *PancakeV3CallerSession) LmPoolDeployer() (common.Address, error) {
	return _PancakeV3.Contract.LmPoolDeployer(&_PancakeV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3 *PancakeV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3 *PancakeV3Session) Owner() (common.Address, error) {
	return _PancakeV3.Contract.Owner(&_PancakeV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3 *PancakeV3CallerSession) Owner() (common.Address, error) {
	return _PancakeV3.Contract.Owner(&_PancakeV3.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3 *PancakeV3Caller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3 *PancakeV3Session) PoolDeployer() (common.Address, error) {
	return _PancakeV3.Contract.PoolDeployer(&_PancakeV3.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3 *PancakeV3CallerSession) PoolDeployer() (common.Address, error) {
	return _PancakeV3.Contract.PoolDeployer(&_PancakeV3.CallOpts)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3 *PancakeV3Transactor) CollectProtocol(opts *bind.TransactOpts, pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "collectProtocol", pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3 *PancakeV3Session) CollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3.Contract.CollectProtocol(&_PancakeV3.TransactOpts, pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3 *PancakeV3TransactorSession) CollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3.Contract.CollectProtocol(&_PancakeV3.TransactOpts, pool, recipient, amount0Requested, amount1Requested)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3 *PancakeV3Transactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3 *PancakeV3Session) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3.Contract.CreatePool(&_PancakeV3.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3 *PancakeV3TransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3.Contract.CreatePool(&_PancakeV3.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3 *PancakeV3Transactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3 *PancakeV3Session) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3.Contract.EnableFeeAmount(&_PancakeV3.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3 *PancakeV3TransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3.Contract.EnableFeeAmount(&_PancakeV3.TransactOpts, fee, tickSpacing)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3 *PancakeV3Transactor) SetFeeAmountExtraInfo(opts *bind.TransactOpts, fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "setFeeAmountExtraInfo", fee, whitelistRequested, enabled)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3 *PancakeV3Session) SetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetFeeAmountExtraInfo(&_PancakeV3.TransactOpts, fee, whitelistRequested, enabled)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3 *PancakeV3TransactorSession) SetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetFeeAmountExtraInfo(&_PancakeV3.TransactOpts, fee, whitelistRequested, enabled)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3 *PancakeV3Transactor) SetFeeProtocol(opts *bind.TransactOpts, pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "setFeeProtocol", pool, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3 *PancakeV3Session) SetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetFeeProtocol(&_PancakeV3.TransactOpts, pool, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3 *PancakeV3TransactorSession) SetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetFeeProtocol(&_PancakeV3.TransactOpts, pool, feeProtocol0, feeProtocol1)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3 *PancakeV3Transactor) SetLmPool(opts *bind.TransactOpts, pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "setLmPool", pool, lmPool)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3 *PancakeV3Session) SetLmPool(pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetLmPool(&_PancakeV3.TransactOpts, pool, lmPool)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3 *PancakeV3TransactorSession) SetLmPool(pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetLmPool(&_PancakeV3.TransactOpts, pool, lmPool)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3 *PancakeV3Transactor) SetLmPoolDeployer(opts *bind.TransactOpts, _lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "setLmPoolDeployer", _lmPoolDeployer)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3 *PancakeV3Session) SetLmPoolDeployer(_lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetLmPoolDeployer(&_PancakeV3.TransactOpts, _lmPoolDeployer)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3 *PancakeV3TransactorSession) SetLmPoolDeployer(_lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetLmPoolDeployer(&_PancakeV3.TransactOpts, _lmPoolDeployer)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3 *PancakeV3Transactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3 *PancakeV3Session) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetOwner(&_PancakeV3.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3 *PancakeV3TransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetOwner(&_PancakeV3.TransactOpts, _owner)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3 *PancakeV3Transactor) SetWhiteListAddress(opts *bind.TransactOpts, user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3.contract.Transact(opts, "setWhiteListAddress", user, verified)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3 *PancakeV3Session) SetWhiteListAddress(user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetWhiteListAddress(&_PancakeV3.TransactOpts, user, verified)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3 *PancakeV3TransactorSession) SetWhiteListAddress(user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3.Contract.SetWhiteListAddress(&_PancakeV3.TransactOpts, user, verified)
}

// PancakeV3FeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the PancakeV3 contract.
type PancakeV3FeeAmountEnabledIterator struct {
	Event *PancakeV3FeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *PancakeV3FeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FeeAmountEnabled)
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
		it.Event = new(PancakeV3FeeAmountEnabled)
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
func (it *PancakeV3FeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FeeAmountEnabled represents a FeeAmountEnabled event raised by the PancakeV3 contract.
type PancakeV3FeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3 *PancakeV3Filterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*PancakeV3FeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _PancakeV3.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FeeAmountEnabledIterator{contract: _PancakeV3.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3 *PancakeV3Filterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *PancakeV3FeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _PancakeV3.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FeeAmountEnabled)
				if err := _PancakeV3.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3 *PancakeV3Filterer) ParseFeeAmountEnabled(log types.Log) (*PancakeV3FeeAmountEnabled, error) {
	event := new(PancakeV3FeeAmountEnabled)
	if err := _PancakeV3.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3FeeAmountExtraInfoUpdatedIterator is returned from FilterFeeAmountExtraInfoUpdated and is used to iterate over the raw logs and unpacked data for FeeAmountExtraInfoUpdated events raised by the PancakeV3 contract.
type PancakeV3FeeAmountExtraInfoUpdatedIterator struct {
	Event *PancakeV3FeeAmountExtraInfoUpdated // Event containing the contract specifics and raw log

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
func (it *PancakeV3FeeAmountExtraInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3FeeAmountExtraInfoUpdated)
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
		it.Event = new(PancakeV3FeeAmountExtraInfoUpdated)
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
func (it *PancakeV3FeeAmountExtraInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3FeeAmountExtraInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3FeeAmountExtraInfoUpdated represents a FeeAmountExtraInfoUpdated event raised by the PancakeV3 contract.
type PancakeV3FeeAmountExtraInfoUpdated struct {
	Fee                *big.Int
	WhitelistRequested bool
	Enabled            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountExtraInfoUpdated is a free log retrieval operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3 *PancakeV3Filterer) FilterFeeAmountExtraInfoUpdated(opts *bind.FilterOpts, fee []*big.Int) (*PancakeV3FeeAmountExtraInfoUpdatedIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3.contract.FilterLogs(opts, "FeeAmountExtraInfoUpdated", feeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3FeeAmountExtraInfoUpdatedIterator{contract: _PancakeV3.contract, event: "FeeAmountExtraInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeAmountExtraInfoUpdated is a free log subscription operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3 *PancakeV3Filterer) WatchFeeAmountExtraInfoUpdated(opts *bind.WatchOpts, sink chan<- *PancakeV3FeeAmountExtraInfoUpdated, fee []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3.contract.WatchLogs(opts, "FeeAmountExtraInfoUpdated", feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3FeeAmountExtraInfoUpdated)
				if err := _PancakeV3.contract.UnpackLog(event, "FeeAmountExtraInfoUpdated", log); err != nil {
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

// ParseFeeAmountExtraInfoUpdated is a log parse operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3 *PancakeV3Filterer) ParseFeeAmountExtraInfoUpdated(log types.Log) (*PancakeV3FeeAmountExtraInfoUpdated, error) {
	event := new(PancakeV3FeeAmountExtraInfoUpdated)
	if err := _PancakeV3.contract.UnpackLog(event, "FeeAmountExtraInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PancakeV3 contract.
type PancakeV3OwnerChangedIterator struct {
	Event *PancakeV3OwnerChanged // Event containing the contract specifics and raw log

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
func (it *PancakeV3OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3OwnerChanged)
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
		it.Event = new(PancakeV3OwnerChanged)
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
func (it *PancakeV3OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3OwnerChanged represents a OwnerChanged event raised by the PancakeV3 contract.
type PancakeV3OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3 *PancakeV3Filterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PancakeV3OwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeV3.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3OwnerChangedIterator{contract: _PancakeV3.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3 *PancakeV3Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PancakeV3OwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeV3.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3OwnerChanged)
				if err := _PancakeV3.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3 *PancakeV3Filterer) ParseOwnerChanged(log types.Log) (*PancakeV3OwnerChanged, error) {
	event := new(PancakeV3OwnerChanged)
	if err := _PancakeV3.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3PoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PancakeV3 contract.
type PancakeV3PoolCreatedIterator struct {
	Event *PancakeV3PoolCreated // Event containing the contract specifics and raw log

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
func (it *PancakeV3PoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3PoolCreated)
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
		it.Event = new(PancakeV3PoolCreated)
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
func (it *PancakeV3PoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3PoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3PoolCreated represents a PoolCreated event raised by the PancakeV3 contract.
type PancakeV3PoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3 *PancakeV3Filterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*PancakeV3PoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3PoolCreatedIterator{contract: _PancakeV3.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3 *PancakeV3Filterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PancakeV3PoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3PoolCreated)
				if err := _PancakeV3.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3 *PancakeV3Filterer) ParsePoolCreated(log types.Log) (*PancakeV3PoolCreated, error) {
	event := new(PancakeV3PoolCreated)
	if err := _PancakeV3.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3SetLmPoolDeployerIterator is returned from FilterSetLmPoolDeployer and is used to iterate over the raw logs and unpacked data for SetLmPoolDeployer events raised by the PancakeV3 contract.
type PancakeV3SetLmPoolDeployerIterator struct {
	Event *PancakeV3SetLmPoolDeployer // Event containing the contract specifics and raw log

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
func (it *PancakeV3SetLmPoolDeployerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3SetLmPoolDeployer)
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
		it.Event = new(PancakeV3SetLmPoolDeployer)
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
func (it *PancakeV3SetLmPoolDeployerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3SetLmPoolDeployerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3SetLmPoolDeployer represents a SetLmPoolDeployer event raised by the PancakeV3 contract.
type PancakeV3SetLmPoolDeployer struct {
	LmPoolDeployer common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetLmPoolDeployer is a free log retrieval operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3 *PancakeV3Filterer) FilterSetLmPoolDeployer(opts *bind.FilterOpts, lmPoolDeployer []common.Address) (*PancakeV3SetLmPoolDeployerIterator, error) {

	var lmPoolDeployerRule []interface{}
	for _, lmPoolDeployerItem := range lmPoolDeployer {
		lmPoolDeployerRule = append(lmPoolDeployerRule, lmPoolDeployerItem)
	}

	logs, sub, err := _PancakeV3.contract.FilterLogs(opts, "SetLmPoolDeployer", lmPoolDeployerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3SetLmPoolDeployerIterator{contract: _PancakeV3.contract, event: "SetLmPoolDeployer", logs: logs, sub: sub}, nil
}

// WatchSetLmPoolDeployer is a free log subscription operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3 *PancakeV3Filterer) WatchSetLmPoolDeployer(opts *bind.WatchOpts, sink chan<- *PancakeV3SetLmPoolDeployer, lmPoolDeployer []common.Address) (event.Subscription, error) {

	var lmPoolDeployerRule []interface{}
	for _, lmPoolDeployerItem := range lmPoolDeployer {
		lmPoolDeployerRule = append(lmPoolDeployerRule, lmPoolDeployerItem)
	}

	logs, sub, err := _PancakeV3.contract.WatchLogs(opts, "SetLmPoolDeployer", lmPoolDeployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3SetLmPoolDeployer)
				if err := _PancakeV3.contract.UnpackLog(event, "SetLmPoolDeployer", log); err != nil {
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

// ParseSetLmPoolDeployer is a log parse operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3 *PancakeV3Filterer) ParseSetLmPoolDeployer(log types.Log) (*PancakeV3SetLmPoolDeployer, error) {
	event := new(PancakeV3SetLmPoolDeployer)
	if err := _PancakeV3.contract.UnpackLog(event, "SetLmPoolDeployer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3WhiteListAddedIterator is returned from FilterWhiteListAdded and is used to iterate over the raw logs and unpacked data for WhiteListAdded events raised by the PancakeV3 contract.
type PancakeV3WhiteListAddedIterator struct {
	Event *PancakeV3WhiteListAdded // Event containing the contract specifics and raw log

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
func (it *PancakeV3WhiteListAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3WhiteListAdded)
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
		it.Event = new(PancakeV3WhiteListAdded)
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
func (it *PancakeV3WhiteListAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3WhiteListAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3WhiteListAdded represents a WhiteListAdded event raised by the PancakeV3 contract.
type PancakeV3WhiteListAdded struct {
	User     common.Address
	Verified bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhiteListAdded is a free log retrieval operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3 *PancakeV3Filterer) FilterWhiteListAdded(opts *bind.FilterOpts, user []common.Address) (*PancakeV3WhiteListAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeV3.contract.FilterLogs(opts, "WhiteListAdded", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3WhiteListAddedIterator{contract: _PancakeV3.contract, event: "WhiteListAdded", logs: logs, sub: sub}, nil
}

// WatchWhiteListAdded is a free log subscription operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3 *PancakeV3Filterer) WatchWhiteListAdded(opts *bind.WatchOpts, sink chan<- *PancakeV3WhiteListAdded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeV3.contract.WatchLogs(opts, "WhiteListAdded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3WhiteListAdded)
				if err := _PancakeV3.contract.UnpackLog(event, "WhiteListAdded", log); err != nil {
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

// ParseWhiteListAdded is a log parse operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3 *PancakeV3Filterer) ParseWhiteListAdded(log types.Log) (*PancakeV3WhiteListAdded, error) {
	event := new(PancakeV3WhiteListAdded)
	if err := _PancakeV3.contract.UnpackLog(event, "WhiteListAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
