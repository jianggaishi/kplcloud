/**
 * @Time: 2020/3/27 17:34
 * @Author: solacowa@gmail.com
 * @File: responsestatus
 * @Software: GoLand
 */

package encode

import (
	"github.com/pkg/errors"
)

type ResStatus string

var ResponseMessage = map[ResStatus]int{
	Invalid:        400,
	InvalidParams:  400,
	ErrParamsPhone: 401,
	ErrBadRoute:    401,
	ErrSystem:      500,
	ErrNotfound:    404,
	ErrLimiter:     429,

	ErrAccountNotFound:         404,
	ErrAccountLogin:            1002,
	ErrAccountLoginIsNull:      1003,
	ErrAccountNotLogin:         501,
	ErrAccountASD:              1004,
	ErrAccountLocked:           1005,
	ErrAuthNotLogin:            501,
	ErrAuthLogin:               1006,
	ErrAuthCheckCaptchaCode:    1007,
	ErrAuthCheckCaptchaNotnull: 1008,
	ErrAuthRegisterExists:      1009,
	ErrAuthRegisterSave:        1010,
	ErrAuthTimeout:             1011,

	// 系统API
	ErrSysRoleNotfound:     2001,
	ErrSysRoleSave:         2002,
	ErrSysRoleUserNotfound: 2003,
	ErrSysRoleUser:         2004,
	ErrSysRoleUserLen:      2005,
	ErrSysRoleUserDelete:   2006,
	ErrSysUserNotfound:     2007,

	ErrSysSettingNotfound: 2020,
	ErrSysSettingExists:   2021,
	ErrSysSettingDelete:   2022,
	ErrSysSettingSave:     2023,

	ErrInstallDbConnect:  901,
	ErrInstallDbDrive:    902,
	ErrInstallUploadPath: 903,
	ErrInstallWebPath:    904,
	ErrInstallUpload:     905,

	ErrClusterAdd:           920,
	ErrClusterConnect:       921,
	ErrClusterNotfound:      922,
	ErrClusterParams:        923,
	ErrClusterList:          924,
	ErrClusterNotPermission: 925,

	ErrNamespaceNotfound: 940,
	ErrNameNotfound:      941,
	ErrNamespaceExists:   942,
	ErrNamespaceSave:     943,

	ErrDeploymentSyncList:    960,
	ErrDeploymentGetNotfound: 961,

	ErrConfigMapSyncList: 980,

	ErrStorageClassSync:     1100,
	ErrStorageClassNotfound: 1101,
	ErrStorageClassSyncPv:   1102,
	ErrStorageClassExists:   1103,
	ErrStorageClassCreate:   1104,
	ErrStorageClassUpdate:   1105,

	ErrSecretMarshal:   1120,
	ErrSecretImageSave: 1121,
	ErrSecretDelete:    1122,

	ErrNodeCordon:   1140,
	ErrNodeNotfound: 1141,
	ErrNodeDrain:    1142,

	ErrPersistentVolumeClaimList:     1160,
	ErrPersistentVolumeClaimCreate:   1161,
	ErrPersistentVolumeClaimNotfound: 1162,
	ErrPersistentVolumeClaimExists:   1163,
	ErrPersistentVolumeClaimDelete:   1164,

	ErrTemplateSave: 1180,

	ErrRegistryDelete:   1200,
	ErrRegistryUpdate:   1201,
	ErrRegistryNotfound: 1202,
}

const (
	// 公共错误信息
	Invalid        ResStatus = "invalid"
	InvalidParams  ResStatus = "请求参数错误"
	ErrNotfound    ResStatus = "不存在"
	ErrBadRoute    ResStatus = "请求路由错误"
	ErrParamsPhone ResStatus = "手机格式不正确"
	ErrLimiter     ResStatus = "太快了,等我一会儿..."

	ErrInstallDbConnect  ResStatus = "数据库连接失败"
	ErrInstallDbDrive    ResStatus = "暂不支持其他数据库"
	ErrInstallUploadPath ResStatus = "文件目录未配置"
	ErrInstallWebPath    ResStatus = "Web目录未配置"
	ErrInstallUpload     ResStatus = "上传文件失败"

	ErrClusterAdd           ResStatus = "集群添加错误"
	ErrClusterConnect       ResStatus = "集群链接错误"
	ErrClusterNotfound      ResStatus = "集群不存在"
	ErrClusterParams        ResStatus = "集群参数错误"
	ErrClusterList          ResStatus = "集群列表错误"
	ErrClusterNotPermission ResStatus = "无权访问该集群"
	ErrClusterDelete        ResStatus = "集群删除失败"

	ErrSecretMarshal   ResStatus = "转换错误"
	ErrSecretImageSave ResStatus = "添加镜像Secret错误"
	ErrSecretDelete    ResStatus = "Secret删除错误"
	ErrSecretExists    ResStatus = "Secret已存在"

	ErrNamespaceNotfound ResStatus = "空间不存在"
	ErrNameNotfound      ResStatus = "名称不存在"
	ErrNamespaceExists   ResStatus = "空间已存在"
	ErrNamespaceSave     ResStatus = "空间保存失败"
	ErrNamespaceList     ResStatus = "空间获取出错"

	ErrRegistryNotfound ResStatus = "镜像仓库不存在"
	ErrRegistryUpdate   ResStatus = "镜像仓库更新失败"
	ErrRegistryDelete   ResStatus = "镜像仓库删除失败"

	ErrDeploymentSyncList    ResStatus = "同步失败"
	ErrDeploymentGetNotfound ResStatus = "项目不存在"

	ErrConfigMapSyncList ResStatus = "ConfigMap同步失败"

	ErrStorageClassSync     ResStatus = "StorageClass同步失败"
	ErrStorageClassNotfound ResStatus = "StorageClass不存在"
	ErrStorageClassSyncPv   ResStatus = "StorageClass Pv同步失败"
	ErrStorageClassExists   ResStatus = "StorageClass 已存在"
	ErrStorageClassCreate   ResStatus = "StorageClass 创建失败"
	ErrStorageClassDelete   ResStatus = "StorageClass 删除失败"
	ErrStorageClassUpdate   ResStatus = "StorageClass 更新失败"

	ErrPersistentVolumeClaimList     ResStatus = "存储卷声明列表获取失败"
	ErrPersistentVolumeClaimCreate   ResStatus = "存储卷声明创建失败"
	ErrPersistentVolumeClaimNotfound ResStatus = "存储卷声明不存在"
	ErrPersistentVolumeClaimExists   ResStatus = "存储卷声明已存在"
	ErrPersistentVolumeClaimDelete   ResStatus = "存储卷声明删除失败"

	ErrPodNotfound ResStatus = "Pods可能不存在"

	ErrTerminalPodsNotfound ResStatus = "没有找到相应标签的pods"
	ErrTerminalToken        ResStatus = "生成token失败"

	ErrAutoScaleMetricsServerNotfound ResStatus = "metrics-server不存在无法创建HPA"

	// 中间件错误信息
	ErrSystem                  ResStatus = "系统错误"
	ErrAccountNotLogin         ResStatus = "用户没登录"
	ErrAuthNotLogin            ResStatus = "请先登录"
	ErrAccountLoginIsNull      ResStatus = "用户名和密码不能为空"
	ErrAccountLogin            ResStatus = "用户名或密码错误"
	ErrAccountNotFound         ResStatus = "账号不存在"
	ErrAccountASD              ResStatus = "权限验证失败"
	ErrAccountLocked           ResStatus = "用户已被锁定"
	ErrAccountNamespace        ResStatus = "用户空间错误"
	ErrAuthLogin               ResStatus = "登录失败"
	ErrAuthCheckCaptchaCode    ResStatus = "图形验证码错误"
	ErrAuthCheckCaptchaNotnull ResStatus = "图形验证码不能为空"
	ErrAuthRegisterExists      ResStatus = "注册用户已存在"
	ErrAuthRegisterSave        ResStatus = "注册失败,请联系管理员"
	ErrAuthTimeout             ResStatus = "授权已过期"

	// 系统API
	ErrSysRoleNotfound     ResStatus = "角色不存在"
	ErrSysRoleSave         ResStatus = "角色保证错误"
	ErrSysRoleUserNotfound ResStatus = "角色用户不存在"
	ErrSysRoleUser         ResStatus = "用户角色配置失败"
	ErrSysRoleUserLen      ResStatus = "请选择用户"
	ErrSysRoleUserDelete   ResStatus = "角色删除失败"
	ErrSysUserNotfound     ResStatus = "用户不存在"

	ErrSysSettingNotfound ResStatus = "配置不存在"
	ErrSysSettingDelete   ResStatus = "删除失败"
	ErrSysSettingExists   ResStatus = "配置已存在"
	ErrSysSettingSave     ResStatus = "配置保存失败"

	ErrSysGroupExists        ResStatus = "组已存在"
	ErrSysGroupNotfound      ResStatus = "组不存在"
	ErrSysGroupNotPermission ResStatus = "您无法操作,请让组管理员操作"
	ErrSysGroupSave          ResStatus = "保存失败"
	ErrSysGroupDelete        ResStatus = "删除失败"

	ErrNodeCordon    ResStatus = "操作失败"
	ErrNodeNotfound  ResStatus = "节点不存在"
	ErrNodeDrain     ResStatus = "驱逐失败"
	ErrNodeDelete    ResStatus = "节点删除失败"
	ErrNodeScheduled ResStatus = "请将节点设置为不可调度状态"

	ErrTemplateSave ResStatus = "模版保存失败"
)

func (c ResStatus) String() string {
	return string(c)
}

func (c ResStatus) Error() error {
	return errors.New(string(c))
}

func (c ResStatus) Wrap(err error) error {
	return errors.Wrap(err, string(c))
}
