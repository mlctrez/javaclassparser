package cpool

// ConstBase contains commonality across all constants in the constant pool
type ConstBase struct {
	Pool ConstantPool
	Tag  uint8
	Type string
}
