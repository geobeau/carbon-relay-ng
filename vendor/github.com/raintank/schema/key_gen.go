package schema

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Key) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes((z)[:])
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Key) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes((z)[:])
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Key) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (z)[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Key) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (z)[:])
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Key) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (16 * (msgp.ByteSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MKey) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			err = dc.ReadExactBytes((z.Key)[:])
			if err != nil {
				return
			}
		case "Org":
			z.Org, err = dc.ReadUint32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MKey) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Key"
	err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Key)[:])
	if err != nil {
		return
	}
	// write "Org"
	err = en.Append(0xa3, 0x4f, 0x72, 0x67)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Org)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MKey) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Key"
	o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendBytes(o, (z.Key)[:])
	// string "Org"
	o = append(o, 0xa3, 0x4f, 0x72, 0x67)
	o = msgp.AppendUint32(o, z.Org)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			bts, err = msgp.ReadExactBytes(bts, (z.Key)[:])
			if err != nil {
				return
			}
		case "Org":
			z.Org, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MKey) Msgsize() (s int) {
	s = 1 + 4 + msgp.ArrayHeaderSize + (16 * (msgp.ByteSize)) + 4 + msgp.Uint32Size
	return
}
