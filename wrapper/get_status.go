package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

// GetStatus mengambil status pesanan dari layanan OTPWeb menggunakan kunci API yang diberikan dan ID pesanan.
// Fungsi ini memerlukan satu parameter yaitu orderID (ID pesanan).
// Fungsi ini mengembalikan struktur data model.GetStatus yang berisi informasi status pesanan yang diperoleh,
// dan struktur data model.Error jika terjadi kesalahan selama pengambilan status pesanan.
func (w *Wrapper) GetStatus(orderID string) (model.GetStatus, model.Error) {
	getStatus, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_status&order_id=%s", w.APIkey, orderID))

	if err != nil {
		fmt.Println(err)
		return model.GetStatus{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getStatus.Body)
	var data model.GetStatus

	body, err := io.ReadAll(getStatus.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetStatus{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.GetStatus{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetStatus{}, errorResp
	}

	return data, model.Error{}
}

// ChangeStatus mengubah status pesanan pada layanan OTPWeb menggunakan kunci API yang diberikan, ID pesanan, dan status baru.
// Fungsi ini memerlukan dua parameter yaitu orderID (ID pesanan) dan status (status baru yang akan diatur).
// Fungsi ini mengembalikan struktur data model.Error yang berisi informasi status perubahan yang diperoleh.
func (w *Wrapper) ChangeStatus(orderID string, status string) model.Error {
	getStatus, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=set_status&order_id=%s&status=%s", w.APIkey, orderID, status))

	if err != nil {
		fmt.Println(err)
		return model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getStatus.Body)
	var data model.Error

	body, err := io.ReadAll(getStatus.Body)
	if err != nil {
		fmt.Println(err)
		return model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return errorResp
	}

	return data
}
