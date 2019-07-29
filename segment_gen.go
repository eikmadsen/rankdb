package rankdb

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Segment) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 8 {
		err = msgp.ArrayError{Wanted: 8, Got: zb0001}
		return
	}
	{
		var zb0002 uint32
		zb0002, err = dc.ReadUint32()
		if err != nil {
			err = msgp.WrapError(err, "ID")
			return
		}
		z.ID = SegmentID(zb0002)
	}
	z.Min, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "Min")
		return
	}
	z.Max, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "Max")
		return
	}
	z.MinTie, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "MinTie")
		return
	}
	z.MaxTie, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "MaxTie")
		return
	}
	z.N, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "N")
		return
	}
	z.Updated, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "Updated")
		return
	}
	err = z.Parent.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Parent")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Segment) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 8
	err = en.Append(0x98)
	if err != nil {
		return
	}
	err = en.WriteUint32(uint32(z.ID))
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	err = en.WriteUint64(z.Min)
	if err != nil {
		err = msgp.WrapError(err, "Min")
		return
	}
	err = en.WriteUint64(z.Max)
	if err != nil {
		err = msgp.WrapError(err, "Max")
		return
	}
	err = en.WriteUint32(z.MinTie)
	if err != nil {
		err = msgp.WrapError(err, "MinTie")
		return
	}
	err = en.WriteUint32(z.MaxTie)
	if err != nil {
		err = msgp.WrapError(err, "MaxTie")
		return
	}
	err = en.WriteInt(z.N)
	if err != nil {
		err = msgp.WrapError(err, "N")
		return
	}
	err = en.WriteInt64(z.Updated)
	if err != nil {
		err = msgp.WrapError(err, "Updated")
		return
	}
	err = z.Parent.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Parent")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Segment) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 8
	o = append(o, 0x98)
	o = msgp.AppendUint32(o, uint32(z.ID))
	o = msgp.AppendUint64(o, z.Min)
	o = msgp.AppendUint64(o, z.Max)
	o = msgp.AppendUint32(o, z.MinTie)
	o = msgp.AppendUint32(o, z.MaxTie)
	o = msgp.AppendInt(o, z.N)
	o = msgp.AppendInt64(o, z.Updated)
	o, err = z.Parent.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Parent")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Segment) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 8 {
		err = msgp.ArrayError{Wanted: 8, Got: zb0001}
		return
	}
	{
		var zb0002 uint32
		zb0002, bts, err = msgp.ReadUint32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "ID")
			return
		}
		z.ID = SegmentID(zb0002)
	}
	z.Min, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Min")
		return
	}
	z.Max, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Max")
		return
	}
	z.MinTie, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MinTie")
		return
	}
	z.MaxTie, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MaxTie")
		return
	}
	z.N, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "N")
		return
	}
	z.Updated, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Updated")
		return
	}
	bts, err = z.Parent.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Parent")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Segment) Msgsize() (s int) {
	s = 1 + msgp.Uint32Size + msgp.Uint64Size + msgp.Uint64Size + msgp.Uint32Size + msgp.Uint32Size + msgp.IntSize + msgp.Int64Size + z.Parent.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SegmentID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint32
		zb0001, err = dc.ReadUint32()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = SegmentID(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z SegmentID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint32(uint32(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SegmentID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint32(o, uint32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SegmentID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint32
		zb0001, bts, err = msgp.ReadUint32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = SegmentID(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SegmentID) Msgsize() (s int) {
	s = msgp.Uint32Size
	return
}
