// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_token\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_token\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_hashes\",\"type\":\"string[]\"}],\"name\":\"put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StorageFuncSigs maps the 4-byte function signature to its string representation.
var StorageFuncSigs = map[string]string{
	"693ec85e": "get(string)",
	"aa6ca808": "getTokens()",
	"e9a650f1": "put(string,string[])",
}

// StorageBin is the compiled bytecode used for deploying new contracts.
var StorageBin = "0x608060405234801561001057600080fd5b50610813806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063693ec85e14610046578063aa6ca8081461006f578063e9a650f114610077575b600080fd5b610059610054366004610590565b61008c565b6040516100669190610711565b60405180910390f35b610059610182565b61008a6100853660046105cb565b61025a565b005b606060008260405161009e9190610685565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156101775760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156101635780601f1061013857610100808354040283529160200191610163565b820191906000526020600020905b81548152906001019060200180831161014657829003601f168201915b5050505050815260200190600101906100cc565b505050509050919050565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156102515760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561023d5780601f106102125761010080835404028352916020019161023d565b820191906000526020600020905b81548152906001019060200180831161022057829003601f168201915b5050505050815260200190600101906101a6565b50505050905090565b60005b60015481101561035157826040516020016102789190610685565b604051602081830303815290604052805190602001206001828154811061029b57fe5b906000526020600020016040516020016102b591906106a1565b6040516020818303038152906040528051906020012014156103495760005b8251811015610342576000846040516102ed9190610685565b908152602001604051809103902083828151811061030757fe5b6020908102919091018101518254600181018455600093845292829020815161033994919091019291909101906103cd565b506001016102d4565b50506103c9565b60010161025d565b50806000836040516103639190610685565b90815260200160405180910390209080519060200190610384929190610459565b5060018054808201825560009190915282516103c7917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019060208501906103cd565b505b5050565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826104035760008555610449565b82601f1061041c57805160ff1916838001178555610449565b82800160010185558215610449579182015b8281111561044957825182559160200191906001019061042e565b506104559291506104b2565b5090565b8280548282559060005260206000209081019282156104a6579160200282015b828111156104a657825180516104969184916020909101906103cd565b5091602001919060010190610479565b506104559291506104c7565b5b8082111561045557600081556001016104b3565b808211156104555760006104db82826104e4565b506001016104c7565b50805460018160011615610100020316600290046000825580601f1061050a5750610528565b601f01602090049060005260206000209081019061052891906104b2565b50565b600082601f83011261053b578081fd5b813567ffffffffffffffff81111561054f57fe5b610562601f8201601f1916602001610789565b818152846020838601011115610576578283fd5b816020850160208301379081016020019190915292915050565b6000602082840312156105a1578081fd5b813567ffffffffffffffff8111156105b7578182fd5b6105c38482850161052b565b949350505050565b600080604083850312156105dd578081fd5b823567ffffffffffffffff808211156105f4578283fd5b6106008683870161052b565b9350602091508185013581811115610616578384fd5b8501601f81018713610626578384fd5b80358281111561063257fe5b61063f8485830201610789565b8181528481019350828501865b83811015610675576106638b88843588010161052b565b8652948601949086019060010161064c565b5096999098509650505050505050565b600082516106978184602087016107ad565b9190910192915050565b60008083546001808216600081146106c057600181146106d757610706565b60ff198316865260028304607f1686019350610706565b600283048786526020808720875b838110156106fe5781548a8201529085019082016106e5565b505050860193505b509195945050505050565b6000602080830181845280855180835260408601915060408482028701019250838701855b8281101561077c57878503603f190184528151805180875261075d818989018a85016107ad565b601f01601f191695909501860194509285019290850190600101610736565b5092979650505050505050565b60405181810167ffffffffffffffff811182821017156107a557fe5b604052919050565b60005b838110156107c85781810151838201526020016107b0565b838111156107d7576000848401525b5050505056fea2646970667358221220151372c11161807d1e8329e2eb2ba63b191b3a33654ea9df421ce75a2272f10964736f6c63430007060033"

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _token) view returns(string[])
func (_Storage *StorageCaller) Get(opts *bind.CallOpts, _token string) ([]string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "get", _token)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _token) view returns(string[])
func (_Storage *StorageSession) Get(_token string) ([]string, error) {
	return _Storage.Contract.Get(&_Storage.CallOpts, _token)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _token) view returns(string[])
func (_Storage *StorageCallerSession) Get(_token string) ([]string, error) {
	return _Storage.Contract.Get(&_Storage.CallOpts, _token)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(string[])
func (_Storage *StorageCaller) GetTokens(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(string[])
func (_Storage *StorageSession) GetTokens() ([]string, error) {
	return _Storage.Contract.GetTokens(&_Storage.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(string[])
func (_Storage *StorageCallerSession) GetTokens() ([]string, error) {
	return _Storage.Contract.GetTokens(&_Storage.CallOpts)
}

// Put is a paid mutator transaction binding the contract method 0xe9a650f1.
//
// Solidity: function put(string _token, string[] _hashes) returns()
func (_Storage *StorageTransactor) Put(opts *bind.TransactOpts, _token string, _hashes []string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "put", _token, _hashes)
}

// Put is a paid mutator transaction binding the contract method 0xe9a650f1.
//
// Solidity: function put(string _token, string[] _hashes) returns()
func (_Storage *StorageSession) Put(_token string, _hashes []string) (*types.Transaction, error) {
	return _Storage.Contract.Put(&_Storage.TransactOpts, _token, _hashes)
}

// Put is a paid mutator transaction binding the contract method 0xe9a650f1.
//
// Solidity: function put(string _token, string[] _hashes) returns()
func (_Storage *StorageTransactorSession) Put(_token string, _hashes []string) (*types.Transaction, error) {
	return _Storage.Contract.Put(&_Storage.TransactOpts, _token, _hashes)
}
