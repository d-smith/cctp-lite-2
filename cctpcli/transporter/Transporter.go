// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transporter

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

// TransporterMetaData contains all meta data concerning the Transporter contract.
var TransporterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_remoteAttestor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestPreviouslyProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnrecognizedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedBodyVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedDestinationDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"}],\"name\":\"DepositForBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"}],\"name\":\"depositForBurn\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBodyVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteAttestor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TransporterABI is the input ABI used to generate the binding from.
// Deprecated: Use TransporterMetaData.ABI instead.
var TransporterABI = TransporterMetaData.ABI

// Transporter is an auto generated Go binding around an Ethereum contract.
type Transporter struct {
	TransporterCaller     // Read-only binding to the contract
	TransporterTransactor // Write-only binding to the contract
	TransporterFilterer   // Log filterer for contract events
}

// TransporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransporterSession struct {
	Contract     *Transporter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransporterCallerSession struct {
	Contract *TransporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TransporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransporterTransactorSession struct {
	Contract     *TransporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TransporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransporterRaw struct {
	Contract *Transporter // Generic contract binding to access the raw methods on
}

// TransporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransporterCallerRaw struct {
	Contract *TransporterCaller // Generic read-only contract binding to access the raw methods on
}

// TransporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransporterTransactorRaw struct {
	Contract *TransporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransporter creates a new instance of Transporter, bound to a specific deployed contract.
func NewTransporter(address common.Address, backend bind.ContractBackend) (*Transporter, error) {
	contract, err := bindTransporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transporter{TransporterCaller: TransporterCaller{contract: contract}, TransporterTransactor: TransporterTransactor{contract: contract}, TransporterFilterer: TransporterFilterer{contract: contract}}, nil
}

// NewTransporterCaller creates a new read-only instance of Transporter, bound to a specific deployed contract.
func NewTransporterCaller(address common.Address, caller bind.ContractCaller) (*TransporterCaller, error) {
	contract, err := bindTransporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransporterCaller{contract: contract}, nil
}

// NewTransporterTransactor creates a new write-only instance of Transporter, bound to a specific deployed contract.
func NewTransporterTransactor(address common.Address, transactor bind.ContractTransactor) (*TransporterTransactor, error) {
	contract, err := bindTransporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransporterTransactor{contract: contract}, nil
}

// NewTransporterFilterer creates a new log filterer instance of Transporter, bound to a specific deployed contract.
func NewTransporterFilterer(address common.Address, filterer bind.ContractFilterer) (*TransporterFilterer, error) {
	contract, err := bindTransporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransporterFilterer{contract: contract}, nil
}

// bindTransporter binds a generic wrapper to an already deployed contract.
func bindTransporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransporterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transporter *TransporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transporter.Contract.TransporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transporter *TransporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transporter.Contract.TransporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transporter *TransporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transporter.Contract.TransporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transporter *TransporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transporter *TransporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transporter *TransporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transporter.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Transporter *TransporterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Transporter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Transporter *TransporterSession) LocalDomain() (uint32, error) {
	return _Transporter.Contract.LocalDomain(&_Transporter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Transporter *TransporterCallerSession) LocalDomain() (uint32, error) {
	return _Transporter.Contract.LocalDomain(&_Transporter.CallOpts)
}

// MessageBodyVersion is a free data retrieval call binding the contract method 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (_Transporter *TransporterCaller) MessageBodyVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Transporter.contract.Call(opts, &out, "messageBodyVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MessageBodyVersion is a free data retrieval call binding the contract method 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (_Transporter *TransporterSession) MessageBodyVersion() (uint32, error) {
	return _Transporter.Contract.MessageBodyVersion(&_Transporter.CallOpts)
}

// MessageBodyVersion is a free data retrieval call binding the contract method 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (_Transporter *TransporterCallerSession) MessageBodyVersion() (uint32, error) {
	return _Transporter.Contract.MessageBodyVersion(&_Transporter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_Transporter *TransporterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Transporter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_Transporter *TransporterSession) NextAvailableNonce() (uint64, error) {
	return _Transporter.Contract.NextAvailableNonce(&_Transporter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_Transporter *TransporterCallerSession) NextAvailableNonce() (uint64, error) {
	return _Transporter.Contract.NextAvailableNonce(&_Transporter.CallOpts)
}

// RemoteAttestor is a free data retrieval call binding the contract method 0x429e8a82.
//
// Solidity: function remoteAttestor() view returns(address)
func (_Transporter *TransporterCaller) RemoteAttestor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transporter.contract.Call(opts, &out, "remoteAttestor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteAttestor is a free data retrieval call binding the contract method 0x429e8a82.
//
// Solidity: function remoteAttestor() view returns(address)
func (_Transporter *TransporterSession) RemoteAttestor() (common.Address, error) {
	return _Transporter.Contract.RemoteAttestor(&_Transporter.CallOpts)
}

// RemoteAttestor is a free data retrieval call binding the contract method 0x429e8a82.
//
// Solidity: function remoteAttestor() view returns(address)
func (_Transporter *TransporterCallerSession) RemoteAttestor() (common.Address, error) {
	return _Transporter.Contract.RemoteAttestor(&_Transporter.CallOpts)
}

// RemoteDomain is a free data retrieval call binding the contract method 0x961681dc.
//
// Solidity: function remoteDomain() view returns(uint32)
func (_Transporter *TransporterCaller) RemoteDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Transporter.contract.Call(opts, &out, "remoteDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RemoteDomain is a free data retrieval call binding the contract method 0x961681dc.
//
// Solidity: function remoteDomain() view returns(uint32)
func (_Transporter *TransporterSession) RemoteDomain() (uint32, error) {
	return _Transporter.Contract.RemoteDomain(&_Transporter.CallOpts)
}

// RemoteDomain is a free data retrieval call binding the contract method 0x961681dc.
//
// Solidity: function remoteDomain() view returns(uint32)
func (_Transporter *TransporterCallerSession) RemoteDomain() (uint32, error) {
	return _Transporter.Contract.RemoteDomain(&_Transporter.CallOpts)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken) returns(uint64)
func (_Transporter *TransporterTransactor) DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address) (*types.Transaction, error) {
	return _Transporter.contract.Transact(opts, "depositForBurn", amount, destinationDomain, mintRecipient, burnToken)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken) returns(uint64)
func (_Transporter *TransporterSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address) (*types.Transaction, error) {
	return _Transporter.Contract.DepositForBurn(&_Transporter.TransactOpts, amount, destinationDomain, mintRecipient, burnToken)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken) returns(uint64)
func (_Transporter *TransporterTransactorSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address) (*types.Transaction, error) {
	return _Transporter.Contract.DepositForBurn(&_Transporter.TransactOpts, amount, destinationDomain, mintRecipient, burnToken)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool)
func (_Transporter *TransporterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _Transporter.contract.Transact(opts, "receiveMessage", message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool)
func (_Transporter *TransporterSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Transporter.Contract.ReceiveMessage(&_Transporter.TransactOpts, message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool)
func (_Transporter *TransporterTransactorSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Transporter.Contract.ReceiveMessage(&_Transporter.TransactOpts, message, attestation)
}

// Recover is a paid mutator transaction binding the contract method 0x1ed13d1b.
//
// Solidity: function recover(bytes message, bytes attestation) returns(address)
func (_Transporter *TransporterTransactor) Recover(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _Transporter.contract.Transact(opts, "recover", message, attestation)
}

// Recover is a paid mutator transaction binding the contract method 0x1ed13d1b.
//
// Solidity: function recover(bytes message, bytes attestation) returns(address)
func (_Transporter *TransporterSession) Recover(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Transporter.Contract.Recover(&_Transporter.TransactOpts, message, attestation)
}

// Recover is a paid mutator transaction binding the contract method 0x1ed13d1b.
//
// Solidity: function recover(bytes message, bytes attestation) returns(address)
func (_Transporter *TransporterTransactorSession) Recover(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Transporter.Contract.Recover(&_Transporter.TransactOpts, message, attestation)
}

// TransporterDepositForBurnIterator is returned from FilterDepositForBurn and is used to iterate over the raw logs and unpacked data for DepositForBurn events raised by the Transporter contract.
type TransporterDepositForBurnIterator struct {
	Event *TransporterDepositForBurn // Event containing the contract specifics and raw log

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
func (it *TransporterDepositForBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransporterDepositForBurn)
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
		it.Event = new(TransporterDepositForBurn)
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
func (it *TransporterDepositForBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransporterDepositForBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransporterDepositForBurn represents a DepositForBurn event raised by the Transporter contract.
type TransporterDepositForBurn struct {
	Nonce             uint64
	BurnToken         common.Address
	Amount            *big.Int
	Depositor         common.Address
	MintRecipient     [32]byte
	DestinationDomain uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDepositForBurn is a free log retrieval operation binding the contract event 0x820119b4d0d835b37528cc59497a58521bbf8a10bde645a60dc16acffd0223e3.
//
// Solidity: event DepositForBurn(uint64 indexed nonce, address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain)
func (_Transporter *TransporterFilterer) FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*TransporterDepositForBurnIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Transporter.contract.FilterLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &TransporterDepositForBurnIterator{contract: _Transporter.contract, event: "DepositForBurn", logs: logs, sub: sub}, nil
}

// WatchDepositForBurn is a free log subscription operation binding the contract event 0x820119b4d0d835b37528cc59497a58521bbf8a10bde645a60dc16acffd0223e3.
//
// Solidity: event DepositForBurn(uint64 indexed nonce, address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain)
func (_Transporter *TransporterFilterer) WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *TransporterDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Transporter.contract.WatchLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransporterDepositForBurn)
				if err := _Transporter.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
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

// ParseDepositForBurn is a log parse operation binding the contract event 0x820119b4d0d835b37528cc59497a58521bbf8a10bde645a60dc16acffd0223e3.
//
// Solidity: event DepositForBurn(uint64 indexed nonce, address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain)
func (_Transporter *TransporterFilterer) ParseDepositForBurn(log types.Log) (*TransporterDepositForBurn, error) {
	event := new(TransporterDepositForBurn)
	if err := _Transporter.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransporterMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the Transporter contract.
type TransporterMessageSentIterator struct {
	Event *TransporterMessageSent // Event containing the contract specifics and raw log

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
func (it *TransporterMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransporterMessageSent)
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
		it.Event = new(TransporterMessageSent)
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
func (it *TransporterMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransporterMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransporterMessageSent represents a MessageSent event raised by the Transporter contract.
type TransporterMessageSent struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_Transporter *TransporterFilterer) FilterMessageSent(opts *bind.FilterOpts) (*TransporterMessageSentIterator, error) {

	logs, sub, err := _Transporter.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &TransporterMessageSentIterator{contract: _Transporter.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_Transporter *TransporterFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *TransporterMessageSent) (event.Subscription, error) {

	logs, sub, err := _Transporter.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransporterMessageSent)
				if err := _Transporter.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_Transporter *TransporterFilterer) ParseMessageSent(log types.Log) (*TransporterMessageSent, error) {
	event := new(TransporterMessageSent)
	if err := _Transporter.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
