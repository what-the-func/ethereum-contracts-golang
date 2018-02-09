// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TronTokenABI is the input ABI used to generate the binding from.
const TronTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addressFounder\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TronTokenBin is the compiled bytecode used for deploying new contracts.
const TronTokenBin = `0x606060405260408051908101604052600681527f54726f6e697800000000000000000000000000000000000000000000000000006020820152600090805161004b929160200190610155565b5060408051908101604052600381527f545258000000000000000000000000000000000000000000000000000000000060208201526001908051610093929160200190610155565b506006600281905560006005558054600160a860020a031916905534156100b957600080fd5b604051602080610b288339810160405280805160068054600160a060020a033381166101000261010060a860020a03199092169190911790915567016345785d8a00006005819055908216600081815260036020526040808220849055939550909350917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91905190815260200160405180910390a3506101f0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061019657805160ff19168380011785556101c3565b828001600101855582156101c3579182015b828111156101c35782518255916020019190600101906101a8565b506101cf9291506101d3565b5090565b6101ed91905b808211156101cf57600081556001016101d9565b90565b610929806101ff6000396000f3006060604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d457806307da68f51461015e578063095ea7b31461017357806318160ddd146101a957806323b872dd146101ce578063313ce567146101f657806342966c681461020957806370a082311461021f57806375f12b211461023e57806395d89b4114610251578063a9059cbb14610264578063be9a655514610286578063c47f002714610299578063dd62ed3e146102ea575b600080fd5b34156100df57600080fd5b6100e761030f565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012357808201518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016957600080fd5b6101716103ad565b005b341561017e57600080fd5b610195600160a060020a03600435166024356103d9565b604051901515815260200160405180910390f35b34156101b457600080fd5b6101bc61049f565b60405190815260200160405180910390f35b34156101d957600080fd5b610195600160a060020a03600435811690602435166044356104a5565b341561020157600080fd5b6101bc6105d6565b341561021457600080fd5b6101716004356105dc565b341561022a57600080fd5b6101bc600160a060020a0360043516610685565b341561024957600080fd5b610195610697565b341561025c57600080fd5b6100e76106a0565b341561026f57600080fd5b610195600160a060020a036004351660243561070b565b341561029157600080fd5b6101716107e8565b34156102a457600080fd5b61017160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061081195505050505050565b34156102f557600080fd5b6101bc600160a060020a0360043581169060243516610845565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a55780601f1061037a576101008083540402835291602001916103a5565b820191906000526020600020905b81548152906001019060200180831161038857829003601f168201915b505050505081565b60065433600160a060020a0390811661010090920416146103ca57fe5b6006805460ff19166001179055565b60065460009060ff16156103e957fe5b600160a060020a03331615156103fb57fe5b81158061042b5750600160a060020a03338116600090815260046020908152604080832093871683529290522054155b151561043657600080fd5b600160a060020a03338116600081815260046020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60055481565b60065460009060ff16156104b557fe5b600160a060020a03331615156104c757fe5b600160a060020a038416600090815260036020526040902054829010156104ed57600080fd5b600160a060020a038316600090815260036020526040902054828101101561051457600080fd5b600160a060020a03808516600090815260046020908152604080832033909416835292905220548290101561054857600080fd5b600160a060020a03808416600081815260036020908152604080832080548801905588851680845281842080548990039055600483528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60025481565b600160a060020a0333166000908152600360205260409020548190101561060257600080fd5b600160a060020a033316600081815260036020526040808220805485900390558180527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff8054850190559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9084905190815260200160405180910390a350565b60036020526000908152604090205481565b60065460ff1681565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103a55780601f1061037a576101008083540402835291602001916103a5565b60065460009060ff161561071b57fe5b600160a060020a033316151561072d57fe5b600160a060020a0333166000908152600360205260409020548290101561075357600080fd5b600160a060020a038316600090815260036020526040902054828101101561077a57600080fd5b600160a060020a033381166000818152600360205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b60065433600160a060020a03908116610100909204161461080557fe5b6006805460ff19169055565b60065433600160a060020a03908116610100909204161461082e57fe5b6000818051610841929160200190610862565b5050565b600460209081526000928352604080842090915290825290205481565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108a357805160ff19168380011785556108d0565b828001600101855582156108d0579182015b828111156108d05782518255916020019190600101906108b5565b506108dc9291506108e0565b5090565b6108fa91905b808211156108dc57600081556001016108e6565b905600a165627a7a72305820a67dd93f272d17e75b4d0df2e65eb6a908f5c33c0c17b8ee999ca39d7d5cdb890029`

// DeployTronToken deploys a new Ethereum contract, binding an instance of TronToken to it.
func DeployTronToken(auth *bind.TransactOpts, backend bind.ContractBackend, _addressFounder common.Address) (common.Address, *types.Transaction, *TronToken, error) {
	parsed, err := abi.JSON(strings.NewReader(TronTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TronTokenBin), backend, _addressFounder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TronToken{TronTokenCaller: TronTokenCaller{contract: contract}, TronTokenTransactor: TronTokenTransactor{contract: contract}}, nil
}

// TronToken is an auto generated Go binding around an Ethereum contract.
type TronToken struct {
	TronTokenCaller     // Read-only binding to the contract
	TronTokenTransactor // Write-only binding to the contract
}

// TronTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TronTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TronTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TronTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TronTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TronTokenSession struct {
	Contract     *TronToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TronTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TronTokenCallerSession struct {
	Contract *TronTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TronTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TronTokenTransactorSession struct {
	Contract     *TronTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TronTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TronTokenRaw struct {
	Contract *TronToken // Generic contract binding to access the raw methods on
}

// TronTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TronTokenCallerRaw struct {
	Contract *TronTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TronTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TronTokenTransactorRaw struct {
	Contract *TronTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTronToken creates a new instance of TronToken, bound to a specific deployed contract.
func NewTronToken(address common.Address, backend bind.ContractBackend) (*TronToken, error) {
	contract, err := bindTronToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TronToken{TronTokenCaller: TronTokenCaller{contract: contract}, TronTokenTransactor: TronTokenTransactor{contract: contract}}, nil
}

// NewTronTokenCaller creates a new read-only instance of TronToken, bound to a specific deployed contract.
func NewTronTokenCaller(address common.Address, caller bind.ContractCaller) (*TronTokenCaller, error) {
	contract, err := bindTronToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TronTokenCaller{contract: contract}, nil
}

// NewTronTokenTransactor creates a new write-only instance of TronToken, bound to a specific deployed contract.
func NewTronTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TronTokenTransactor, error) {
	contract, err := bindTronToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TronTokenTransactor{contract: contract}, nil
}

// bindTronToken binds a generic wrapper to an already deployed contract.
func bindTronToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TronTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TronToken *TronTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TronToken.Contract.TronTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TronToken *TronTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.Contract.TronTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TronToken *TronTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TronToken.Contract.TronTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TronToken *TronTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TronToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TronToken *TronTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TronToken *TronTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TronToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TronToken *TronTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TronToken *TronTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TronToken.Contract.Allowance(&_TronToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TronToken *TronTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TronToken.Contract.Allowance(&_TronToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TronToken *TronTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TronToken *TronTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TronToken.Contract.BalanceOf(&_TronToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TronToken *TronTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TronToken.Contract.BalanceOf(&_TronToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TronToken *TronTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TronToken *TronTokenSession) Decimals() (*big.Int, error) {
	return _TronToken.Contract.Decimals(&_TronToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TronToken *TronTokenCallerSession) Decimals() (*big.Int, error) {
	return _TronToken.Contract.Decimals(&_TronToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TronToken *TronTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TronToken *TronTokenSession) Name() (string, error) {
	return _TronToken.Contract.Name(&_TronToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TronToken *TronTokenCallerSession) Name() (string, error) {
	return _TronToken.Contract.Name(&_TronToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TronToken *TronTokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TronToken *TronTokenSession) Stopped() (bool, error) {
	return _TronToken.Contract.Stopped(&_TronToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TronToken *TronTokenCallerSession) Stopped() (bool, error) {
	return _TronToken.Contract.Stopped(&_TronToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TronToken *TronTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TronToken *TronTokenSession) Symbol() (string, error) {
	return _TronToken.Contract.Symbol(&_TronToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TronToken *TronTokenCallerSession) Symbol() (string, error) {
	return _TronToken.Contract.Symbol(&_TronToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TronToken *TronTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TronToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TronToken *TronTokenSession) TotalSupply() (*big.Int, error) {
	return _TronToken.Contract.TotalSupply(&_TronToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TronToken *TronTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TronToken.Contract.TotalSupply(&_TronToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TronToken *TronTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TronToken *TronTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Approve(&_TronToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TronToken *TronTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Approve(&_TronToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_TronToken *TronTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_TronToken *TronTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Burn(&_TronToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_TronToken *TronTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Burn(&_TronToken.TransactOpts, _value)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_TronToken *TronTokenTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_TronToken *TronTokenSession) SetName(_name string) (*types.Transaction, error) {
	return _TronToken.Contract.SetName(&_TronToken.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_TronToken *TronTokenTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _TronToken.Contract.SetName(&_TronToken.TransactOpts, _name)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_TronToken *TronTokenTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_TronToken *TronTokenSession) Start() (*types.Transaction, error) {
	return _TronToken.Contract.Start(&_TronToken.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_TronToken *TronTokenTransactorSession) Start() (*types.Transaction, error) {
	return _TronToken.Contract.Start(&_TronToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TronToken *TronTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TronToken *TronTokenSession) Stop() (*types.Transaction, error) {
	return _TronToken.Contract.Stop(&_TronToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TronToken *TronTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _TronToken.Contract.Stop(&_TronToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_TronToken *TronTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_TronToken *TronTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Transfer(&_TronToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_TronToken *TronTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.Transfer(&_TronToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TronToken *TronTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TronToken *TronTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.TransferFrom(&_TronToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TronToken *TronTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TronToken.Contract.TransferFrom(&_TronToken.TransactOpts, _from, _to, _value)
}
