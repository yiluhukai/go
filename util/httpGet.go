package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"yiluhuakai/logger"
	"yiluhuakai/pop/model"
)

func HttpGet(url string) (result *model.ResTenXun, err error) {
	result = new(model.ResTenXun)
	resp, err := http.Get(url)
	if err != nil {
		logger.LogError("requerst %v failed:%v", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError("parse response body failed:%v", err)
		return
	}
	err = json.Unmarshal(body, result)

	return
}
