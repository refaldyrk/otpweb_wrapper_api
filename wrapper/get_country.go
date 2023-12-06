package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

func (w *Wrapper) GetCountry() (model.GetCountry, error) {
	getCountry, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=country", w.APIkey))

	if err != nil {
		return model.GetCountry{}, err
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
		return model.GetCountry{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return model.GetCountry{}, err
	}

	return data, nil
}
