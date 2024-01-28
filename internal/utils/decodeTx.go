package utils

import (
	"encoding/json"
	"net/http"
)

type TxResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Value   struct {
		Uuid      string `json:"uuid"`
		Block     string `json:"block"`
		Validator string `json:"validator"`
		Status    string `json:"status"`
		Info      string `json:"info"`
		Txs       []struct {
			TxHash       string `json:"tx_hash"`
			Status       string `json:"status"`
			RevertMsg    string `json:"revert_msg"`
			AcceptRevert bool   `json:"accept_revert"`
			Created      int    `json:"created"`
		} `json:"txs"`
		Created int `json:"created"`
	} `json:"value"`
}

func GetTxResponse(hash string) (TxResponse, error) {
	url := "https://explorer.48.club/api/v1/puissant/" + hash
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TxResponse{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TxResponse{}, err
	}
	defer resp.Body.Close()

	var response TxResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return TxResponse{}, err
	}

	return response, nil
}
