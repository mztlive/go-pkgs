package wrapper

import (
	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/wechat/mini"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
)

func (m *SilenceperMini) CreateWXAQRCode(coderParams mini.QRCoder) (response []byte, err error) {
	silenceperParam := qrcode.QRCoder{}
	if err := copier.Copy(&silenceperParam, &coderParams); err != nil {
		return nil, err
	}

	return m.engine.GetQRCode().CreateWXAQRCode(silenceperParam)
}

func (m *SilenceperMini) GetWXACode(coderParams mini.QRCoder) (response []byte, err error) {
	silenceperParam := qrcode.QRCoder{}
	if err := copier.Copy(&silenceperParam, &coderParams); err != nil {
		return nil, err
	}

	return m.engine.GetQRCode().GetWXACode(silenceperParam)
}

func (m *SilenceperMini) GetWXACodeUnlimit(coderParams mini.QRCoder) (response []byte, err error) {
	silenceperParam := qrcode.QRCoder{}
	if err := copier.Copy(&silenceperParam, &coderParams); err != nil {
		return nil, err
	}

	return m.engine.GetQRCode().GetWXACodeUnlimit(silenceperParam)
}
