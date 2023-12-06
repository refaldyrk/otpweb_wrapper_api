package wrapper

import (
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/otpweb_wrapper_api/model"
	"io"
	"net/http"
)

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
