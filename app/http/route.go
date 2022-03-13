package http

import (
	"github.com/sherlockjjobs/sigma/app/http/module/demo"
	"github.com/sherlockjjobs/sigma/framework/contract"
	"github.com/sherlockjjobs/sigma/framework/gin"
	ginSwagger "github.com/sherlockjjobs/sigma/framework/middleware/gin-swagger"
	"github.com/sherlockjjobs/sigma/framework/middleware/gin-swagger/swaggerFiles"
	"github.com/sherlockjjobs/sigma/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 动态路由定义
	demo.Register(r)
}
