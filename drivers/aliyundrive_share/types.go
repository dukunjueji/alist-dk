package aliyundrive_share

import (
	"time"

	"github.com/alist-org/alist/v3/internal/model"
)

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ShareTokenResp struct {
	ShareToken string    `json:"share_token"`
	ExpireTime time.Time `json:"expire_time"`
	ExpiresIn  int       `json:"expires_in"`
}

type ListResp struct {
	Items             []File `json:"items"`
	NextMarker        string `json:"next_marker"`
	PunishedFileCount int    `json:"punished_file_count"`
}

type File struct {
	DriveId      string    `json:"drive_id"`
	DomainId     string    `json:"domain_id"`
	FileId       string    `json:"file_id"`
	ShareId      string    `json:"share_id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ParentFileId string    `json:"parent_file_id"`
	Size         int64     `json:"size"`
	Thumbnail    string    `json:"thumbnail"`
}

func fileToObj(f File) *model.ObjThumb {
	return &model.ObjThumb{
		Object: model.Object{
			ID:       f.FileId,
			Name:     f.Name,
			Size:     f.Size,
			Modified: f.UpdatedAt,
			IsFolder: f.Type == "folder",
		},
		Thumbnail: model.Thumbnail{Thumbnail: f.Thumbnail},
	}
}

type ShareLinkResp struct {
	DownloadUrl string `json:"download_url"`
	Url         string `json:"url"`
	Thumbnail   string `json:"thumbnail"`
}
type VideoPreviewPlayInfo struct {
	Category             string            `json:"category"`
	Meta                 VideoMeta         `json:"meta"`
	LiveTranscodingTasks []TranscodingTask `json:"live_transcoding_task_list"`
}

type VideoMeta struct {
	Duration float64 `json:"duration"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
}

type TranscodingTask struct {
	TemplateID             string `json:"template_id"`
	TemplateName           string `json:"template_name"`
	TemplateWidth          int    `json:"template_width"`
	TemplateHeight         int    `json:"template_height"`
	Status                 string `json:"status"`
	Stage                  string `json:"stage"`
	PreviewURL             string `json:"preview_url"`
	Url                    string `json:"url"`
	KeepOriginalResolution bool   `json:"keep_original_resolution,omitempty"`
}

type VideoInfo struct {
	ShareID                     string               `json:"share_id"`
	FileID                      string               `json:"file_id"`
	Category                    string               `json:"category"`
	VideoPreviewPlayInfo        VideoPreviewPlayInfo `json:"video_preview_play_info"`
	PunishFlag                  int                  `json:"punish_flag"`
	MetaNamePunishFlag          int                  `json:"meta_name_punish_flag"`
	MetaNameInvestigationStatus int                  `json:"meta_name_investigation_status"`
}
type ResponseBody struct {
	DomainID string `json:"domain_id"`
	DriveID  string `json:"drive_id"`
	FileID   string `json:"file_id"`
}

type Response struct {
	Body   ResponseBody `json:"body"`
	ID     string       `json:"id"`
	Status int          `json:"status"`
}
type User struct {
	DriveId string `json:"default_drive_id"`
}
type DistributorCouponInfo struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	ButtonText      string `json:"buttonText"`
	ButtonSchemaURL string `json:"buttonSchemaUrl"`
	DisplayValidity string `json:"displayValidity"`
	MaxSaving       string `json:"maxSaving"`
	DisplayCurrency string `json:"displayCurrency"`
}

type CopyFileRes struct {
	Responses             []Response            `json:"responses"`
	DistributorCouponInfo DistributorCouponInfo `json:"distributorCouponInfo"`
}
