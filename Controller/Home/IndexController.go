package Home

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"webchat/Controller/Base"
)

type IndexController struct {

	this Base.BaseController
}

/**
	当前登录用户列表
 */
type userList struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data userData `json:"data"`

}

/**
	用户信息
 */
type userData struct {
	Mine userMine 	`json:"mine"`
	Friend []UserFried `json:"friend"`
	Group []GroupList `json:"group"`
}

/**
	登录用户详情
 */
type userMine struct {
	Username string `json:"username"`
	ID int `json:"id"`
	Status string `json:"status"`
	Sign string	`json:"sign"`
	Avatar string `json:"avatar"` //图片地址
}

/**
	所属朋友
 */
type UserFried struct {
	Groupname string  `json:"groupname"`
	ID int `json:"id"`
	Online int `json:"online"`
	List []FriedList `json:"list"`
}

/**
	朋友详情
 */
type FriedList struct {
	Username string `json:"username"`
	ID int `json:"id"`
	Avatar string `json:"avatar"`
	Sign string `json:"sign"`
}

/**
	分组列表
 */
type GroupList struct {
	Groupname string `json:"groupname"`
	ID int `json:"id"`
	Avatar string `json:"avatar"`
}



func (this *IndexController) chatIndex()  {

	fmt.Print("哈哈哈")
}

func ChatIndex(c *gin.Context)  {

	fmt.Print("我被请求到了")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func UseList(c *gin.Context)  {

	var data  userList //信息

	//临时模拟
	var mine userMine

	mine.Username = "纸飞机"
	mine.ID = 1000
	mine.Status = "online"
	mine.Sign = "在深邃的编码世界，做一枚轻盈的纸飞机"
	mine.Avatar = "http://cdn.firstlinkapp.com/upload/2016_6/1465575923433_33812.jpg"

	var flist []FriedList

	flist = make([]FriedList,1)

	flist[0].ID = 100001
	flist[0].Username = "贤心"
	flist[0].Avatar = "http://tp1.sinaimg.cn/1571889140/180/40030060651/1"
	flist[0].Sign  = "这些都是测试数据，实际使用请严格按照该格式返回"


	var grouplist []GroupList

	grouplist = make([]GroupList,2)
	grouplist[0].Avatar = "前端群"
	grouplist[0].ID = 101
	grouplist[0].Avatar = "http://tp2.sinaimg.cn/2211874245/180/40050524279/0"
	grouplist[1].Avatar = "Fly社区官方群"
	grouplist[1].ID = 102
	grouplist[1].Avatar = "http://tp2.sinaimg.cn/5488749285/50/5719808192/1"

	var friend []UserFried

	friend = make([]UserFried,1)
	friend[0].ID = 2000
	friend[0].Groupname = "前端码屌"
	friend[0].Online = 2
	friend[0].List = flist

	var userInfo userData

	userInfo.Group = grouplist
	userInfo.Friend = friend

	userInfo.Mine = mine

	data.Code = 0
	data.Data = userInfo
	data.Msg = "获取成功"

	a,_ := json.Marshal(data)

	fmt.Print(string(a))

	c.JSON(200,data)
}

func UserMembers(c *gin.Context)  {

	var data string = `{
  "code": 0
  ,"msg": ""
  ,"data": {
    "owner": {
      "username": "贤心"
      ,"id": "100001"
      ,"avatar": "http://tp1.sinaimg.cn/1571889140/180/40030060651/1"
      ,"sign": "这些都是测试数据，实际使用请严格按照该格式返回"
    }
    ,"members": 12
    ,"list": [{
      "username": "贤心"
      ,"id": "100001"
      ,"avatar": "http://tp1.sinaimg.cn/1571889140/180/40030060651/1"
      ,"sign": "这些都是测试数据，实际使用请严格按照该格式返回"
    },{
      "username": "Z_子晴"
      ,"id": "108101"
      ,"avatar": "http://tva3.sinaimg.cn/crop.0.0.512.512.180/8693225ajw8f2rt20ptykj20e80e8weu.jpg"
      ,"sign": "微电商达人"
    },{
      "username": "Lemon_CC"
      ,"id": "102101"
      ,"avatar": "http://tp2.sinaimg.cn/1833062053/180/5643591594/0"
      ,"sign": ""
    },{
      "username": "马小云"
      ,"id": "168168"
      ,"avatar": "http://tp4.sinaimg.cn/2145291155/180/5601307179/1"
      ,"sign": "让天下没有难写的代码"
    },{
      "username": "徐小峥"
      ,"id": "666666"
      ,"avatar": "http://tp2.sinaimg.cn/1783286485/180/5677568891/1"
      ,"sign": "代码在囧途，也要写到底"
    },{
      "username": "罗玉凤"
      ,"id": "121286"
      ,"avatar": "http://tp1.sinaimg.cn/1241679004/180/5743814375/0"
      ,"sign": "在自己实力不济的时候，不要去相信什么媒体和记者。他们不是善良的人，有时候候他们的采访对当事人而言就是陷阱"
    },{
      "username": "长泽梓Azusa"
      ,"id": "100001222"
      ,"avatar": "http://tva1.sinaimg.cn/crop.0.0.180.180.180/86b15b6cjw1e8qgp5bmzyj2050050aa8.jpg"
      ,"sign": "我是日本女艺人长泽あずさ"
    },{
      "username": "大鱼_MsYuyu"
      ,"id": "12123454"
      ,"avatar": "http://tp1.sinaimg.cn/5286730964/50/5745125631/0"
      ,"sign": "我瘋了！這也太準了吧  超級笑點低"
    },{
      "username": "谢楠"
      ,"id": "10034001"
      ,"avatar": "http://tp4.sinaimg.cn/1665074831/180/5617130952/0"
      ,"sign": ""
    },{
      "username": "柏雪近在它香"
      ,"id": "3435343"
      ,"avatar": "http://tp2.sinaimg.cn/2518326245/180/5636099025/0"
      ,"sign": ""
    },{
      "username": "林心如"
      ,"id": "76543"
      ,"avatar": "http://tp3.sinaimg.cn/1223762662/180/5741707953/0"
      ,"sign": "我爱贤心"
    },{
      "username": "佟丽娅"
      ,"id": "4803920"
      ,"avatar": "http://tp4.sinaimg.cn/1345566427/180/5730976522/0"
      ,"sign": "我也爱贤心吖吖啊"
    }]
  }
}`

	c.JSON(200,data)
}

