package menu

import (
	"context"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/iot-synergy/oms-core/rpc/ent/menu"

	"github.com/iot-synergy/oms-core/rpc/internal/svc"
	"github.com/iot-synergy/oms-core/rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/oms-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *core.IDReq) (*core.BaseResp, error) {
	exist, err := l.svcCtx.DB.Menu.Query().Where(menu.ParentID(in.Id)).Exist(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if exist {
		logx.Errorw("delete menu failed, please check its children had been deleted",
			logx.Field("menuId", in.Id))
		return nil, errorx.NewInvalidArgumentError("menu.deleteChildrenDesc")
	}

	err = l.svcCtx.DB.Menu.DeleteOneID(in.Id).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
