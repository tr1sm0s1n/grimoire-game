// SPDX-License-Identifier: MIT
pragma solidity 0.8.0;

import "./contracts/token/ERC721/ERC721.sol";
import "./contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "./contracts/access/Ownable.sol";

contract Grimoire is ERC721, ERC721Burnable, Ownable {
    constructor() ERC721("GRIMOIRE Players", "GP") Ownable(msg.sender) {}

    uint256 count;

    struct NFT {
        address player;
        uint256 score;
        uint256 id;
    }

    mapping(string => NFT) public NFTDetails;
    mapping(string => uint) public BetAmount;
    mapping(string => uint) public ClaimAmount;

    function _baseURI() internal pure override returns (string memory) {
        return
            "https://white-provincial-bison-737.mypinata.cloud/ipfs/QmYdu6fxKnecnmwLngE3RrdwuFpMW3LYFJnh7CAGSSEFW9";
    }

    function mint(address _to, string memory _userID) public onlyOwner {
        NFTDetails[_userID] = NFT(_to, 0, ++count);
        _safeMint(_to, count);
    }

    function update(string memory _userID, uint256 _score) public onlyOwner {
        NFTDetails[_userID].score = _score;
    }

    function setBet(string memory _userID, uint _amount) public {
        BetAmount[_userID] = _amount;
    }

    function closeBet(string memory _userID, uint _amount) public {
        ClaimAmount[_userID] = _amount;
    }
}
