package cpool

// ConstBase contains commonality across all constants in the constant pool
type ConstBase struct {
	// Index is the location in the constant pool for this constant
	Index uint16
	// Pool contains the constant pool for this constant and is used to lookup other constants in String()
	Pool  ConstantPool
	// Tag is the constant pool tag which designates the type of constant
	Tag   uint8
	// Type is the java spec type name
	Type  string
}
