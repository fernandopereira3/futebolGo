package main

import (
	"fmt"
	"math/rand"
)

const time1 string = "Argentina"
const time2 string = "Germany"
const time3 string = "Japan"
const time4 string = "Brazil"
var time1Pontos float32
var time2Pontos float32
var time3Pontos float32
var time4Pontos float32

// string,float32,string,float32,string,float32,string,float32
func rodada1() (string,float32,string,float32,string,float32,string,float32){

	time1gols := rand.Intn(6)
	time2gols := rand.Intn(6)
	time3gols := rand.Intn(6)
	time4gols := rand.Intn(6)
	 

	for rodada := 0; rodada <= 9; rodada++{

	// time1 X time2 
	if time1gols == time2gols{
		time1Pontos += 1.5
		time2Pontos += 1.5
	}else if time1gols > time2gols{
		time1Pontos += 3
		time2Pontos += 0
	}else if time2gols > time1gols{
		time1Pontos += 0
		time2Pontos += 3
	}
	// time1 X time3
	if time1gols == time3gols{
		time1Pontos += 1.5 
		time3Pontos += 1.5
	}else if time1gols > time3gols{
		time1Pontos += 3
		time3Pontos += 0
	}else if time3gols > time1gols{
		time1Pontos += 0
		time3Pontos += 3
	}
	// // time1 X time4
	if time1gols == time4gols{
		time1Pontos += 1.5 
		time4Pontos += 1.5
	}else if time1gols > time4gols{
		time1Pontos += 3
		time4Pontos += 0
	}else if time4gols > time1gols{
		time1Pontos += 0
		time4Pontos += 3
	}
	// rodada2
	// time2 X time1
	if time2gols == time1gols{
		time1Pontos += 1.5
		time2Pontos += 1.5
	}else if time1gols > time2gols{
		time1Pontos += 3
		time2Pontos += 0
	}else if time2gols > time1gols{
		time1Pontos += 0
		time2Pontos += 3
	}
	// time2 X time3
	if time2gols == time3gols{
		time2Pontos += 1.5 
		time3Pontos += 1.5
	}else if time2gols > time3gols{
		time2Pontos += 3
		time3Pontos += 0
	}else if time3gols > time2gols{
		time2Pontos += 0
		time3Pontos += 3
	}
	// // time2 X time4
	if time2gols == time4gols{
		time1Pontos += 1.5 
		time4Pontos += 1.5
	}else if time2gols > time4gols{
		time2Pontos += 3
		time4Pontos += 0
	}else if time4gols > time2gols{
		time2Pontos += 0
		time4Pontos += 3
	}

	// rodada3
	// time3 X time1
	if time3gols == time1gols{
		time3Pontos += 1.5
		time2Pontos += 1.5
	}else if time3gols > time1gols{
		time3Pontos += 3
		time1Pontos += 0
	}else if time1gols > time3gols{
		time3Pontos += 0
		time1Pontos += 3
	}
	// time2 X time3
	if time2gols == time3gols{
		time2Pontos += 1.5 
		time3Pontos += 1.5
	}else if time2gols > time3gols{
		time2Pontos += 3
		time3Pontos += 0
	}else if time3gols > time2gols{
		time2Pontos += 0
		time3Pontos += 3
	}
	// // time3 X time4
	if time3gols == time4gols{
		time3Pontos += 1.5 
		time4Pontos += 1.5
	}else if time3gols > time4gols{
		time3Pontos += 3
		time4Pontos += 0
	}else if time4gols > time3gols{
		time3Pontos += 0
		time4Pontos += 3
	}

	// rodada4
	// time4 X time1
	if time4gols == time1gols{
		time1Pontos += 1.5
		time4Pontos += 1.5
	}else if time1gols > time4gols{
		time1Pontos += 3
		time4Pontos += 0
	}else if time4gols > time1gols{
		time1Pontos += 0
		time4Pontos += 3
	}
	// time4 X time2
	if time4gols == time2gols{
		time4Pontos += 1.5 
		time2Pontos += 1.5
	}else if time4gols > time2gols{
		time4Pontos += 3
		time2Pontos += 0
	}else if time2gols > time4gols{
		time4Pontos += 0
		time2Pontos += 3
	}
	// // time4 X time3
	if time4gols == time3gols{
		time3Pontos += 1.5 
		time4Pontos += 1.5
	}else if time3gols > time4gols{
		time3Pontos += 3
		time4Pontos += 0
	}else if time4gols > time3gols{
		time3Pontos += 0
		time4Pontos += 3
	}
	}
	return time1,time1Pontos,time2,time2Pontos,time3,time3Pontos,time4,time4Pontos
	
}


func main() {
	fmt.Println(rodada1())
}