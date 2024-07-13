// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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

// GrimoireMetaData contains all meta data concerning the Grimoire contract.
var GrimoireMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"NFTDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_userID\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_userID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_score\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50336040518060400160405280601081526020017f4752494d4f49524520506c6179657273000000000000000000000000000000008152506040518060400160405280600281526020017f4750000000000000000000000000000000000000000000000000000000000000815250815f90816200008e919062000455565b508060019081620000a0919062000455565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160362000116575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016200010d91906200057c565b60405180910390fd5b62000127816200012e60201b60201c565b5062000597565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200026d57607f821691505b60208210810362000283576200028262000228565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620002e77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002aa565b620002f38683620002aa565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200033d6200033762000331846200030b565b62000314565b6200030b565b9050919050565b5f819050919050565b62000358836200031d565b62000370620003678262000344565b848454620002b6565b825550505050565b5f90565b6200038662000378565b620003938184846200034d565b505050565b5b81811015620003ba57620003ae5f826200037c565b60018101905062000399565b5050565b601f8211156200040957620003d38162000289565b620003de846200029b565b81016020851015620003ee578190505b62000406620003fd856200029b565b83018262000398565b50505b505050565b5f82821c905092915050565b5f6200042b5f19846008026200040e565b1980831691505092915050565b5f6200044583836200041a565b9150826002028217905092915050565b6200046082620001f1565b67ffffffffffffffff8111156200047c576200047b620001fb565b5b62000488825462000255565b62000495828285620003be565b5f60209050601f831160018114620004cb575f8415620004b6578287015190505b620004c2858262000438565b86555062000531565b601f198416620004db8662000289565b5f5b828110156200050457848901518255600182019150602085019450602081019050620004dd565b8683101562000524578489015162000520601f8916826200041a565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620005648262000539565b9050919050565b620005768162000558565b82525050565b5f602082019050620005915f8301846200056b565b92915050565b6121c980620005a55f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c806370a08231116100ab578063b88d4fde1161006f578063b88d4fde1461032c578063c87b56dd14610348578063d0def52114610378578063e985e9c514610394578063f2fde38b146103c45761012a565b806370a082311461029a578063715018a6146102ca5780638da5cb5b146102d457806395d89b41146102f2578063a22cb465146103105761012a565b806342842e0e116100f257806342842e0e146101e457806342966c68146102005780634ee5c2ee1461021c5780636352211e146102385780636cccbefe146102685761012a565b806301ffc9a71461012e57806306fdde031461015e578063081812fc1461017c578063095ea7b3146101ac57806323b872dd146101c8575b5f80fd5b610148600480360381019061014391906117ed565b6103e0565b6040516101559190611832565b60405180910390f35b6101666104c1565b60405161017391906118d5565b60405180910390f35b61019660048036038101906101919190611928565b610550565b6040516101a39190611992565b60405180910390f35b6101c660048036038101906101c191906119d5565b61056b565b005b6101e260048036038101906101dd9190611a13565b610581565b005b6101fe60048036038101906101f99190611a13565b610680565b005b61021a60048036038101906102159190611928565b61069f565b005b61023660048036038101906102319190611b8f565b6106b5565b005b610252600480360381019061024d9190611928565b6106e7565b60405161025f9190611992565b60405180910390f35b610282600480360381019061027d9190611be9565b6106f8565b60405161029193929190611c3f565b60405180910390f35b6102b460048036038101906102af9190611c74565b610755565b6040516102c19190611c9f565b60405180910390f35b6102d261080b565b005b6102dc61081e565b6040516102e99190611992565b60405180910390f35b6102fa610846565b60405161030791906118d5565b60405180910390f35b61032a60048036038101906103259190611ce2565b6108d6565b005b61034660048036038101906103419190611dbe565b6108ec565b005b610362600480360381019061035d9190611928565b610909565b60405161036f91906118d5565b60405180910390f35b610392600480360381019061038d9190611e3e565b610929565b005b6103ae60048036038101906103a99190611e98565b6109fe565b6040516103bb9190611832565b60405180910390f35b6103de60048036038101906103d99190611c74565b610a8c565b005b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806104aa57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806104ba57506104b982610b10565b5b9050919050565b60605f80546104cf90611f03565b80601f01602080910402602001604051908101604052809291908181526020018280546104fb90611f03565b80156105465780601f1061051d57610100808354040283529160200191610546565b820191905f5260205f20905b81548152906001019060200180831161052957829003601f168201915b5050505050905090565b5f61055a82610b79565b5061056482610bff565b9050919050565b61057d8282610578610c38565b610c3f565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105f1575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105e89190611992565b60405180910390fd5b5f61060483836105ff610c38565b610c51565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461067a578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161067193929190611f33565b60405180910390fd5b50505050565b61069a83838360405180602001604052805f8152506108ec565b505050565b6106b15f826106ac610c38565b610c51565b5050565b6106bd610e5c565b806008836040516106ce9190611fa2565b9081526020016040518091039020600101819055505050565b5f6106f182610b79565b9050919050565b6008818051602081018201805184825260208301602085012081835280955050505050505f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107c6575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016107bd9190611992565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610813610e5c565b61081c5f610ee3565b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606001805461085590611f03565b80601f016020809104026020016040519081016040528092919081815260200182805461088190611f03565b80156108cc5780601f106108a3576101008083540402835291602001916108cc565b820191905f5260205f20905b8154815290600101906020018083116108af57829003601f168201915b5050505050905090565b6108e86108e1610c38565b8383610fa6565b5050565b6108f7848484610581565b6109038484848461110f565b50505050565b606061091482610b79565b505f61091e6112c1565b905080915050919050565b610931610e5c565b60405180606001604052808373ffffffffffffffffffffffffffffffffffffffff1681526020015f815260200160075f815461096c90611fe5565b9190508190558152506008826040516109859190611fa2565b90815260200160405180910390205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050506109fa826007546112e1565b5050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b610a94610e5c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b04575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610afb9190611992565b60405180910390fd5b610b0d81610ee3565b50565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f80610b84836112fe565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bf657826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610bed9190611c9f565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610c4c8383836001611337565b505050565b5f80610c5c846112fe565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610c9d57610c9c8184866114f6565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d2857610cdc5f855f80611337565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610da757600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b610e64610c38565b73ffffffffffffffffffffffffffffffffffffffff16610e8261081e565b73ffffffffffffffffffffffffffffffffffffffff1614610ee157610ea5610c38565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610ed89190611992565b60405180910390fd5b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361101657816040517f5b08ba1800000000000000000000000000000000000000000000000000000000815260040161100d9190611992565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516111029190611832565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b11156112bb578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02611152610c38565b8685856040518563ffffffff1660e01b8152600401611174949392919061207e565b6020604051808303815f875af19250505080156111af57506040513d601f19601f820116820180604052508101906111ac91906120dc565b60015b611230573d805f81146111dd576040519150601f19603f3d011682016040523d82523d5f602084013e6111e2565b606091505b505f81510361122857836040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161121f9190611992565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146112b957836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016112b09190611992565b60405180910390fd5b505b50505050565b60606040518060a001604052806065815260200161212f60659139905090565b6112fa828260405180602001604052805f8152506115b9565b5050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061136f57505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156114a1575f61137e84610b79565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156113e857508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156113fb57506113f981846109fe565b155b1561143d57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016114349190611992565b60405180910390fd5b811561149f57838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6115018383836115d4565b6115b4575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361157557806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161156c9190611c9f565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016115ab929190612107565b60405180910390fd5b505050565b6115c38383611694565b6115cf5f84848461110f565b505050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561168b57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061164c575061164b84846109fe565b5b8061168a57508273ffffffffffffffffffffffffffffffffffffffff1661167283610bff565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611704575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016116fb9190611992565b60405180910390fd5b5f61171083835f610c51565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611782575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016117799190611992565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6117cc81611798565b81146117d6575f80fd5b50565b5f813590506117e7816117c3565b92915050565b5f6020828403121561180257611801611790565b5b5f61180f848285016117d9565b91505092915050565b5f8115159050919050565b61182c81611818565b82525050565b5f6020820190506118455f830184611823565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015611882578082015181840152602081019050611867565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6118a78261184b565b6118b18185611855565b93506118c1818560208601611865565b6118ca8161188d565b840191505092915050565b5f6020820190508181035f8301526118ed818461189d565b905092915050565b5f819050919050565b611907816118f5565b8114611911575f80fd5b50565b5f81359050611922816118fe565b92915050565b5f6020828403121561193d5761193c611790565b5b5f61194a84828501611914565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61197c82611953565b9050919050565b61198c81611972565b82525050565b5f6020820190506119a55f830184611983565b92915050565b6119b481611972565b81146119be575f80fd5b50565b5f813590506119cf816119ab565b92915050565b5f80604083850312156119eb576119ea611790565b5b5f6119f8858286016119c1565b9250506020611a0985828601611914565b9150509250929050565b5f805f60608486031215611a2a57611a29611790565b5b5f611a37868287016119c1565b9350506020611a48868287016119c1565b9250506040611a5986828701611914565b9150509250925092565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611aa18261188d565b810181811067ffffffffffffffff82111715611ac057611abf611a6b565b5b80604052505050565b5f611ad2611787565b9050611ade8282611a98565b919050565b5f67ffffffffffffffff821115611afd57611afc611a6b565b5b611b068261188d565b9050602081019050919050565b828183375f83830152505050565b5f611b33611b2e84611ae3565b611ac9565b905082815260208101848484011115611b4f57611b4e611a67565b5b611b5a848285611b13565b509392505050565b5f82601f830112611b7657611b75611a63565b5b8135611b86848260208601611b21565b91505092915050565b5f8060408385031215611ba557611ba4611790565b5b5f83013567ffffffffffffffff811115611bc257611bc1611794565b5b611bce85828601611b62565b9250506020611bdf85828601611914565b9150509250929050565b5f60208284031215611bfe57611bfd611790565b5b5f82013567ffffffffffffffff811115611c1b57611c1a611794565b5b611c2784828501611b62565b91505092915050565b611c39816118f5565b82525050565b5f606082019050611c525f830186611983565b611c5f6020830185611c30565b611c6c6040830184611c30565b949350505050565b5f60208284031215611c8957611c88611790565b5b5f611c96848285016119c1565b91505092915050565b5f602082019050611cb25f830184611c30565b92915050565b611cc181611818565b8114611ccb575f80fd5b50565b5f81359050611cdc81611cb8565b92915050565b5f8060408385031215611cf857611cf7611790565b5b5f611d05858286016119c1565b9250506020611d1685828601611cce565b9150509250929050565b5f67ffffffffffffffff821115611d3a57611d39611a6b565b5b611d438261188d565b9050602081019050919050565b5f611d62611d5d84611d20565b611ac9565b905082815260208101848484011115611d7e57611d7d611a67565b5b611d89848285611b13565b509392505050565b5f82601f830112611da557611da4611a63565b5b8135611db5848260208601611d50565b91505092915050565b5f805f8060808587031215611dd657611dd5611790565b5b5f611de3878288016119c1565b9450506020611df4878288016119c1565b9350506040611e0587828801611914565b925050606085013567ffffffffffffffff811115611e2657611e25611794565b5b611e3287828801611d91565b91505092959194509250565b5f8060408385031215611e5457611e53611790565b5b5f611e61858286016119c1565b925050602083013567ffffffffffffffff811115611e8257611e81611794565b5b611e8e85828601611b62565b9150509250929050565b5f8060408385031215611eae57611ead611790565b5b5f611ebb858286016119c1565b9250506020611ecc858286016119c1565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611f1a57607f821691505b602082108103611f2d57611f2c611ed6565b5b50919050565b5f606082019050611f465f830186611983565b611f536020830185611c30565b611f606040830184611983565b949350505050565b5f81905092915050565b5f611f7c8261184b565b611f868185611f68565b9350611f96818560208601611865565b80840191505092915050565b5f611fad8284611f72565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fef826118f5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361202157612020611fb8565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f6120508261202c565b61205a8185612036565b935061206a818560208601611865565b6120738161188d565b840191505092915050565b5f6080820190506120915f830187611983565b61209e6020830186611983565b6120ab6040830185611c30565b81810360608301526120bd8184612046565b905095945050505050565b5f815190506120d6816117c3565b92915050565b5f602082840312156120f1576120f0611790565b5b5f6120fe848285016120c8565b91505092915050565b5f60408201905061211a5f830185611983565b6121276020830184611c30565b939250505056fe68747470733a2f2f77686974652d70726f76696e6369616c2d6269736f6e2d3733372e6d7970696e6174612e636c6f75642f697066732f516d5964753666784b6e65636e6d774c6e674533527264777546704d57334c59464a6e6837434147535345465739a2646970667358221220e42a0e7e1347ea697747f60fbe819ebb11357ac62e901b7fd948beec990d0c1064736f6c63430008140033",
}

// GrimoireABI is the input ABI used to generate the binding from.
// Deprecated: Use GrimoireMetaData.ABI instead.
var GrimoireABI = GrimoireMetaData.ABI

// GrimoireBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GrimoireMetaData.Bin instead.
var GrimoireBin = GrimoireMetaData.Bin

// DeployGrimoire deploys a new Ethereum contract, binding an instance of Grimoire to it.
func DeployGrimoire(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Grimoire, error) {
	parsed, err := GrimoireMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GrimoireBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Grimoire{GrimoireCaller: GrimoireCaller{contract: contract}, GrimoireTransactor: GrimoireTransactor{contract: contract}, GrimoireFilterer: GrimoireFilterer{contract: contract}}, nil
}

// Grimoire is an auto generated Go binding around an Ethereum contract.
type Grimoire struct {
	GrimoireCaller     // Read-only binding to the contract
	GrimoireTransactor // Write-only binding to the contract
	GrimoireFilterer   // Log filterer for contract events
}

// GrimoireCaller is an auto generated read-only Go binding around an Ethereum contract.
type GrimoireCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrimoireTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GrimoireTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrimoireFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GrimoireFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrimoireSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GrimoireSession struct {
	Contract     *Grimoire         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GrimoireCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GrimoireCallerSession struct {
	Contract *GrimoireCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GrimoireTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GrimoireTransactorSession struct {
	Contract     *GrimoireTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GrimoireRaw is an auto generated low-level Go binding around an Ethereum contract.
type GrimoireRaw struct {
	Contract *Grimoire // Generic contract binding to access the raw methods on
}

// GrimoireCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GrimoireCallerRaw struct {
	Contract *GrimoireCaller // Generic read-only contract binding to access the raw methods on
}

// GrimoireTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GrimoireTransactorRaw struct {
	Contract *GrimoireTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGrimoire creates a new instance of Grimoire, bound to a specific deployed contract.
func NewGrimoire(address common.Address, backend bind.ContractBackend) (*Grimoire, error) {
	contract, err := bindGrimoire(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Grimoire{GrimoireCaller: GrimoireCaller{contract: contract}, GrimoireTransactor: GrimoireTransactor{contract: contract}, GrimoireFilterer: GrimoireFilterer{contract: contract}}, nil
}

// NewGrimoireCaller creates a new read-only instance of Grimoire, bound to a specific deployed contract.
func NewGrimoireCaller(address common.Address, caller bind.ContractCaller) (*GrimoireCaller, error) {
	contract, err := bindGrimoire(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GrimoireCaller{contract: contract}, nil
}

// NewGrimoireTransactor creates a new write-only instance of Grimoire, bound to a specific deployed contract.
func NewGrimoireTransactor(address common.Address, transactor bind.ContractTransactor) (*GrimoireTransactor, error) {
	contract, err := bindGrimoire(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GrimoireTransactor{contract: contract}, nil
}

// NewGrimoireFilterer creates a new log filterer instance of Grimoire, bound to a specific deployed contract.
func NewGrimoireFilterer(address common.Address, filterer bind.ContractFilterer) (*GrimoireFilterer, error) {
	contract, err := bindGrimoire(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GrimoireFilterer{contract: contract}, nil
}

// bindGrimoire binds a generic wrapper to an already deployed contract.
func bindGrimoire(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GrimoireMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Grimoire *GrimoireRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Grimoire.Contract.GrimoireCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Grimoire *GrimoireRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grimoire.Contract.GrimoireTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Grimoire *GrimoireRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Grimoire.Contract.GrimoireTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Grimoire *GrimoireCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Grimoire.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Grimoire *GrimoireTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grimoire.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Grimoire *GrimoireTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Grimoire.Contract.contract.Transact(opts, method, params...)
}

// NFTDetails is a free data retrieval call binding the contract method 0x6cccbefe.
//
// Solidity: function NFTDetails(string ) view returns(address player, uint256 score, uint256 id)
func (_Grimoire *GrimoireCaller) NFTDetails(opts *bind.CallOpts, arg0 string) (struct {
	Player common.Address
	Score  *big.Int
	Id     *big.Int
}, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "NFTDetails", arg0)

	outstruct := new(struct {
		Player common.Address
		Score  *big.Int
		Id     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Player = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Score = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NFTDetails is a free data retrieval call binding the contract method 0x6cccbefe.
//
// Solidity: function NFTDetails(string ) view returns(address player, uint256 score, uint256 id)
func (_Grimoire *GrimoireSession) NFTDetails(arg0 string) (struct {
	Player common.Address
	Score  *big.Int
	Id     *big.Int
}, error) {
	return _Grimoire.Contract.NFTDetails(&_Grimoire.CallOpts, arg0)
}

// NFTDetails is a free data retrieval call binding the contract method 0x6cccbefe.
//
// Solidity: function NFTDetails(string ) view returns(address player, uint256 score, uint256 id)
func (_Grimoire *GrimoireCallerSession) NFTDetails(arg0 string) (struct {
	Player common.Address
	Score  *big.Int
	Id     *big.Int
}, error) {
	return _Grimoire.Contract.NFTDetails(&_Grimoire.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Grimoire *GrimoireCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Grimoire *GrimoireSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Grimoire.Contract.BalanceOf(&_Grimoire.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Grimoire *GrimoireCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Grimoire.Contract.BalanceOf(&_Grimoire.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Grimoire *GrimoireCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Grimoire *GrimoireSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Grimoire.Contract.GetApproved(&_Grimoire.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Grimoire *GrimoireCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Grimoire.Contract.GetApproved(&_Grimoire.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Grimoire *GrimoireCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Grimoire *GrimoireSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Grimoire.Contract.IsApprovedForAll(&_Grimoire.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Grimoire *GrimoireCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Grimoire.Contract.IsApprovedForAll(&_Grimoire.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Grimoire *GrimoireCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Grimoire *GrimoireSession) Name() (string, error) {
	return _Grimoire.Contract.Name(&_Grimoire.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Grimoire *GrimoireCallerSession) Name() (string, error) {
	return _Grimoire.Contract.Name(&_Grimoire.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Grimoire *GrimoireCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Grimoire *GrimoireSession) Owner() (common.Address, error) {
	return _Grimoire.Contract.Owner(&_Grimoire.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Grimoire *GrimoireCallerSession) Owner() (common.Address, error) {
	return _Grimoire.Contract.Owner(&_Grimoire.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Grimoire *GrimoireCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Grimoire *GrimoireSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Grimoire.Contract.OwnerOf(&_Grimoire.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Grimoire *GrimoireCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Grimoire.Contract.OwnerOf(&_Grimoire.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Grimoire *GrimoireCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Grimoire *GrimoireSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Grimoire.Contract.SupportsInterface(&_Grimoire.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Grimoire *GrimoireCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Grimoire.Contract.SupportsInterface(&_Grimoire.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Grimoire *GrimoireCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Grimoire *GrimoireSession) Symbol() (string, error) {
	return _Grimoire.Contract.Symbol(&_Grimoire.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Grimoire *GrimoireCallerSession) Symbol() (string, error) {
	return _Grimoire.Contract.Symbol(&_Grimoire.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Grimoire *GrimoireCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Grimoire.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Grimoire *GrimoireSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Grimoire.Contract.TokenURI(&_Grimoire.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Grimoire *GrimoireCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Grimoire.Contract.TokenURI(&_Grimoire.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.Approve(&_Grimoire.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.Approve(&_Grimoire.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Grimoire *GrimoireSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.Burn(&_Grimoire.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.Burn(&_Grimoire.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _userID) returns()
func (_Grimoire *GrimoireTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _userID string) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "mint", _to, _userID)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _userID) returns()
func (_Grimoire *GrimoireSession) Mint(_to common.Address, _userID string) (*types.Transaction, error) {
	return _Grimoire.Contract.Mint(&_Grimoire.TransactOpts, _to, _userID)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _userID) returns()
func (_Grimoire *GrimoireTransactorSession) Mint(_to common.Address, _userID string) (*types.Transaction, error) {
	return _Grimoire.Contract.Mint(&_Grimoire.TransactOpts, _to, _userID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Grimoire *GrimoireTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Grimoire *GrimoireSession) RenounceOwnership() (*types.Transaction, error) {
	return _Grimoire.Contract.RenounceOwnership(&_Grimoire.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Grimoire *GrimoireTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Grimoire.Contract.RenounceOwnership(&_Grimoire.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.SafeTransferFrom(&_Grimoire.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.SafeTransferFrom(&_Grimoire.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Grimoire *GrimoireTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Grimoire *GrimoireSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Grimoire.Contract.SafeTransferFrom0(&_Grimoire.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Grimoire *GrimoireTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Grimoire.Contract.SafeTransferFrom0(&_Grimoire.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Grimoire *GrimoireTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Grimoire *GrimoireSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Grimoire.Contract.SetApprovalForAll(&_Grimoire.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Grimoire *GrimoireTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Grimoire.Contract.SetApprovalForAll(&_Grimoire.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.TransferFrom(&_Grimoire.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Grimoire *GrimoireTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.TransferFrom(&_Grimoire.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Grimoire *GrimoireTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Grimoire *GrimoireSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Grimoire.Contract.TransferOwnership(&_Grimoire.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Grimoire *GrimoireTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Grimoire.Contract.TransferOwnership(&_Grimoire.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x4ee5c2ee.
//
// Solidity: function update(string _userID, uint256 _score) returns()
func (_Grimoire *GrimoireTransactor) Update(opts *bind.TransactOpts, _userID string, _score *big.Int) (*types.Transaction, error) {
	return _Grimoire.contract.Transact(opts, "update", _userID, _score)
}

// Update is a paid mutator transaction binding the contract method 0x4ee5c2ee.
//
// Solidity: function update(string _userID, uint256 _score) returns()
func (_Grimoire *GrimoireSession) Update(_userID string, _score *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.Update(&_Grimoire.TransactOpts, _userID, _score)
}

// Update is a paid mutator transaction binding the contract method 0x4ee5c2ee.
//
// Solidity: function update(string _userID, uint256 _score) returns()
func (_Grimoire *GrimoireTransactorSession) Update(_userID string, _score *big.Int) (*types.Transaction, error) {
	return _Grimoire.Contract.Update(&_Grimoire.TransactOpts, _userID, _score)
}

// GrimoireApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Grimoire contract.
type GrimoireApprovalIterator struct {
	Event *GrimoireApproval // Event containing the contract specifics and raw log

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
func (it *GrimoireApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrimoireApproval)
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
		it.Event = new(GrimoireApproval)
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
func (it *GrimoireApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrimoireApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrimoireApproval represents a Approval event raised by the Grimoire contract.
type GrimoireApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Grimoire *GrimoireFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GrimoireApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Grimoire.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GrimoireApprovalIterator{contract: _Grimoire.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Grimoire *GrimoireFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GrimoireApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Grimoire.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrimoireApproval)
				if err := _Grimoire.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Grimoire *GrimoireFilterer) ParseApproval(log types.Log) (*GrimoireApproval, error) {
	event := new(GrimoireApproval)
	if err := _Grimoire.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GrimoireApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Grimoire contract.
type GrimoireApprovalForAllIterator struct {
	Event *GrimoireApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GrimoireApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrimoireApprovalForAll)
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
		it.Event = new(GrimoireApprovalForAll)
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
func (it *GrimoireApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrimoireApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrimoireApprovalForAll represents a ApprovalForAll event raised by the Grimoire contract.
type GrimoireApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Grimoire *GrimoireFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GrimoireApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Grimoire.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GrimoireApprovalForAllIterator{contract: _Grimoire.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Grimoire *GrimoireFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GrimoireApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Grimoire.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrimoireApprovalForAll)
				if err := _Grimoire.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Grimoire *GrimoireFilterer) ParseApprovalForAll(log types.Log) (*GrimoireApprovalForAll, error) {
	event := new(GrimoireApprovalForAll)
	if err := _Grimoire.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GrimoireOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Grimoire contract.
type GrimoireOwnershipTransferredIterator struct {
	Event *GrimoireOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GrimoireOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrimoireOwnershipTransferred)
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
		it.Event = new(GrimoireOwnershipTransferred)
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
func (it *GrimoireOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrimoireOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrimoireOwnershipTransferred represents a OwnershipTransferred event raised by the Grimoire contract.
type GrimoireOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Grimoire *GrimoireFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GrimoireOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Grimoire.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GrimoireOwnershipTransferredIterator{contract: _Grimoire.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Grimoire *GrimoireFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GrimoireOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Grimoire.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrimoireOwnershipTransferred)
				if err := _Grimoire.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Grimoire *GrimoireFilterer) ParseOwnershipTransferred(log types.Log) (*GrimoireOwnershipTransferred, error) {
	event := new(GrimoireOwnershipTransferred)
	if err := _Grimoire.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GrimoireTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Grimoire contract.
type GrimoireTransferIterator struct {
	Event *GrimoireTransfer // Event containing the contract specifics and raw log

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
func (it *GrimoireTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrimoireTransfer)
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
		it.Event = new(GrimoireTransfer)
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
func (it *GrimoireTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrimoireTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrimoireTransfer represents a Transfer event raised by the Grimoire contract.
type GrimoireTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Grimoire *GrimoireFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GrimoireTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Grimoire.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GrimoireTransferIterator{contract: _Grimoire.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Grimoire *GrimoireFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GrimoireTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Grimoire.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrimoireTransfer)
				if err := _Grimoire.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Grimoire *GrimoireFilterer) ParseTransfer(log types.Log) (*GrimoireTransfer, error) {
	event := new(GrimoireTransfer)
	if err := _Grimoire.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
