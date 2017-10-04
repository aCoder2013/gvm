package classfile

import "strings"

type ClassFile struct {
	magic           uint32
	minorVersion    uint16
	majorVersion    uint16
	constantPool    ConstantPool
	accessFlags     uint16
	thisClass       uint16
	superClass      uint16
	interfacesCount uint16
	interfaces      []uint16
	fieldsCount     uint16
	fields          []MemberInfo
	methodsCount    uint16
	methods         []MemberInfo
	attributesCount uint16
	attributes      AttributeInfo
}

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
)

/**
	Constant pool related data structure
 */
type ConstantPool struct {
	constantInfoList []ConstantInfoReader
}

func (cp *ConstantPool) getUtf8(index int) string {
	return cp.constantInfoList[index].(*ConstantUtf8Info).value
}

func (self *ConstantPool) getClassName(index uint16) string {
	classInfo := self.constantInfoList[index].(*ConstantClassInfo)
	return self.getUtf8(int(classInfo.nameIndex))
}

func (self *ConstantPool) getNameAndType(index uint16) (name, _type string) {
	ntInfo := self.constantInfoList[index].(*ConstantNameAndTypeIndex)
	name = self.getUtf8(int(ntInfo.nameIndex))
	_type = self.getUtf8(int(ntInfo.descriptorIndex))
	return
}

/**
	field_info && method_info
 */
type MemberInfo struct {
	cp              *ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []AttributeInfo
}

func (self *MemberInfo) String() string {
	return strings.Join([]string{self.cp.getUtf8(int(self.nameIndex)), self.cp.getUtf8(int(self.descriptorIndex))}, ",")
}

/*
	Attribute info
 */
type AttributeInfo interface {
	readInfo(reader *ClassFileReader)
}

type AttributeDeprecatedInfo struct {
	cp    *ConstantPool
	value string
}

func (self *AttributeDeprecatedInfo) readInfo(reader *ClassFileReader) {
	self.value = self.cp.getUtf8(int(reader.ReadUint16()))
}

type AttributeRuntimeVisibleAnnotationsInfo struct {
	attribute_name string
}

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	cp             *ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	AttributeTable
}

func (self *CodeAttribute) readInfo(reader *ClassFileReader) {
	self.maxStack = reader.ReadUint16()
	self.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	self.code = reader.ReadBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readAttributes(reader *ClassFileReader, cp *ConstantPool) []AttributeInfo {
	attributesCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassFileReader, cp *ConstantPool) AttributeInfo {
	attrNameIndex := reader.ReadUint16()
	attrLen := reader.ReadUint32()
	attrName := cp.getUtf8(int(attrNameIndex))
	attrInfo := newAttributeInfo(attrName, cp)
	if attrInfo == nil {
		attrInfo = &UnparsedAttribute{
			name:   attrName,
			length: attrLen,
		}
	}
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, cp *ConstantPool) AttributeInfo {
	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
		// case "MethodParameters":
		// case "RuntimeInvisibleAnnotations":
		// case "RuntimeInvisibleParameterAnnotations":
		// case "RuntimeInvisibleTypeAnnotations":
		// case "RuntimeVisibleAnnotations":
		// case "RuntimeVisibleParameterAnnotations":
		// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
		// case "SourceDebugExtension":
		// case "StackMapTable":
	case "Synthetic":
		return _attrSynthetic
	default:
		return nil // undefined attr
	}
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassFileReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}
	return exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
