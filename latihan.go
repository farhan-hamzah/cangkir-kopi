package main
import "fmt"

func cangkirKopi(k1, k2 *int, g1, g2 int){
	var i int
	i = 0
	for *k1 >= g1 && *k2 >= g2 {
		*k1 -= g1
		*k2 -= g2
		i++
	} 
	fmt.Println(i)
}
func main(){
	var k1, k2, g1, g2 int
	fmt.Scan(&k1, &k2, &g1, &g2)
	cangkirKopi(&k1, &k2, g1, g2)
}


