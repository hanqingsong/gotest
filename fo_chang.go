package main
import "fmt"
func main(){
	items := [...]int{10,20,30,40,50}
	for i,item:=range items{
		item *=2
		fmt.Printf("%d\n",item)
		items[i]=item
	}
	for _,item:=range items{
		fmt.Printf("%d\n",item)
	}
}
