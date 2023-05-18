package comment

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCommentLogic {
	return &NewCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCommentLogic) NewComment(req *types.NewCommentReq) (resp *types.NewCommentResp, err error) {
	resp = new(types.NewCommentResp)
	userId := l.ctx.Value("userId").(string)

	// 获取回复用户id
	replyToId := ""
	if req.Scope == "comment" {
		replyTo, err := l.svcCtx.CommentRPC.RetrieveCommentById(l.ctx, &pb.RetrieveCommentByIdRequest{Id: req.Id})
		if err != nil {
			return nil, err
		}
		replyToId = replyTo.Comment.AuthorId
	}

	_, err = l.svcCtx.CommentRPC.CreateComment(l.ctx, &pb.CreateCommentRequest{
		Text:     req.Text,
		AuthorId: userId,
		ReplyTo:  replyToId,
		Type:     req.Scope,
		ParentId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}
