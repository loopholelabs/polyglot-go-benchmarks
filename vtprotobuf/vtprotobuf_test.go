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

package vtprotobuf

import (
	"github.com/loopholelabs/polyglot-go-benchmarks"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()
	val := benchmarks.GenerateVTProtobuf()
	var buf []byte
	var err error
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchmarks.TestSize; j++ {
			buf, err = val.MarshalVT()
			if err != nil {
				b.Error(err)
			}
			buf = buf[:0]
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()
	val := benchmarks.GenerateVTProtobuf()
	buf, err := val.MarshalVT()
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchmarks.TestSize; j++ {
			err = val.UnmarshalVT(buf)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkDecodeValidate(b *testing.B) {
	b.ReportAllocs()
	val := benchmarks.GenerateVTProtobuf()
	buf, err := val.MarshalVT()
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < benchmarks.TestSize; j++ {
			err = val.UnmarshalVT(buf)
			if err != nil {
				b.Error(err)
			}
			if !benchmarks.ValidateVTProtobuf(val) {
				b.Error("Validation failed")
			}
		}
	}
}
