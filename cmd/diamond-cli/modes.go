package main

var modes []*Command = []*Command{
	{
		Name: "deploy",
		SubCommands: []*Command{
			{
				Name:        "diamond",
				Description: "Deploy a new diamond",
				SubCommands: []*Command{
					{
						Name:        "owner",
						Description: "Specify the Ethereum address of the owner",
					},
				},
			},
			{
				Name:        "facet",
				Description: "Deploy a facet to use in an existing diamond",
				SubCommands: []*Command{
					{
						Name:        "metadata",
						Description: "Path to contract metadata file",
					},
					{
						Name:        "constructor-args",
						Description: "Comma-separated list of constructor arguments",
					},
				},
			},
			{
				Name:        "init",
				Description: "Deploy a facet to use in an existing diamond",
				SubCommands: []*Command{},
			},
		},
	},

	{
		Name: "cut",
		SubCommands: []*Command{
			{
				Name:        "add",
				Description: "Add a new facet with specified function selectors",
				SubCommands: []*Command{
					{
						Name:        "address",
						Description: "Specify the Ethereum address of a facet",
					},
					{
						Name:        "selectors",
						Description: "Specify the function selectors",
					},
				},
			},
			{
				Name:        "replace",
				Description: "Replace selectors of an existing facet",
				SubCommands: []*Command{
					{
						Name:        "address",
						Description: "Specify the Ethereum address of a facet",
					},
					{
						Name:        "selectors",
						Description: "Specify the function selectors",
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove selectors from the diamond",
				SubCommands: []*Command{
					{
						Name:        "selectors",
						Description: "Specify the function selectors",
					},
				},
			},
		},
	},

	{
		Name: "loupe",
		SubCommands: []*Command{
			{
				Name:        "facets",
				Description: "Show all facets and their selectors",
			},
			{
				Name:        "addresses",
				Description: "Show all facet addresses used by a diamond",
			},
			{
				Name:        "facet-selectors",
				Description: "Show all function selectors provided by a facet",
				SubCommands: []*Command{
					{
						Name:        "address",
						Description: "Specify the Ethereum address of a facet",
					},
				},
			},
			{
				Name:        "facet-address",
				Description: "Show the facet that supports the given selector",
				SubCommands: []*Command{
					{
						Name:        "selector",
						Description: "Specify the function selector",
					},
				},
			},
			{
				Name:        "supports-interface",
				Description: "Show if the contract implements an interface",
				SubCommands: []*Command{
					{
						Name:        "id",
						Description: "Specify the interface identifier",
					},
				},
			},
		},
	},

	{
		Name: "test",
		SubCommands: []*Command{
			{
				Name:        "diamond",
				Description: "Deploy a new diamond",
				SubCommands: []*Command{
					{
						Name:        "owner",
						Description: "Specify the Ethereum address of the owner",
					},
				},
			},
		},
	},
}

func selectMode(selected string) *Command {
	defaultMode := &Command{
		SubCommands: []*Command{
			{
				Name:        "help",
				Description: "Show help message",
			},
			{
				Name:        "exit",
				Description: "Exit the interactive mode",
			},
		},
	}

	for _, mode := range modes {
		if mode.Name == selected {
			mode.SubCommands = append(mode.SubCommands, defaultMode.SubCommands...)
			return mode
		}
	}

	return nil
}
