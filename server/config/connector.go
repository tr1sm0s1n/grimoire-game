import (
    "log"
    
    "github.com/stellar/go/clients/horizonclient"
    "github.com/stellar/go/keypair"
    "github.com/stellar/go/network"
    "github.com/stellar/go/txnbuild"
)

// Make a keypair for a known account from a secret seed
kp, _ := keypair.Parse("SBPQUZ6G4FZNWFHKUWC5BEYWF6R52E3SEP7R3GWYSM2XTKGF5LNTWW4R")

// Get the current state of the account from the network
client := horizonclient.DefaultTestNetClient
ar := horizonclient.AccountRequest{AccountID: kp.Address()}
sourceAccount, err := client.AccountDetail(ar)
if err != nil {
    log.Fatalln(err)
}

// Build an operation to create and fund a new account
op := txnbuild.CreateAccount{
    Destination: "GCCOBXW2XQNUSL467IEILE6MMCNRR66SSVL4YQADUNYYNUVREF3FIV2Z",
    Amount:      "10",
}

// Construct the transaction that holds the operations to execute on the network
tx, err := txnbuild.NewTransaction(
    txnbuild.TransactionParams{
        SourceAccount:        &sourceAccount,
        IncrementSequenceNum: true,
        Operations:           []txnbuild.Operation{&op},
        BaseFee:              txnbuild.MinBaseFee,
        Timebounds:           txnbuild.NewTimeout(300),
    },
)
if err != nil {
    log.Fatalln(err)
)

// Sign the transaction
tx, err = tx.Sign(network.TestNetworkPassphrase, kp.(*keypair.Full))
if err != nil {
    log.Fatalln(err)
)

// Get the base 64 encoded transaction envelope
txe, err := tx.Base64()
if err != nil {
    log.Fatalln(err)
}

// Send the transaction to the network
resp, err := client.SubmitTransactionXDR(txe)
if err != nil {
    log.Fatalln(err)
}