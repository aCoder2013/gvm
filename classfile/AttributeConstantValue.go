package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassFileReader) {
	self.constantValueIndex = reader.ReadUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
