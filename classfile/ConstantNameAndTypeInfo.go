package classfile

/**
	u1 tag;
    u2 name_index;
    u2 descriptor_index;
 */

type ConstantNameAndTypeIndex struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeIndex) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	self.nameIndex = classFileReader.ReadUint16()
	self.descriptorIndex = classFileReader.ReadUint16()
}
