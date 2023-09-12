package arbcompressfuzz

// Edit if desired. Code generated by "go-fuzz-fill-utils ./blsSignatures/".

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto/bls12381"
	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/blsSignatures"
)

func Fuzz_PublicKey_ToTrusted(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pubKey *bls12381.PointG2
		var validityProof *bls12381.PointG1
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&pubKey, &validityProof)
		if pubKey == nil || validityProof == nil {
			return
		}

		p1, err := blsSignatures.NewPublicKey(pubKey, validityProof)
		if err != nil {
			return
		}
		p1.ToTrusted()
	})
}

func Fuzz_AggregatePublicKeys(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pubKeys []blsSignatures.PublicKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&pubKeys)

		blsSignatures.AggregatePublicKeys(pubKeys)
	})
}

func Fuzz_AggregateSignatures(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var sigs []blsSignatures.Signature
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&sigs)

		blsSignatures.AggregateSignatures(sigs)
	})
}

func Fuzz_KeyValidityProof(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pubKey *bls12381.PointG2
		var privateKey blsSignatures.PrivateKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&pubKey, &privateKey)
		if pubKey == nil {
			return
		}

		blsSignatures.KeyValidityProof(pubKey, privateKey)
	})
}

func Fuzz_NewPublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pubKey *bls12381.PointG2
		var validityProof *bls12381.PointG1
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&pubKey, &validityProof)
		if pubKey == nil || validityProof == nil {
			return
		}

		blsSignatures.NewPublicKey(pubKey, validityProof)
	})
}

func Fuzz_NewTrustedPublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pubKey *bls12381.PointG2
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&pubKey)
		if pubKey == nil {
			return
		}

		blsSignatures.NewTrustedPublicKey(pubKey)
	})
}

func Fuzz_PrivateKeyFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		blsSignatures.PrivateKeyFromBytes(in)
	})
}

func Fuzz_PrivateKeyToBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var priv blsSignatures.PrivateKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&priv)

		blsSignatures.PrivateKeyToBytes(priv)
	})
}

func Fuzz_PublicKeyFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte, trustedSource bool) {
		blsSignatures.PublicKeyFromBytes(in, trustedSource)
	})
}

func Fuzz_PublicKeyFromPrivateKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var privateKey blsSignatures.PrivateKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&privateKey)

		blsSignatures.PublicKeyFromPrivateKey(privateKey)
	})
}

func Fuzz_PublicKeyToBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var pub blsSignatures.PublicKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&pub)

		blsSignatures.PublicKeyToBytes(pub)
	})
}

func Fuzz_SignMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var priv blsSignatures.PrivateKey
		var message []byte
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&priv, &message)

		blsSignatures.SignMessage(priv, message)
	})
}

func Fuzz_SignatureFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		blsSignatures.SignatureFromBytes(in)
	})
}

func Fuzz_SignatureToBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var sig blsSignatures.Signature
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&sig)

		blsSignatures.SignatureToBytes(sig)
	})
}

func Fuzz_VerifyAggregatedSignatureDifferentMessages(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var sig blsSignatures.Signature
		var messages [][]byte
		var pubKeys []blsSignatures.PublicKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&sig, &messages, &pubKeys)

		blsSignatures.VerifyAggregatedSignatureDifferentMessages(sig, messages, pubKeys)
	})
}

func Fuzz_VerifyAggregatedSignatureSameMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var sig blsSignatures.Signature
		var message []byte
		var pubKeys []blsSignatures.PublicKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&sig, &message, &pubKeys)

		blsSignatures.VerifyAggregatedSignatureSameMessage(sig, message, pubKeys)
	})
}

func Fuzz_VerifySignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var sig blsSignatures.Signature
		var message []byte
		var publicKey blsSignatures.PublicKey
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&sig, &message, &publicKey)

		blsSignatures.VerifySignature(sig, message, publicKey)
	})
}