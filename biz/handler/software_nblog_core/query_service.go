// Code generated by hertz generator.

package software_nblog_core

import (
	"context"
	"nblog.org.cn/software_nblog_core/biz/err_code"
	"nblog.org.cn/software_nblog_core/biz/service/log"
	"nblog.org.cn/software_nblog_core/utils"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	core "nblog.org.cn/software_nblog_core/biz/model/software_nblog_core"
)

// QueryLog .
// @router /api/admin/log [GET]
func QueryLog(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.QueryLogReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, utils.ConvErr(err_code.ParaError))
		return
	}

	resp := log.QueryLog(ctx, &req)

	c.JSON(http.StatusOK, resp)
}
