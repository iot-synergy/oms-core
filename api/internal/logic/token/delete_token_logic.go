package token

import (
	"context"

	"github.com/iot-synergy/oms-core/api/internal/svc"
	"github.com/iot-synergy/oms-core/api/internal/types"
	"github.com/iot-synergy/oms-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTokenLogic) DeleteToken(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.DeleteToken(l.ctx, &core.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
