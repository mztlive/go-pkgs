package wrapper

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/wechat/mini"
)

func (m *SilenceperMini) Code2Session(ctx context.Context, jsCode string) (result *mini.ResCode2Session, err error) {

	code2Session, err := m.engine.GetAuth().Code2Session(jsCode)
	if err != nil {
		return nil, err
	}

	result = &mini.ResCode2Session{}
	copier.Copy(result, &code2Session)

	return result, nil
}
