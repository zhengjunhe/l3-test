package keys

import (
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_keys(t *testing.T) {

	//privateKeySli := common.Hex2Bytes("4aa5560f09a2bd30898447a7a5dcb20ab0c058e5c09269daeb22a81b31e83607")
	privateKeySli := common.Hex2Bytes("c5d9dfd6ca71a1a1606828caabed3e587465acdb4591cc305d8bf0a37e9e01c0")

	ethPri := secp256k1.PrivKeyFromBytes(privateKeySli)

	fmt.Println("SerializeUncompressed", common.Bytes2Hex(ethPri.PubKey().SerializeUncompressed()))
	fmt.Println("SerializeCompressed", common.Bytes2Hex(ethPri.PubKey().SerializeCompressed()))

	assert.Equal(t, nil, nil)
}

//04
//
//b1822a0cfb926a60a203e1000c9054275f5c8ef946659682d1687cbb9d99f7d8824c8750b9f3711c0dc3284c186641e4011d7562ad0263a553bc1bf8ccca5177

func Test_key2address(t *testing.T) {

	//privateKeySli := common.Hex2Bytes("60279c35f089e07a9573b683d5e47d4d9fb42cc5288d33263ae14ac05201c40e")
	privateKeySli := common.Hex2Bytes("534049ad41dd0c181c4e15665291ca00867f7950f6302f6ee49731d4e4d62c9a")

	ethPri := secp256k1.PrivKeyFromBytes(privateKeySli)

	fmt.Println("SerializeUncompressed", common.Bytes2Hex(ethPri.PubKey().SerializeUncompressed()))
	fmt.Println("SerializeCompressed", common.Bytes2Hex(ethPri.PubKey().SerializeCompressed()))

	pub, err := crypto.UnmarshalPubkey(ethPri.PubKey().SerializeUncompressed())
	assert.Equal(t, err, nil)
	addr := crypto.PubkeyToAddress(*pub)
	fmt.Println("addr", addr)
}
