package reward

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lucky-draw/internal/lucky-drawn/wechat"
	"math/rand"
	"time"
)

var jsonUsersFile = ""
func init()  {
	//Get original users info from wechat, then save it into "../wechat/users_original.txt"
	filename := wechat.GetAllData();
	jsonUsersFile = wechat.FormatUsersToJson(filename)
}

func GetAllUsers() wechat.Users {
	usersByte, err := ioutil.ReadFile(jsonUsersFile)
	users := wechat.Users{}
	if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal(usersByte, &users)
		if err != nil {
			panic(err)
		}
	}
	return users
}

func GetAwardUser(users wechat.Users)(user wechat.User)  {
	userList := users.UserList
	size := len(userList)
	rand.Seed(time.Now().Unix())
	awardIndex := rand.Intn(size)
	fmt.Println(fmt.Sprintf("awardIndex: %d" ,awardIndex))
	user = userList[awardIndex]
	return user
}

func getAwardUser_weight(users map[string]int64) (name string) {
	type awardUser struct {
		name   string
		offset int64
		count  int64
	}

	userSli := make([]*awardUser, 0,len(users))
	var sumCount int64 = 0
	for n, c := range users {
		a := awardUser{
			name:   n,
			offset: sumCount,
			count:  c,
		}
		//整理所有用户的count数据为数轴
		userSli = append(userSli, &a)
		sumCount += c
	}

	awardIndex := rand.Int63n(sumCount)
	for _, u := range userSli {
		//判断获奖index落在那个用户区间内
		if u.offset+u.count>awardIndex {
			name = u.name
			return
		}
	}
	return
}
