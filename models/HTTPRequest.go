package models

// HTTPRequest is ..
type HTTPRequest struct {
	TimeStamp        string `json:"timestamp"`
	RequestID        string `json:"requestid"`
	OriginalClientIP string `json:"originalClientIP"`
	Method           string `json:"method"`
	URI              string `json:"uri"`
	HTTPVersion      string `json:"httpVersion"`
	Headers          string `json:"headers"`
	Body             string `json:"-"`
}
