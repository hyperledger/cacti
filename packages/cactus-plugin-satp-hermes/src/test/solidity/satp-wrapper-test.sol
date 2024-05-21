// SPDX-License-Identifier: GPL-3.0
        
pragma solidity >=0.4.22 <0.9.0;

// This import is automatically injected by Remix
import "remix_tests.sol"; 

// This import is required to use custom transaction context
// Although it may fail compilation in 'Solidity Compiler' plugin
// But it will work fine in 'Solidity Unit Testing' plugin
import "remix_accounts.sol";
import "../contracts/satp-wrapper.sol";
import "./../contracts/satp-erc20.sol";

contract SATPWrapTest {

    SATPContract contract1;
    SATPContract contract2;
    
    SATPWrapperContract wrapperContract;
    
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    function beforeEach() public {
        // Remix does not offer a set of methods that can change the msg.sender so every contract owner is the same

        wrapperContract = new SATPWrapperContract(address(this));

        contract1 = new SATPContract(address(wrapperContract), "ID1");
        contract2 = new SATPContract(address(this), "ID2");      
    }

    function checkWrap() public {
        Token memory token = Token(address(contract1), TokenType.ERC20, contract1.getId(), address(this), 99);

        wrapperContract.wrap(token, address(contract1));

        Token memory tokenReceived = wrapperContract.getToken(address(contract1));

        Assert.equal(tokenReceived.assetContract, token.assetContract, "Tokens don't match");

        Assert.equal(wrapperContract.getAllAssetsIDs()[0], contract1.getId(), "Ids don't match");
    }

    function checkUnwrap() public {
        Token memory token = Token(address(contract1), TokenType.ERC20, contract1.getId(), address(this), 99);

        wrapperContract.wrap(token, address(contract1));

        wrapperContract.unwrap(address(contract1));

        Token memory tokenReceived = wrapperContract.getToken(address(contract1));

        Assert.notEqual(tokenReceived.assetContract, token.assetContract, "Tokens don't match");
    }

    function checkMint() public {
        Token memory token = Token(address(contract1), TokenType.ERC20, contract1.getId(), address(this), 0);

        wrapperContract.wrap(token, address(contract1));

        wrapperContract.mint(address(contract1), 10);
        
        Assert.equal(contract1.balanceOf(address(wrapperContract)), 10, "Token not minted");
    }

    function checkBurn() public {
        Token memory token = Token(address(contract1), TokenType.ERC20, contract1.getId(), address(this), 0);

        wrapperContract.wrap(token, address(contract1));

        wrapperContract.mint(address(contract1), 10);

        wrapperContract.burn(address(contract1), 10);        

        Assert.equal(contract1.balanceOf(address(this)), 0, "Token not burned");
    }

    function checkAssign() public {
        Token memory token = Token(address(contract1), TokenType.ERC20, contract1.getId(), address(this), 0);

        wrapperContract.wrap(token, address(contract1));

        wrapperContract.mint(address(contract1), 10);

        wrapperContract.assign(address(contract1), address(TestsAccounts.getAccount(1)), 10);        
        
        Assert.equal(contract1.balanceOf(TestsAccounts.getAccount(1)), 10, "Token not assign");
        Assert.equal(wrapperContract.getToken(address(contract1)).assetContract, address(0), "Token was not unwrapped");
    }

    function checkLock() public {
        contract2.mint(address(this), 10);

        contract2.giveRole(address(wrapperContract));

        contract2.approve(address(wrapperContract), 5);

        Token memory token = Token(address(contract2), TokenType.ERC20, contract2.getId(), address(this), 0);

        wrapperContract.wrap(token, address(contract2));

        wrapperContract.lock(address(contract2), 5);

        Token memory tokenReceived = wrapperContract.getToken(address(contract2));

        Assert.equal(contract2.balanceOf(address(this)), 5, "Amount was not transfered to the wrapper contract address");
        Assert.equal(tokenReceived.amount, 5, "Token not unlocked");
    }

    function checkUnlock() public {
        contract2.mint(address(this), 10);

        contract2.giveRole(address(wrapperContract));

        contract2.approve(address(wrapperContract), 5);

        Token memory token = Token(address(contract2), TokenType.ERC20, contract2.getId(), address(this), 0);

        wrapperContract.wrap(token, address(contract2));

        wrapperContract.lock(address(contract2), 5);

        wrapperContract.unlock(address(contract2), 5);

        Token memory tokenReceived = wrapperContract.getToken(address(contract2));

        Assert.equal(contract2.balanceOf(address(this)), 10, "Amount was not transfered to the owner contract address");
        Assert.equal(contract2.balanceOf(address(wrapperContract)), 0, "Amount was not transfered to the owner contract address");
        Assert.equal(tokenReceived.amount, 0, "Token not unlocked");
    }
}
    