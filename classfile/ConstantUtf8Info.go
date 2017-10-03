package classfile

type ConstantUtf8Info struct {
	value string
}

func (self *ConstantUtf8Info) ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile) {
	length := classFileReader.ReadUint16()
	bytes := classFileReader.ReadBytes(uint32(length))
	value := classFileReader.ReadUTF8(bytes)
	self.value = value
}
