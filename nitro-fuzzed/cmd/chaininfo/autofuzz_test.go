package chaininfofuzz

// Edit if desired. Code generated by "go-fuzz-fill-utils ./chaininfo".

import (
	"math/big"
	"testing"

	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/cmd/chaininfo"
)

func Fuzz_GetChainConfig(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var chainId *big.Int
		var chainName string
		var genesisBlockNum uint64
		var l2ChainInfoFiles []string
		var l2ChainInfoJson string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&chainId, &chainName, &genesisBlockNum, &l2ChainInfoFiles, &l2ChainInfoJson)
		if chainId == nil {
			return
		}

		chaininfo.GetChainConfig(chainId, chainName, genesisBlockNum, l2ChainInfoFiles, l2ChainInfoJson)
	})
}

func Fuzz_GetRollupAddressesConfig(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var chainId uint64
		var chainName string
		var l2ChainInfoFiles []string
		var l2ChainInfoJson string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&chainId, &chainName, &l2ChainInfoFiles, &l2ChainInfoJson)

		chaininfo.GetRollupAddressesConfig(chainId, chainName, l2ChainInfoFiles, l2ChainInfoJson)
	})
}

func Fuzz_ProcessChainInfo(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var chainId uint64
		var chainName string
		var l2ChainInfoFiles []string
		var l2ChainInfoJson string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&chainId, &chainName, &l2ChainInfoFiles, &l2ChainInfoJson)

		chaininfo.ProcessChainInfo(chainId, chainName, l2ChainInfoFiles, l2ChainInfoJson)
	})
}