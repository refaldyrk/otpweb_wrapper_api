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
