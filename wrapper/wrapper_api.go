package wrapper

type Wrapper struct {
	APIkey string
}

// NewWrapper membuat dan mengembalikan instans baru dari Wrapper dengan menginisialisasi API key.
// Fungsi ini memerlukan satu parameter yaitu apikey (kunci API yang akan digunakan).
// Fungsi ini mengembalikan pointer ke Wrapper yang telah diinisialisasi dengan API key yang diberikan.
func NewWrapper(apikey string) *Wrapper {
	return &Wrapper{APIkey: apikey}
}
