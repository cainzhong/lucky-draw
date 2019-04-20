package wechat

import (
	"github.com/imroc/req"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"lucky-draw/internal/lucky-drawn"
	"path"
	"regexp"
	"strings"
)

const (
	//replace them
	GET_ALL_DATA_URL = "https://mp.weixin.qq.com/cgi-bin/user_tag?action=get_all_data&lang=zh_CN&token=1840900442"
	WECHAT_COOKIE    = "noticeLoginFlag=1; remember_acct=cainzhong%40aliyun.com; ua_id=SBLR0Bo1PAKB7LthAAAAACVBrSF6-gma9RkWePKj16I=; mm_lang=zh_CN; pgv_pvi=3183542272; pgv_si=s9203781632; ticket_id=gh_bb2dbc84136b; cert=5YYgvG489OLMRk3C50yFUumgV8weCAZ3; rewardsn=; wxtokenkey=777; bizuin=3564864391; mmad_session=b2c021e92ad8928c30e99f22fc2931b99556fea4c957af084e14949a9c5f666dc38df13e3c51088b8432af8bb5a07892d65e3a7be638fea663d38d562076244e6c16752d877df7b29c235bbdfdac2649f6c51559a3f5e6247385f58268ccd99111de1c56c245721266e7088080fefde3; pgv_info=ssid=s4121212104; pgv_pvid=9472565614; ts_uid=5186855000; uuid=2d7902529c3f2154c1980b18afc2b1b8; ticket=89ed7b1ecf4fa44aa22918dc75699641e8089b05; data_bizuin=3564864391; data_ticket=1D5YPwSX7jEJ5QCj2F0hqyKfl6TO4q8bzscG6Gc8q88Bg57wF4yRv6KKFyNC7/nj; slave_sid=MXhUVEVRc3lvVGlhdl9uZ2NTc0VQUk5fa3pTSnVUWXdibjlaTFVtcDdoN0NjQ1BjTUR2Qng5bWoxVEZaOEszbzNKNldwdkFSVUhQcWhzMHA2cFRJc3dYSFhYdnVRSTgzcHRJM0RWS1pZWWZCMWpVRk1ycVV1OTVOSkx4cGR4ZU5DZzcyRjFiZzVieTU0TWNW; slave_user=gh_bb2dbc84136b; xid=fa4eb1e5f433bdc1a0bcdde34b543242; openid2ticket_oh4N71BLyWpAhNUUbAcJGMQq-D4k=ujTNPrYtl0YoTwydJlibeNz67Hg58Hpg+fAGrU8nplI=noticeLoginFlag=1; remember_acct=cainzhong%40aliyun.com; ua_id=SBLR0Bo1PAKB7LthAAAAACVBrSF6-gma9RkWePKj16I=; mm_lang=zh_CN; pgv_pvi=3183542272; pgv_si=s9203781632; ticket_id=gh_bb2dbc84136b; cert=5YYgvG489OLMRk3C50yFUumgV8weCAZ3; rewardsn=; wxtokenkey=777; bizuin=3564864391; mmad_session=b2c021e92ad8928c30e99f22fc2931b99556fea4c957af084e14949a9c5f666dc38df13e3c51088b8432af8bb5a07892d65e3a7be638fea663d38d562076244e6c16752d877df7b29c235bbdfdac2649f6c51559a3f5e6247385f58268ccd99111de1c56c245721266e7088080fefde3; pgv_info=ssid=s4121212104; pgv_pvid=9472565614; ts_uid=5186855000; uuid=2d7902529c3f2154c1980b18afc2b1b8; ticket=89ed7b1ecf4fa44aa22918dc75699641e8089b05; data_bizuin=3564864391; data_ticket=1D5YPwSX7jEJ5QCj2F0hqyKfl6TO4q8bzscG6Gc8q88Bg57wF4yRv6KKFyNC7/nj; slave_sid=MXhUVEVRc3lvVGlhdl9uZ2NTc0VQUk5fa3pTSnVUWXdibjlaTFVtcDdoN0NjQ1BjTUR2Qng5bWoxVEZaOEszbzNKNldwdkFSVUhQcWhzMHA2cFRJc3dYSFhYdnVRSTgzcHRJM0RWS1pZWWZCMWpVRk1ycVV1OTVOSkx4cGR4ZU5DZzcyRjFiZzVieTU0TWNW; slave_user=gh_bb2dbc84136b; xid=fa4eb1e5f433bdc1a0bcdde34b543242; openid2ticket_oh4N71BLyWpAhNUUbAcJGMQq-D4k=ujTNPrYtl0YoTwydJlibeNz67Hg58Hpg+fAGrU8nplI="
	BASE_DIR = "C:/repositories/cainzhong/src/lucky-draw/"
)

func init() {
	const webchatCertFile = "/assets/wechat.cer"
	localCertFile := path.Join(path.Dir(BASE_DIR), webchatCertFile)
	client := lucky_drawn.InitSystemCertPool(localCertFile)
	req.SetClient(client)
}

func GenerateWechatRequestHeaders() req.Header {
	headers := make(map[string]string)
	headers["Cookie"] = WECHAT_COOKIE
	return headers
}

/**
Get original users info from wechat, then save it into a text file.
*/
func GetAllData() string {
	logger.Info("Trying to get all data from wechat...")
	headers := GenerateWechatRequestHeaders()
	resp, err := req.Get(GET_ALL_DATA_URL, headers)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Response().Body.Close()

	b, _ := ioutil.ReadAll(resp.Response().Body)
	text := string(b)
	//匹配wx.cgiData=开头，seajs.use结尾的字符串，需要再把开头和结尾去掉
	regExp := `(wx.cgiData=)([\w\W]*)(seajs.use)`
	reg := regexp.MustCompile(regExp)
	matchedText := string(reg.Find([]byte(text)))
	matchedText = strings.ReplaceAll(matchedText, "wx.cgiData=", "")
	matchedText = strings.ReplaceAll(matchedText, "seajs.use", "")
	matchedText = strings.ReplaceAll(matchedText, ";", "")
	logger.Info("Find matched text\n %s", matchedText)
	filename := path.Join(path.Dir(BASE_DIR), "./users_original.txt")
	logger.Info("Save it into file %s", filename)
	ioutil.WriteFile(filename, []byte(matchedText), 0644)
	return filename
}

func FormatUsersToJson(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error(err)
	}
	jsonUsers := JsonDecode(string(b))

	jsonFilename := path.Join(path.Dir(BASE_DIR), "./users.json")
	logger.Info("Save the formated JSON into file %s", jsonFilename)
	ioutil.WriteFile(jsonFilename, []byte(jsonUsers), 0644)
	return jsonFilename
}
