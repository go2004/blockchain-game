//package Test

//import "reflect"

package main

import (
	"fmt"
	_ "reflect"
	"time"
	"math/rand"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Skills interface {
	reading()
	running()
}

type Student struct {
	Name string
	Age  int
}

func (self Student) runing() {
	fmt.Printf("%s is running\n", self.Name)
}
func (self Student) reading() {
	fmt.Printf("%s is reading\n", self.Name)
}

//生成随机字符串
func GetRandomString(strlen int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < strlen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func main() {
	/*	stu1 := Student{Name:"wd",Age:22}
		inf := new(Skills)
		stu_type := reflect.TypeOf(stu1)
		inf_type := reflect.TypeOf(inf).Elem()   // 特别说明，引用类型需要用Elem()获取指针所指的对象类型
		fmt.Println(stu_type.String())  //main.Student
		fmt.Println(stu_type.Name()) //Student
		fmt.Println(stu_type.PkgPath()) //main
		fmt.Println(stu_type.Kind()) //struct
		fmt.Println(stu_type.Size())  //24
		fmt.Println(inf_type.NumMethod())  //2
		fmt.Println(inf_type.Method(0),inf_type.Method(0).Name)  // {reading main func() <invalid Value> 0} reading
		fmt.Println(inf_type.MethodByName("reading")) //{reading main func() <invalid Value> 0} true
	*/
	/*
		str := "abc123"

		//方法一
		data := []byte(str)
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

		fmt.Println(md5str1)
		fmt.Println(GetRandomString(10))
	*/

	/*
		data := make(map[int] string)
		data[1]= "a"
		data[2]= "b"
		data[3]= "c"
		data[4]= "d"


		if _, ok := data[4]; ok {
			fmt.Println("r=",data[4])
		}

		*/
/*
	type blockMachine struct {
		Index     int
		Timestamp string
		PlayerID  uint
		DataID    int
		Price     int //单价
		Count     int //数量
		Hash      string
		PrevHash  string
	}


	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/blockchain?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("connect db error")
	}
	defer db.Close()
*/
/*
	//读矿机连数据
	var list []blockMachine

	db.Table("block_machines").Select("index,timestamp,player_id,data_id,price,count,hash,prev_hash").Scan(&list)
	println(list)
	*/
	//fmt.Sprintf("a %s", "string")

	//矿机数据 （列表）
	type worker struct {
		//位置 0	矿机名称
		//位置 1	矿机算力
		//位置 2	过去1小时算力
		//位置 3	过去1小时过期拒绝数
		//位置 4	过去24小时算力
		//位置 5	过去24小时过期拒绝数
		//位置 6	最近提交Share时间
		//位置 7	扩展字段
	}

	//历史收益 （字典）
	type value_workers struct {

		//矿机 1    历史收益
		//矿机 2    历史收益
	}
	type  Hashratehistory struct {
		Time string
		value float32
	}
	type Message struct {
		//Balance	float32	//	余额
		//Paid float32	//已支付工资
		//Payout_history	float32	//支付记录
		//Value	float32	//总收益
		Value_last_day	float32	//过去24小时收益
		//Stale_hashes_rejected_last_day	float32	//过去24小时过期拒绝数
		//Hashes_last_day	float32	//过去24小时算力
		//Hashrate	float32	//实时算力
		Hashrate_history  string	//	历史算力（过去24小时）
		//Worker_length	float32	//矿机总数
		//Worker_length_online	float32	//在线矿机数
		//Workers		float32	//矿机数据 （列表）
		//Value_workers	float32	//历史收益 （字典）
	}

	resp, err := http.Get("http://api.f2pool.com/bitcoin/user")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	/*message := Message{}
	err = json.Unmarshal(body,&message)
	fmt.Println(string(body))
	fmt.Println(message.Hashrate_history)
	*/
	/*
	var f interface{}
	json.Unmarshal(body, &f)

	m := f.(map[string]interface{})
	fmt.Println(m["hashrate_history"])  // 读取 json 内容
*/
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err == nil {
		fmt.Println(dat["hashrate_history"])
	}
}
