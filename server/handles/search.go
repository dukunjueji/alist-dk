package handles

import (
	"fmt"
	json "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
	"github.com/alist-org/alist/v3/internal/search"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/alist-org/alist/v3/server/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SearchReq struct {
	model.SearchReq
	Password string `json:"password"`
}

type SearchResp struct {
	model.SearchNode
	Type int `json:"type"`
}

func Search(c *gin.Context) {
	var (
		req SearchReq
		err error
	)
	if err = c.ShouldBind(&req); err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	user := c.MustGet("user").(*model.User)
	req.Parent, err = user.JoinPath(req.Parent)
	if err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	if err := req.Validate(); err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	nodes, total, err := search.Search(c, req.SearchReq)
	if err != nil {
		common.ErrorResp(c, err, 500)
		return
	}
	var filteredNodes []model.SearchNode
	for _, node := range nodes {
		if !strings.HasPrefix(node.Parent, user.BasePath) {
			continue
		}
		meta, err := op.GetNearestMeta(node.Parent)
		if err != nil && !errors.Is(errors.Cause(err), errs.MetaNotFound) {
			continue
		}
		if !common.CanAccess(user, meta, path.Join(node.Parent, node.Name), req.Password) {
			continue
		}
		filteredNodes = append(filteredNodes, node)
	}
	common.SuccessResp(c, common.PageResp{
		Content: utils.MustSliceConvert(filteredNodes, nodeToSearchResp),
		Total:   total,
	})
}

func nodeToSearchResp(node model.SearchNode) SearchResp {
	return SearchResp{
		SearchNode: node,
		Type:       utils.GetObjType(node.Name, node.IsDir),
	}
}
func HotSearch(c *gin.Context) {

	client := &http.Client{}
	keyword := c.Query("cat")
	//fmt.Println("创建请求时keyword:", keyword)
	req, err := http.NewRequest("GET", "https://api.web.360kan.com/v1/rank?cat="+keyword, nil)
	if err != nil {
		fmt.Println("创建请求时出错:", err)
		return
	}
	// 添加自定义的请求头
	req.Header.Add("Referer", "https://www.360kan.com/rank/general")

	// 发送 HTTP 请求
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// 解析JSON数据
	var tvShow TVShow
	if err := json.Unmarshal(body, &tvShow); err != nil {
		fmt.Println("解析JSON时出错:", err)
		return
	}

	common.SuccessResp(c, tvShow)
}

type TVShow struct {
	Data []struct {
		Title       string   `json:"title"`
		Comment     string   `json:"comment"`
		UpInfo      string   `json:"upinfo"`
		DoubanScore string   `json:"doubanscore"`
		ID          int      `json:"id"`
		Cat         int      `json:"cat"`
		PV          string   `json:"pv"`
		Cover       string   `json:"cover"`
		URL         string   `json:"url"`
		Percent     string   `json:"percent"`
		EntID       string   `json:"ent_id"`
		MovieCat    []string `json:"moviecat"`
		VIP         bool     `json:"vip"`
		Description string   `json:"description"`
		PubDate     string   `json:"pubdate"`
	} `json:"data"`
}
