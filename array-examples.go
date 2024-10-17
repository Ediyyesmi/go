package tekrar

import "fmt"

func Demo1(){

	var m, n int= 0,0
	var tek[5] int
	var cift[5] int 
	for i:=0;i<10;i++ {
		if i%2!=0 {
			tek[n]=i
			n++
		}else{
			cift[m]=i
			m++
		}
	}
	
	fmt.Println("tek sayilar: ",tek)
	fmt.Println("cift sayilar: ",cift)


}
