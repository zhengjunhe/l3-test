// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControl contract.
type AccessControlRoleAdminChangedIterator struct {
	Event *AccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleAdminChanged)
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
		it.Event = new(AccessControlRoleAdminChanged)
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
func (it *AccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControl contract.
type AccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleAdminChangedIterator{contract: _AccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlRoleAdminChanged, error) {
	event := new(AccessControlRoleAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"LogBridgeBankSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prophecyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBridge.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_btcSender\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogNewProphecyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"LogOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumBridge.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"LogProphecyCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeBank\",\"outputs\":[{\"internalType\":\"contractBridgeBank\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"completeBurnClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"}],\"name\":\"completeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"completeLockClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasBridgeBank\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"}],\"name\":\"isProphecyClaimActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"}],\"name\":\"isProphecyClaimValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"isValidClaimType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prophecyClaimCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prophecyClaims\",\"outputs\":[{\"internalType\":\"enumBridge.ClaimType\",\"name\":\"claimType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"btcSender\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originalValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumBridge.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bridgeBank\",\"type\":\"address\"}],\"name\":\"setBridgeBank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_claimType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_btcSender\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_originalValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setNewProphecyClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"contractValset\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620019dd380380620019dd83398101604081905262000033916200009e565b5f600481905580546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556002805460ff60a01b19908116909155600380549091169055620000d4565b80516001600160a01b038116811462000099575f80fd5b919050565b5f8060408385031215620000b0575f80fd5b620000bb8362000082565b9150620000cb6020840162000082565b90509250929050565b6118fb80620000e25f395ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c80637adbf9731161009e5780638ea5352d1161006e5780638ea5352d1461021d578063aa760bc214610234578063b47010f614610247578063f16e609a1461026e578063fb7831f214610281575f80fd5b80637adbf973146101d15780637dc0d1d0146101e45780637f54af0c146101f7578063814c92c31461020a575f80fd5b8063570ca735116100d9578063570ca735146101755780636536be5a1461018757806369294a4e146101aa57806374c9c3e2146101be575f80fd5b80630e41f3731461010a57806333429e051461013a57806341d197141461014f57806348164a6f14610162575b5f80fd5b60035461011d906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61014d61014836600461127b565b610295565b005b61014d61015d36600461127b565b610315565b61014d6101703660046112c0565b61045d565b5f5461011d906001600160a01b031681565b61019a6101953660046112c0565b610570565b6040519015158152602001610131565b60035461019a90600160a01b900460ff1681565b61014d6101cc3660046113b3565b6105f3565b61014d6101df36600461147c565b610938565b60025461011d906001600160a01b031681565b60015461011d906001600160a01b031681565b61014d61021836600461147c565b610a5e565b61022660045481565b604051908152602001610131565b61019a6102423660046112c0565b610b82565b61025a6102553660046112c0565b610bb1565b60405161013198979695949392919061150c565b61019a61027c36600461158b565b610d19565b60025461019a90600160a01b900460ff1681565b6002546001600160a01b031633146102c85760405162461bcd60e51b81526004016102bf906115a4565b60405180910390fd5b60026102d5848484610d79565b7f7d9a5a85a52a06cdda079e1e12e99f2466f201b2e14ba09ef83ff625949921b985826040516103069291906115eb565b60405180910390a15050505050565b6002546001600160a01b0316331461033f5760405162461bcd60e51b81526004016102bf906115a4565b600161034a85610de7565b600354604051636f0fccab60e01b81526001600160a01b0385811660048301525f921690636f0fccab906024015f60405180830381865afa158015610391573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103b891908101906115ff565b600354604051630be2724760e21b81529192506001600160a01b031690632f89c91c906103ef908890889086908990600401611671565b5f604051808303815f87803b158015610406575f80fd5b505af1158015610418573d5f803e3d5ffd5b505050507f7d9a5a85a52a06cdda079e1e12e99f2466f201b2e14ba09ef83ff625949921b9868360405161044d9291906115eb565b60405180910390a1505050505050565b8061046781610b82565b6104b35760405162461bcd60e51b815260206004820152601c60248201527f50726f706865637920636c61696d206973206e6f74206163746976650000000060448201526064016102bf565b6002546001600160a01b031633146104dd5760405162461bcd60e51b81526004016102bf906115a4565b5f82815260056020526040902060078101805460ff19166002908117909155905460ff1690600190829081111561051657610516611497565b036105295761052483610de7565b610532565b61053283611030565b7f7d9a5a85a52a06cdda079e1e12e99f2466f201b2e14ba09ef83ff625949921b983826040516105639291906115eb565b60405180910390a1505050565b6001545f82815260056020526040808220600301549051631015428760e21b81526001600160a01b039182166004820152919216906340550a1c90602401602060405180830381865afa1580156105c9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105ed91906116ad565b92915050565b600254600160a01b900460ff161515600114801561061f5750600354600160a01b900460ff1615156001145b6106a05760405162461bcd60e51b815260206004820152604660248201527f546865204f70657261746f72206d7573742073657420746865206f7261636c6560448201527f20616e64206272696467652062616e6b20666f722062726964676520616374696064820152653b30ba34b7b760d11b608482015260a4016102bf565b6002546001600160a01b031633146106f05760405162461bcd60e51b815260206004820152601360248201527226bab9ba103132903a34329037b930b1b6329760691b60448201526064016102bf565b6004546106fe9060016116cc565b6004555f60ff8816600281111561071757610717611497565b9050600281600281111561072d5761072d611497565b036107a4576003546040516370e2a8d960e01b81526001600160a01b03909116906370e2a8d9906107629086906004016116eb565b602060405180830381865afa15801561077d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a191906116fd565b93505b5f6040518061010001604052808360028111156107c3576107c3611497565b8152602081018a90526001600160a01b03808a16604083015288811660608301528716608082015260a0810186905260c0810185905260e001600190525f8b8152600560205260409020815181549293508392829060ff1916600183600281111561083057610830611497565b021790555060208201516001820190610849908261179a565b5060408201516002820180546001600160a01b039283166001600160a01b0319918216179091556060840151600384018054918416918316919091179055608084015160048401805491909316911617905560a082015160058201906108af908261179a565b5060c0820151600682015560e082015160078201805460ff191660018360038111156108dd576108dd611497565b02179055509050507f4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8600454838a8a8a8a8a8a604051610924989796959493929190611856565b60405180910390a150505050505050505050565b5f546001600160a01b031633146109895760405162461bcd60e51b815260206004820152601560248201527426bab9ba103132903a34329037b832b930ba37b91760591b60448201526064016102bf565b600254600160a01b900460ff16156109fd5760405162461bcd60e51b815260206004820152603160248201527f546865204f7261636c652063616e6e6f742062652075706461746564206f6e6360448201527019481a5d081a185cc81899595b881cd95d607a1b60648201526084016102bf565b600280546001600160a01b038381166001600160a81b031990921691909117600160a01b1791829055604051911681527f6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0906020015b60405180910390a150565b5f546001600160a01b03163314610aaf5760405162461bcd60e51b815260206004820152601560248201527426bab9ba103132903a34329037b832b930ba37b91760591b60448201526064016102bf565b600354600160a01b900460ff1615610b285760405162461bcd60e51b815260206004820152603660248201527f546865204272696467652042616e6b2063616e6e6f742062652075706461746560448201527519081bdb98d9481a5d081a185cc81899595b881cd95d60521b60648201526084016102bf565b600380546001600160a01b038381166001600160a81b031990921691909117600160a01b1791829055604051911681527fc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa90602001610a53565b5f60015f8381526005602052604090206007015460ff166003811115610baa57610baa611497565b1492915050565b60056020525f90815260409020805460018201805460ff9092169291610bd690611718565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0290611718565b8015610c4d5780601f10610c2457610100808354040283529160200191610c4d565b820191905f5260205f20905b815481529060010190602001808311610c3057829003601f168201915b5050505060028301546003840154600485015460058601805495966001600160a01b03948516969385169550939091169291610c8890611718565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb490611718565b8015610cff5780601f10610cd657610100808354040283529160200191610cff565b820191905f5260205f20905b815481529060010190602001808311610ce257829003601f168201915b50505050600683015460079093015491929160ff16905088565b5f808260ff166002811115610d3057610d30611497565b90506001816002811115610d4657610d46611497565b1480610d6357506002816002811115610d6157610d61611497565b145b15610d715750600192915050565b505f92915050565b60035460405163094b4cd760e11b81526001600160a01b0385811660048301528481166024830152604482018490529091169063129699ae906064015f604051808303815f87803b158015610dcc575f80fd5b505af1158015610dde573d5f803e3d5ffd5b50505050505050565b5f81815260056020526040808220815161010081019092528054829060ff166002811115610e1757610e17611497565b6002811115610e2857610e28611497565b8152602001600182018054610e3c90611718565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6890611718565b8015610eb35780601f10610e8a57610100808354040283529160200191610eb3565b820191905f5260205f20905b815481529060010190602001808311610e9657829003601f168201915b505050918352505060028201546001600160a01b0390811660208301526003830154811660408301526004830154166060820152600582018054608090920191610efc90611718565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2890611718565b8015610f735780601f10610f4a57610100808354040283529160200191610f73565b820191905f5260205f20905b815481529060010190602001808311610f5657829003601f168201915b505050918352505060068201546020820152600782015460409091019060ff166003811115610fa457610fa4611497565b6003811115610fb557610fb5611497565b905250600354604080830151608084015160a085015160c08601519351630be2724760e21b81529596506001600160a01b0390941694632f89c91c94610fff949091600401611671565b5f604051808303815f87803b158015611016575f80fd5b505af1158015611028573d5f803e3d5ffd5b505050505050565b5f81815260056020526040808220815161010081019092528054829060ff16600281111561106057611060611497565b600281111561107157611071611497565b815260200160018201805461108590611718565b80601f01602080910402602001604051908101604052809291908181526020018280546110b190611718565b80156110fc5780601f106110d3576101008083540402835291602001916110fc565b820191905f5260205f20905b8154815290600101906020018083116110df57829003601f168201915b505050918352505060028201546001600160a01b039081166020830152600383015481166040830152600483015416606082015260058201805460809092019161114590611718565b80601f016020809104026020016040519081016040528092919081815260200182805461117190611718565b80156111bc5780601f10611193576101008083540402835291602001916111bc565b820191905f5260205f20905b81548152906001019060200180831161119f57829003601f168201915b505050918352505060068201546020820152600782015460409091019060ff1660038111156111ed576111ed611497565b60038111156111fe576111fe611497565b905250600354604082810151608084015160c0850151925163094b4cd760e11b81526001600160a01b0392831660048201529082166024820152604481019290925292935091169063129699ae90606401610fff565b6001600160a01b0381168114611268575f80fd5b50565b803561127681611254565b919050565b5f805f806080858703121561128e575f80fd5b8435935060208501356112a081611254565b925060408501356112b081611254565b9396929550929360600135925050565b5f602082840312156112d0575f80fd5b5035919050565b803560ff81168114611276575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611324576113246112e7565b604052919050565b5f67ffffffffffffffff821115611345576113456112e7565b50601f01601f191660200190565b5f6113656113608461132c565b6112fb565b9050828152838383011115611378575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261139d575f80fd5b6113ac83833560208501611353565b9392505050565b5f805f805f805f80610100898b0312156113cb575f80fd5b883597506113db60208a016112d7565b9650604089013567ffffffffffffffff808211156113f7575f80fd5b818b0191508b601f83011261140a575f80fd5b6114198c833560208501611353565b975061142760608c0161126b565b965061143560808c0161126b565b955061144360a08c0161126b565b945060c08b0135915080821115611458575f80fd5b506114658b828c0161138e565b92505060e089013590509295985092959890939650565b5f6020828403121561148c575f80fd5b81356113ac81611254565b634e487b7160e01b5f52602160045260245ffd5b600381106114bb576114bb611497565b9052565b5f5b838110156114d95781810151838201526020016114c1565b50505f910152565b5f81518084526114f88160208601602086016114bf565b601f01601f19169290920160200192915050565b5f61010061151a838c6114ab565b80602084015261152c8184018b6114e1565b6001600160a01b038a8116604086015289811660608601528816608085015283810360a0850152905061155f81876114e1565b9150508360c08301526004831061157857611578611497565b8260e08301529998505050505050505050565b5f6020828403121561159b575f80fd5b6113ac826112d7565b60208082526027908201527f4f6e6c7920746865204f7261636c65206d617920636f6d706c6574652070726f6040820152667068656369657360c81b606082015260800190565b828152604081016113ac60208301846114ab565b5f6020828403121561160f575f80fd5b815167ffffffffffffffff811115611625575f80fd5b8201601f81018413611635575f80fd5b80516116436113608261132c565b818152856020838501011115611657575f80fd5b6116688260208301602086016114bf565b95945050505050565b6001600160a01b038581168252841660208201526080604082018190525f9061169c908301856114e1565b905082606083015295945050505050565b5f602082840312156116bd575f80fd5b815180151581146113ac575f80fd5b808201808211156105ed57634e487b7160e01b5f52601160045260245ffd5b602081525f6113ac60208301846114e1565b5f6020828403121561170d575f80fd5b81516113ac81611254565b600181811c9082168061172c57607f821691505b60208210810361174a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115611795575f81815260208120601f850160051c810160208610156117765750805b601f850160051c820191505b8181101561102857828155600101611782565b505050565b815167ffffffffffffffff8111156117b4576117b46112e7565b6117c8816117c28454611718565b84611750565b602080601f8311600181146117fb575f84156117e45750858301515b5f19600386901b1c1916600185901b178555611028565b5f85815260208120601f198616915b828110156118295788860151825594840194600190910190840161180a565b508582101561184657878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f6101008a835261186a602084018b6114ab565b80604084015261187c8184018a6114e1565b6001600160a01b0389811660608601528881166080860152871660a085015283810360c085015290506118af81866114e1565b9150508260e0830152999850505050505050505056fea26469706673582212202204ecf3bb492b82670036e38216a85cb1e106eef86644c87d2b3306483954ca64736f6c63430008140033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _valset common.Address) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, _operator, _valset)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_Bridge *BridgeCaller) BridgeBank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeBank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_Bridge *BridgeSession) BridgeBank() (common.Address, error) {
	return _Bridge.Contract.BridgeBank(&_Bridge.CallOpts)
}

// BridgeBank is a free data retrieval call binding the contract method 0x0e41f373.
//
// Solidity: function bridgeBank() view returns(address)
func (_Bridge *BridgeCallerSession) BridgeBank() (common.Address, error) {
	return _Bridge.Contract.BridgeBank(&_Bridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_Bridge *BridgeCaller) HasBridgeBank(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasBridgeBank")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_Bridge *BridgeSession) HasBridgeBank() (bool, error) {
	return _Bridge.Contract.HasBridgeBank(&_Bridge.CallOpts)
}

// HasBridgeBank is a free data retrieval call binding the contract method 0x69294a4e.
//
// Solidity: function hasBridgeBank() view returns(bool)
func (_Bridge *BridgeCallerSession) HasBridgeBank() (bool, error) {
	return _Bridge.Contract.HasBridgeBank(&_Bridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_Bridge *BridgeCaller) HasOracle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasOracle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_Bridge *BridgeSession) HasOracle() (bool, error) {
	return _Bridge.Contract.HasOracle(&_Bridge.CallOpts)
}

// HasOracle is a free data retrieval call binding the contract method 0xfb7831f2.
//
// Solidity: function hasOracle() view returns(bool)
func (_Bridge *BridgeCallerSession) HasOracle() (bool, error) {
	return _Bridge.Contract.HasOracle(&_Bridge.CallOpts)
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xaa760bc2.
//
// Solidity: function isProphecyClaimActive(bytes32 _claimID) view returns(bool)
func (_Bridge *BridgeCaller) IsProphecyClaimActive(opts *bind.CallOpts, _claimID [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isProphecyClaimActive", _claimID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xaa760bc2.
//
// Solidity: function isProphecyClaimActive(bytes32 _claimID) view returns(bool)
func (_Bridge *BridgeSession) IsProphecyClaimActive(_claimID [32]byte) (bool, error) {
	return _Bridge.Contract.IsProphecyClaimActive(&_Bridge.CallOpts, _claimID)
}

// IsProphecyClaimActive is a free data retrieval call binding the contract method 0xaa760bc2.
//
// Solidity: function isProphecyClaimActive(bytes32 _claimID) view returns(bool)
func (_Bridge *BridgeCallerSession) IsProphecyClaimActive(_claimID [32]byte) (bool, error) {
	return _Bridge.Contract.IsProphecyClaimActive(&_Bridge.CallOpts, _claimID)
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x6536be5a.
//
// Solidity: function isProphecyClaimValidatorActive(bytes32 _claimID) view returns(bool)
func (_Bridge *BridgeCaller) IsProphecyClaimValidatorActive(opts *bind.CallOpts, _claimID [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isProphecyClaimValidatorActive", _claimID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x6536be5a.
//
// Solidity: function isProphecyClaimValidatorActive(bytes32 _claimID) view returns(bool)
func (_Bridge *BridgeSession) IsProphecyClaimValidatorActive(_claimID [32]byte) (bool, error) {
	return _Bridge.Contract.IsProphecyClaimValidatorActive(&_Bridge.CallOpts, _claimID)
}

// IsProphecyClaimValidatorActive is a free data retrieval call binding the contract method 0x6536be5a.
//
// Solidity: function isProphecyClaimValidatorActive(bytes32 _claimID) view returns(bool)
func (_Bridge *BridgeCallerSession) IsProphecyClaimValidatorActive(_claimID [32]byte) (bool, error) {
	return _Bridge.Contract.IsProphecyClaimValidatorActive(&_Bridge.CallOpts, _claimID)
}

// IsValidClaimType is a free data retrieval call binding the contract method 0xf16e609a.
//
// Solidity: function isValidClaimType(uint8 _claimType) pure returns(bool)
func (_Bridge *BridgeCaller) IsValidClaimType(opts *bind.CallOpts, _claimType uint8) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isValidClaimType", _claimType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidClaimType is a free data retrieval call binding the contract method 0xf16e609a.
//
// Solidity: function isValidClaimType(uint8 _claimType) pure returns(bool)
func (_Bridge *BridgeSession) IsValidClaimType(_claimType uint8) (bool, error) {
	return _Bridge.Contract.IsValidClaimType(&_Bridge.CallOpts, _claimType)
}

// IsValidClaimType is a free data retrieval call binding the contract method 0xf16e609a.
//
// Solidity: function isValidClaimType(uint8 _claimType) pure returns(bool)
func (_Bridge *BridgeCallerSession) IsValidClaimType(_claimType uint8) (bool, error) {
	return _Bridge.Contract.IsValidClaimType(&_Bridge.CallOpts, _claimType)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Bridge *BridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Bridge *BridgeSession) Operator() (common.Address, error) {
	return _Bridge.Contract.Operator(&_Bridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Bridge *BridgeCallerSession) Operator() (common.Address, error) {
	return _Bridge.Contract.Operator(&_Bridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Bridge *BridgeCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Bridge *BridgeSession) Oracle() (common.Address, error) {
	return _Bridge.Contract.Oracle(&_Bridge.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Bridge *BridgeCallerSession) Oracle() (common.Address, error) {
	return _Bridge.Contract.Oracle(&_Bridge.CallOpts)
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() view returns(uint256)
func (_Bridge *BridgeCaller) ProphecyClaimCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "prophecyClaimCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() view returns(uint256)
func (_Bridge *BridgeSession) ProphecyClaimCount() (*big.Int, error) {
	return _Bridge.Contract.ProphecyClaimCount(&_Bridge.CallOpts)
}

// ProphecyClaimCount is a free data retrieval call binding the contract method 0x8ea5352d.
//
// Solidity: function prophecyClaimCount() view returns(uint256)
func (_Bridge *BridgeCallerSession) ProphecyClaimCount() (*big.Int, error) {
	return _Bridge.Contract.ProphecyClaimCount(&_Bridge.CallOpts)
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xb47010f6.
//
// Solidity: function prophecyClaims(bytes32 ) view returns(uint8 claimType, bytes btcSender, address ethereumReceiver, address originalValidator, address tokenAddress, string name, uint256 amount, uint8 status)
func (_Bridge *BridgeCaller) ProphecyClaims(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ClaimType         uint8
	BtcSender         []byte
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Name              string
	Amount            *big.Int
	Status            uint8
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "prophecyClaims", arg0)

	outstruct := new(struct {
		ClaimType         uint8
		BtcSender         []byte
		EthereumReceiver  common.Address
		OriginalValidator common.Address
		TokenAddress      common.Address
		Name              string
		Amount            *big.Int
		Status            uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClaimType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.BtcSender = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.EthereumReceiver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OriginalValidator = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TokenAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// ProphecyClaims is a free data retrieval call binding the contract method 0xb47010f6.
//
// Solidity: function prophecyClaims(bytes32 ) view returns(uint8 claimType, bytes btcSender, address ethereumReceiver, address originalValidator, address tokenAddress, string name, uint256 amount, uint8 status)
func (_Bridge *BridgeSession) ProphecyClaims(arg0 [32]byte) (struct {
	ClaimType         uint8
	BtcSender         []byte
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Name              string
	Amount            *big.Int
	Status            uint8
}, error) {
	return _Bridge.Contract.ProphecyClaims(&_Bridge.CallOpts, arg0)
}

// ProphecyClaims is a free data retrieval call binding the contract method 0xb47010f6.
//
// Solidity: function prophecyClaims(bytes32 ) view returns(uint8 claimType, bytes btcSender, address ethereumReceiver, address originalValidator, address tokenAddress, string name, uint256 amount, uint8 status)
func (_Bridge *BridgeCallerSession) ProphecyClaims(arg0 [32]byte) (struct {
	ClaimType         uint8
	BtcSender         []byte
	EthereumReceiver  common.Address
	OriginalValidator common.Address
	TokenAddress      common.Address
	Name              string
	Amount            *big.Int
	Status            uint8
}, error) {
	return _Bridge.Contract.ProphecyClaims(&_Bridge.CallOpts, arg0)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Bridge *BridgeCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "valset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Bridge *BridgeSession) Valset() (common.Address, error) {
	return _Bridge.Contract.Valset(&_Bridge.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Bridge *BridgeCallerSession) Valset() (common.Address, error) {
	return _Bridge.Contract.Valset(&_Bridge.CallOpts)
}

// CompleteBurnClaim is a paid mutator transaction binding the contract method 0x41d19714.
//
// Solidity: function completeBurnClaim(bytes32 _claimID, address _ethereumReceiver, address _tokenAddress, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) CompleteBurnClaim(opts *bind.TransactOpts, _claimID [32]byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "completeBurnClaim", _claimID, _ethereumReceiver, _tokenAddress, _amount)
}

// CompleteBurnClaim is a paid mutator transaction binding the contract method 0x41d19714.
//
// Solidity: function completeBurnClaim(bytes32 _claimID, address _ethereumReceiver, address _tokenAddress, uint256 _amount) returns()
func (_Bridge *BridgeSession) CompleteBurnClaim(_claimID [32]byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteBurnClaim(&_Bridge.TransactOpts, _claimID, _ethereumReceiver, _tokenAddress, _amount)
}

// CompleteBurnClaim is a paid mutator transaction binding the contract method 0x41d19714.
//
// Solidity: function completeBurnClaim(bytes32 _claimID, address _ethereumReceiver, address _tokenAddress, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) CompleteBurnClaim(_claimID [32]byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteBurnClaim(&_Bridge.TransactOpts, _claimID, _ethereumReceiver, _tokenAddress, _amount)
}

// CompleteClaim is a paid mutator transaction binding the contract method 0x48164a6f.
//
// Solidity: function completeClaim(bytes32 _claimID) returns()
func (_Bridge *BridgeTransactor) CompleteClaim(opts *bind.TransactOpts, _claimID [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "completeClaim", _claimID)
}

// CompleteClaim is a paid mutator transaction binding the contract method 0x48164a6f.
//
// Solidity: function completeClaim(bytes32 _claimID) returns()
func (_Bridge *BridgeSession) CompleteClaim(_claimID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteClaim(&_Bridge.TransactOpts, _claimID)
}

// CompleteClaim is a paid mutator transaction binding the contract method 0x48164a6f.
//
// Solidity: function completeClaim(bytes32 _claimID) returns()
func (_Bridge *BridgeTransactorSession) CompleteClaim(_claimID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteClaim(&_Bridge.TransactOpts, _claimID)
}

// CompleteLockClaim is a paid mutator transaction binding the contract method 0x33429e05.
//
// Solidity: function completeLockClaim(bytes32 _claimID, address _ethereumReceiver, address _tokenAddress, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) CompleteLockClaim(opts *bind.TransactOpts, _claimID [32]byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "completeLockClaim", _claimID, _ethereumReceiver, _tokenAddress, _amount)
}

// CompleteLockClaim is a paid mutator transaction binding the contract method 0x33429e05.
//
// Solidity: function completeLockClaim(bytes32 _claimID, address _ethereumReceiver, address _tokenAddress, uint256 _amount) returns()
func (_Bridge *BridgeSession) CompleteLockClaim(_claimID [32]byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteLockClaim(&_Bridge.TransactOpts, _claimID, _ethereumReceiver, _tokenAddress, _amount)
}

// CompleteLockClaim is a paid mutator transaction binding the contract method 0x33429e05.
//
// Solidity: function completeLockClaim(bytes32 _claimID, address _ethereumReceiver, address _tokenAddress, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) CompleteLockClaim(_claimID [32]byte, _ethereumReceiver common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteLockClaim(&_Bridge.TransactOpts, _claimID, _ethereumReceiver, _tokenAddress, _amount)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_Bridge *BridgeTransactor) SetBridgeBank(opts *bind.TransactOpts, _bridgeBank common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setBridgeBank", _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_Bridge *BridgeSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetBridgeBank(&_Bridge.TransactOpts, _bridgeBank)
}

// SetBridgeBank is a paid mutator transaction binding the contract method 0x814c92c3.
//
// Solidity: function setBridgeBank(address _bridgeBank) returns()
func (_Bridge *BridgeTransactorSession) SetBridgeBank(_bridgeBank common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetBridgeBank(&_Bridge.TransactOpts, _bridgeBank)
}

// SetNewProphecyClaim is a paid mutator transaction binding the contract method 0x74c9c3e2.
//
// Solidity: function setNewProphecyClaim(bytes32 _claimID, uint8 _claimType, bytes _btcSender, address _ethereumReceiver, address _originalValidator, address _tokenAddress, string _name, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) SetNewProphecyClaim(opts *bind.TransactOpts, _claimID [32]byte, _claimType uint8, _btcSender []byte, _ethereumReceiver common.Address, _originalValidator common.Address, _tokenAddress common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setNewProphecyClaim", _claimID, _claimType, _btcSender, _ethereumReceiver, _originalValidator, _tokenAddress, _name, _amount)
}

// SetNewProphecyClaim is a paid mutator transaction binding the contract method 0x74c9c3e2.
//
// Solidity: function setNewProphecyClaim(bytes32 _claimID, uint8 _claimType, bytes _btcSender, address _ethereumReceiver, address _originalValidator, address _tokenAddress, string _name, uint256 _amount) returns()
func (_Bridge *BridgeSession) SetNewProphecyClaim(_claimID [32]byte, _claimType uint8, _btcSender []byte, _ethereumReceiver common.Address, _originalValidator common.Address, _tokenAddress common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetNewProphecyClaim(&_Bridge.TransactOpts, _claimID, _claimType, _btcSender, _ethereumReceiver, _originalValidator, _tokenAddress, _name, _amount)
}

// SetNewProphecyClaim is a paid mutator transaction binding the contract method 0x74c9c3e2.
//
// Solidity: function setNewProphecyClaim(bytes32 _claimID, uint8 _claimType, bytes _btcSender, address _ethereumReceiver, address _originalValidator, address _tokenAddress, string _name, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) SetNewProphecyClaim(_claimID [32]byte, _claimType uint8, _btcSender []byte, _ethereumReceiver common.Address, _originalValidator common.Address, _tokenAddress common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetNewProphecyClaim(&_Bridge.TransactOpts, _claimID, _claimType, _btcSender, _ethereumReceiver, _originalValidator, _tokenAddress, _name, _amount)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Bridge *BridgeTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Bridge *BridgeSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetOracle(&_Bridge.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Bridge *BridgeTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetOracle(&_Bridge.TransactOpts, _oracle)
}

// BridgeLogBridgeBankSetIterator is returned from FilterLogBridgeBankSet and is used to iterate over the raw logs and unpacked data for LogBridgeBankSet events raised by the Bridge contract.
type BridgeLogBridgeBankSetIterator struct {
	Event *BridgeLogBridgeBankSet // Event containing the contract specifics and raw log

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
func (it *BridgeLogBridgeBankSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogBridgeBankSet)
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
		it.Event = new(BridgeLogBridgeBankSet)
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
func (it *BridgeLogBridgeBankSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogBridgeBankSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogBridgeBankSet represents a LogBridgeBankSet event raised by the Bridge contract.
type BridgeLogBridgeBankSet struct {
	BridgeBank common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeBankSet is a free log retrieval operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_Bridge *BridgeFilterer) FilterLogBridgeBankSet(opts *bind.FilterOpts) (*BridgeLogBridgeBankSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return &BridgeLogBridgeBankSetIterator{contract: _Bridge.contract, event: "LogBridgeBankSet", logs: logs, sub: sub}, nil
}

// WatchLogBridgeBankSet is a free log subscription operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_Bridge *BridgeFilterer) WatchLogBridgeBankSet(opts *bind.WatchOpts, sink chan<- *BridgeLogBridgeBankSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LogBridgeBankSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogBridgeBankSet)
				if err := _Bridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
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

// ParseLogBridgeBankSet is a log parse operation binding the contract event 0xc8b65043fb196ac032b79a435397d1d14a96b4e9d12e366c3b1f550cb01d2dfa.
//
// Solidity: event LogBridgeBankSet(address _bridgeBank)
func (_Bridge *BridgeFilterer) ParseLogBridgeBankSet(log types.Log) (*BridgeLogBridgeBankSet, error) {
	event := new(BridgeLogBridgeBankSet)
	if err := _Bridge.contract.UnpackLog(event, "LogBridgeBankSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogNewProphecyClaimIterator is returned from FilterLogNewProphecyClaim and is used to iterate over the raw logs and unpacked data for LogNewProphecyClaim events raised by the Bridge contract.
type BridgeLogNewProphecyClaimIterator struct {
	Event *BridgeLogNewProphecyClaim // Event containing the contract specifics and raw log

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
func (it *BridgeLogNewProphecyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogNewProphecyClaim)
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
		it.Event = new(BridgeLogNewProphecyClaim)
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
func (it *BridgeLogNewProphecyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogNewProphecyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogNewProphecyClaim represents a LogNewProphecyClaim event raised by the Bridge contract.
type BridgeLogNewProphecyClaim struct {
	ProphecyID       *big.Int
	ClaimType        uint8
	BtcSender        []byte
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	Name             string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewProphecyClaim is a free log retrieval operation binding the contract event 0x4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, bytes _btcSender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, string _name, uint256 _amount)
func (_Bridge *BridgeFilterer) FilterLogNewProphecyClaim(opts *bind.FilterOpts) (*BridgeLogNewProphecyClaimIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LogNewProphecyClaim")
	if err != nil {
		return nil, err
	}
	return &BridgeLogNewProphecyClaimIterator{contract: _Bridge.contract, event: "LogNewProphecyClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewProphecyClaim is a free log subscription operation binding the contract event 0x4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, bytes _btcSender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, string _name, uint256 _amount)
func (_Bridge *BridgeFilterer) WatchLogNewProphecyClaim(opts *bind.WatchOpts, sink chan<- *BridgeLogNewProphecyClaim) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LogNewProphecyClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogNewProphecyClaim)
				if err := _Bridge.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
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

// ParseLogNewProphecyClaim is a log parse operation binding the contract event 0x4c4b04a2b190e6bb01b6243f150fc76174861acd19cf98841801baaff5262dd8.
//
// Solidity: event LogNewProphecyClaim(uint256 _prophecyID, uint8 _claimType, bytes _btcSender, address _ethereumReceiver, address _validatorAddress, address _tokenAddress, string _name, uint256 _amount)
func (_Bridge *BridgeFilterer) ParseLogNewProphecyClaim(log types.Log) (*BridgeLogNewProphecyClaim, error) {
	event := new(BridgeLogNewProphecyClaim)
	if err := _Bridge.contract.UnpackLog(event, "LogNewProphecyClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogOracleSetIterator is returned from FilterLogOracleSet and is used to iterate over the raw logs and unpacked data for LogOracleSet events raised by the Bridge contract.
type BridgeLogOracleSetIterator struct {
	Event *BridgeLogOracleSet // Event containing the contract specifics and raw log

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
func (it *BridgeLogOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogOracleSet)
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
		it.Event = new(BridgeLogOracleSet)
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
func (it *BridgeLogOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogOracleSet represents a LogOracleSet event raised by the Bridge contract.
type BridgeLogOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogOracleSet is a free log retrieval operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_Bridge *BridgeFilterer) FilterLogOracleSet(opts *bind.FilterOpts) (*BridgeLogOracleSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return &BridgeLogOracleSetIterator{contract: _Bridge.contract, event: "LogOracleSet", logs: logs, sub: sub}, nil
}

// WatchLogOracleSet is a free log subscription operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_Bridge *BridgeFilterer) WatchLogOracleSet(opts *bind.WatchOpts, sink chan<- *BridgeLogOracleSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LogOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogOracleSet)
				if err := _Bridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
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

// ParseLogOracleSet is a log parse operation binding the contract event 0x6efb0434342713e2e9b1501dbebf76b4ed18406ea77ab5d56535cc26dec3adc0.
//
// Solidity: event LogOracleSet(address _oracle)
func (_Bridge *BridgeFilterer) ParseLogOracleSet(log types.Log) (*BridgeLogOracleSet, error) {
	event := new(BridgeLogOracleSet)
	if err := _Bridge.contract.UnpackLog(event, "LogOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogProphecyCompletedIterator is returned from FilterLogProphecyCompleted and is used to iterate over the raw logs and unpacked data for LogProphecyCompleted events raised by the Bridge contract.
type BridgeLogProphecyCompletedIterator struct {
	Event *BridgeLogProphecyCompleted // Event containing the contract specifics and raw log

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
func (it *BridgeLogProphecyCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogProphecyCompleted)
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
		it.Event = new(BridgeLogProphecyCompleted)
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
func (it *BridgeLogProphecyCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogProphecyCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogProphecyCompleted represents a LogProphecyCompleted event raised by the Bridge contract.
type BridgeLogProphecyCompleted struct {
	ClaimID   [32]byte
	ClaimType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyCompleted is a free log retrieval operation binding the contract event 0x7d9a5a85a52a06cdda079e1e12e99f2466f201b2e14ba09ef83ff625949921b9.
//
// Solidity: event LogProphecyCompleted(bytes32 _claimID, uint8 _claimType)
func (_Bridge *BridgeFilterer) FilterLogProphecyCompleted(opts *bind.FilterOpts) (*BridgeLogProphecyCompletedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return &BridgeLogProphecyCompletedIterator{contract: _Bridge.contract, event: "LogProphecyCompleted", logs: logs, sub: sub}, nil
}

// WatchLogProphecyCompleted is a free log subscription operation binding the contract event 0x7d9a5a85a52a06cdda079e1e12e99f2466f201b2e14ba09ef83ff625949921b9.
//
// Solidity: event LogProphecyCompleted(bytes32 _claimID, uint8 _claimType)
func (_Bridge *BridgeFilterer) WatchLogProphecyCompleted(opts *bind.WatchOpts, sink chan<- *BridgeLogProphecyCompleted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LogProphecyCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogProphecyCompleted)
				if err := _Bridge.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
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

// ParseLogProphecyCompleted is a log parse operation binding the contract event 0x7d9a5a85a52a06cdda079e1e12e99f2466f201b2e14ba09ef83ff625949921b9.
//
// Solidity: event LogProphecyCompleted(bytes32 _claimID, uint8 _claimType)
func (_Bridge *BridgeFilterer) ParseLogProphecyCompleted(log types.Log) (*BridgeLogProphecyCompleted, error) {
	event := new(BridgeLogProphecyCompleted)
	if err := _Bridge.contract.UnpackLog(event, "LogProphecyCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankMetaData contains all meta data concerning the BridgeBank contract.
var BridgeBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btcBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenDeployer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bridgeServiceFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"LogBridgeTokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ownerFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"LogBtcTokenBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"LogLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"LogNewBridgeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogUnlock\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addToken2LockList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeServiceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeTokenCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgeTokenWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btcBridge\",\"outputs\":[{\"internalType\":\"contractBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burnBridgeTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_percents\",\"type\":\"uint8\"}],\"name\":\"configLockedTokenOfflineSave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_offlineSave\",\"type\":\"address\"}],\"name\":\"configOfflineSaveAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"configplatformCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createNewBridgeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgetoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toChainID\",\"type\":\"uint256\"}],\"name\":\"enableBridgeToken2Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getLockedTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getToken2address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getToken2addressV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getofflineSaveCfg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"hasBridgeTokenCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasSetPlatformCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_intendedRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintBridgeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offlineSave\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offlineSaveCfgs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_percents\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformCoin\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bridgeServiceFee\",\"type\":\"uint256\"}],\"name\":\"setBridgeServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenDeployer\",\"type\":\"address\"}],\"name\":\"setTokenDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"token2WithdrawCfg\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"token2address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenAddrAllow2name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenAllow2Lock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600b805461ffff19166150051790553480156200001f575f80fd5b506040516200446b3803806200446b83398101604081905262000042916200014e565b5f808055600555600b80546001600160a01b03808916620100000262010000600160b01b031990921691909117909155600d80548783166001600160a01b031991821617909155600e805487841692169190911790556010805491851661010002610100600160a81b031990921691909117905580620001085760405162461bcd60e51b815260206004820152601960248201527f496e76616c6964205f6272696467655365727669636546656500000000000000604482015260640160405180910390fd5b601155600c80546001600160a01b0319166001600160a01b039290921691909117905550620001d292505050565b6001600160a01b03811681146200014b575f80fd5b50565b5f805f805f8060c0878903121562000164575f80fd5b8651620001718162000136565b6020880151909650620001848162000136565b6040880151909550620001978162000136565b6060880151909450620001aa8162000136565b6080880151909350620001bd8162000136565b8092505060a087015190509295509295509295565b61428b80620001e05f395ff3fe60806040526004361062000276575f3560e01c806379380cd2116200014a578063b3f0067411620000ba578063d2168d201162000078578063d2168d201462000899578063dbda621214620008d5578063efdcd97414620008f9578063efea27ab146200091d578063fad60627146200094157620002d5565b8063b3f0067414620007e5578063b5a9096e1462000806578063b86247d7146200081d578063c4c2709e146200084c578063d1d008ae146200087e57620002d5565b806385f0bdc7116200010857806385f0bdc7146200071b578063880183e914620007535780639849085514620007775780639eadc7cc146200078e578063a80de1a414620007c157620002d5565b806379380cd2146200068457806379c50c94146200069b5780637a8d31e514620006b25780637dc0d1d014620006d6578063800823a914620006f757620002d5565b8063328470ab11620001e657806359bf3a3311620001a457806359bf3a3314620005ac5780636f0fccab14620005de5780637056c493146200061157806370e2a8d914620006495780637750c9f0146200066d57620002d5565b8063328470ab14620004d45780633f4d568114620004f957806344aea0de146200052b5780634e25d152146200054f578063570ca735146200058557620002d5565b80631c4336a511620002345780631c4336a514620004015780631ffcdacc146200044557806324ade6d714620004665780632a2dae0a146200048a5780632f89c91c14620004b057620002d5565b80630a1f9b6614620003285780630b5ba4751462000369578063129699ae146200038d57806314f54f9a14620003b15780631ba1c7b914620003dd57620002d5565b36620002d5576006546001600160a01b03163314620002d35760405162461bcd60e51b815260206004820152601460248201527326bab9ba1031329037b7363ca7b3333634b7329760611b60448201526064015b60405180910390fd5b005b6006546001600160a01b03163314620002d35760405162461bcd60e51b815260206004820152601460248201527326bab9ba1031329037b7363ca7b3333634b7329760611b6044820152606401620002ca565b34801562000334575f80fd5b506200034c620003463660046200289b565b62000962565b6040516001600160a01b0390911681526020015b60405180910390f35b34801562000375575f80fd5b50620002d3620003873660046200289b565b620009ab565b34801562000399575f80fd5b50620002d3620003ab366004620028e9565b62000a60565b348015620003bd575f80fd5b50601054620003cc9060ff1681565b604051901515815260200162000360565b348015620003e9575f80fd5b50620002d3620003fb3660046200292c565b62000a9f565b3480156200040d575f80fd5b50620004256200041f3660046200289b565b62000ae2565b604080516001600160a01b03909316835290151560208301520162000360565b34801562000451575f80fd5b50600e546200034c906001600160a01b031681565b34801562000472575f80fd5b50620002d36200048436600462002995565b62000b88565b34801562000496575f80fd5b506010546200034c9061010090046001600160a01b031681565b348015620004bc575f80fd5b50620002d3620004ce36600462002a03565b62000ce3565b348015620004e0575f80fd5b50620004ea5f5481565b60405190815260200162000360565b34801562000505575f80fd5b50620003cc6200051736600462002a71565b60016020525f908152604090205460ff1681565b34801562000537575f80fd5b506200034c6200054936600462002a96565b62000f18565b3480156200055b575f80fd5b50620005736200056d36600462002a71565b62000fc3565b60405162000360949392919062002b5f565b34801562000591575f80fd5b50600b546200034c906201000090046001600160a01b031681565b348015620005b8575f80fd5b50620003cc620005ca36600462002b9e565b60026020525f908152604090205460ff1681565b348015620005ea575f80fd5b5062000602620005fc36600462002a71565b62001087565b60405162000360919062002bb6565b3480156200061d575f80fd5b506200034c6200062f36600462002b9e565b60086020525f90815260409020546001600160a01b031681565b34801562000655575f80fd5b506200034c620006673660046200289b565b6200113a565b620002d36200067e366004620028e9565b620011b1565b34801562000690575f80fd5b506200060262001561565b620002d3620006ac36600462002bca565b620015f5565b348015620006be575f80fd5b50620002d3620006d036600462002b9e565b620017a2565b348015620006e2575f80fd5b50600d546200034c906001600160a01b031681565b34801562000703575f80fd5b50620002d36200071536600462002c14565b6200182b565b34801562000727575f80fd5b506200034c6200073936600462002b9e565b60036020525f90815260409020546001600160a01b031681565b3480156200075f575f80fd5b50620002d36200077136600462002a71565b620018b9565b34801562000783575f80fd5b50620004ea60115481565b3480156200079a575f80fd5b50600b54620007ae90610100900460ff1681565b60405160ff909116815260200162000360565b348015620007cd575f80fd5b5062000602620007df36600462002a71565b62001914565b348015620007f1575f80fd5b50600c546200034c906001600160a01b031681565b34801562000812575f80fd5b50620004ea60055481565b34801562000829575f80fd5b50620004ea6200083b36600462002a71565b60076020525f908152604090205481565b34801562000858575f80fd5b50620003cc6200086a36600462002b9e565b60046020525f908152604090205460ff1681565b3480156200088a575f80fd5b50600b54620007ae9060ff1681565b348015620008a5575f80fd5b50620008bd620008b736600462002a71565b6200192e565b6040805192835260ff90911660208301520162000360565b348015620008e1575f80fd5b50620002d3620008f336600462002a71565b62001a29565b34801562000905575f80fd5b50620002d36200091736600462002a71565b62001a7e565b34801562000929575f80fd5b50620003cc6200093b3660046200289b565b62001b22565b3480156200094d575f80fd5b506006546200034c906001600160a01b031681565b5f808260405160200162000977919062002c41565b60408051601f1981840301815291815281516020928301205f90815260089092529020546001600160a01b03169392505050565b600b546201000090046001600160a01b03163314620009de5760405162461bcd60e51b8152600401620002ca9062002c5e565b60105460ff161562000a415760405162461bcd60e51b815260206004820152602560248201527f54686520706c6174666f726d20436f696e20686173206265656e20636f6e666960448201526419dd5c995960da1b6064820152608401620002ca565b600f62000a4f828262002d12565b50506010805460ff19166001179055565b600e546001600160a01b0316331462000a8d5760405162461bcd60e51b8152600401620002ca9062002ddb565b62000a9a83838362001b65565b505050565b600b546201000090046001600160a01b0316331462000ad25760405162461bcd60e51b8152600401620002ca9062002c5e565b62000ade828262001cbe565b5050565b5f808262000af08162001b22565b62000b0f5760405162461bcd60e51b8152600401620002ca9062002e1e565b5f8460405160200162000b23919062002c41565b60408051601f1981840301815291815281516020928301205f818152600290935291205490915060ff16151560011462000b64575f80935093505062000b82565b5f908152600360205260409020546001600160a01b03169250600191505b50915091565b600b546201000090046001600160a01b0316331462000bbb5760405162461bcd60e51b8152600401620002ca9062002c5e565b6001600160a01b0384161562000c65578280519060200120846001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa15801562000c0f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262000c38919081019062002e5f565b805190602001201462000c5f5760405162461bcd60e51b8152600401620002ca9062002eda565b62000ccf565b60105460ff16151560011462000c8f5760405162461bcd60e51b8152600401620002ca9062002f22565b8251602084012060405162000ca790600f9062002f6b565b60405180910390201462000ccf5760405162461bcd60e51b8152600401620002ca9062002eda565b62000cdd8484848462001dd0565b50505050565b600e546001600160a01b0316331462000d105760405162461bcd60e51b8152600401620002ca9062002ddb565b6001600160a01b0383165f908152600760205260409020548390829081111562000db15760405162461bcd60e51b8152602060048201526044602482018190527f5468652042616e6b20646f6573206e6f7420686f6c6420656e6f756768206c6f908201527f636b656420746f6b656e7320746f2066756c66696c6c2074686973207265717560648201526332b9ba1760e11b608482015260a401620002ca565b84836001600160a01b03821662000e2e578047101562000e285760405162461bcd60e51b815260206004820152602b60248201527f496e73756666696369656e7420657468657265756d2062616c616e636520666f60448201526a39103232b634bb32b93c9760a91b6064820152608401620002ca565b62000f00565b6040516370a0823160e01b815230600482015281906001600160a01b038416906370a0823190602401602060405180830381865afa15801562000e73573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000e99919062002fe5565b101562000f005760405162461bcd60e51b815260206004820152602e60248201527f496e73756666696369656e7420455243323020746f6b656e2062616c616e636560448201526d103337b9103232b634bb32b93c9760911b6064820152608401620002ca565b62000f0e8888888862001ef8565b5050505050505050565b600b545f906201000090046001600160a01b031633148062000f49575060105461010090046001600160a01b031633145b62000fae5760405162461bcd60e51b815260206004820152602e60248201527f4d75737420626520546f6b656e204465706c6f796572206f722042726964676560448201526d2130b7359037b832b930ba37b91760911b6064820152608401620002ca565b62000fbb84848462001fca565b949350505050565b600a6020525f9081526040902080546001820180546001600160a01b03909216929162000ff09062002c95565b80601f01602080910402602001604051908101604052809291908181526020018280546200101e9062002c95565b80156200106d5780601f1062001043576101008083540402835291602001916200106d565b820191905f5260205f20905b8154815290600101906020018083116200104f57829003601f168201915b50505050600283015460039093015491929160ff16905084565b6001600160a01b0381165f90815260096020526040812080546060929190620010b09062002c95565b80601f0160208091040260200160405190810160405280929190818152602001828054620010de9062002c95565b80156200112d5780601f1062001103576101008083540402835291602001916200112d565b820191905f5260205f20905b8154815290600101906020018083116200110f57829003601f168201915b5093979650505050505050565b5f81620011478162001b22565b620011665760405162461bcd60e51b8152600401620002ca9062002e1e565b5f836040516020016200117a919062002c41565b60408051601f1981840301815291815281516020928301205f90815260039092529020546001600160a01b03169250505b50919050565b600554620011c181600162003011565b11620012075760405162461bcd60e51b815260206004820152601460248201527327379030bb30b4b630b13632903737b731b2b99760611b6044820152606401620002ca565b60603415620013e2576001600160a01b038316156200129d5760405162461bcd60e51b8152602060048201526044602482018190527f457468657265756d206465706f73697473207265717569726520746865202774908201527f6f6b656e27206164647265737320746f20626520746865206e756c6c206164646064820152637265737360e01b608482015260a401620002ca565b8134146200131f5760405162461bcd60e51b815260206004820152604260248201527f546865207472616e73616374696f6e732076616c7565206d757374206265206560448201527f7175616c207468652073706563696669656420616d6f756e742028696e207765606482015261692960f01b608482015260a401620002ca565b60105460ff161515600114620013495760405162461bcd60e51b8152600401620002ca9062002f22565b600f8054620013589062002c95565b80601f0160208091040260200160405190810160405280929190818152602001828054620013869062002c95565b8015620013d55780601f10620013ab57610100808354040283529160200191620013d5565b820191905f5260205f20905b815481529060010190602001808311620013b757829003601f168201915b5050505050905062001552565b620013f08333308562002155565b6001600160a01b0383165f9081526009602052604090208054620014149062002c95565b80601f0160208091040260200160405190810160405280929190818152602001828054620014429062002c95565b8015620014915780601f10620014675761010080835404028352916020019162001491565b820191905f5260205f20905b8154815290600101906020018083116200147357829003601f168201915b50505050509050826001600160a01b031660085f83604051602001620014b8919062002c41565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b031614620015525760405162461bcd60e51b815260206004820152603460248201527f54686520746f6b656e206973206e6f7420616c6c6f77656420746f206265206c60448201527337b1b5b2b210333937b69022ba3432b932bab69760611b6064820152608401620002ca565b62000cdd338585848662002282565b600f8054620015709062002c95565b80601f01602080910402602001604051908101604052809291908181526020018280546200159e9062002c95565b8015620015ed5780601f10620015c357610100808354040283529160200191620015ed565b820191905f5260205f20905b815481529060010190602001808311620015cf57829003601f168201915b505050505081565b601154341015620016495760405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420627269646765536572766963654665650000006044820152606401620002ca565b6040516bffffffffffffffffffffffff19606084901b16602082015260348101859052829085906004905f9060540160408051601f198184030181529181528151602092830120835290820192909252015f205460ff161515600114620017025760405162461bcd60e51b815260206004820152602660248201527f436861696e4944206973206e6f7420616c6c6f77656420666f7220627269646760448201526532aa37b5b2b760d11b6064820152608401620002ca565b5f8311620017535760405162461bcd60e51b815260206004820152601a60248201527f616d6f756e74206e6565642067726561746572207468616e20300000000000006044820152606401620002ca565b600c546040516001600160a01b03909116903480156108fc02915f818181858888f193505050501580156200178a573d5f803e3d5ffd5b506200179a86338787876200255f565b505050505050565b600b546201000090046001600160a01b03163314620017d55760405162461bcd60e51b8152600401620002ca9062002c5e565b5f8111620018265760405162461bcd60e51b815260206004820152601960248201527f496e76616c6964205f62726964676553657276696365466565000000000000006044820152606401620002ca565b601155565b600b546201000090046001600160a01b031633146200185e5760405162461bcd60e51b8152600401620002ca9062002c5e565b6040516bffffffffffffffffffffffff19606084901b166020820152603481018290525f9060540160408051601f1981840301815291815281516020928301205f90815260049092529020805460ff19166001179055505050565b600b546201000090046001600160a01b03163314620018ec5760405162461bcd60e51b8152600401620002ca9062002c5e565b601080546001600160a01b0390921661010002610100600160a81b0319909216919091179055565b60096020525f908152604090208054620015709062002c95565b6001600160a01b038082165f908152600a602090815260408083208151608081019092528054909416815260018401805493948594859491929184019190620019779062002c95565b80601f0160208091040260200160405190810160405280929190818152602001828054620019a59062002c95565b8015620019f45780601f10620019ca57610100808354040283529160200191620019f4565b820191905f5260205f20905b815481529060010190602001808311620019d657829003601f168201915b50505091835250506002820154602082015260039091015460ff16604091820152810151606090910151909590945092505050565b600b546201000090046001600160a01b0316331462001a5c5760405162461bcd60e51b8152600401620002ca9062002c5e565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b600b546201000090046001600160a01b0316331462001ab15760405162461bcd60e51b8152600401620002ca9062002c5e565b6001600160a01b03811662001b005760405162461bcd60e51b815260206004820152601460248201527324b73b30b634b2102fb332b2a932b1b2b4bb32b960611b6044820152606401620002ca565b600c80546001600160a01b0319166001600160a01b0392909216919091179055565b5f808260405160200162001b37919062002c41565b60408051601f1981840301815291815281516020928301205f908152600290925290205460ff169392505050565b6001600160a01b0382165f9081526001602052604090205460ff1662001b9f5760405162461bcd60e51b8152600401620002ca906200302d565b6040516340c10f1960e01b81526001600160a01b038481166004830152602482018390528316906340c10f19906044016020604051808303815f875af115801562001bec573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001c12919062003075565b62001c6f5760405162461bcd60e51b815260206004820152602660248201527f417474656d70746564206d696e74206f662062726964676520746f6b656e732060448201526519985a5b195960d21b6064820152608401620002ca565b604080516001600160a01b0384811682526020820184905285168183015290517f4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c9181900360600190a1505050565b5f8160405160200162001cd2919062002c41565b60408051601f1981840301815291815281516020928301205f81815260089093529120549091506001600160a01b0316801562001d885760405162461bcd60e51b815260206004820152604760248201527f54686520746f6b656e2077697468207468652073616d65206e616d652068617360448201527f206265656e20616464656420746f206c6f636b20616c6c6f77206c6973742061606482015266363932b0b23c9760c91b608482015260a401620002ca565b5f82815260086020908152604080832080546001600160a01b0319166001600160a01b03891690811790915583526009909152902062001dc9848262002d12565b5050505050565b600b5460ff9081169082161080159062001df85750600b5460ff610100909104811690821611155b62001e625760405162461bcd60e51b815260206004820152603360248201527f5468652070657263656e747320746f20747269676765722073686f756c642077604482015272697468696e2072616e6765205b352c2038305d60681b6064820152608401620002ca565b604080516080810182526001600160a01b03868116808352602080840188815284860188905260ff871660608601525f928352600a909152939020825181546001600160a01b0319169216919091178155915190918291600182019062001eca908262002d12565b50604082015160028201556060909101516003909101805460ff191660ff9092169190911790555050505050565b6001600160a01b0383165f9081526007602052604090205462001f1d90829062003096565b6001600160a01b0384165f8181526007602052604090209190915562001f78576040516001600160a01b0385169082156108fc029083905f818181858888f1935050505015801562001f71573d5f803e3d5ffd5b5062001f85565b62001f85838583620026aa565b7f802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac8484848460405162001fbc9493929190620030ac565b60405180910390a150505050565b5f8362001fd78162001b22565b15620020305760405162461bcd60e51b815260206004820152602160248201527f546865206e616d6520686173206265656e206372656174656420616c726561646044820152607960f81b6064820152608401620002ca565b5f546200203f90600162003011565b5f819055505f8585856040516200205690620027c3565b6200206493929190620030ea565b604051809103905ff0801580156200207e573d5f803e3d5ffd5b506001600160a01b0381165f908152600160208181526040808420805460ff191690931790925590519293508392620020ba918a910162002c41565b60408051601f1981840301815282825280516020918201205f81815260028352838120805460ff19166001179055600390925291902080546001600160a01b0386166001600160a01b031990911617905591507f438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d773190620021429084908b908b908b9062003126565b60405180910390a1509695505050505050565b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b17905291515f92839290881691620021ba919062002c41565b5f604051808303815f865af19150503d805f8114620021f5576040519150601f19603f3d011682016040523d82523d5f602084013e620021fa565b606091505b5091509150818015620022285750805115806200222857508080602001905181019062002228919062003075565b6200179a5760405162461bcd60e51b8152602060048201526024808201527f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f46416044820152631253115160e21b6064820152608401620002ca565b6005546200229290600162003011565b6005556001600160a01b0383165f90815260076020526040902054620022ba90829062003011565b6001600160a01b0384165f90815260076020526040908190209190915560055490517f4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe916200231491889188918891889188919062003174565b60405180910390a16006546001600160a01b03161562001dc9575f6001600160a01b038416810362002348575047620023b4565b6040516370a0823160e01b81523060048201526001600160a01b038516906370a0823190602401602060405180830381865afa1580156200238b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620023b1919062002fe5565b90505b6001600160a01b038085165f908152600a6020908152604080832081516080810190925280549094168152600184018054939491939192840191620023f99062002c95565b80601f0160208091040260200160405190810160405280929190818152602001828054620024279062002c95565b8015620024765780601f106200244c5761010080835404028352916020019162002476565b820191905f5260205f20905b8154815290600101906020018083116200245857829003601f168201915b50505091835250506002820154602082015260039091015460ff908116604090920191909152600b546060830151929350811691161015620024ba57505062001dc9565b8060400151821015620024cf57505062001dc9565b5f606483836060015160ff16620024e79190620031be565b620024f39190620031d8565b90506001600160a01b0386165f0362002545576006546040516001600160a01b039091169082156108fc029083905f818181858888f193505050501580156200253e573d5f803e3d5ffd5b5062000f0e565b60065462000f0e9087906001600160a01b031683620026aa565b6001600160a01b0382165f9081526001602052604090205460ff16620025995760405162461bcd60e51b8152600401620002ca906200302d565b60405163079cc67960e41b81526001600160a01b038581166004830152602482018390528391908216906379cc6790906044015f604051808303815f87803b158015620025e4575f80fd5b505af1158015620025f7573d5f803e3d5ffd5b505050507f6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e8684836001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200265a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262002683919081019062002e5f565b8589896040516200269a96959493929190620031f8565b60405180910390a1505050505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b17905291515f9283929087169162002707919062002c41565b5f604051808303815f865af19150503d805f811462002742576040519150601f19603f3d011682016040523d82523d5f602084013e62002747565b606091505b5091509150818015620027755750805115806200277557508080602001905181019062002775919062003075565b62001dc95760405162461bcd60e51b815260206004820152601f60248201527f5472616e7366657248656c7065723a205452414e534645525f4641494c4544006044820152606401620002ca565b611014806200324283390190565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715620028115762002811620027d1565b604052919050565b5f67ffffffffffffffff821115620028355762002835620027d1565b50601f01601f191660200190565b5f82601f83011262002853575f80fd5b81356200286a620028648262002819565b620027e5565b8181528460208386010111156200287f575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215620028ac575f80fd5b813567ffffffffffffffff811115620028c3575f80fd5b62000fbb8482850162002843565b6001600160a01b0381168114620028e6575f80fd5b50565b5f805f60608486031215620028fc575f80fd5b83356200290981620028d1565b925060208401356200291b81620028d1565b929592945050506040919091013590565b5f80604083850312156200293e575f80fd5b82356200294b81620028d1565b9150602083013567ffffffffffffffff81111562002967575f80fd5b620029758582860162002843565b9150509250929050565b803560ff8116811462002990575f80fd5b919050565b5f805f8060808587031215620029a9575f80fd5b8435620029b681620028d1565b9350602085013567ffffffffffffffff811115620029d2575f80fd5b620029e08782880162002843565b93505060408501359150620029f8606086016200297f565b905092959194509250565b5f805f806080858703121562002a17575f80fd5b843562002a2481620028d1565b9350602085013562002a3681620028d1565b9250604085013567ffffffffffffffff81111562002a52575f80fd5b62002a608782880162002843565b949793965093946060013593505050565b5f6020828403121562002a82575f80fd5b813562002a8f81620028d1565b9392505050565b5f805f6060848603121562002aa9575f80fd5b833567ffffffffffffffff8082111562002ac1575f80fd5b62002acf8783880162002843565b9450602086013591508082111562002ae5575f80fd5b5062002af48682870162002843565b92505062002b05604085016200297f565b90509250925092565b5f5b8381101562002b2a57818101518382015260200162002b10565b50505f910152565b5f815180845262002b4b81602086016020860162002b0e565b601f01601f19169290920160200192915050565b6001600160a01b03851681526080602082018190525f9062002b849083018662002b32565b905083604083015260ff8316606083015295945050505050565b5f6020828403121562002baf575f80fd5b5035919050565b602081525f62002a8f602083018462002b32565b5f805f806080858703121562002bde575f80fd5b84359350602085013562002bf281620028d1565b9250604085013562002c0481620028d1565b9396929550929360600135925050565b5f806040838503121562002c26575f80fd5b823562002c3381620028d1565b946020939093013593505050565b5f825162002c5481846020870162002b0e565b9190910192915050565b6020808252601c908201527f4d7573742062652042726964676542616e6b206f70657261746f722e00000000604082015260600190565b600181811c9082168062002caa57607f821691505b602082108103620011ab57634e487b7160e01b5f52602260045260245ffd5b601f82111562000a9a575f81815260208120601f850160051c8101602086101562002cf15750805b601f850160051c820191505b818110156200179a5782815560010162002cfd565b815167ffffffffffffffff81111562002d2f5762002d2f620027d1565b62002d478162002d40845462002c95565b8462002cc9565b602080601f83116001811462002d7d575f841562002d655750858301515b5f19600386901b1c1916600185901b1785556200179a565b5f85815260208120601f198616915b8281101562002dad5788860151825594840194600190910190840162002d8c565b508582101562002dcb57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b60208082526023908201527f416363657373207265737472696374656420746f20746865206274632062726960408201526264676560e81b606082015260800190565b60208082526021908201527f546865206e616d6520686173206e6f74206265656e20637265617465642079656040820152601d60fa1b606082015260800190565b5f6020828403121562002e70575f80fd5b815167ffffffffffffffff81111562002e87575f80fd5b8201601f8101841362002e98575f80fd5b805162002ea9620028648262002819565b81815285602083850101111562002ebe575f80fd5b62002ed182602083016020860162002b0e565b95945050505050565b60208082526028908201527f746f6b656e206164647265737320616e64206e616d65206973206e6f7420636f6040820152671b9cda5cdd195b9d60c21b606082015260800190565b60208082526029908201527f54686520706c6174666f726d20436f696e20686173206e6f74206265656e20636040820152681bdb999a59dd5c995960ba1b606082015260800190565b5f80835462002f7a8162002c95565b6001828116801562002f95576001811462002fab5762002fd9565b60ff198416875282151583028701945062002fd9565b875f526020805f205f5b8581101562002fd05781548a82015290840190820162002fb5565b50505082870194505b50929695505050505050565b5f6020828403121562002ff6575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111562003027576200302762002ffd565b92915050565b60208082526028908201527f546f6b656e206d75737420626520612077686974656c6973746564206272696460408201526733b2903a37b5b2b760c11b606082015260800190565b5f6020828403121562003086575f80fd5b8151801515811462002a8f575f80fd5b8181038181111562003027576200302762002ffd565b6001600160a01b038581168252841660208201526080604082018190525f90620030d99083018562002b32565b905082606083015295945050505050565b606081525f620030fe606083018662002b32565b828103602084015262003112818662002b32565b91505060ff83166040830152949350505050565b6001600160a01b03851681526080602082018190525f906200314b9083018662002b32565b82810360408401526200315f818662002b32565b91505060ff8316606083015295945050505050565b6001600160a01b03878116825286811660208301528516604082015260c0606082018190525f90620031a99083018662002b32565b60808301949094525060a00152949350505050565b808202811582820484141762003027576200302762002ffd565b5f82620031f357634e487b7160e01b5f52601260045260245ffd5b500490565b8681525f60018060a01b03808816602084015260c060408401526200322160c084018862002b32565b6060840196909652938416608083015250911660a090910152939250505056fe60a060405234801562000010575f80fd5b50604051620010143803806200101483398101604081905262000033916200020d565b8282600362000043838262000318565b50600462000052828262000318565b50505060ff8116608052620000685f336200009f565b50620000957f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6336200009f565b50505050620003e0565b5f8281526005602090815260408083206001600160a01b038516845290915281205460ff1662000146575f8381526005602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620000fd3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600162000149565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011262000173575f80fd5b81516001600160401b03808211156200019057620001906200014f565b604051601f8301601f19908116603f01168101908282118183101715620001bb57620001bb6200014f565b81604052838152602092508683858801011115620001d7575f80fd5b5f91505b83821015620001fa5785820183015181830184015290820190620001db565b5f93810190920192909252949350505050565b5f805f6060848603121562000220575f80fd5b83516001600160401b038082111562000237575f80fd5b620002458783880162000163565b945060208601519150808211156200025b575f80fd5b506200026a8682870162000163565b925050604084015160ff8116811462000281575f80fd5b809150509250925092565b600181811c90821680620002a157607f821691505b602082108103620002c057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000313575f81815260208120601f850160051c81016020861015620002ee5750805b601f850160051c820191505b818110156200030f57828155600101620002fa565b5050505b505050565b81516001600160401b038111156200033457620003346200014f565b6200034c816200034584546200028c565b84620002c6565b602080601f83116001811462000382575f84156200036a5750858301515b5f19600386901b1c1916600185901b1785556200030f565b5f85815260208120601f198616915b82811015620003b25788860151825594840194600190910190840162000391565b5085821015620003d057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b608051610c1b620003f95f395f6101e90152610c1b5ff3fe608060405234801561000f575f80fd5b5060043610610132575f3560e01c806342966c68116100b4578063a217fddf11610079578063a217fddf146102a2578063a9059cbb146102a9578063aa271e1a146102bc578063d5391393146102cf578063d547741f146102f6578063dd62ed3e14610309575f80fd5b806342966c681461023957806370a082311461024c57806379cc67901461027457806391d148541461028757806395d89b411461029a575f80fd5b8063248a9ca3116100fa578063248a9ca3146101ab5780632f2ff15d146101cd578063313ce567146101e257806336568abe1461021357806340c10f1914610226575f80fd5b806301ffc9a71461013657806306fdde031461015e578063095ea7b31461017357806318160ddd1461018657806323b872dd14610198575b5f80fd5b610149610144366004610a17565b610341565b60405190151581526020015b60405180910390f35b610166610377565b6040516101559190610a45565b610149610181366004610aab565b610407565b6002545b604051908152602001610155565b6101496101a6366004610ad3565b61041e565b61018a6101b9366004610b0c565b5f9081526005602052604090206001015490565b6101e06101db366004610b23565b610441565b005b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610155565b6101e0610221366004610b23565b61046b565b610149610234366004610aab565b6104a3565b6101e0610247366004610b0c565b6104d8565b61018a61025a366004610b4d565b6001600160a01b03165f9081526020819052604090205490565b6101e0610282366004610aab565b6104e5565b610149610295366004610b23565b6104fe565b610166610528565b61018a5f81565b6101496102b7366004610aab565b610537565b6101496102ca366004610b4d565b610544565b61018a7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6101e0610304366004610b23565b61056f565b61018a610317366004610b66565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b5f6001600160e01b03198216637965db0b60e01b148061037157506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606003805461038690610b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546103b290610b8e565b80156103fd5780601f106103d4576101008083540402835291602001916103fd565b820191905f5260205f20905b8154815290600101906020018083116103e057829003601f168201915b5050505050905090565b5f33610414818585610593565b5060019392505050565b5f3361042b8582856105a0565b61043685858561061b565b506001949350505050565b5f8281526005602052604090206001015461045b81610678565b6104658383610682565b50505050565b6001600160a01b03811633146104945760405163334bd91960e11b815260040160405180910390fd5b61049e8282610713565b505050565b5f7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66104ce81610678565b610414848461077e565b6104e233826107b2565b50565b6104f08233836105a0565b6104fa82826107b2565b5050565b5f9182526005602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606004805461038690610b8e565b5f3361041481858561061b565b5f6103717f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6836104fe565b5f8281526005602052604090206001015461058981610678565b6104658383610713565b61049e83838360016107e6565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811015610465578181101561060d57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61046584848484035f6107e6565b6001600160a01b03831661064457604051634b637e8f60e11b81525f6004820152602401610604565b6001600160a01b03821661066d5760405163ec442f0560e01b81525f6004820152602401610604565b61049e8383836108b8565b6104e281336109de565b5f61068d83836104fe565b61070c575f8381526005602090815260408083206001600160a01b03861684529091529020805460ff191660011790556106c43390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610371565b505f610371565b5f61071e83836104fe565b1561070c575f8381526005602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610371565b6001600160a01b0382166107a75760405163ec442f0560e01b81525f6004820152602401610604565b6104fa5f83836108b8565b6001600160a01b0382166107db57604051634b637e8f60e11b81525f6004820152602401610604565b6104fa825f836108b8565b6001600160a01b03841661080f5760405163e602df0560e01b81525f6004820152602401610604565b6001600160a01b03831661083857604051634a1406b160e11b81525f6004820152602401610604565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561046557826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516108aa91815260200190565b60405180910390a350505050565b6001600160a01b0383166108e2578060025f8282546108d79190610bc6565b909155506109529050565b6001600160a01b0383165f90815260208190526040902054818110156109345760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610604565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661096e5760028054829003905561098c565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516109d191815260200190565b60405180910390a3505050565b6109e882826104fe565b6104fa5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610604565b5f60208284031215610a27575f80fd5b81356001600160e01b031981168114610a3e575f80fd5b9392505050565b5f6020808352835180828501525f5b81811015610a7057858101830151858201604001528201610a54565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610aa6575f80fd5b919050565b5f8060408385031215610abc575f80fd5b610ac583610a90565b946020939093013593505050565b5f805f60608486031215610ae5575f80fd5b610aee84610a90565b9250610afc60208501610a90565b9150604084013590509250925092565b5f60208284031215610b1c575f80fd5b5035919050565b5f8060408385031215610b34575f80fd5b82359150610b4460208401610a90565b90509250929050565b5f60208284031215610b5d575f80fd5b610a3e82610a90565b5f8060408385031215610b77575f80fd5b610b8083610a90565b9150610b4460208401610a90565b600181811c90821680610ba257607f821691505b602082108103610bc057634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561037157634e487b7160e01b5f52601160045260245ffdfea26469706673582212207e983d1a5d95ef0a10825dcc962d6f4842fb299e188d00202589abd9c1aaa83464736f6c63430008140033a2646970667358221220bca8a9cbfbe636eea487d3b27f1f421538e3cabd6ce122874a7cde5cbf09526564736f6c63430008140033",
}

// BridgeBankABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBankMetaData.ABI instead.
var BridgeBankABI = BridgeBankMetaData.ABI

// BridgeBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeBankMetaData.Bin instead.
var BridgeBankBin = BridgeBankMetaData.Bin

// DeployBridgeBank deploys a new Ethereum contract, binding an instance of BridgeBank to it.
func DeployBridgeBank(auth *bind.TransactOpts, backend bind.ContractBackend, _operatorAddress common.Address, _oracleAddress common.Address, _btcBridgeAddress common.Address, _tokenDeployer common.Address, _feeReceiver common.Address, _bridgeServiceFee *big.Int) (common.Address, *types.Transaction, *BridgeBank, error) {
	parsed, err := BridgeBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBankBin), backend, _operatorAddress, _oracleAddress, _btcBridgeAddress, _tokenDeployer, _feeReceiver, _bridgeServiceFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// BridgeBank is an auto generated Go binding around an Ethereum contract.
type BridgeBank struct {
	BridgeBankCaller     // Read-only binding to the contract
	BridgeBankTransactor // Write-only binding to the contract
	BridgeBankFilterer   // Log filterer for contract events
}

// BridgeBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBankSession struct {
	Contract     *BridgeBank       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBankCallerSession struct {
	Contract *BridgeBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBankTransactorSession struct {
	Contract     *BridgeBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBankRaw struct {
	Contract *BridgeBank // Generic contract binding to access the raw methods on
}

// BridgeBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBankCallerRaw struct {
	Contract *BridgeBankCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBankTransactorRaw struct {
	Contract *BridgeBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBank creates a new instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBank(address common.Address, backend bind.ContractBackend) (*BridgeBank, error) {
	contract, err := bindBridgeBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBank{BridgeBankCaller: BridgeBankCaller{contract: contract}, BridgeBankTransactor: BridgeBankTransactor{contract: contract}, BridgeBankFilterer: BridgeBankFilterer{contract: contract}}, nil
}

// NewBridgeBankCaller creates a new read-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankCaller(address common.Address, caller bind.ContractCaller) (*BridgeBankCaller, error) {
	contract, err := bindBridgeBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankCaller{contract: contract}, nil
}

// NewBridgeBankTransactor creates a new write-only instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBankTransactor, error) {
	contract, err := bindBridgeBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBankTransactor{contract: contract}, nil
}

// NewBridgeBankFilterer creates a new log filterer instance of BridgeBank, bound to a specific deployed contract.
func NewBridgeBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBankFilterer, error) {
	contract, err := bindBridgeBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBankFilterer{contract: contract}, nil
}

// bindBridgeBank binds a generic wrapper to an already deployed contract.
func bindBridgeBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.BridgeBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.BridgeBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBank *BridgeBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBank *BridgeBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBank *BridgeBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBank.Contract.contract.Transact(opts, method, params...)
}

// BridgeServiceFee is a free data retrieval call binding the contract method 0x98490855.
//
// Solidity: function bridgeServiceFee() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) BridgeServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "bridgeServiceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeServiceFee is a free data retrieval call binding the contract method 0x98490855.
//
// Solidity: function bridgeServiceFee() view returns(uint256)
func (_BridgeBank *BridgeBankSession) BridgeServiceFee() (*big.Int, error) {
	return _BridgeBank.Contract.BridgeServiceFee(&_BridgeBank.CallOpts)
}

// BridgeServiceFee is a free data retrieval call binding the contract method 0x98490855.
//
// Solidity: function bridgeServiceFee() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) BridgeServiceFee() (*big.Int, error) {
	return _BridgeBank.Contract.BridgeServiceFee(&_BridgeBank.CallOpts)
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) BridgeTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "bridgeTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() view returns(uint256)
func (_BridgeBank *BridgeBankSession) BridgeTokenCount() (*big.Int, error) {
	return _BridgeBank.Contract.BridgeTokenCount(&_BridgeBank.CallOpts)
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) BridgeTokenCount() (*big.Int, error) {
	return _BridgeBank.Contract.BridgeTokenCount(&_BridgeBank.CallOpts)
}

// BridgeTokenCreated is a free data retrieval call binding the contract method 0x59bf3a33.
//
// Solidity: function bridgeTokenCreated(bytes32 ) view returns(bool)
func (_BridgeBank *BridgeBankCaller) BridgeTokenCreated(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "bridgeTokenCreated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeTokenCreated is a free data retrieval call binding the contract method 0x59bf3a33.
//
// Solidity: function bridgeTokenCreated(bytes32 ) view returns(bool)
func (_BridgeBank *BridgeBankSession) BridgeTokenCreated(arg0 [32]byte) (bool, error) {
	return _BridgeBank.Contract.BridgeTokenCreated(&_BridgeBank.CallOpts, arg0)
}

// BridgeTokenCreated is a free data retrieval call binding the contract method 0x59bf3a33.
//
// Solidity: function bridgeTokenCreated(bytes32 ) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) BridgeTokenCreated(arg0 [32]byte) (bool, error) {
	return _BridgeBank.Contract.BridgeTokenCreated(&_BridgeBank.CallOpts, arg0)
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) view returns(bool)
func (_BridgeBank *BridgeBankCaller) BridgeTokenWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "bridgeTokenWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) view returns(bool)
func (_BridgeBank *BridgeBankSession) BridgeTokenWhitelist(arg0 common.Address) (bool, error) {
	return _BridgeBank.Contract.BridgeTokenWhitelist(&_BridgeBank.CallOpts, arg0)
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) BridgeTokenWhitelist(arg0 common.Address) (bool, error) {
	return _BridgeBank.Contract.BridgeTokenWhitelist(&_BridgeBank.CallOpts, arg0)
}

// BtcBridge is a free data retrieval call binding the contract method 0x1ffcdacc.
//
// Solidity: function btcBridge() view returns(address)
func (_BridgeBank *BridgeBankCaller) BtcBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "btcBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcBridge is a free data retrieval call binding the contract method 0x1ffcdacc.
//
// Solidity: function btcBridge() view returns(address)
func (_BridgeBank *BridgeBankSession) BtcBridge() (common.Address, error) {
	return _BridgeBank.Contract.BtcBridge(&_BridgeBank.CallOpts)
}

// BtcBridge is a free data retrieval call binding the contract method 0x1ffcdacc.
//
// Solidity: function btcBridge() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) BtcBridge() (common.Address, error) {
	return _BridgeBank.Contract.BtcBridge(&_BridgeBank.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_BridgeBank *BridgeBankCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_BridgeBank *BridgeBankSession) FeeReceiver() (common.Address, error) {
	return _BridgeBank.Contract.FeeReceiver(&_BridgeBank.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) FeeReceiver() (common.Address, error) {
	return _BridgeBank.Contract.FeeReceiver(&_BridgeBank.CallOpts)
}

// GetLockedTokenAddress is a free data retrieval call binding the contract method 0x0a1f9b66.
//
// Solidity: function getLockedTokenAddress(string _name) view returns(address)
func (_BridgeBank *BridgeBankCaller) GetLockedTokenAddress(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "getLockedTokenAddress", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLockedTokenAddress is a free data retrieval call binding the contract method 0x0a1f9b66.
//
// Solidity: function getLockedTokenAddress(string _name) view returns(address)
func (_BridgeBank *BridgeBankSession) GetLockedTokenAddress(_name string) (common.Address, error) {
	return _BridgeBank.Contract.GetLockedTokenAddress(&_BridgeBank.CallOpts, _name)
}

// GetLockedTokenAddress is a free data retrieval call binding the contract method 0x0a1f9b66.
//
// Solidity: function getLockedTokenAddress(string _name) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) GetLockedTokenAddress(_name string) (common.Address, error) {
	return _BridgeBank.Contract.GetLockedTokenAddress(&_BridgeBank.CallOpts, _name)
}

// GetToken2address is a free data retrieval call binding the contract method 0x70e2a8d9.
//
// Solidity: function getToken2address(string _name) view returns(address)
func (_BridgeBank *BridgeBankCaller) GetToken2address(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "getToken2address", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken2address is a free data retrieval call binding the contract method 0x70e2a8d9.
//
// Solidity: function getToken2address(string _name) view returns(address)
func (_BridgeBank *BridgeBankSession) GetToken2address(_name string) (common.Address, error) {
	return _BridgeBank.Contract.GetToken2address(&_BridgeBank.CallOpts, _name)
}

// GetToken2address is a free data retrieval call binding the contract method 0x70e2a8d9.
//
// Solidity: function getToken2address(string _name) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) GetToken2address(_name string) (common.Address, error) {
	return _BridgeBank.Contract.GetToken2address(&_BridgeBank.CallOpts, _name)
}

// GetToken2addressV2 is a free data retrieval call binding the contract method 0x1c4336a5.
//
// Solidity: function getToken2addressV2(string _name) view returns(address, bool)
func (_BridgeBank *BridgeBankCaller) GetToken2addressV2(opts *bind.CallOpts, _name string) (common.Address, bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "getToken2addressV2", _name)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetToken2addressV2 is a free data retrieval call binding the contract method 0x1c4336a5.
//
// Solidity: function getToken2addressV2(string _name) view returns(address, bool)
func (_BridgeBank *BridgeBankSession) GetToken2addressV2(_name string) (common.Address, bool, error) {
	return _BridgeBank.Contract.GetToken2addressV2(&_BridgeBank.CallOpts, _name)
}

// GetToken2addressV2 is a free data retrieval call binding the contract method 0x1c4336a5.
//
// Solidity: function getToken2addressV2(string _name) view returns(address, bool)
func (_BridgeBank *BridgeBankCallerSession) GetToken2addressV2(_name string) (common.Address, bool, error) {
	return _BridgeBank.Contract.GetToken2addressV2(&_BridgeBank.CallOpts, _name)
}

// GetTokenName is a free data retrieval call binding the contract method 0x6f0fccab.
//
// Solidity: function getTokenName(address _tokenAddress) view returns(string)
func (_BridgeBank *BridgeBankCaller) GetTokenName(opts *bind.CallOpts, _tokenAddress common.Address) (string, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "getTokenName", _tokenAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenName is a free data retrieval call binding the contract method 0x6f0fccab.
//
// Solidity: function getTokenName(address _tokenAddress) view returns(string)
func (_BridgeBank *BridgeBankSession) GetTokenName(_tokenAddress common.Address) (string, error) {
	return _BridgeBank.Contract.GetTokenName(&_BridgeBank.CallOpts, _tokenAddress)
}

// GetTokenName is a free data retrieval call binding the contract method 0x6f0fccab.
//
// Solidity: function getTokenName(address _tokenAddress) view returns(string)
func (_BridgeBank *BridgeBankCallerSession) GetTokenName(_tokenAddress common.Address) (string, error) {
	return _BridgeBank.Contract.GetTokenName(&_BridgeBank.CallOpts, _tokenAddress)
}

// GetofflineSaveCfg is a free data retrieval call binding the contract method 0xd2168d20.
//
// Solidity: function getofflineSaveCfg(address _token) view returns(uint256, uint8)
func (_BridgeBank *BridgeBankCaller) GetofflineSaveCfg(opts *bind.CallOpts, _token common.Address) (*big.Int, uint8, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "getofflineSaveCfg", _token)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetofflineSaveCfg is a free data retrieval call binding the contract method 0xd2168d20.
//
// Solidity: function getofflineSaveCfg(address _token) view returns(uint256, uint8)
func (_BridgeBank *BridgeBankSession) GetofflineSaveCfg(_token common.Address) (*big.Int, uint8, error) {
	return _BridgeBank.Contract.GetofflineSaveCfg(&_BridgeBank.CallOpts, _token)
}

// GetofflineSaveCfg is a free data retrieval call binding the contract method 0xd2168d20.
//
// Solidity: function getofflineSaveCfg(address _token) view returns(uint256, uint8)
func (_BridgeBank *BridgeBankCallerSession) GetofflineSaveCfg(_token common.Address) (*big.Int, uint8, error) {
	return _BridgeBank.Contract.GetofflineSaveCfg(&_BridgeBank.CallOpts, _token)
}

// HasBridgeTokenCreated is a free data retrieval call binding the contract method 0xefea27ab.
//
// Solidity: function hasBridgeTokenCreated(string _name) view returns(bool)
func (_BridgeBank *BridgeBankCaller) HasBridgeTokenCreated(opts *bind.CallOpts, _name string) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "hasBridgeTokenCreated", _name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBridgeTokenCreated is a free data retrieval call binding the contract method 0xefea27ab.
//
// Solidity: function hasBridgeTokenCreated(string _name) view returns(bool)
func (_BridgeBank *BridgeBankSession) HasBridgeTokenCreated(_name string) (bool, error) {
	return _BridgeBank.Contract.HasBridgeTokenCreated(&_BridgeBank.CallOpts, _name)
}

// HasBridgeTokenCreated is a free data retrieval call binding the contract method 0xefea27ab.
//
// Solidity: function hasBridgeTokenCreated(string _name) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) HasBridgeTokenCreated(_name string) (bool, error) {
	return _BridgeBank.Contract.HasBridgeTokenCreated(&_BridgeBank.CallOpts, _name)
}

// HasSetPlatformCoin is a free data retrieval call binding the contract method 0x14f54f9a.
//
// Solidity: function hasSetPlatformCoin() view returns(bool)
func (_BridgeBank *BridgeBankCaller) HasSetPlatformCoin(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "hasSetPlatformCoin")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSetPlatformCoin is a free data retrieval call binding the contract method 0x14f54f9a.
//
// Solidity: function hasSetPlatformCoin() view returns(bool)
func (_BridgeBank *BridgeBankSession) HasSetPlatformCoin() (bool, error) {
	return _BridgeBank.Contract.HasSetPlatformCoin(&_BridgeBank.CallOpts)
}

// HasSetPlatformCoin is a free data retrieval call binding the contract method 0x14f54f9a.
//
// Solidity: function hasSetPlatformCoin() view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) HasSetPlatformCoin() (bool, error) {
	return _BridgeBank.Contract.HasSetPlatformCoin(&_BridgeBank.CallOpts)
}

// HighThreshold is a free data retrieval call binding the contract method 0x9eadc7cc.
//
// Solidity: function highThreshold() view returns(uint8)
func (_BridgeBank *BridgeBankCaller) HighThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "highThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HighThreshold is a free data retrieval call binding the contract method 0x9eadc7cc.
//
// Solidity: function highThreshold() view returns(uint8)
func (_BridgeBank *BridgeBankSession) HighThreshold() (uint8, error) {
	return _BridgeBank.Contract.HighThreshold(&_BridgeBank.CallOpts)
}

// HighThreshold is a free data retrieval call binding the contract method 0x9eadc7cc.
//
// Solidity: function highThreshold() view returns(uint8)
func (_BridgeBank *BridgeBankCallerSession) HighThreshold() (uint8, error) {
	return _BridgeBank.Contract.HighThreshold(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "lockNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockNonce() (*big.Int, error) {
	return _BridgeBank.Contract.LockNonce(&_BridgeBank.CallOpts)
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) view returns(uint256)
func (_BridgeBank *BridgeBankCaller) LockedFunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "lockedFunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) view returns(uint256)
func (_BridgeBank *BridgeBankSession) LockedFunds(arg0 common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.LockedFunds(&_BridgeBank.CallOpts, arg0)
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) view returns(uint256)
func (_BridgeBank *BridgeBankCallerSession) LockedFunds(arg0 common.Address) (*big.Int, error) {
	return _BridgeBank.Contract.LockedFunds(&_BridgeBank.CallOpts, arg0)
}

// LowThreshold is a free data retrieval call binding the contract method 0xd1d008ae.
//
// Solidity: function lowThreshold() view returns(uint8)
func (_BridgeBank *BridgeBankCaller) LowThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "lowThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LowThreshold is a free data retrieval call binding the contract method 0xd1d008ae.
//
// Solidity: function lowThreshold() view returns(uint8)
func (_BridgeBank *BridgeBankSession) LowThreshold() (uint8, error) {
	return _BridgeBank.Contract.LowThreshold(&_BridgeBank.CallOpts)
}

// LowThreshold is a free data retrieval call binding the contract method 0xd1d008ae.
//
// Solidity: function lowThreshold() view returns(uint8)
func (_BridgeBank *BridgeBankCallerSession) LowThreshold() (uint8, error) {
	return _BridgeBank.Contract.LowThreshold(&_BridgeBank.CallOpts)
}

// OfflineSave is a free data retrieval call binding the contract method 0xfad60627.
//
// Solidity: function offlineSave() view returns(address)
func (_BridgeBank *BridgeBankCaller) OfflineSave(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "offlineSave")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OfflineSave is a free data retrieval call binding the contract method 0xfad60627.
//
// Solidity: function offlineSave() view returns(address)
func (_BridgeBank *BridgeBankSession) OfflineSave() (common.Address, error) {
	return _BridgeBank.Contract.OfflineSave(&_BridgeBank.CallOpts)
}

// OfflineSave is a free data retrieval call binding the contract method 0xfad60627.
//
// Solidity: function offlineSave() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) OfflineSave() (common.Address, error) {
	return _BridgeBank.Contract.OfflineSave(&_BridgeBank.CallOpts)
}

// OfflineSaveCfgs is a free data retrieval call binding the contract method 0x4e25d152.
//
// Solidity: function offlineSaveCfgs(address ) view returns(address token, string name, uint256 _threshold, uint8 _percents)
func (_BridgeBank *BridgeBankCaller) OfflineSaveCfgs(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Threshold *big.Int
	Percents  uint8
}, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "offlineSaveCfgs", arg0)

	outstruct := new(struct {
		Token     common.Address
		Name      string
		Threshold *big.Int
		Percents  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Threshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Percents = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// OfflineSaveCfgs is a free data retrieval call binding the contract method 0x4e25d152.
//
// Solidity: function offlineSaveCfgs(address ) view returns(address token, string name, uint256 _threshold, uint8 _percents)
func (_BridgeBank *BridgeBankSession) OfflineSaveCfgs(arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Threshold *big.Int
	Percents  uint8
}, error) {
	return _BridgeBank.Contract.OfflineSaveCfgs(&_BridgeBank.CallOpts, arg0)
}

// OfflineSaveCfgs is a free data retrieval call binding the contract method 0x4e25d152.
//
// Solidity: function offlineSaveCfgs(address ) view returns(address token, string name, uint256 _threshold, uint8 _percents)
func (_BridgeBank *BridgeBankCallerSession) OfflineSaveCfgs(arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Threshold *big.Int
	Percents  uint8
}, error) {
	return _BridgeBank.Contract.OfflineSaveCfgs(&_BridgeBank.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Operator() (common.Address, error) {
	return _BridgeBank.Contract.Operator(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Oracle() (common.Address, error) {
	return _BridgeBank.Contract.Oracle(&_BridgeBank.CallOpts)
}

// PlatformCoin is a free data retrieval call binding the contract method 0x79380cd2.
//
// Solidity: function platformCoin() view returns(string)
func (_BridgeBank *BridgeBankCaller) PlatformCoin(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "platformCoin")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PlatformCoin is a free data retrieval call binding the contract method 0x79380cd2.
//
// Solidity: function platformCoin() view returns(string)
func (_BridgeBank *BridgeBankSession) PlatformCoin() (string, error) {
	return _BridgeBank.Contract.PlatformCoin(&_BridgeBank.CallOpts)
}

// PlatformCoin is a free data retrieval call binding the contract method 0x79380cd2.
//
// Solidity: function platformCoin() view returns(string)
func (_BridgeBank *BridgeBankCallerSession) PlatformCoin() (string, error) {
	return _BridgeBank.Contract.PlatformCoin(&_BridgeBank.CallOpts)
}

// Token2WithdrawCfg is a free data retrieval call binding the contract method 0xc4c2709e.
//
// Solidity: function token2WithdrawCfg(bytes32 ) view returns(bool)
func (_BridgeBank *BridgeBankCaller) Token2WithdrawCfg(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "token2WithdrawCfg", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Token2WithdrawCfg is a free data retrieval call binding the contract method 0xc4c2709e.
//
// Solidity: function token2WithdrawCfg(bytes32 ) view returns(bool)
func (_BridgeBank *BridgeBankSession) Token2WithdrawCfg(arg0 [32]byte) (bool, error) {
	return _BridgeBank.Contract.Token2WithdrawCfg(&_BridgeBank.CallOpts, arg0)
}

// Token2WithdrawCfg is a free data retrieval call binding the contract method 0xc4c2709e.
//
// Solidity: function token2WithdrawCfg(bytes32 ) view returns(bool)
func (_BridgeBank *BridgeBankCallerSession) Token2WithdrawCfg(arg0 [32]byte) (bool, error) {
	return _BridgeBank.Contract.Token2WithdrawCfg(&_BridgeBank.CallOpts, arg0)
}

// Token2address is a free data retrieval call binding the contract method 0x85f0bdc7.
//
// Solidity: function token2address(bytes32 ) view returns(address)
func (_BridgeBank *BridgeBankCaller) Token2address(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "token2address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token2address is a free data retrieval call binding the contract method 0x85f0bdc7.
//
// Solidity: function token2address(bytes32 ) view returns(address)
func (_BridgeBank *BridgeBankSession) Token2address(arg0 [32]byte) (common.Address, error) {
	return _BridgeBank.Contract.Token2address(&_BridgeBank.CallOpts, arg0)
}

// Token2address is a free data retrieval call binding the contract method 0x85f0bdc7.
//
// Solidity: function token2address(bytes32 ) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) Token2address(arg0 [32]byte) (common.Address, error) {
	return _BridgeBank.Contract.Token2address(&_BridgeBank.CallOpts, arg0)
}

// TokenAddrAllow2name is a free data retrieval call binding the contract method 0xa80de1a4.
//
// Solidity: function tokenAddrAllow2name(address ) view returns(string)
func (_BridgeBank *BridgeBankCaller) TokenAddrAllow2name(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "tokenAddrAllow2name", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenAddrAllow2name is a free data retrieval call binding the contract method 0xa80de1a4.
//
// Solidity: function tokenAddrAllow2name(address ) view returns(string)
func (_BridgeBank *BridgeBankSession) TokenAddrAllow2name(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.TokenAddrAllow2name(&_BridgeBank.CallOpts, arg0)
}

// TokenAddrAllow2name is a free data retrieval call binding the contract method 0xa80de1a4.
//
// Solidity: function tokenAddrAllow2name(address ) view returns(string)
func (_BridgeBank *BridgeBankCallerSession) TokenAddrAllow2name(arg0 common.Address) (string, error) {
	return _BridgeBank.Contract.TokenAddrAllow2name(&_BridgeBank.CallOpts, arg0)
}

// TokenAllow2Lock is a free data retrieval call binding the contract method 0x7056c493.
//
// Solidity: function tokenAllow2Lock(bytes32 ) view returns(address)
func (_BridgeBank *BridgeBankCaller) TokenAllow2Lock(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "tokenAllow2Lock", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAllow2Lock is a free data retrieval call binding the contract method 0x7056c493.
//
// Solidity: function tokenAllow2Lock(bytes32 ) view returns(address)
func (_BridgeBank *BridgeBankSession) TokenAllow2Lock(arg0 [32]byte) (common.Address, error) {
	return _BridgeBank.Contract.TokenAllow2Lock(&_BridgeBank.CallOpts, arg0)
}

// TokenAllow2Lock is a free data retrieval call binding the contract method 0x7056c493.
//
// Solidity: function tokenAllow2Lock(bytes32 ) view returns(address)
func (_BridgeBank *BridgeBankCallerSession) TokenAllow2Lock(arg0 [32]byte) (common.Address, error) {
	return _BridgeBank.Contract.TokenAllow2Lock(&_BridgeBank.CallOpts, arg0)
}

// TokenDeployer is a free data retrieval call binding the contract method 0x2a2dae0a.
//
// Solidity: function tokenDeployer() view returns(address)
func (_BridgeBank *BridgeBankCaller) TokenDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBank.contract.Call(opts, &out, "tokenDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDeployer is a free data retrieval call binding the contract method 0x2a2dae0a.
//
// Solidity: function tokenDeployer() view returns(address)
func (_BridgeBank *BridgeBankSession) TokenDeployer() (common.Address, error) {
	return _BridgeBank.Contract.TokenDeployer(&_BridgeBank.CallOpts)
}

// TokenDeployer is a free data retrieval call binding the contract method 0x2a2dae0a.
//
// Solidity: function tokenDeployer() view returns(address)
func (_BridgeBank *BridgeBankCallerSession) TokenDeployer() (common.Address, error) {
	return _BridgeBank.Contract.TokenDeployer(&_BridgeBank.CallOpts)
}

// AddToken2LockList is a paid mutator transaction binding the contract method 0x1ba1c7b9.
//
// Solidity: function addToken2LockList(address _token, string _name) returns()
func (_BridgeBank *BridgeBankTransactor) AddToken2LockList(opts *bind.TransactOpts, _token common.Address, _name string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "addToken2LockList", _token, _name)
}

// AddToken2LockList is a paid mutator transaction binding the contract method 0x1ba1c7b9.
//
// Solidity: function addToken2LockList(address _token, string _name) returns()
func (_BridgeBank *BridgeBankSession) AddToken2LockList(_token common.Address, _name string) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddToken2LockList(&_BridgeBank.TransactOpts, _token, _name)
}

// AddToken2LockList is a paid mutator transaction binding the contract method 0x1ba1c7b9.
//
// Solidity: function addToken2LockList(address _token, string _name) returns()
func (_BridgeBank *BridgeBankTransactorSession) AddToken2LockList(_token common.Address, _name string) (*types.Transaction, error) {
	return _BridgeBank.Contract.AddToken2LockList(&_BridgeBank.TransactOpts, _token, _name)
}

// BurnBridgeTokens is a paid mutator transaction binding the contract method 0x79c50c94.
//
// Solidity: function burnBridgeTokens(uint256 _chainID, address _receiver, address _bridgeTokenAddress, uint256 _amount) payable returns()
func (_BridgeBank *BridgeBankTransactor) BurnBridgeTokens(opts *bind.TransactOpts, _chainID *big.Int, _receiver common.Address, _bridgeTokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "burnBridgeTokens", _chainID, _receiver, _bridgeTokenAddress, _amount)
}

// BurnBridgeTokens is a paid mutator transaction binding the contract method 0x79c50c94.
//
// Solidity: function burnBridgeTokens(uint256 _chainID, address _receiver, address _bridgeTokenAddress, uint256 _amount) payable returns()
func (_BridgeBank *BridgeBankSession) BurnBridgeTokens(_chainID *big.Int, _receiver common.Address, _bridgeTokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.BurnBridgeTokens(&_BridgeBank.TransactOpts, _chainID, _receiver, _bridgeTokenAddress, _amount)
}

// BurnBridgeTokens is a paid mutator transaction binding the contract method 0x79c50c94.
//
// Solidity: function burnBridgeTokens(uint256 _chainID, address _receiver, address _bridgeTokenAddress, uint256 _amount) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) BurnBridgeTokens(_chainID *big.Int, _receiver common.Address, _bridgeTokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.BurnBridgeTokens(&_BridgeBank.TransactOpts, _chainID, _receiver, _bridgeTokenAddress, _amount)
}

// ConfigLockedTokenOfflineSave is a paid mutator transaction binding the contract method 0x24ade6d7.
//
// Solidity: function configLockedTokenOfflineSave(address _token, string _name, uint256 _threshold, uint8 _percents) returns()
func (_BridgeBank *BridgeBankTransactor) ConfigLockedTokenOfflineSave(opts *bind.TransactOpts, _token common.Address, _name string, _threshold *big.Int, _percents uint8) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "configLockedTokenOfflineSave", _token, _name, _threshold, _percents)
}

// ConfigLockedTokenOfflineSave is a paid mutator transaction binding the contract method 0x24ade6d7.
//
// Solidity: function configLockedTokenOfflineSave(address _token, string _name, uint256 _threshold, uint8 _percents) returns()
func (_BridgeBank *BridgeBankSession) ConfigLockedTokenOfflineSave(_token common.Address, _name string, _threshold *big.Int, _percents uint8) (*types.Transaction, error) {
	return _BridgeBank.Contract.ConfigLockedTokenOfflineSave(&_BridgeBank.TransactOpts, _token, _name, _threshold, _percents)
}

// ConfigLockedTokenOfflineSave is a paid mutator transaction binding the contract method 0x24ade6d7.
//
// Solidity: function configLockedTokenOfflineSave(address _token, string _name, uint256 _threshold, uint8 _percents) returns()
func (_BridgeBank *BridgeBankTransactorSession) ConfigLockedTokenOfflineSave(_token common.Address, _name string, _threshold *big.Int, _percents uint8) (*types.Transaction, error) {
	return _BridgeBank.Contract.ConfigLockedTokenOfflineSave(&_BridgeBank.TransactOpts, _token, _name, _threshold, _percents)
}

// ConfigOfflineSaveAccount is a paid mutator transaction binding the contract method 0xdbda6212.
//
// Solidity: function configOfflineSaveAccount(address _offlineSave) returns()
func (_BridgeBank *BridgeBankTransactor) ConfigOfflineSaveAccount(opts *bind.TransactOpts, _offlineSave common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "configOfflineSaveAccount", _offlineSave)
}

// ConfigOfflineSaveAccount is a paid mutator transaction binding the contract method 0xdbda6212.
//
// Solidity: function configOfflineSaveAccount(address _offlineSave) returns()
func (_BridgeBank *BridgeBankSession) ConfigOfflineSaveAccount(_offlineSave common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ConfigOfflineSaveAccount(&_BridgeBank.TransactOpts, _offlineSave)
}

// ConfigOfflineSaveAccount is a paid mutator transaction binding the contract method 0xdbda6212.
//
// Solidity: function configOfflineSaveAccount(address _offlineSave) returns()
func (_BridgeBank *BridgeBankTransactorSession) ConfigOfflineSaveAccount(_offlineSave common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.ConfigOfflineSaveAccount(&_BridgeBank.TransactOpts, _offlineSave)
}

// ConfigplatformCoin is a paid mutator transaction binding the contract method 0x0b5ba475.
//
// Solidity: function configplatformCoin(string _name) returns()
func (_BridgeBank *BridgeBankTransactor) ConfigplatformCoin(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "configplatformCoin", _name)
}

// ConfigplatformCoin is a paid mutator transaction binding the contract method 0x0b5ba475.
//
// Solidity: function configplatformCoin(string _name) returns()
func (_BridgeBank *BridgeBankSession) ConfigplatformCoin(_name string) (*types.Transaction, error) {
	return _BridgeBank.Contract.ConfigplatformCoin(&_BridgeBank.TransactOpts, _name)
}

// ConfigplatformCoin is a paid mutator transaction binding the contract method 0x0b5ba475.
//
// Solidity: function configplatformCoin(string _name) returns()
func (_BridgeBank *BridgeBankTransactorSession) ConfigplatformCoin(_name string) (*types.Transaction, error) {
	return _BridgeBank.Contract.ConfigplatformCoin(&_BridgeBank.TransactOpts, _name)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0x44aea0de.
//
// Solidity: function createNewBridgeToken(string _name, string _symbol, uint8 decimals) returns(address)
func (_BridgeBank *BridgeBankTransactor) CreateNewBridgeToken(opts *bind.TransactOpts, _name string, _symbol string, decimals uint8) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "createNewBridgeToken", _name, _symbol, decimals)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0x44aea0de.
//
// Solidity: function createNewBridgeToken(string _name, string _symbol, uint8 decimals) returns(address)
func (_BridgeBank *BridgeBankSession) CreateNewBridgeToken(_name string, _symbol string, decimals uint8) (*types.Transaction, error) {
	return _BridgeBank.Contract.CreateNewBridgeToken(&_BridgeBank.TransactOpts, _name, _symbol, decimals)
}

// CreateNewBridgeToken is a paid mutator transaction binding the contract method 0x44aea0de.
//
// Solidity: function createNewBridgeToken(string _name, string _symbol, uint8 decimals) returns(address)
func (_BridgeBank *BridgeBankTransactorSession) CreateNewBridgeToken(_name string, _symbol string, decimals uint8) (*types.Transaction, error) {
	return _BridgeBank.Contract.CreateNewBridgeToken(&_BridgeBank.TransactOpts, _name, _symbol, decimals)
}

// EnableBridgeToken2Withdraw is a paid mutator transaction binding the contract method 0x800823a9.
//
// Solidity: function enableBridgeToken2Withdraw(address _bridgetoken, uint256 _toChainID) returns()
func (_BridgeBank *BridgeBankTransactor) EnableBridgeToken2Withdraw(opts *bind.TransactOpts, _bridgetoken common.Address, _toChainID *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "enableBridgeToken2Withdraw", _bridgetoken, _toChainID)
}

// EnableBridgeToken2Withdraw is a paid mutator transaction binding the contract method 0x800823a9.
//
// Solidity: function enableBridgeToken2Withdraw(address _bridgetoken, uint256 _toChainID) returns()
func (_BridgeBank *BridgeBankSession) EnableBridgeToken2Withdraw(_bridgetoken common.Address, _toChainID *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.EnableBridgeToken2Withdraw(&_BridgeBank.TransactOpts, _bridgetoken, _toChainID)
}

// EnableBridgeToken2Withdraw is a paid mutator transaction binding the contract method 0x800823a9.
//
// Solidity: function enableBridgeToken2Withdraw(address _bridgetoken, uint256 _toChainID) returns()
func (_BridgeBank *BridgeBankTransactorSession) EnableBridgeToken2Withdraw(_bridgetoken common.Address, _toChainID *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.EnableBridgeToken2Withdraw(&_BridgeBank.TransactOpts, _bridgetoken, _toChainID)
}

// Lock is a paid mutator transaction binding the contract method 0x7750c9f0.
//
// Solidity: function lock(address _recipient, address _token, uint256 _amount) payable returns()
func (_BridgeBank *BridgeBankTransactor) Lock(opts *bind.TransactOpts, _recipient common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "lock", _recipient, _token, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x7750c9f0.
//
// Solidity: function lock(address _recipient, address _token, uint256 _amount) payable returns()
func (_BridgeBank *BridgeBankSession) Lock(_recipient common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Lock(&_BridgeBank.TransactOpts, _recipient, _token, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x7750c9f0.
//
// Solidity: function lock(address _recipient, address _token, uint256 _amount) payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Lock(_recipient common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Lock(&_BridgeBank.TransactOpts, _recipient, _token, _amount)
}

// MintBridgeTokens is a paid mutator transaction binding the contract method 0x129699ae.
//
// Solidity: function mintBridgeTokens(address _intendedRecipient, address _bridgeTokenAddress, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactor) MintBridgeTokens(opts *bind.TransactOpts, _intendedRecipient common.Address, _bridgeTokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "mintBridgeTokens", _intendedRecipient, _bridgeTokenAddress, _amount)
}

// MintBridgeTokens is a paid mutator transaction binding the contract method 0x129699ae.
//
// Solidity: function mintBridgeTokens(address _intendedRecipient, address _bridgeTokenAddress, uint256 _amount) returns()
func (_BridgeBank *BridgeBankSession) MintBridgeTokens(_intendedRecipient common.Address, _bridgeTokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.MintBridgeTokens(&_BridgeBank.TransactOpts, _intendedRecipient, _bridgeTokenAddress, _amount)
}

// MintBridgeTokens is a paid mutator transaction binding the contract method 0x129699ae.
//
// Solidity: function mintBridgeTokens(address _intendedRecipient, address _bridgeTokenAddress, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) MintBridgeTokens(_intendedRecipient common.Address, _bridgeTokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.MintBridgeTokens(&_BridgeBank.TransactOpts, _intendedRecipient, _bridgeTokenAddress, _amount)
}

// SetBridgeServiceFee is a paid mutator transaction binding the contract method 0x7a8d31e5.
//
// Solidity: function setBridgeServiceFee(uint256 _bridgeServiceFee) returns()
func (_BridgeBank *BridgeBankTransactor) SetBridgeServiceFee(opts *bind.TransactOpts, _bridgeServiceFee *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "setBridgeServiceFee", _bridgeServiceFee)
}

// SetBridgeServiceFee is a paid mutator transaction binding the contract method 0x7a8d31e5.
//
// Solidity: function setBridgeServiceFee(uint256 _bridgeServiceFee) returns()
func (_BridgeBank *BridgeBankSession) SetBridgeServiceFee(_bridgeServiceFee *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetBridgeServiceFee(&_BridgeBank.TransactOpts, _bridgeServiceFee)
}

// SetBridgeServiceFee is a paid mutator transaction binding the contract method 0x7a8d31e5.
//
// Solidity: function setBridgeServiceFee(uint256 _bridgeServiceFee) returns()
func (_BridgeBank *BridgeBankTransactorSession) SetBridgeServiceFee(_bridgeServiceFee *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetBridgeServiceFee(&_BridgeBank.TransactOpts, _bridgeServiceFee)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_BridgeBank *BridgeBankTransactor) SetFeeReceiver(opts *bind.TransactOpts, _feeReceiver common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "setFeeReceiver", _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_BridgeBank *BridgeBankSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetFeeReceiver(&_BridgeBank.TransactOpts, _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_BridgeBank *BridgeBankTransactorSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetFeeReceiver(&_BridgeBank.TransactOpts, _feeReceiver)
}

// SetTokenDeployer is a paid mutator transaction binding the contract method 0x880183e9.
//
// Solidity: function setTokenDeployer(address _tokenDeployer) returns()
func (_BridgeBank *BridgeBankTransactor) SetTokenDeployer(opts *bind.TransactOpts, _tokenDeployer common.Address) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "setTokenDeployer", _tokenDeployer)
}

// SetTokenDeployer is a paid mutator transaction binding the contract method 0x880183e9.
//
// Solidity: function setTokenDeployer(address _tokenDeployer) returns()
func (_BridgeBank *BridgeBankSession) SetTokenDeployer(_tokenDeployer common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetTokenDeployer(&_BridgeBank.TransactOpts, _tokenDeployer)
}

// SetTokenDeployer is a paid mutator transaction binding the contract method 0x880183e9.
//
// Solidity: function setTokenDeployer(address _tokenDeployer) returns()
func (_BridgeBank *BridgeBankTransactorSession) SetTokenDeployer(_tokenDeployer common.Address) (*types.Transaction, error) {
	return _BridgeBank.Contract.SetTokenDeployer(&_BridgeBank.TransactOpts, _tokenDeployer)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f89c91c.
//
// Solidity: function unlock(address _recipient, address _token, string _name, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactor) Unlock(opts *bind.TransactOpts, _recipient common.Address, _token common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.contract.Transact(opts, "unlock", _recipient, _token, _name, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f89c91c.
//
// Solidity: function unlock(address _recipient, address _token, string _name, uint256 _amount) returns()
func (_BridgeBank *BridgeBankSession) Unlock(_recipient common.Address, _token common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Unlock(&_BridgeBank.TransactOpts, _recipient, _token, _name, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f89c91c.
//
// Solidity: function unlock(address _recipient, address _token, string _name, uint256 _amount) returns()
func (_BridgeBank *BridgeBankTransactorSession) Unlock(_recipient common.Address, _token common.Address, _name string, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeBank.Contract.Unlock(&_BridgeBank.TransactOpts, _recipient, _token, _name, _amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.Contract.Fallback(&_BridgeBank.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBank.Contract.Fallback(&_BridgeBank.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBank *BridgeBankTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBank.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBank *BridgeBankSession) Receive() (*types.Transaction, error) {
	return _BridgeBank.Contract.Receive(&_BridgeBank.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBank *BridgeBankTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeBank.Contract.Receive(&_BridgeBank.TransactOpts)
}

// BridgeBankLogBridgeTokenMintIterator is returned from FilterLogBridgeTokenMint and is used to iterate over the raw logs and unpacked data for LogBridgeTokenMint events raised by the BridgeBank contract.
type BridgeBankLogBridgeTokenMintIterator struct {
	Event *BridgeBankLogBridgeTokenMint // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogBridgeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogBridgeTokenMint)
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
		it.Event = new(BridgeBankLogBridgeTokenMint)
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
func (it *BridgeBankLogBridgeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogBridgeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogBridgeTokenMint represents a LogBridgeTokenMint event raised by the BridgeBank contract.
type BridgeBankLogBridgeTokenMint struct {
	Token       common.Address
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeTokenMint is a free log retrieval operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) FilterLogBridgeTokenMint(opts *bind.FilterOpts) (*BridgeBankLogBridgeTokenMintIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogBridgeTokenMintIterator{contract: _BridgeBank.contract, event: "LogBridgeTokenMint", logs: logs, sub: sub}, nil
}

// WatchLogBridgeTokenMint is a free log subscription operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) WatchLogBridgeTokenMint(opts *bind.WatchOpts, sink chan<- *BridgeBankLogBridgeTokenMint) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogBridgeTokenMint)
				if err := _BridgeBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
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

// ParseLogBridgeTokenMint is a log parse operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_BridgeBank *BridgeBankFilterer) ParseLogBridgeTokenMint(log types.Log) (*BridgeBankLogBridgeTokenMint, error) {
	event := new(BridgeBankLogBridgeTokenMint)
	if err := _BridgeBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogBtcTokenBurnIterator is returned from FilterLogBtcTokenBurn and is used to iterate over the raw logs and unpacked data for LogBtcTokenBurn events raised by the BridgeBank contract.
type BridgeBankLogBtcTokenBurnIterator struct {
	Event *BridgeBankLogBtcTokenBurn // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogBtcTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogBtcTokenBurn)
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
		it.Event = new(BridgeBankLogBtcTokenBurn)
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
func (it *BridgeBankLogBtcTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogBtcTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogBtcTokenBurn represents a LogBtcTokenBurn event raised by the BridgeBank contract.
type BridgeBankLogBtcTokenBurn struct {
	ChainID   *big.Int
	Token     common.Address
	Name      string
	Amount    *big.Int
	OwnerFrom common.Address
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogBtcTokenBurn is a free log retrieval operation binding the contract event 0x6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e.
//
// Solidity: event LogBtcTokenBurn(uint256 _chainID, address _token, string _name, uint256 _amount, address _ownerFrom, address _receiver)
func (_BridgeBank *BridgeBankFilterer) FilterLogBtcTokenBurn(opts *bind.FilterOpts) (*BridgeBankLogBtcTokenBurnIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogBtcTokenBurn")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogBtcTokenBurnIterator{contract: _BridgeBank.contract, event: "LogBtcTokenBurn", logs: logs, sub: sub}, nil
}

// WatchLogBtcTokenBurn is a free log subscription operation binding the contract event 0x6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e.
//
// Solidity: event LogBtcTokenBurn(uint256 _chainID, address _token, string _name, uint256 _amount, address _ownerFrom, address _receiver)
func (_BridgeBank *BridgeBankFilterer) WatchLogBtcTokenBurn(opts *bind.WatchOpts, sink chan<- *BridgeBankLogBtcTokenBurn) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogBtcTokenBurn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogBtcTokenBurn)
				if err := _BridgeBank.contract.UnpackLog(event, "LogBtcTokenBurn", log); err != nil {
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

// ParseLogBtcTokenBurn is a log parse operation binding the contract event 0x6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e.
//
// Solidity: event LogBtcTokenBurn(uint256 _chainID, address _token, string _name, uint256 _amount, address _ownerFrom, address _receiver)
func (_BridgeBank *BridgeBankFilterer) ParseLogBtcTokenBurn(log types.Log) (*BridgeBankLogBtcTokenBurn, error) {
	event := new(BridgeBankLogBtcTokenBurn)
	if err := _BridgeBank.contract.UnpackLog(event, "LogBtcTokenBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogLockIterator is returned from FilterLogLock and is used to iterate over the raw logs and unpacked data for LogLock events raised by the BridgeBank contract.
type BridgeBankLogLockIterator struct {
	Event *BridgeBankLogLock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogLock)
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
		it.Event = new(BridgeBankLogLock)
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
func (it *BridgeBankLogLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogLock represents a LogLock event raised by the BridgeBank contract.
type BridgeBankLogLock struct {
	From  common.Address
	To    common.Address
	Token common.Address
	Name  string
	Value *big.Int
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogLock is a free log retrieval operation binding the contract event 0x4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe.
//
// Solidity: event LogLock(address _from, address _to, address _token, string _name, uint256 _value, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) FilterLogLock(opts *bind.FilterOpts) (*BridgeBankLogLockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogLockIterator{contract: _BridgeBank.contract, event: "LogLock", logs: logs, sub: sub}, nil
}

// WatchLogLock is a free log subscription operation binding the contract event 0x4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe.
//
// Solidity: event LogLock(address _from, address _to, address _token, string _name, uint256 _value, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) WatchLogLock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogLock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogLock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
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

// ParseLogLock is a log parse operation binding the contract event 0x4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe.
//
// Solidity: event LogLock(address _from, address _to, address _token, string _name, uint256 _value, uint256 _nonce)
func (_BridgeBank *BridgeBankFilterer) ParseLogLock(log types.Log) (*BridgeBankLogLock, error) {
	event := new(BridgeBankLogLock)
	if err := _BridgeBank.contract.UnpackLog(event, "LogLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogNewBridgeTokenIterator is returned from FilterLogNewBridgeToken and is used to iterate over the raw logs and unpacked data for LogNewBridgeToken events raised by the BridgeBank contract.
type BridgeBankLogNewBridgeTokenIterator struct {
	Event *BridgeBankLogNewBridgeToken // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogNewBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogNewBridgeToken)
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
		it.Event = new(BridgeBankLogNewBridgeToken)
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
func (it *BridgeBankLogNewBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogNewBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogNewBridgeToken represents a LogNewBridgeToken event raised by the BridgeBank contract.
type BridgeBankLogNewBridgeToken struct {
	Token    common.Address
	Name     string
	Symbol   string
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNewBridgeToken is a free log retrieval operation binding the contract event 0x438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d7731.
//
// Solidity: event LogNewBridgeToken(address _token, string _name, string _symbol, uint8 decimals)
func (_BridgeBank *BridgeBankFilterer) FilterLogNewBridgeToken(opts *bind.FilterOpts) (*BridgeBankLogNewBridgeTokenIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogNewBridgeToken")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogNewBridgeTokenIterator{contract: _BridgeBank.contract, event: "LogNewBridgeToken", logs: logs, sub: sub}, nil
}

// WatchLogNewBridgeToken is a free log subscription operation binding the contract event 0x438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d7731.
//
// Solidity: event LogNewBridgeToken(address _token, string _name, string _symbol, uint8 decimals)
func (_BridgeBank *BridgeBankFilterer) WatchLogNewBridgeToken(opts *bind.WatchOpts, sink chan<- *BridgeBankLogNewBridgeToken) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogNewBridgeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogNewBridgeToken)
				if err := _BridgeBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
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

// ParseLogNewBridgeToken is a log parse operation binding the contract event 0x438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d7731.
//
// Solidity: event LogNewBridgeToken(address _token, string _name, string _symbol, uint8 decimals)
func (_BridgeBank *BridgeBankFilterer) ParseLogNewBridgeToken(log types.Log) (*BridgeBankLogNewBridgeToken, error) {
	event := new(BridgeBankLogNewBridgeToken)
	if err := _BridgeBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBankLogUnlockIterator is returned from FilterLogUnlock and is used to iterate over the raw logs and unpacked data for LogUnlock events raised by the BridgeBank contract.
type BridgeBankLogUnlockIterator struct {
	Event *BridgeBankLogUnlock // Event containing the contract specifics and raw log

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
func (it *BridgeBankLogUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBankLogUnlock)
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
		it.Event = new(BridgeBankLogUnlock)
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
func (it *BridgeBankLogUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBankLogUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBankLogUnlock represents a LogUnlock event raised by the BridgeBank contract.
type BridgeBankLogUnlock struct {
	To    common.Address
	Token common.Address
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogUnlock is a free log retrieval operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _name, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) FilterLogUnlock(opts *bind.FilterOpts) (*BridgeBankLogUnlockIterator, error) {

	logs, sub, err := _BridgeBank.contract.FilterLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return &BridgeBankLogUnlockIterator{contract: _BridgeBank.contract, event: "LogUnlock", logs: logs, sub: sub}, nil
}

// WatchLogUnlock is a free log subscription operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _name, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) WatchLogUnlock(opts *bind.WatchOpts, sink chan<- *BridgeBankLogUnlock) (event.Subscription, error) {

	logs, sub, err := _BridgeBank.contract.WatchLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBankLogUnlock)
				if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
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

// ParseLogUnlock is a log parse operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _name, uint256 _value)
func (_BridgeBank *BridgeBankFilterer) ParseLogUnlock(log types.Log) (*BridgeBankLogUnlock, error) {
	event := new(BridgeBankLogUnlock)
	if err := _BridgeBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenMetaData contains all meta data concerning the BridgeToken contract.
var BridgeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b50604051620010143803806200101483398101604081905262000033916200020d565b8282600362000043838262000318565b50600462000052828262000318565b50505060ff8116608052620000685f336200009f565b50620000957f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6336200009f565b50505050620003e0565b5f8281526005602090815260408083206001600160a01b038516845290915281205460ff1662000146575f8381526005602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620000fd3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600162000149565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011262000173575f80fd5b81516001600160401b03808211156200019057620001906200014f565b604051601f8301601f19908116603f01168101908282118183101715620001bb57620001bb6200014f565b81604052838152602092508683858801011115620001d7575f80fd5b5f91505b83821015620001fa5785820183015181830184015290820190620001db565b5f93810190920192909252949350505050565b5f805f6060848603121562000220575f80fd5b83516001600160401b038082111562000237575f80fd5b620002458783880162000163565b945060208601519150808211156200025b575f80fd5b506200026a8682870162000163565b925050604084015160ff8116811462000281575f80fd5b809150509250925092565b600181811c90821680620002a157607f821691505b602082108103620002c057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000313575f81815260208120601f850160051c81016020861015620002ee5750805b601f850160051c820191505b818110156200030f57828155600101620002fa565b5050505b505050565b81516001600160401b038111156200033457620003346200014f565b6200034c816200034584546200028c565b84620002c6565b602080601f83116001811462000382575f84156200036a5750858301515b5f19600386901b1c1916600185901b1785556200030f565b5f85815260208120601f198616915b82811015620003b25788860151825594840194600190910190840162000391565b5085821015620003d057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b608051610c1b620003f95f395f6101e90152610c1b5ff3fe608060405234801561000f575f80fd5b5060043610610132575f3560e01c806342966c68116100b4578063a217fddf11610079578063a217fddf146102a2578063a9059cbb146102a9578063aa271e1a146102bc578063d5391393146102cf578063d547741f146102f6578063dd62ed3e14610309575f80fd5b806342966c681461023957806370a082311461024c57806379cc67901461027457806391d148541461028757806395d89b411461029a575f80fd5b8063248a9ca3116100fa578063248a9ca3146101ab5780632f2ff15d146101cd578063313ce567146101e257806336568abe1461021357806340c10f1914610226575f80fd5b806301ffc9a71461013657806306fdde031461015e578063095ea7b31461017357806318160ddd1461018657806323b872dd14610198575b5f80fd5b610149610144366004610a17565b610341565b60405190151581526020015b60405180910390f35b610166610377565b6040516101559190610a45565b610149610181366004610aab565b610407565b6002545b604051908152602001610155565b6101496101a6366004610ad3565b61041e565b61018a6101b9366004610b0c565b5f9081526005602052604090206001015490565b6101e06101db366004610b23565b610441565b005b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610155565b6101e0610221366004610b23565b61046b565b610149610234366004610aab565b6104a3565b6101e0610247366004610b0c565b6104d8565b61018a61025a366004610b4d565b6001600160a01b03165f9081526020819052604090205490565b6101e0610282366004610aab565b6104e5565b610149610295366004610b23565b6104fe565b610166610528565b61018a5f81565b6101496102b7366004610aab565b610537565b6101496102ca366004610b4d565b610544565b61018a7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6101e0610304366004610b23565b61056f565b61018a610317366004610b66565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b5f6001600160e01b03198216637965db0b60e01b148061037157506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606003805461038690610b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546103b290610b8e565b80156103fd5780601f106103d4576101008083540402835291602001916103fd565b820191905f5260205f20905b8154815290600101906020018083116103e057829003601f168201915b5050505050905090565b5f33610414818585610593565b5060019392505050565b5f3361042b8582856105a0565b61043685858561061b565b506001949350505050565b5f8281526005602052604090206001015461045b81610678565b6104658383610682565b50505050565b6001600160a01b03811633146104945760405163334bd91960e11b815260040160405180910390fd5b61049e8282610713565b505050565b5f7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66104ce81610678565b610414848461077e565b6104e233826107b2565b50565b6104f08233836105a0565b6104fa82826107b2565b5050565b5f9182526005602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606004805461038690610b8e565b5f3361041481858561061b565b5f6103717f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6836104fe565b5f8281526005602052604090206001015461058981610678565b6104658383610713565b61049e83838360016107e6565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811015610465578181101561060d57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61046584848484035f6107e6565b6001600160a01b03831661064457604051634b637e8f60e11b81525f6004820152602401610604565b6001600160a01b03821661066d5760405163ec442f0560e01b81525f6004820152602401610604565b61049e8383836108b8565b6104e281336109de565b5f61068d83836104fe565b61070c575f8381526005602090815260408083206001600160a01b03861684529091529020805460ff191660011790556106c43390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610371565b505f610371565b5f61071e83836104fe565b1561070c575f8381526005602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610371565b6001600160a01b0382166107a75760405163ec442f0560e01b81525f6004820152602401610604565b6104fa5f83836108b8565b6001600160a01b0382166107db57604051634b637e8f60e11b81525f6004820152602401610604565b6104fa825f836108b8565b6001600160a01b03841661080f5760405163e602df0560e01b81525f6004820152602401610604565b6001600160a01b03831661083857604051634a1406b160e11b81525f6004820152602401610604565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561046557826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516108aa91815260200190565b60405180910390a350505050565b6001600160a01b0383166108e2578060025f8282546108d79190610bc6565b909155506109529050565b6001600160a01b0383165f90815260208190526040902054818110156109345760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610604565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661096e5760028054829003905561098c565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516109d191815260200190565b60405180910390a3505050565b6109e882826104fe565b6104fa5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610604565b5f60208284031215610a27575f80fd5b81356001600160e01b031981168114610a3e575f80fd5b9392505050565b5f6020808352835180828501525f5b81811015610a7057858101830151858201604001528201610a54565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610aa6575f80fd5b919050565b5f8060408385031215610abc575f80fd5b610ac583610a90565b946020939093013593505050565b5f805f60608486031215610ae5575f80fd5b610aee84610a90565b9250610afc60208501610a90565b9150604084013590509250925092565b5f60208284031215610b1c575f80fd5b5035919050565b5f8060408385031215610b34575f80fd5b82359150610b4460208401610a90565b90509250929050565b5f60208284031215610b5d575f80fd5b610a3e82610a90565b5f8060408385031215610b77575f80fd5b610b8083610a90565b9150610b4460208401610a90565b600181811c90821680610ba257607f821691505b602082108103610bc057634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561037157634e487b7160e01b5f52601160045260245ffdfea26469706673582212207e983d1a5d95ef0a10825dcc962d6f4842fb299e188d00202589abd9c1aaa83464736f6c63430008140033",
}

// BridgeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTokenMetaData.ABI instead.
var BridgeTokenABI = BridgeTokenMetaData.ABI

// BridgeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTokenMetaData.Bin instead.
var BridgeTokenBin = BridgeTokenMetaData.Bin

// DeployBridgeToken deploys a new Ethereum contract, binding an instance of BridgeToken to it.
func DeployBridgeToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *BridgeToken, error) {
	parsed, err := BridgeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTokenBin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeToken{BridgeTokenCaller: BridgeTokenCaller{contract: contract}, BridgeTokenTransactor: BridgeTokenTransactor{contract: contract}, BridgeTokenFilterer: BridgeTokenFilterer{contract: contract}}, nil
}

// BridgeToken is an auto generated Go binding around an Ethereum contract.
type BridgeToken struct {
	BridgeTokenCaller     // Read-only binding to the contract
	BridgeTokenTransactor // Write-only binding to the contract
	BridgeTokenFilterer   // Log filterer for contract events
}

// BridgeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTokenSession struct {
	Contract     *BridgeToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTokenCallerSession struct {
	Contract *BridgeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTokenTransactorSession struct {
	Contract     *BridgeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTokenRaw struct {
	Contract *BridgeToken // Generic contract binding to access the raw methods on
}

// BridgeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTokenCallerRaw struct {
	Contract *BridgeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTokenTransactorRaw struct {
	Contract *BridgeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeToken creates a new instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeToken(address common.Address, backend bind.ContractBackend) (*BridgeToken, error) {
	contract, err := bindBridgeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeToken{BridgeTokenCaller: BridgeTokenCaller{contract: contract}, BridgeTokenTransactor: BridgeTokenTransactor{contract: contract}, BridgeTokenFilterer: BridgeTokenFilterer{contract: contract}}, nil
}

// NewBridgeTokenCaller creates a new read-only instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenCaller(address common.Address, caller bind.ContractCaller) (*BridgeTokenCaller, error) {
	contract, err := bindBridgeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenCaller{contract: contract}, nil
}

// NewBridgeTokenTransactor creates a new write-only instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTokenTransactor, error) {
	contract, err := bindBridgeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenTransactor{contract: contract}, nil
}

// NewBridgeTokenFilterer creates a new log filterer instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTokenFilterer, error) {
	contract, err := bindBridgeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenFilterer{contract: contract}, nil
}

// bindBridgeToken binds a generic wrapper to an already deployed contract.
func bindBridgeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeToken *BridgeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeToken.Contract.BridgeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeToken *BridgeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.Contract.BridgeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeToken *BridgeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeToken.Contract.BridgeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeToken *BridgeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeToken *BridgeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeToken *BridgeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeToken *BridgeTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeToken *BridgeTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeToken.Contract.DEFAULTADMINROLE(&_BridgeToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeToken *BridgeTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeToken.Contract.DEFAULTADMINROLE(&_BridgeToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BridgeToken *BridgeTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BridgeToken *BridgeTokenSession) MINTERROLE() ([32]byte, error) {
	return _BridgeToken.Contract.MINTERROLE(&_BridgeToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_BridgeToken *BridgeTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _BridgeToken.Contract.MINTERROLE(&_BridgeToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeToken *BridgeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeToken *BridgeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.Allowance(&_BridgeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.Allowance(&_BridgeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeToken *BridgeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeToken *BridgeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.BalanceOf(&_BridgeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.BalanceOf(&_BridgeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BridgeToken *BridgeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BridgeToken *BridgeTokenSession) Decimals() (uint8, error) {
	return _BridgeToken.Contract.Decimals(&_BridgeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BridgeToken *BridgeTokenCallerSession) Decimals() (uint8, error) {
	return _BridgeToken.Contract.Decimals(&_BridgeToken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeToken *BridgeTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeToken *BridgeTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeToken.Contract.GetRoleAdmin(&_BridgeToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeToken *BridgeTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeToken.Contract.GetRoleAdmin(&_BridgeToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeToken *BridgeTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeToken *BridgeTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeToken.Contract.HasRole(&_BridgeToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeToken.Contract.HasRole(&_BridgeToken.CallOpts, role, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_BridgeToken *BridgeTokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_BridgeToken *BridgeTokenSession) IsMinter(account common.Address) (bool, error) {
	return _BridgeToken.Contract.IsMinter(&_BridgeToken.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _BridgeToken.Contract.IsMinter(&_BridgeToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeToken *BridgeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeToken *BridgeTokenSession) Name() (string, error) {
	return _BridgeToken.Contract.Name(&_BridgeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeToken *BridgeTokenCallerSession) Name() (string, error) {
	return _BridgeToken.Contract.Name(&_BridgeToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeToken *BridgeTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeToken *BridgeTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeToken.Contract.SupportsInterface(&_BridgeToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeToken.Contract.SupportsInterface(&_BridgeToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeToken *BridgeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeToken *BridgeTokenSession) Symbol() (string, error) {
	return _BridgeToken.Contract.Symbol(&_BridgeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeToken *BridgeTokenCallerSession) Symbol() (string, error) {
	return _BridgeToken.Contract.Symbol(&_BridgeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeToken *BridgeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeToken *BridgeTokenSession) TotalSupply() (*big.Int, error) {
	return _BridgeToken.Contract.TotalSupply(&_BridgeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BridgeToken.Contract.TotalSupply(&_BridgeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Approve(&_BridgeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Approve(&_BridgeToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BridgeToken *BridgeTokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BridgeToken *BridgeTokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Burn(&_BridgeToken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_BridgeToken *BridgeTokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Burn(&_BridgeToken.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BridgeToken *BridgeTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BridgeToken *BridgeTokenSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.BurnFrom(&_BridgeToken.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_BridgeToken *BridgeTokenTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.BurnFrom(&_BridgeToken.TransactOpts, account, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeToken *BridgeTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeToken *BridgeTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.GrantRole(&_BridgeToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeToken *BridgeTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.GrantRole(&_BridgeToken.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_BridgeToken *BridgeTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Mint(&_BridgeToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Mint(&_BridgeToken.TransactOpts, to, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeToken *BridgeTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeToken *BridgeTokenSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.RenounceRole(&_BridgeToken.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeToken *BridgeTokenTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.RenounceRole(&_BridgeToken.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeToken *BridgeTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeToken *BridgeTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.RevokeRole(&_BridgeToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeToken *BridgeTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.RevokeRole(&_BridgeToken.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Transfer(&_BridgeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Transfer(&_BridgeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.TransferFrom(&_BridgeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.TransferFrom(&_BridgeToken.TransactOpts, from, to, value)
}

// BridgeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BridgeToken contract.
type BridgeTokenApprovalIterator struct {
	Event *BridgeTokenApproval // Event containing the contract specifics and raw log

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
func (it *BridgeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenApproval)
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
		it.Event = new(BridgeTokenApproval)
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
func (it *BridgeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenApproval represents a Approval event raised by the BridgeToken contract.
type BridgeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeToken *BridgeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BridgeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenApprovalIterator{contract: _BridgeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeToken *BridgeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BridgeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenApproval)
				if err := _BridgeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BridgeToken *BridgeTokenFilterer) ParseApproval(log types.Log) (*BridgeTokenApproval, error) {
	event := new(BridgeTokenApproval)
	if err := _BridgeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeToken contract.
type BridgeTokenRoleAdminChangedIterator struct {
	Event *BridgeTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenRoleAdminChanged)
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
		it.Event = new(BridgeTokenRoleAdminChanged)
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
func (it *BridgeTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeToken contract.
type BridgeTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeToken *BridgeTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenRoleAdminChangedIterator{contract: _BridgeToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeToken *BridgeTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenRoleAdminChanged)
				if err := _BridgeToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeToken *BridgeTokenFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeTokenRoleAdminChanged, error) {
	event := new(BridgeTokenRoleAdminChanged)
	if err := _BridgeToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeToken contract.
type BridgeTokenRoleGrantedIterator struct {
	Event *BridgeTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenRoleGranted)
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
		it.Event = new(BridgeTokenRoleGranted)
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
func (it *BridgeTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenRoleGranted represents a RoleGranted event raised by the BridgeToken contract.
type BridgeTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeToken *BridgeTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenRoleGrantedIterator{contract: _BridgeToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeToken *BridgeTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenRoleGranted)
				if err := _BridgeToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeToken *BridgeTokenFilterer) ParseRoleGranted(log types.Log) (*BridgeTokenRoleGranted, error) {
	event := new(BridgeTokenRoleGranted)
	if err := _BridgeToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeToken contract.
type BridgeTokenRoleRevokedIterator struct {
	Event *BridgeTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenRoleRevoked)
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
		it.Event = new(BridgeTokenRoleRevoked)
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
func (it *BridgeTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenRoleRevoked represents a RoleRevoked event raised by the BridgeToken contract.
type BridgeTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeToken *BridgeTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenRoleRevokedIterator{contract: _BridgeToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeToken *BridgeTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenRoleRevoked)
				if err := _BridgeToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeToken *BridgeTokenFilterer) ParseRoleRevoked(log types.Log) (*BridgeTokenRoleRevoked, error) {
	event := new(BridgeTokenRoleRevoked)
	if err := _BridgeToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BridgeToken contract.
type BridgeTokenTransferIterator struct {
	Event *BridgeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BridgeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenTransfer)
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
		it.Event = new(BridgeTokenTransfer)
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
func (it *BridgeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenTransfer represents a Transfer event raised by the BridgeToken contract.
type BridgeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeToken *BridgeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BridgeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenTransferIterator{contract: _BridgeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeToken *BridgeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BridgeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenTransfer)
				if err := _BridgeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BridgeToken *BridgeTokenFilterer) ParseTransfer(log types.Log) (*BridgeTokenTransfer, error) {
	event := new(BridgeTokenTransfer)
	if err := _BridgeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212208b6fb0c80a71929ce34f0ab9da38313378c0e20e9b56702e96eb838b08b50eaf64736f6c63430008140033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableMetaData contains all meta data concerning the ERC20Burnable contract.
var ERC20BurnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableMetaData.ABI instead.
var ERC20BurnableABI = ERC20BurnableMetaData.ABI

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20BurnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, from, to, value)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBankMetaData contains all meta data concerning the EthereumBank contract.
var EthereumBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"LogBridgeTokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ownerFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"LogBtcTokenBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"LogNewBridgeToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeTokenCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgeTokenWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getToken2address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getToken2addressV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"hasBridgeTokenCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"token2WithdrawCfg\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"token2address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b505f80556104ad806100205f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c806370e2a8d91161005857806370e2a8d91461012a57806385f0bdc714610155578063c4c2709e1461017d578063efea27ab1461019f575f80fd5b80631c4336a514610089578063328470ab146100c05780633f4d5681146100d657806359bf3a3314610108575b5f80fd5b61009c61009736600461031b565b6101b2565b604080516001600160a01b0390931683529015156020830152015b60405180910390f35b6100c85f5481565b6040519081526020016100b7565b6100f86100e43660046103c6565b60016020525f908152604090205460ff1681565b60405190151581526020016100b7565b6100f86101163660046103f3565b60026020525f908152604090205460ff1681565b61013d61013836600461031b565b610258565b6040516001600160a01b0390911681526020016100b7565b61013d6101633660046103f3565b60036020525f90815260409020546001600160a01b031681565b6100f861018b3660046103f3565b60046020525f908152604090205460ff1681565b6100f86101ad36600461031b565b6102c6565b5f80826101be816102c6565b6101e35760405162461bcd60e51b81526004016101da9061040a565b60405180910390fd5b5f846040516020016101f5919061044b565b60408051601f1981840301815291815281516020928301205f818152600290935291205490915060ff161515600114610234575f809350935050610252565b5f908152600360205260409020546001600160a01b03169250600191505b50915091565b5f81610263816102c6565b61027f5760405162461bcd60e51b81526004016101da9061040a565b5f83604051602001610291919061044b565b60408051808303601f1901815291815281516020928301205f90815260039092529020546001600160a01b0316949350505050565b5f80826040516020016102d9919061044b565b60408051601f1981840301815291815281516020928301205f908152600290925290205460ff169392505050565b634e487b7160e01b5f52604160045260245ffd5b5f6020828403121561032b575f80fd5b813567ffffffffffffffff80821115610342575f80fd5b818401915084601f830112610355575f80fd5b81358181111561036757610367610307565b604051601f8201601f19908116603f0116810190838211818310171561038f5761038f610307565b816040528281528760208487010111156103a7575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f602082840312156103d6575f80fd5b81356001600160a01b03811681146103ec575f80fd5b9392505050565b5f60208284031215610403575f80fd5b5035919050565b60208082526021908201527f546865206e616d6520686173206e6f74206265656e20637265617465642079656040820152601d60fa1b606082015260800190565b5f82515f5b8181101561046a5760208186018101518583015201610450565b505f92019182525091905056fea26469706673582212207847bae5f4185f24b00333abdcb627b984323abc4dc9704d6545ebeb9d63c38064736f6c63430008140033",
}

// EthereumBankABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumBankMetaData.ABI instead.
var EthereumBankABI = EthereumBankMetaData.ABI

// EthereumBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthereumBankMetaData.Bin instead.
var EthereumBankBin = EthereumBankMetaData.Bin

// DeployEthereumBank deploys a new Ethereum contract, binding an instance of EthereumBank to it.
func DeployEthereumBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthereumBank, error) {
	parsed, err := EthereumBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthereumBankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthereumBank{EthereumBankCaller: EthereumBankCaller{contract: contract}, EthereumBankTransactor: EthereumBankTransactor{contract: contract}, EthereumBankFilterer: EthereumBankFilterer{contract: contract}}, nil
}

// EthereumBank is an auto generated Go binding around an Ethereum contract.
type EthereumBank struct {
	EthereumBankCaller     // Read-only binding to the contract
	EthereumBankTransactor // Write-only binding to the contract
	EthereumBankFilterer   // Log filterer for contract events
}

// EthereumBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumBankSession struct {
	Contract     *EthereumBank     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumBankCallerSession struct {
	Contract *EthereumBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EthereumBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumBankTransactorSession struct {
	Contract     *EthereumBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EthereumBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumBankRaw struct {
	Contract *EthereumBank // Generic contract binding to access the raw methods on
}

// EthereumBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumBankCallerRaw struct {
	Contract *EthereumBankCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumBankTransactorRaw struct {
	Contract *EthereumBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereumBank creates a new instance of EthereumBank, bound to a specific deployed contract.
func NewEthereumBank(address common.Address, backend bind.ContractBackend) (*EthereumBank, error) {
	contract, err := bindEthereumBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthereumBank{EthereumBankCaller: EthereumBankCaller{contract: contract}, EthereumBankTransactor: EthereumBankTransactor{contract: contract}, EthereumBankFilterer: EthereumBankFilterer{contract: contract}}, nil
}

// NewEthereumBankCaller creates a new read-only instance of EthereumBank, bound to a specific deployed contract.
func NewEthereumBankCaller(address common.Address, caller bind.ContractCaller) (*EthereumBankCaller, error) {
	contract, err := bindEthereumBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumBankCaller{contract: contract}, nil
}

// NewEthereumBankTransactor creates a new write-only instance of EthereumBank, bound to a specific deployed contract.
func NewEthereumBankTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumBankTransactor, error) {
	contract, err := bindEthereumBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumBankTransactor{contract: contract}, nil
}

// NewEthereumBankFilterer creates a new log filterer instance of EthereumBank, bound to a specific deployed contract.
func NewEthereumBankFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumBankFilterer, error) {
	contract, err := bindEthereumBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumBankFilterer{contract: contract}, nil
}

// bindEthereumBank binds a generic wrapper to an already deployed contract.
func bindEthereumBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthereumBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumBank *EthereumBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumBank.Contract.EthereumBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumBank *EthereumBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBank.Contract.EthereumBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumBank *EthereumBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumBank.Contract.EthereumBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthereumBank *EthereumBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthereumBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthereumBank *EthereumBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthereumBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthereumBank *EthereumBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthereumBank.Contract.contract.Transact(opts, method, params...)
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() view returns(uint256)
func (_EthereumBank *EthereumBankCaller) BridgeTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "bridgeTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() view returns(uint256)
func (_EthereumBank *EthereumBankSession) BridgeTokenCount() (*big.Int, error) {
	return _EthereumBank.Contract.BridgeTokenCount(&_EthereumBank.CallOpts)
}

// BridgeTokenCount is a free data retrieval call binding the contract method 0x328470ab.
//
// Solidity: function bridgeTokenCount() view returns(uint256)
func (_EthereumBank *EthereumBankCallerSession) BridgeTokenCount() (*big.Int, error) {
	return _EthereumBank.Contract.BridgeTokenCount(&_EthereumBank.CallOpts)
}

// BridgeTokenCreated is a free data retrieval call binding the contract method 0x59bf3a33.
//
// Solidity: function bridgeTokenCreated(bytes32 ) view returns(bool)
func (_EthereumBank *EthereumBankCaller) BridgeTokenCreated(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "bridgeTokenCreated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeTokenCreated is a free data retrieval call binding the contract method 0x59bf3a33.
//
// Solidity: function bridgeTokenCreated(bytes32 ) view returns(bool)
func (_EthereumBank *EthereumBankSession) BridgeTokenCreated(arg0 [32]byte) (bool, error) {
	return _EthereumBank.Contract.BridgeTokenCreated(&_EthereumBank.CallOpts, arg0)
}

// BridgeTokenCreated is a free data retrieval call binding the contract method 0x59bf3a33.
//
// Solidity: function bridgeTokenCreated(bytes32 ) view returns(bool)
func (_EthereumBank *EthereumBankCallerSession) BridgeTokenCreated(arg0 [32]byte) (bool, error) {
	return _EthereumBank.Contract.BridgeTokenCreated(&_EthereumBank.CallOpts, arg0)
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) view returns(bool)
func (_EthereumBank *EthereumBankCaller) BridgeTokenWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "bridgeTokenWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) view returns(bool)
func (_EthereumBank *EthereumBankSession) BridgeTokenWhitelist(arg0 common.Address) (bool, error) {
	return _EthereumBank.Contract.BridgeTokenWhitelist(&_EthereumBank.CallOpts, arg0)
}

// BridgeTokenWhitelist is a free data retrieval call binding the contract method 0x3f4d5681.
//
// Solidity: function bridgeTokenWhitelist(address ) view returns(bool)
func (_EthereumBank *EthereumBankCallerSession) BridgeTokenWhitelist(arg0 common.Address) (bool, error) {
	return _EthereumBank.Contract.BridgeTokenWhitelist(&_EthereumBank.CallOpts, arg0)
}

// GetToken2address is a free data retrieval call binding the contract method 0x70e2a8d9.
//
// Solidity: function getToken2address(string _name) view returns(address)
func (_EthereumBank *EthereumBankCaller) GetToken2address(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "getToken2address", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken2address is a free data retrieval call binding the contract method 0x70e2a8d9.
//
// Solidity: function getToken2address(string _name) view returns(address)
func (_EthereumBank *EthereumBankSession) GetToken2address(_name string) (common.Address, error) {
	return _EthereumBank.Contract.GetToken2address(&_EthereumBank.CallOpts, _name)
}

// GetToken2address is a free data retrieval call binding the contract method 0x70e2a8d9.
//
// Solidity: function getToken2address(string _name) view returns(address)
func (_EthereumBank *EthereumBankCallerSession) GetToken2address(_name string) (common.Address, error) {
	return _EthereumBank.Contract.GetToken2address(&_EthereumBank.CallOpts, _name)
}

// GetToken2addressV2 is a free data retrieval call binding the contract method 0x1c4336a5.
//
// Solidity: function getToken2addressV2(string _name) view returns(address, bool)
func (_EthereumBank *EthereumBankCaller) GetToken2addressV2(opts *bind.CallOpts, _name string) (common.Address, bool, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "getToken2addressV2", _name)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetToken2addressV2 is a free data retrieval call binding the contract method 0x1c4336a5.
//
// Solidity: function getToken2addressV2(string _name) view returns(address, bool)
func (_EthereumBank *EthereumBankSession) GetToken2addressV2(_name string) (common.Address, bool, error) {
	return _EthereumBank.Contract.GetToken2addressV2(&_EthereumBank.CallOpts, _name)
}

// GetToken2addressV2 is a free data retrieval call binding the contract method 0x1c4336a5.
//
// Solidity: function getToken2addressV2(string _name) view returns(address, bool)
func (_EthereumBank *EthereumBankCallerSession) GetToken2addressV2(_name string) (common.Address, bool, error) {
	return _EthereumBank.Contract.GetToken2addressV2(&_EthereumBank.CallOpts, _name)
}

// HasBridgeTokenCreated is a free data retrieval call binding the contract method 0xefea27ab.
//
// Solidity: function hasBridgeTokenCreated(string _name) view returns(bool)
func (_EthereumBank *EthereumBankCaller) HasBridgeTokenCreated(opts *bind.CallOpts, _name string) (bool, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "hasBridgeTokenCreated", _name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBridgeTokenCreated is a free data retrieval call binding the contract method 0xefea27ab.
//
// Solidity: function hasBridgeTokenCreated(string _name) view returns(bool)
func (_EthereumBank *EthereumBankSession) HasBridgeTokenCreated(_name string) (bool, error) {
	return _EthereumBank.Contract.HasBridgeTokenCreated(&_EthereumBank.CallOpts, _name)
}

// HasBridgeTokenCreated is a free data retrieval call binding the contract method 0xefea27ab.
//
// Solidity: function hasBridgeTokenCreated(string _name) view returns(bool)
func (_EthereumBank *EthereumBankCallerSession) HasBridgeTokenCreated(_name string) (bool, error) {
	return _EthereumBank.Contract.HasBridgeTokenCreated(&_EthereumBank.CallOpts, _name)
}

// Token2WithdrawCfg is a free data retrieval call binding the contract method 0xc4c2709e.
//
// Solidity: function token2WithdrawCfg(bytes32 ) view returns(bool)
func (_EthereumBank *EthereumBankCaller) Token2WithdrawCfg(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "token2WithdrawCfg", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Token2WithdrawCfg is a free data retrieval call binding the contract method 0xc4c2709e.
//
// Solidity: function token2WithdrawCfg(bytes32 ) view returns(bool)
func (_EthereumBank *EthereumBankSession) Token2WithdrawCfg(arg0 [32]byte) (bool, error) {
	return _EthereumBank.Contract.Token2WithdrawCfg(&_EthereumBank.CallOpts, arg0)
}

// Token2WithdrawCfg is a free data retrieval call binding the contract method 0xc4c2709e.
//
// Solidity: function token2WithdrawCfg(bytes32 ) view returns(bool)
func (_EthereumBank *EthereumBankCallerSession) Token2WithdrawCfg(arg0 [32]byte) (bool, error) {
	return _EthereumBank.Contract.Token2WithdrawCfg(&_EthereumBank.CallOpts, arg0)
}

// Token2address is a free data retrieval call binding the contract method 0x85f0bdc7.
//
// Solidity: function token2address(bytes32 ) view returns(address)
func (_EthereumBank *EthereumBankCaller) Token2address(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EthereumBank.contract.Call(opts, &out, "token2address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token2address is a free data retrieval call binding the contract method 0x85f0bdc7.
//
// Solidity: function token2address(bytes32 ) view returns(address)
func (_EthereumBank *EthereumBankSession) Token2address(arg0 [32]byte) (common.Address, error) {
	return _EthereumBank.Contract.Token2address(&_EthereumBank.CallOpts, arg0)
}

// Token2address is a free data retrieval call binding the contract method 0x85f0bdc7.
//
// Solidity: function token2address(bytes32 ) view returns(address)
func (_EthereumBank *EthereumBankCallerSession) Token2address(arg0 [32]byte) (common.Address, error) {
	return _EthereumBank.Contract.Token2address(&_EthereumBank.CallOpts, arg0)
}

// EthereumBankLogBridgeTokenMintIterator is returned from FilterLogBridgeTokenMint and is used to iterate over the raw logs and unpacked data for LogBridgeTokenMint events raised by the EthereumBank contract.
type EthereumBankLogBridgeTokenMintIterator struct {
	Event *EthereumBankLogBridgeTokenMint // Event containing the contract specifics and raw log

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
func (it *EthereumBankLogBridgeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBankLogBridgeTokenMint)
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
		it.Event = new(EthereumBankLogBridgeTokenMint)
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
func (it *EthereumBankLogBridgeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBankLogBridgeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBankLogBridgeTokenMint represents a LogBridgeTokenMint event raised by the EthereumBank contract.
type EthereumBankLogBridgeTokenMint struct {
	Token       common.Address
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogBridgeTokenMint is a free log retrieval operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_EthereumBank *EthereumBankFilterer) FilterLogBridgeTokenMint(opts *bind.FilterOpts) (*EthereumBankLogBridgeTokenMintIterator, error) {

	logs, sub, err := _EthereumBank.contract.FilterLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return &EthereumBankLogBridgeTokenMintIterator{contract: _EthereumBank.contract, event: "LogBridgeTokenMint", logs: logs, sub: sub}, nil
}

// WatchLogBridgeTokenMint is a free log subscription operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_EthereumBank *EthereumBankFilterer) WatchLogBridgeTokenMint(opts *bind.WatchOpts, sink chan<- *EthereumBankLogBridgeTokenMint) (event.Subscription, error) {

	logs, sub, err := _EthereumBank.contract.WatchLogs(opts, "LogBridgeTokenMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBankLogBridgeTokenMint)
				if err := _EthereumBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
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

// ParseLogBridgeTokenMint is a log parse operation binding the contract event 0x4ce045d816af1004ed2e4feb414f2bdaf7e3c92341dd98275e665e861d19551c.
//
// Solidity: event LogBridgeTokenMint(address _token, uint256 _amount, address _beneficiary)
func (_EthereumBank *EthereumBankFilterer) ParseLogBridgeTokenMint(log types.Log) (*EthereumBankLogBridgeTokenMint, error) {
	event := new(EthereumBankLogBridgeTokenMint)
	if err := _EthereumBank.contract.UnpackLog(event, "LogBridgeTokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBankLogBtcTokenBurnIterator is returned from FilterLogBtcTokenBurn and is used to iterate over the raw logs and unpacked data for LogBtcTokenBurn events raised by the EthereumBank contract.
type EthereumBankLogBtcTokenBurnIterator struct {
	Event *EthereumBankLogBtcTokenBurn // Event containing the contract specifics and raw log

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
func (it *EthereumBankLogBtcTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBankLogBtcTokenBurn)
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
		it.Event = new(EthereumBankLogBtcTokenBurn)
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
func (it *EthereumBankLogBtcTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBankLogBtcTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBankLogBtcTokenBurn represents a LogBtcTokenBurn event raised by the EthereumBank contract.
type EthereumBankLogBtcTokenBurn struct {
	ChainID   *big.Int
	Token     common.Address
	Name      string
	Amount    *big.Int
	OwnerFrom common.Address
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogBtcTokenBurn is a free log retrieval operation binding the contract event 0x6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e.
//
// Solidity: event LogBtcTokenBurn(uint256 _chainID, address _token, string _name, uint256 _amount, address _ownerFrom, address _receiver)
func (_EthereumBank *EthereumBankFilterer) FilterLogBtcTokenBurn(opts *bind.FilterOpts) (*EthereumBankLogBtcTokenBurnIterator, error) {

	logs, sub, err := _EthereumBank.contract.FilterLogs(opts, "LogBtcTokenBurn")
	if err != nil {
		return nil, err
	}
	return &EthereumBankLogBtcTokenBurnIterator{contract: _EthereumBank.contract, event: "LogBtcTokenBurn", logs: logs, sub: sub}, nil
}

// WatchLogBtcTokenBurn is a free log subscription operation binding the contract event 0x6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e.
//
// Solidity: event LogBtcTokenBurn(uint256 _chainID, address _token, string _name, uint256 _amount, address _ownerFrom, address _receiver)
func (_EthereumBank *EthereumBankFilterer) WatchLogBtcTokenBurn(opts *bind.WatchOpts, sink chan<- *EthereumBankLogBtcTokenBurn) (event.Subscription, error) {

	logs, sub, err := _EthereumBank.contract.WatchLogs(opts, "LogBtcTokenBurn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBankLogBtcTokenBurn)
				if err := _EthereumBank.contract.UnpackLog(event, "LogBtcTokenBurn", log); err != nil {
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

// ParseLogBtcTokenBurn is a log parse operation binding the contract event 0x6e2650ab6cd6a0b8638139579035840a510887df676e703ee72c99e220f78c5e.
//
// Solidity: event LogBtcTokenBurn(uint256 _chainID, address _token, string _name, uint256 _amount, address _ownerFrom, address _receiver)
func (_EthereumBank *EthereumBankFilterer) ParseLogBtcTokenBurn(log types.Log) (*EthereumBankLogBtcTokenBurn, error) {
	event := new(EthereumBankLogBtcTokenBurn)
	if err := _EthereumBank.contract.UnpackLog(event, "LogBtcTokenBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumBankLogNewBridgeTokenIterator is returned from FilterLogNewBridgeToken and is used to iterate over the raw logs and unpacked data for LogNewBridgeToken events raised by the EthereumBank contract.
type EthereumBankLogNewBridgeTokenIterator struct {
	Event *EthereumBankLogNewBridgeToken // Event containing the contract specifics and raw log

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
func (it *EthereumBankLogNewBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumBankLogNewBridgeToken)
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
		it.Event = new(EthereumBankLogNewBridgeToken)
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
func (it *EthereumBankLogNewBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumBankLogNewBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumBankLogNewBridgeToken represents a LogNewBridgeToken event raised by the EthereumBank contract.
type EthereumBankLogNewBridgeToken struct {
	Token    common.Address
	Name     string
	Symbol   string
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNewBridgeToken is a free log retrieval operation binding the contract event 0x438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d7731.
//
// Solidity: event LogNewBridgeToken(address _token, string _name, string _symbol, uint8 decimals)
func (_EthereumBank *EthereumBankFilterer) FilterLogNewBridgeToken(opts *bind.FilterOpts) (*EthereumBankLogNewBridgeTokenIterator, error) {

	logs, sub, err := _EthereumBank.contract.FilterLogs(opts, "LogNewBridgeToken")
	if err != nil {
		return nil, err
	}
	return &EthereumBankLogNewBridgeTokenIterator{contract: _EthereumBank.contract, event: "LogNewBridgeToken", logs: logs, sub: sub}, nil
}

// WatchLogNewBridgeToken is a free log subscription operation binding the contract event 0x438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d7731.
//
// Solidity: event LogNewBridgeToken(address _token, string _name, string _symbol, uint8 decimals)
func (_EthereumBank *EthereumBankFilterer) WatchLogNewBridgeToken(opts *bind.WatchOpts, sink chan<- *EthereumBankLogNewBridgeToken) (event.Subscription, error) {

	logs, sub, err := _EthereumBank.contract.WatchLogs(opts, "LogNewBridgeToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumBankLogNewBridgeToken)
				if err := _EthereumBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
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

// ParseLogNewBridgeToken is a log parse operation binding the contract event 0x438a0232a8154e2b334120305ea77327d9e18a5a3d12cf4ba9e84891fe8d7731.
//
// Solidity: event LogNewBridgeToken(address _token, string _name, string _symbol, uint8 decimals)
func (_EthereumBank *EthereumBankFilterer) ParseLogNewBridgeToken(log types.Log) (*EthereumBankLogNewBridgeToken, error) {
	event := new(EthereumBankLogNewBridgeToken)
	if err := _EthereumBank.contract.UnpackLog(event, "LogNewBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlMetaData contains all meta data concerning the IAccessControl contract.
var IAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlMetaData.ABI instead.
var IAccessControlABI = IAccessControlMetaData.ABI

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// IAccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IAccessControl contract.
type IAccessControlRoleAdminChangedIterator struct {
	Event *IAccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleAdminChanged)
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
		it.Event = new(IAccessControlRoleAdminChanged)
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
func (it *IAccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the IAccessControl contract.
type IAccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IAccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleAdminChangedIterator{contract: _IAccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleAdminChanged)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*IAccessControlRoleAdminChanged, error) {
	event := new(IAccessControlRoleAdminChanged)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControl contract.
type IAccessControlRoleGrantedIterator struct {
	Event *IAccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleGranted)
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
		it.Event = new(IAccessControlRoleGranted)
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
func (it *IAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleGranted represents a RoleGranted event raised by the IAccessControl contract.
type IAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleGrantedIterator{contract: _IAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleGranted)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleGranted(log types.Log) (*IAccessControlRoleGranted, error) {
	event := new(IAccessControlRoleGranted)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControl contract.
type IAccessControlRoleRevokedIterator struct {
	Event *IAccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleRevoked)
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
		it.Event = new(IAccessControlRoleRevoked)
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
func (it *IAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleRevoked represents a RoleRevoked event raised by the IAccessControl contract.
type IAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleRevokedIterator{contract: _IAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleRevoked)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlRoleRevoked, error) {
	event := new(IAccessControlRoleRevoked)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155ErrorsMetaData contains all meta data concerning the IERC1155Errors contract.
var IERC1155ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"}]",
}

// IERC1155ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1155ErrorsMetaData.ABI instead.
var IERC1155ErrorsABI = IERC1155ErrorsMetaData.ABI

// IERC1155Errors is an auto generated Go binding around an Ethereum contract.
type IERC1155Errors struct {
	IERC1155ErrorsCaller     // Read-only binding to the contract
	IERC1155ErrorsTransactor // Write-only binding to the contract
	IERC1155ErrorsFilterer   // Log filterer for contract events
}

// IERC1155ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155ErrorsSession struct {
	Contract     *IERC1155Errors   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155ErrorsCallerSession struct {
	Contract *IERC1155ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC1155ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155ErrorsTransactorSession struct {
	Contract     *IERC1155ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC1155ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155ErrorsRaw struct {
	Contract *IERC1155Errors // Generic contract binding to access the raw methods on
}

// IERC1155ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155ErrorsCallerRaw struct {
	Contract *IERC1155ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155ErrorsTransactorRaw struct {
	Contract *IERC1155ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155Errors creates a new instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155Errors(address common.Address, backend bind.ContractBackend) (*IERC1155Errors, error) {
	contract, err := bindIERC1155Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155Errors{IERC1155ErrorsCaller: IERC1155ErrorsCaller{contract: contract}, IERC1155ErrorsTransactor: IERC1155ErrorsTransactor{contract: contract}, IERC1155ErrorsFilterer: IERC1155ErrorsFilterer{contract: contract}}, nil
}

// NewIERC1155ErrorsCaller creates a new read-only instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC1155ErrorsCaller, error) {
	contract, err := bindIERC1155Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155ErrorsCaller{contract: contract}, nil
}

// NewIERC1155ErrorsTransactor creates a new write-only instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155ErrorsTransactor, error) {
	contract, err := bindIERC1155Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155ErrorsTransactor{contract: contract}, nil
}

// NewIERC1155ErrorsFilterer creates a new log filterer instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155ErrorsFilterer, error) {
	contract, err := bindIERC1155Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155ErrorsFilterer{contract: contract}, nil
}

// bindIERC1155Errors binds a generic wrapper to an already deployed contract.
func bindIERC1155Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1155ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Errors *IERC1155ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Errors.Contract.IERC1155ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Errors *IERC1155ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.IERC1155ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Errors *IERC1155ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.IERC1155ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Errors *IERC1155ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Errors *IERC1155ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Errors *IERC1155ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.contract.Transact(opts, method, params...)
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ErrorsMetaData contains all meta data concerning the IERC20Errors contract.
var IERC20ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"}]",
}

// IERC20ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20ErrorsMetaData.ABI instead.
var IERC20ErrorsABI = IERC20ErrorsMetaData.ABI

// IERC20Errors is an auto generated Go binding around an Ethereum contract.
type IERC20Errors struct {
	IERC20ErrorsCaller     // Read-only binding to the contract
	IERC20ErrorsTransactor // Write-only binding to the contract
	IERC20ErrorsFilterer   // Log filterer for contract events
}

// IERC20ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20ErrorsSession struct {
	Contract     *IERC20Errors     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20ErrorsCallerSession struct {
	Contract *IERC20ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20ErrorsTransactorSession struct {
	Contract     *IERC20ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20ErrorsRaw struct {
	Contract *IERC20Errors // Generic contract binding to access the raw methods on
}

// IERC20ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20ErrorsCallerRaw struct {
	Contract *IERC20ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20ErrorsTransactorRaw struct {
	Contract *IERC20ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Errors creates a new instance of IERC20Errors, bound to a specific deployed contract.
func NewIERC20Errors(address common.Address, backend bind.ContractBackend) (*IERC20Errors, error) {
	contract, err := bindIERC20Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Errors{IERC20ErrorsCaller: IERC20ErrorsCaller{contract: contract}, IERC20ErrorsTransactor: IERC20ErrorsTransactor{contract: contract}, IERC20ErrorsFilterer: IERC20ErrorsFilterer{contract: contract}}, nil
}

// NewIERC20ErrorsCaller creates a new read-only instance of IERC20Errors, bound to a specific deployed contract.
func NewIERC20ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC20ErrorsCaller, error) {
	contract, err := bindIERC20Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20ErrorsCaller{contract: contract}, nil
}

// NewIERC20ErrorsTransactor creates a new write-only instance of IERC20Errors, bound to a specific deployed contract.
func NewIERC20ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20ErrorsTransactor, error) {
	contract, err := bindIERC20Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20ErrorsTransactor{contract: contract}, nil
}

// NewIERC20ErrorsFilterer creates a new log filterer instance of IERC20Errors, bound to a specific deployed contract.
func NewIERC20ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20ErrorsFilterer, error) {
	contract, err := bindIERC20Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20ErrorsFilterer{contract: contract}, nil
}

// bindIERC20Errors binds a generic wrapper to an already deployed contract.
func bindIERC20Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Errors *IERC20ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Errors.Contract.IERC20ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Errors *IERC20ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Errors.Contract.IERC20ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Errors *IERC20ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Errors.Contract.IERC20ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Errors *IERC20ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Errors *IERC20ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Errors *IERC20ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Errors.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataMetaData.ABI instead.
var IERC20MetadataABI = IERC20MetadataMetaData.ABI

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, from, to, value)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ErrorsMetaData contains all meta data concerning the IERC721Errors contract.
var IERC721ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"}]",
}

// IERC721ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721ErrorsMetaData.ABI instead.
var IERC721ErrorsABI = IERC721ErrorsMetaData.ABI

// IERC721Errors is an auto generated Go binding around an Ethereum contract.
type IERC721Errors struct {
	IERC721ErrorsCaller     // Read-only binding to the contract
	IERC721ErrorsTransactor // Write-only binding to the contract
	IERC721ErrorsFilterer   // Log filterer for contract events
}

// IERC721ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ErrorsSession struct {
	Contract     *IERC721Errors    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ErrorsCallerSession struct {
	Contract *IERC721ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC721ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ErrorsTransactorSession struct {
	Contract     *IERC721ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC721ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ErrorsRaw struct {
	Contract *IERC721Errors // Generic contract binding to access the raw methods on
}

// IERC721ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ErrorsCallerRaw struct {
	Contract *IERC721ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ErrorsTransactorRaw struct {
	Contract *IERC721ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Errors creates a new instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721Errors(address common.Address, backend bind.ContractBackend) (*IERC721Errors, error) {
	contract, err := bindIERC721Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Errors{IERC721ErrorsCaller: IERC721ErrorsCaller{contract: contract}, IERC721ErrorsTransactor: IERC721ErrorsTransactor{contract: contract}, IERC721ErrorsFilterer: IERC721ErrorsFilterer{contract: contract}}, nil
}

// NewIERC721ErrorsCaller creates a new read-only instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC721ErrorsCaller, error) {
	contract, err := bindIERC721Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ErrorsCaller{contract: contract}, nil
}

// NewIERC721ErrorsTransactor creates a new write-only instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ErrorsTransactor, error) {
	contract, err := bindIERC721Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ErrorsTransactor{contract: contract}, nil
}

// NewIERC721ErrorsFilterer creates a new log filterer instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ErrorsFilterer, error) {
	contract, err := bindIERC721Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ErrorsFilterer{contract: contract}, nil
}

// bindIERC721Errors binds a generic wrapper to an already deployed contract.
func bindIERC721Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC721ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Errors *IERC721ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Errors.Contract.IERC721ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Errors *IERC721ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Errors.Contract.IERC721ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Errors *IERC721ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Errors.Contract.IERC721ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Errors *IERC721ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Errors *IERC721ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Errors *IERC721ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Errors.Contract.contract.Transact(opts, method, params...)
}

// JuBankMetaData contains all meta data concerning the JuBank contract.
var JuBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"LogLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"LogUnlock\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getLockedTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getofflineSaveCfg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offlineSave\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offlineSaveCfgs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_percents\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenAddrAllow2name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenAllow2Lock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526006805461ffff191661500517905534801561001e575f80fd5b505f805561075b8061002f5f395ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c8063a80de1a41161006e578063a80de1a414610169578063b5a9096e1461017c578063b86247d714610192578063d1d008ae146101b1578063d2168d20146101be578063fad60627146101e8575f80fd5b80630a1f9b66146100aa5780634e25d152146100da5780636f0fccab146100fd5780637056c4931461011d5780639eadc7cc14610145575b5f80fd5b6100bd6100b8366004610547565b6101fb565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ed6100e83660046105f2565b610242565b6040516100d1949392919061066c565b61011061010b3660046105f2565b6102fe565b6040516100d191906106a9565b6100bd61012b3660046106bb565b60036020525f90815260409020546001600160a01b031681565b60065461015790610100900460ff1681565b60405160ff90911681526020016100d1565b6101106101773660046105f2565b6103a9565b6101845f5481565b6040519081526020016100d1565b6101846101a03660046105f2565b60026020525f908152604090205481565b6006546101579060ff1681565b6101d16101cc3660046105f2565b610440565b6040805192835260ff9091166020830152016100d1565b6001546100bd906001600160a01b031681565b5f808260405160200161020e91906106d2565b60408051601f1981840301815291815281516020928301205f90815260039092529020546001600160a01b03169392505050565b60056020525f9081526040902080546001820180546001600160a01b03909216929161026d906106ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610299906106ed565b80156102e45780601f106102bb576101008083540402835291602001916102e4565b820191905f5260205f20905b8154815290600101906020018083116102c757829003601f168201915b50505050600283015460039093015491929160ff16905084565b6001600160a01b0381165f90815260046020526040812080546060929190610325906106ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610351906106ed565b801561039c5780601f106103735761010080835404028352916020019161039c565b820191905f5260205f20905b81548152906001019060200180831161037f57829003601f168201915b5093979650505050505050565b60046020525f9081526040902080546103c1906106ed565b80601f01602080910402602001604051908101604052809291908181526020018280546103ed906106ed565b80156104385780601f1061040f57610100808354040283529160200191610438565b820191905f5260205f20905b81548152906001019060200180831161041b57829003601f168201915b505050505081565b6001600160a01b038082165f9081526005602090815260408083208151608081019092528054909416815260018401805493948594859491929184019190610487906106ed565b80601f01602080910402602001604051908101604052809291908181526020018280546104b3906106ed565b80156104fe5780601f106104d5576101008083540402835291602001916104fe565b820191905f5260205f20905b8154815290600101906020018083116104e157829003601f168201915b50505091835250506002820154602082015260039091015460ff16604091820152810151606090910151909590945092505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610557575f80fd5b813567ffffffffffffffff8082111561056e575f80fd5b818401915084601f830112610581575f80fd5b81358181111561059357610593610533565b604051601f8201601f19908116603f011681019083821181831017156105bb576105bb610533565b816040528281528760208487010111156105d3575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f60208284031215610602575f80fd5b81356001600160a01b0381168114610618575f80fd5b9392505050565b5f5b83811015610639578181015183820152602001610621565b50505f910152565b5f815180845261065881602086016020860161061f565b601f01601f19169290920160200192915050565b6001600160a01b03851681526080602082018190525f9061068f90830186610641565b905083604083015260ff8316606083015295945050505050565b602081525f6106186020830184610641565b5f602082840312156106cb575f80fd5b5035919050565b5f82516106e381846020870161061f565b9190910192915050565b600181811c9082168061070157607f821691505b60208210810361071f57634e487b7160e01b5f52602260045260245ffd5b5091905056fea264697066735822122014245905a019c49c9b3a78e89c2f534a32adc715ec7a90e825bfb5fb10e6b0ad64736f6c63430008140033",
}

// JuBankABI is the input ABI used to generate the binding from.
// Deprecated: Use JuBankMetaData.ABI instead.
var JuBankABI = JuBankMetaData.ABI

// JuBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JuBankMetaData.Bin instead.
var JuBankBin = JuBankMetaData.Bin

// DeployJuBank deploys a new Ethereum contract, binding an instance of JuBank to it.
func DeployJuBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JuBank, error) {
	parsed, err := JuBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JuBankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JuBank{JuBankCaller: JuBankCaller{contract: contract}, JuBankTransactor: JuBankTransactor{contract: contract}, JuBankFilterer: JuBankFilterer{contract: contract}}, nil
}

// JuBank is an auto generated Go binding around an Ethereum contract.
type JuBank struct {
	JuBankCaller     // Read-only binding to the contract
	JuBankTransactor // Write-only binding to the contract
	JuBankFilterer   // Log filterer for contract events
}

// JuBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type JuBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JuBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JuBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JuBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JuBankSession struct {
	Contract     *JuBank           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JuBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JuBankCallerSession struct {
	Contract *JuBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JuBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JuBankTransactorSession struct {
	Contract     *JuBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JuBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type JuBankRaw struct {
	Contract *JuBank // Generic contract binding to access the raw methods on
}

// JuBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JuBankCallerRaw struct {
	Contract *JuBankCaller // Generic read-only contract binding to access the raw methods on
}

// JuBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JuBankTransactorRaw struct {
	Contract *JuBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJuBank creates a new instance of JuBank, bound to a specific deployed contract.
func NewJuBank(address common.Address, backend bind.ContractBackend) (*JuBank, error) {
	contract, err := bindJuBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JuBank{JuBankCaller: JuBankCaller{contract: contract}, JuBankTransactor: JuBankTransactor{contract: contract}, JuBankFilterer: JuBankFilterer{contract: contract}}, nil
}

// NewJuBankCaller creates a new read-only instance of JuBank, bound to a specific deployed contract.
func NewJuBankCaller(address common.Address, caller bind.ContractCaller) (*JuBankCaller, error) {
	contract, err := bindJuBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JuBankCaller{contract: contract}, nil
}

// NewJuBankTransactor creates a new write-only instance of JuBank, bound to a specific deployed contract.
func NewJuBankTransactor(address common.Address, transactor bind.ContractTransactor) (*JuBankTransactor, error) {
	contract, err := bindJuBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JuBankTransactor{contract: contract}, nil
}

// NewJuBankFilterer creates a new log filterer instance of JuBank, bound to a specific deployed contract.
func NewJuBankFilterer(address common.Address, filterer bind.ContractFilterer) (*JuBankFilterer, error) {
	contract, err := bindJuBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JuBankFilterer{contract: contract}, nil
}

// bindJuBank binds a generic wrapper to an already deployed contract.
func bindJuBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JuBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuBank *JuBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuBank.Contract.JuBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuBank *JuBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuBank.Contract.JuBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuBank *JuBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuBank.Contract.JuBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JuBank *JuBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JuBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JuBank *JuBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JuBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JuBank *JuBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JuBank.Contract.contract.Transact(opts, method, params...)
}

// GetLockedTokenAddress is a free data retrieval call binding the contract method 0x0a1f9b66.
//
// Solidity: function getLockedTokenAddress(string _name) view returns(address)
func (_JuBank *JuBankCaller) GetLockedTokenAddress(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "getLockedTokenAddress", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLockedTokenAddress is a free data retrieval call binding the contract method 0x0a1f9b66.
//
// Solidity: function getLockedTokenAddress(string _name) view returns(address)
func (_JuBank *JuBankSession) GetLockedTokenAddress(_name string) (common.Address, error) {
	return _JuBank.Contract.GetLockedTokenAddress(&_JuBank.CallOpts, _name)
}

// GetLockedTokenAddress is a free data retrieval call binding the contract method 0x0a1f9b66.
//
// Solidity: function getLockedTokenAddress(string _name) view returns(address)
func (_JuBank *JuBankCallerSession) GetLockedTokenAddress(_name string) (common.Address, error) {
	return _JuBank.Contract.GetLockedTokenAddress(&_JuBank.CallOpts, _name)
}

// GetTokenName is a free data retrieval call binding the contract method 0x6f0fccab.
//
// Solidity: function getTokenName(address _tokenAddress) view returns(string)
func (_JuBank *JuBankCaller) GetTokenName(opts *bind.CallOpts, _tokenAddress common.Address) (string, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "getTokenName", _tokenAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenName is a free data retrieval call binding the contract method 0x6f0fccab.
//
// Solidity: function getTokenName(address _tokenAddress) view returns(string)
func (_JuBank *JuBankSession) GetTokenName(_tokenAddress common.Address) (string, error) {
	return _JuBank.Contract.GetTokenName(&_JuBank.CallOpts, _tokenAddress)
}

// GetTokenName is a free data retrieval call binding the contract method 0x6f0fccab.
//
// Solidity: function getTokenName(address _tokenAddress) view returns(string)
func (_JuBank *JuBankCallerSession) GetTokenName(_tokenAddress common.Address) (string, error) {
	return _JuBank.Contract.GetTokenName(&_JuBank.CallOpts, _tokenAddress)
}

// GetofflineSaveCfg is a free data retrieval call binding the contract method 0xd2168d20.
//
// Solidity: function getofflineSaveCfg(address _token) view returns(uint256, uint8)
func (_JuBank *JuBankCaller) GetofflineSaveCfg(opts *bind.CallOpts, _token common.Address) (*big.Int, uint8, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "getofflineSaveCfg", _token)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetofflineSaveCfg is a free data retrieval call binding the contract method 0xd2168d20.
//
// Solidity: function getofflineSaveCfg(address _token) view returns(uint256, uint8)
func (_JuBank *JuBankSession) GetofflineSaveCfg(_token common.Address) (*big.Int, uint8, error) {
	return _JuBank.Contract.GetofflineSaveCfg(&_JuBank.CallOpts, _token)
}

// GetofflineSaveCfg is a free data retrieval call binding the contract method 0xd2168d20.
//
// Solidity: function getofflineSaveCfg(address _token) view returns(uint256, uint8)
func (_JuBank *JuBankCallerSession) GetofflineSaveCfg(_token common.Address) (*big.Int, uint8, error) {
	return _JuBank.Contract.GetofflineSaveCfg(&_JuBank.CallOpts, _token)
}

// HighThreshold is a free data retrieval call binding the contract method 0x9eadc7cc.
//
// Solidity: function highThreshold() view returns(uint8)
func (_JuBank *JuBankCaller) HighThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "highThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HighThreshold is a free data retrieval call binding the contract method 0x9eadc7cc.
//
// Solidity: function highThreshold() view returns(uint8)
func (_JuBank *JuBankSession) HighThreshold() (uint8, error) {
	return _JuBank.Contract.HighThreshold(&_JuBank.CallOpts)
}

// HighThreshold is a free data retrieval call binding the contract method 0x9eadc7cc.
//
// Solidity: function highThreshold() view returns(uint8)
func (_JuBank *JuBankCallerSession) HighThreshold() (uint8, error) {
	return _JuBank.Contract.HighThreshold(&_JuBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_JuBank *JuBankCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "lockNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_JuBank *JuBankSession) LockNonce() (*big.Int, error) {
	return _JuBank.Contract.LockNonce(&_JuBank.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_JuBank *JuBankCallerSession) LockNonce() (*big.Int, error) {
	return _JuBank.Contract.LockNonce(&_JuBank.CallOpts)
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) view returns(uint256)
func (_JuBank *JuBankCaller) LockedFunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "lockedFunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) view returns(uint256)
func (_JuBank *JuBankSession) LockedFunds(arg0 common.Address) (*big.Int, error) {
	return _JuBank.Contract.LockedFunds(&_JuBank.CallOpts, arg0)
}

// LockedFunds is a free data retrieval call binding the contract method 0xb86247d7.
//
// Solidity: function lockedFunds(address ) view returns(uint256)
func (_JuBank *JuBankCallerSession) LockedFunds(arg0 common.Address) (*big.Int, error) {
	return _JuBank.Contract.LockedFunds(&_JuBank.CallOpts, arg0)
}

// LowThreshold is a free data retrieval call binding the contract method 0xd1d008ae.
//
// Solidity: function lowThreshold() view returns(uint8)
func (_JuBank *JuBankCaller) LowThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "lowThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LowThreshold is a free data retrieval call binding the contract method 0xd1d008ae.
//
// Solidity: function lowThreshold() view returns(uint8)
func (_JuBank *JuBankSession) LowThreshold() (uint8, error) {
	return _JuBank.Contract.LowThreshold(&_JuBank.CallOpts)
}

// LowThreshold is a free data retrieval call binding the contract method 0xd1d008ae.
//
// Solidity: function lowThreshold() view returns(uint8)
func (_JuBank *JuBankCallerSession) LowThreshold() (uint8, error) {
	return _JuBank.Contract.LowThreshold(&_JuBank.CallOpts)
}

// OfflineSave is a free data retrieval call binding the contract method 0xfad60627.
//
// Solidity: function offlineSave() view returns(address)
func (_JuBank *JuBankCaller) OfflineSave(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "offlineSave")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OfflineSave is a free data retrieval call binding the contract method 0xfad60627.
//
// Solidity: function offlineSave() view returns(address)
func (_JuBank *JuBankSession) OfflineSave() (common.Address, error) {
	return _JuBank.Contract.OfflineSave(&_JuBank.CallOpts)
}

// OfflineSave is a free data retrieval call binding the contract method 0xfad60627.
//
// Solidity: function offlineSave() view returns(address)
func (_JuBank *JuBankCallerSession) OfflineSave() (common.Address, error) {
	return _JuBank.Contract.OfflineSave(&_JuBank.CallOpts)
}

// OfflineSaveCfgs is a free data retrieval call binding the contract method 0x4e25d152.
//
// Solidity: function offlineSaveCfgs(address ) view returns(address token, string name, uint256 _threshold, uint8 _percents)
func (_JuBank *JuBankCaller) OfflineSaveCfgs(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Threshold *big.Int
	Percents  uint8
}, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "offlineSaveCfgs", arg0)

	outstruct := new(struct {
		Token     common.Address
		Name      string
		Threshold *big.Int
		Percents  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Threshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Percents = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// OfflineSaveCfgs is a free data retrieval call binding the contract method 0x4e25d152.
//
// Solidity: function offlineSaveCfgs(address ) view returns(address token, string name, uint256 _threshold, uint8 _percents)
func (_JuBank *JuBankSession) OfflineSaveCfgs(arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Threshold *big.Int
	Percents  uint8
}, error) {
	return _JuBank.Contract.OfflineSaveCfgs(&_JuBank.CallOpts, arg0)
}

// OfflineSaveCfgs is a free data retrieval call binding the contract method 0x4e25d152.
//
// Solidity: function offlineSaveCfgs(address ) view returns(address token, string name, uint256 _threshold, uint8 _percents)
func (_JuBank *JuBankCallerSession) OfflineSaveCfgs(arg0 common.Address) (struct {
	Token     common.Address
	Name      string
	Threshold *big.Int
	Percents  uint8
}, error) {
	return _JuBank.Contract.OfflineSaveCfgs(&_JuBank.CallOpts, arg0)
}

// TokenAddrAllow2name is a free data retrieval call binding the contract method 0xa80de1a4.
//
// Solidity: function tokenAddrAllow2name(address ) view returns(string)
func (_JuBank *JuBankCaller) TokenAddrAllow2name(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "tokenAddrAllow2name", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenAddrAllow2name is a free data retrieval call binding the contract method 0xa80de1a4.
//
// Solidity: function tokenAddrAllow2name(address ) view returns(string)
func (_JuBank *JuBankSession) TokenAddrAllow2name(arg0 common.Address) (string, error) {
	return _JuBank.Contract.TokenAddrAllow2name(&_JuBank.CallOpts, arg0)
}

// TokenAddrAllow2name is a free data retrieval call binding the contract method 0xa80de1a4.
//
// Solidity: function tokenAddrAllow2name(address ) view returns(string)
func (_JuBank *JuBankCallerSession) TokenAddrAllow2name(arg0 common.Address) (string, error) {
	return _JuBank.Contract.TokenAddrAllow2name(&_JuBank.CallOpts, arg0)
}

// TokenAllow2Lock is a free data retrieval call binding the contract method 0x7056c493.
//
// Solidity: function tokenAllow2Lock(bytes32 ) view returns(address)
func (_JuBank *JuBankCaller) TokenAllow2Lock(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _JuBank.contract.Call(opts, &out, "tokenAllow2Lock", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAllow2Lock is a free data retrieval call binding the contract method 0x7056c493.
//
// Solidity: function tokenAllow2Lock(bytes32 ) view returns(address)
func (_JuBank *JuBankSession) TokenAllow2Lock(arg0 [32]byte) (common.Address, error) {
	return _JuBank.Contract.TokenAllow2Lock(&_JuBank.CallOpts, arg0)
}

// TokenAllow2Lock is a free data retrieval call binding the contract method 0x7056c493.
//
// Solidity: function tokenAllow2Lock(bytes32 ) view returns(address)
func (_JuBank *JuBankCallerSession) TokenAllow2Lock(arg0 [32]byte) (common.Address, error) {
	return _JuBank.Contract.TokenAllow2Lock(&_JuBank.CallOpts, arg0)
}

// JuBankLogLockIterator is returned from FilterLogLock and is used to iterate over the raw logs and unpacked data for LogLock events raised by the JuBank contract.
type JuBankLogLockIterator struct {
	Event *JuBankLogLock // Event containing the contract specifics and raw log

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
func (it *JuBankLogLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuBankLogLock)
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
		it.Event = new(JuBankLogLock)
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
func (it *JuBankLogLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuBankLogLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuBankLogLock represents a LogLock event raised by the JuBank contract.
type JuBankLogLock struct {
	From  common.Address
	To    common.Address
	Token common.Address
	Name  string
	Value *big.Int
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogLock is a free log retrieval operation binding the contract event 0x4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe.
//
// Solidity: event LogLock(address _from, address _to, address _token, string _name, uint256 _value, uint256 _nonce)
func (_JuBank *JuBankFilterer) FilterLogLock(opts *bind.FilterOpts) (*JuBankLogLockIterator, error) {

	logs, sub, err := _JuBank.contract.FilterLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return &JuBankLogLockIterator{contract: _JuBank.contract, event: "LogLock", logs: logs, sub: sub}, nil
}

// WatchLogLock is a free log subscription operation binding the contract event 0x4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe.
//
// Solidity: event LogLock(address _from, address _to, address _token, string _name, uint256 _value, uint256 _nonce)
func (_JuBank *JuBankFilterer) WatchLogLock(opts *bind.WatchOpts, sink chan<- *JuBankLogLock) (event.Subscription, error) {

	logs, sub, err := _JuBank.contract.WatchLogs(opts, "LogLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuBankLogLock)
				if err := _JuBank.contract.UnpackLog(event, "LogLock", log); err != nil {
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

// ParseLogLock is a log parse operation binding the contract event 0x4150c77f33761fbde38a274b3cbecce67bc2060ac09687dd6eab0cffe90dfefe.
//
// Solidity: event LogLock(address _from, address _to, address _token, string _name, uint256 _value, uint256 _nonce)
func (_JuBank *JuBankFilterer) ParseLogLock(log types.Log) (*JuBankLogLock, error) {
	event := new(JuBankLogLock)
	if err := _JuBank.contract.UnpackLog(event, "LogLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JuBankLogUnlockIterator is returned from FilterLogUnlock and is used to iterate over the raw logs and unpacked data for LogUnlock events raised by the JuBank contract.
type JuBankLogUnlockIterator struct {
	Event *JuBankLogUnlock // Event containing the contract specifics and raw log

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
func (it *JuBankLogUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JuBankLogUnlock)
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
		it.Event = new(JuBankLogUnlock)
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
func (it *JuBankLogUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JuBankLogUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JuBankLogUnlock represents a LogUnlock event raised by the JuBank contract.
type JuBankLogUnlock struct {
	To    common.Address
	Token common.Address
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogUnlock is a free log retrieval operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _name, uint256 _value)
func (_JuBank *JuBankFilterer) FilterLogUnlock(opts *bind.FilterOpts) (*JuBankLogUnlockIterator, error) {

	logs, sub, err := _JuBank.contract.FilterLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return &JuBankLogUnlockIterator{contract: _JuBank.contract, event: "LogUnlock", logs: logs, sub: sub}, nil
}

// WatchLogUnlock is a free log subscription operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _name, uint256 _value)
func (_JuBank *JuBankFilterer) WatchLogUnlock(opts *bind.WatchOpts, sink chan<- *JuBankLogUnlock) (event.Subscription, error) {

	logs, sub, err := _JuBank.contract.WatchLogs(opts, "LogUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JuBankLogUnlock)
				if err := _JuBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
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

// ParseLogUnlock is a log parse operation binding the contract event 0x802cd873de701272ec903860b690986bd460b5bcd57e30ac1fdfdeece10528ac.
//
// Solidity: event LogUnlock(address _to, address _token, string _name, uint256 _value)
func (_JuBank *JuBankFilterer) ParseLogUnlock(log types.Log) (*JuBankLogUnlock, error) {
	event := new(JuBankLogUnlock)
	if err := _JuBank.contract.UnpackLog(event, "LogUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_valset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btcBridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"LogNewOracleClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"LogProphecyProcessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"btcBridge\",\"outputs\":[{\"internalType\":\"contractBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"}],\"name\":\"checkBridgeProphecy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hasMadeClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumOracle.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"_ethereumReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_srcTxHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_claimID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature2\",\"type\":\"bytes\"}],\"name\":\"newOracleClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracleClaimValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"valset\",\"outputs\":[{\"internalType\":\"contractValset\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610fb1380380610fb183398101604081905261002e9161008a565b600280546001600160a01b039485166001600160a01b0319918216179091555f805492851692821692909217909155600180549290931691161790556100ca565b80516001600160a01b0381168114610085575f80fd5b919050565b5f805f6060848603121561009c575f80fd5b6100a58461006f565b92506100b36020850161006f565b91506100c16040850161006f565b90509250925092565b610eda806100d75f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c806371fb93581161005857806371fb93581461010b5780637f54af0c146101205780638350549f14610133578063e83def4314610146575f80fd5b8063129c8a23146100895780631ffcdacc146100c05780633408e470146100ea578063570ca735146100f8575b5f80fd5b6100ab610097366004610bad565b60046020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b5f546100d2906001600160a01b031681565b6040516001600160a01b0390911681526020016100b7565b6040514681526020016100b7565b6002546100d2906001600160a01b031681565b61011e610119366004610c3e565b610176565b005b6001546100d2906001600160a01b031681565b6100d2610141366004610d1a565b61078f565b610159610154366004610bad565b6107c3565b6040805193151584526020840192909252908201526060016100b7565b600154604051631015428760e21b81523360048201526001600160a01b03909116906340550a1c90602401602060405180830381865afa1580156101bc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e09190610d3a565b6102315760405162461bcd60e51b815260206004820152601b60248201527f4d75737420626520616e206163746976652076616c696461746f72000000000060448201526064015b60405180910390fd5b844689898d8d8b60405160200161024d96959493929190610d60565b604051602081830303815290604052805190602001201461029e5760405162461bcd60e51b815260206004820152600b60248201526a31b630b4b6a4a21032b93960a91b6044820152606401610228565b5f8581526004602052604090205460ff16156102f15760405162461bcd60e51b8152602060048201526012602482015271111d5c1b1a58d85d19590818db185a5b525160721b6044820152606401610228565b600180546040516319045a2560e01b81525f916001600160a01b0316906319045a2590610326908a908a908a90600401610da3565b602060405180830381865afa158015610341573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103659190610dd8565b600154604051631015428760e21b81526001600160a01b0380841660048301529293509116906340550a1c90602401602060405180830381865afa1580156103af573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103d39190610d3a565b1561040e576001600160a01b03811633036104005760405162461bcd60e51b815260040161022890610df3565b61040b600183610e35565b91505b6001546040516319045a2560e01b81525f916001600160a01b0316906319045a2590610442908b9089908990600401610da3565b602060405180830381865afa15801561045d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104819190610dd8565b600154604051631015428760e21b81526001600160a01b0380841660048301529293509116906340550a1c90602401602060405180830381865afa1580156104cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104ef9190610d3a565b15610563576001600160a01b038116330361051c5760405162461bcd60e51b815260040161022890610df3565b8260020361055557816001600160a01b0316816001600160a01b0316036105555760405162461bcd60e51b815260040161022890610df3565b610560600184610e35565b92505b60038310156105ac5760405162461bcd60e51b81526020600482015260156024820152744e6f7420656e6f756768207369676e61747572657360581b6044820152606401610228565b60018e60028111156105c0576105c0610e4e565b14806105dd575060028e60028111156105db576105db610e4e565b145b61061e5760405162461bcd60e51b8152602060048201526012602482015271496e76616c696420636c61696d207479706560701b6044820152606401610228565b60018e600281111561063257610632610e4e565b036106ab575f5460405163107465c560e21b8152600481018a90526001600160a01b038f811660248301528e81166044830152606482018c9052909116906341d19714906084015f604051808303815f87803b158015610690575f80fd5b505af11580156106a2573d5f803e3d5ffd5b5050505061071b565b5f546040516333429e0560e01b8152600481018a90526001600160a01b038f811660248301528e81166044830152606482018c9052909116906333429e05906084015f604051808303815f87803b158015610704575f80fd5b505af1158015610716573d5f803e3d5ffd5b505050505b5f8881526004602052604090819020805460ff19166001179055517f3fe675dd72294e681d888f819675033de7f7d2e84dd7da60daab7af319c7934590610777908a9033909182526001600160a01b0316602082015260400190565b60405180910390a15050505050505050505050505050565b6003602052815f5260405f2081815481106107a8575f80fd5b5f918252602090912001546001600160a01b03169150829050565b6002545f90819081906001600160a01b0316331461081b5760405162461bcd60e51b815260206004820152601560248201527426bab9ba103132903a34329037b832b930ba37b91760591b6044820152606401610228565b5f5460405163553b05e160e11b81526004810186905285916001600160a01b03169063aa760bc290602401602060405180830381865afa158015610861573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108859190610d3a565b15156001146108ee5760405162461bcd60e51b815260206004820152602f60248201527f5468652070726f7068656379206d7573742062652070656e64696e6720666f7260448201526e103a3434b99037b832b930ba34b7b760891b6064820152608401610228565b5f5460405163553b05e160e11b8152600481018790526001600160a01b039091169063aa760bc290602401602060405180830381865afa158015610934573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109589190610d3a565b15156001146109a95760405162461bcd60e51b815260206004820181905260248201527f43616e206f6e6c7920636865636b206163746976652070726f706865636965736044820152606401610228565b6109b2856109c0565b935093509350509193909250565b5f805f805f90505f60015f9054906101000a90046001600160a01b03166001600160a01b031663db3ad22c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a18573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3c9190610e62565b90505f5b5f87815260036020526040902054811015610b7f575f878152600360205260408120805483908110610a7457610a74610e79565b5f91825260209091200154600154604051631015428760e21b81526001600160a01b039283166004820181905293509116906340550a1c90602401602060405180830381865afa158015610aca573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aee9190610d3a565b15610b6c576001546040516311cda46960e21b81526001600160a01b0383811660048301529091169063473691a490602401602060405180830381865afa158015610b3b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5f9190610e62565b610b699085610e35565b93505b50610b78816001610e35565b9050610a40565b505f610b8c836003610e8d565b90505f610b9a836002610e8d565b8083101599929850965090945050505050565b5f60208284031215610bbd575f80fd5b5035919050565b803560038110610bd2575f80fd5b919050565b6001600160a01b0381168114610beb575f80fd5b50565b8035610bd281610bd7565b5f8083601f840112610c09575f80fd5b50813567ffffffffffffffff811115610c20575f80fd5b602083019150836020828501011115610c37575f80fd5b9250929050565b5f805f805f805f805f805f6101008c8e031215610c59575f80fd5b610c628c610bc4565b9a50610c7060208d01610bee565b9950610c7e60408d01610bee565b985067ffffffffffffffff8060608e01351115610c99575f80fd5b610ca98e60608f01358f01610bf9565b909950975060808d0135965060a08d0135955060c08d0135811015610ccc575f80fd5b610cdc8e60c08f01358f01610bf9565b909550935060e08d0135811015610cf1575f80fd5b50610d028d60e08e01358e01610bf9565b81935080925050509295989b509295989b9093969950565b5f8060408385031215610d2b575f80fd5b50508035926020909101359150565b5f60208284031215610d4a575f80fd5b81518015158114610d59575f80fd5b9392505050565b868152848660208301376bffffffffffffffffffffffff19606094851b81169590910160208101959095529190921b166034830152604882015260680192915050565b83815260406020820152816040820152818360608301375f818301606090810191909152601f909201601f1916010192915050565b5f60208284031215610de8575f80fd5b8151610d5981610bd7565b602080825260149082015273223ab83634b1b0ba32b2103b30b634b230ba37b960611b604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610e4857610e48610e21565b92915050565b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215610e72575f80fd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b8082028115828204841417610e4857610e48610e2156fea2646970667358221220629c22954740fcc707945afa0eba93d56873262f177a28017ee9494bf83c3c1064736f6c63430008140033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _valset common.Address, _btcBridge common.Address) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend, _operator, _valset, _btcBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// BtcBridge is a free data retrieval call binding the contract method 0x1ffcdacc.
//
// Solidity: function btcBridge() view returns(address)
func (_Oracle *OracleCaller) BtcBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "btcBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcBridge is a free data retrieval call binding the contract method 0x1ffcdacc.
//
// Solidity: function btcBridge() view returns(address)
func (_Oracle *OracleSession) BtcBridge() (common.Address, error) {
	return _Oracle.Contract.BtcBridge(&_Oracle.CallOpts)
}

// BtcBridge is a free data retrieval call binding the contract method 0x1ffcdacc.
//
// Solidity: function btcBridge() view returns(address)
func (_Oracle *OracleCallerSession) BtcBridge() (common.Address, error) {
	return _Oracle.Contract.BtcBridge(&_Oracle.CallOpts)
}

// CheckBridgeProphecy is a free data retrieval call binding the contract method 0xe83def43.
//
// Solidity: function checkBridgeProphecy(bytes32 _claimID) view returns(bool, uint256, uint256)
func (_Oracle *OracleCaller) CheckBridgeProphecy(opts *bind.CallOpts, _claimID [32]byte) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "checkBridgeProphecy", _claimID)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CheckBridgeProphecy is a free data retrieval call binding the contract method 0xe83def43.
//
// Solidity: function checkBridgeProphecy(bytes32 _claimID) view returns(bool, uint256, uint256)
func (_Oracle *OracleSession) CheckBridgeProphecy(_claimID [32]byte) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.CheckBridgeProphecy(&_Oracle.CallOpts, _claimID)
}

// CheckBridgeProphecy is a free data retrieval call binding the contract method 0xe83def43.
//
// Solidity: function checkBridgeProphecy(bytes32 _claimID) view returns(bool, uint256, uint256)
func (_Oracle *OracleCallerSession) CheckBridgeProphecy(_claimID [32]byte) (bool, *big.Int, *big.Int, error) {
	return _Oracle.Contract.CheckBridgeProphecy(&_Oracle.CallOpts, _claimID)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Oracle *OracleCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Oracle *OracleSession) GetChainId() (*big.Int, error) {
	return _Oracle.Contract.GetChainId(&_Oracle.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Oracle *OracleCallerSession) GetChainId() (*big.Int, error) {
	return _Oracle.Contract.GetChainId(&_Oracle.CallOpts)
}

// HasMadeClaim is a free data retrieval call binding the contract method 0x129c8a23.
//
// Solidity: function hasMadeClaim(bytes32 ) view returns(bool)
func (_Oracle *OracleCaller) HasMadeClaim(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "hasMadeClaim", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMadeClaim is a free data retrieval call binding the contract method 0x129c8a23.
//
// Solidity: function hasMadeClaim(bytes32 ) view returns(bool)
func (_Oracle *OracleSession) HasMadeClaim(arg0 [32]byte) (bool, error) {
	return _Oracle.Contract.HasMadeClaim(&_Oracle.CallOpts, arg0)
}

// HasMadeClaim is a free data retrieval call binding the contract method 0x129c8a23.
//
// Solidity: function hasMadeClaim(bytes32 ) view returns(bool)
func (_Oracle *OracleCallerSession) HasMadeClaim(arg0 [32]byte) (bool, error) {
	return _Oracle.Contract.HasMadeClaim(&_Oracle.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleSession) Operator() (common.Address, error) {
	return _Oracle.Contract.Operator(&_Oracle.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Oracle *OracleCallerSession) Operator() (common.Address, error) {
	return _Oracle.Contract.Operator(&_Oracle.CallOpts)
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x8350549f.
//
// Solidity: function oracleClaimValidators(bytes32 , uint256 ) view returns(address)
func (_Oracle *OracleCaller) OracleClaimValidators(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "oracleClaimValidators", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x8350549f.
//
// Solidity: function oracleClaimValidators(bytes32 , uint256 ) view returns(address)
func (_Oracle *OracleSession) OracleClaimValidators(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Oracle.Contract.OracleClaimValidators(&_Oracle.CallOpts, arg0, arg1)
}

// OracleClaimValidators is a free data retrieval call binding the contract method 0x8350549f.
//
// Solidity: function oracleClaimValidators(bytes32 , uint256 ) view returns(address)
func (_Oracle *OracleCallerSession) OracleClaimValidators(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Oracle.Contract.OracleClaimValidators(&_Oracle.CallOpts, arg0, arg1)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleCaller) Valset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "valset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleSession) Valset() (common.Address, error) {
	return _Oracle.Contract.Valset(&_Oracle.CallOpts)
}

// Valset is a free data retrieval call binding the contract method 0x7f54af0c.
//
// Solidity: function valset() view returns(address)
func (_Oracle *OracleCallerSession) Valset() (common.Address, error) {
	return _Oracle.Contract.Valset(&_Oracle.CallOpts)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x71fb9358.
//
// Solidity: function newOracleClaim(uint8 _claimType, address _ethereumReceiver, address _tokenAddress, bytes _srcTxHash, uint256 _amount, bytes32 _claimID, bytes _signature1, bytes _signature2) returns()
func (_Oracle *OracleTransactor) NewOracleClaim(opts *bind.TransactOpts, _claimType uint8, _ethereumReceiver common.Address, _tokenAddress common.Address, _srcTxHash []byte, _amount *big.Int, _claimID [32]byte, _signature1 []byte, _signature2 []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "newOracleClaim", _claimType, _ethereumReceiver, _tokenAddress, _srcTxHash, _amount, _claimID, _signature1, _signature2)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x71fb9358.
//
// Solidity: function newOracleClaim(uint8 _claimType, address _ethereumReceiver, address _tokenAddress, bytes _srcTxHash, uint256 _amount, bytes32 _claimID, bytes _signature1, bytes _signature2) returns()
func (_Oracle *OracleSession) NewOracleClaim(_claimType uint8, _ethereumReceiver common.Address, _tokenAddress common.Address, _srcTxHash []byte, _amount *big.Int, _claimID [32]byte, _signature1 []byte, _signature2 []byte) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleClaim(&_Oracle.TransactOpts, _claimType, _ethereumReceiver, _tokenAddress, _srcTxHash, _amount, _claimID, _signature1, _signature2)
}

// NewOracleClaim is a paid mutator transaction binding the contract method 0x71fb9358.
//
// Solidity: function newOracleClaim(uint8 _claimType, address _ethereumReceiver, address _tokenAddress, bytes _srcTxHash, uint256 _amount, bytes32 _claimID, bytes _signature1, bytes _signature2) returns()
func (_Oracle *OracleTransactorSession) NewOracleClaim(_claimType uint8, _ethereumReceiver common.Address, _tokenAddress common.Address, _srcTxHash []byte, _amount *big.Int, _claimID [32]byte, _signature1 []byte, _signature2 []byte) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleClaim(&_Oracle.TransactOpts, _claimType, _ethereumReceiver, _tokenAddress, _srcTxHash, _amount, _claimID, _signature1, _signature2)
}

// OracleLogNewOracleClaimIterator is returned from FilterLogNewOracleClaim and is used to iterate over the raw logs and unpacked data for LogNewOracleClaim events raised by the Oracle contract.
type OracleLogNewOracleClaimIterator struct {
	Event *OracleLogNewOracleClaim // Event containing the contract specifics and raw log

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
func (it *OracleLogNewOracleClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLogNewOracleClaim)
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
		it.Event = new(OracleLogNewOracleClaim)
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
func (it *OracleLogNewOracleClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLogNewOracleClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLogNewOracleClaim represents a LogNewOracleClaim event raised by the Oracle contract.
type OracleLogNewOracleClaim struct {
	ClaimID          [32]byte
	ValidatorAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewOracleClaim is a free log retrieval operation binding the contract event 0x9573b45c7fa80d6b556ba50357ea808a3c3304c99959a440de796362c0b0e57b.
//
// Solidity: event LogNewOracleClaim(bytes32 _claimID, address _validatorAddress)
func (_Oracle *OracleFilterer) FilterLogNewOracleClaim(opts *bind.FilterOpts) (*OracleLogNewOracleClaimIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return &OracleLogNewOracleClaimIterator{contract: _Oracle.contract, event: "LogNewOracleClaim", logs: logs, sub: sub}, nil
}

// WatchLogNewOracleClaim is a free log subscription operation binding the contract event 0x9573b45c7fa80d6b556ba50357ea808a3c3304c99959a440de796362c0b0e57b.
//
// Solidity: event LogNewOracleClaim(bytes32 _claimID, address _validatorAddress)
func (_Oracle *OracleFilterer) WatchLogNewOracleClaim(opts *bind.WatchOpts, sink chan<- *OracleLogNewOracleClaim) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LogNewOracleClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLogNewOracleClaim)
				if err := _Oracle.contract.UnpackLog(event, "LogNewOracleClaim", log); err != nil {
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

// ParseLogNewOracleClaim is a log parse operation binding the contract event 0x9573b45c7fa80d6b556ba50357ea808a3c3304c99959a440de796362c0b0e57b.
//
// Solidity: event LogNewOracleClaim(bytes32 _claimID, address _validatorAddress)
func (_Oracle *OracleFilterer) ParseLogNewOracleClaim(log types.Log) (*OracleLogNewOracleClaim, error) {
	event := new(OracleLogNewOracleClaim)
	if err := _Oracle.contract.UnpackLog(event, "LogNewOracleClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleLogProphecyProcessedIterator is returned from FilterLogProphecyProcessed and is used to iterate over the raw logs and unpacked data for LogProphecyProcessed events raised by the Oracle contract.
type OracleLogProphecyProcessedIterator struct {
	Event *OracleLogProphecyProcessed // Event containing the contract specifics and raw log

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
func (it *OracleLogProphecyProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLogProphecyProcessed)
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
		it.Event = new(OracleLogProphecyProcessed)
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
func (it *OracleLogProphecyProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLogProphecyProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLogProphecyProcessed represents a LogProphecyProcessed event raised by the Oracle contract.
type OracleLogProphecyProcessed struct {
	ClaimID   [32]byte
	Submitter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProphecyProcessed is a free log retrieval operation binding the contract event 0x3fe675dd72294e681d888f819675033de7f7d2e84dd7da60daab7af319c79345.
//
// Solidity: event LogProphecyProcessed(bytes32 _claimID, address _submitter)
func (_Oracle *OracleFilterer) FilterLogProphecyProcessed(opts *bind.FilterOpts) (*OracleLogProphecyProcessedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LogProphecyProcessed")
	if err != nil {
		return nil, err
	}
	return &OracleLogProphecyProcessedIterator{contract: _Oracle.contract, event: "LogProphecyProcessed", logs: logs, sub: sub}, nil
}

// WatchLogProphecyProcessed is a free log subscription operation binding the contract event 0x3fe675dd72294e681d888f819675033de7f7d2e84dd7da60daab7af319c79345.
//
// Solidity: event LogProphecyProcessed(bytes32 _claimID, address _submitter)
func (_Oracle *OracleFilterer) WatchLogProphecyProcessed(opts *bind.WatchOpts, sink chan<- *OracleLogProphecyProcessed) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LogProphecyProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLogProphecyProcessed)
				if err := _Oracle.contract.UnpackLog(event, "LogProphecyProcessed", log); err != nil {
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

// ParseLogProphecyProcessed is a log parse operation binding the contract event 0x3fe675dd72294e681d888f819675033de7f7d2e84dd7da60daab7af319c79345.
//
// Solidity: event LogProphecyProcessed(bytes32 _claimID, address _submitter)
func (_Oracle *OracleFilterer) ParseLogProphecyProcessed(log types.Log) (*OracleLogProphecyProcessed, error) {
	event := new(OracleLogProphecyProcessed)
	if err := _Oracle.contract.UnpackLog(event, "LogProphecyProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212206b9dedae230e32030be7f4b33f6eb493594d1636c6d51b25126c65d453a2f09864736f6c63430008140033",
}

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}

// ValsetMetaData contains all meta data concerning the Valset contract.
var ValsetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_initValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_initPowers\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValidatorPowerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_currentValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValsetReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newValsetVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPower\",\"type\":\"uint256\"}],\"name\":\"LogValsetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Debug_Packed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Debug_ethMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorPower\",\"type\":\"uint256\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"getValidatorPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_valsetVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"recoverGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newValidatorPower\",\"type\":\"uint256\"}],\"name\":\"updateValidatorPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620014f0380380620014f083398101604081905262000033916200049a565b5f80546001600160a01b0319166001600160a01b0385161781556002556200005c828262000065565b505050620005b1565b5f546001600160a01b03163314620000c45760405162461bcd60e51b815260206004820152601560248201527f4d75737420626520746865206f70657261746f722e000000000000000000000060448201526064015b60405180910390fd5b80518251146200012f5760405162461bcd60e51b815260206004820152602f60248201527f45766572792076616c696461746f72206d7573742068617665206120636f727260448201526e32b9b837b73234b733903837bbb2b960891b6064820152608401620000bb565b62000139620001f0565b5f5b8251811015620001a4576200018f8382815181106200015e576200015e62000577565b60200260200101518383815181106200017b576200017b62000577565b60200260200101516200025060201b60201c565b6200019c8160016200058b565b90506200013b565b5060025460035460015460408051938452602084019290925282820152517f3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa69181900360600190a15050565b600254620002009060016200058b565b60028190555f6003819055600181905560408051928352602083018290528201527fd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f29060600160405180910390a1565b5f600254836040516020016200027d92919091825260601b6001600160601b031916602082015260340190565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915060ff1615620002f65760405162461bcd60e51b815260206004820152601360248201527f76616c646961746f7220697320616374697665000000000000000000000000006044820152606401620000bb565b600354620003069060016200058b565b600355600154620003199083906200058b565b60019081555f828152600460209081526040808320805460ff1916851790556005825291829020859055600254600354935483516001600160a01b038916815292830187905292820152606081019290925260808201527f1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a9060a00160405180910390a1505050565b80516001600160a01b0381168114620003b9575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715620003fd57620003fd620003be565b604052919050565b5f6001600160401b03821115620004205762000420620003be565b5060051b60200190565b5f82601f8301126200043a575f80fd5b81516020620004536200044d8362000405565b620003d2565b82815260059290921b8401810191818101908684111562000472575f80fd5b8286015b848110156200048f578051835291830191830162000476565b509695505050505050565b5f805f60608486031215620004ad575f80fd5b620004b884620003a2565b602085810151919450906001600160401b0380821115620004d7575f80fd5b818701915087601f830112620004eb575f80fd5b8151620004fc6200044d8262000405565b81815260059190911b8301840190848101908a8311156200051b575f80fd5b938501935b8285101562000544576200053485620003a2565b8252938501939085019062000520565b60408a015190975094505050808311156200055d575f80fd5b50506200056d868287016200042a565b9150509250925092565b634e487b7160e01b5f52603260045260245ffd5b80820180821115620005ab57634e487b7160e01b5f52601160045260245ffd5b92915050565b610f3180620005bf5f395ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c8063788cf92f11610093578063b872c52311610063578063b872c5231461021a578063db3ad22c14610239578063fa21089614610242578063fc6c1f02146102a2575f80fd5b8063788cf92f146101c95780638d56c37d146101dc5780639bdafcb3146101e5578063b5672be314610207575f80fd5b806340a141ff116100ce57806340a141ff1461017e578063473691a414610191578063570ca735146101a457806361bb63f3146101b6575f80fd5b80630f43a677146100ff57806319045a251461011b5780632e75293b1461014657806340550a1c1461015b575b5f80fd5b61010860035481565b6040519081526020015b60405180910390f35b61012e610129366004610b4a565b6102b5565b6040516001600160a01b039091168152602001610112565b610159610154366004610bfe565b6102d6565b005b61016e610169366004610c26565b610439565b6040519015158152602001610112565b61015961018c366004610c26565b61047e565b61010861019f366004610c26565b6105d8565b5f5461012e906001600160a01b031681565b6101086101c4366004610c46565b61061a565b6101596101d7366004610ce8565b61066c565b61010860025481565b61016e6101f3366004610c46565b60046020525f908152604090205460ff1681565b610159610215366004610da2565b6107a6565b610108610228366004610c46565b60056020525f908152604090205481565b61010860015481565b610295610250366004610c46565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8082019390935281518082039093018352605c01905290565b6040516101129190610dcc565b6101596102b0366004610bfe565b61088a565b5f806102c08461061a565b90506102cc81846108c1565b9150505b92915050565b5f546001600160a01b031633146103085760405162461bcd60e51b81526004016102ff90610e17565b60405180910390fd5b5f6002548360405160200161031e929190610e46565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915060ff166103ab5760405162461bcd60e51b815260206004820152602e60248201527f43616e206f6e6c79207570646174652074686520706f776572206f662061637460448201526d6976652076616c646961746f727360901b60648201526084016102ff565b5f818152600560205260409020546001546103c7908290610e7a565b60018190556103d7908490610e8d565b60019081555f8381526005602052604090819020859055600254600354925491517f335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d53409361042b938993899390929190610ea0565b60405180910390a150505050565b5f8060025483604051602001610450929190610e46565b60408051601f1981840301815291815281516020928301205f908152600490925290205460ff169392505050565b5f546001600160a01b031633146104a75760405162461bcd60e51b81526004016102ff90610e17565b5f600254826040516020016104bd929190610e46565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915060ff1661053d5760405162461bcd60e51b815260206004820152602160248201527f43616e206f6e6c792072656d6f7665206163746976652076616c646961746f726044820152607360f81b60648201526084016102ff565b600160035461054c9190610e7a565b6003555f8181526005602052604090205460015461056a9190610e7a565b60019081555f828152600460209081526040808320805460ff191690556005909152808220829055600254600354935491517f1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c946105cc948894909392610ea0565b60405180910390a15050565b5f80600254836040516020016105ef929190610e46565b60408051601f1981840301815291815281516020928301205f90815260059092529020549392505050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018290525f90605c01604051602081830303815290604052805190602001209050919050565b5f546001600160a01b031633146106955760405162461bcd60e51b81526004016102ff90610e17565b80518251146106fe5760405162461bcd60e51b815260206004820152602f60248201527f45766572792076616c696461746f72206d7573742068617665206120636f727260448201526e32b9b837b73234b733903837bbb2b960891b60648201526084016102ff565b610706610990565b5f5b825181101561075f5761074d83828151811061072657610726610ece565b602002602001015183838151811061074057610740610ece565b60200260200101516109ee565b610758816001610e8d565b9050610708565b50600254600354600154604080519384526020840192909252908201527f3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6906060016105cc565b5f546001600160a01b031633146107cf5760405162461bcd60e51b81526004016102ff90610e17565b600254821061083e5760405162461bcd60e51b815260206004820152603560248201527f476173207265636f76657279206f6e6c7920616c6c6f77656420666f722070726044820152746576696f75732076616c696461746f72207365747360581b60648201526084016102ff565b5f8282604051602001610852929190610e46565b60408051601f1981840301815291815281516020928301205f90815260048352818120805460ff191690556005909252812055505050565b5f546001600160a01b031633146108b35760405162461bcd60e51b81526004016102ff90610e17565b6108bd82826109ee565b5050565b5f805f8084516041146108d9575f93505050506102d0565b5050506020820151604083015160608401515f1a601b81101561090457610901601b82610ee2565b90505b8060ff16601b1415801561091c57508060ff16601c14155b1561092c575f93505050506102d0565b604080515f81526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa15801561097c573d5f803e3d5ffd5b5050506020604051035193505050506102d0565b60025461099e906001610e8d565b60028190555f6003819055600181905560408051928352602083018290528201527fd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f29060600160405180910390a1565b5f60025483604051602001610a04929190610e46565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915060ff1615610a715760405162461bcd60e51b815260206004820152601360248201527276616c646961746f722069732061637469766560681b60448201526064016102ff565b600354610a7f906001610e8d565b600355600154610a90908390610e8d565b60019081555f828152600460209081526040808320805460ff191685179055600590915290819020849055600254600354925491517f1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a93610af8938893889390929190610ea0565b60405180910390a1505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610b4257610b42610b05565b604052919050565b5f8060408385031215610b5b575f80fd5b8235915060208084013567ffffffffffffffff80821115610b7a575f80fd5b818601915086601f830112610b8d575f80fd5b813581811115610b9f57610b9f610b05565b610bb1601f8201601f19168501610b19565b91508082528784828501011115610bc6575f80fd5b80848401858401375f848284010152508093505050509250929050565b80356001600160a01b0381168114610bf9575f80fd5b919050565b5f8060408385031215610c0f575f80fd5b610c1883610be3565b946020939093013593505050565b5f60208284031215610c36575f80fd5b610c3f82610be3565b9392505050565b5f60208284031215610c56575f80fd5b5035919050565b5f67ffffffffffffffff821115610c7657610c76610b05565b5060051b60200190565b5f82601f830112610c8f575f80fd5b81356020610ca4610c9f83610c5d565b610b19565b82815260059290921b84018101918181019086841115610cc2575f80fd5b8286015b84811015610cdd5780358352918301918301610cc6565b509695505050505050565b5f8060408385031215610cf9575f80fd5b823567ffffffffffffffff80821115610d10575f80fd5b818501915085601f830112610d23575f80fd5b81356020610d33610c9f83610c5d565b82815260059290921b84018101918181019089841115610d51575f80fd5b948201945b83861015610d7657610d6786610be3565b82529482019490820190610d56565b96505086013592505080821115610d8b575f80fd5b50610d9885828601610c80565b9150509250929050565b5f8060408385031215610db3575f80fd5b82359150610dc360208401610be3565b90509250929050565b5f6020808352835180828501525f5b81811015610df757858101830151858201604001528201610ddb565b505f604082860101526040601f19601f8301168501019250505092915050565b60208082526015908201527426bab9ba103132903a34329037b832b930ba37b91760591b604082015260600190565b91825260601b6bffffffffffffffffffffffff1916602082015260340190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156102d0576102d0610e66565b808201808211156102d0576102d0610e66565b6001600160a01b03959095168552602085019390935260408401919091526060830152608082015260a00190565b634e487b7160e01b5f52603260045260245ffd5b60ff81811683821601908111156102d0576102d0610e6656fea26469706673582212203dd1834af08e0a1b7d5943304aa07288cc7ab3cd5c13fc93bf7d3d7ff3581ca964736f6c63430008140033",
}

// ValsetABI is the input ABI used to generate the binding from.
// Deprecated: Use ValsetMetaData.ABI instead.
var ValsetABI = ValsetMetaData.ABI

// ValsetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValsetMetaData.Bin instead.
var ValsetBin = ValsetMetaData.Bin

// DeployValset deploys a new Ethereum contract, binding an instance of Valset to it.
func DeployValset(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _initValidators []common.Address, _initPowers []*big.Int) (common.Address, *types.Transaction, *Valset, error) {
	parsed, err := ValsetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValsetBin), backend, _operator, _initValidators, _initPowers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Valset{ValsetCaller: ValsetCaller{contract: contract}, ValsetTransactor: ValsetTransactor{contract: contract}, ValsetFilterer: ValsetFilterer{contract: contract}}, nil
}

// Valset is an auto generated Go binding around an Ethereum contract.
type Valset struct {
	ValsetCaller     // Read-only binding to the contract
	ValsetTransactor // Write-only binding to the contract
	ValsetFilterer   // Log filterer for contract events
}

// ValsetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValsetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValsetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValsetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValsetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValsetSession struct {
	Contract     *Valset           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValsetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValsetCallerSession struct {
	Contract *ValsetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValsetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValsetTransactorSession struct {
	Contract     *ValsetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValsetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValsetRaw struct {
	Contract *Valset // Generic contract binding to access the raw methods on
}

// ValsetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValsetCallerRaw struct {
	Contract *ValsetCaller // Generic read-only contract binding to access the raw methods on
}

// ValsetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValsetTransactorRaw struct {
	Contract *ValsetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValset creates a new instance of Valset, bound to a specific deployed contract.
func NewValset(address common.Address, backend bind.ContractBackend) (*Valset, error) {
	contract, err := bindValset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Valset{ValsetCaller: ValsetCaller{contract: contract}, ValsetTransactor: ValsetTransactor{contract: contract}, ValsetFilterer: ValsetFilterer{contract: contract}}, nil
}

// NewValsetCaller creates a new read-only instance of Valset, bound to a specific deployed contract.
func NewValsetCaller(address common.Address, caller bind.ContractCaller) (*ValsetCaller, error) {
	contract, err := bindValset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValsetCaller{contract: contract}, nil
}

// NewValsetTransactor creates a new write-only instance of Valset, bound to a specific deployed contract.
func NewValsetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValsetTransactor, error) {
	contract, err := bindValset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValsetTransactor{contract: contract}, nil
}

// NewValsetFilterer creates a new log filterer instance of Valset, bound to a specific deployed contract.
func NewValsetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValsetFilterer, error) {
	contract, err := bindValset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValsetFilterer{contract: contract}, nil
}

// bindValset binds a generic wrapper to an already deployed contract.
func bindValset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValsetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Valset *ValsetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Valset.Contract.ValsetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Valset *ValsetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Valset.Contract.ValsetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Valset *ValsetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Valset.Contract.ValsetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Valset *ValsetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Valset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Valset *ValsetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Valset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Valset *ValsetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Valset.Contract.contract.Transact(opts, method, params...)
}

// DebugPacked is a free data retrieval call binding the contract method 0xfa210896.
//
// Solidity: function Debug_Packed(bytes32 hash) pure returns(bytes)
func (_Valset *ValsetCaller) DebugPacked(opts *bind.CallOpts, hash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "Debug_Packed", hash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DebugPacked is a free data retrieval call binding the contract method 0xfa210896.
//
// Solidity: function Debug_Packed(bytes32 hash) pure returns(bytes)
func (_Valset *ValsetSession) DebugPacked(hash [32]byte) ([]byte, error) {
	return _Valset.Contract.DebugPacked(&_Valset.CallOpts, hash)
}

// DebugPacked is a free data retrieval call binding the contract method 0xfa210896.
//
// Solidity: function Debug_Packed(bytes32 hash) pure returns(bytes)
func (_Valset *ValsetCallerSession) DebugPacked(hash [32]byte) ([]byte, error) {
	return _Valset.Contract.DebugPacked(&_Valset.CallOpts, hash)
}

// DebugEthMessageHash is a free data retrieval call binding the contract method 0x61bb63f3.
//
// Solidity: function Debug_ethMessageHash(bytes32 hash) pure returns(bytes32)
func (_Valset *ValsetCaller) DebugEthMessageHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "Debug_ethMessageHash", hash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DebugEthMessageHash is a free data retrieval call binding the contract method 0x61bb63f3.
//
// Solidity: function Debug_ethMessageHash(bytes32 hash) pure returns(bytes32)
func (_Valset *ValsetSession) DebugEthMessageHash(hash [32]byte) ([32]byte, error) {
	return _Valset.Contract.DebugEthMessageHash(&_Valset.CallOpts, hash)
}

// DebugEthMessageHash is a free data retrieval call binding the contract method 0x61bb63f3.
//
// Solidity: function Debug_ethMessageHash(bytes32 hash) pure returns(bytes32)
func (_Valset *ValsetCallerSession) DebugEthMessageHash(hash [32]byte) ([32]byte, error) {
	return _Valset.Contract.DebugEthMessageHash(&_Valset.CallOpts, hash)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_Valset *ValsetCaller) CurrentValsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "currentValsetVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_Valset *ValsetSession) CurrentValsetVersion() (*big.Int, error) {
	return _Valset.Contract.CurrentValsetVersion(&_Valset.CallOpts)
}

// CurrentValsetVersion is a free data retrieval call binding the contract method 0x8d56c37d.
//
// Solidity: function currentValsetVersion() view returns(uint256)
func (_Valset *ValsetCallerSession) CurrentValsetVersion() (*big.Int, error) {
	return _Valset.Contract.CurrentValsetVersion(&_Valset.CallOpts)
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_Valset *ValsetCaller) GetValidatorPower(opts *bind.CallOpts, _validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "getValidatorPower", _validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_Valset *ValsetSession) GetValidatorPower(_validatorAddress common.Address) (*big.Int, error) {
	return _Valset.Contract.GetValidatorPower(&_Valset.CallOpts, _validatorAddress)
}

// GetValidatorPower is a free data retrieval call binding the contract method 0x473691a4.
//
// Solidity: function getValidatorPower(address _validatorAddress) view returns(uint256)
func (_Valset *ValsetCallerSession) GetValidatorPower(_validatorAddress common.Address) (*big.Int, error) {
	return _Valset.Contract.GetValidatorPower(&_Valset.CallOpts, _validatorAddress)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_Valset *ValsetCaller) IsActiveValidator(opts *bind.CallOpts, _validatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "isActiveValidator", _validatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_Valset *ValsetSession) IsActiveValidator(_validatorAddress common.Address) (bool, error) {
	return _Valset.Contract.IsActiveValidator(&_Valset.CallOpts, _validatorAddress)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddress) view returns(bool)
func (_Valset *ValsetCallerSession) IsActiveValidator(_validatorAddress common.Address) (bool, error) {
	return _Valset.Contract.IsActiveValidator(&_Valset.CallOpts, _validatorAddress)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Valset *ValsetCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Valset *ValsetSession) Operator() (common.Address, error) {
	return _Valset.Contract.Operator(&_Valset.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Valset *ValsetCallerSession) Operator() (common.Address, error) {
	return _Valset.Contract.Operator(&_Valset.CallOpts)
}

// Powers is a free data retrieval call binding the contract method 0xb872c523.
//
// Solidity: function powers(bytes32 ) view returns(uint256)
func (_Valset *ValsetCaller) Powers(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "powers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Powers is a free data retrieval call binding the contract method 0xb872c523.
//
// Solidity: function powers(bytes32 ) view returns(uint256)
func (_Valset *ValsetSession) Powers(arg0 [32]byte) (*big.Int, error) {
	return _Valset.Contract.Powers(&_Valset.CallOpts, arg0)
}

// Powers is a free data retrieval call binding the contract method 0xb872c523.
//
// Solidity: function powers(bytes32 ) view returns(uint256)
func (_Valset *ValsetCallerSession) Powers(arg0 [32]byte) (*big.Int, error) {
	return _Valset.Contract.Powers(&_Valset.CallOpts, arg0)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _message, bytes _signature) pure returns(address)
func (_Valset *ValsetCaller) Recover(opts *bind.CallOpts, _message [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "recover", _message, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _message, bytes _signature) pure returns(address)
func (_Valset *ValsetSession) Recover(_message [32]byte, _signature []byte) (common.Address, error) {
	return _Valset.Contract.Recover(&_Valset.CallOpts, _message, _signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _message, bytes _signature) pure returns(address)
func (_Valset *ValsetCallerSession) Recover(_message [32]byte, _signature []byte) (common.Address, error) {
	return _Valset.Contract.Recover(&_Valset.CallOpts, _message, _signature)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Valset *ValsetCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "totalPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Valset *ValsetSession) TotalPower() (*big.Int, error) {
	return _Valset.Contract.TotalPower(&_Valset.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() view returns(uint256)
func (_Valset *ValsetCallerSession) TotalPower() (*big.Int, error) {
	return _Valset.Contract.TotalPower(&_Valset.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Valset *ValsetCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Valset *ValsetSession) ValidatorCount() (*big.Int, error) {
	return _Valset.Contract.ValidatorCount(&_Valset.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Valset *ValsetCallerSession) ValidatorCount() (*big.Int, error) {
	return _Valset.Contract.ValidatorCount(&_Valset.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(bool)
func (_Valset *ValsetCaller) Validators(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Valset.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(bool)
func (_Valset *ValsetSession) Validators(arg0 [32]byte) (bool, error) {
	return _Valset.Contract.Validators(&_Valset.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x9bdafcb3.
//
// Solidity: function validators(bytes32 ) view returns(bool)
func (_Valset *ValsetCallerSession) Validators(arg0 [32]byte) (bool, error) {
	return _Valset.Contract.Validators(&_Valset.CallOpts, arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_Valset *ValsetTransactor) AddValidator(opts *bind.TransactOpts, _validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "addValidator", _validatorAddress, _validatorPower)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_Valset *ValsetSession) AddValidator(_validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.AddValidator(&_Valset.TransactOpts, _validatorAddress, _validatorPower)
}

// AddValidator is a paid mutator transaction binding the contract method 0xfc6c1f02.
//
// Solidity: function addValidator(address _validatorAddress, uint256 _validatorPower) returns()
func (_Valset *ValsetTransactorSession) AddValidator(_validatorAddress common.Address, _validatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.AddValidator(&_Valset.TransactOpts, _validatorAddress, _validatorPower)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_Valset *ValsetTransactor) RecoverGas(opts *bind.TransactOpts, _valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "recoverGas", _valsetVersion, _validatorAddress)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_Valset *ValsetSession) RecoverGas(_valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RecoverGas(&_Valset.TransactOpts, _valsetVersion, _validatorAddress)
}

// RecoverGas is a paid mutator transaction binding the contract method 0xb5672be3.
//
// Solidity: function recoverGas(uint256 _valsetVersion, address _validatorAddress) returns()
func (_Valset *ValsetTransactorSession) RecoverGas(_valsetVersion *big.Int, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RecoverGas(&_Valset.TransactOpts, _valsetVersion, _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_Valset *ValsetTransactor) RemoveValidator(opts *bind.TransactOpts, _validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "removeValidator", _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_Valset *ValsetSession) RemoveValidator(_validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RemoveValidator(&_Valset.TransactOpts, _validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorAddress) returns()
func (_Valset *ValsetTransactorSession) RemoveValidator(_validatorAddress common.Address) (*types.Transaction, error) {
	return _Valset.Contract.RemoveValidator(&_Valset.TransactOpts, _validatorAddress)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_Valset *ValsetTransactor) UpdateValidatorPower(opts *bind.TransactOpts, _validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "updateValidatorPower", _validatorAddress, _newValidatorPower)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_Valset *ValsetSession) UpdateValidatorPower(_validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValidatorPower(&_Valset.TransactOpts, _validatorAddress, _newValidatorPower)
}

// UpdateValidatorPower is a paid mutator transaction binding the contract method 0x2e75293b.
//
// Solidity: function updateValidatorPower(address _validatorAddress, uint256 _newValidatorPower) returns()
func (_Valset *ValsetTransactorSession) UpdateValidatorPower(_validatorAddress common.Address, _newValidatorPower *big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValidatorPower(&_Valset.TransactOpts, _validatorAddress, _newValidatorPower)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_Valset *ValsetTransactor) UpdateValset(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Valset.contract.Transact(opts, "updateValset", _validators, _powers)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_Valset *ValsetSession) UpdateValset(_validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValset(&_Valset.TransactOpts, _validators, _powers)
}

// UpdateValset is a paid mutator transaction binding the contract method 0x788cf92f.
//
// Solidity: function updateValset(address[] _validators, uint256[] _powers) returns()
func (_Valset *ValsetTransactorSession) UpdateValset(_validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Valset.Contract.UpdateValset(&_Valset.TransactOpts, _validators, _powers)
}

// ValsetLogValidatorAddedIterator is returned from FilterLogValidatorAdded and is used to iterate over the raw logs and unpacked data for LogValidatorAdded events raised by the Valset contract.
type ValsetLogValidatorAddedIterator struct {
	Event *ValsetLogValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ValsetLogValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetLogValidatorAdded)
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
		it.Event = new(ValsetLogValidatorAdded)
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
func (it *ValsetLogValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetLogValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetLogValidatorAdded represents a LogValidatorAdded event raised by the Valset contract.
type ValsetLogValidatorAdded struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogValidatorAdded is a free log retrieval operation binding the contract event 0x1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a.
//
// Solidity: event LogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterLogValidatorAdded(opts *bind.FilterOpts) (*ValsetLogValidatorAddedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "LogValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &ValsetLogValidatorAddedIterator{contract: _Valset.contract, event: "LogValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogValidatorAdded is a free log subscription operation binding the contract event 0x1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a.
//
// Solidity: event LogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchLogValidatorAdded(opts *bind.WatchOpts, sink chan<- *ValsetLogValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "LogValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetLogValidatorAdded)
				if err := _Valset.contract.UnpackLog(event, "LogValidatorAdded", log); err != nil {
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

// ParseLogValidatorAdded is a log parse operation binding the contract event 0x1a396bcf647502e902dce665d58a0c1b25f982f193ab9a1d0f1500d8d927bf2a.
//
// Solidity: event LogValidatorAdded(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseLogValidatorAdded(log types.Log) (*ValsetLogValidatorAdded, error) {
	event := new(ValsetLogValidatorAdded)
	if err := _Valset.contract.UnpackLog(event, "LogValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValsetLogValidatorPowerUpdatedIterator is returned from FilterLogValidatorPowerUpdated and is used to iterate over the raw logs and unpacked data for LogValidatorPowerUpdated events raised by the Valset contract.
type ValsetLogValidatorPowerUpdatedIterator struct {
	Event *ValsetLogValidatorPowerUpdated // Event containing the contract specifics and raw log

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
func (it *ValsetLogValidatorPowerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetLogValidatorPowerUpdated)
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
		it.Event = new(ValsetLogValidatorPowerUpdated)
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
func (it *ValsetLogValidatorPowerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetLogValidatorPowerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetLogValidatorPowerUpdated represents a LogValidatorPowerUpdated event raised by the Valset contract.
type ValsetLogValidatorPowerUpdated struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogValidatorPowerUpdated is a free log retrieval operation binding the contract event 0x335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d5340.
//
// Solidity: event LogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterLogValidatorPowerUpdated(opts *bind.FilterOpts) (*ValsetLogValidatorPowerUpdatedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "LogValidatorPowerUpdated")
	if err != nil {
		return nil, err
	}
	return &ValsetLogValidatorPowerUpdatedIterator{contract: _Valset.contract, event: "LogValidatorPowerUpdated", logs: logs, sub: sub}, nil
}

// WatchLogValidatorPowerUpdated is a free log subscription operation binding the contract event 0x335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d5340.
//
// Solidity: event LogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchLogValidatorPowerUpdated(opts *bind.WatchOpts, sink chan<- *ValsetLogValidatorPowerUpdated) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "LogValidatorPowerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetLogValidatorPowerUpdated)
				if err := _Valset.contract.UnpackLog(event, "LogValidatorPowerUpdated", log); err != nil {
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

// ParseLogValidatorPowerUpdated is a log parse operation binding the contract event 0x335940ce4119f8aae891d73dba74510a3d51f6210134d058237f26e6a31d5340.
//
// Solidity: event LogValidatorPowerUpdated(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseLogValidatorPowerUpdated(log types.Log) (*ValsetLogValidatorPowerUpdated, error) {
	event := new(ValsetLogValidatorPowerUpdated)
	if err := _Valset.contract.UnpackLog(event, "LogValidatorPowerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValsetLogValidatorRemovedIterator is returned from FilterLogValidatorRemoved and is used to iterate over the raw logs and unpacked data for LogValidatorRemoved events raised by the Valset contract.
type ValsetLogValidatorRemovedIterator struct {
	Event *ValsetLogValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ValsetLogValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetLogValidatorRemoved)
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
		it.Event = new(ValsetLogValidatorRemoved)
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
func (it *ValsetLogValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetLogValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetLogValidatorRemoved represents a LogValidatorRemoved event raised by the Valset contract.
type ValsetLogValidatorRemoved struct {
	Validator            common.Address
	Power                *big.Int
	CurrentValsetVersion *big.Int
	ValidatorCount       *big.Int
	TotalPower           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogValidatorRemoved is a free log retrieval operation binding the contract event 0x1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c.
//
// Solidity: event LogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterLogValidatorRemoved(opts *bind.FilterOpts) (*ValsetLogValidatorRemovedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "LogValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &ValsetLogValidatorRemovedIterator{contract: _Valset.contract, event: "LogValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogValidatorRemoved is a free log subscription operation binding the contract event 0x1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c.
//
// Solidity: event LogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchLogValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValsetLogValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "LogValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetLogValidatorRemoved)
				if err := _Valset.contract.UnpackLog(event, "LogValidatorRemoved", log); err != nil {
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

// ParseLogValidatorRemoved is a log parse operation binding the contract event 0x1241fb43a101ff98ab819a1882097d4ccada51ba60f326c1981cc48840f55b8c.
//
// Solidity: event LogValidatorRemoved(address _validator, uint256 _power, uint256 _currentValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseLogValidatorRemoved(log types.Log) (*ValsetLogValidatorRemoved, error) {
	event := new(ValsetLogValidatorRemoved)
	if err := _Valset.contract.UnpackLog(event, "LogValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValsetLogValsetResetIterator is returned from FilterLogValsetReset and is used to iterate over the raw logs and unpacked data for LogValsetReset events raised by the Valset contract.
type ValsetLogValsetResetIterator struct {
	Event *ValsetLogValsetReset // Event containing the contract specifics and raw log

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
func (it *ValsetLogValsetResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetLogValsetReset)
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
		it.Event = new(ValsetLogValsetReset)
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
func (it *ValsetLogValsetResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetLogValsetResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetLogValsetReset represents a LogValsetReset event raised by the Valset contract.
type ValsetLogValsetReset struct {
	NewValsetVersion *big.Int
	ValidatorCount   *big.Int
	TotalPower       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogValsetReset is a free log retrieval operation binding the contract event 0xd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f2.
//
// Solidity: event LogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterLogValsetReset(opts *bind.FilterOpts) (*ValsetLogValsetResetIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "LogValsetReset")
	if err != nil {
		return nil, err
	}
	return &ValsetLogValsetResetIterator{contract: _Valset.contract, event: "LogValsetReset", logs: logs, sub: sub}, nil
}

// WatchLogValsetReset is a free log subscription operation binding the contract event 0xd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f2.
//
// Solidity: event LogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchLogValsetReset(opts *bind.WatchOpts, sink chan<- *ValsetLogValsetReset) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "LogValsetReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetLogValsetReset)
				if err := _Valset.contract.UnpackLog(event, "LogValsetReset", log); err != nil {
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

// ParseLogValsetReset is a log parse operation binding the contract event 0xd870653e19f161500290fd0c4ca41bf5cf2bcb1ba66448f41c66c512dabd65f2.
//
// Solidity: event LogValsetReset(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseLogValsetReset(log types.Log) (*ValsetLogValsetReset, error) {
	event := new(ValsetLogValsetReset)
	if err := _Valset.contract.UnpackLog(event, "LogValsetReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValsetLogValsetUpdatedIterator is returned from FilterLogValsetUpdated and is used to iterate over the raw logs and unpacked data for LogValsetUpdated events raised by the Valset contract.
type ValsetLogValsetUpdatedIterator struct {
	Event *ValsetLogValsetUpdated // Event containing the contract specifics and raw log

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
func (it *ValsetLogValsetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValsetLogValsetUpdated)
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
		it.Event = new(ValsetLogValsetUpdated)
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
func (it *ValsetLogValsetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValsetLogValsetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValsetLogValsetUpdated represents a LogValsetUpdated event raised by the Valset contract.
type ValsetLogValsetUpdated struct {
	NewValsetVersion *big.Int
	ValidatorCount   *big.Int
	TotalPower       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogValsetUpdated is a free log retrieval operation binding the contract event 0x3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6.
//
// Solidity: event LogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) FilterLogValsetUpdated(opts *bind.FilterOpts) (*ValsetLogValsetUpdatedIterator, error) {

	logs, sub, err := _Valset.contract.FilterLogs(opts, "LogValsetUpdated")
	if err != nil {
		return nil, err
	}
	return &ValsetLogValsetUpdatedIterator{contract: _Valset.contract, event: "LogValsetUpdated", logs: logs, sub: sub}, nil
}

// WatchLogValsetUpdated is a free log subscription operation binding the contract event 0x3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6.
//
// Solidity: event LogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) WatchLogValsetUpdated(opts *bind.WatchOpts, sink chan<- *ValsetLogValsetUpdated) (event.Subscription, error) {

	logs, sub, err := _Valset.contract.WatchLogs(opts, "LogValsetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValsetLogValsetUpdated)
				if err := _Valset.contract.UnpackLog(event, "LogValsetUpdated", log); err != nil {
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

// ParseLogValsetUpdated is a log parse operation binding the contract event 0x3a7ef0da3179668af8114719645585b5a37092ef2d66f187dcf63d83a221eaa6.
//
// Solidity: event LogValsetUpdated(uint256 _newValsetVersion, uint256 _validatorCount, uint256 _totalPower)
func (_Valset *ValsetFilterer) ParseLogValsetUpdated(log types.Log) (*ValsetLogValsetUpdated, error) {
	event := new(ValsetLogValsetUpdated)
	if err := _Valset.contract.UnpackLog(event, "LogValsetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

