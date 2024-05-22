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

    function wrap(Token memory token, address assetContract) external onlyOwner {
        console.log("Wrap token: %s", assetContract);
        if(tokens[assetContract].assetContract != address(0)) {
            console.log("Token Already Wrapped");
            revert TokenAlreadyWrapped(assetContract);
        }
        tokens[assetContract] = token;

        ids.push(token.tokenId);
    }

    function wrap2(address assetContract, TokenType tokenType, string memory tokenId, address owner) external onlyOwner {
        console.log("Wrap token: %s", assetContract);
        if(tokens[assetContract].assetContract != address(0)) {
            console.log("Token Already Wrapped");
            revert TokenAlreadyWrapped(assetContract);
        }
        tokens[assetContract] = Token(assetContract, type, tokenId, owner, 0);

        ids.push(tokenId);
    }

    function unwrap(address assetContract) external onlyOwner {
        console.log("Unwrap token: %s", assetContract);
        deleteFromArray(tokens[assetContract].tokenId);
        console.log("Deleted from array token: %s", assetContract);
        delete tokens[assetContract];
    }

    function lock(address assetContract, uint256 amount) external onlyOwner returns (bool success) {
        console.log("Lock token: %s\n amount: %s", assetContract, amount);
        if(tokens[assetContract].assetContract == address(0)){
            revert TokenNotAvailable(assetContract);
        }

        (bool success,) = assetContract.call(abi.encodeWithSignature("transfer(address,address,uint256)", tokens[assetContract].owner, address(this), amount));

        if(success) {
            tokens[assetContract].amount += amount; 
            return true;
        }

        revert TokenNotLocked(assetContract);
    } 

    function unlock(address assetContract, uint256 amount) external onlyOwner returns (bool success) { //ammount
        console.log("Unlock token: %s\n amount: %s", assetContract, amount);
        if(tokens[assetContract].assetContract == address(0)){
            revert TokenNotAvailable(assetContract);
        }

        if(tokens[assetContract].amount < amount) {
            revert InsuficientAmountLocked(assetContract, amount);
        }

        (bool successAprov,) = assetContract.call(abi.encodeWithSignature("approve(address,uint256)", address(this), amount));
        if(!successAprov) {
            revert TokenNotUnlocked(assetContract);
        }

        (bool success,) = assetContract.call(abi.encodeWithSignature("transfer(address,address,uint256)", address(this), tokens[assetContract].owner, tokens[assetContract].amount));
        if(success) {
            tokens[assetContract].amount -= amount;
            return true;
        }

        revert TokenNotUnlocked(assetContract);
    } 
    
    function getAllAssetsIDs() external view returns (string[] memory) {
        return ids;
    }

    function mint(address assetContract, uint256 amount) external onlyOwner {
        console.log("Wrapper Mint to: %s\n amount: %s", assetContract, amount);

        if(tokens[assetContract].assetContract == address(0)){
            revert TokenNotAvailable(assetContract);
        }

        (bool success, ) = assetContract.call(abi.encodeWithSignature("mint(address,uint256)", address(this), amount));
        
        require(success, "mint asset call failed");

        tokens[assetContract].amount = amount; 
    }

    function burn(address assetContract, uint256 amount) external onlyOwner {
        console.log("Wrapper Burn from: %s\n amount: %s", assetContract, amount);

        require(tokens[assetContract].amount >= amount, "burn asset asset is not locked");

        (bool success, ) = assetContract.call(abi.encodeWithSignature("burn(address,uint256)",  address(this), amount));

        require(success, "burn asset call failed");

        tokens[assetContract].amount -= amount; 
    }

    function assign(address assetContract, address receiver_account, uint256 amount) external onlyOwner {
        console.log("Wrapper Assign from: %s\n to: %s\n amount: %s", assetContract, receiver_account, amount);

        require(tokens[assetContract].amount >= amount, "assign asset asset is not locked");

        (bool success, ) = assetContract.call(abi.encodeWithSignature("assign(address,address,uint256)", address(this), receiver_account, amount));

        require(success, "assign asset call failed");
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

    function getToken(address assetContract) view public returns (Token memory token) {
        return tokens[assetContract];
    }
}