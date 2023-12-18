# Go Diamond

## Introduction

Go Diamond is a CLI-tool for interacting with Diamond smart contracts based on the [ERC-2535 standard](https://eips.ethereum.org/EIPS/eip-2535). Built with Go and Foundry, it simplifies the deployment, management, and inspection of Diamond proxy contracts without requiring in-depth knowledge of the standard.  
_The use of Javascript in this project is deliberately avoided by all means._

## Features

1. Deploy: deploy new Diamond contracts with initial facets and function selectors.
2. Cut: Add, replace, and remove facets in a Diamond contract.
3. Loupe: Inspect a Diamond contract and its facets.

## Getting Started
Ensure to have Go and Foundry installed and configured.

**Clone the Repository:**
  ```bash
   git clone git@github.com:nvrmndmnm/godiamond.git
   cd godiamond
```

**Install Dependencies:**
  ```bash
go mod download
```

**Build:**
  ```bash
go build -o diamond-cli ./cmd/diamond-cli/
```

## Usage
**CLI**
```
Usage:
    diamond-cli deploy [options]
    diamond-cli cut [options]
    diamond-cli loupe [options]

Options:
    --rpc <name>          string    RPC identifier
    --chain-id <id>       int       Chain ID (default: 0)
    -c --config <path>    string    Load config file (default: "config.yaml")
    -d --debug                      Enable debug mode (default: disabled)
    -h --help                       Show help
```

**Deploy**
```
Commands:
    diamond <owner>                        Deploy a new diamond
    facet <metadata> <constructor-args>    Deploy a facet to use in an existing diamond
    init                                   Deploy initial set of contracts specified by the standard 
    help                                   Show help
    exit                                   Exit the deploy mode

Arguments:
    --owner              string    Ethereum address of the diamond owner
	--metadata           string    Path to contract metadata file
	--constructor-args   string    Comma-separated list of constructor arguments
```
**Cut**
```
Commands:
    add <address> <selectors>      Add a new facet with specified function selectors
    replace <address> <selectors>  Replace selectors of an existing facet
    remove <selectors>             Remove selectors from the diamond
    help                           Show help
    exit                           Exit the cut mode

Arguments:
    --address     string    Ethereum address of the facet
    --selectors   string    Comma-separated list of 4-byte function selectors
```
**Loupe**
```
Commands:
    facets                        Show all facets and their selectors
    addresses                     Show all facet addresses used by a diamond
    facet-selectors <address>     Show all function selectors provided by a facet
    facet-address <selector>      Show the facet that supports the given selector
    help                          Show help
    exit                          Exit the loupe mode

Arguments:
    --address           string    Ethereum address of the facet
    --selector          string    4-byte function selector representation 
```

## Examples

## Testing

Run Go tests:
```bash
go test ./...
```

Run Diamond contract tests with Foundry:
```bash
forge test
```


## Disclaimer
The project is under active development at the moment. Feedback and contributions are welcome.

For more information on Diamonds refer to [Awesome Diamonds](https://github.com/mudgen/awesome-diamonds).  
Big props to [Nick Mudge](https://github.com/mudgen/) for such a beautiful idea and implementation.

Licensed under the [MIT License](LICENSE).