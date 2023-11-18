package main

import (
	"fmt"
	"math"
)

func main(){
    vip:=4
	regular:=0
	student:=0
	day:=21
	fmt.Println(getTicket(vip, regular, student, day))
}

func getTicket(vipTicket int, regularTicket int, studentTicket int, day int) float32{

	var getDiscount float32

	totalTicket:= vipTicket + regularTicket + studentTicket

	vipPrize := vipTicket * 30
	regularPrize := regularTicket * 20
	studentPrize := studentTicket * 10

	sumTotal :=  vipPrize + regularPrize + studentPrize

	if sumTotal >=100{
		if day % 2 == 0{
			if totalTicket <5{
				getDiscount = float32(math.Round(float64(sumTotal))) - (float32(math.Round(float64(sumTotal))) * 10 / 100)
			}else{
				getDiscount = float32(math.Round(float64(sumTotal))) - (float32(math.Round(float64(sumTotal))) * 20 / 100)
			}
		} else{
			if totalTicket <5{
				getDiscount = float32(math.Round(float64(sumTotal))) - (float32(math.Round(float64(sumTotal))) * 15 / 100)
			}else{
				getDiscount = float32(math.Round(float64(sumTotal))) - (float32(math.Round(float64(sumTotal))) * 25 / 100)
			}
		}
	} else{
		getDiscount = float32(sumTotal)
	}
	
	return float32(math.Round(float64(getDiscount))) 
}
