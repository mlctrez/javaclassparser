package cpool

// TODO: transition constants pool to use proper error handling
func failErr(err error) {
	if err != nil {
		panic(err)
	}
}
