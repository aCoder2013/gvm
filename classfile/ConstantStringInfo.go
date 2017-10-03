package classfile

/**
 	u1 tag;
    u2 string_index;
 */

type ConstantStringInfo struct {
	stringIndex uint16
}

func (self *ConstantStringInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	self.stringIndex = classFileReader.ReadUint16()
}
