package aliyundrive_share

import (
	"context"
	"fmt"
	"github.com/Xhofe/rateg"
	"github.com/alist-org/alist/v3/drivers/base"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
	"github.com/alist-org/alist/v3/pkg/cron"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AliyundriveShare struct {
	model.Storage
	Addition
	AccessToken string
	ShareToken  string
	DriveId     string
	UserDriveId string
	cron        *cron.Cron

	limitList func(ctx context.Context, dir model.Obj) ([]model.Obj, error)
	limitLink func(ctx context.Context, file model.Obj) (*model.Link, error)
}

func (d *AliyundriveShare) GetDriveId() string {
	//TODO implement me
	panic("implement me")
}

func (d *AliyundriveShare) Config() driver.Config {
	return config
}

func (d *AliyundriveShare) GetAddition() driver.Additional {
	return &d.Addition
}

func (d *AliyundriveShare) Init(ctx context.Context) error {
	err := d.refreshToken()
	if err != nil {
		return err
	}
	err = d.getShareToken()
	if err != nil {
		return err
	}
	d.cron = cron.NewCron(time.Hour * 2)
	d.cron.Do(func() {
		err := d.refreshToken()
		if err != nil {
			log.Errorf("%+v", err)
		}
	})
	d.limitList = rateg.LimitFnCtx(d.list, rateg.LimitFnOption{
		Limit:  4,
		Bucket: 1,
	})
	d.limitLink = rateg.LimitFnCtx(d.link, rateg.LimitFnOption{
		Limit:  1,
		Bucket: 1,
	})
	var res User
	_, _ = d.request("https://api.aliyundrive.com/v2/user/get", http.MethodPost, func(req *resty.Request) {
		req.SetBody(base.Json{}).SetResult(&res)
	})
	d.UserDriveId = res.DriveId
	// Unmarshal JSON data into Json (map[string]interface{})
	return nil
}

func (d *AliyundriveShare) Drop(ctx context.Context) error {
	if d.cron != nil {
		d.cron.Stop()
	}
	d.DriveId = ""
	return nil
}

func (d *AliyundriveShare) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	if d.limitList == nil {
		return nil, fmt.Errorf("driver not init")
	}
	return d.limitList(ctx, dir)
}

func (d *AliyundriveShare) list(ctx context.Context, dir model.Obj) ([]model.Obj, error) {
	files, err := d.getFiles(dir.GetID())
	if err != nil {
		return nil, err
	}
	return utils.SliceConvert(files, func(src File) (model.Obj, error) {
		return fileToObj(src), nil
	})
}

func (d *AliyundriveShare) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	if d.limitLink == nil {
		return nil, fmt.Errorf("driver not init")
	}
	return d.limitLink(ctx, file)
}

func (d *AliyundriveShare) link(ctx context.Context, file model.Obj) (*model.Link, error) {
	//data := base.Json{
	//	"category":          "live_transcoding",
	//	"get_preview_url":   true,
	//	"get_subtitle_info": true,
	//	"template_id":       "",
	//	"file_id":           file.GetID(),
	//	// // Only ten minutes lifetime
	//	"share_id": d.ShareId,
	//}
	FileRes, _ := CopyFile(file, d)

	link, err := GetOpenDwnUrl(ctx, FileRes)

	log.Debugf("link:%s,err%s", link.URL, err)

	DeleteFile(d, FileRes)

	return &model.Link{
		Header: http.Header{
			"Referer": []string{"https://www.aliyundrive.com/"},
		},
		URL: link.URL,
	}, nil
}

func GetOpenDwnUrl(ctx context.Context, FileRes CopyFileRes) (*model.Link, error) {
	storageAilYunOpen, _, _ := op.GetStorageAndActualPath("/root")
	var rootObj model.Obj
	rootObj = &model.Object{
		ID:       FileRes.Responses[0].Body.FileID,
		Name:     FileRes.Responses[0].Body.FileID,
		Size:     0,
		IsFolder: true,
	}
	linkArgs := model.LinkArgs{
		IP:     "your_ip_address",
		Header: http.Header{"Authorization": []string{"Bearer your_token"}},
		Type:   "application/json",
		HttpReq: &http.Request{
			Method: "GET",
			// Initialize other fields of the http.Request if needed
		},
	}
	link, err1 := storageAilYunOpen.Link(ctx, rootObj, linkArgs)
	return link, err1
}

func CopyFile(file model.Obj, d *AliyundriveShare) (CopyFileRes, *model.Link) {
	var FileRes CopyFileRes
	fmt.Println("https://api.aliyundrive.com/adrive/v2/batch  start:", file.GetID()+d.UserDriveId)
	_, err := d.request("https://api.aliyundrive.com/adrive/v2/batch", http.MethodPost, func(req *resty.Request) {
		req.SetHeader(CanaryHeaderKey, CanaryHeaderValue).
			SetBody(base.Json{
				"requests": []base.Json{
					{
						"headers": base.Json{
							"Content-Type": "application/json",
						},
						"method": "POST",
						"id":     0,
						"body": base.Json{
							"file_id":           file.GetID(),
							"share_id":          d.ShareId,
							"to_drive_id":       d.UserDriveId,
							"to_parent_file_id": "root",
						},
						"url": "/file/copy",
					},
				},
				"resource": "file",
			}).SetResult(&FileRes)
	})
	if err != nil {
		return CopyFileRes{}, nil
	}
	return FileRes, nil
}

func DeleteFile(d *AliyundriveShare, FileRes CopyFileRes) {
	_, _ = d.request("https://api.aliyundrive.com/adrive/v2/batch", http.MethodPost, func(req *resty.Request) {
		req.SetHeader(CanaryHeaderKey, CanaryHeaderValue).
			SetBody(base.Json{
				"requests": []base.Json{
					{
						"headers": base.Json{
							"Content-Type": "application/json",
						},
						"method": "POST",
						"id":     FileRes.Responses[0].Body.FileID,
						"body": base.Json{
							"file_id":  FileRes.Responses[0].Body.FileID,
							"drive_id": d.UserDriveId,
						},
						"url": "/file/delete",
					},
				},
				"resource": "file",
			}).SetResult(&FileRes)
	})
}

func (d *AliyundriveShare) Other(ctx context.Context, args model.OtherArgs) (interface{}, error) {
	var resp base.Json
	var url string
	data := base.Json{
		"share_id": d.ShareId,
		"file_id":  args.Obj.GetID(),
	}
	switch args.Method {
	case "doc_preview":
		url = "https://api.aliyundrive.com/v2/file/get_office_preview_url"
	case "video_preview":
		url = "https://api.aliyundrive.com/v2/file/get_video_preview_play_info"
		data["category"] = "live_transcoding"
	default:
		return nil, errs.NotSupport
	}
	_, err := d.request(url, http.MethodPost, func(req *resty.Request) {
		req.SetBody(data).SetResult(&resp)
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

var _ driver.Driver = (*AliyundriveShare)(nil)
