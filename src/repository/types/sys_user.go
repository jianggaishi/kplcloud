/**
 * @Time : 3/4/21 5:26 PM
 * @Author : solacowa@gmail.com
 * @File : sys_user
 * @Software: GoLand
 */

package types

import (
	"strings"
	"time"
)

type SysUser struct {
	Id                int64      `gorm:"column:id;primary_key" json:"id"`
	Username          string     `gorm:"column:username;null;comment:'姓名'" json:"username"`                              // 姓名
	Mobile            string     `gorm:"column:mobile;null;comment:'手机号'" json:"mobile"`                                 // 手机号
	LoginName         string     `gorm:"column:login_name;notnull;index;size:16;unique;comment:'登录名'" json:"login_name"` // 登录号
	Email             string     `gorm:"column:email;notnull;index;size:24;unique;comment:'邮箱'" json:"email"`            // 邮箱
	Password          string     `gorm:"column:password;null;size:128;comment:'密码'" json:"password"`                     // 密码
	Locked            bool       `gorm:"column:locked;null;default:false;comment:'是否锁定'" json:"locked"`                  // 是否锁定
	Expired           bool       `gorm:"column:expired;null;default:false;comment:'是否过期'" json:"expired"`                // 是否过期
	ConfirmationToken string     `gorm:"column:confirmation_token;null" json:"confirmation_token"`                       // 确认TOKEN
	WechatOpenId      string     `gorm:"column:wechat_openid;null;comment:'微信openId'" json:"wechat_openid"`              // 微信OPENID
	LastLogin         *time.Time `gorm:"column:last_login;null;comment:'最后登录时间'" json:"last_login"`                      // last_login
	Remark            string     `gorm:"column:remark;null;size:1000;comment:'备注'" json:"remark"`
	ExpiresAt         *time.Time `gorm:"column:expires_at;type:datetime" json:"expires_at"`     // 过期时间
	CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at" form:"updated_at"` // 更新时间
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                   // 删除时间

	SysRoles   []SysRole   `gorm:"many2many:sys_user_roles;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:role_id;jointable_foreignkey:sys_user_id;" json:"sys_roles"`
	Namespaces []Namespace `gorm:"many2many:sys_user_namespaces;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:namespace_id;jointable_foreignkey:sys_user_id;" json:"sys_namespaces"`
	Clusters   []Cluster   `gorm:"many2many:sys_user_cluster;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:cluster_id;jointable_foreignkey:user_id;" json:"clusters"`
	SysGroups  []SysRole   `gorm:"many2many:sys_group_users;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:group_id;jointable_foreignkey:sys_user_id;" json:"sys_groups"`
}

// TableName set table
func (*SysUser) TableName() string {
	return "sys_user"
}

func (s *SysUser) GroupIds() []int64 {
	var ids []int64
	for _, v := range s.SysGroups {
		ids = append(ids, v.Id)
	}
	return ids
}

func (s *SysUser) IsAdmin() bool {
	var isAdmin bool
	for _, v := range s.SysRoles {
		if strings.EqualFold(v.Name, "super.admin") {
			return true
		}
	}
	return isAdmin
}
