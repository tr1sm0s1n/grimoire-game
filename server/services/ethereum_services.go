package services

import (
	"context"
	"os"
	"time"

	"github.com/DEMYSTIF/gin-postgres-api/config"
	"github.com/DEMYSTIF/gin-postgres-api/middlewares"
	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func IssueNFT(player *models.Player) error {
	eth, err := config.ConnectEthereum(os.Getenv("ETH_URL"))
	if err != nil {
		return err
	}

	instance, err := config.EthereumInstance(eth,  os.Getenv("ETH_ADDR"))
	if err != nil {
		return err
	}

	auth, err := middlewares.AuthGenerator(eth)
	if err != nil {
		return err
	}

	addrHex := common.HexToAddress(player.Address)
	trx, err := instance.Mint(auth, addrHex, player.UserName)
	if err != nil {
		return err
	}

	for {
		time.Sleep(1 * time.Second)

		txr, err := eth.TransactionReceipt(context.Background(), trx.Hash())
		if err != nil {
			if err == ethereum.NotFound {
				continue
			}
			return err
		}

		if txr.Status == 1 {
			break
		}
	}

	return nil
}