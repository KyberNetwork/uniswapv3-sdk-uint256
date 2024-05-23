package msgpencode

import (
	"math/big"
	"math/bits"
)

const (
	negSign        = 255
	wordSizeInByte = bits.UintSize / 8
)

func EncodeInt(x *big.Int) []byte {
	if x == nil {
		return nil
	}

	numWords := len(x.Bits())
	b := make([]byte, 1 /* sign */ +(numWords*wordSizeInByte) /* words */)
	x.FillBytes(b[1:])
	if x.Sign() < 0 {
		b[0] = negSign
	} else {
		b[0] = 0
	}
	return b
}

func DecodeInt(b []byte) *big.Int {
	if b == nil {
		return nil
	}

	z := new(big.Int)
	z.SetBytes(b[1:])
	if b[0] == negSign {
		z.Neg(z)
	}
	return z
}
