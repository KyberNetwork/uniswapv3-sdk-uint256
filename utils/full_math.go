package utils

import (
	"errors"
	"math/big"

	"github.com/holiman/uint256"
)

var (
	ErrMulDivOverflow = errors.New("muldiv overflow")
	One               = big.NewInt(1)
)

// MulDivRoundingUp Calculates ceil(a×b÷denominator) with full precision
func MulDivRoundingUp(a, b, denominator *uint256.Int) (*uint256.Int, error) {
	var result Uint256
	return &result, MulDivRoundingUpV2(a, b, denominator, &result)
}

func MulDivRoundingUpV2(a, b, denominator, result *uint256.Int) error {
	var remainder Uint256
	err := MulDivV2(a, b, denominator, result, &remainder)
	if err != nil {
		return err
	}

	if !remainder.IsZero() {
		if result.Cmp(MaxUint256) == 0 {
			return ErrInvariant
		}
		result.AddUint64(result, 1)
	}
	return nil
}

// MulDivV2 z=floor(a×b÷denominator), r=a×b%denominator
// (pass remainder=nil if not required)
// (the main usage for `remainder` is to be used in `MulDivRoundingUpV2` to determine if we need to round up, so it won't have to call MulMod again)
func MulDivV2(x, y, d, z, r *uint256.Int) error {
	if x.IsZero() || y.IsZero() || d.IsZero() {
		z.Clear()
		return nil
	}
	p := umul(x, y)

	var quot [8]uint64
	rem := udivrem(quot[:], p[:], d)
	if r != nil {
		r.Set(&rem)
	}

	copy(z[:], quot[:4])

	if (quot[4] | quot[5] | quot[6] | quot[7]) != 0 {
		return ErrMulDivOverflow
	}
	return nil
}

// MulDiv Calculates floor(a×b÷denominator) with full precision
func MulDiv(a, b, denominator *uint256.Int) (*uint256.Int, error) {
	result, overflow := new(uint256.Int).MulDivOverflow(a, b, denominator)
	if overflow {
		return nil, ErrMulDivOverflow
	}
	return result, nil
}

// DivRoundingUp Returns ceil(x / y)
func DivRoundingUp(a, denominator, result *uint256.Int) {
	var rem uint256.Int
	result.DivMod(a, denominator, &rem)
	if !rem.IsZero() {
		result.AddUint64(result, 1)
	}
}
