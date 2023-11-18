package main

import(
	"fmt"
	"strings"
)

func main(){
	gmail:= "ridhofahreza@gmail.co.id"
	fmt.Println(getDomain(gmail))
}


func splitGmail(params string)string{
	splitMail:= strings.Split(params, "@")
	getDomain:=splitMail[1]
	return getDomain;
}

func getDomain(params string)string{
	dataParams := params

	getSplit := splitGmail(dataParams)
	firstDot := false
	var getDomain string;
	var getTld string;
	for i:=0; i<len(getSplit); i++{
		if string(getSplit[i]) == "." && !firstDot{
			firstDot = true
			continue
		}
		if firstDot == false{
			getDomain+=string(getSplit[i])
		}
		if firstDot == true{
			getTld+= string(getSplit[i])
		}
	}

	return "Domain:" + getDomain + " TLD:" + getTld
}