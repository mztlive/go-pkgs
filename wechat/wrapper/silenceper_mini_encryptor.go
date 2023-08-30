package wrapper

import (
	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/wechat/mini"
)

func (m *SilenceperMini) Decrypt(sessionKey, encryptedData, iv string) (*mini.PlainData, error) {
	plainData, err := m.engine.GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	result := &mini.PlainData{}
	copier.Copy(result, &plainData)
	return result, nil
}
