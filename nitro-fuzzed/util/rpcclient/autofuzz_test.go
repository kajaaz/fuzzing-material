package util

// Edit if desired. Code generated by "go-fuzz-fill-utils ./rpcclient/".

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/util/rpcclient"
	"github.com/spf13/pflag"
)

func Fuzz_ClientConfig_Validate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var c *rpcclient.ClientConfig
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&c)
		if c == nil {
			return
		}

		c.Validate()
	})
}

func Fuzz_RpcClient_BatchCallContext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var c *rpcclient.RpcClient
		var ctx context.Context
		var b []rpc.BatchElem
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&c, &ctx, &b)
		if c == nil {
			return
		}

		c.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_RpcClient_CallContext because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_RpcClient_Close(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var c *rpcclient.RpcClient
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&c)
		if c == nil {
			return
		}

		c.Close()
	})
}

// skipping Fuzz_RpcClient_EthSubscribe because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_RpcClient_Start(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var c *rpcclient.RpcClient
		var ctx_in context.Context
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&c, &ctx_in)
		if c == nil {
			return
		}

		c.Start(ctx_in)
	})
}

// skipping Fuzz_NewRpcClient because parameters include func, chan, or unsupported interface: github.com/offchainlabs/nitro/util/rpcclient.ClientConfigFetcher

func Fuzz_RPCClientAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		var defaultConfig *rpcclient.ClientConfig
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2, &defaultConfig)
		if f2 == nil || defaultConfig == nil {
			return
		}

		rpcclient.RPCClientAddOptions(prefix, f2, defaultConfig)
	})
}
