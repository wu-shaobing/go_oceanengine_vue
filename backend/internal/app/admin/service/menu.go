package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/errcode"
)

// MenuService 菜单服务
type MenuService struct {
	db *gorm.DB
}

// NewMenuService 创建菜单服务
func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{db: db}
}

// GetTree 获取菜单树
func (s *MenuService) GetTree(ctx context.Context) ([]*dto.MenuTree, error) {
	var menus []*model.Menu
	if err := s.db.WithContext(ctx).Order("sort ASC, id ASC").Find(&menus).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.buildMenuTree(menus, 0), nil
}

// GetUserMenuTree 获取用户菜单树（根据角色权限）
func (s *MenuService) GetUserMenuTree(ctx context.Context, roleID uint64) ([]*dto.MenuTree, error) {
	// 获取角色菜单ID列表
	var roleMenus []model.RoleMenu
	if err := s.db.WithContext(ctx).Where("role_id = ?", roleID).Find(&roleMenus).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	if len(roleMenus) == 0 {
		return []*dto.MenuTree{}, nil
	}

	menuIDs := make([]uint64, len(roleMenus))
	for i, rm := range roleMenus {
		menuIDs[i] = rm.MenuID
	}

	// 获取菜单
	var menus []*model.Menu
	if err := s.db.WithContext(ctx).Where("id IN ? AND status = ?", menuIDs, model.UserStatusEnabled).Order("sort ASC, id ASC").Find(&menus).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return s.buildMenuTree(menus, 0), nil
}

// buildMenuTree 构建菜单树
func (s *MenuService) buildMenuTree(menus []*model.Menu, parentID uint64) []*dto.MenuTree {
	var tree []*dto.MenuTree
	for _, menu := range menus {
		if menu.ParentID == parentID {
			// Visible=1 表示可见，转换为 Hidden=0
			hidden := int8(0)
			if menu.Visible == 0 {
				hidden = 1
			}
			node := &dto.MenuTree{
				ID:         menu.ID,
				ParentID:   menu.ParentID,
				Name:       menu.Name,
				Path:       menu.Path,
				Component:  menu.Component,
				Icon:       menu.Icon,
				Sort:       menu.Sort,
				Type:       menu.Type,
				Permission: menu.Permission,
				Status:     menu.Status,
				Hidden:     hidden,
				Children:   s.buildMenuTree(menus, menu.ID),
			}
			tree = append(tree, node)
		}
	}
	return tree
}

// GetList 获取菜单列表（扁平）
func (s *MenuService) GetList(ctx context.Context) ([]*model.Menu, error) {
	var menus []*model.Menu
	if err := s.db.WithContext(ctx).Order("sort ASC, id ASC").Find(&menus).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return menus, nil
}

// GetByID 获取菜单详情
func (s *MenuService) GetByID(ctx context.Context, id uint64) (*model.Menu, error) {
	var menu model.Menu
	if err := s.db.WithContext(ctx).First(&menu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrMenuNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}
	return &menu, nil
}

// Create 创建菜单
func (s *MenuService) Create(ctx context.Context, req *dto.MenuCreateReq, operatorID uint64) error {
	// Hidden 转换为 Visible: Hidden=1 -> Visible=0
	visible := int8(1)
	if req.Hidden == 1 {
		visible = 0
	}
	menu := &model.Menu{
		ParentID:   req.ParentID,
		Name:       req.Name,
		Path:       req.Path,
		Component:  req.Component,
		Icon:       req.Icon,
		Sort:       req.Sort,
		Type:       req.Type,
		Permission: req.Permission,
		Status:     req.Status,
		Visible:    visible,
		Remark:     req.Remark,
	}

	if err := s.db.WithContext(ctx).Create(menu).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Update 更新菜单
func (s *MenuService) Update(ctx context.Context, req *dto.MenuUpdateReq, operatorID uint64) error {
	var menu model.Menu
	if err := s.db.WithContext(ctx).First(&menu, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrMenuNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 不能将菜单的父级设为自己
	if req.ParentID == req.ID {
		return errcode.New(errcode.ErrInvalidParams)
	}

	// Hidden 转换为 Visible
	visible := int8(1)
	if req.Hidden == 1 {
		visible = 0
	}
	updates := map[string]interface{}{
		"parent_id":  req.ParentID,
		"name":       req.Name,
		"path":       req.Path,
		"component":  req.Component,
		"icon":       req.Icon,
		"sort":       req.Sort,
		"type":       req.Type,
		"permission": req.Permission,
		"status":     req.Status,
		"visible":    visible,
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := s.db.WithContext(ctx).Model(&menu).Updates(updates).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// Delete 删除菜单
func (s *MenuService) Delete(ctx context.Context, id uint64) error {
	// 检查是否有子菜单
	var childCount int64
	if err := s.db.WithContext(ctx).Model(&model.Menu{}).Where("parent_id = ?", id).Count(&childCount).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}
	if childCount > 0 {
		return errcode.New(errcode.ErrMenuHasChildren)
	}

	// 删除角色菜单关联
	if err := s.db.WithContext(ctx).Where("menu_id = ?", id).Delete(&model.RoleMenu{}).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 删除菜单
	result := s.db.WithContext(ctx).Delete(&model.Menu{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrMenuNotFound)
	}
	return nil
}
