package quip_test

import (
	"os"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/fluent/quip"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
)

func Example() {
	nb := basicnode.Prototype.Any.NewBuilder()
	err := quip.BuildMap(nb, 4, func(ma ipld.MapAssembler) {
		quip.MapEntry(ma, "some key", quip.AssignString("some value"))
		quip.MapEntry(ma, "another key", quip.AssignString("another value"))
		quip.MapEntry(ma, "nested map", quip.Map(2, func(ma ipld.MapAssembler) {
			quip.MapEntry(ma, "deeper entries", quip.AssignString("deeper values"))
			quip.MapEntry(ma, "more deeper entries", quip.AssignString("more deeper values"))
		}))
		quip.MapEntry(ma, "nested list", quip.List(2, func(la ipld.ListAssembler) {
			quip.ListEntry(la, quip.AssignInt(1))
			quip.ListEntry(la, quip.AssignInt(2))
		}))
	})
	if err != nil {
		panic(err)
	}
	n := nb.Build()
	dagjson.Encoder(n, os.Stdout)

	// Output:
	// {
	// 	"some key": "some value",
	// 	"another key": "another value",
	// 	"nested map": {
	// 		"deeper entries": "deeper values",
	// 		"more deeper entries": "more deeper values"
	// 	},
	// 	"nested list": [
	// 		1,
	// 		2
	// 	]
	// }
}
