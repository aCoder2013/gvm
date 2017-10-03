package classfile

type ClassFile struct {
	magic           uint32
	minorVersion    uint16
	majorVersion    uint16
	constantPool    ConstantPool
	accessFlags     uint16
	thisClass       uint16
	superClass      uint16
	interfacesCount uint16
	interfaces      uint16
	fieldsCount     uint16
	fields          MemberInfo
	methodsCount    uint16
	methods         MemberInfo
	attributesCount uint16
	attributes      AttributeInfo
}

/**
	Constant pool related data structure
 */
type ConstantPool struct {
	constantPoolCount uint16
	constantInfoList  []ConstantClassInfo
}

/**
	field_info && method_info
 */
type MemberInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []AttributeInfo
}

type AttributeInfo struct {
	attribute_name_index uint16
	attribute_length     uint32
	info                 []uint8
}
