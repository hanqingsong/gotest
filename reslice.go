package main 
import "fmt"
func main(){
	var ar=[10]int{0,11,22,31,41,51,61,71,81,91}
	var a=ar[3:7]
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a=a[0:4]
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	
}
