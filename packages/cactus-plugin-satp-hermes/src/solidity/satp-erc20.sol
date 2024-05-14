// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.15;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@hyperledger/cactus-plugin-bungee-hermes/src/test/typescript/solidity/lock-asset-contract/ITraceableContract.sol";


contract SATPContract is Ownable, ERC20, ITraceableContract{
 
    address public bridge_address;

    mapping (string => uint256) public balanceOf;

    constructor(address _bridge_address) ERC20("SATPToken", "SATP")  Ownable(_bridge_address){
        bridge_address = address(_bridge_address);
    }

    function mint(uint256 amount) external onlyOwner {
        _mint(bridge_address, amount);
    }

    function burn(uint256 amount) external onlyOwner {
        _burn(bridge_address, amount);
    }

    function assign(address receiver_account, uint256 amount) external onlyOwner {
        _transfer(bridge_address, receiver_account, amount);
    }

    function getAllAssetsIDs() external view returns (string[] memory) {
        string[] memory myArray = new string[](1);
        myArray[0] = Strings.toHexString(address(this));
        return myArray;
    }
}