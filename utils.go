package main

import (
	"encoding/json"
	"fmt"
)

func Decode(content string)(Event, error) {
	bucket := &Event{}
	json.Unmarshal([]byte(content), bucket)
	//fmt.Println(bucket)
	return *bucket, nil
}

func Encode(event Event) string {
	bytes, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

/**
  短时记忆	5分钟，20分钟， 1小时
  长时记忆	12小时，1天，3天，7天，14天，30天
 */
func generateEvents(event Event) []Event {
	slots := []int64{
		// 短时记忆槽
		5 * 60,
		20 * 60,
		60 * 60,
		// 长时记忆槽
		12 * 60 * 60,
		86400,
		86400*3,
		86400*7,
		86400*14,
		86400*30,
		86400*90,
	}
	var events []Event
	for _, addseconds := range slots {
		tmp := event
		tmp.Tiptime = tmp.Addtime + addseconds
		events = append(events, tmp)
	}
	return events
}