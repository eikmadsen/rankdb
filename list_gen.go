package rankdb

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *List) DecodeMsg(dc *msgp.Reader) (err error) {
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
	err = z.ID.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	z.Set, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Set")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	if z.Metadata == nil {
		z.Metadata = make(map[string]string, zb0002)
	} else if len(z.Metadata) > 0 {
		for key := range z.Metadata {
			delete(z.Metadata, key)
		}
	}
	for zb0002 > 0 {
		zb0002--
		var za0001 string
		var za0002 string
		za0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		za0002, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0001)
			return
		}
		z.Metadata[za0001] = za0002
	}
	z.SplitSize, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "SplitSize")
		return
	}
	z.MergeSize, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "MergeSize")
		return
	}
	z.LoadIndex, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "LoadIndex")
		return
	}
	err = z.Scores.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Scores")
		return
	}
	err = z.Index.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *List) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 8
	err = en.Append(0x98)
	if err != nil {
		return
	}
	err = z.ID.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	err = en.WriteString(z.Set)
	if err != nil {
		err = msgp.WrapError(err, "Set")
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Metadata)))
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	for za0001, za0002 := range z.Metadata {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0001)
			return
		}
	}
	err = en.WriteInt(z.SplitSize)
	if err != nil {
		err = msgp.WrapError(err, "SplitSize")
		return
	}
	err = en.WriteInt(z.MergeSize)
	if err != nil {
		err = msgp.WrapError(err, "MergeSize")
		return
	}
	err = en.WriteBool(z.LoadIndex)
	if err != nil {
		err = msgp.WrapError(err, "LoadIndex")
		return
	}
	err = z.Scores.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Scores")
		return
	}
	err = z.Index.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *List) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 8
	o = append(o, 0x98)
	o, err = z.ID.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	o = msgp.AppendString(o, z.Set)
	o = msgp.AppendMapHeader(o, uint32(len(z.Metadata)))
	for za0001, za0002 := range z.Metadata {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	o = msgp.AppendInt(o, z.SplitSize)
	o = msgp.AppendInt(o, z.MergeSize)
	o = msgp.AppendBool(o, z.LoadIndex)
	o, err = z.Scores.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Scores")
		return
	}
	o, err = z.Index.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *List) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
	bts, err = z.ID.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	z.Set, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Set")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Metadata")
		return
	}
	if z.Metadata == nil {
		z.Metadata = make(map[string]string, zb0002)
	} else if len(z.Metadata) > 0 {
		for key := range z.Metadata {
			delete(z.Metadata, key)
		}
	}
	for zb0002 > 0 {
		var za0001 string
		var za0002 string
		zb0002--
		za0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
		za0002, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Metadata", za0001)
			return
		}
		z.Metadata[za0001] = za0002
	}
	z.SplitSize, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "SplitSize")
		return
	}
	z.MergeSize, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "MergeSize")
		return
	}
	z.LoadIndex, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "LoadIndex")
		return
	}
	bts, err = z.Scores.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Scores")
		return
	}
	bts, err = z.Index.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Index")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *List) Msgsize() (s int) {
	s = 1 + z.ID.Msgsize() + msgp.StringPrefixSize + len(z.Set) + msgp.MapHeaderSize
	if z.Metadata != nil {
		for za0001, za0002 := range z.Metadata {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += msgp.IntSize + msgp.IntSize + msgp.BoolSize + z.Scores.Msgsize() + z.Index.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ListStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Elements":
			z.Elements, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Elements")
				return
			}
		case "Segments":
			z.Segments, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Segments")
				return
			}
		case "Top":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Top")
					return
				}
				z.Top = nil
			} else {
				if z.Top == nil {
					z.Top = new(RankedElement)
				}
				err = z.Top.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Top")
					return
				}
			}
		case "Bottom":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Bottom")
					return
				}
				z.Bottom = nil
			} else {
				if z.Bottom == nil {
					z.Bottom = new(RankedElement)
				}
				err = z.Bottom.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Bottom")
					return
				}
			}
		case "CacheHits":
			z.CacheHits, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "CacheHits")
				return
			}
		case "CacheMisses":
			z.CacheMisses, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "CacheMisses")
				return
			}
		case "CachePct":
			z.CachePct, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "CachePct")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ListStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Elements"
	err = en.Append(0x87, 0xa8, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Elements)
	if err != nil {
		err = msgp.WrapError(err, "Elements")
		return
	}
	// write "Segments"
	err = en.Append(0xa8, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Segments)
	if err != nil {
		err = msgp.WrapError(err, "Segments")
		return
	}
	// write "Top"
	err = en.Append(0xa3, 0x54, 0x6f, 0x70)
	if err != nil {
		return
	}
	if z.Top == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Top.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Top")
			return
		}
	}
	// write "Bottom"
	err = en.Append(0xa6, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d)
	if err != nil {
		return
	}
	if z.Bottom == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Bottom.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Bottom")
			return
		}
	}
	// write "CacheHits"
	err = en.Append(0xa9, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x69, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.CacheHits)
	if err != nil {
		err = msgp.WrapError(err, "CacheHits")
		return
	}
	// write "CacheMisses"
	err = en.Append(0xab, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.CacheMisses)
	if err != nil {
		err = msgp.WrapError(err, "CacheMisses")
		return
	}
	// write "CachePct"
	err = en.Append(0xa8, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x63, 0x74)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.CachePct)
	if err != nil {
		err = msgp.WrapError(err, "CachePct")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ListStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Elements"
	o = append(o, 0x87, 0xa8, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendInt(o, z.Elements)
	// string "Segments"
	o = append(o, 0xa8, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendInt(o, z.Segments)
	// string "Top"
	o = append(o, 0xa3, 0x54, 0x6f, 0x70)
	if z.Top == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Top.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Top")
			return
		}
	}
	// string "Bottom"
	o = append(o, 0xa6, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d)
	if z.Bottom == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Bottom.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Bottom")
			return
		}
	}
	// string "CacheHits"
	o = append(o, 0xa9, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x69, 0x74, 0x73)
	o = msgp.AppendUint64(o, z.CacheHits)
	// string "CacheMisses"
	o = append(o, 0xab, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x65, 0x73)
	o = msgp.AppendUint64(o, z.CacheMisses)
	// string "CachePct"
	o = append(o, 0xa8, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x63, 0x74)
	o = msgp.AppendFloat64(o, z.CachePct)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ListStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Elements":
			z.Elements, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Elements")
				return
			}
		case "Segments":
			z.Segments, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Segments")
				return
			}
		case "Top":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Top = nil
			} else {
				if z.Top == nil {
					z.Top = new(RankedElement)
				}
				bts, err = z.Top.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Top")
					return
				}
			}
		case "Bottom":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Bottom = nil
			} else {
				if z.Bottom == nil {
					z.Bottom = new(RankedElement)
				}
				bts, err = z.Bottom.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Bottom")
					return
				}
			}
		case "CacheHits":
			z.CacheHits, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CacheHits")
				return
			}
		case "CacheMisses":
			z.CacheMisses, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CacheMisses")
				return
			}
		case "CachePct":
			z.CachePct, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CachePct")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ListStats) Msgsize() (s int) {
	s = 1 + 9 + msgp.IntSize + 9 + msgp.IntSize + 4
	if z.Top == nil {
		s += msgp.NilSize
	} else {
		s += z.Top.Msgsize()
	}
	s += 7
	if z.Bottom == nil {
		s += msgp.NilSize
	} else {
		s += z.Bottom.Msgsize()
	}
	s += 10 + msgp.Uint64Size + 12 + msgp.Uint64Size + 9 + msgp.Float64Size
	return
}
