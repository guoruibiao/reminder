package main

import "github.com/garyburd/redigo/redis"

type Event struct {
	Type string
	//Status int // 想了想，鸡肋功能，更新会导致Redis的sorted_set的member变化
	Title string // 简要描述
	Description string // 详细内容
	Addtime int64 // 添加时间
	Tiptime int64 // 提醒时间
	Master string // 所属人，会应用到Redis的key前缀
}

type RedisConfigs struct {
	address string
	network string
	options redis.DialOption
}
type ServerConfigs struct {
	addr string
}