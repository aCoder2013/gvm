package classfile

type ConstantInfoReader interface {
	ReadInfo(classFileReader *ClassFileReader, classFile *ClassFile)
}
