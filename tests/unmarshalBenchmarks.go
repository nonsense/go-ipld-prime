package tests

import (
	"bytes"
	"testing"

	refmtjson "github.com/polydawn/refmt/json"

	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/encoding"
	"github.com/ipld/go-ipld-prime/tests/corpus"
)

// All of the marshalling and unmarshalling benchmark specs use JSON.
// This does mean we're measuring a bunch of stuff that has nothing to do
//  with the core operations of the Node/NodeBuilder interface.
// We do this so that:
// - we get a reasonable picture of how much time is spent in the IPLD Data Model
//    versus how much time is spent in the serialization efforts;
// - we can make direct comparisons to the standard library json marshalling
//    and unmarshalling, thus having a back-of-the-envelope baseline to compare.

var sink interface{}

func BenchmarkSpec_Unmarshal_Map3StrInt(b *testing.B, nb ipld.NodeBuilder) {
	msg := corpus.Map3StrInt()
	b.ResetTimer()

	var err error
	for i := 0; i < b.N; i++ {
		sink, err = encoding.Unmarshal(nb, refmtjson.NewDecoder(bytes.NewBufferString(msg)))
	}
	if err != nil {
		b.Fatalf("unmarshal errored: %s", err)
	}
}
