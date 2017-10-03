package classfile

import (
	"math"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	bytes := classFileReader.ReadUint32()
	self.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	bytes := classFileReader.ReadUint32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	bytes := classFileReader.ReadUint64()
	self.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	bytes := classFileReader.ReadUint64()
	self.val = math.Float64frombits(bytes)
}
