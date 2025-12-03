package oceanengine

import (
	"context"
)

// EnterpriseClient 企业号API客户端
type EnterpriseClient struct {
	client *Client
}

// NewEnterpriseClient 创建企业号客户端
func (c *Client) Enterprise() *EnterpriseClient {
	return &EnterpriseClient{client: c}
}

// ==================== 账号信息 ====================

// EnterpriseInfo 企业号信息
type EnterpriseInfo struct {
	DouyinID      string `json:"douyin_id"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	FollowerCount int64  `json:"follower_count"`
	TotalFav      int64  `json:"total_favorited"`
	VideoCount    int64  `json:"aweme_count"`
	IsVerified    bool   `json:"is_verified"`
}

// GetInfo 获取企业号信息
func (e *EnterpriseClient) GetInfo(ctx context.Context, accessToken string, douyinID string) (*EnterpriseInfo, error) {
	path := "/enterprise/info/"
	params := map[string]interface{}{
		"open_id": douyinID,
	}

	var result struct {
		Data EnterpriseInfo `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 绑定管理 ====================

// BindAccount 绑定账号信息
type BindAccount struct {
	DouyinID   string `json:"douyin_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	BindTime   string `json:"bind_time"`
	BindStatus int    `json:"bind_status"`
}

// GetBindList 获取绑定列表
func (e *EnterpriseClient) GetBindList(ctx context.Context, accessToken string, advertiserID uint64) ([]BindAccount, error) {
	path := "/enterprise/bind/list/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
	}

	var result struct {
		Data struct {
			List []BindAccount `json:"list"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 评论管理 ====================

// Comment 评论信息
type Comment struct {
	CommentID  string `json:"comment_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UserID     string `json:"user_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	LikeCount  int64  `json:"like_count"`
	ReplyCount int64  `json:"reply_count"`
	IsTop      bool   `json:"is_top"`
	ItemID     string `json:"item_id"` // 关联视频ID
}

// GetCommentList 获取评论列表
func (e *EnterpriseClient) GetCommentList(ctx context.Context, accessToken string, openID string, itemID string, cursor int64, count int) ([]Comment, int64, bool, error) {
	path := "/enterprise/comment/list/"
	params := map[string]interface{}{
		"open_id": openID,
		"item_id": itemID,
		"cursor":  cursor,
		"count":   count,
	}

	var result struct {
		Data struct {
			List    []Comment `json:"list"`
			Cursor  int64     `json:"cursor"`
			HasMore bool      `json:"has_more"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, false, err
	}
	return result.Data.List, result.Data.Cursor, result.Data.HasMore, nil
}

// ReplyComment 回复评论
func (e *EnterpriseClient) ReplyComment(ctx context.Context, accessToken string, openID string, itemID string, commentID string, content string) (string, error) {
	path := "/enterprise/comment/reply/"
	data := map[string]interface{}{
		"open_id":    openID,
		"item_id":    itemID,
		"comment_id": commentID,
		"content":    content,
	}

	var result struct {
		Data struct {
			CommentID string `json:"comment_id"`
		} `json:"data"`
	}
	err := e.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return "", err
	}
	return result.Data.CommentID, nil
}

// TopComment 置顶评论
func (e *EnterpriseClient) TopComment(ctx context.Context, accessToken string, openID string, itemID string, commentID string, isTop bool) error {
	path := "/enterprise/comment/top/"
	data := map[string]interface{}{
		"open_id":    openID,
		"item_id":    itemID,
		"comment_id": commentID,
		"top":        isTop,
	}
	return e.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 视频管理 ====================

// Video 视频信息
type Video struct {
	ItemID       string `json:"item_id"`
	Title        string `json:"title"`
	Cover        string `json:"cover"`
	VideoStatus  int    `json:"video_status"`
	CreateTime   string `json:"create_time"`
	PlayCount    int64  `json:"play_count"`
	DiggCount    int64  `json:"digg_count"`
	CommentCount int64  `json:"comment_count"`
	ShareCount   int64  `json:"share_count"`
}

// GetVideoList 获取视频列表
func (e *EnterpriseClient) GetVideoList(ctx context.Context, accessToken string, openID string, cursor int64, count int) ([]Video, int64, bool, error) {
	path := "/enterprise/item/list/"
	params := map[string]interface{}{
		"open_id": openID,
		"cursor":  cursor,
		"count":   count,
	}

	var result struct {
		Data struct {
			List    []Video `json:"list"`
			Cursor  int64   `json:"cursor"`
			HasMore bool    `json:"has_more"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, false, err
	}
	return result.Data.List, result.Data.Cursor, result.Data.HasMore, nil
}

// ==================== 数据概览 ====================

// OverviewData 数据概览
type OverviewData struct {
	Date          string `json:"date"`
	NewFollower   int64  `json:"new_follower"`
	TotalFollower int64  `json:"total_follower"`
	NewVideo      int64  `json:"new_video"`
	VideoView     int64  `json:"video_view"`
	VideoLike     int64  `json:"video_like"`
	VideoComment  int64  `json:"video_comment"`
	VideoShare    int64  `json:"video_share"`
	ProfileView   int64  `json:"profile_view"`
}

// GetOverviewData 获取数据概览
func (e *EnterpriseClient) GetOverviewData(ctx context.Context, accessToken string, openID string, dateType string) (*OverviewData, error) {
	path := "/enterprise/data/overview/"
	params := map[string]interface{}{
		"open_id":   openID,
		"date_type": dateType, // 7, 15, 30
	}

	var result struct {
		Data OverviewData `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 操作日志 ====================

// OperationLog 操作日志
type OperationLog struct {
	LogID         string `json:"log_id"`
	OperatorID    string `json:"operator_id"`
	OperatorName  string `json:"operator_name"`
	OperationType string `json:"operation_type"`
	OperationTime string `json:"operation_time"`
	Detail        string `json:"detail"`
}

// GetOperationLog 获取操作日志
func (e *EnterpriseClient) GetOperationLog(ctx context.Context, accessToken string, advertiserID uint64, startDate, endDate string, page, pageSize int) ([]OperationLog, int, error) {
	path := "/enterprise/operation/log/"
	params := map[string]interface{}{
		"advertiser_id": advertiserID,
		"start_date":    startDate,
		"end_date":      endDate,
		"page":          page,
		"page_size":     pageSize,
	}

	var result struct {
		Data struct {
			List     []OperationLog `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 流量来源 ====================

// FlowSourceData 流量来源数据
type FlowSourceData struct {
	SourceName string  `json:"source_name"`
	Count      int64   `json:"count"`
	Ratio      float64 `json:"ratio"`
}

// GetFlowCategoryData 获取流量来源数据
func (e *EnterpriseClient) GetFlowCategoryData(ctx context.Context, accessToken string, openID string, dateType string) ([]FlowSourceData, error) {
	path := "/enterprise/data/flow/category/"
	params := map[string]interface{}{
		"open_id":   openID,
		"date_type": dateType,
	}

	var result struct {
		Data struct {
			List []FlowSourceData `json:"list"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 视频分析 ====================

// VideoAnalytics 视频分析数据
type VideoAnalytics struct {
	ItemID          string  `json:"item_id"`
	Title           string  `json:"title"`
	PlayCount       int64   `json:"play_count"`
	DiggCount       int64   `json:"digg_count"`
	CommentCount    int64   `json:"comment_count"`
	ShareCount      int64   `json:"share_count"`
	AvgPlayDuration float64 `json:"avg_play_duration"`
	FinishRate      float64 `json:"finish_rate"`
}

// GetVideoAnalytics 获取视频分析数据
func (e *EnterpriseClient) GetVideoAnalytics(ctx context.Context, accessToken string, openID string, itemID string) (*VideoAnalytics, error) {
	path := "/enterprise/data/item/"
	params := map[string]interface{}{
		"open_id": openID,
		"item_id": itemID,
	}

	var result struct {
		Data VideoAnalytics `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 评论回复列表 ====================

// CommentReply 评论回复
type CommentReply struct {
	ReplyID    string `json:"reply_id"`
	CommentID  string `json:"comment_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UserID     string `json:"user_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
}

// GetCommentReplyList 获取评论回复列表
func (e *EnterpriseClient) GetCommentReplyList(ctx context.Context, accessToken string, openID string, itemID string, commentID string, cursor int64, count int) ([]CommentReply, int64, bool, error) {
	path := "/enterprise/comment/reply/list/"
	params := map[string]interface{}{
		"open_id":    openID,
		"item_id":    itemID,
		"comment_id": commentID,
		"cursor":     cursor,
		"count":      count,
	}

	var result struct {
		Data struct {
			List    []CommentReply `json:"list"`
			Cursor  int64          `json:"cursor"`
			HasMore bool           `json:"has_more"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, false, err
	}
	return result.Data.List, result.Data.Cursor, result.Data.HasMore, nil
}

// ==================== 评论详情 ====================

// CommentDetail 评论详情
type CommentDetail struct {
	CommentID  string `json:"comment_id"`
	ItemID     string `json:"item_id"`
	ItemTitle  string `json:"item_title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UserID     string `json:"user_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	LikeCount  int64  `json:"like_count"`
	ReplyCount int64  `json:"reply_count"`
}

// GetCommentDetail 获取评论详情
func (e *EnterpriseClient) GetCommentDetail(ctx context.Context, accessToken string, openID string, commentIDs []string) ([]CommentDetail, error) {
	path := "/enterprise/comment/detail/"
	params := map[string]interface{}{
		"open_id":     openID,
		"comment_ids": commentIDs,
	}

	var result struct {
		Data struct {
			List []CommentDetail `json:"list"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return result.Data.List, nil
}

// ==================== 纵横组织账户 ====================

// MajordomoAdvertiser 纵横组织下的广告主
type MajordomoAdvertiser struct {
	AdvertiserID   uint64 `json:"advertiser_id"`
	AdvertiserName string `json:"advertiser_name"`
	Company        string `json:"company"`
	Status         int    `json:"status"`
}

// GetMajordomoAdvertiserList 获取纵横组织下广告主列表
func (e *EnterpriseClient) GetMajordomoAdvertiserList(ctx context.Context, accessToken string, majordomoID uint64, page, pageSize int) ([]MajordomoAdvertiser, int, error) {
	path := "/majordomo/advertiser/select/"
	params := map[string]interface{}{
		"majordomo_id": majordomoID,
		"page":         page,
		"page_size":    pageSize,
	}

	var result struct {
		Data struct {
			List     []MajordomoAdvertiser `json:"list"`
			PageInfo struct {
				TotalNumber int `json:"total_number"`
			} `json:"page_info"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, err
	}
	return result.Data.List, result.Data.PageInfo.TotalNumber, nil
}

// ==================== 私信管理 ====================

// PrivateMessage 私信消息
type PrivateMessage struct {
	MsgID      string `json:"msg_id"`
	Content    string `json:"content"`
	MsgType    string `json:"msg_type"`
	SendTime   string `json:"send_time"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	IsRead     bool   `json:"is_read"`
}

// SendPrivateMessage 发送私信
func (e *EnterpriseClient) SendPrivateMessage(ctx context.Context, accessToken string, openID string, toUserID string, msgType string, content string) (string, error) {
	path := "/enterprise/im/message/send/"
	data := map[string]interface{}{
		"open_id":    openID,
		"to_user_id": toUserID,
		"msg_type":   msgType,
		"content":    content,
	}

	var result struct {
		Data struct {
			MsgID string `json:"msg_id"`
		} `json:"data"`
	}
	err := e.client.PostWithToken(ctx, accessToken, path, data, &result)
	if err != nil {
		return "", err
	}
	return result.Data.MsgID, nil
}

// GetPrivateMessageList 获取私信列表
func (e *EnterpriseClient) GetPrivateMessageList(ctx context.Context, accessToken string, openID string, conversationID string, cursor int64, count int) ([]PrivateMessage, int64, bool, error) {
	path := "/enterprise/im/message/list/"
	params := map[string]interface{}{
		"open_id":         openID,
		"conversation_id": conversationID,
		"cursor":          cursor,
		"count":           count,
	}

	var result struct {
		Data struct {
			List    []PrivateMessage `json:"list"`
			Cursor  int64            `json:"cursor"`
			HasMore bool             `json:"has_more"`
		} `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, 0, false, err
	}
	return result.Data.List, result.Data.Cursor, result.Data.HasMore, nil
}

// ==================== 粉丝分析 ====================

// FansPortrait 粉丝画像
type FansPortrait struct {
	GenderRatio map[string]float64 `json:"gender_ratio"`
	AgeRatio    map[string]float64 `json:"age_ratio"`
	ProvinceTop []ProvinceData     `json:"province_top"`
	CityTop     []CityData         `json:"city_top"`
	DeviceRatio map[string]float64 `json:"device_ratio"`
	InterestTop []InterestData     `json:"interest_top"`
}

type ProvinceData struct {
	Province string  `json:"province"`
	Ratio    float64 `json:"ratio"`
}

type CityData struct {
	City  string  `json:"city"`
	Ratio float64 `json:"ratio"`
}

type InterestData struct {
	Interest string  `json:"interest"`
	Ratio    float64 `json:"ratio"`
}

// GetFansPortrait 获取粉丝画像
func (e *EnterpriseClient) GetFansPortrait(ctx context.Context, accessToken string, openID string) (*FansPortrait, error) {
	path := "/enterprise/data/fans/portrait/"
	params := map[string]interface{}{
		"open_id": openID,
	}

	var result struct {
		Data FansPortrait `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// ==================== 视频置顶 ====================

// SetVideoTop 设置视频置顶
func (e *EnterpriseClient) SetVideoTop(ctx context.Context, accessToken string, openID string, itemID string, isTop bool) error {
	path := "/enterprise/item/top/"
	data := map[string]interface{}{
		"open_id": openID,
		"item_id": itemID,
		"is_top":  isTop,
	}
	return e.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteVideo 删除视频
func (e *EnterpriseClient) DeleteVideo(ctx context.Context, accessToken string, openID string, itemID string) error {
	path := "/enterprise/item/delete/"
	data := map[string]interface{}{
		"open_id": openID,
		"item_id": itemID,
	}
	return e.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// ==================== 评论操作 ====================

// UpdateCommentReply 更新评论回复
func (e *EnterpriseClient) UpdateCommentReply(ctx context.Context, accessToken string, openID string, commentID string, content string) error {
	path := "/enterprise/comment/reply/update/"
	data := map[string]interface{}{
		"open_id":    openID,
		"comment_id": commentID,
		"content":    content,
	}
	return e.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// HideComment 隐藏评论
func (e *EnterpriseClient) HideComment(ctx context.Context, accessToken string, openID string, commentID string, isHide bool) error {
	path := "/enterprise/comment/hide/"
	data := map[string]interface{}{
		"open_id":    openID,
		"comment_id": commentID,
		"is_hide":    isHide,
	}
	return e.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// DeleteComment 删除评论
func (e *EnterpriseClient) DeleteComment(ctx context.Context, accessToken string, openID string, commentID string) error {
	path := "/enterprise/comment/delete/"
	data := map[string]interface{}{
		"open_id":    openID,
		"comment_id": commentID,
	}
	return e.client.PostWithToken(ctx, accessToken, path, data, nil)
}

// GetOverviewDataByDateRange 按日期范围获取数据概览
func (e *EnterpriseClient) GetOverviewDataByDateRange(ctx context.Context, accessToken string, openID string, startDate string, endDate string) (*OverviewData, error) {
	path := "/enterprise/data/overview/"
	params := map[string]interface{}{
		"open_id":    openID,
		"start_date": startDate,
		"end_date":   endDate,
	}

	var result struct {
		Data OverviewData `json:"data"`
	}
	err := e.client.GetWithToken(ctx, accessToken, path, params, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}
