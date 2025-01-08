package handler

import (
	"github.com/coscms/webcore/library/navigate"
	"github.com/webx-top/echo"
)

var LeftNavigate = &navigate.Item{
	Display: true,
	Name:    echo.T(`FTP账号`),
	Action:  `ftp`,
	Icon:    `users`,
	Children: &navigate.List{
		{
			Display: true,
			Name:    echo.T(`账号管理`),
			Action:  `account`,
		},
		{
			Display: true,
			Name:    echo.T(`添加账号`),
			Action:  `account_add`,
			Icon:    `plus`,
		},
		{
			Display: true,
			Name:    echo.T(`用户组`),
			Action:  `group`,
		},
		{
			Display: true,
			Name:    echo.T(`添加用户组`),
			Action:  `group_add`,
			Icon:    `plus`,
		},
		{
			Display: false,
			Name:    echo.T(`修改账号`),
			Action:  `account_edit`,
			Icon:    ``,
		},
		{
			Display: false,
			Name:    echo.T(`删除账号`),
			Action:  `account_delete`,
			Icon:    ``,
		},
		{
			Display: false,
			Name:    echo.T(`修改分组`),
			Action:  `group_edit`,
			Icon:    ``,
		},
		{
			Display: false,
			Name:    echo.T(`删除分组`),
			Action:  `group_delete`,
			Icon:    ``,
		},
		{
			Display: false,
			Name:    echo.T(`重启FTP服务`),
			Action:  `restart`,
			Icon:    ``,
		},
		{
			Display: false,
			Name:    echo.T(`关闭FTP服务`),
			Action:  `stop`,
			Icon:    ``,
		},
		{
			Display: false,
			Name:    echo.T(`查看FTP动态`),
			Action:  `log`,
			Icon:    ``,
		},
	},
}
