package gowechat

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// SubmitAudit 提交审核
type MiniProgramSubmitAuditRequest struct {
	ItemList      []MiniProgramAuditItem      `json:"item_list,omitempty"`      // 审核项列表（选填，至多填写 5 项）
	PreviewInfo   MiniProgramAuditPreviewInfo `json:"preview_info,omitempty"`   // 预览信息（小程序页面截图和操作录屏）
	VersionDesc   string                      `json:"version_desc,omitempty"`   // 小程序版本说明和功能解释
	FeedbackInfo  string                      `json:"feedback_info,omitempty"`  // 反馈内容，至多 200 字
	FeedbackStuff string                      `json:"feedback_stuff,omitempty"` // 用 | 分割的 media_id 列表，至多 5 张图片, 可以通过新增临时素材接口上传而得到
	UGCDeclare    MiniProgramAuditUGCDeclare  `json:"ugc_declare,omitempty"`    // 用户生成内容场景（UGC）信息安全声明
}

type MiniProgramSubmitAuditResponse struct {
	util.CommonError
	AuditID int `json:"auditid,omitempty"` // 审核编号
}

func (miniprogram *MiniProgram) SubmitAudit(req *MiniProgramSubmitAuditRequest) (*MiniProgramSubmitAuditResponse, error) {
	accessToken, err := miniprogram.MiniProgram.GetContext().GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/submit_audit?access_token=%s", accessToken)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	resp := &MiniProgramSubmitAuditResponse{}
	err = util.DecodeWithError(response, resp, "SubmitAudit")
	return resp, err
}

type MiniProgramAuditItem struct {
	Address     string `json:"address,omitempty"`      // 小程序的页面，可通过获取小程序的页面列表接口获得
	Tag         string `json:"tag,omitempty"`          // 小程序的标签，用空格分隔，标签至多 10 个，标签长度至多 20
	FirstClass  string `json:"first_class,omitempty"`  // 一级类目名称
	SecondClass string `json:"second_class,omitempty"` // 二级类目名称
	ThirdClass  string `json:"third_class,omitempty"`  // 三级类目名称
	FirstID     string `json:"first_id,omitempty"`     // 一级类目的 ID
	SecondID    string `json:"second_id,omitempty"`    // 二级类目的 ID
	ThirdID     string `json:"third_id,omitempty"`     // 三级类目的 ID
	Title       string `json:"title,omitempty"`        // 小程序页面的标题,标题长度至多 32
}

type MiniProgramAuditPreviewInfo struct {
	VideoIDList []string `json:"video_id_list,omitempty"` // 录屏mediaid列表，可以通过提审素材上传接口获得
	PicIDList   []string `json:"pic_id_list,omitempty"`   // 截屏mediaid列表，可以通过提审素材上传接口获得
}

type MiniProgramAuditUGCDeclare struct {
	Scene          []int  `json:"scene,omitempty"`            // UGC场景 0,不涉及用户生成内容, 1.用户资料,2.图片,3.视频,4.文本,5其他, 可多选,当scene填0时无需填写下列字段
	OtherSceneDesc string `json:"other_scene_desc,omitempty"` // 当scene选其他时的说明,不超时256字
	Method         []int  `json:"method,omitempty"`           // 内容安全机制 1.使用平台建议的内容安全API,2.使用其他的内容审核产品,3.通过人工审核把关,4.未做内容审核把关
	HasAuditTeam   int    `json:"has_audit_team,omitempty"`   // 是否有审核团队, 0.无,1.有,默认0
	AuditDesc      string `json:"audit_desc,omitempty"`       // 说明当前对UGC内容的审核机制,不超过256字
}
