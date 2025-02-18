package main

import "testing"

func TestBurn(t *testing.T) {

	mnemonic := "cupboard kiss effort champion faith bread firm cruise tissue decide all rotate"
	testKey := "7939624566468cfa3cb2c9f39d5ad83bdc7cf4356bfd1a7b8094abda6b0699d1"
	registerAddr := "0x76887513A7148277E064489DE680f4B9D545dCaD"
	tokenAddr := "0xFF7a3338fC5Ce5f303e1B211BABAC246F0e7edC7"
	burn(mnemonic, testKey, registerAddr, tokenAddr, "http://52.74.204.233:8545", 2, 97, 0.1)
}
