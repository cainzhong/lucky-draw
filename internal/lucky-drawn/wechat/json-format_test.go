package wechat

import (
	"fmt"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	users := `{
        lang: 'zh_CN',
        group_list : [
                                                                                {
                        name : "未分组",
                        cnt : "29" * 1,
                        create_time : "1552729425" * 1,
                        id : "0"
                    }
                                                        ,                    {
                        name : "星标组",
                        cnt : "0" * 1,
                        create_time : "1552729425" * 1,
                        id : "2"
                    }
                                                        ,                    {
                        name : "屏蔽组",
                        cnt : "0" * 1,
                        create_time : "" * 1,
                        id : "1"
                    }
                                                        ],
        user_list : [
                                                                                {
                        id : "oh4N71JpRL5pgCq6m1WmR0MWPf-k",
                        nick_name : "陈付梅",
                        remark_name : "",
                        create_time : "1554955701",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71BWXgbrgrkw3DoLGaTuqhuo",
                        nick_name : "Jun",
                        remark_name : "",
                        create_time : "1554686826",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71LHvoiMc1DhxTE0dMpb0VFQ",
                        nick_name : "盛夏",
                        remark_name : "",
                        create_time : "1554643594",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71EJodcN_MaeKMWi9R4w2K4E",
                        nick_name : "什么都想吃小朋友",
                        remark_name : "",
                        create_time : "1554635707",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71IKJ8zTSFzMTEwGETNY8VfA",
                        nick_name : "骏杰",
                        remark_name : "",
                        create_time : "1554634029",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71E-Ut5yJbevttpqHT8goLCc",
                        nick_name : "小方方方要努力",
                        remark_name : "",
                        create_time : "1554633863",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71NTKVm2DiGfzmfkjDOmfOJc",
                        nick_name : "赵丽英",
                        remark_name : "",
                        create_time : "1554633837",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71IXrnORZBrkWjUsq4vzmW0I",
                        nick_name : "SUSIE🌱",
                        remark_name : "",
                        create_time : "1554633816",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71F2bqp2gKtc2NZ1tVnd2oP4",
                        nick_name : "L.K",
                        remark_name : "",
                        create_time : "1554512623",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71OPTBoCdUOHzP4mM2oPLyZ4",
                        nick_name : "小浦",
                        remark_name : "",
                        create_time : "1554475847",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71IXmHBIbSYlhZlgQEc8-7KM",
                        nick_name : "宽字一",
                        remark_name : "",
                        create_time : "1554174461",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71FCgasBOi9n7dM5wgFYNCEw",
                        nick_name : "大白羊",
                        remark_name : "",
                        create_time : "1554107944",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71LzRIIYbr9tYuDlr3y12JaA",
                        nick_name : "Camille🤡",
                        remark_name : "",
                        create_time : "1554107913",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71Ba1B2-MjzjyGdrsRLvjgQM",
                        nick_name : "意中人😸",
                        remark_name : "",
                        create_time : "1554102239",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71M_4OYWC3IPKpOQHidKzs1M",
                        nick_name : "wilson",
                        remark_name : "",
                        create_time : "1554019070",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71H__hgt-G0wagHHvLu1rAlQ",
                        nick_name : "王赢～Seiya 🌕🐳",
                        remark_name : "",
                        create_time : "1554009687",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71LcUMAdia1ch6njvT1jd-Wk",
                        nick_name : "博文",
                        remark_name : "",
                        create_time : "1554009382",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71MUcoVIHopOjjnD2Z7jTB7M",
                        nick_name : "袁笑笑",
                        remark_name : "",
                        create_time : "1554009294",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71I7G0zKvZJT3sFhkpF30zsY",
                        nick_name : "🌈好运smile咘婷🍀",
                        remark_name : "",
                        create_time : "1553990248",
                        group_id : [
                                                                                ]
                    }
                                                        ,                    {
                        id : "oh4N71DRgWfbFjm1HU66_CQzin5s",
                        nick_name : "Angelina",
                        remark_name : "",
                        create_time : "1553954599",
                        group_id : [
                                                                                ]
                    }
                                                        ],
        total_user_num : "29"
    }`
	fmt.Println(JsonDecode(users))
}
