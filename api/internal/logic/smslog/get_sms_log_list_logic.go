package smslog

import (
	"context"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/oms-core/api/internal/svc"
	"github.com/iot-synergy/oms-core/api/internal/types"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSmsLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogListLogic {
	return &GetSmsLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSmsLogListLogic) GetSmsLogList(req *types.SmsLogListReq) (resp *types.SmsLogListResp, err error) {
	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	data, err := l.svcCtx.McmsRpc.GetSmsLogList(l.ctx,
		&mcms.SmsLogListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			PhoneNumber: req.PhoneNumber,
			Content:     req.Content,
			Provider:    req.Provider,
			SendStatus:  req.SendStatus,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.SmsLogListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.SmsLogInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				PhoneNumber: v.PhoneNumber,
				Content:     v.Content,
				SendStatus:  v.SendStatus,
				Provider:    v.Provider,
			})
	}
	return resp, nil
}
