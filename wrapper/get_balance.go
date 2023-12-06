package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

func (w *Wrapper) GetBalance() (model.GetBalance, model.Error) {
	getBalance, err := http.Get(fmt.Sprintf("https://otpweb.com/api?api_key=%s&action=balance", w.APIkey))
	if err != nil {
		fmt.Println(err)
		return model.GetBalance{}, model.Error{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(getBalance.Body)
	var data model.GetBalance

	body, err := io.ReadAll(getBalance.Body)
	if err != nil {
		fmt.Println(err)
		return model.GetBalance{}, model.Error{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return model.GetBalance{}, model.Error{}
	}

	if !data.Status {
		var errorResp model.Error
		json.Unmarshal(body, &errorResp)

		return model.GetBalance{}, errorResp
	}

	return data, model.Error{}
}
