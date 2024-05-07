package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ikun666/kun_chat/chat/api/internal/global"
	"github.com/ikun666/kun_chat/chat/api/internal/svc"
	"github.com/ikun666/kun_chat/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMsgsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMsgsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMsgsLogic {
	return &GetMsgsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMsgsLogic) GetMsgs(req *types.GetMsgRequest) (resp *types.GetMsgResponse, err error) {
	// todo: add your logic here and delete this line
	// RedisMsg 获取缓存里面的消息

	userIdStr := strconv.Itoa(int(req.FormId))
	targetIdStr := strconv.Itoa(int(req.TargetId))

	//拼接记录名称
	var key string
	if req.FormId < 1000000000 {
		key = "msg_g_" + userIdStr
	} else {
		if req.FormId < req.TargetId {
			key = "msg_f_" + userIdStr + "_" + targetIdStr
		} else {
			key = "msg_f_" + targetIdStr + "_" + userIdStr
		}

	}

	var msgs []string
	if req.Reverse {
		msgs, err = global.Redis.Zrange(key, req.Start, req.Stop)
	} else {
		msgs, err = global.Redis.Zrevrange(key, req.Start, req.Stop)
	}
	if err != nil {
		fmt.Println(err) //没有找到
		return nil, err
	}
	// fmt.Println("聊天记录：", msgs)
	return &types.GetMsgResponse{Msgs: msgs}, nil

}
