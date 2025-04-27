package api

import (
	init_client "client/init"
	"client/kitex-gens/seckill/kitex_gen/seckill"
	"client/response"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

func CreateSecKillEventHandler(ctx context.Context, c *app.RequestContext) {
	var event seckill.CreateSecKillRequest
	//1.从请求中获取参数
	token := c.GetHeader("Authorization")
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	event.Token = string(token)
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewSecKillClient.CreateSecKill(ctx, &event)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(consts.StatusInternalServerError, response.RpcServerConnectTimeOut)
			return
		} else {
			c.JSON(consts.StatusInternalServerError, response.InternalError(err))
			return
		}
	}
	if resp != nil {
		if resp.Status.Code == "500" { //如果是内部错误
			c.JSON(consts.StatusInternalServerError, resp.Status)
			return
		} else if resp.Status.Code != "10000" { //如果是参数错误
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
	}
	//3.返回结果
	c.JSON(consts.StatusOK, resp)
}

func SecKillHandler(ctx context.Context, c *app.RequestContext) {
	var event seckill.SecKillRequest
	//1.从请求中获取参数
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewSecKillClient.SecKill(ctx, &event)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(consts.StatusInternalServerError, response.RpcServerConnectTimeOut)
			return
		} else {
			c.JSON(consts.StatusInternalServerError, response.InternalError(err))
			return
		}
	}
	if resp != nil {
		if resp.Status.Code == "500" { //如果是内部错误
			c.JSON(consts.StatusInternalServerError, resp.Status)
			return
		} else if resp.Status.Code != "10000" { //如果是参数错误
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
	}
	//3.返回结果
	c.JSON(consts.StatusOK, resp)
}

func GetOrderStatusHandler(ctx context.Context, c *app.RequestContext) {
	var req seckill.GetOrderStatusRequest
	//1.从请求中获取参数
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewSecKillClient.GetOrderStatus(ctx, &req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(consts.StatusInternalServerError, response.RpcServerConnectTimeOut)
			return
		} else {
			c.JSON(consts.StatusInternalServerError, response.InternalError(err))
			return
		}
	}
	if resp != nil {
		if resp.Status.Code == "500" { //如果是内部错误
			c.JSON(consts.StatusInternalServerError, resp.Status)
			return
		} else if resp.Status.Code != "10000" { //如果是参数错误
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
	}
	//3.返回结果
	c.JSON(consts.StatusOK, resp)
}
