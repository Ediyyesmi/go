package main

import "fmt"

func main() {

	sayi := 7
	tahmin := 0

	fmt.Println("sayiyi tahmin edin. 5 hak var")
	for i := 1; i <= 5; i++ {
		fmt.Scanln(&tahmin)
		if sayi != tahmin {
			if sayi < tahmin {
				fmt.Println("daha kucuk bir sayi giriniz")
			} else {
				fmt.Println("daha buyuk bir sayi giriniz")
			}

		} else {
			fmt.Println("cevabiniz dogru")
			break
		}
		if i == 5 {
			fmt.Println("hakkiniz kalmadi. dogru sayi: ", sayi)
		}
	}

}
