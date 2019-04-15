package reward

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetAllUsers(t *testing.T) {
	users := GetAllUsers()
	fmt.Println(fmt.Sprintf("%+v",users))
}

func TestGetAwardUser(t *testing.T) {
	users := GetAllUsers()
	user:=GetAwardUser(users)
	fmt.Println(fmt.Sprintf("Congratulations! The award user is %+v",user))
}

func Test_getAwardUser_weight(t *testing.T) {
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 6,
		"c": 3,
		"d": 12,
		"f": 1,
	}

	rand.Seed(time.Now().Unix())
	awardCount := make(map[string]int)
	for i := 0; i <= 100000; i++ {
		awardName := getAwardUser_weight(users)
		if count, ok := awardCount[awardName]; ok {
			awardCount[awardName] = count + 1
		} else {
			awardCount[awardName] = 0
		}
	}
	for n,c := range awardCount {
		fmt.Printf("%v:%v \n",n,c)
	}
}
