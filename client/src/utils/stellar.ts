import {
  Contract,
  SorobanRpc,
  TransactionBuilder,
  Networks,
  BASE_FEE,
  nativeToScVal,
  Address,
  xdr,
} from '@stellar/stellar-sdk'
import { userSignTransaction } from './freighter'

let rpcUrl = 'https://soroban-testnet.stellar.org'

let contractAddress = 'CABKK7CKUV6LSSCREM2TABPFKLA5KCQLGBTYL7WAUBPNFTHFHGO57DUC'

// coverting Account Address to ScVal form
const accountToScVal = (account: string) => new Address(account).toScVal()

// coverting String to ScVal form
const stringToScValString = (value: string) => {
  return nativeToScVal(value)
}

const numberToU64 = (value: number) => {
  return nativeToScVal(value, { type: 'u64' })
}

let params = {
  fee: BASE_FEE,
  networkPassphrase: Networks.TESTNET,
}

async function contractInt(
  caller: string,
  functName: string,
  values: xdr.ScVal[] | xdr.ScVal
) {
  const provider = new SorobanRpc.Server(rpcUrl, { allowHttp: true })
  const sourceAccount = await provider.getAccount(caller)
  const contract = new Contract(contractAddress)
  let buildTx

  if (values == null) {
    buildTx = new TransactionBuilder(sourceAccount, params)
      .addOperation(contract.call(functName))
      .setTimeout(30)
      .build()
  } else if (Array.isArray(values)) {
    buildTx = new TransactionBuilder(sourceAccount, params)
      .addOperation(contract.call(functName, ...values))
      .setTimeout(30)
      .build()
  } else {
    buildTx = new TransactionBuilder(sourceAccount, params)
      .addOperation(contract.call(functName, values))
      .setTimeout(30)
      .build()
  }

  let _buildTx = await provider.prepareTransaction(buildTx)

  let prepareTx = _buildTx.toXDR() // pre-encoding (converting it to XDR format)

  let signedTx = await userSignTransaction(prepareTx, 'TESTNET', caller)
  
  let tx = TransactionBuilder.fromXDR(signedTx, Networks.TESTNET)

  try {
    let sendTx = await provider.sendTransaction(tx).catch(function (err) {
      console.error('Catch-1', err)
      return err
    })
    if (sendTx.errorResult) {
      throw new Error('Unable to submit transaction')
    }
    if (sendTx.status === 'PENDING') {
      let txResponse = await provider.getTransaction(sendTx.hash)
      //   we will continously checking the transaction status until it gets successfull added to the blockchain ledger or it gets rejected
      while (txResponse.status === 'NOT_FOUND') {
        txResponse = await provider.getTransaction(sendTx.hash)
        await new Promise((resolve) => setTimeout(resolve, 100))
      }
      if (txResponse.status === 'SUCCESS') {
        let result = txResponse.returnValue
        return result
      }
    }
  } catch (err) {
    console.log('Catch-2', err)
    return
  }
}

async function mint(caller: string, id: string, addr: string, score: number) {
  let idScVal = stringToScValString(id)
  let addrScVal = accountToScVal(addr)
  let scoreScVal = numberToU64(score)
  let values = [idScVal, addrScVal, scoreScVal]

  try {
    const res = await contractInt(caller, 'mint', values)
    console.log(res)
    return res
  } catch (error) {
    console.log(error)
    throw new Error("error: cannot mint");
  }
}

async function getNftDetails(caller: string, id: string) {
  let values = stringToScValString(id)

  try {
    let result = await contractInt(caller, 'get_nft_details', values)
    console.log(result)
    return result
  } catch (error) {
    console.log('Unable to fetch Your Pass Status!!')
  }
}

export { mint, getNftDetails }
