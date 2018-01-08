package main

type Balance interface{
	 DoBalance([]string) string //传递一个ip列表，返回一个ip地址
} 
