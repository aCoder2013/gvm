package classfile

const (
	JVM_CONSTANT_Utf8               = 1
	JVM_CONSTANT_Unicode            = 2 /* unused */
	JVM_CONSTANT_Integer            = 3
	JVM_CONSTANT_Float              = 4
	JVM_CONSTANT_Long               = 5
	JVM_CONSTANT_Double             = 6
	JVM_CONSTANT_Class              = 7
	JVM_CONSTANT_String             = 8
	JVM_CONSTANT_Fieldref           = 9
	JVM_CONSTANT_Methodref          = 10
	JVM_CONSTANT_InterfaceMethodref = 11
	JVM_CONSTANT_NameAndType        = 12
	JVM_CONSTANT_MethodHandle       = 15 // JSR 292
	JVM_CONSTANT_MethodType         = 16 // JSR 292
	//JVM_CONSTANT_(unused)             = 17  // JSR 292 early drafts only
	JVM_CONSTANT_InvokeDynamic = 18 // JSR 292
)
