// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "forge-std/StdJson.sol";

import "../src/Diamond.sol";
import "../src/facets/DiamondCutFacet.sol";
import "../src/facets/DiamondLoupeFacet.sol";
import "../src/facets/Test1Facet.sol";
import "../src/facets/Test2Facet.sol";
import "../src/upgradeInitializers/DiamondInit.sol";
import "../src/interfaces/IDiamondCut.sol";

contract DiamondTest is Test {
    uint mainnetFork;

    Diamond diamond;
    DiamondInit diamondInit;
    DiamondCutFacet diamondCutFacet;
    DiamondLoupeFacet diamondLoupeFacet;
    Test1Facet test1Facet;
    Test2Facet test2Facet;
   
    IDiamondCut ICut;
    IDiamondLoupe ILoupe;

    using stdJson for string;

    function setUp() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address owner = vm.addr(deployerPrivateKey);
        vm.startBroadcast(deployerPrivateKey);
        
        diamondCutFacet = new DiamondCutFacet();
        diamondLoupeFacet = new DiamondLoupeFacet();
        test1Facet = new Test1Facet();
        test2Facet = new Test2Facet();
        diamond = new Diamond(owner, address(diamondCutFacet));

        address[] memory facetAddressList;
        IDiamondCut.FacetCut[] memory cut = new IDiamondCut.FacetCut[](1);

        bytes4[] memory selectors = getSelectors("DiamondLoupeFacet");

        cut[0] = (
            IDiamondCut.FacetCut({
                facetAddress: address(diamondLoupeFacet),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: selectors
            })
        );

        ILoupe = IDiamondLoupe(address(diamond));
        ICut = IDiamondCut(address(diamond));
        ICut.diamondCut(cut, address(0x0), "");

        vm.stopBroadcast();
    }

    // 1 'should have three facets -- call to facetAddresses function'
    function testThreeFacets() public {}

    // 2 'facets should have the right function selectors -- call to facetFunctionSelectors function'
    function testFacetFunctionSelectors() public {}

    // 3 'selectors should be associated to facets correctly -- multiple calls to facetAddress function'
    function testSelectorsAssociatedToFacets() public {}

    // 4 'should add test1 functions'
    function testAddTest1Functions() public {}

    // 5 'should test function call'
    function testFunctionCall() public {}

    // 6 'should replace supportsInterface function'
    function testReplaceSupportsInterfaceFunction() public {}

    // 7 'should add test2 functions'
    function testAddTest2Functions() public {}

    // 8 'should remove some test2 functions'
    function testRemoveSomeTest2Functions() public {}

    // 9 'should remove some test1 functions'
    function testRemoveSomeTest1Functions() public {}

    // 10 'remove all functions and facets except \'diamondCut\' and \'facets\''
    function testRemoveAllExceptDiamondCutAndFacets() public {}

    // 11 'add most functions and facets'
    function testAddMostFunctionsAndFacets() public {}

    function getSelectors(
        string memory contractName
    ) internal returns (bytes4[] memory) {
        string memory root = vm.projectRoot();
        string memory path = string.concat(
            root,
            "/contracts/out/",
            contractName,
            ".sol/",
            contractName,
            ".json"
        );
        string memory json = vm.readFile(path);
        bool exists = vm.keyExists(json, ".methodIdentifiers");
        assertTrue(exists);

        bytes memory methodIdentifiers = json.parseRaw(
            "$.methodIdentifiers[*]"
        );
        bytes[] memory decodedSelectors = abi.decode(
            methodIdentifiers,
            (bytes[])
        );

        bytes4[] memory selectors = new bytes4[](decodedSelectors.length);
        for (uint i = 0; i < decodedSelectors.length; i++) {
            bytes memory temp = bytes.concat(decodedSelectors[i]);
            selectors[i] = bytes4(temp);
        }

        return selectors;
    }
}
