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

package benchmarks

import (
	"bytes"
	"fmt"
	polyglot "github.com/loopholelabs/polyglot-go-benchmarks/polyglot/benchmark"
	protobuf "github.com/loopholelabs/polyglot-go-benchmarks/protobuf/benchmark"
)

const (
	TestSize       = 1024
	Message        = "Hello, World!"
	MapKey0        = "Hello, World!"
	MapKey1        = "Hello Other World!"
	MapValue0      = 1.0
	MapValue1      = 2.0
	RepeatedFormat = "Hello Embedded World %d!"
	ENUMMessage    = "ENUM MESSAGE"
)

var (
	ByteMessage = []byte("Hello, World!")
)

func GeneratePolyglot() *polyglot.Benchmark {
	b := polyglot.NewBenchmark()
	b.Message = Message
	b.Embedded.EmbeddedData = ByteMessage
	b.EmbeddedMap.BuiltMap = polyglot.NewMapBuiltMapMap(2)
	b.EmbeddedMap.BuiltMap[MapKey0] = MapValue0
	b.EmbeddedMap.BuiltMap[MapKey1] = MapValue1
	for i := 0; i < 10; i++ {
		b.Repeated = append(b.Repeated, &polyglot.EmbeddedMessage{EmbeddedData: []byte(fmt.Sprintf(RepeatedFormat, i))})
	}

	b.EnumMessage.Message = ENUMMessage
	b.EnumMessage.Embedded = polyglot.EnumMessageUNIVERSAL

	return b
}

func GenerateProtobuf() *protobuf.Benchmark {
	b := &protobuf.Benchmark{
		Message: Message,
		Embedded: &protobuf.EmbeddedMessage{
			EmbeddedData: ByteMessage,
		},
		EmbeddedMap: &protobuf.Map{
			BuiltMap: map[string]float64{
				MapKey0: MapValue0,
				MapKey1: MapValue1,
			},
		},
		EnumMessage: &protobuf.EnumMessage{
			Message:  ENUMMessage,
			Embedded: protobuf.EnumMessage_UNIVERSAL,
		},
	}
	for i := 0; i < 10; i++ {
		b.Repeated = append(b.Repeated, &protobuf.EmbeddedMessage{EmbeddedData: []byte(fmt.Sprintf(RepeatedFormat, i))})
	}

	return b
}

func Generate() *Benchmark {
	b := &Benchmark{
		Message: Message,
		Embedded: &EmbeddedMessage{
			EmbeddedData: ByteMessage,
		},
		EmbeddedMap: &Map{
			BuiltMap: map[string]float64{
				MapKey0: MapValue0,
				MapKey1: MapValue1,
			},
		},
		EnumMessage: &EnumMessage{
			Message:  ENUMMessage,
			Embedded: EnumMessageUNIVERSAL,
		},
	}
	for i := 0; i < 10; i++ {
		b.Repeated = append(b.Repeated, &EmbeddedMessage{EmbeddedData: []byte(fmt.Sprintf(RepeatedFormat, i))})
	}

	return b
}

func ValidatePolyglot(b *polyglot.Benchmark) bool {
	if b.Message != Message {
		return false
	}

	if !bytes.Equal(b.Embedded.EmbeddedData, ByteMessage) {
		return false
	}

	if b.EmbeddedMap.BuiltMap[MapKey0] != MapValue0 {
		return false
	}

	if b.EmbeddedMap.BuiltMap[MapKey1] != MapValue1 {
		return false
	}

	for i := 0; i < 10; i++ {
		if !bytes.Equal(b.Repeated[i].EmbeddedData, []byte(fmt.Sprintf(RepeatedFormat, i))) {
			return false
		}
	}

	if b.EnumMessage.Message != ENUMMessage {
		return false
	}

	if b.EnumMessage.Embedded != polyglot.EnumMessageUNIVERSAL {
		return false
	}

	return true
}

func ValidateProtobuf(b *protobuf.Benchmark) bool {
	if b.Message != Message {
		return false
	}

	if !bytes.Equal(b.Embedded.EmbeddedData, ByteMessage) {
		return false
	}

	if b.EmbeddedMap.BuiltMap[MapKey0] != MapValue0 {
		return false
	}

	if b.EmbeddedMap.BuiltMap[MapKey1] != MapValue1 {
		return false
	}

	for i := 0; i < 10; i++ {
		if !bytes.Equal(b.Repeated[i].EmbeddedData, []byte(fmt.Sprintf(RepeatedFormat, i))) {
			return false
		}
	}

	if b.EnumMessage.Message != ENUMMessage {
		return false
	}

	if b.EnumMessage.Embedded != protobuf.EnumMessage_UNIVERSAL {
		return false
	}

	return true
}

func Validate(b *Benchmark) bool {
	if b.Message != Message {
		return false
	}

	if !bytes.Equal(b.Embedded.EmbeddedData, ByteMessage) {
		return false
	}

	if b.EmbeddedMap.BuiltMap[MapKey0] != MapValue0 {
		return false
	}

	if b.EmbeddedMap.BuiltMap[MapKey1] != MapValue1 {
		return false
	}

	for i := 0; i < 10; i++ {
		if !bytes.Equal(b.Repeated[i].EmbeddedData, []byte(fmt.Sprintf(RepeatedFormat, i))) {
			return false
		}
	}

	if b.EnumMessage.Message != ENUMMessage {
		return false
	}

	if b.EnumMessage.Embedded != EnumMessageUNIVERSAL {
		return false
	}

	return true
}
