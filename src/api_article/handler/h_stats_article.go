package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/pb"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"net/http"
)

var HandlerStatsArticleSet = wire.NewSet(wire.Struct(new(HandlerStatsArticle), "*"))

type HandlerStatsArticle struct {
	Logger      core.ILogger
	Res         core.IResponse
	RedisClient *redis.Client
	RocketMQ    core.IRocketMQ
}

func (h *HandlerStatsArticle) GetStatsArticleMain(c *gin.Context) {
	cfg := core.GetConfig()
	// 设置元信息
	ctx, err := core.SetMetadataForDownstreamFromHttpHeaders(context.Background(), c, cfg.Downstreams[0].Name, h.RedisClient)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "设置下游请求元信息失败", err)
		return
	}
	// 拨号连接
	addr := fmt.Sprintf("%s:%s", cfg.Downstreams[0].Endpoint.Host, cfg.Downstreams[0].Endpoint.Port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "连接拨号失败", err)
		return
	}
	// 发送请求
	client := pb.NewStatsServiceClient(conn)
	res, err := client.GetStatsArticleMain(ctx, &pb.Empty{})
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "桩函数请求调用失败", err)
		return
	}
	h.Res.ResSuccess(c, res)
}
