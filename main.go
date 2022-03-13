// Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/sherlockjjobs/sigma/app/console"
	"github.com/sherlockjjobs/sigma/app/http"
	"github.com/sherlockjjobs/sigma/framework"
	"github.com/sherlockjjobs/sigma/framework/provider/app"
	"github.com/sherlockjjobs/sigma/framework/provider/cache"
	"github.com/sherlockjjobs/sigma/framework/provider/config"
	"github.com/sherlockjjobs/sigma/framework/provider/distributed"
	"github.com/sherlockjjobs/sigma/framework/provider/env"
	"github.com/sherlockjjobs/sigma/framework/provider/id"
	"github.com/sherlockjjobs/sigma/framework/provider/kernel"
	"github.com/sherlockjjobs/sigma/framework/provider/log"
	"github.com/sherlockjjobs/sigma/framework/provider/orm"
	"github.com/sherlockjjobs/sigma/framework/provider/redis"
	"github.com/sherlockjjobs/sigma/framework/provider/ssh"
	"github.com/sherlockjjobs/sigma/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&id.HadeIDProvider{})
	container.Bind(&trace.HadeTraceProvider{})
	container.Bind(&log.HadeLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.HadeCacheProvider{})
	container.Bind(&ssh.SSHProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
