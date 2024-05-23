//go:generate go run github.com/tinylib/msgp -unexported -tests=false -v
//msgp:tuple Tick
//msgp:shim *uint256.Int as:[]byte using:msgpencode.EncodeUint256/msgpencode.DecodeUint256
//msgp:shim *utils.Int128 as:[]byte using:msgpencode.EncodeInt256/msgpencode.DecodeInt256
//msgp:ignore TickDataProvider

package entities

import (
	"github.com/KyberNetwork/uniswapv3-sdk-uint256/utils"
	"github.com/holiman/uint256"
)

type Tick struct {
	Index          int
	LiquidityGross *uint256.Int
	LiquidityNet   *utils.Int128
}

// Provides information about ticks
type TickDataProvider interface {
	/**
	 * Return information corresponding to a specific tick
	 * @param tick the tick to load
	 */
	GetTick(tick int) (Tick, error)

	/**
	 * Return the next tick that is initialized within a single word
	 * @param tick The current tick
	 * @param lte Whether the next tick should be lte the current tick
	 * @param tickSpacing The tick spacing of the pool
	 */
	NextInitializedTickWithinOneWord(tick int, lte bool, tickSpacing int) (int, bool, error)

	// NextInitializedTickIndex return the next tick that is initialized
	NextInitializedTickIndex(tick int, lte bool) (int, bool, error)
}
