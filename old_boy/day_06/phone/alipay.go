package main

import "fmt"

type ALiPay struct {

}

func (wcP *ALiPay)pay(userId int64, money float32) error  {
	fmt.Println("1, 连接到阿里支付的服务器")
	fmt.Println("2, 找到对应的用户")
	fmt.Println("3, 检查余额是否充足")
	fmt.Println("4, 扣款")
	fmt.Println("5, 返回扣款是否成功")
	return nil
}