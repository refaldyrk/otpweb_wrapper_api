package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

func (w *Wrapper) GetService(countryID int) (model.GetService, error) {
	getService, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_service&country_id=%d", w.APIkey, countryID))

	if err != nil {
		return model.GetService{}, err
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
		return model.GetService{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return model.GetService{}, err
	}

	return data, nil
}

func (w *Wrapper) GetSpecialService() (model.GetSpecialService, error) {
	getSpecialService, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=get_spesialService", w.APIkey))

	if err != nil {
		return model.GetSpecialService{}, err
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
		return model.GetSpecialService{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return model.GetSpecialService{}, err
	}

	return data, nil
}
