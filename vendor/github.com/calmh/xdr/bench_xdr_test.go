// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package xdr_test

import (
	"github.com/calmh/xdr"
)

/*

XDRBenchStruct Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                         I1 (64 bits)                          +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                              I2                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|         16 zero bits          |              I3               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                 24 zero bits                  |      I4       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  Bs0 (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  Bs1 (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Number of Is0                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
|                         Is0 (n items)                         |
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   S0 (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   S1 (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct XDRBenchStruct {
	unsigned hyper I1;
	unsigned int I2;
	unsigned int I3;
	unsigned int I4;
	opaque Bs0<128>;
	opaque Bs1<>;
	int Is0<>;
	string S0<128>;
	string S1<>;
}

*/

func (o XDRBenchStruct) XDRSize() int {
	return 8 + 4 + 4 + 4 +
		4 + len(o.Bs0) + xdr.Padding(len(o.Bs0)) +
		4 + len(o.Bs1) + xdr.Padding(len(o.Bs1)) +
		4 + len(o.Is0)*4 +
		4 + len(o.S0) + xdr.Padding(len(o.S0)) +
		4 + len(o.S1) + xdr.Padding(len(o.S1))
}

func (o XDRBenchStruct) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o XDRBenchStruct) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o XDRBenchStruct) MarshalXDRInto(m *xdr.Marshaller) error {
	m.MarshalUint64(o.I1)
	m.MarshalUint32(o.I2)
	m.MarshalUint16(o.I3)
	m.MarshalUint8(o.I4)
	if l := len(o.Bs0); l > 128 {
		return xdr.ElementSizeExceeded("Bs0", l, 128)
	}
	m.MarshalBytes(o.Bs0)
	m.MarshalBytes(o.Bs1)
	m.MarshalUint32(uint32(len(o.Is0)))
	for i := range o.Is0 {
		m.MarshalUint32(uint32(o.Is0[i]))
	}
	if l := len(o.S0); l > 128 {
		return xdr.ElementSizeExceeded("S0", l, 128)
	}
	m.MarshalString(o.S0)
	m.MarshalString(o.S1)
	return m.Error
}

func (o *XDRBenchStruct) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *XDRBenchStruct) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.I1 = u.UnmarshalUint64()
	o.I2 = u.UnmarshalUint32()
	o.I3 = u.UnmarshalUint16()
	o.I4 = u.UnmarshalUint8()
	o.Bs0 = u.UnmarshalBytesMax(128)
	o.Bs1 = u.UnmarshalBytes()
	_Is0Size := int(u.UnmarshalUint32())
	if _Is0Size < 0 {
		return xdr.ElementSizeExceeded("Is0", _Is0Size, 0)
	} else if _Is0Size == 0 {
		o.Is0 = nil
	} else {
		if _Is0Size <= len(o.Is0) {
			o.Is0 = o.Is0[:_Is0Size]
		} else {
			o.Is0 = make([]int32, _Is0Size)
		}
		for i := range o.Is0 {
			o.Is0[i] = int32(u.UnmarshalUint32())
		}
	}
	o.S0 = u.UnmarshalStringMax(128)
	o.S1 = u.UnmarshalString()
	return u.Error
}
