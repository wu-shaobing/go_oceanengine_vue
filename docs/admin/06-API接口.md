# Admin API 接口参考

本文汇总 Admin 后台相关 API（认证 + 系统管理）。统一响应结构见“响应格式”。

## 认证 (/api/v1/auth)

- POST /api/v1/auth/login
  - 入参: JSON dto.LoginReq { username, password, captcha_id?, captcha_code? }
  - 出参: dto.LoginResp { access_token, refresh_token, expires_in, user }
- POST /api/v1/auth/refresh
  - 入参: JSON dto.RefreshTokenReq { refresh_token }
  - 出参: dto.LoginResp
- GET /api/v1/auth/userinfo 需 Bearer
  - 出参: dto.UserInfo
- POST /api/v1/auth/logout 需 Bearer
  - 出参: 空

## 系统用户 (/api/v1/system/users)

- GET    /api/v1/system/users
  - Query: username, nickname, phone, status, role_id, page, page_size
  - 出参: PageData<dto.UserListResp>
- GET    /api/v1/system/users/{id}
  - 出参: model.User
- POST   /api/v1/system/users
  - 入参: JSON dto.UserCreateReq
- PUT    /api/v1/system/users/{id}
  - 入参: JSON dto.UserUpdateReq
- DELETE /api/v1/system/users/{id}
- POST   /api/v1/system/users/{id}/reset-password
  - 入参: JSON dto.UserResetPasswordReq
- POST   /api/v1/system/users/change-password
  - 入参: JSON dto.UserChangePasswordReq

## 系统角色 (/api/v1/system/roles)

- GET    /api/v1/system/roles
  - Query: name, code, status, page, page_size
  - 出参: PageData<model.Role>
- GET    /api/v1/system/roles/all
  - 出参: []model.Role
- GET    /api/v1/system/roles/{id}
  - 出参: model.Role
- POST   /api/v1/system/roles
  - 入参: JSON dto.RoleCreateReq
- PUT    /api/v1/system/roles/{id}
  - 入参: JSON dto.RoleUpdateReq
- DELETE /api/v1/system/roles/{id}
- GET    /api/v1/system/roles/{id}/menus
  - 出参: []uint64 (菜单ID)
- PUT    /api/v1/system/roles/{id}/menus
  - 入参: JSON dto.RoleMenuUpdateReq { menu_ids }

## 系统菜单 (/api/v1/system/menus)

- GET    /api/v1/system/menus
  - 出参: []model.Menu
- GET    /api/v1/system/menus/tree
  - 出参: []dto.MenuTree
- GET    /api/v1/system/menus/user
  - 出参: []dto.MenuTree (按当前用户角色过滤)
- GET    /api/v1/system/menus/{id}
  - 出参: model.Menu
- POST   /api/v1/system/menus
  - 入参: JSON dto.MenuCreateReq
- PUT    /api/v1/system/menus/{id}
  - 入参: JSON dto.MenuUpdateReq
- DELETE /api/v1/system/menus/{id}

## 操作日志 (/api/v1/system/logs)

- GET    /api/v1/system/logs/operation
  - Query: user_id, username, module, action, status, start_time, end_time, page, page_size
  - 出参: PageData<dto.OperationLogResp>
- GET    /api/v1/system/logs/modules
  - 出参: []string
- DELETE /api/v1/system/logs/operation
  - 入参: JSON dto.OperationLogDeleteReq { before_time }

## 响应格式

所有接口返回统一结构 `response.Response`：

```json
{
  "code": 0,
  "message": "成功",
  "data": {},
  "request_id": "...",
  "timestamp": 1733731200000
}
```

分页返回 `response.PageData` 包装在 `data` 内：

```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "list": [ /* ... */ ],
    "total": 123,
    "page": 1,
    "page_size": 20
  },
  "request_id": "...",
  "timestamp": 1733731200000
}
```

## 成功与错误码

- 成功：`code = 0`
- 常见错误：参数错误 `ErrInvalidParam`、未授权 `ErrUnauthorized`、权限不足 `ErrPermissionDeny`、资源不存在 `ErrNotFound`、服务器错误 `ErrInternalServer`（详见 `backend/pkg/errcode`）。

## 备注

- 所有 `/api/v1/system/*` 端点均需 Bearer Token（登录后获取）。
- 时间格式推荐 `2006-01-02 15:04:05`。