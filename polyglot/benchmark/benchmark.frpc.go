// Code generated by fRPC Go v0.5.1, DO NOT EDIT.
// source: benchmark.proto

package benchmark

import (
	"errors"
	"github.com/loopholelabs/polyglot-go"
)

var (
	NilDecode = errors.New("cannot decode into a nil root struct")
)

type EnumMessageEmbeddedEnum uint32

const (
	EnumMessageUNIVERSAL = EnumMessageEmbeddedEnum(0)
	EnumMessageWEB       = EnumMessageEmbeddedEnum(1)
	EnumMessageIMAGES    = EnumMessageEmbeddedEnum(2)
	EnumMessageLOCAL     = EnumMessageEmbeddedEnum(3)
	EnumMessageNEWS      = EnumMessageEmbeddedEnum(4)
	EnumMessagePRODUCTS  = EnumMessageEmbeddedEnum(5)
	EnumMessageVIDEO     = EnumMessageEmbeddedEnum(6)
)

type EnumMessage struct {
	error  error
	ignore bool

	Message  string
	Embedded EnumMessageEmbeddedEnum
}

func NewEnumMessage() *EnumMessage {
	return &EnumMessage{}
}

func (x *EnumMessage) Error(b *polyglot.Buffer, err error) {
	polyglot.Encoder(b).Error(err)
}

func (x *EnumMessage) Encode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else if x.error != nil {
		polyglot.Encoder(b).Error(x.error)
	} else {
		polyglot.Encoder(b).Bool(x.ignore).String(x.Message).Uint32(uint32(x.Embedded))
	}
}

func (x *EnumMessage) Decode(b []byte) error {
	if x == nil {
		return NilDecode
	}
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return x.decode(d)
}

func (x *EnumMessage) decode(d *polyglot.Decoder) error {
	if d.Nil() {
		return nil
	}
	var err error
	x.error, err = d.Error()
	if err != nil {
		x.ignore, err = d.Bool()
		if err != nil {
			return err
		}
		x.Message, err = d.String()
		if err != nil {
			return err
		}
		var EmbeddedTemp uint32
		EmbeddedTemp, err = d.Uint32()
		x.Embedded = EnumMessageEmbeddedEnum(EmbeddedTemp)
		if err != nil {
			return err
		}
	}
	return nil
}

type EmbeddedMessage struct {
	error  error
	ignore bool

	EmbeddedData []byte
}

func NewEmbeddedMessage() *EmbeddedMessage {
	return &EmbeddedMessage{}
}

func (x *EmbeddedMessage) Error(b *polyglot.Buffer, err error) {
	polyglot.Encoder(b).Error(err)
}

func (x *EmbeddedMessage) Encode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else if x.error != nil {
		polyglot.Encoder(b).Error(x.error)
	} else {
		polyglot.Encoder(b).Bool(x.ignore).Bytes(x.EmbeddedData)
	}
}

func (x *EmbeddedMessage) Decode(b []byte) error {
	if x == nil {
		return NilDecode
	}
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return x.decode(d)
}

func (x *EmbeddedMessage) decode(d *polyglot.Decoder) error {
	if d.Nil() {
		return nil
	}
	var err error
	x.error, err = d.Error()
	if err != nil {
		x.ignore, err = d.Bool()
		if err != nil {
			return err
		}
		x.EmbeddedData, err = d.Bytes(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

type MapBuiltMapMap map[string]float64

func NewMapBuiltMapMap(size uint32) map[string]float64 {
	return make(map[string]float64, size)
}

func (x MapBuiltMapMap) Encode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else {
		polyglot.Encoder(b).Map(uint32(len(x)), polyglot.StringKind, polyglot.Float64Kind)
		for k, v := range x {
			polyglot.Encoder(b).String(k)
			polyglot.Encoder(b).Float64(v)
		}
	}
}

func (x MapBuiltMapMap) decode(d *polyglot.Decoder, size uint32) error {
	if size == 0 {
		return nil
	}
	var k string
	var v float64
	var err error
	for i := uint32(0); i < size; i++ {
		k, err = d.String()
		if err != nil {
			return err
		}
		v, err = d.Float64()
		if err != nil {
			return err
		}
		x[k] = v
	}
	return nil
}

type Map struct {
	error  error
	ignore bool

	BuiltMap MapBuiltMapMap
}

func NewMap() *Map {
	return &Map{}
}

func (x *Map) Error(b *polyglot.Buffer, err error) {
	polyglot.Encoder(b).Error(err)
}

func (x *Map) Encode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else if x.error != nil {
		polyglot.Encoder(b).Error(x.error)
	} else {
		polyglot.Encoder(b).Bool(x.ignore)
		x.BuiltMap.Encode(b)
	}
}

func (x *Map) Decode(b []byte) error {
	if x == nil {
		return NilDecode
	}
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return x.decode(d)
}

func (x *Map) decode(d *polyglot.Decoder) error {
	if d.Nil() {
		return nil
	}
	var err error
	x.error, err = d.Error()
	if err != nil {
		x.ignore, err = d.Bool()
		if err != nil {
			return err
		}
		if !d.Nil() {
			BuiltMapSize, err := d.Map(polyglot.StringKind, polyglot.Float64Kind)
			if err != nil {
				return err
			}
			x.BuiltMap = NewMapBuiltMapMap(BuiltMapSize)
			err = x.BuiltMap.decode(d, BuiltMapSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Benchmark struct {
	error  error
	ignore bool

	Message     string
	Embedded    *EmbeddedMessage
	EmbeddedMap *Map
	EnumMessage *EnumMessage
	Repeated    []*EmbeddedMessage
}

func NewBenchmark() *Benchmark {
	return &Benchmark{
		Embedded:    NewEmbeddedMessage(),
		EmbeddedMap: NewMap(),
		EnumMessage: NewEnumMessage(),
	}
}

func (x *Benchmark) Error(b *polyglot.Buffer, err error) {
	polyglot.Encoder(b).Error(err)
}

func (x *Benchmark) Encode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else if x.error != nil {
		polyglot.Encoder(b).Error(x.error)
	} else {
		polyglot.Encoder(b).Bool(x.ignore).String(x.Message)
		polyglot.Encoder(b).Slice(uint32(len(x.Repeated)), polyglot.AnyKind)
		for _, v := range x.Repeated {
			v.Encode(b)
		}

		x.Embedded.Encode(b)
		x.EmbeddedMap.Encode(b)
		x.EnumMessage.Encode(b)
	}
}

func (x *Benchmark) Decode(b []byte) error {
	if x == nil {
		return NilDecode
	}
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return x.decode(d)
}

func (x *Benchmark) decode(d *polyglot.Decoder) error {
	if d.Nil() {
		return nil
	}
	var err error
	x.error, err = d.Error()
	if err != nil {
		x.ignore, err = d.Bool()
		if err != nil {
			return err
		}
		x.Message, err = d.String()
		if err != nil {
			return err
		}
		var sliceSize uint32
		sliceSize, err = d.Slice(polyglot.AnyKind)
		if err != nil {
			return err
		}
		if uint32(len(x.Repeated)) != sliceSize {
			x.Repeated = make([]*EmbeddedMessage, sliceSize)
		}
		for i := uint32(0); i < sliceSize; i++ {
			if x.Repeated[i] == nil {
				x.Repeated[i] = NewEmbeddedMessage()
			}
			err = x.Repeated[i].decode(d)
			if err != nil {
				return err
			}
		}
		if x.Embedded == nil {
			x.Embedded = NewEmbeddedMessage()
		}
		err = x.Embedded.decode(d)
		if err != nil {
			return err
		}
		if x.EmbeddedMap == nil {
			x.EmbeddedMap = NewMap()
		}
		err = x.EmbeddedMap.decode(d)
		if err != nil {
			return err
		}
		if x.EnumMessage == nil {
			x.EnumMessage = NewEnumMessage()
		}
		err = x.EnumMessage.decode(d)
		if err != nil {
			return err
		}
	}
	return nil
}
