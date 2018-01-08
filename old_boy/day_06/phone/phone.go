package main

import "fmt"

type Phone struct {
	PayMap map [string]Pay
}

func (p *Phone) OpenWeiChatPay(){
	weChatPay := &WeChatPay{}
	p.PayMap["wechatPay"] = weChatPay

}

func (p *Phone)OpenALiPay(){
	aLiPay := &ALiPay{}
	p.PayMap["aLiPay"] = aLiPay

}

func (p *Phone)PayMoney(name string, money float32) (err error)  {
	/*
	name:支付的方式 ，支付宝/微信
	money:支付金额
	 */
	 pay, ok := p.PayMap[name]
	if !ok {
		err = fmt.Errorf("不支持[%s]支付方式", name)
		return
		}
	err = pay.pay(123,money)
	return

	}