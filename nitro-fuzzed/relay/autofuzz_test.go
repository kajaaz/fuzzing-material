package arbcompressfuzz

// Edit if desired. Code generated by "go-fuzz-fill-utils ./relay".

import (
	"context"
	"testing"

	"github.com/infosecual/go-fuzz-fill-utils/fuzzer"
	"github.com/offchainlabs/nitro/broadcaster"
	"github.com/offchainlabs/nitro/relay"
	"github.com/spf13/pflag"
)

func Fuzz_MessageQueue_AddBroadcastMessages(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var q *relay.MessageQueue
		var feedMessages []*broadcaster.BroadcastFeedMessage
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&q, &feedMessages)
		if q == nil {
			return
		}

		q.AddBroadcastMessages(feedMessages)
	})
}

func Fuzz_Relay_GetListenerAddr(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *relay.Relay
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.GetListenerAddr()
	})
}

func Fuzz_Relay_Start(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *relay.Relay
		var ctx context.Context
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r, &ctx)
		if r == nil {
			return
		}

		r.Start(ctx)
	})
}

func Fuzz_Relay_StopAndWait(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r *relay.Relay
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&r)
		if r == nil {
			return
		}

		r.StopAndWait()
	})
}

func Fuzz_ConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var f1 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&f1)
		if f1 == nil {
			return
		}

		relay.ConfigAddOptions(f1)
	})
}

func Fuzz_L2ConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		relay.L2ConfigAddOptions(prefix, f2)
	})
}

// skipping Fuzz_NewRelay because parameters include func, chan, or unsupported interface: chan error

func Fuzz_NodeConfigAddOptions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var prefix string
		var f2 *pflag.FlagSet
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&prefix, &f2)
		if f2 == nil {
			return
		}

		relay.NodeConfigAddOptions(prefix, f2)
	})
}

func Fuzz_ParseRelay(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var _x1 context.Context
		var args []string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&_x1, &args)

		relay.ParseRelay(_x1, args)
	})
}
