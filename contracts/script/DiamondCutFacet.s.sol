// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/facets/DiamondCutFacet.sol";

contract DeployScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        DiamondCutFacet dcf = new DiamondCutFacet();
        console.log(address(dcf));
        
        vm.stopBroadcast();
    }
}
