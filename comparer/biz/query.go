package biz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Query(conf ConfigAPI) ([]byte, error) {
	queryUrl := conf.Addr + "/query?db=" + conf.DB + "&q=" + url.QueryEscape(conf.SQL)

	resp, err := http.Get(queryUrl)
	if err != nil {
		var ret []byte
		return ret, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("request [%v] error: %v", queryUrl, resp)
		var ret []byte
		return ret, err
	}

	return ioutil.ReadAll(resp.Body)
}
