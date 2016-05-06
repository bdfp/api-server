package server

//HTTPErrorResponse dd
type HTTPErrorResponse struct {
	Err    bool   `json:"err"`
	Reason string `json:"reason"`
}
