package http

import (
	"encoding/json"
	"net/http"
)

func DoJson(req *http.Request, dst interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return err
	}

	return nil
}
