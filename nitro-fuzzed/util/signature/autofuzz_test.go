package util

// Edit if desired. Code generated by "go-fuzz-fill-utils ./signature/".

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/util/signature"
	"github.com/spf13/pflag"
)

func Fuzz_SignVerify_SignMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var v *signature.SignVerify
		var d2 [][]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&v, &d2)
		if v == nil {
			return
		}

		v.SignMessage(d2...)
	})
}

func Fuzz_SignVerify_VerifySignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var v *signature.SignVerify
		var ctx context.Context
		var s3 []byte
		var d4 [][]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&v, &ctx, &s3, &d4)
		if v == nil {
			return
		}

		v.VerifySignature(ctx, s3, d4...)
	})
}

func Fuzz_SimpleHmac_SignMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var config *signature.SimpleHmacConfig
		var d2 [][]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&config, &d2)
		if config == nil {
			return
		}

		h, err := signature.NewSimpleHmac(config)
		if err != nil {
			return
		}
		h.SignMessage(d2...)
	})
}

func Fuzz_SimpleHmac_VerifySignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var config *signature.SimpleHmacConfig
		var sig []byte
		var d3 [][]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&config, &sig, &d3)
		if config == nil {
			return
		}

		h, err := signature.NewSimpleHmac(config)
		if err != nil {
			return
		}
		h.VerifySignature(sig, d3...)
	})
}

func Fuzz_Verifier_VerifyData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var v *signature.Verifier
		var ctx context.Context
		var s3 []byte
		var d4 [][]byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&v, &ctx, &s3, &d4)
		if v == nil {
			return
		}

		v.VerifyData(ctx, s3, d4...)
	})
}

func Fuzz_Verifier_VerifyHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var v *signature.Verifier
		var ctx context.Context
		var s3 []byte
		var hash common.Hash
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&v, &ctx, &s3, &hash)
		if v == nil {
			return
		}

		v.VerifyHash(ctx, s3, hash)
	})
}

func Fuzz_DangerousFeedVerifierConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		signature.DangerousFeedVerifierConfigAddOptions(prefix, f2)
	})
}

func Fuzz_DataSignerFromPrivateKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var privateKey *ecdsa.PrivateKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&privateKey)
		if privateKey == nil {
			return
		}

		signature.DataSignerFromPrivateKey(privateKey)
	})
}

func Fuzz_FeedVerifierConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		signature.FeedVerifierConfigAddOptions(prefix, f2)
	})
}

func Fuzz_LoadSigningKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyConfig string) {
		signature.LoadSigningKey(keyConfig)
	})
}

// skipping Fuzz_NewSignVerify because parameters include func, chan, or unsupported interface: github.com/offchainlabs/nitro/util/signature.DataSignerFunc

func Fuzz_NewSimpleHmac(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var config *signature.SimpleHmacConfig
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&config)
		if config == nil {
			return
		}

		signature.NewSimpleHmac(config)
	})
}

// skipping Fuzz_NewVerifier because parameters include func, chan, or unsupported interface: github.com/offchainlabs/nitro/util/contracts.BatchPosterVerifierInterface

func Fuzz_SignVerifyConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		signature.SignVerifyConfigAddOptions(prefix, f2)
	})
}

func Fuzz_SimpleHmacConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		signature.SimpleHmacConfigAddOptions(prefix, f2)
	})
}

func Fuzz_SimpleHmacDangerousConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		signature.SimpleHmacDangerousConfigAddOptions(prefix, f2)
	})
}