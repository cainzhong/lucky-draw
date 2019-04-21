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
	TOKEN            = "2129137112"
	GET_ALL_DATA_URL = "https://mp.weixin.qq.com/cgi-bin/user_tag?action=get_all_data&lang=zh_CN&token=" + TOKEN
	WECHAT_COOKIE    = "noticeLoginFlag=1; remember_acct=cainzhong%40aliyun.com; ua_id=SBLR0Bo1PAKB7LthAAAAACVBrSF6-gma9RkWePKj16I=; mm_lang=zh_CN; pgv_pvi=3183542272; pgv_si=s9203781632; cert=5YYgvG489OLMRk3C50yFUumgV8weCAZ3; rewardsn=; wxtokenkey=777; mmad_session=b2c021e92ad8928c30e99f22fc2931b99556fea4c957af084e14949a9c5f666dc38df13e3c51088b8432af8bb5a07892d65e3a7be638fea663d38d562076244e6c16752d877df7b29c235bbdfdac2649f6c51559a3f5e6247385f58268ccd99111de1c56c245721266e7088080fefde3; pgv_info=ssid=s4121212104; pgv_pvid=9472565614; ts_uid=5186855000; openid2ticket_oh4N71BLyWpAhNUUbAcJGMQq-D4k=ujTNPrYtl0YoTwydJlibeNz67Hg58Hpg+fAGrU8nplI=; uuid=5ed22cfb71f6158d817954a3df106729; data_bizuin=3564864391; bizuin=3564864391; data_ticket=lH2MbwxSuI+99LqEgCyccxizmv/0kM1+gf9o40q5yciX/M92WEPuOSFLhhsTWSc8; slave_sid=RU5LMTRab3JZVXJralpEMHRqWWxxbkdkU0lqYlYyVno5Nnd2XzhOVFhOMnpXcnd1S3pXZVp4Ym5XeDh1ZkZQRkNzVUVxWnZEVDZsTEtQVV9RUEFlWktvV0l0OVQ4cm5OQ3Z6dXViRGcyMng3dXRhc3c1djZBenhuQ00wUWZxWFpma2dNZ1k2WnpMdlJLZFhW; slave_user=gh_bb2dbc84136b; xid=0c74ec355461d870326f7bf30b2608aa"
	BASE_DIR         = "C:/repositories/cainzhong/src/lucky-draw/"
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
	filename := path.Join(path.Dir(BASE_DIR), "/assets/users_original.txt")
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

	jsonFilename := path.Join(path.Dir(BASE_DIR), "/assets/users.json")
	logger.Info("Save the formated JSON into file %s", jsonFilename)
	ioutil.WriteFile(jsonFilename, []byte(jsonUsers), 0644)
	return jsonFilename
}

func GetHeadImg(fakeid string) []byte {
	logger.Info("Trying to get head image from wechat...")
	headers := GenerateWechatRequestHeaders()
	url := "https://mp.weixin.qq.com/misc/getheadimg?fakeid=" + fakeid + "&token=" + TOKEN + "&lang=zh_CN"
	resp, err := req.Get(url, headers)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Response().Body.Close()

	b, _ := ioutil.ReadAll(resp.Response().Body)
	logger.Info("Got image size: %d", len(b))
	return b
}
