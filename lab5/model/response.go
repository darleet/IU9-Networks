package model

type Response struct {
	Bin string `json:"bin"`
	Hex string `json:"hex"`
}

type ErrResponse struct {
	Err string `json:"error"`
}
