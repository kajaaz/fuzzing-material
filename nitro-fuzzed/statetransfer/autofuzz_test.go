package arbcompressfuzz

// Edit if desired. Code generated by "go-fuzz-fill-utils ./statetransfer/".

import (
	"testing"

	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/statetransfer"
)

func Fuzz_FieldReader_Close(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var f1 *statetransfer.FieldReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&f1)
		if f1 == nil {
			return
		}

		f1.Close()
	})
}

func Fuzz_FieldReader_More(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var f1 *statetransfer.FieldReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&f1)
		if f1 == nil {
			return
		}

		f1.More()
	})
}

func Fuzz_JsonAccountDataReaderr_GetNext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonAccountDataReaderr
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNext()
	})
}

func Fuzz_JsonAddressReader_GetNext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonAddressReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNext()
	})
}

func Fuzz_JsonInitDataReader_Close(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.Close()
	})
}

func Fuzz_JsonInitDataReader_GetAccountDataReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetAccountDataReader()
	})
}

func Fuzz_JsonInitDataReader_GetAddressTableReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetAddressTableReader()
	})
}

func Fuzz_JsonInitDataReader_GetNextBlockNumber(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNextBlockNumber()
	})
}

func Fuzz_JsonInitDataReader_GetRetryableDataReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetRetryableDataReader()
	})
}

func Fuzz_JsonListReader_Close(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var l *statetransfer.JsonListReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&l)
		if l == nil {
			return
		}

		l.Close()
	})
}

func Fuzz_JsonListReader_More(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var l *statetransfer.JsonListReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&l)
		if l == nil {
			return
		}

		l.More()
	})
}

func Fuzz_JsonRetryableDataReader_GetNext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.JsonRetryableDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNext()
	})
}

func Fuzz_MemoryAccountDataReaderr_GetNext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryAccountDataReaderr
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNext()
	})
}

func Fuzz_MemoryAddressReader_GetNext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryAddressReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNext()
	})
}

func Fuzz_MemoryInitDataReader_Close(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.Close()
	})
}

func Fuzz_MemoryInitDataReader_GetAccountDataReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetAccountDataReader()
	})
}

func Fuzz_MemoryInitDataReader_GetAddressTableReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetAddressTableReader()
	})
}

func Fuzz_MemoryInitDataReader_GetNextBlockNumber(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNextBlockNumber()
	})
}

func Fuzz_MemoryInitDataReader_GetRetryableDataReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryInitDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetRetryableDataReader()
	})
}

func Fuzz_MemoryRetryableDataReader_GetNext(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *statetransfer.MemoryRetryableDataReader
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetNext()
	})
}

func Fuzz_NewJsonInitDataReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, filepath string) {
		statetransfer.NewJsonInitDataReader(filepath)
	})
}

func Fuzz_NewMemoryInitDataReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var d1 *statetransfer.ArbosInitializationInfo
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&d1)
		if d1 == nil {
			return
		}

		statetransfer.NewMemoryInitDataReader(d1)
	})
}
