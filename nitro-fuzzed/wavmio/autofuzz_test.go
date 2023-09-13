package arbcompressfuzz

// Edit if desired. Code generated by "go-fuzz-fill-utils ./wavmio/".

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/wavmio"
)

func Fuzz_ReadDelayedInboxMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, seqNum uint64) {
		wavmio.ReadDelayedInboxMessage(seqNum)
	})
}

func Fuzz_ReadInboxMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, msgNum uint64) {
		wavmio.ReadInboxMessage(msgNum)
	})
}

func Fuzz_ResolvePreImage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var hash common.Hash
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&hash)

		wavmio.ResolvePreImage(hash)
	})
}

func Fuzz_SetLastBlockHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var hash [32]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&hash)

		wavmio.SetLastBlockHash(hash)
	})
}

func Fuzz_SetPositionWithinMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, pos uint64) {
		wavmio.SetPositionWithinMessage(pos)
	})
}

func Fuzz_SetSendRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var hash [32]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&hash)

		wavmio.SetSendRoot(hash)
	})
}