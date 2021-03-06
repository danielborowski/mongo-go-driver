package bson

import "testing"

type encodetest struct {
	Field1String  string
	Field1Int64   int64
	Field1Float64 float64
	Field2String  string
	Field2Int64   int64
	Field2Float64 float64
	Field3String  string
	Field3Int64   int64
	Field3Float64 float64
	Field4String  string
	Field4Int64   int64
	Field4Float64 float64
}

type nestedtest1 struct {
	Nested nestedtest2
}

type nestedtest2 struct {
	Nested nestedtest3
}

type nestedtest3 struct {
	Nested nestedtest4
}

type nestedtest4 struct {
	Nested nestedtest5
}

type nestedtest5 struct {
	Nested nestedtest6
}

type nestedtest6 struct {
	Nested nestedtest7
}

type nestedtest7 struct {
	Nested nestedtest8
}

type nestedtest8 struct {
	Nested nestedtest9
}

type nestedtest9 struct {
	Nested nestedtest10
}

type nestedtest10 struct {
	Nested nestedtest11
}

type nestedtest11 struct {
	Nested encodetest
}

var encodetestInstance = encodetest{
	Field1String:  "foo",
	Field1Int64:   1,
	Field1Float64: 3.0,
	Field2String:  "bar",
	Field2Int64:   2,
	Field2Float64: 3.1,
	Field3String:  "baz",
	Field3Int64:   3,
	Field3Float64: 3.14,
	Field4String:  "qux",
	Field4Int64:   4,
	Field4Float64: 3.141,
}

var nestedInstance = nestedtest1{
	nestedtest2{
		nestedtest3{
			nestedtest4{
				nestedtest5{
					nestedtest6{
						nestedtest7{
							nestedtest8{
								nestedtest9{
									nestedtest10{
										nestedtest11{
											encodetest{
												Field1String:  "foo",
												Field1Int64:   1,
												Field1Float64: 3.0,
												Field2String:  "bar",
												Field2Int64:   2,
												Field2Float64: 3.1,
												Field3String:  "baz",
												Field3Int64:   3,
												Field3Float64: 3.14,
												Field4String:  "qux",
												Field4Int64:   4,
												Field4Float64: 3.141,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func BenchmarkEncodingv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(encodetestInstance)
	}
}

func BenchmarkEncodingv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Marshalv2(encodetestInstance)
	}
}

func BenchmarkEncodingv2ToDocument(b *testing.B) {
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf, _ = Marshalv2(encodetestInstance)
		_, _ = ReadDocument(buf)
	}
}

func BenchmarkEncodingDocument(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = MarshalDocument(encodetestInstance)
	}
}

func BenchmarkEncodingv1Nested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(nestedInstance)
	}
}

func BenchmarkEncodingv2Nested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Marshalv2(nestedInstance)
	}
}

func BenchmarkEncodingv2ToDocumentNested(b *testing.B) {
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf, _ = Marshalv2(nestedInstance)
		_, _ = ReadDocument(buf)
	}
}

func BenchmarkEncodingDocumentNested(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = MarshalDocument(nestedInstance)
	}
}
