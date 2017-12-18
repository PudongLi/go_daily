package lib
import "fmt"

func SayHello(){
	fmt.Println("hello")
}
var(
	Name = "lipd"
)
func init(){
	fmt.Println("lib init...")
}