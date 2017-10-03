package classfile

/**
	 u1 tag;
     u2 name_index;
 */

type ConstantClassInfo struct {
	nameIndex uint16
}

func (self *ConstantClassInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	self.nameIndex = classFileReader.ReadUint16()
}
