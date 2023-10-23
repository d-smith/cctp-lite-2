// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fiddy

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

// FiddyMetaData contains all meta data concerning the Fiddy contract.
var FiddyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000d5338038062000d538339810160408190526200003491620002b4565b8060405180604001604052806009815260200168119a59191e50d95b9d60ba1b815250604051806040016040528060048152602001634644444360e01b81525081600390816200008591906200038b565b5060046200009482826200038b565b5050506001600160a01b038116620000c757604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b620000d281620000f1565b50620000ea816a084595161401484a00000062000143565b506200047f565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0382166200016f5760405163ec442f0560e01b815260006004820152602401620000be565b6200017d6000838362000181565b5050565b6001600160a01b038316620001b0578060026000828254620001a4919062000457565b90915550620002249050565b6001600160a01b03831660009081526020819052604090205481811015620002055760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000be565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216620002425760028054829003905562000261565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002a791815260200190565b60405180910390a3505050565b600060208284031215620002c757600080fd5b81516001600160a01b0381168114620002df57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200031157607f821691505b6020821081036200033257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200038657600081815260208120601f850160051c81016020861015620003615750805b601f850160051c820191505b8181101562000382578281556001016200036d565b5050505b505050565b81516001600160401b03811115620003a757620003a7620002e6565b620003bf81620003b88454620002fc565b8462000338565b602080601f831160018114620003f75760008415620003de5750858301515b600019600386901b1c1916600185901b17855562000382565b600085815260208120601f198616915b82811015620004285788860151825594840194600190910190840162000407565b5085821015620004475787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200047957634e487b7160e01b600052601160045260246000fd5b92915050565b6108c4806200048f6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806370a082311161008c57806395d89b411161006657806395d89b41146101aa578063a9059cbb146101b2578063dd62ed3e146101c5578063f2fde38b146101fe57600080fd5b806370a082311461015e578063715018a6146101875780638da5cb5b1461018f57600080fd5b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461011557806323b872dd14610127578063313ce5671461013a57806340c10f1914610149575b600080fd5b6100dc610211565b6040516100e9919061070e565b60405180910390f35b610105610100366004610778565b6102a3565b60405190151581526020016100e9565b6002545b6040519081526020016100e9565b6101056101353660046107a2565b6102bd565b604051601281526020016100e9565b61015c610157366004610778565b6102e1565b005b61011961016c3660046107de565b6001600160a01b031660009081526020819052604090205490565b61015c6102f7565b6005546040516001600160a01b0390911681526020016100e9565b6100dc61030b565b6101056101c0366004610778565b61031a565b6101196101d3366004610800565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61015c61020c3660046107de565b610328565b60606003805461022090610833565b80601f016020809104026020016040519081016040528092919081815260200182805461024c90610833565b80156102995780601f1061026e57610100808354040283529160200191610299565b820191906000526020600020905b81548152906001019060200180831161027c57829003601f168201915b5050505050905090565b6000336102b181858561036b565b60019150505b92915050565b6000336102cb85828561037d565b6102d68585856103fb565b506001949350505050565b6102e961045a565b6102f38282610487565b5050565b6102ff61045a565b61030960006104bd565b565b60606004805461022090610833565b6000336102b18185856103fb565b61033061045a565b6001600160a01b03811661035f57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b610368816104bd565b50565b610378838383600161050f565b505050565b6001600160a01b0383811660009081526001602090815260408083209386168352929052205460001981146103f557818110156103e657604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610356565b6103f58484848403600061050f565b50505050565b6001600160a01b03831661042557604051634b637e8f60e11b815260006004820152602401610356565b6001600160a01b03821661044f5760405163ec442f0560e01b815260006004820152602401610356565b6103788383836105e4565b6005546001600160a01b031633146103095760405163118cdaa760e01b8152336004820152602401610356565b6001600160a01b0382166104b15760405163ec442f0560e01b815260006004820152602401610356565b6102f3600083836105e4565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0384166105395760405163e602df0560e01b815260006004820152602401610356565b6001600160a01b03831661056357604051634a1406b160e11b815260006004820152602401610356565b6001600160a01b03808516600090815260016020908152604080832093871683529290522082905580156103f557826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105d691815260200190565b60405180910390a350505050565b6001600160a01b03831661060f578060026000828254610604919061086d565b909155506106819050565b6001600160a01b038316600090815260208190526040902054818110156106625760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610356565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b03821661069d576002805482900390556106bc565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161070191815260200190565b60405180910390a3505050565b600060208083528351808285015260005b8181101561073b5785810183015185820160400152820161071f565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461077357600080fd5b919050565b6000806040838503121561078b57600080fd5b6107948361075c565b946020939093013593505050565b6000806000606084860312156107b757600080fd5b6107c08461075c565b92506107ce6020850161075c565b9150604084013590509250925092565b6000602082840312156107f057600080fd5b6107f98261075c565b9392505050565b6000806040838503121561081357600080fd5b61081c8361075c565b915061082a6020840161075c565b90509250929050565b600181811c9082168061084757607f821691505b60208210810361086757634e487b7160e01b600052602260045260246000fd5b50919050565b808201808211156102b757634e487b7160e01b600052601160045260246000fdfea2646970667358221220a6424f3b77426a19594813a954ef9a7352d464f665e7aee2ad90cb9eae23ba7664736f6c63430008140033",
}

// FiddyABI is the input ABI used to generate the binding from.
// Deprecated: Use FiddyMetaData.ABI instead.
var FiddyABI = FiddyMetaData.ABI

// FiddyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiddyMetaData.Bin instead.
var FiddyBin = FiddyMetaData.Bin

// DeployFiddy deploys a new Ethereum contract, binding an instance of Fiddy to it.
func DeployFiddy(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address) (common.Address, *types.Transaction, *Fiddy, error) {
	parsed, err := FiddyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiddyBin), backend, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fiddy{FiddyCaller: FiddyCaller{contract: contract}, FiddyTransactor: FiddyTransactor{contract: contract}, FiddyFilterer: FiddyFilterer{contract: contract}}, nil
}

// Fiddy is an auto generated Go binding around an Ethereum contract.
type Fiddy struct {
	FiddyCaller     // Read-only binding to the contract
	FiddyTransactor // Write-only binding to the contract
	FiddyFilterer   // Log filterer for contract events
}

// FiddyCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiddyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiddyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiddyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiddyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiddyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiddySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiddySession struct {
	Contract     *Fiddy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FiddyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiddyCallerSession struct {
	Contract *FiddyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FiddyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiddyTransactorSession struct {
	Contract     *FiddyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FiddyRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiddyRaw struct {
	Contract *Fiddy // Generic contract binding to access the raw methods on
}

// FiddyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiddyCallerRaw struct {
	Contract *FiddyCaller // Generic read-only contract binding to access the raw methods on
}

// FiddyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiddyTransactorRaw struct {
	Contract *FiddyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiddy creates a new instance of Fiddy, bound to a specific deployed contract.
func NewFiddy(address common.Address, backend bind.ContractBackend) (*Fiddy, error) {
	contract, err := bindFiddy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fiddy{FiddyCaller: FiddyCaller{contract: contract}, FiddyTransactor: FiddyTransactor{contract: contract}, FiddyFilterer: FiddyFilterer{contract: contract}}, nil
}

// NewFiddyCaller creates a new read-only instance of Fiddy, bound to a specific deployed contract.
func NewFiddyCaller(address common.Address, caller bind.ContractCaller) (*FiddyCaller, error) {
	contract, err := bindFiddy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiddyCaller{contract: contract}, nil
}

// NewFiddyTransactor creates a new write-only instance of Fiddy, bound to a specific deployed contract.
func NewFiddyTransactor(address common.Address, transactor bind.ContractTransactor) (*FiddyTransactor, error) {
	contract, err := bindFiddy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiddyTransactor{contract: contract}, nil
}

// NewFiddyFilterer creates a new log filterer instance of Fiddy, bound to a specific deployed contract.
func NewFiddyFilterer(address common.Address, filterer bind.ContractFilterer) (*FiddyFilterer, error) {
	contract, err := bindFiddy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiddyFilterer{contract: contract}, nil
}

// bindFiddy binds a generic wrapper to an already deployed contract.
func bindFiddy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FiddyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fiddy *FiddyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fiddy.Contract.FiddyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fiddy *FiddyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fiddy.Contract.FiddyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fiddy *FiddyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fiddy.Contract.FiddyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fiddy *FiddyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fiddy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fiddy *FiddyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fiddy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fiddy *FiddyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fiddy.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Fiddy *FiddyCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Fiddy *FiddySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Fiddy.Contract.Allowance(&_Fiddy.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Fiddy *FiddyCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Fiddy.Contract.Allowance(&_Fiddy.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Fiddy *FiddyCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Fiddy *FiddySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Fiddy.Contract.BalanceOf(&_Fiddy.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Fiddy *FiddyCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Fiddy.Contract.BalanceOf(&_Fiddy.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Fiddy *FiddyCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Fiddy *FiddySession) Decimals() (uint8, error) {
	return _Fiddy.Contract.Decimals(&_Fiddy.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Fiddy *FiddyCallerSession) Decimals() (uint8, error) {
	return _Fiddy.Contract.Decimals(&_Fiddy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fiddy *FiddyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fiddy *FiddySession) Name() (string, error) {
	return _Fiddy.Contract.Name(&_Fiddy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fiddy *FiddyCallerSession) Name() (string, error) {
	return _Fiddy.Contract.Name(&_Fiddy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fiddy *FiddyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fiddy *FiddySession) Owner() (common.Address, error) {
	return _Fiddy.Contract.Owner(&_Fiddy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fiddy *FiddyCallerSession) Owner() (common.Address, error) {
	return _Fiddy.Contract.Owner(&_Fiddy.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fiddy *FiddyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fiddy *FiddySession) Symbol() (string, error) {
	return _Fiddy.Contract.Symbol(&_Fiddy.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fiddy *FiddyCallerSession) Symbol() (string, error) {
	return _Fiddy.Contract.Symbol(&_Fiddy.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fiddy *FiddyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fiddy.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fiddy *FiddySession) TotalSupply() (*big.Int, error) {
	return _Fiddy.Contract.TotalSupply(&_Fiddy.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fiddy *FiddyCallerSession) TotalSupply() (*big.Int, error) {
	return _Fiddy.Contract.TotalSupply(&_Fiddy.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Fiddy *FiddyTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Fiddy *FiddySession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.Approve(&_Fiddy.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Fiddy *FiddyTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.Approve(&_Fiddy.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Fiddy *FiddyTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fiddy.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Fiddy *FiddySession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.Mint(&_Fiddy.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Fiddy *FiddyTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.Mint(&_Fiddy.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fiddy *FiddyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fiddy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fiddy *FiddySession) RenounceOwnership() (*types.Transaction, error) {
	return _Fiddy.Contract.RenounceOwnership(&_Fiddy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fiddy *FiddyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Fiddy.Contract.RenounceOwnership(&_Fiddy.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Fiddy *FiddyTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Fiddy *FiddySession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.Transfer(&_Fiddy.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Fiddy *FiddyTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.Transfer(&_Fiddy.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Fiddy *FiddyTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Fiddy *FiddySession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.TransferFrom(&_Fiddy.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Fiddy *FiddyTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Fiddy.Contract.TransferFrom(&_Fiddy.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fiddy *FiddyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Fiddy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fiddy *FiddySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fiddy.Contract.TransferOwnership(&_Fiddy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fiddy *FiddyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fiddy.Contract.TransferOwnership(&_Fiddy.TransactOpts, newOwner)
}

// FiddyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Fiddy contract.
type FiddyApprovalIterator struct {
	Event *FiddyApproval // Event containing the contract specifics and raw log

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
func (it *FiddyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiddyApproval)
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
		it.Event = new(FiddyApproval)
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
func (it *FiddyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiddyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiddyApproval represents a Approval event raised by the Fiddy contract.
type FiddyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Fiddy *FiddyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FiddyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Fiddy.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FiddyApprovalIterator{contract: _Fiddy.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Fiddy *FiddyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FiddyApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Fiddy.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiddyApproval)
				if err := _Fiddy.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Fiddy *FiddyFilterer) ParseApproval(log types.Log) (*FiddyApproval, error) {
	event := new(FiddyApproval)
	if err := _Fiddy.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiddyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Fiddy contract.
type FiddyOwnershipTransferredIterator struct {
	Event *FiddyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FiddyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiddyOwnershipTransferred)
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
		it.Event = new(FiddyOwnershipTransferred)
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
func (it *FiddyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiddyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiddyOwnershipTransferred represents a OwnershipTransferred event raised by the Fiddy contract.
type FiddyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Fiddy *FiddyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiddyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fiddy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiddyOwnershipTransferredIterator{contract: _Fiddy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Fiddy *FiddyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FiddyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fiddy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiddyOwnershipTransferred)
				if err := _Fiddy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Fiddy *FiddyFilterer) ParseOwnershipTransferred(log types.Log) (*FiddyOwnershipTransferred, error) {
	event := new(FiddyOwnershipTransferred)
	if err := _Fiddy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiddyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Fiddy contract.
type FiddyTransferIterator struct {
	Event *FiddyTransfer // Event containing the contract specifics and raw log

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
func (it *FiddyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiddyTransfer)
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
		it.Event = new(FiddyTransfer)
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
func (it *FiddyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiddyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiddyTransfer represents a Transfer event raised by the Fiddy contract.
type FiddyTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Fiddy *FiddyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FiddyTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Fiddy.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FiddyTransferIterator{contract: _Fiddy.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Fiddy *FiddyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FiddyTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Fiddy.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiddyTransfer)
				if err := _Fiddy.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Fiddy *FiddyFilterer) ParseTransfer(log types.Log) (*FiddyTransfer, error) {
	event := new(FiddyTransfer)
	if err := _Fiddy.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
