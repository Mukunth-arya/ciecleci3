package main



import (

	"testing"
)

func Testsample(t* testing.T)bool{


     var a = []int{} 

	 if a == nil{
		return true
		t.Errorf("Soory value as expected")
	 }
    return true

}