package mock

type MockRetriever struct{
	Description string
}

func (t MockRetriever)Get(url string)string{
	return t.Description
}
