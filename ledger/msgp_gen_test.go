// +build !skip_msgp_testing

package ledger

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/algorand/go-algorand/protocol"
	"github.com/algorand/go-algorand/test/partitiontest"
	"github.com/algorand/msgp/msgp"
)

func TestMarshalUnmarshalCatchpointFileHeader(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := CatchpointFileHeader{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingCatchpointFileHeader(t *testing.T) {
	protocol.RunEncodingTest(t, &CatchpointFileHeader{})
}

func BenchmarkMarshalMsgCatchpointFileHeader(b *testing.B) {
	v := CatchpointFileHeader{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgCatchpointFileHeader(b *testing.B) {
	v := CatchpointFileHeader{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalCatchpointFileHeader(b *testing.B) {
	v := CatchpointFileHeader{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalcatchpointFileBalancesChunk(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := catchpointFileBalancesChunk{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingcatchpointFileBalancesChunk(t *testing.T) {
	protocol.RunEncodingTest(t, &catchpointFileBalancesChunk{})
}

func BenchmarkMarshalMsgcatchpointFileBalancesChunk(b *testing.B) {
	v := catchpointFileBalancesChunk{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgcatchpointFileBalancesChunk(b *testing.B) {
	v := catchpointFileBalancesChunk{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalcatchpointFileBalancesChunk(b *testing.B) {
	v := catchpointFileBalancesChunk{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalencodedBalanceRecord(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := encodedBalanceRecord{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingencodedBalanceRecord(t *testing.T) {
	protocol.RunEncodingTest(t, &encodedBalanceRecord{})
}

func BenchmarkMarshalMsgencodedBalanceRecord(b *testing.B) {
	v := encodedBalanceRecord{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgencodedBalanceRecord(b *testing.B) {
	v := encodedBalanceRecord{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalencodedBalanceRecord(b *testing.B) {
	v := encodedBalanceRecord{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}