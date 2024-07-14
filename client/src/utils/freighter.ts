import { signTransaction } from '@stellar/freighter-api'

export const userSignTransaction = async (
  xdr: string,
  network: string,
  signWith: string
) => {
  let signedTransaction = ''
  let error = ''

  try {
    signedTransaction = await signTransaction(xdr, {
      network,
      accountToSign: signWith,
    })    
  } catch (e: any) {
    error = e
  }

  if (error) {
    return error
  }

  return signedTransaction
}
