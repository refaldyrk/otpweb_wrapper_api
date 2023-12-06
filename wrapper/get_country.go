package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

// GetCountry mengambil data negara dari layanan OTPWeb menggunakan kunci API yang diberikan.
// Fungsi ini mengembalikan struktur data model.GetCountry yang berisi informasi negara,
// dan struktur data model.Error jika terjadi kesalahan selama pengambilan data negara.
func (w *Wrapper) GetCountry() (model.GetCountry, model.Error) {
	getCountry, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=country", w.APIkey))

	if err != nil {
		fmt.Println(err)
		return model.GetCountry{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getCountry.Body)
	var data model.GetCountry

	body, err := io.ReadAll(getCountry.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetCountry{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return model.GetCountry{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetCountry{}, errorResp
	}

	return data, model.Error{}
}
