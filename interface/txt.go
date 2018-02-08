package text

type RW interface{
	Write(s string)
	RO
}

type RO interface{
	Read()string
	Print()
}
