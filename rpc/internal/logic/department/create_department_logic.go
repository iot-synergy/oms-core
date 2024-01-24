package department

import (
	"context"

	"github.com/iot-synergy/synergy-common/utils/pointy"

	"github.com/iot-synergy/oms-core/rpc/internal/svc"
	"github.com/iot-synergy/oms-core/rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/oms-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/iot-synergy/synergy-common/i18n"
)

type CreateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(in *core.DepartmentInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Department.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilSort(in.Sort).
		SetNotNilName(in.Name).
		SetNotNilAncestors(in.Ancestors).
		SetNotNilLeader(in.Leader).
		SetNotNilPhone(in.Phone).
		SetNotNilEmail(in.Email).
		SetNotNilRemark(in.Remark).
		SetNotNilParentID(in.ParentId).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
