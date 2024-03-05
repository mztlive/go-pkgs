package wrapper

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/wechat/mini/dataanalysis"
	"github.com/silenceper/wechat/v2/miniprogram/analysis"
)

func (m *SilenceperMini) GetAnalysisDailyRetain(ctx context.Context, beginDate, endDate string) (*dataanalysis.ResAnalysisRetain, error) {
	var (
		out *dataanalysis.ResAnalysisRetain
		err error

		result analysis.ResAnalysisRetain
	)

	if result, err = m.engine.GetAnalysis().GetAnalysisDailyRetain(beginDate, endDate); err != nil {
		return nil, fmt.Errorf("GetAnalysisDailyRetain Error, err=%w", err)
	}

	if err = copier.Copy(out, &result); err != nil {
		return nil, fmt.Errorf("copier.Copy Error, err=%w", err)
	}

	return out, nil
}
