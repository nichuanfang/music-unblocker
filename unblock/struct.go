package unblock

import "fmt"

type Config struct {
	// 字符
	str string
	// 年龄
	age int
}

func (receiver *Config) name() {
	receiver.str = "12"
	fmt.Println(receiver.str)
}

func AppMain() {
	c := &Config{
		str: "ds",
		age: 12,
	}
	print(c)
}
