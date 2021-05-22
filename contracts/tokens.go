package contracts

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const contractAddress = "0x9aa34941436414ebc637e29db6be1dce2361cea3"
const nodeURL = "https://rinkeby.infura.io/v3/2d41bc4968a14c24960446e532b702ed"
const key = `{"address":"a26eb2364fa0a455bd01d223e063be913c3b1ef8","crypto":{"cipher":"aes-128-ctr","ciphertext":"8b936280aefa4d157fc49c042c1b4d26a013289da0982b5e7c12b87e03dd5ab6","cipherparams":{"iv":"a1e0808b4feff6e790ac59c9d1e7bce0"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ed04a928f1c04849545a24118aedd5134abe07c98c80585e0f043dad9bf9cff1"},"mac":"6cde4fa9327b9ca0ee8140c71057e9c043d1283c718ae943f7e637ef8c63e6e4"},"id":"cb1436d8-0a1b-40cb-8a9d-7429e4883c6a","version":3}`

func GetTokens() ([]string, error) {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial(nodeURL)

	if err != nil {
		return []string{}, errors.New("Unable to connect to network:%v\n")
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
	contract, err := NewStorage(common.HexToAddress(contractAddress), blockchain)
	if err != nil {
		return []string{}, errors.New("Unable to bind to deployed instance of contract:%v\n")
	}

	return contract.GetTokens(nil)
}

func Put(token string, hashes []string) error {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial(nodeURL)

	if err != nil {
		return errors.New("Unable to connect to network:%v\n")
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
	contract, err := NewStorage(common.HexToAddress(contractAddress), blockchain)
	if err != nil {
		return errors.New("Unable to bind to deployed instance of contract:%v\n")
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "metamask")
	if err != nil {
		return errors.New("Failed to create authorized transactor")
	}

	_, err = contract.Put(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, token, hashes)

	if err != nil {
		return err
	}
	return nil
}

func Get(token string) ([]string, error) {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial(nodeURL)

	if err != nil {
		return []string{}, errors.New("Unable to connect to network:%v\n")
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
	contract, err := NewStorage(common.HexToAddress(contractAddress), blockchain)
	if err != nil {
		return []string{}, errors.New("Unable to bind to deployed instance of contract:%v\n")
	}

	return contract.Get(nil, token)
}
