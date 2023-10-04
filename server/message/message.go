package message

func (m *Message) Parse() ([]byte, error) {
	return []byte(m.Msg), nil
}

func (m *Message) GetType() string {
	return MESSAGE
}
