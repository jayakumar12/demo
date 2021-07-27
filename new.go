package main
import  "fmt"
func main(){
	var x[10]int=[10]int{4,4,5,5,6,6,6,6,7,8}
	
	for i:=0;i<10;i++{
	   if  i==0{
		   if !(x[i+1]==x[i]){
		   fmt.Println(x[i])
		   }
	   }else if i==9{
		   if !(x[i-1]==x[i]){
		   fmt.Println(x[i])
		   }
	   }else{
		   if !(x[i-1]==x[i] ||x[i+1]==x[i]) {
		   fmt.Println(x[i])
		   }
	   }
	}
}
