package dictionarydetail

import (
	"context"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/pointy"

	"github.com/iot-synergy/oms-core/rpc/internal/svc"
	"github.com/iot-synergy/oms-core/rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/oms-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictionaryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictionaryDetailLogic {
	return &UpdateDictionaryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDictionaryDetailLogic) UpdateDictionaryDetail(in *core.DictionaryDetailInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.DictionaryDetail.UpdateOneID(*in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilTitle(in.Title).
		SetNotNilKey(in.Key).
		SetNotNilValue(in.Value).
		SetNotNilSort(in.Sort).
		SetNotNilDictionaryID(in.DictionaryId).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
