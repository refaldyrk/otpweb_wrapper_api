package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

func (w *Wrapper) GetOperator() (model.GetOperator, error) {
	getOperator, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=operators", w.APIkey))

	if err != nil {
		return model.GetOperator{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getOperator.Body)
	var data model.GetOperator

	body, err := io.ReadAll(getOperator.Body)
	if err != nil {
		return model.GetOperator{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return model.GetOperator{}, err
	}

	return data, nil
}
