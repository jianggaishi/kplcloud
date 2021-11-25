/**
 * @Time : 24/04/2020 11:49 AM
 * @Author : solacowa@gmail.com
 * @File : service_gen_table
 * @Software: GoLand
 */

package server

import (
	"context"
	"github.com/go-kit/kit/log/level"
	"github.com/spf13/cobra"

	"github.com/kplcloud/kplcloud/src/repository/types"
)

var (
	generateCmd = &cobra.Command{
		Use:               "generate command <args> [flags]",
		Short:             "生成命令",
		SilenceErrors:     false,
		DisableAutoGenTag: false,
		Example: `## 生成命令
可用的配置类型：
[table, init-data]

kplcloud generate -h
`,
	}

	genTableCmd = &cobra.Command{
		Use:               `table <args> [flags]`,
		Short:             "生成数据库表",
		SilenceErrors:     false,
		DisableAutoGenTag: false,
		Example: `
kplcloud generate table all
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 关闭资源连接
			defer func() {
				_ = level.Debug(logger).Log("db", "close", "err", db.Close())
				if rds != nil {
					_ = rds.Close(context.Background())
				}
			}()

			if len(args) > 0 && args[0] == "all" {
				return generateTable()
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return installPre()
		},
	}

	genInitDataCmd = &cobra.Command{
		Use:               `init-data <args> [flags]`,
		Short:             "生成数据",
		SilenceErrors:     false,
		DisableAutoGenTag: false,
		Example: `
kplcloud generate init-data
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 关闭资源连接
			defer func() {
				_ = level.Debug(logger).Log("db", "close", "err", db.Close())
				if rds != nil {
					_ = level.Debug(logger).Log("redis", "close", "err", rds.Close(context.Background()))
				}
			}()

			return generateInitData()
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return installPre()
		},
	}
)

func generateInitData() (err error) {
	// 初始化数据
	//authRsaPublicKey, authRsaPrivateKey, err := util.GenRsaKey()
	//if err != nil {
	//	_ = level.Error(logger).Log("util", "GenRsaKey", "err", err.Error())
	//	return
	//}
	//ctx := context.Background()
	//publicKey := strings.TrimSpace(string(authRsaPublicKey))
	//privateKey := strings.TrimSpace(string(authRsaPrivateKey))
	//publicKey = strings.Trim(publicKey, "\n")
	//privateKey = strings.Trim(privateKey, "\n")
	//_ = logger.Log("add", "data", "publicKey", store.SysSetting().Add(ctx, "AUTH_RSA_PUBLIC_KEY", publicKey, "公钥"))
	//_ = logger.Log("add", "data", "privateKey", store.SysSetting().Add(ctx, "AUTH_RSA_PRIVATE_KEY", privateKey, "私钥"))
	//
	//if cacheSvc != nil {
	//	_ = cacheSvc.Del(context.Background(), "auth:publicKey")
	//	_ = cacheSvc.Del(context.Background(), "auth:privateKey")
	//}

	return
}

func generateTable() (err error) {
	_ = logger.Log("migrate", "table", "SysRole", db.AutoMigrate(types.SysRole{}).Error)
	_ = logger.Log("migrate", "table", "SysUser", db.AutoMigrate(types.SysUser{}).Error)
	_ = logger.Log("migrate", "table", "SysPermission", db.AutoMigrate(types.SysPermission{}).Error)
	_ = logger.Log("migrate", "table", "SysNamespace", db.AutoMigrate(types.SysNamespace{}).Error)
	_ = logger.Log("migrate", "table", "SysSetting", db.AutoMigrate(types.SysSetting{}).Error)
	_ = logger.Log("migrate", "table", "Cluster", db.AutoMigrate(types.Cluster{}).Error)
	_ = logger.Log("migrate", "table", "Nodes", db.AutoMigrate(types.Nodes{}).Error)
	_ = logger.Log("migrate", "table", "Label", db.AutoMigrate(types.Label{}).Error)
	_ = logger.Log("migrate", "table", "Namespace", db.AutoMigrate(types.Namespace{}).Error)
	_ = logger.Log("migrate", "table", "ConfigMap", db.AutoMigrate(types.ConfigMap{}).Error)
	_ = logger.Log("migrate", "table", "Secrets", db.AutoMigrate(types.Secret{}).Error)
	_ = logger.Log("migrate", "table", "Data", db.AutoMigrate(types.Data{}).Error)
	_ = logger.Log("migrate", "table", "StorageClass", db.AutoMigrate(types.StorageClass{}).Error)
	_ = logger.Log("migrate", "table", "Registry", db.AutoMigrate(types.Registry{}).Error)
	_ = logger.Log("migrate", "table", "K8sTemplate", db.AutoMigrate(types.K8sTemplate{}).Error)
	_ = logger.Log("migrate", "table", "Audit", db.AutoMigrate(types.Audit{}).Error)
	_ = logger.Log("migrate", "table", "Application", db.AutoMigrate(types.Application{}).Error)
	_ = logger.Log("migrate", "table", "PersistentVolumeClaim", db.AutoMigrate(types.PersistentVolumeClaim{}).Error)
	return
}
