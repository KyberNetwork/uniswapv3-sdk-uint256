package entities

import (
	"testing"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// BenchmarkSortsBefore
// BenchmarkSortsBefore/daoleno_SortsBefore
// BenchmarkSortsBefore/daoleno_SortsBefore-16         	  565605	      2026 ns/op
// BenchmarkSortsBefore/KyberSwap_SortsBefore
// BenchmarkSortsBefore/KyberSwap_SortsBefore-16       	160554603	         7.101 ns/op
func BenchmarkSortsBefore(b *testing.B) {
	tokenA := entities.NewToken(1, common.HexToAddress("0xB8c77482e45F1F44dE1745F52C74426C631bDD52"),
		18, "BNB", "BNB")
	tokenB := entities.NewToken(1, common.HexToAddress("0xb62132e35a6c13ee1ee0f84dc5d40bad8d815206"),
		18, "NEXO", "Nexo")
	var before bool
	var err error

	b.Run("daoleno SortsBefore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			before, err = tokenA.SortsBefore(tokenB)
		}
		b.StopTimer()
		assert.False(b, before)
		assert.NoError(b, err)
	})

	b.Run("KyberSwap SortsBefore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			before, err = SortsBefore(tokenA, tokenB)
		}
		b.StopTimer()
		assert.False(b, before)
		assert.NoError(b, err)
	})
}
