package main

import (
	"fmt"

	"github.com/caistrong/gopl-homework/src/chapter4/work4_11/editorin"
)

// run command
// go run ./src/chapter4/work4_11/work4_11.go

// skip 4-11
// 我觉得这道题是个非常好的题目，基本上涵盖了一个常见go web client要做的事情(CRUD)，另外还包含Auth以及系统调用(调用文本编辑器获取输入)
// 问题是要把这道题写完太耗时间了（我懒），其中还包含了一些重复的发起http请求，处理http响应的工作，比较耗时，因此打算暂时跳过该题，后续有时间再来补上
func main() {
	sensitiveBytes, err := editorin.CaptureInputFromEditor(editorin.GetPreferredEditorFromEnvironment)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	fmt.Printf("%v\n", string(sensitiveBytes[:]))
}
