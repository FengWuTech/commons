package consul

import "github.com/hashicorp/consul/api"

var clt *api.Client

func Setup(rc string, tk string) {
	if rc == "" {
		panic("注册中心不能为空")
	}
	config := api.DefaultConfig()
	config.Address = rc
	config.Token = tk
	client, err := api.NewClient(config)
	if err != nil {
		panic("consul初始化失败")
	}

	clt = client
}

func Client() *api.Client {
	return clt
}
