package contracts

import (
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Deploy Smart contract
func Deploy() {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/2d41bc4968a14c24960446e532b702ed")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "metamask")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _ := DeployStorage(
		auth,
		blockchain,
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)

}
