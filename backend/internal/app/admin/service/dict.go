package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"oceanengine-backend/internal/app/admin/dto"
	"oceanengine-backend/internal/app/admin/model"
	"oceanengine-backend/pkg/errcode"
)

// DictService 字典服务
type DictService struct {
	db *gorm.DB
}

// NewDictService 创建字典服务
func NewDictService(db *gorm.DB) *DictService {
	return &DictService{db: db}
}

// ========== 字典类型 ==========

// GetTypeList 获取字典类型列表
func (s *DictService) GetTypeList(ctx context.Context, req *dto.DictTypeListReq) ([]*dto.DictTypeResp, int64, error) {
	var types []*model.DictType
	var total int64

	query := s.db.WithContext(ctx).Model(&model.DictType{})

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Type != "" {
		query = query.Where("type LIKE ?", "%"+req.Type+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("id DESC").Find(&types).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.DictTypeResp, len(types))
	for i, t := range types {
		result[i] = &dto.DictTypeResp{
			ID:        t.ID,
			Name:      t.Name,
			Type:      t.Type,
			Status:    t.Status,
			Remark:    t.Remark,
			CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetTypeByID 获取字典类型详情
func (s *DictService) GetTypeByID(ctx context.Context, id uint64) (*dto.DictTypeResp, error) {
	var dictType model.DictType
	if err := s.db.WithContext(ctx).First(&dictType, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.DictTypeResp{
		ID:        dictType.ID,
		Name:      dictType.Name,
		Type:      dictType.Type,
		Status:    dictType.Status,
		Remark:    dictType.Remark,
		CreatedAt: dictType.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// CreateType 创建字典类型
func (s *DictService) CreateType(ctx context.Context, req *dto.DictTypeCreateReq) error {
	// 检查类型是否已存在
	var count int64
	s.db.WithContext(ctx).Model(&model.DictType{}).Where("type = ?", req.Type).Count(&count)
	if count > 0 {
		return errcode.New(errcode.ErrAlreadyExists)
	}

	dictType := &model.DictType{
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Remark: req.Remark,
	}

	if err := s.db.WithContext(ctx).Create(dictType).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// UpdateType 更新字典类型
func (s *DictService) UpdateType(ctx context.Context, req *dto.DictTypeUpdateReq) error {
	var dictType model.DictType
	if err := s.db.WithContext(ctx).First(&dictType, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	// 如果类型发生变化，检查是否冲突
	if req.Type != "" && req.Type != dictType.Type {
		var count int64
		s.db.WithContext(ctx).Model(&model.DictType{}).Where("type = ? AND id != ?", req.Type, req.ID).Count(&count)
		if count > 0 {
			return errcode.New(errcode.ErrAlreadyExists)
		}
	}

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Type != "" {
		updates["type"] = req.Type
	}
	updates["status"] = req.Status
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	if err := s.db.WithContext(ctx).Model(&dictType).Updates(updates).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// DeleteType 删除字典类型
func (s *DictService) DeleteType(ctx context.Context, id uint64) error {
	// 检查是否有关联的字典数据
	var count int64
	var dictType model.DictType
	if err := s.db.WithContext(ctx).First(&dictType, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	s.db.WithContext(ctx).Model(&model.DictData{}).Where("dict_type = ?", dictType.Type).Count(&count)
	if count > 0 {
		return errcode.NewWithMessage(errcode.ErrInvalidParams, "存在关联的字典数据，无法删除")
	}

	if err := s.db.WithContext(ctx).Delete(&model.DictType{}, id).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// ========== 字典数据 ==========

// GetDataList 获取字典数据列表
func (s *DictService) GetDataList(ctx context.Context, req *dto.DictDataListReq) ([]*dto.DictDataResp, int64, error) {
	var dataList []*model.DictData
	var total int64

	query := s.db.WithContext(ctx).Model(&model.DictData{}).Where("dict_type = ?", req.DictType)

	if req.Label != "" {
		query = query.Where("label LIKE ?", "%"+req.Label+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	offset := (req.GetPage() - 1) * req.GetPageSize()
	if err := query.Offset(offset).Limit(req.GetPageSize()).Order("sort ASC, id DESC").Find(&dataList).Error; err != nil {
		return nil, 0, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.DictDataResp, len(dataList))
	for i, d := range dataList {
		result[i] = &dto.DictDataResp{
			ID:        d.ID,
			DictType:  d.DictType,
			Label:     d.Label,
			Value:     d.Value,
			Sort:      d.Sort,
			Status:    d.Status,
			IsDefault: d.IsDefault,
			Remark:    d.Remark,
			CssClass:  d.CssClass,
			ListClass: d.ListClass,
			CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return result, total, nil
}

// GetDataByType 根据类型获取字典数据（不分页）
func (s *DictService) GetDataByType(ctx context.Context, dictType string) ([]*dto.DictDataResp, error) {
	var dataList []*model.DictData
	if err := s.db.WithContext(ctx).
		Where("dict_type = ? AND status = ?", dictType, model.DictStatusEnabled).
		Order("sort ASC, id DESC").
		Find(&dataList).Error; err != nil {
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	result := make([]*dto.DictDataResp, len(dataList))
	for i, d := range dataList {
		result[i] = &dto.DictDataResp{
			ID:        d.ID,
			DictType:  d.DictType,
			Label:     d.Label,
			Value:     d.Value,
			Sort:      d.Sort,
			Status:    d.Status,
			IsDefault: d.IsDefault,
			CssClass:  d.CssClass,
			ListClass: d.ListClass,
		}
	}

	return result, nil
}

// GetDataByID 获取字典数据详情
func (s *DictService) GetDataByID(ctx context.Context, id uint64) (*dto.DictDataResp, error) {
	var data model.DictData
	if err := s.db.WithContext(ctx).First(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.New(errcode.ErrNotFound)
		}
		return nil, errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return &dto.DictDataResp{
		ID:        data.ID,
		DictType:  data.DictType,
		Label:     data.Label,
		Value:     data.Value,
		Sort:      data.Sort,
		Status:    data.Status,
		IsDefault: data.IsDefault,
		Remark:    data.Remark,
		CssClass:  data.CssClass,
		ListClass: data.ListClass,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// CreateData 创建字典数据
func (s *DictService) CreateData(ctx context.Context, req *dto.DictDataCreateReq) error {
	// 检查字典类型是否存在
	var count int64
	s.db.WithContext(ctx).Model(&model.DictType{}).Where("type = ?", req.DictType).Count(&count)
	if count == 0 {
		return errcode.NewWithMessage(errcode.ErrNotFound, "字典类型不存在")
	}

	data := &model.DictData{
		DictType:  req.DictType,
		Label:     req.Label,
		Value:     req.Value,
		Sort:      req.Sort,
		Status:    req.Status,
		IsDefault: req.IsDefault,
		Remark:    req.Remark,
		CssClass:  req.CssClass,
		ListClass: req.ListClass,
	}

	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// UpdateData 更新字典数据
func (s *DictService) UpdateData(ctx context.Context, req *dto.DictDataUpdateReq) error {
	var data model.DictData
	if err := s.db.WithContext(ctx).First(&data, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errcode.New(errcode.ErrNotFound)
		}
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	updates := map[string]interface{}{
		"sort":       req.Sort,
		"status":     req.Status,
		"is_default": req.IsDefault,
	}
	if req.Label != "" {
		updates["label"] = req.Label
	}
	if req.Value != "" {
		updates["value"] = req.Value
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}
	if req.CssClass != "" {
		updates["css_class"] = req.CssClass
	}
	if req.ListClass != "" {
		updates["list_class"] = req.ListClass
	}

	if err := s.db.WithContext(ctx).Model(&data).Updates(updates).Error; err != nil {
		return errcode.Wrap(errcode.ErrInternalServer, err)
	}

	return nil
}

// DeleteData 删除字典数据
func (s *DictService) DeleteData(ctx context.Context, id uint64) error {
	result := s.db.WithContext(ctx).Delete(&model.DictData{}, id)
	if result.Error != nil {
		return errcode.Wrap(errcode.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return errcode.New(errcode.ErrNotFound)
	}
	return nil
}
