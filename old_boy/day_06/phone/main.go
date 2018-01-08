package main

import "fmt"

func main() {

	phone := &Phone{
		PayMap:make(map[string]Pay,16),
	}
	phone.OpenWeiChatPay()

	err := phone.PayMoney("wechatPay", 10)
	if err != nil {
		fmt.Println("支付失败，失败原因：", err)
	}
	fmt.Println("支付成功")
}
