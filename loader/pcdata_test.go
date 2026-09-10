package loader

import (
	"bytes"
	"testing"
)

func TestPcdataMarshalBinarySingleSafeEntry(t *testing.T) {
	data, err := (Pcdata{{PC: 32, Val: PCDATA_UnsafePointSafe}}).MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	want := []byte{0, 32, 0}
	if !bytes.Equal(data, want) {
		t.Fatalf("unexpected encoded table: got %x want %x", data, want)
	}
}

func TestPcdataMarshalBinarySingleUnsafeEntry(t *testing.T) {
	data, err := (Pcdata{{PC: 32, Val: PCDATA_UnsafePointUnsafe}}).MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	want := []byte{1, 32, 0}
	if !bytes.Equal(data, want) {
		t.Fatalf("unexpected encoded table: got %x want %x", data, want)
	}
}

func TestBuildLoadFuncNoPreemptEncodesSafeUnsafePoint(t *testing.T) {
	fn := buildLoadFunc(false, LoadOneItem{FuncName: "spin"}, 32, 0)

	data, err := fn.PcUnsafePoint.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) <= 1 {
		t.Fatalf("encoded unsafe-point table has no entries: %x", data)
	}
}
