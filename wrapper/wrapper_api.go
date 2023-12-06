package wrapper

type Wrapper struct {
	APIkey string
}

func NewWrapper(apikey string) *Wrapper {
	return &Wrapper{APIkey: apikey}
}
