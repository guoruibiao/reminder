package main

import "fmt"

func testAddEvent() {
	redisConfigs := RedisConfigs{
		network:"tcp",
		address:"127.0.0.1:6379",
	}
	event := Event{
		Type: "a",
		Description: "测试",
		Master:"biao",
	}
	isSuccess, err := addEvent(event, redisConfigs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("操作结果：", isSuccess)
}

func testGetEvents() {
	redisConfigs := RedisConfigs{
		network:"tcp",
		address:"127.0.0.1:6379",
	}
	events, _ := getEvents(redisConfigs, "biao", "1555164515", "1555207604")
	for _, event := range events {
		fmt.Println(event)
	}
}
