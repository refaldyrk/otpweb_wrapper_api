package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

func (w *Wrapper) GetNumber(countryID int, serviceID, operator string) (model.GetNumber, model.Error) {
	getNumber, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_number&country_id=%d&service_id=%s&operator=%s", w.APIkey, countryID, serviceID, operator))

	if err != nil {
		fmt.Println(err)
		return model.GetNumber{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getNumber.Body)
	var data model.GetNumber

	body, err := io.ReadAll(getNumber.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetNumber{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.GetNumber{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetNumber{}, errorResp
	}

	return data, model.Error{}
}

func (w *Wrapper) GetSpecialNumber(serviceID string) (model.GetSpecialNumber, model.Error) {
	getNumber, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_spesialNumber&service_id=%s", w.APIkey, serviceID))

	if err != nil {
		fmt.Println(err)
		return model.GetSpecialNumber{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getNumber.Body)
	var data model.GetSpecialNumber

	body, err := io.ReadAll(getNumber.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetSpecialNumber{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.GetSpecialNumber{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetSpecialNumber{}, errorResp
	}

	return data, model.Error{}
}
