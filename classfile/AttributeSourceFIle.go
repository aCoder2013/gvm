package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	cp              *ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassFileReader) {
	self.sourceFileIndex = reader.ReadUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(int(self.sourceFileIndex))
}
