package main

import "fmt"

func main() {

	sayi:=0 
	fmt.Println("bir sayi giriniz")
	fmt.Scanln(&sayi)
	
	for i:=2;i<=sayi/2;i++ {
		if sayi % i == 0 {
			fmt.Println("sayi asal degil")
			return

		}else{
			fmt.Println("sayi asal")
			return
		}
	}

}
