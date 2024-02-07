// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "forge-std/StdJson.sol";

import "../src/Diamond.sol";
import "../src/facets/DiamondCutFacet.sol";
import "../src/facets/DiamondLoupeFacet.sol";
import "../src/facets/OwnershipFacet.sol";
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
    OwnershipFacet ownershipFacet;
    Test1Facet test1Facet;
    Test2Facet test2Facet;

    address[] facetAddresses;
    mapping(address => bytes4[]) testSelectors;

    using stdJson for string;

    function setUp() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address owner = vm.addr(deployerPrivateKey);
        vm.startBroadcast(deployerPrivateKey);

        diamondCutFacet = new DiamondCutFacet();
        diamondLoupeFacet = new DiamondLoupeFacet();
        ownershipFacet = new OwnershipFacet();
        test1Facet = new Test1Facet();
        test2Facet = new Test2Facet();
        diamond = new Diamond(owner, address(diamondCutFacet));

        IDiamondCut.FacetCut[] memory cut = new IDiamondCut.FacetCut[](2);

        bytes4[] memory functionSelectors = new bytes4[](5);
        functionSelectors[0] = IDiamondLoupe.facets.selector; // hex"7a0ed627"
        functionSelectors[1] = IDiamondLoupe.facetFunctionSelectors.selector; // hex"adfca15e"
        functionSelectors[2] = IDiamondLoupe.facetAddresses.selector; // hex"52ef6b2c"
        functionSelectors[3] = IDiamondLoupe.facetAddress.selector; // hex"cdffacc6"
        functionSelectors[4] = IERC165.supportsInterface.selector; // hex"01ffc9a7"

        cut[0] = IDiamondCut.FacetCut({
            facetAddress: address(diamondLoupeFacet),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });

        functionSelectors = new bytes4[](2);
        functionSelectors[0] = OwnershipFacet.transferOwnership.selector; // hex"f2fde38b"
        functionSelectors[1] = OwnershipFacet.owner.selector; // hex"8da5cb5b"

        cut[1] = IDiamondCut.FacetCut({
            facetAddress: address(ownershipFacet),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });

        IDiamondCut(address(diamond)).diamondCut(cut, address(0), "");

        testSelectors[address(diamondCutFacet)] = [bytes4(hex"1f931c1c")];
        testSelectors[address(diamondLoupeFacet)] = [
            bytes4(hex"7a0ed627"),
            bytes4(hex"adfca15e"),
            bytes4(hex"52ef6b2c"),
            bytes4(hex"cdffacc6"),
            bytes4(hex"01ffc9a7")
        ];
        testSelectors[address(ownershipFacet)] = [
            bytes4(hex"f2fde38b"),
            bytes4(hex"8da5cb5b")
        ];

        // vm.stopBroadcast();
    }

    // 1 'should have three facets -- call to facetAddresses function'
    function testThreeFacets() public {
        facetAddresses = IDiamondLoupe(address(diamond)).facetAddresses();
        assertEq(facetAddresses.length, 3);
    }

    // 2 'facets should have the right function selectors -- call to facetFunctionSelectors function'
    function testFacetFunctionSelectors() public {
        bytes4[] memory selectors = IDiamondLoupe(address(diamond))
            .facetFunctionSelectors(address(diamondCutFacet));
        assertEq(
            testSelectors[address(diamondCutFacet)].length,
            selectors.length
        );
        assertEq(testSelectors[address(diamondCutFacet)][0], selectors[0]);

        selectors = IDiamondLoupe(address(diamond)).facetFunctionSelectors(
            address(diamondLoupeFacet)
        );
        assertEq(
            testSelectors[address(diamondLoupeFacet)].length,
            selectors.length
        );
        assertEq(testSelectors[address(diamondLoupeFacet)][0], selectors[0]);

        selectors = IDiamondLoupe(address(diamond)).facetFunctionSelectors(
            address(ownershipFacet)
        );
        assertEq(
            testSelectors[address(ownershipFacet)].length,
            selectors.length
        );
        assertEq(testSelectors[address(ownershipFacet)][0], selectors[0]);
    }

    // 3 'selectors should be associated to facets correctly -- multiple calls to facetAddress function'
    function testSelectorsAssociatedToFacets() public {
        assertEq(
            address(diamondCutFacet),
            IDiamondLoupe(address(diamond)).facetAddress(bytes4(hex"1f931c1c"))
        );
        assertEq(
            address(diamondLoupeFacet),
            IDiamondLoupe(address(diamond)).facetAddress(bytes4(hex"cdffacc6"))
        );
        assertEq(
            address(diamondLoupeFacet),
            IDiamondLoupe(address(diamond)).facetAddress(bytes4(hex"01ffc9a7"))
        );
        assertEq(
            address(ownershipFacet),
            IDiamondLoupe(address(diamond)).facetAddress(bytes4(hex"f2fde38b"))
        );
    }

    // 4 'should add test1 functions'
    function testAddTest1Functions() public {
        bytes4[] memory functionSelectors = new bytes4[](1);

        functionSelectors[0] = Test1Facet.test1Func1.selector;
        IDiamondCut.FacetCut[] memory cut = new IDiamondCut.FacetCut[](1);
        cut[0] = IDiamondCut.FacetCut({
            facetAddress: address(test1Facet),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        IDiamondCut(address(diamond)).diamondCut(cut, address(0), "");
        bytes4[] memory selectors = IDiamondLoupe(address(diamond)).facetFunctionSelectors(
            address(test1Facet)
        );
        assertEq(functionSelectors[0], selectors[0]);
    }

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

        string[] memory decodedSelectors = abi.decode(
            methodIdentifiers,
            (string[])
        );
        
        bytes4[] memory selectors = new bytes4[](decodedSelectors.length);
        for (uint i = 0; i < decodedSelectors.length; i++) {
            bytes memory temp = abi.encodePacked(decodedSelectors[i]);
            selectors[i] = bytes4(temp);
        }
        console.log(vm.toString(selectors[1]));

        return selectors;
    }
}
