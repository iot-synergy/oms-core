package department

import (
	"context"

	"github.com/iot-synergy/oms-core/rpc/ent/user"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/oms-core/rpc/ent/department"

	"github.com/iot-synergy/oms-core/rpc/internal/svc"
	"github.com/iot-synergy/oms-core/rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/oms-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/iot-synergy/synergy-common/i18n"
)

type DeleteDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDepartmentLogic {
	return &DeleteDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDepartmentLogic) DeleteDepartment(in *core.IDsReq) (*core.BaseResp, error) {
	exist, err := l.svcCtx.DB.Department.Query().Where(department.ParentIDIn(in.Ids...)).Exist(l.ctx)
	if exist {
		logx.Errorw("delete department failed, please check its children had been deleted",
			logx.Field("departmentId", in.Ids))
		return nil, errorx.NewInvalidArgumentError("department.deleteDepartmentChildrenFirst")
	}

	checkUser, err := l.svcCtx.DB.User.Query().Where(user.DepartmentIDIn(in.Ids...)).Exist(l.ctx)
	if checkUser {
		logx.Errorw("delete department failed, there are users belongs to the department", logx.Field("departmentId", in.Ids))
		return nil, errorx.NewInvalidArgumentError("department.deleteDepartmentUserFirst")
	}

	_, err = l.svcCtx.DB.Department.Delete().Where(department.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
