package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func SigStringToByte4(sig string) [4]byte {
	sig1, err := hexutil.Decode(sig)

	if err != nil {
		fmt.Println(sig, "格式错误")
		return [4]byte{}
	}

	return [4]byte{sig1[0], sig1[1], sig1[2], sig1[3]}

}

func SigBytesToString(sig1 [4]byte) string {

	var sig = []byte{sig1[0], sig1[1], sig1[2], sig1[3]}

	return hexutil.Encode(sig)
}

func HexToInt64(h string) int64 {

	if h[:2] == "0x" {
		h = h[2:]
	}

	num, ok := new(big.Int).SetString(h, 16)

	if !ok {
		return 0
	} else {
		return num.Int64()
	}

}

func Hash(data []byte) string {

	h := crypto.Keccak256Hash(data)
	return h.String()
}
