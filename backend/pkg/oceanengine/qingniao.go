package oceanengine

import (
	"context"
	"encoding/json"
	"fmt"
)

// QingniaoService 青鸟线索通服务
type QingniaoService struct {
	client *Client
}

// NewQingniaoService 创建青鸟线索通服务
func NewQingniaoService(client *Client) *QingniaoService {
	return &QingniaoService{client: client}
}

// ==================== 表单管理 ====================

// QingniaoForm 青鸟表单
type QingniaoForm struct {
	FormID        int64               `json:"form_id"`
	AdvertiserID  int64               `json:"advertiser_id"`
	FormName      string              `json:"form_name"`
	FormType      string              `json:"form_type"`
	Status        int                 `json:"status"`
	CreateTime    string              `json:"create_time"`
	ModifyTime    string              `json:"modify_time"`
	FormElements  []QingniaoFormField `json:"form_elements"`
	ButtonText    string              `json:"button_text"`
	SubmitMessage string              `json:"submit_message"`
}

// QingniaoFormField 表单字段
type QingniaoFormField struct {
	FieldID     int64    `json:"field_id"`
	FieldType   string   `json:"field_type"`
	FieldName   string   `json:"field_name"`
	Required    bool     `json:"required"`
	Options     []string `json:"options,omitempty"`
	Placeholder string   `json:"placeholder,omitempty"`
	MaxLength   int      `json:"max_length,omitempty"`
}

// FormListRequest 获取表单列表请求
type FormListRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	FormName     string `json:"form_name,omitempty"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

// FormListResponse 获取表单列表响应
type FormListResponse struct {
	List     []QingniaoForm `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetFormList 获取表单列表
func (s *QingniaoService) GetFormList(ctx context.Context, req *FormListRequest) (*FormListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.FormName != "" {
		params["form_name"] = req.FormName
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/form/list/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result FormListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// FormDetailReq 获取表单详情请求
type FormDetailReq struct {
	AdvertiserID int64 `json:"advertiser_id"`
	FormID       int64 `json:"form_id"`
}

// GetFormDetail 获取表单详情
func (s *QingniaoService) GetFormDetail(ctx context.Context, req *FormDetailReq) (*QingniaoForm, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"form_id":       req.FormID,
	}

	resp, err := s.client.Get(ctx, "/2/clue/form/detail/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result QingniaoForm
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// FormCreateRequest 创建表单请求
type FormCreateRequest struct {
	AdvertiserID  int64               `json:"advertiser_id"`
	FormName      string              `json:"form_name"`
	FormType      string              `json:"form_type"`
	FormElements  []QingniaoFormField `json:"form_elements"`
	ButtonText    string              `json:"button_text,omitempty"`
	SubmitMessage string              `json:"submit_message,omitempty"`
}

// CreateForm 创建表单
func (s *QingniaoService) CreateForm(ctx context.Context, req *FormCreateRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/clue/form/create/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		FormID int64 `json:"form_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.FormID, nil
}

// FormUpdateRequest 更新表单请求
type FormUpdateRequest struct {
	AdvertiserID  int64               `json:"advertiser_id"`
	FormID        int64               `json:"form_id"`
	FormName      string              `json:"form_name,omitempty"`
	FormElements  []QingniaoFormField `json:"form_elements,omitempty"`
	ButtonText    string              `json:"button_text,omitempty"`
	SubmitMessage string              `json:"submit_message,omitempty"`
}

// UpdateForm 更新表单
func (s *QingniaoService) UpdateForm(ctx context.Context, req *FormUpdateRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/clue/form/update/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		FormID int64 `json:"form_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.FormID, nil
}

// FormDeleteRequest 删除表单请求
type FormDeleteRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	FormID       int64 `json:"form_id"`
}

// DeleteForm 删除表单
func (s *QingniaoService) DeleteForm(ctx context.Context, req *FormDeleteRequest) (bool, error) {
	resp, err := s.client.Post(ctx, "/2/clue/form/delete/", req)
	if err != nil {
		return false, err
	}

	if !resp.IsSuccess() {
		return false, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return true, nil
}

// ==================== 卡券管理 ====================

// Coupon 卡券信息
type Coupon struct {
	CouponID       int64  `json:"coupon_id"`
	AdvertiserID   int64  `json:"advertiser_id"`
	CouponName     string `json:"coupon_name"`
	CouponType     string `json:"coupon_type"`
	DiscountAmount int64  `json:"discount_amount"`
	ValidStartTime string `json:"valid_start_time"`
	ValidEndTime   string `json:"valid_end_time"`
	TotalCount     int    `json:"total_count"`
	UsedCount      int    `json:"used_count"`
	Status         int    `json:"status"`
	CreateTime     string `json:"create_time"`
	ModifyTime     string `json:"modify_time"`
}

// CouponCreateRequest 创建卡券请求
type CouponCreateRequest struct {
	AdvertiserID   int64  `json:"advertiser_id"`
	CouponName     string `json:"coupon_name"`
	CouponType     string `json:"coupon_type"`
	DiscountAmount int64  `json:"discount_amount,omitempty"`
	ValidStartTime string `json:"valid_start_time"`
	ValidEndTime   string `json:"valid_end_time"`
	TotalCount     int    `json:"total_count"`
	Description    string `json:"description,omitempty"`
	UseCondition   string `json:"use_condition,omitempty"`
}

// CreateCoupon 创建卡券
func (s *QingniaoService) CreateCoupon(ctx context.Context, req *CouponCreateRequest) (int64, error) {
	resp, err := s.client.Post(ctx, "/2/clue/coupon/create/", req)
	if err != nil {
		return 0, err
	}

	if !resp.IsSuccess() {
		return 0, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result struct {
		CouponID int64 `json:"coupon_id"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return result.CouponID, nil
}

// CouponListRequest 获取卡券列表请求
type CouponListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// CouponListResponse 获取卡券列表响应
type CouponListResponse struct {
	List     []Coupon `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetCouponList 获取卡券列表
func (s *QingniaoService) GetCouponList(ctx context.Context, req *CouponListRequest) (*CouponListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/coupon/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CouponListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CouponDetailRequest 获取卡券详情请求
type CouponDetailRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	CouponID     int64 `json:"coupon_id"`
}

// GetCouponDetail 获取卡券详情
func (s *QingniaoService) GetCouponDetail(ctx context.Context, req *CouponDetailRequest) (*Coupon, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"coupon_id":     req.CouponID,
	}

	resp, err := s.client.Get(ctx, "/2/clue/coupon/detail/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result Coupon
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CouponUpdateRequest 更新卡券请求
type CouponUpdateRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	CouponID     int64  `json:"coupon_id"`
	CouponName   string `json:"coupon_name,omitempty"`
	ValidEndTime string `json:"valid_end_time,omitempty"`
	TotalCount   int    `json:"total_count,omitempty"`
	Description  string `json:"description,omitempty"`
	UseCondition string `json:"use_condition,omitempty"`
}

// UpdateCoupon 更新卡券
func (s *QingniaoService) UpdateCoupon(ctx context.Context, req *CouponUpdateRequest) error {
	resp, err := s.client.Post(ctx, "/2/clue/coupon/update/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// CouponCodeUploadRequest 上传券码请求
type CouponCodeUploadRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	CouponID     int64    `json:"coupon_id"`
	Codes        []string `json:"codes"`
}

// CouponCodeUploadResponse 上传券码响应
type CouponCodeUploadResponse struct {
	SuccessCount int      `json:"success_count"`
	FailCount    int      `json:"fail_count"`
	FailCodes    []string `json:"fail_codes"`
}

// UploadCouponCode 上传券码
func (s *QingniaoService) UploadCouponCode(ctx context.Context, req *CouponCodeUploadRequest) (*CouponCodeUploadResponse, error) {
	resp, err := s.client.Post(ctx, "/2/clue/coupon/code/upload/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CouponCodeUploadResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CouponCode 券码记录
type CouponCode struct {
	Code         string `json:"code"`
	Status       int    `json:"status"`
	ConsumeTime  string `json:"consume_time"`
	EmployeeID   int64  `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
}

// CouponCodeGetRequest 查询券码记录请求
type CouponCodeGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	CouponID     int64 `json:"coupon_id"`
	Status       int   `json:"status,omitempty"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// CouponCodeGetResponse 查询券码记录响应
type CouponCodeGetResponse struct {
	List     []CouponCode `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetCouponCodes 查询券码记录
func (s *QingniaoService) GetCouponCodes(ctx context.Context, req *CouponCodeGetRequest) (*CouponCodeGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"coupon_id":     req.CouponID,
	}
	if req.Status > 0 {
		params["status"] = req.Status
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/coupon/code/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result CouponCodeGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// CouponCodeConsumeRequest 核销券码请求
type CouponCodeConsumeRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	CouponID     int64  `json:"coupon_id"`
	Code         string `json:"code"`
	EmployeeID   int64  `json:"employee_id,omitempty"`
}

// ConsumeCouponCode 核销券码
func (s *QingniaoService) ConsumeCouponCode(ctx context.Context, req *CouponCodeConsumeRequest) error {
	resp, err := s.client.Post(ctx, "/2/clue/coupon/code/consume/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ==================== 核销员管理 ====================

// Employee 核销员
type Employee struct {
	EmployeeID   int64  `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	PhoneNumber  string `json:"phone_number"`
	Status       int    `json:"status"`
	CreateTime   string `json:"create_time"`
}

// EmployeeGetRequest 查询核销员请求
type EmployeeGetRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	CouponID     int64 `json:"coupon_id"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// EmployeeGetResponse 查询核销员响应
type EmployeeGetResponse struct {
	List     []Employee `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetEmployees 查询核销员
func (s *QingniaoService) GetEmployees(ctx context.Context, req *EmployeeGetRequest) (*EmployeeGetResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"coupon_id":     req.CouponID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/coupon/employee/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result EmployeeGetResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// EmployeeCreateRequest 添加核销员请求
type EmployeeCreateRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	CouponID     int64  `json:"coupon_id"`
	EmployeeName string `json:"employee_name"`
	PhoneNumber  string `json:"phone_number"`
}

// EmployeeCreateResponse 添加核销员响应
type EmployeeCreateResponse struct {
	EmployeeID int64 `json:"employee_id"`
}

// CreateEmployee 添加核销员
func (s *QingniaoService) CreateEmployee(ctx context.Context, req *EmployeeCreateRequest) (*EmployeeCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/clue/coupon/employee/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result EmployeeCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// EmployeeDeleteRequest 删除核销员请求
type EmployeeDeleteRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	CouponID     int64 `json:"coupon_id"`
	EmployeeID   int64 `json:"employee_id"`
}

// DeleteEmployee 删除核销员
func (s *QingniaoService) DeleteEmployee(ctx context.Context, req *EmployeeDeleteRequest) error {
	resp, err := s.client.Post(ctx, "/2/clue/coupon/employee/delete/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// ==================== 智能电话管理 ====================

// QingniaoSmartPhone 智能电话
type QingniaoSmartPhone struct {
	SmartPhoneID int64  `json:"smart_phone_id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	AreaCode     string `json:"area_code"`
	BindStatus   int    `json:"bind_status"`
	CreateTime   string `json:"create_time"`
	ModifyTime   string `json:"modify_time"`
}

// SmartPhoneCreateRequest 创建智能电话请求
type SmartPhoneCreateRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	AreaCode     string `json:"area_code,omitempty"`
}

// SmartPhoneCreateResponse 创建智能电话响应
type SmartPhoneCreateResponse struct {
	SmartPhoneID int64 `json:"smart_phone_id"`
}

// CreateSmartPhone 创建智能电话
func (s *QingniaoService) CreateSmartPhone(ctx context.Context, req *SmartPhoneCreateRequest) (*SmartPhoneCreateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/clue/smartphone/create/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result SmartPhoneCreateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// QingniaoSmartPhoneListRequest 获取智能电话列表请求
type QingniaoSmartPhoneListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// QingniaoSmartPhoneListResponse 获取智能电话列表响应
type QingniaoSmartPhoneListResponse struct {
	List     []QingniaoSmartPhone `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetSmartPhoneList 获取智能电话列表
func (s *QingniaoService) GetSmartPhoneList(ctx context.Context, req *QingniaoSmartPhoneListRequest) (*QingniaoSmartPhoneListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/smartphone/get/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result QingniaoSmartPhoneListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// SmartPhoneDeleteRequest 删除智能电话请求
type SmartPhoneDeleteRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	SmartPhoneID int64 `json:"smart_phone_id"`
}

// DeleteSmartPhone 删除智能电话
func (s *QingniaoService) DeleteSmartPhone(ctx context.Context, req *SmartPhoneDeleteRequest) error {
	resp, err := s.client.Post(ctx, "/2/clue/smartphone/delete/", req)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	return nil
}

// SmartPhoneRecord 智能电话拨打记录
type SmartPhoneRecord struct {
	RecordID     int64  `json:"record_id"`
	SmartPhoneID int64  `json:"smart_phone_id"`
	CallerNumber string `json:"caller_number"`
	CalleeNumber string `json:"callee_number"`
	Duration     int    `json:"duration"`
	Status       int    `json:"status"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	RecordURL    string `json:"record_url"`
}

// SmartPhoneRecordRequest 查询拨打记录请求
type SmartPhoneRecordRequest struct {
	AdvertiserID int64  `json:"advertiser_id"`
	SmartPhoneID int64  `json:"smart_phone_id,omitempty"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

// SmartPhoneRecordResponse 查询拨打记录响应
type SmartPhoneRecordResponse struct {
	List     []SmartPhoneRecord `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetSmartPhoneRecords 查询智能电话拨打记录
func (s *QingniaoService) GetSmartPhoneRecords(ctx context.Context, req *SmartPhoneRecordRequest) (*SmartPhoneRecordResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"start_time":    req.StartTime,
		"end_time":      req.EndTime,
	}
	if req.SmartPhoneID > 0 {
		params["smart_phone_id"] = req.SmartPhoneID
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/smartphone/record/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result SmartPhoneRecordResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// ==================== 微信加粉组件 ====================

// WechatPool 微信库微信号
type WechatPool struct {
	WechatID   string `json:"wechat_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// WechatPoolListRequest 获取微信库微信号列表请求
type WechatPoolListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// WechatPoolListResponse 获取微信库微信号列表响应
type WechatPoolListResponse struct {
	List     []WechatPool `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetWechatPoolList 获取微信库微信号列表
func (s *QingniaoService) GetWechatPoolList(ctx context.Context, req *WechatPoolListRequest) (*WechatPoolListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/wechat/pool/list/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result WechatPoolListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// WechatInstance 微信号码包
type WechatInstance struct {
	InstanceID   int64        `json:"instance_id"`
	InstanceName string       `json:"instance_name"`
	WechatList   []WechatPool `json:"wechat_list"`
	Status       int          `json:"status"`
	CreateTime   string       `json:"create_time"`
	ModifyTime   string       `json:"modify_time"`
}

// WechatInstanceListRequest 获取微信号码包列表请求
type WechatInstanceListRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	Page         int   `json:"page,omitempty"`
	PageSize     int   `json:"page_size,omitempty"`
}

// WechatInstanceListResponse 获取微信号码包列表响应
type WechatInstanceListResponse struct {
	List     []WechatInstance `json:"list"`
	PageInfo struct {
		Page        int `json:"page"`
		PageSize    int `json:"page_size"`
		TotalNumber int `json:"total_number"`
		TotalPage   int `json:"total_page"`
	} `json:"page_info"`
}

// GetWechatInstanceList 获取微信号码包列表
func (s *QingniaoService) GetWechatInstanceList(ctx context.Context, req *WechatInstanceListRequest) (*WechatInstanceListResponse, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
	}
	if req.Page > 0 {
		params["page"] = req.Page
	}
	if req.PageSize > 0 {
		params["page_size"] = req.PageSize
	}

	resp, err := s.client.Get(ctx, "/2/clue/wechat/instance/list/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result WechatInstanceListResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// WechatInstanceDetailRequest 获取微信号码包详情请求
type WechatInstanceDetailRequest struct {
	AdvertiserID int64 `json:"advertiser_id"`
	InstanceID   int64 `json:"instance_id"`
}

// GetWechatInstanceDetail 获取微信号码包详情
func (s *QingniaoService) GetWechatInstanceDetail(ctx context.Context, req *WechatInstanceDetailRequest) (*WechatInstance, error) {
	params := map[string]interface{}{
		"advertiser_id": req.AdvertiserID,
		"instance_id":   req.InstanceID,
	}

	resp, err := s.client.Get(ctx, "/2/clue/wechat/instance/detail/", params)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result WechatInstance
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}

// WechatInstanceUpdateRequest 更新微信号码包请求
type WechatInstanceUpdateRequest struct {
	AdvertiserID int64    `json:"advertiser_id"`
	InstanceID   int64    `json:"instance_id"`
	InstanceName string   `json:"instance_name,omitempty"`
	WechatIDs    []string `json:"wechat_ids,omitempty"`
}

// WechatInstanceUpdateResponse 更新微信号码包响应
type WechatInstanceUpdateResponse struct {
	InstanceID int64 `json:"instance_id"`
}

// UpdateWechatInstance 更新微信号码包
func (s *QingniaoService) UpdateWechatInstance(ctx context.Context, req *WechatInstanceUpdateRequest) (*WechatInstanceUpdateResponse, error) {
	resp, err := s.client.Post(ctx, "/2/clue/wechat/instance/update/", req)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("api error: code=%d, message=%s", resp.Code, resp.Message)
	}

	var result WechatInstanceUpdateResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %w", err)
	}

	return &result, nil
}
