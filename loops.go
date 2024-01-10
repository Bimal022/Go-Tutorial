package main
import(
	"fmt"
)

func main(){
	for i:=1; i<10; i++{
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	/**********Break And Continue**********/
	for i:=0; i<20; i++{
		if i == 10{
			break
		}else if i%2 != 0{
			continue
		}else{
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	i:=1
	for i<=10{
		fmt.Printf("2 * %d = %d", i, i*2)
		fmt.Println()
		i+=1
	}
}