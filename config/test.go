package main

import (
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
	"github.com/stellar/go/xdr"
)

func main() {

	// Make a keypair for a known account from a secret seed
	// club waste muscle trumpet huge obey fold math guess broccoli trial caution
	kp, adderr := keypair.Parse("SANRELHUNYIUTGL25NWDGOYGF2TBCPKQWG23KGHJ5JSIQTJ7JJFO7DBB")

	// kp, adderr := keypair.Parse("GANHJTZUIKQCUEPGPUIOGR6DRVF2XFDKKU5VDVGBYTW2WEOUMNMZU7CO")
	// fmt.Println("tets", kp)
	if adderr != nil {
		fmt.Println("addres", adderr)
	}
	// Get the current state of the account from the network
	client := horizonclient.DefaultTestNetClient
	ar := horizonclient.AccountRequest{AccountID: kp.Address()} //give address here
	sourceAccount, err := client.AccountDetail(ar)
	if err != nil {
		log.Fatalln(err)
	}
	// asset := txnbuild.NativeAsset{}

	/////
	// Build an operation to create and fund a new account
	// op := txnbuild.Payment{
	// 	Destination:   "GDDDEQJO64LIDHS4NJOXAZJAQXS3ZDIEM37IOR4PKVU4A2WR6EET3T6W",
	// 	Amount:        "1000",
	// 	Asset:         asset,
	// 	SourceAccount: "GANHJTZUIKQCUEPGPUIOGR6DRVF2XFDKKU5VDVGBYTW2WEOUMNMZU7CO",
	// }
	///
	accountId := xdr.MustAddress("CABKK7CKUV6LSSCREM2TABPFKLA5KCQLGBTYL7WAUBPNFTHFHGO57DUC")
	id := xdr.ScString("new")
	aa := xdr.MustAddress("GDDDEQJO64LIDHS4NJOXAZJAQXS3ZDIEM37IOR4PKVU4A2WR6EET3T6W")
	// addr := xdr.ScAddress("GDDDEQJO64LIDHS4NJOXAZJAQXS3ZDIEM37IOR4PKVU4A2WR6EET3T6W")
	addr := xdr.ScAddress{AccountId: &aa}
	score := xdr.Uint64(5)
	// Convert the contract address to the appropriate XDR type (e.g., xdr.AccountId, xdr.ScAddress, or other types used by your InvokeContractArgs)

	invokeHostFunctionOp := txnbuild.InvokeHostFunction{
		HostFunction: xdr.HostFunction{
			Type: xdr.HostFunctionTypeHostFunctionTypeInvokeContract,
			InvokeContract: &xdr.InvokeContractArgs{
				ContractAddress: xdr.ScAddress{
					Type:      xdr.ScAddressTypeScAddressTypeContract,
					AccountId: &accountId,
				},
				FunctionName: "mint",
				Args: xdr.ScVec{
					xdr.ScVal{
						Type: xdr.ScValTypeScvString,
						Str:  &id,
					},
					xdr.ScVal{
						Type:    xdr.ScValTypeScvAddress,
						Address: &addr,
					},
					xdr.ScVal{
						Type: xdr.ScValTypeScvU64,
						U64:  &score,
					},
				},
			},

			// InvokeContract: &xdr.InvokeContractArgs{ContractAddress: xdrContractAddress},
		},
		Auth:          []xdr.SorobanAuthorizationEntry{},
		SourceAccount: "GANHJTZUIKQCUEPGPUIOGR6DRVF2XFDKKU5VDVGBYTW2WEOUMNMZU7CO",
	}

	// Construct the transaction that holds the operations to execute on the network
	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           []txnbuild.Operation{&invokeHostFunctionOp},
			BaseFee:              txnbuild.MinBaseFee,
			Preconditions:        txnbuild.Preconditions{TimeBounds: txnbuild.NewInfiniteTimeout()},
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	// kps, keyErr := keypair.Random()
	// kps, keyErr := keypair.ParseFull("club waste muscle trumpet huge obey fold math guess broccoli trial caution")
	// kps, keyErr := keypair.ParseFull("clubwastemuscletrumpethugeobeyfoldmathguessbroccolitrialcaution")

	// if keyErr != nil {
	// 	fmt.Println("keyerr", keyErr)
	// }

	// Sign the transaction
	tx, err = tx.Sign(network.TestNetworkPassphrase, kp.(*keypair.Full))
	if err != nil {
		log.Fatalln("sign err", err)
	}

	// Get the base 64 encoded transaction envelope
	txe, err := tx.Base64()
	if err != nil {
		log.Fatalln("base err", err)
	}

	// Send the transaction to the network
	resp, err := client.SubmitTransactionXDR(txe)
	if err != nil {
		log.Fatalln("submit err", err)
	}
	fmt.Println("final response", resp)

}
