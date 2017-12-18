package main
/*
使用按位与、按位或、异或 来定义类是属性，减少类的字段
设置为false时需要预先判断
 */

type User struct {
	name string
	pre uint8
	/*
	0000 0000
	0000 0001 IsHongMing
	0000 0010 IsDaRen
	0000 0100 IsVip
	0000 1000 IsSuperVip
	0001 0000 IsLanZuan
	0100 0000 IsHongZuan
	 */
}

//const (
//	HongMing = 1
//	DaRen = 1 << 1
//	Vip = 1 << 2
//	SuperVip = 1 << 3
//	LanZuan = 1 << 4
//	HongZuan = 1 << 5
//	)

func SetPre(user User, isSet bool, posi uint8) User{
	/*
	parameters:
		user : user
		isSet : set the prerogative to true or false
		posi : the position of prerogative
	return:
		user
	 */
	if isSet == true{
		user.pre = user.pre | posi
	}else {
		user.pre = user.pre ^ posi
	}
	return user
}

func GetPre(user User, posi uint8) bool{
	/*
	parameters:
		user : user
		posi : the position of prerogative
	return:
		true or false
	 */
	result := user.pre & posi
	return result == posi
}

//func SetHongMing (user User, isSet bool) User{
//	if isSet == true{
//		user.pre = user.pre | HongMing
//	}else {
//		user.pre = user.pre ^ HongMing
//	}
//	return user
//}
//func GetHongMing(user User) bool{
//	result := user.pre & 1
//	return result == 1
//}
//
//func SetDaRen(user User, isSet bool) User{
//	if isSet == true{
//		user.pre = user.pre | DaRen
//	}else {
//		user.pre = user.pre ^ DaRen
//	}
//	return user
//}
//
//func GetDaRen(user User) bool{
//	result := user.pre & DaRen
//	return result == DaRen
//}
//
//
//func SetVip(user User, isSet bool) User{
//	if isSet {
//		user.pre = user.pre | Vip
//	}else {
//		user.pre = user.pre ^ Vip
//	}
//	return user
//}
//
//func GetVip(user User) bool{
//	result := user.pre & Vip
//	return  result == Vip
//}
