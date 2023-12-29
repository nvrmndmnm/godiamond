// config_test.go
package main

import (
	"testing"
)

func TestValidateStandardContracts(t *testing.T) {
	config := Config{
		Contracts: map[string]ContractConfig{
			"diamond":       {},
			"diamond_init":  {},
			"cut_facet":     {},
			"loupe_facet":   {},
			"extra_contract":{},
		},
	}

	err := config.validateStandardContracts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	delete(config.Contracts, "diamond")
	err = config.validateStandardContracts()
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
