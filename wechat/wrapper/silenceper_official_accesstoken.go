package wrapper

func (m *SilenceperOfficialAccount) GetAccessToken() (string, error) {
	return m.engine.GetAccessToken()
}
