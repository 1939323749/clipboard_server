package message

func (r *Response) Parse() ([]byte, error) {
	return []byte(r.Msg), nil
}

func (r *Response) GetType() string {
	return RESPONSE
}
