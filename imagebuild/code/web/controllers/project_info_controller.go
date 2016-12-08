/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package controllers

import (
	log "github.com/Sirupsen/logrus"
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
项目信息
 */
type ProjectInfoController struct {
	BasicController
}

func (c *ProjectInfoController) Get() {
	log.Info("ProjectInfoController: %s", c.Ctx.Request.Form)

	project := c.GetString("projectName")
	creator := c.Operator()

	if creator == "" || project == "" {
		log.Error("creator,projectName should not be empy when building project")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			-1,
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	c.Layout = "project.tpl"
	c.TplName = "info.tpl"
	code, info := models.AppServer.GetProjectInfo(project)
	resp := models.BuildResponse(
		code,
		info,
		errors.ErrorCodeToMessage(code))

	c.Data["json"] = resp
	c.ServeJSON(true)
	return
}
