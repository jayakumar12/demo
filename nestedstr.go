package main
import "fmt"

type student struct{
 name string
 branch string
 year int
}
type teacher struct{
	name string
	subject string
	exp int
	details student
	}
type  Saiyan struct{
	name string
	power int 
}
func main(){
	result:=teacher{
		name:"jaya",
		subject:"golang",
		exp:5,
		details:student{
			name:"hari",
			branch:"biotech",
			year:2021},
	}

	fmt.Println("deatils of teacher:")
	fmt.Println(result.name)
	fmt.Println(result.subject)
	fmt.Println(result.exp)
	fmt.Println("details of student:")
	fmt.Println(result.details.name)
	fmt.Println(result.details.branch)
	fmt.Println(result.details.year)

	goke:=Saiyan{"Goku",90000}
	Super(goke)
	fmt.Println(goke.power) //value not chnaged because super and power are pointing to different locations 
	goke1:=&Saiyan{"topu",10000}
	Super1(goke1)
	fmt.Println(goke1.power)
	goke2:=&Saiyan{"ramesh",2000}
	goke2.super()
	fmt.Println(goke2.power)

}
func Super(s Saiyan){
	s.power+=10000
}
func Super1(s *Saiyan){  //pointing to same location of power so updated
	s.power+=9000
}
func (s *Saiyan)super(){
	s.power+=1000
}
