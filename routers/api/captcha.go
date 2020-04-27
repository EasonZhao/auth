package api

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/securityin/auth/pkg/app"
)

const (
	// CaptchaLen 验证码长度
	CaptchaLen = 6
)

// GetCaptcha 获取验证码
// @Summary 获取验证码
// @accept application/x-www-form-urlencoded
// @Tags captcha
// @Produce json
// @Success 200 {body} string "{"ID":"验证码ID","CaptchaURL":"访问路径"}"
// @Failure 204 {body} string "{"code":204,"data":null,"msg":"错误信息"}"
// @Router /captcha [post]
func GetCaptcha(c *gin.Context) {
	appG := app.GetGin(c)
	captchaID := captcha.NewLen(CaptchaLen)

	data := &struct {
		ID         string
		CaptchaURL string
	}{
		captchaID,
		"http://" + c.Request.Host + "/captcha/" + captchaID + ".png",
	}
	appG.ResponseSuc(data)
}

// GetCaptchaImage 获取验证码图片
// @Summary 获取验证码图片
// @accept application/x-www-form-urlencoded
// @Tags captcha
// @Produce mage/png
// @Success 200 {body} string "图片数据"
// @Router /captcha/:captchaId [post]
func GetCaptchaImage(c *gin.Context) {
	//captchaID := c.Param("captchaId")
	dir, file := path.Split(c.Request.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//fmt.Println("reload : " + r.FormValue("reload"))
	if c.Request.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(c.Request.FormValue("lang"))
	download := path.Base(dir) == "download"

	if serve(c.Writer, c.Request, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(c.Writer, c.Request)
	}
}

func serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
