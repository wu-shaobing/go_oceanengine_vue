package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/errcode"
)

// RoleService 角色服务
type RoleService struct {
	db *gorm.DB
}

// NewRoleService 创建角色服务
func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{db: db}
}

// GetList 获取角色列表
func (s *RoleService) GetList(ctx context.Context, req *dto.RoleListReq) ([]*model.Role, int64, error) {
	var roles []*model.Role
	var total int64

	query := s.db.WithContext(ctx).Model(&model.Role{})

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		query = query.Where("code LIKE ?", "%"+req.Code+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("sort ASC, id ASC").Find(&roles).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return roles, total, nil
}

// GetAll 获取所有角色（不分页）
func (s *RoleService) GetAll(ctx context.Context) ([]*model.Role, error) {
	var roles []*model.Role
	if err := s.db.WithContext(ctx).Where("status = ?", model.UserStatusEnabled).Order("sort ASC, id ASC").Find(&roles).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return roles, nil
}

// GetByID 获取角色详情
func (s *RoleService) GetByID(ctx context.Context, id uint64) (*model.Role, error) {
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrRoleNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &role, nil
}

// Create 创建角色
func (s *RoleService) Create(ctx context.Context, req *dto.RoleCreateReq, operatorID uint64) error {
	// 检查角色 Key 是否已存在
	var count int64
	if err := s.db.WithContext(ctx).Model(&model.Role{}).Where("`key` = ?", req.Code).Count(&count).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	if count > 0 {
		return errcode.New(errcode.ErrRoleExists)
	}

	role := &model.Role{
		Name:      req.Name,
		Key:       req.Code,
		Status:    req.Status,
		Sort:      req.Sort,
		Remark:    req.Remark,
		CreatedBy: operatorID,
		UpdatedBy: operatorID,
	}

	if err := s.db.WithContext(ctx).Create(role).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Update 更新角色
func (s *RoleService) Update(ctx context.Context, req *dto.RoleUpdateReq, operatorID uint64) error {
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrRoleNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 如果修改了 Key，检查是否重复
	if req.Code != "" && req.Code != role.Key {
		var count int64
		if err := s.db.WithContext(ctx).Model(&model.Role{}).Where("`key` = ? AND id != ?", req.Code, req.ID).Count(&count).Error; err != nil {
			return errcode.Wrap(errcode.ErrInternalServer, err)
		}
		if count > 0 {
			return errcode.New(errcode.ErrRoleExists)
		}
	}

	updates := map[string]interface{}{
		"updated_by": operatorID,
	}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Code != "" {
		updates["key"] = req.Code
	}
	updates["status"] = req.Status
	updates["sort"] = req.Sort
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := s.db.WithContext(ctx).Model(&role).Updates(updates).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Delete 删除角色
func (s *RoleService) Delete(ctx context.Context, id uint64) error {
	// 检查是否有用户使用该角色
	var userCount int64
	if err := s.db.WithContext(ctx).Model(&model.User{}).Where("role_id = ?", id).Count(&userCount).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	if userCount > 0 {
		return errcode.New(errcode.ErrRoleInUse)
	}

	// 删除角色菜单关联
	if err := s.db.WithContext(ctx).Where("role_id = ?", id).Delete(&model.RoleMenu{}).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 删除角色
	result := s.db.WithContext(ctx).Delete(&model.Role{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrRoleNotFound)
	}
	return nil
}

// GetRoleMenus 获取角色菜单ID列表
func (s *RoleService) GetRoleMenus(ctx context.Context, roleID uint64) ([]uint64, error) {
	var roleMenus []model.RoleMenu
	if err := s.db.WithContext(ctx).Where("role_id = ?", roleID).Find(&roleMenus).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	menuIDs := make([]uint64, len(roleMenus))
	for i, rm := range roleMenus {
		menuIDs[i] = rm.MenuID
	}
	return menuIDs, nil
}

// UpdateRoleMenus 更新角色菜单
func (s *RoleService) UpdateRoleMenus(ctx context.Context, roleID uint64, menuIDs []uint64) error {
	// 检查角色是否存在
	var role model.Role
	if err := s.db.WithContext(ctx).First(&role, roleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrRoleNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 使用事务更新
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除原有关联
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RoleMenu{}).Error; err != nil {
			return errcode.Wrap(errcode.ErrInternalServer, err)
		}

		// 创建新关联
		if len(menuIDs) > 0 {
			roleMenus := make([]*model.RoleMenu, len(menuIDs))
			for i, menuID := range menuIDs {
				roleMenus[i] = &model.RoleMenu{
					RoleID: roleID,
					MenuID: menuID,
				}
			}
			if err := tx.Create(&roleMenus).Error; err != nil {
				return errcode.Wrap(errcode.ErrInternalServer, err)
			}
		}

		return nil
	})
}
