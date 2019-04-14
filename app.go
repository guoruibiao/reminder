package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
		"net/http"
	"github.com/gin-gonic/gin"
)

const KEY_PREFIX = "zset:"

func addEvent(event Event, redisConfigs RedisConfigs) (bool, error){
	client, err := redis.Dial(redisConfigs.network, redisConfigs.address)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	eventstr := Encode(event)
	score := event.Addtime
	if score == 0 {
		event.Addtime = time.Now().Unix()
	}
	key := KEY_PREFIX + event.Master
	events := generateEvents(event)
	for _, event := range events {
		eventstr = Encode(event)
		_, err = client.Do("zadd", key, event.Tiptime, eventstr)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func getEvents(redisConfigs RedisConfigs, master string, starttime string, endtime string)  ([]Event, error) {
	client, err := redis.Dial(redisConfigs.network, redisConfigs.address)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	key := KEY_PREFIX + master
	content, err := redis.Values(client.Do("zrangebyscore", key, starttime, endtime))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var ret []Event
	for _, v := range content {
		//fmt.Printf("%s ", v.([]byte))
		tmp, _ := Decode(string(v.([]byte)))
		ret = append(ret, tmp)
	}
	return ret, nil
}


func main() {
	app := gin.Default()
	//加载模板
	app.Static("/static", "/Users/biao/go/src/memory/static")


	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	//定义路由
	app.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 添加备忘事件
	app.GET("/addevent", func(context *gin.Context) {
		title := context.DefaultQuery("title", "no title")
		description := context.DefaultQuery("description", "no description")
		master := "biao"
		event := Event{
			Type:"test",
			Title:title,
			Description:description,
			Addtime:time.Now().Unix(),
			Tiptime:time.Now().Unix() + 10000,
			Master:master,
		}
		isSuccess, err := addEvent(event, redisConfigs)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err})
		}else{
			context.JSON(http.StatusOK, gin.H{"ret": isSuccess})
		}
	})
	// 获取备忘事件
	app.GET("/getevent", func(context *gin.Context) {
		starttime := context.DefaultQuery("starttime", "0")
		endtime := context.DefaultQuery("endtime", string(int(time.Now().Unix())))
		fmt.Println(starttime, endtime)
		master := "biao"
		events, err := getEvents(redisConfigs, master, starttime, endtime)
		fmt.Println(events)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err, "starttime":starttime, "endtime":endtime})
		}else{
			context.JSON(http.StatusOK, gin.H{"ret": events, "starttime":starttime, "endtime":endtime})
		}
	})
	app.Run(serverConfigs.addr)
}
