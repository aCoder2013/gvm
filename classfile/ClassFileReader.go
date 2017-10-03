package classfile

import (
	"encoding/binary"
	"fmt"
	"unicode/utf16"
)

type ClassFileReader struct {
	data []byte
}

func (self *ClassFileReader) ReadUint8() uint8 {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *ClassFileReader) ReadUint16() uint16 {
	b := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return b
}

func (self *ClassFileReader) ReadUint32() uint32 {
	b := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return b
}

func (self *ClassFileReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassFileReader) ReadUint16s() []uint16 {
	n := self.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.ReadUint16()
	}
	return s
}

func (self *ClassFileReader) ReadBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}

/**
	decode MUTF-8 :mutf8 -> utf16 -> utf32 -> string
 */
func (self *ClassFileReader) ReadUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
