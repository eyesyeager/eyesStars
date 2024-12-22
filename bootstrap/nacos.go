package bootstrap

import (
	"encoding/json"
	"eyesStars/global"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

/**
 * @author eyesYeager
 * @date 2023/2/8 13:12
 */

func InitNacos() {
	// 服务端配置
	sc := []constant.ServerConfig{
		{
			IpAddr: global.Native.Nacos.Ip,
			Port:   global.Native.Nacos.Port,
		},
	}

	// 客户端配置
	cc := constant.ClientConfig{
		Username:            global.Native.Nacos.Username,
		Password:            global.Native.Nacos.Password,
		NamespaceId:         global.Native.Nacos.NamespaceId,
		TimeoutMs:           global.Native.Nacos.Timeout,
		NotLoadCacheAtStart: global.Native.Nacos.NotLoadCacheAtStart,
		LogDir:              global.Native.Nacos.LogDir,
		CacheDir:            global.Native.Nacos.CacheDir,
		LogLevel:            global.Native.Nacos.LogLevel,
	}

	// 创建动态配置客户端
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		panic(err)
	}

	// 拉取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: global.Native.Nacos.DataId,
		Group:  global.Native.Nacos.Group,
	})
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(content), &global.Config)
	if err != nil {
		panic(err)
	}
}
