package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "N0000000025;0000445;-1;8007;50017;;;;;;{-1,-1,-1};0;;-1;;-1;-1;;13624083777;0;;240;;;;;20170502120224;20170502120224;0;3;;0;;;;;;;;;0;;;;;;;;;0502121130700000100100;2;;;;99;;;701001;;700137325000;;;;1;2;;2;0;0;;;;40288000;;;7000001;;;;100;0;;;;;;;;0;0;;;0;;-1;20170502121747;GP20170502147.250;9999;;;;;;;;;1;;;;;01;;;;;;;;0;;;;;0;;;;;;;;;26;0;;20170502121747;;;;;;;;;;0;0;0;;0;;;;;;;;;;;;;;;;;250;;;;;;;;;;85:0:0;;;;;;;;;;0;;;700137325000;;;;;;"
	strList := strings.Split(str, ";")
	fmt.Println(strList[18])



	}
