package config

import (
	"github.com/DEMYSTIF/gin-postgres-api/lib"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectEthereum(chainURL string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(chainURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func EthereumInstance(client *ethclient.Client, address string) (*lib.Grimoire, error) {
	contractAddress := common.HexToAddress(address)
	instance, err := lib.NewGrimoire(contractAddress, client)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
