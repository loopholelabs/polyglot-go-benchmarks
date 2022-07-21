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

type Benchmark struct {
	Message     string
	Embedded    *EmbeddedMessage
	EmbeddedMap *Map
	EnumMessage *EnumMessage
	Repeated    []*EmbeddedMessage
}

type EmbeddedMessage struct {
	EmbeddedData []byte
}

type Map struct {
	BuiltMap MapBuiltMapMap
}

type MapBuiltMapMap map[string]float64

type EnumMessageEmbeddedEnum uint32

const (
	EnumMessageUNIVERSAL = EnumMessageEmbeddedEnum(0)
)

type EnumMessage struct {
	error  error
	ignore bool

	Message  string
	Embedded EnumMessageEmbeddedEnum
}
