package app

// Error : error check
func Error(e error) {
	if e != nil {
		panic(e)
	}
}
