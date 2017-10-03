package classfile

import "fmt"

const JAVA_CLASSFILE_MAGIC = 0xCAFEBABE

const JAVA_MIN_SUPPORTED_VERSION = 45

const JAVA_MAX_SUPPORTED_VERSION = 53

const JAVA_5_VERSION = 49;

const JAVA_6_VERSION = 50;

const JAVA_7_VERSION = 51;

const JAVA_8_VERSION = 51;

const JAVA_9_VERSION = 53;

func Parse(data []byte) *ClassFile {
	classFile := ClassFile{}
	classFileReader := ClassFileReader{data}
	parseMagic(&classFile, &classFileReader)
	parseVersion(&classFileReader, &classFile)
	parseConstantPool(&classFileReader, &classFile)
	return &classFile
}

func parseMagic(classFile *ClassFile, classFileReader *ClassFileReader) {
	classFile.magic = classFileReader.ReadUint32()
}

func parseVersion(classFileReader *ClassFileReader, classFile *ClassFile) {
	minorVersion := classFileReader.ReadUint16()
	majorVersion := classFileReader.ReadUint16()
	classFile.minorVersion = minorVersion
	classFile.majorVersion = majorVersion
}

func parseConstantPool(classFileReader *ClassFileReader, classFile *ClassFile) {
	constantPoolCount := int(classFileReader.ReadUint16())
	if constantPoolCount <= 1 {
		panic("Illegal constant pool size " + string(constantPoolCount) + " .")
	}
	constantInfoList := make([]ConstantInfoReader, constantPoolCount)
	for i := 1; i < constantPoolCount; i++ {
		constantInfo := getConstantInfoByTag(classFileReader)
		constantInfo.ReadInfo(classFileReader, classFile)
		constantInfoList[i] = constantInfo
	}
	fmt.Println(constantInfoList)
}

func getConstantInfoByTag(classFileReader *ClassFileReader) ConstantInfoReader {
	tag := classFileReader.ReadUint8()
	switch tag {
	case JVM_CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case JVM_CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case JVM_CONSTANT_Float:
		return &ConstantFloatInfo{}
	case JVM_CONSTANT_Long:
		return &ConstantLongInfo{}
	case JVM_CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case JVM_CONSTANT_Class:
		return &ConstantClassInfo{}
	case JVM_CONSTANT_String:
		return &ConstantStringInfo{}
	case JVM_CONSTANT_Fieldref:
		return &ConstantMemberRefInfo{}
	case JVM_CONSTANT_Methodref:
		return &ConstantMemberRefInfo{}
	case JVM_CONSTANT_InterfaceMethodref:
		return &ConstantMemberRefInfo{}
	case JVM_CONSTANT_NameAndType:
		return &ConstantNameAndTypeIndex{}
	case JVM_CONSTANT_MethodHandle: //todo
	case JVM_CONSTANT_MethodType: //todo
	case JVM_CONSTANT_InvokeDynamic: //todo
	}
	panic("java.lang.ClassFormatError: unknown constant pool tag ")
}
