package file

import (
	"context"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
	"server/common/xerr"
	"server/common/xjwt"
	"server/common/xupload"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileTokenLogic {
	return &FileTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FileToken
// Author [SliverFlow]
// @desc  获取临时文件上传 token
func (l *FileTokenLogic) FileToken(req *types.FileTokenReq) (resp *types.FileTokenReply, err error) {
	c := l.svcCtx.Config.StsAliOss

	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		logx.Error("获取 uuid 失败")
		return nil, xerr.NewErrMsg("获取 uuid 失败")
	}

	rep, err := xupload.AliStsResponse(&c)
	if err != nil {
		return nil, xerr.NewErrMsg("临时上传凭证获取失败")
	}

	return &types.FileTokenReply{
		AccessKeyId:     *rep.Body.Credentials.AccessKeyId,
		AccessKeySecret: *rep.Body.Credentials.AccessKeySecret,
		StsToken:        *rep.Body.Credentials.SecurityToken,
		Region:          c.Region,
		Bucket:          c.Bucket,
		FilePath:        "/" + req.Path + "/" + uuid + "/",
	}, nil
}
