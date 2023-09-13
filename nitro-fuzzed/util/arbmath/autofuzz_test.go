package util

// Edit if desired. Code generated by "go-fuzz-fill-utils ./arbmath/".

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/util/arbmath"
)

func Fuzz_ApproxExpBasisPoints(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value arbmath.Bips
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value)

		arbmath.ApproxExpBasisPoints(value)
	})
}

func Fuzz_ApproxSquareRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.ApproxSquareRoot(value)
	})
}

func Fuzz_BalancePerEther(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var balance *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&balance)
		if balance == nil {
			return
		}

		arbmath.BalancePerEther(balance)
	})
}

func Fuzz_BigAbs(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value)
		if value == nil {
			return
		}

		arbmath.BigAbs(value)
	})
}

func Fuzz_BigAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var augend *big.Int
		var addend *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&augend, &addend)
		if augend == nil || addend == nil {
			return
		}

		arbmath.BigAdd(augend, addend)
	})
}

func Fuzz_BigAddByUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var augend *big.Int
		var addend uint64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&augend, &addend)
		if augend == nil {
			return
		}

		arbmath.BigAddByUint(augend, addend)
	})
}

func Fuzz_BigAddFloat(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var augend *big.Float
		var addend *big.Float
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&augend, &addend)
		if augend == nil || addend == nil {
			return
		}

		arbmath.BigAddFloat(augend, addend)
	})
}

func Fuzz_BigDiv(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var dividend *big.Int
		var divisor *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&dividend, &divisor)
		if dividend == nil || divisor == nil {
			return
		}

		arbmath.BigDiv(dividend, divisor)
	})
}

func Fuzz_BigDivByInt(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var dividend *big.Int
		var divisor int64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&dividend, &divisor)
		if dividend == nil {
			return
		}

		arbmath.BigDivByInt(dividend, divisor)
	})
}

func Fuzz_BigDivByUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var dividend *big.Int
		var divisor uint64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&dividend, &divisor)
		if dividend == nil {
			return
		}

		arbmath.BigDivByUint(dividend, divisor)
	})
}

func Fuzz_BigEquals(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var first *big.Int
		var second *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&first, &second)
		if first == nil || second == nil {
			return
		}

		arbmath.BigEquals(first, second)
	})
}

func Fuzz_BigFloatMulByUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var multiplicand *big.Float
		var multiplier uint64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&multiplicand, &multiplier)
		if multiplicand == nil {
			return
		}

		arbmath.BigFloatMulByUint(multiplicand, multiplier)
	})
}

func Fuzz_BigGreaterThan(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var first *big.Int
		var second *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&first, &second)
		if first == nil || second == nil {
			return
		}

		arbmath.BigGreaterThan(first, second)
	})
}

func Fuzz_BigLessThan(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var first *big.Int
		var second *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&first, &second)
		if first == nil || second == nil {
			return
		}

		arbmath.BigLessThan(first, second)
	})
}

func Fuzz_BigMax(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var first *big.Int
		var second *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&first, &second)
		if first == nil || second == nil {
			return
		}

		arbmath.BigMax(first, second)
	})
}

func Fuzz_BigMin(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var first *big.Int
		var second *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&first, &second)
		if first == nil || second == nil {
			return
		}

		arbmath.BigMin(first, second)
	})
}

func Fuzz_BigMul(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var multiplicand *big.Int
		var multiplier *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&multiplicand, &multiplier)
		if multiplicand == nil || multiplier == nil {
			return
		}

		arbmath.BigMul(multiplicand, multiplier)
	})
}

func Fuzz_BigMulByBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		var bips arbmath.Bips
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value, &bips)
		if value == nil {
			return
		}

		arbmath.BigMulByBips(value, bips)
	})
}

func Fuzz_BigMulByFrac(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		var numerator int64
		var denominator int64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value, &numerator, &denominator)
		if value == nil {
			return
		}

		arbmath.BigMulByFrac(value, numerator, denominator)
	})
}

func Fuzz_BigMulByInt(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var multiplicand *big.Int
		var multiplier int64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&multiplicand, &multiplier)
		if multiplicand == nil {
			return
		}

		arbmath.BigMulByInt(multiplicand, multiplier)
	})
}

func Fuzz_BigMulByUfrac(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		var numerator uint64
		var denominator uint64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value, &numerator, &denominator)
		if value == nil {
			return
		}

		arbmath.BigMulByUfrac(value, numerator, denominator)
	})
}

func Fuzz_BigMulByUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var multiplicand *big.Int
		var multiplier uint64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&multiplicand, &multiplier)
		if multiplicand == nil {
			return
		}

		arbmath.BigMulByUint(multiplicand, multiplier)
	})
}

func Fuzz_BigMulFloat(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var multiplicand *big.Float
		var multiplier *big.Float
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&multiplicand, &multiplier)
		if multiplicand == nil || multiplier == nil {
			return
		}

		arbmath.BigMulFloat(multiplicand, multiplier)
	})
}

func Fuzz_BigSub(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var minuend *big.Int
		var subtrahend *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&minuend, &subtrahend)
		if minuend == nil || subtrahend == nil {
			return
		}

		arbmath.BigSub(minuend, subtrahend)
	})
}

func Fuzz_BigSubByUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var minuend *big.Int
		var subtrahend uint64
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&minuend, &subtrahend)
		if minuend == nil {
			return
		}

		arbmath.BigSubByUint(minuend, subtrahend)
	})
}

func Fuzz_BigToBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var natural *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&natural)
		if natural == nil {
			return
		}

		arbmath.BigToBips(natural)
	})
}

func Fuzz_BigToUintOrPanic(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value)
		if value == nil {
			return
		}

		arbmath.BigToUintOrPanic(value)
	})
}

func Fuzz_BigToUintSaturating(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value)
		if value == nil {
			return
		}

		arbmath.BigToUintSaturating(value)
	})
}

func Fuzz_ConcatByteSlices(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var slices [][]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&slices)

		arbmath.ConcatByteSlices(slices...)
	})
}

func Fuzz_FlipBit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var d1 common.Hash
		var bit byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&d1, &bit)

		arbmath.FlipBit(d1, bit)
	})
}

func Fuzz_FloatToBig(f *testing.F) {
	f.Fuzz(func(t *testing.T, value float64) {
		arbmath.FloatToBig(value)
	})
}

func Fuzz_IntMulByBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value int64
		var bips arbmath.Bips
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value, &bips)

		arbmath.IntMulByBips(value, bips)
	})
}

func Fuzz_Log2ceil(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.Log2ceil(value)
	})
}

// skipping Fuzz_MaxInt because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_MinInt because parameters include func, chan, or unsupported interface: T

func Fuzz_NaturalToBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, natural int64) {
		arbmath.NaturalToBips(natural)
	})
}

func Fuzz_NextOrCurrentPowerOf2(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.NextOrCurrentPowerOf2(value)
	})
}

func Fuzz_NextPowerOf2(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.NextPowerOf2(value)
	})
}

func Fuzz_PercentToBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, percentage int64) {
		arbmath.PercentToBips(percentage)
	})
}

func Fuzz_SaturatingAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, augend int64, addend int64) {
		arbmath.SaturatingAdd(augend, addend)
	})
}

func Fuzz_SaturatingCast(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.SaturatingCast(value)
	})
}

func Fuzz_SaturatingCastToBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.SaturatingCastToBips(value)
	})
}

func Fuzz_SaturatingCastToUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value *big.Int
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value)
		if value == nil {
			return
		}

		arbmath.SaturatingCastToUint(value)
	})
}

func Fuzz_SaturatingMul(f *testing.F) {
	f.Fuzz(func(t *testing.T, multiplicand int64, multiplier int64) {
		arbmath.SaturatingMul(multiplicand, multiplier)
	})
}

func Fuzz_SaturatingSub(f *testing.F) {
	f.Fuzz(func(t *testing.T, minuend int64, subtrahend int64) {
		arbmath.SaturatingSub(minuend, subtrahend)
	})
}

func Fuzz_SaturatingUAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, augend uint64, addend uint64) {
		arbmath.SaturatingUAdd(augend, addend)
	})
}

func Fuzz_SaturatingUCast(f *testing.F) {
	f.Fuzz(func(t *testing.T, value int64) {
		arbmath.SaturatingUCast(value)
	})
}

func Fuzz_SaturatingUMul(f *testing.F) {
	f.Fuzz(func(t *testing.T, multiplicand uint64, multiplier uint64) {
		arbmath.SaturatingUMul(multiplicand, multiplier)
	})
}

func Fuzz_SaturatingUSub(f *testing.F) {
	f.Fuzz(func(t *testing.T, minuend uint64, subtrahend uint64) {
		arbmath.SaturatingUSub(minuend, subtrahend)
	})
}

func Fuzz_SquareFloat(f *testing.F) {
	f.Fuzz(func(t *testing.T, value float64) {
		arbmath.SquareFloat(value)
	})
}

func Fuzz_SquareUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.SquareUint(value)
	})
}

func Fuzz_UfracToBigFloat(f *testing.F) {
	f.Fuzz(func(t *testing.T, numerator uint64, denominator uint64) {
		arbmath.UfracToBigFloat(numerator, denominator)
	})
}

func Fuzz_Uint32ToBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint32) {
		arbmath.Uint32ToBytes(value)
	})
}

func Fuzz_UintMulByBips(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value uint64
		var bips arbmath.Bips
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&value, &bips)

		arbmath.UintMulByBips(value, bips)
	})
}

func Fuzz_UintToBig(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.UintToBig(value)
	})
}

func Fuzz_UintToBigFloat(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.UintToBigFloat(value)
	})
}

func Fuzz_UintToBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		arbmath.UintToBytes(value)
	})
}

func Fuzz_WordsForBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, nbytes uint64) {
		arbmath.WordsForBytes(nbytes)
	})
}
