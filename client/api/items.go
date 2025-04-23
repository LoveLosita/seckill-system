package api

import (
	initclient "client/init"
	"client/kitex-gens/items/kitex_gen/items"
	"client/response"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"time"
)

func GetItemInfo(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取参数
	strID := c.Query("id")
	if strID == "" {
		c.JSON(consts.StatusBadRequest, response.MissingParam)
		return
	}
	//2.调用服务端接口
	intID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	var getItemInfoReq items.GetItemInfoRequest
	getItemInfoReq.Id = intID
	resp, err := initclient.NewItemClient.GetItemInfo(ctx, &getItemInfoReq)
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

func GetItemList(ctx context.Context, c *app.RequestContext) {
	//1.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	var getItemListReq items.GetItemListRequest
	resp, err := initclient.NewItemClient.GetItemList(ctx, &getItemListReq)
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
	//2.返回结果
	c.JSON(consts.StatusOK, resp)
}

func AddItem(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取参数
	var addItemReq items.AddItemRequest
	err := c.BindJSON(&addItemReq)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := initclient.NewItemClient.AddItem(ctx, &addItemReq)
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
	c.JSON(consts.StatusOK, resp)
}

func UpdateItem(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取参数
	var updateItemReq items.UpdateItemRequest
	err := c.BindJSON(&updateItemReq)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := initclient.NewItemClient.UpdateItem(ctx, &updateItemReq)
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
	c.JSON(consts.StatusOK, resp)
}

func DeleteItem(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取参数
	strID := c.Query("id")
	if strID == "" {
		c.JSON(consts.StatusBadRequest, response.MissingParam)
		return
	}
	//2.调用服务端接口
	intID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	var deleteItemReq items.DeleteItemRequest
	deleteItemReq.Id = intID
	resp, err := initclient.NewItemClient.DeleteItem(ctx, &deleteItemReq)
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
	c.JSON(consts.StatusOK, resp)
}
