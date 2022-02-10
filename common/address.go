package common

import (
	"github.com/btcsuite/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
)

var hrp = "iaa"

type MyAddress struct {
	common.Address
}

func ToBechAddress(address common.Address) (string, error) {
	return ConvertAndEncode(hrp, address.Bytes())
}

func ToHexAddress(bech string) (common.Address, error) {
	_, data, err := DecodeAndConvert(bech)
	if err != nil {
		return [20]byte{}, err
	}
	return common.BytesToAddress(data), nil
}

//ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode(hrp, converted)

}

//DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes
func DecodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(bech)
	if err != nil {
		return "", nil, err
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, err
	}
	return hrp, converted, nil
}
