/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package protobuf

import (
	"github.com/loopholelabs/polyglot-go-benchmarks"
	"google.golang.org/protobuf/proto"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()
	val := benchmarks.GenerateProtobuf()
	marshaller := proto.MarshalOptions{Deterministic: true}
	var buf []byte
	var err error
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchmarks.TestSize; j++ {
			buf, err = marshaller.MarshalAppend(buf, val)
			if err != nil {
				b.Error(err)
			}
			buf = buf[:0]
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()
	val := benchmarks.GenerateProtobuf()
	buf, err := proto.Marshal(val)
	if err != nil {
		b.Error(err)
	}
	unmarshaller := proto.UnmarshalOptions{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchmarks.TestSize; j++ {
			err = unmarshaller.Unmarshal(buf, val)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkDecodeValidate(b *testing.B) {
	b.ReportAllocs()
	val := benchmarks.GenerateProtobuf()
	buf, err := proto.Marshal(val)
	if err != nil {
		b.Error(err)
	}
	unmarshaller := proto.UnmarshalOptions{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchmarks.TestSize; j++ {
			err = unmarshaller.Unmarshal(buf, val)
			if err != nil {
				b.Error(err)
			}
			if !benchmarks.ValidateProtobuf(val) {
				b.Error("Validation failed")
			}
		}
	}
}
