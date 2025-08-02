package main

import "fmt"

func main() {
	errCh := make(chan error)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				errCh <- fmt.Errorf("子协程出错：%v", e)
			}
		}()
		// 子协程里故意触发panic
		panic("出错了")
	}()
	// 父协程通过channel接错误
	if err := <-errCh; err != nil {
		fmt.Println("抓到子协程错误：", err)
	}
}
