package emaillog

import (
	"context"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/oms-core/api/internal/svc"
	"github.com/iot-synergy/oms-core/api/internal/types"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmailLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmailLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmailLogLogic {
	return &CreateEmailLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmailLogLogic) CreateEmailLog(req *types.EmailLogInfo) (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.CreateEmailLog(l.ctx,
		&mcms.EmailLogInfo{
			Target:     req.Target,
			Subject:    req.Subject,
			Content:    req.Content,
			SendStatus: req.SendStatus,
			Provider:   req.Provider,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
