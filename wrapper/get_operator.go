package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

func (w *Wrapper) GetOperator() (model.GetOperator, model.Error) {
	getOperator, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=operators", w.APIkey))

	if err != nil {
		fmt.Println(err)
		return model.GetOperator{}, model.Error{}
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
		fmt.Println(err)
		return model.GetOperator{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return model.GetOperator{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetOperator{}, errorResp
	}

	return data, model.Error{}
}
