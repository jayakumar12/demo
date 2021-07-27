package main
import "fmt"
func main(){
    var name string
    name="jayakumar"
    fmt.Println(name)
    fmt.Printf("%q %T",name,name)
    x:=2000
    fmt.Println(x)
    fmt.Printf("%v",x)
    a:=10
    b:=20
    if a>b{
       fmt.Println("a is greater:",a)
    }else{
      fmt.Println("b is greater:",b)
    }
}


