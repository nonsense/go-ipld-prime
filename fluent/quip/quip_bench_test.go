package quip_test

import (
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/fluent"
	"github.com/ipld/go-ipld-prime/fluent/quip"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
)

func BenchmarkQuip(b *testing.B) {
	b.ReportAllocs()

	var n ipld.Node
	var err error
	for i := 0; i < b.N; i++ {
		n, err = f1()
	}
	_ = n
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	var n ipld.Node
	var err error
	serial := `[{
		"destination": "/",
		"type": "overlay",
		"source": "none",
		"options": [
			"lowerdir=/",
			"upperdir=/tmp/overlay-root/upper",
			"workdir=/tmp/overlay-root/work"
		]
	}]`
	r := strings.NewReader(serial)
	for i := 0; i < b.N; i++ {
		nb := basicnode.Prototype.Any.NewBuilder()
		err = dagjson.Decoder(nb, r)
		n = nb.Build()
		r.Reset(serial)
	}
	_ = n
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkFluent(b *testing.B) {
	var n ipld.Node
	var err error
	for i := 0; i < b.N; i++ {
		n, err = fluent.BuildList(basicnode.Prototype.Any, -1, func(la fluent.ListAssembler) {
			la.AssembleValue().CreateMap(4, func(ma fluent.MapAssembler) {
				ma.AssembleEntry("destination").AssignString("/")
				ma.AssembleEntry("type").AssignString("overlay")
				ma.AssembleEntry("source").AssignString("none")
				ma.AssembleEntry("options").CreateList(-1, func(la fluent.ListAssembler) {
					la.AssembleValue().AssignString("lowerdir=" + "/")
					la.AssembleValue().AssignString("upperdir=" + "/tmp/overlay-root/upper")
					la.AssembleValue().AssignString("workdir=" + "/tmp/overlay-root/work")
				})
			})
		})
	}
	_ = n
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkReflect(b *testing.B) {
	var n ipld.Node
	var err error
	val := []interface{}{
		map[string]interface{}{
			"destination": "/",
			"type":        "overlay",
			"source":      "none",
			"options": []string{
				"lowerdir=/",
				"upperdir=/tmp/overlay-root/upper",
				"workdir=/tmp/overlay-root/work",
			},
		},
	}
	for i := 0; i < b.N; i++ {
		n, err = fluent.Reflect(basicnode.Prototype.Any, val)
	}
	_ = n
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkReflectIncludingInitialization(b *testing.B) {
	var n ipld.Node
	var err error
	for i := 0; i < b.N; i++ {
		n, err = fluent.Reflect(basicnode.Prototype.Any, []interface{}{
			map[string]interface{}{
				"destination": "/",
				"type":        "overlay",
				"source":      "none",
				"options": []string{
					"lowerdir=/",
					"upperdir=/tmp/overlay-root/upper",
					"workdir=/tmp/overlay-root/work",
				},
			},
		})
	}
	_ = n
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkAgonizinglyBare(b *testing.B) {
	var n ipld.Node
	var err error
	for i := 0; i < b.N; i++ {
		n, err = fab()
	}
	_ = n
	if err != nil {
		b.Fatal(err)
	}
}

func fab() (ipld.Node, error) {
	nb := basicnode.Prototype.Any.NewBuilder()
	la1, err := nb.BeginList(-1)
	if err != nil {
		return nil, err
	}
	ma, err := la1.AssembleValue().BeginMap(4)
	if err != nil {
		return nil, err
	}
	va, err := ma.AssembleEntry("destination")
	if err != nil {
		return nil, err
	}
	err = va.AssignString("/")
	if err != nil {
		return nil, err
	}
	va, err = ma.AssembleEntry("type")
	if err != nil {
		return nil, err
	}
	err = va.AssignString("overlay")
	if err != nil {
		return nil, err
	}
	va, err = ma.AssembleEntry("source")
	if err != nil {
		return nil, err
	}
	err = va.AssignString("none")
	if err != nil {
		return nil, err
	}
	va, err = ma.AssembleEntry("options")
	if err != nil {
		return nil, err
	}
	la2, err := va.BeginList(-4)
	if err != nil {
		return nil, err
	}
	err = la2.AssembleValue().AssignString("lowerdir=" + "/")
	if err != nil {
		return nil, err
	}
	err = la2.AssembleValue().AssignString("upperdir=" + "/tmp/overlay-root/upper")
	if err != nil {
		return nil, err
	}
	err = la2.AssembleValue().AssignString("workdir=" + "/tmp/overlay-root/work")
	if err != nil {
		return nil, err
	}
	err = la2.Finish()
	if err != nil {
		return nil, err
	}
	err = ma.Finish()
	if err != nil {
		return nil, err
	}
	err = la1.Finish()
	if err != nil {
		return nil, err
	}
	return nb.Build(), nil
}

func f1() (ipld.Node, error) {
	nb := basicnode.Prototype.Any.NewBuilder()
	err := quip.BuildList(nb, -1, func(la ipld.ListAssembler) {
		f2(la.AssembleValue(),
			"/",
			"overlay",
			"none",
			[]string{
				"lowerdir=" + "/",
				"upperdir=" + "/tmp/overlay-root/upper",
				"workdir=" + "/tmp/overlay-root/work",
			},
		)
	})
	if err != nil {
		return nil, err
	}
	return nb.Build(), nil
}

func f2(na ipld.NodeAssembler, a string, b string, c string, d []string) error {
	return quip.BuildMap(na, 4, func(ma ipld.MapAssembler) {
		quip.MapEntry(ma, "destination", quip.AssignString(a))
		quip.MapEntry(ma, "type", quip.AssignString(b))
		quip.MapEntry(ma, "source", quip.AssignString(c))
		quip.MapEntry(ma, "options", func(va ipld.NodeAssembler) {
			quip.List(va, int64(len(d)), func(la ipld.ListAssembler) {
				for _, s := range d {
					quip.ListEntry(la, quip.AssignString(s))
				}
			})
		})
	})
}
