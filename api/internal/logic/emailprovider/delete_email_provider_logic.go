package emailprovider

import (
	"context"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/oms-core/api/internal/svc"
	"github.com/iot-synergy/oms-core/api/internal/types"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailProviderLogic {
	return &DeleteEmailProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmailProviderLogic) DeleteEmailProvider(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.DeleteEmailProvider(l.ctx, &mcms.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
