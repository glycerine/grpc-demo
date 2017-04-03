package api

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

// DecodeMsg implements msgp.Decodable
func (z *BcastGetReply) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FromID":
			z.FromID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Ki":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Ki = nil
			} else {
				if z.Ki == nil {
					z.Ki = new(KeyInv)
				}
				err = z.Ki.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Err":
			z.Err, err = dc.ReadString()
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
func (z *BcastGetReply) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "FromID"
	err = en.Append(0x83, 0xa6, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FromID)
	if err != nil {
		return
	}
	// write "Ki"
	err = en.Append(0xa2, 0x4b, 0x69)
	if err != nil {
		return err
	}
	if z.Ki == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Ki.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Err"
	err = en.Append(0xa3, 0x45, 0x72, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Err)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BcastGetReply) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "FromID"
	o = append(o, 0x83, 0xa6, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44)
	o = msgp.AppendString(o, z.FromID)
	// string "Ki"
	o = append(o, 0xa2, 0x4b, 0x69)
	if z.Ki == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Ki.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Err"
	o = append(o, 0xa3, 0x45, 0x72, 0x72)
	o = msgp.AppendString(o, z.Err)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BcastGetReply) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FromID":
			z.FromID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Ki":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Ki = nil
			} else {
				if z.Ki == nil {
					z.Ki = new(KeyInv)
				}
				bts, err = z.Ki.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Err":
			z.Err, bts, err = msgp.ReadStringBytes(bts)
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
func (z *BcastGetReply) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.FromID) + 3
	if z.Ki == nil {
		s += msgp.NilSize
	} else {
		s += z.Ki.Msgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Err)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BcastGetRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FromID":
			z.FromID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Key":
			z.Key, err = dc.ReadBytes(z.Key)
			if err != nil {
				return
			}
		case "Who":
			z.Who, err = dc.ReadString()
			if err != nil {
				return
			}
		case "IncludeValue":
			z.IncludeValue, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "ReplyGrpcHost":
			z.ReplyGrpcHost, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ReplyGrpcXPort":
			z.ReplyGrpcXPort, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "ReplyGrpcIPort":
			z.ReplyGrpcIPort, err = dc.ReadInt()
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
func (z *BcastGetRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "FromID"
	err = en.Append(0x87, 0xa6, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FromID)
	if err != nil {
		return
	}
	// write "Key"
	err = en.Append(0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Key)
	if err != nil {
		return
	}
	// write "Who"
	err = en.Append(0xa3, 0x57, 0x68, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Who)
	if err != nil {
		return
	}
	// write "IncludeValue"
	err = en.Append(0xac, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.IncludeValue)
	if err != nil {
		return
	}
	// write "ReplyGrpcHost"
	err = en.Append(0xad, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x70, 0x63, 0x48, 0x6f, 0x73, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyGrpcHost)
	if err != nil {
		return
	}
	// write "ReplyGrpcXPort"
	err = en.Append(0xae, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x70, 0x63, 0x58, 0x50, 0x6f, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ReplyGrpcXPort)
	if err != nil {
		return
	}
	// write "ReplyGrpcIPort"
	err = en.Append(0xae, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x70, 0x63, 0x49, 0x50, 0x6f, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ReplyGrpcIPort)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BcastGetRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "FromID"
	o = append(o, 0x87, 0xa6, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44)
	o = msgp.AppendString(o, z.FromID)
	// string "Key"
	o = append(o, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendBytes(o, z.Key)
	// string "Who"
	o = append(o, 0xa3, 0x57, 0x68, 0x6f)
	o = msgp.AppendString(o, z.Who)
	// string "IncludeValue"
	o = append(o, 0xac, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendBool(o, z.IncludeValue)
	// string "ReplyGrpcHost"
	o = append(o, 0xad, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x70, 0x63, 0x48, 0x6f, 0x73, 0x74)
	o = msgp.AppendString(o, z.ReplyGrpcHost)
	// string "ReplyGrpcXPort"
	o = append(o, 0xae, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x70, 0x63, 0x58, 0x50, 0x6f, 0x72, 0x74)
	o = msgp.AppendInt(o, z.ReplyGrpcXPort)
	// string "ReplyGrpcIPort"
	o = append(o, 0xae, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x70, 0x63, 0x49, 0x50, 0x6f, 0x72, 0x74)
	o = msgp.AppendInt(o, z.ReplyGrpcIPort)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BcastGetRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FromID":
			z.FromID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Key":
			z.Key, bts, err = msgp.ReadBytesBytes(bts, z.Key)
			if err != nil {
				return
			}
		case "Who":
			z.Who, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "IncludeValue":
			z.IncludeValue, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "ReplyGrpcHost":
			z.ReplyGrpcHost, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ReplyGrpcXPort":
			z.ReplyGrpcXPort, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ReplyGrpcIPort":
			z.ReplyGrpcIPort, bts, err = msgp.ReadIntBytes(bts)
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
func (z *BcastGetRequest) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.FromID) + 4 + msgp.BytesPrefixSize + len(z.Key) + 4 + msgp.StringPrefixSize + len(z.Who) + 13 + msgp.BoolSize + 14 + msgp.StringPrefixSize + len(z.ReplyGrpcHost) + 15 + msgp.IntSize + 15 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BcastSetReply) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Err":
			z.Err, err = dc.ReadString()
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
func (z BcastSetReply) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Err"
	err = en.Append(0x81, 0xa3, 0x45, 0x72, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Err)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BcastSetReply) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Err"
	o = append(o, 0x81, 0xa3, 0x45, 0x72, 0x72)
	o = msgp.AppendString(o, z.Err)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BcastSetReply) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Err":
			z.Err, bts, err = msgp.ReadStringBytes(bts)
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
func (z BcastSetReply) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Err)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BcastSetRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FromID":
			z.FromID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Ki":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Ki = nil
			} else {
				if z.Ki == nil {
					z.Ki = new(KeyInv)
				}
				err = z.Ki.DecodeMsg(dc)
				if err != nil {
					return
				}
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
func (z *BcastSetRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "FromID"
	err = en.Append(0x82, 0xa6, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FromID)
	if err != nil {
		return
	}
	// write "Ki"
	err = en.Append(0xa2, 0x4b, 0x69)
	if err != nil {
		return err
	}
	if z.Ki == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Ki.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BcastSetRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "FromID"
	o = append(o, 0x82, 0xa6, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44)
	o = msgp.AppendString(o, z.FromID)
	// string "Ki"
	o = append(o, 0xa2, 0x4b, 0x69)
	if z.Ki == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Ki.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BcastSetRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "FromID":
			z.FromID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Ki":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Ki = nil
			} else {
				if z.Ki == nil {
					z.Ki = new(KeyInv)
				}
				bts, err = z.Ki.UnmarshalMsg(bts)
				if err != nil {
					return
				}
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
func (z *BcastSetRequest) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.FromID) + 3
	if z.Ki == nil {
		s += msgp.NilSize
	} else {
		s += z.Ki.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *KeyInv) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, err = dc.ReadBytes(z.Key)
			if err != nil {
				return
			}
		case "Who":
			z.Who, err = dc.ReadString()
			if err != nil {
				return
			}
		case "When":
			z.When, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "Size":
			z.Size, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Blake2b":
			z.Blake2b, err = dc.ReadBytes(z.Blake2b)
			if err != nil {
				return
			}
		case "Val":
			z.Val, err = dc.ReadBytes(z.Val)
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
func (z *KeyInv) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Key"
	err = en.Append(0x86, 0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Key)
	if err != nil {
		return
	}
	// write "Who"
	err = en.Append(0xa3, 0x57, 0x68, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Who)
	if err != nil {
		return
	}
	// write "When"
	err = en.Append(0xa4, 0x57, 0x68, 0x65, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.When)
	if err != nil {
		return
	}
	// write "Size"
	err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Size)
	if err != nil {
		return
	}
	// write "Blake2b"
	err = en.Append(0xa7, 0x42, 0x6c, 0x61, 0x6b, 0x65, 0x32, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Blake2b)
	if err != nil {
		return
	}
	// write "Val"
	err = en.Append(0xa3, 0x56, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Val)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *KeyInv) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Key"
	o = append(o, 0x86, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendBytes(o, z.Key)
	// string "Who"
	o = append(o, 0xa3, 0x57, 0x68, 0x6f)
	o = msgp.AppendString(o, z.Who)
	// string "When"
	o = append(o, 0xa4, 0x57, 0x68, 0x65, 0x6e)
	o = msgp.AppendTime(o, z.When)
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.Size)
	// string "Blake2b"
	o = append(o, 0xa7, 0x42, 0x6c, 0x61, 0x6b, 0x65, 0x32, 0x62)
	o = msgp.AppendBytes(o, z.Blake2b)
	// string "Val"
	o = append(o, 0xa3, 0x56, 0x61, 0x6c)
	o = msgp.AppendBytes(o, z.Val)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KeyInv) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, bts, err = msgp.ReadBytesBytes(bts, z.Key)
			if err != nil {
				return
			}
		case "Who":
			z.Who, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "When":
			z.When, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Size":
			z.Size, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Blake2b":
			z.Blake2b, bts, err = msgp.ReadBytesBytes(bts, z.Blake2b)
			if err != nil {
				return
			}
		case "Val":
			z.Val, bts, err = msgp.ReadBytesBytes(bts, z.Val)
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
func (z *KeyInv) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Key) + 4 + msgp.StringPrefixSize + len(z.Who) + 5 + msgp.TimeSize + 5 + msgp.Int64Size + 8 + msgp.BytesPrefixSize + len(z.Blake2b) + 4 + msgp.BytesPrefixSize + len(z.Val)
	return
}
