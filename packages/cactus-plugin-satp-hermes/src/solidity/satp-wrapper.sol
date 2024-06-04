// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.15;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@hyperledger/cactus-plugin-bungee-hermes/src/test/typescript/solidity/lock-asset-contract/ITraceableContract.sol";
import "./satp-erc20.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

import "hardhat/console.sol";

enum TokenType { ERC20, ERC721, ERC1155 }
 
struct Token {
    address assetContract;
    TokenType tokenType;
    string tokenId;
    address owner;
    uint amount; //ammount that is approved by the contract owner 
}

error TokenNotAvailable(address token);

error TokenAlreadyWrapped(address token);

error TokenNotLocked(address token);

error TokenNotUnlocked(address token);

error InsuficientAmountLocked(address token, uint256 amount);

contract SATPWrapperContract is Ownable, ITraceableContract{

    mapping (address  => Token) public tokens; //contract address to Token

    string[] ids;

    address public bridge_address;

    constructor(address _bridge_address)  Ownable(_bridge_address) {
        bridge_address = address(_bridge_address);
    }

    function wrap(address assetContract, TokenType tokenType, string memory tokenId, address owner) external onlyOwner returns (bool success) {
        console.log("Wrap token: %s", assetContract);
        if(tokens[assetContract].assetContract != address(0)) {
            console.log("Token Already Wrapped");
            revert TokenAlreadyWrapped(assetContract);
        }
        tokens[tokenId] = Token(assetContract, type, tokenId, owner, 0);

        ids.push(tokenId);
        return true;
    }

    function unwrap(string tokenId) external onlyOwner returns (bool success) {
        console.log("Unwrap token: %s", tokenId);
        deleteFromArray(tokens[tokenId].tokenId);
        console.log("Deleted from array token: %s", tokenId);
        delete tokens[tokenId];

        return true;
    }

    function lock(string tokenId, uint256 amount) external onlyOwner returns (bool success) {
        console.log("Lock token: %s\n amount: %s", tokenId, amount);
        if(tokens[tokenId].assetContract == address(0)){
            revert TokenNotAvailable(tokenId);
        }

        (bool success,) = tokens[tokenId].assetContract.call(abi.encodeWithSignature("transfer(address,address,uint256)", tokens[tokenId].owner, address(this), amount));

        if(success) {
            tokens[tokenId].amount += amount; 
            return true;
        }

        revert TokenNotLocked(tokenId);
    } 

    function unlock(string tokenId, uint256 amount) external onlyOwner returns (bool success) { //ammount
        console.log("Unlock token: %s\n amount: %s", tokenId, amount);
        if(tokens[tokenId].assetContract == address(0)){
            revert TokenNotAvailable(tokenId);
        }

        if(tokens[tokenId].amount < amount) {
            revert InsuficientAmountLocked(tokenId, amount);
        }

        (bool successAprov,) = tokens[tokenId].assetContract.call(abi.encodeWithSignature("approve(address,uint256)", address(this), amount));
        if(!successAprov) {
            revert TokenNotUnlocked(tokenId);
        }

        (bool success,) = tokens[tokenId].assetContract.call(abi.encodeWithSignature("transfer(address,address,uint256)", address(this), tokens[tokenId].owner, tokens[tokenId].amount));
        if(success) {
            tokens[tokenId].amount -= amount;
            return true;
        }

        revert TokenNotUnlocked(tokenId);
    } 
    
    function getAllAssetsIDs() external view returns (string[] memory) {
        return ids;
    }

    function mint(string tokenId, uint256 amount) external onlyOwner returns (bool success) {
        console.log("Wrapper Mint to: %s\n amount: %s", assetContract, amount);

        if(tokens[tokenId].assetContract == address(0)){
            revert TokenNotAvailable(tokenId);
        }

        (bool success, ) = tokens[tokenId].assetContract.call(abi.encodeWithSignature("mint(address,uint256)", address(this), amount));
        
        require(success, "mint asset call failed");

        tokens[tokenId].amount = amount; 

        return true;
    }

    function burn(string tokenId, uint256 amount) external onlyOwner returns (bool success) {
        console.log("Wrapper Burn from: %s\n amount: %s", tokenId, amount);

        require(tokens[tokenId].amount >= amount, "burn asset asset is not locked");

        (bool success, ) = tokens[tokenId].assetContract.call(abi.encodeWithSignature("burn(address,uint256)",  address(this), amount));

        require(success, "burn asset call failed");

        tokens[tokenId].amount -= amount; 

        return true;
    }

    function assign(string tokenId, address receiver_account, uint256 amount) external onlyOwner returns (bool success) {
        console.log("Wrapper Assign from: %s\n to: %s\n amount: %s", tokenId, receiver_account, amount);

        require(tokens[tokenId].amount >= amount, "assign asset asset is not locked");

        (bool success, ) = tokens[tokenId].assetContract.call(abi.encodeWithSignature("assign(address,address,uint256)", address(this), receiver_account, amount));

        require(success, "assign asset call failed");

        return true;
    }   

    function deleteFromArray(string memory id) internal  {
        for (uint256 i = 0; i < ids.length; i++) {
            if (Strings.equal(ids[i], id)) {
                ids[i] = ids[ids.length - 1];
                ids.pop();
                break;
            }
        }
    }

    function getToken(string tokenId) view public returns (Token memory token) {
        return tokens[tokenId];
    }
}