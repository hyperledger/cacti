// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.15;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@hyperledger/cactus-plugin-bungee-hermes/src/test/typescript/solidity/lock-asset-contract/ITraceableContract.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

enum TokenType { ERC20, ERC721, ERC1155 }
 
struct Token {
    address assetContract;
    TokenType tokenType;
    string tokenId;
    bool locked;
}

error TokenNotAvailable(address token);

error TokenAlreadyWrapped(address token);

contract SATPWrapperContract is Ownable, ITraceableContract{

    mapping (address  => Token) public tokens; //contract address to Token

    string[] ids;

    address public bridge_address;

    constructor(address _bridge_address)  Ownable(_bridge_address) {
        bridge_address = address(_bridge_address);
    }

    function wrap(Token memory token, address assetContract) external onlyOwner {

        if(tokens[assetContract].assetContract != address(0)) {
            revert TokenAlreadyWrapped(assetContract);
        }
        tokens[assetContract] = token;
    }

    function unwrap(address assetContract) external onlyOwner {
        delete tokens[assetContract];
    }

    function lock(address tokenAddress) external onlyOwner returns (bool success) {
        if(tokens[tokenAddress].assetContract == address(0)){
            revert TokenNotAvailable(tokenAddress);
        }

        tokens[tokenAddress].locked = true;
        return true;
    } 

    function unlock(address tokenAddress) external onlyOwner returns (bool success) {
        if(tokens[tokenAddress].assetContract == address(0)){
            revert TokenNotAvailable(tokenAddress);
        }

        tokens[tokenAddress].locked = false;
        return true;
    } 
    
    function getAllAssetsIDs() external view returns (string[] memory) {
        return ids;
    }

    function mint(address assetContract, uint256 amount) external onlyOwner {
        require(tokens[assetContract].locked, "mint asset asset is not locked");
        bytes memory encodedFunctionCall = abi.encode("mint(uint256)",  amount);
        (bool success,) = assetContract.call(encodedFunctionCall);

        require(success, "mint asset call failed");
    }

    function burn(address assetContract, uint256 amount) external onlyOwner {
        require(tokens[assetContract].locked, "burn asset asset is not locked");
        bytes memory encodedFunctionCall = abi.encode("burn(uint256)",  amount);
        (bool success,) = assetContract.call(encodedFunctionCall);

        require(success, "burn asset call failed");
    }

    function assign(address assetContract, address receiver_account, uint256 amount) external onlyOwner {
        require(tokens[assetContract].locked, "assign asset asset is not locked");

        bytes memory encodedFunctionCall = abi.encode("assign(receiver_account,uint256)", receiver_account, amount);
        (bool success,) = assetContract.call(encodedFunctionCall);

        require(success, "assign asset call failed");

        this.unwrap(assetContract);
    }
   
}