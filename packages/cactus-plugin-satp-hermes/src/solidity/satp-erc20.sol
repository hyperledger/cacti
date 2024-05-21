// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.15;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@hyperledger/cactus-plugin-bungee-hermes/src/test/typescript/solidity/lock-asset-contract/ITraceableContract.sol";

import "hardhat/console.sol";


contract SATPContract is AccessControl, ERC20, ITraceableContract{

    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");
    bytes32 public constant OWNER_ROLE = keccak256("OWNER_ROLE");

    string public id;

    mapping (string => uint256) public balanceOf;

    constructor(address _owner, string memory _id) ERC20("SATPToken", "SATP") {
        _grantRole(OWNER_ROLE, _owner);
        _grantRole(BRIDGE_ROLE, _owner);

        id = _id;
    }

    function mint(address account, uint256 amount) external onlyRole(BRIDGE_ROLE) {
        console.log("Mint to: %s\n amount: %s", account, amount);
        _mint(account, amount);
    }

    function burn(address account, uint256 amount) external onlyRole(BRIDGE_ROLE) {
        console.log("Burn from: %s\n amount: %s", account, amount);
        _burn(account, amount);
    }

    function assign(address from, address recipient, uint256 amount) external onlyRole(BRIDGE_ROLE) {
        console.log("Assing from: %s\n to: %s \n amount: %s", from, recipient, amount);
        require(from == _msgSender(), "The msgSender is not the owner");
        _transfer(from, recipient, amount);
    }

    function transfer(address from, address recipient, uint256 amount) external onlyRole(BRIDGE_ROLE) {
        console.log("transfer from: %s\n to: %s \n amount: %s", from, recipient, amount);
        transferFrom(from, recipient, amount);
    }

    function getAllAssetsIDs() external view returns (string[] memory) {
        string[] memory myArray = new string[](1);
        myArray[0] = id;
        return myArray;
    }

    function getId() view public returns (string memory) {
        return id;
    }

    function giveRole(address account) external onlyRole(OWNER_ROLE) {
        _grantRole(BRIDGE_ROLE, account);
    }
}