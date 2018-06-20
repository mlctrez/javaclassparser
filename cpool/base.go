package cpool

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

type ConstBase struct {
	Pool ConstantPool
	Tag  uint8
	Type string
}
