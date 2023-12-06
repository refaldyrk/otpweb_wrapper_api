package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

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
