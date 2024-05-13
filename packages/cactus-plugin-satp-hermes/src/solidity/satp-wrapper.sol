// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.15;

import "@openzeppelin/contracts/access/Ownable.sol";

enum TokenType { ERC20, ERC721, ERC1155 }
 
struct Token {
    address assetContract;
    TokenType tokenType;
    uint256 tokenId;
    uint256 totalAmount;
}

contract SATPWrapperContract is Ownable, ITraceableContract{

    function wrap(
        Token[] memory tokensToWrap,
        string calldata uriForWrappedToken,
        address recipient
    ) external payable returns (uint256 tokenId);

    function unwrap(
        uint256 tokenId,
        address recipient
    ) external;

   
}