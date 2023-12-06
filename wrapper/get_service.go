package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

// GetService mengambil daftar layanan dari layanan OTPWeb menggunakan kunci API yang diberikan dan ID negara.
// Fungsi ini memerlukan satu parameter yaitu countryID (ID negara).
// Fungsi ini mengembalikan struktur data model.GetService yang berisi daftar layanan yang diperoleh,
// dan struktur data model.Error jika terjadi kesalahan selama pengambilan daftar layanan.
func (w *Wrapper) GetService(countryID int) (model.GetService, model.Error) {
	getService, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_service&country_id=%d", w.APIkey, countryID))

	if err != nil {
		fmt.Println(err)
		return model.GetService{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getService.Body)
	var data model.GetService

	body, err := io.ReadAll(getService.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetService{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.GetService{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetService{}, errorResp
	}

	return data, model.Error{}
}

// GetSpecialService mengambil daftar layanan khusus dari layanan OTPWeb menggunakan kunci API yang diberikan.
// Fungsi ini tidak memerlukan parameter tambahan.
// Fungsi ini mengembalikan struktur data model.GetSpecialService yang berisi daftar layanan khusus yang diperoleh,
// dan struktur data model.Error jika terjadi kesalahan selama pengambilan daftar layanan khusus.
func (w *Wrapper) GetSpecialService() (model.GetSpecialService, model.Error) {
	getSpecialService, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_spesialService", w.APIkey))

	if err != nil {
		fmt.Println(err)
		return model.GetSpecialService{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getSpecialService.Body)
	var data model.GetSpecialService

	body, err := io.ReadAll(getSpecialService.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetSpecialService{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.GetSpecialService{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetSpecialService{}, errorResp
	}

	return data, model.Error{}
}
