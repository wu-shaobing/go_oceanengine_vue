package errcode

// 通用错误码 (00xxxx)
const (
	Success           = 0      // 成功
	ErrUnknown        = 1      // 未知错误
	ErrInvalidParam   = 100001 // 参数错误
	ErrInvalidParams  = 100001 // 参数错误（别名）
	ErrNotFound       = 100002 // 资源不存在
	ErrAlreadyExists  = 100003 // 资源已存在
	ErrPermissionDeny = 100004 // 权限不足
	ErrInternalServer = 100005 // 服务器内部错误
	ErrServiceUnavail = 100006 // 服务不可用
	ErrTimeout        = 100007 // 请求超时
	ErrTooManyRequest = 100008 // 请求过于频繁
)

// 认证授权错误码 (10xxxx)
const (
	ErrUnauthorized        = 100100 // 未登录
	ErrTokenInvalid        = 100101 // Token 无效
	ErrTokenExpired        = 100102 // Token 已过期
	ErrLoginFailed         = 100103 // 登录失败
	ErrCaptchaInvalid      = 100104 // 验证码错误
	ErrPasswordWrong       = 100105 // 密码错误
	ErrAccountLocked       = 100106 // 账号已锁定
	ErrAccountDisabled     = 100107 // 账号已禁用
	ErrRefreshTokenInvalid = 100108 // 刷新 Token 无效
)

// 用户管理错误码 (20xxxx)
const (
	ErrUserNotFound    = 200001 // 用户不存在
	ErrUserExists      = 200002 // 用户已存在
	ErrUsernameInvalid = 200003 // 用户名格式错误
	ErrPasswordInvalid = 200004 // 密码格式错误
	ErrEmailInvalid    = 200005 // 邮箱格式错误
	ErrPhoneInvalid    = 200006 // 手机号格式错误
)

// 角色管理错误码 (21xxxx)
const (
	ErrRoleNotFound = 210001 // 角色不存在
	ErrRoleExists   = 210002 // 角色已存在
	ErrRoleInUse    = 210003 // 角色正在使用中
)

// 菜单管理错误码 (22xxxx)
const (
	ErrMenuNotFound    = 220001 // 菜单不存在
	ErrMenuExists      = 220002 // 菜单已存在
	ErrMenuHasChildren = 220003 // 菜单存在子菜单
)

// 广告主管理错误码 (30xxxx)
const (
	ErrAdvertiserNotFound   = 300001 // 广告主不存在
	ErrAdvertiserExists     = 300002 // 广告主已存在
	ErrAdvertiserDisabled   = 300003 // 广告主已禁用
	ErrAdvertiserAuthFailed = 300004 // 广告主授权失败
	ErrAdvertiserSyncFailed = 300005 // 广告主同步失败
)

// 广告系列错误码 (31xxxx)
const (
	ErrCampaignNotFound   = 310001 // 广告系列不存在
	ErrCampaignExists     = 310002 // 广告系列已存在
	ErrCampaignCreateFail = 310003 // 创建广告系列失败
	ErrCampaignUpdateFail = 310004 // 更新广告系列失败
	ErrCampaignDeleteFail = 310005 // 删除广告系列失败
	ErrCampaignStatusFail = 310006 // 状态变更失败
)

// 广告组错误码 (32xxxx)
const (
	ErrAdNotFound   = 320001 // 广告组不存在
	ErrAdExists     = 320002 // 广告组已存在
	ErrAdCreateFail = 320003 // 创建广告组失败
	ErrAdUpdateFail = 320004 // 更新广告组失败
	ErrAdDeleteFail = 320005 // 删除广告组失败
)

// 创意错误码 (33xxxx)
const (
	ErrCreativeNotFound   = 330001 // 创意不存在
	ErrCreativeExists     = 330002 // 创意已存在
	ErrCreativeCreateFail = 330003 // 创建创意失败
)

// 素材错误码 (34xxxx)
const (
	ErrMaterialNotFound    = 340001 // 素材不存在
	ErrMaterialUploadFail  = 340002 // 素材上传失败
	ErrMaterialSizeLimit   = 340003 // 素材大小超限
	ErrMaterialTypeInvalid = 340004 // 素材类型不支持
)

// 报表错误码 (40xxxx)
const (
	ErrReportQueryFail  = 400001 // 报表查询失败
	ErrReportExportFail = 400002 // 报表导出失败
	ErrReportDateRange  = 400003 // 日期范围错误
)

// Ocean Engine API 错误码 (90xxxx)
const (
	ErrOEAPIFailed        = 900001 // API 调用失败
	ErrOETokenInvalid     = 900002 // Ocean Engine Token 无效
	ErrOETokenExpired     = 900003 // Ocean Engine Token 过期
	ErrOERateLimit        = 900004 // 请求频率限制
	ErrOEParamInvalid     = 900005 // 参数错误
	ErrOEResourceNotFound = 900006 // 资源不存在
)

// 错误消息映射
var messages = map[int]string{
	Success:           "成功",
	ErrUnknown:        "未知错误",
	ErrInvalidParam:   "参数错误",
	ErrNotFound:       "资源不存在",
	ErrAlreadyExists:  "资源已存在",
	ErrPermissionDeny: "权限不足",
	ErrInternalServer: "服务器内部错误",
	ErrServiceUnavail: "服务不可用",
	ErrTimeout:        "请求超时",
	ErrTooManyRequest: "请求过于频繁",

	ErrUnauthorized:        "请先登录",
	ErrTokenInvalid:        "Token 无效",
	ErrTokenExpired:        "Token 已过期",
	ErrLoginFailed:         "登录失败",
	ErrCaptchaInvalid:      "验证码错误",
	ErrPasswordWrong:       "密码错误",
	ErrAccountLocked:       "账号已锁定，请稍后再试",
	ErrAccountDisabled:     "账号已禁用",
	ErrRefreshTokenInvalid: "刷新 Token 无效",

	ErrUserNotFound:    "用户不存在",
	ErrUserExists:      "用户已存在",
	ErrUsernameInvalid: "用户名格式错误",
	ErrPasswordInvalid: "密码格式不符合要求",

	ErrRoleNotFound: "角色不存在",
	ErrRoleExists:   "角色已存在",
	ErrRoleInUse:    "角色正在使用中，无法删除",

	ErrMenuNotFound:    "菜单不存在",
	ErrMenuExists:      "菜单已存在",
	ErrMenuHasChildren: "菜单存在子菜单，无法删除",

	ErrAdvertiserNotFound:   "广告主不存在",
	ErrAdvertiserExists:     "广告主已存在",
	ErrAdvertiserDisabled:   "广告主已禁用",
	ErrAdvertiserAuthFailed: "广告主授权失败",
	ErrAdvertiserSyncFailed: "广告主同步失败",

	ErrCampaignNotFound:   "广告系列不存在",
	ErrCampaignExists:     "广告系列已存在",
	ErrCampaignCreateFail: "创建广告系列失败",
	ErrCampaignUpdateFail: "更新广告系列失败",
	ErrCampaignDeleteFail: "删除广告系列失败",
	ErrCampaignStatusFail: "状态变更失败",

	ErrAdNotFound:   "广告组不存在",
	ErrAdExists:     "广告组已存在",
	ErrAdCreateFail: "创建广告组失败",
	ErrAdUpdateFail: "更新广告组失败",
	ErrAdDeleteFail: "删除广告组失败",

	ErrCreativeNotFound:   "创意不存在",
	ErrCreativeExists:     "创意已存在",
	ErrCreativeCreateFail: "创建创意失败",

	ErrMaterialNotFound:    "素材不存在",
	ErrMaterialUploadFail:  "素材上传失败",
	ErrMaterialSizeLimit:   "素材大小超限",
	ErrMaterialTypeInvalid: "素材类型不支持",

	ErrReportQueryFail:  "报表查询失败",
	ErrReportExportFail: "报表导出失败",
	ErrReportDateRange:  "日期范围错误",

	ErrOEAPIFailed:        "Ocean Engine API 调用失败",
	ErrOETokenInvalid:     "广告主授权已失效，请重新授权",
	ErrOETokenExpired:     "广告主授权已过期，请重新授权",
	ErrOERateLimit:        "请求过于频繁，请稍后再试",
	ErrOEParamInvalid:     "Ocean Engine 参数错误",
	ErrOEResourceNotFound: "Ocean Engine 资源不存在",
}

// Message 获取错误消息
func Message(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "未知错误"
}
