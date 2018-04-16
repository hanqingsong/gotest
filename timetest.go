package main
import( 
	"time"
	"fmt"
)
func main(){
	start:=time.Now()
	fmt.Printf("time:%s\n",start.Local())
}
