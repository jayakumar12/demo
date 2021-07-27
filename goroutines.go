package main
import ("fmt"
        "time")
func display(str string){
    for w:=0;w<6;w++{
    fmt.Println(str)
    time.Sleep(1*time.Second)
    }
}
func main(){
	go display("welcome")
	display("HELLO World")
}


