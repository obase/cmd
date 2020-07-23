package cmd

import (
	"flag"
	"fmt"
	"testing"
)

//func init()  {
//	os.Args = []string{os.Args[0], "test", "-help"}
//}

func TestExec(t *testing.T) {

	Add("test", func() int {
		flag.Bool("option", false, "这是一个测试选项")
		fmt.Println("init test")
		return 0
	}, func() int {
		fmt.Println("exec test")
		return 0
	}, "测试")
	Add("testtesttest", func() int {
		flag.Bool("option", false, "这是一个测试选项")
		fmt.Println("init test")
		return 0
	}, func() int {
		fmt.Println("exec test")
		return 0
	}, "测试")
	Exec("dataflow")
}
