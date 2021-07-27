package main
import ("fmt"
        "time")
func portal1(chanell1 chan string){
	time.Sleep(3*time.Second)
	chanell1<-"welcome to chanell1"
}
func portal2(chanell2 chan string){
  time.Sleep(9*time.Second)
  chanell2<-"welcome to chanell2"
}
func main(){
	/*fmt.Println("start main method")
	//creating channnel
        ch:=make(chan int)
	go myfunc(ch)
	ch<-23
	fmt.Println("end main method")*/
	r1:=make(chan string)
	r2:=make(chan string)
	go portal1(r1)
	go portal2(r2)
	select{
	case op1:=<-r1:
		fmt.Println(op1)
	case op2:=<-r2:
		fmt.Println(op2)
	}
}
