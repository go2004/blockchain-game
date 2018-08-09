package main

import (
	"github.com/opesun/goquery"
	"fmt"
	"strings"
	"strconv"
	"github.com/name5566/leaf/timer"
	"time"
	"sort"
)


//%%游戏币排名
type RankInfo struct {
	//角色信息
	playerID	uint	//玩家ID
	name        string	//玩家名字
	style       uint32	//外形头像
	value		float32	//数据
	noID        uint32	 //排名
}

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (s Persons) Len() int           { return len(s) }
func (s Persons) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Persons) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	/*conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	//data := []byte(`{
	//	"Hello": {
	//		"Name": "炒在工"
	//	}
	//}`)

	test := &msg.U2LS_Login{
		SdkAccount: *proto.String("SdkAccount"),
		SdkClientId:*proto.String("SdkClientId"),
		SdkAccessToken:*proto.String("SdkAccessToken"),
	}

	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	// len + data
	len := 4+ len(data)
	m := make([]byte, len)

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len)+2)
	binary.BigEndian.PutUint16(m[2:], 0)

	copy(m[4:], data)

	// 发送消息
	conn.Write(m)

	m1 := make([] byte, 1024)
	len1, err := conn.Read(m1)
	fmt.Print(len1,string(m1))
	*/
	//chain.Init()
	//chain.WriteBlockMachine(uint(1001245), 12)

	/*
	//map 序列化为json
	type value_workers struct {

		name string
		dic map[string]string
	}

	var p2 *value_workers = &value_workers{name:"abcd"}

	var m map[string]string = make(map[string]string)
	m["Go"] = "No.1"
	m["Java"] = "No.2"
	m["C"] = "No.3"
	p2.dic = m
	if bs, err := json.Marshal(p2); err != nil {
		panic(err)
	} else {
		//result --> {"C":"No.3","Go":"No.1","Java":"No.2"}
		fmt.Println(string(bs))
	}
*/

	/*
		jsonStr := `{
		"hashes_last_hour": 0,
		"worker_length": 0,
		"balance": 0.000062923771232195,
		"stale_hashes_rejected_last_day": 0,
		"paid": 0,
		"hashrate": 0,
		"last": {},
		"hashrate_history": {
			"2018-07-30T17:30:00Z": 0,
			"2018-07-30T19:30:00Z": 0,
			"2018-07-31T02:30:00Z": 0
		},
		"hashes_last_day": 0,
		"payout_history": [],
		"stale_hashes_rejected_last_hour": 0,
		"value": 0.000062923771232195,
		"value_last_day": 0,
		"workers": [],
		"worker_length_online": 0
	}`


		//json str 转map
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
			fmt.Println("==============json str 转map=======================")
			//fmt.Println(dat)
			fmt.Println(dat["balance"])
		}
	*/


	var url = "https://www.antpool.com/earnComparison.htm"
	p, err := goquery.ParseUrl(url)
	if err != nil {
		panic(err)
	} else {
		t := p.Find("span").Text()
		startpos :=strings.Index(t, "BTC:0.")
		endpos :=strings.Index(t, "BCH:0.")

		btcStr := t[startpos+4:endpos]
		vtc, _ := strconv.ParseFloat(btcStr, 32)
		fmt.Println(vtc)
		//for i := 0; i < t.Length(); i++ {
		//	d := t.Eq(i).Attr("href")
		//	fmt.Println(d)
		//}
	}

	var a uint
	a= 1
	println("a=", a << 28)

/*
	v := 3.1415926535
	s1 := strconv.FormatFloat(v, 'f', -1, 64)
	println("s1:",s1)

	aChan := make(chan int, 1)
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("ticked at %v\n", time.Now())
			}
		}

	}()
	//阻塞主线程
	<-aChan

*/



	p1 := []Person{Person{"Lily", 20}, Person{"Bing", 18}, Person{"Tom", 23}, Person{"Vivy", 16}, Person{"John", 18}}

	sort.Sort(sort.Reverse(Persons(p1))) //sort.Reverse 生成递减序列
	fmt.Println(p)



}





func ExampleCronExpr() {
	cronExpr, err := timer.NewCronExpr("0 * * * *")
	if err != nil {
		return
	}

	fmt.Println(cronExpr.Next(time.Date(
		2000, 1, 1,
		20, 10, 5,
		0, time.UTC,
	)))

	// Output:
	// 2000-01-01 21:00:00 +0000 UTC
}