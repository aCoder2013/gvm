package classfile

/**
	u1 tag;
    u2 class_index;
    u2 name_and_type_index;
 */

type ConstantMemberRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberRefInfo) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	self.classIndex = classFileReader.ReadUint16()
	self.nameAndTypeIndex = classFileReader.ReadUint16()
}
