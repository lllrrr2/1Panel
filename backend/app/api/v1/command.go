package v1

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/constant"
	"github.com/gin-gonic/gin"
)

// @Tags Command
// @Summary Create command
// @Accept json
// @Param request body dto.CommandOperate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command [post]
// @x-panel-log {"bodyKeys":["name","command"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"创建快捷命令 [name][command]","formatEN":"create quick command [name][command]"}
func (b *BaseApi) CreateCommand(c *gin.Context) {
	var req dto.CommandOperate
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := commandService.Create(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Redis Command
// @Summary Save redis command
// @Accept json
// @Param request body dto.RedisCommand true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/redis [post]
// @x-panel-log {"bodyKeys":["name","command"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"保存 redis 快捷命令 [name][command]","formatEN":"save quick command for redis [name][command]"}
func (b *BaseApi) SaveRedisCommand(c *gin.Context) {
	var req dto.RedisCommand
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := commandService.SaveRedisCommand(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Command
// @Summary Page commands
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/search [post]
func (b *BaseApi) SearchCommand(c *gin.Context) {
	var req dto.SearchCommandWithPage
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	total, list, err := commandService.SearchWithPage(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})
}

// @Tags Redis Command
// @Summary Page redis commands
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/redis/search [post]
func (b *BaseApi) SearchRedisCommand(c *gin.Context) {
	var req dto.SearchWithPage
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	total, list, err := commandService.SearchRedisCommandWithPage(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})
}

// @Tags Command
// @Summary Tree commands
// @Accept json
// @Success 200 {array} dto.CommandTree
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/tree [get]
func (b *BaseApi) SearchCommandTree(c *gin.Context) {
	list, err := commandService.SearchForTree()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	helper.SuccessWithData(c, list)
}

// @Tags Redis Command
// @Summary List redis commands
// @Success 200 {array} dto.RedisCommand
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/redis [get]
func (b *BaseApi) ListRedisCommand(c *gin.Context) {
	list, err := commandService.ListRedisCommand()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	helper.SuccessWithData(c, list)
}

// @Tags Command
// @Summary List commands
// @Success 200 {object} dto.CommandInfo
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command [get]
func (b *BaseApi) ListCommand(c *gin.Context) {
	list, err := commandService.List()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}

	helper.SuccessWithData(c, list)
}

// @Tags Command
// @Summary Delete command
// @Accept json
// @Param request body dto.BatchDeleteReq true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/del [post]
// @x-panel-log {"bodyKeys":["ids"],"paramKeys":[],"BeforeFunctions":[{"input_column":"id","input_value":"ids","isList":true,"db":"commands","output_column":"name","output_value":"names"}],"formatZH":"删除快捷命令 [names]","formatEN":"delete quick command [names]"}
func (b *BaseApi) DeleteCommand(c *gin.Context) {
	var req dto.BatchDeleteReq
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := commandService.Delete(req.Ids); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Redis Command
// @Summary Delete redis command
// @Accept json
// @Param request body dto.BatchDeleteReq true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/redis/del [post]
// @x-panel-log {"bodyKeys":["ids"],"paramKeys":[],"BeforeFunctions":[{"input_column":"id","input_value":"ids","isList":true,"db":"redis_commands","output_column":"name","output_value":"names"}],"formatZH":"删除 redis 快捷命令 [names]","formatEN":"delete quick command of redis [names]"}
func (b *BaseApi) DeleteRedisCommand(c *gin.Context) {
	var req dto.BatchDeleteReq
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := commandService.DeleteRedisCommand(req.Ids); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Command
// @Summary Update command
// @Accept json
// @Param request body dto.CommandOperate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/command/update [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"更新快捷命令 [name]","formatEN":"update quick command [name]"}
func (b *BaseApi) UpdateCommand(c *gin.Context) {
	var req dto.CommandOperate
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	upMap := make(map[string]interface{})
	upMap["name"] = req.Name
	upMap["group_id"] = req.GroupID
	upMap["command"] = req.Command
	if err := commandService.Update(req.ID, upMap); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}
