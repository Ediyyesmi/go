package tekrarlar

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	dosyam, _ := os.Create("log.txt")
	log.SetOutput(dosyam)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
}

func control(err error) {
	if err != nil {
		panic(err)
	}
}

func Demo3() {

	var admininfo int = 1357
	var adminpass int = 7878
	var studentinfo int = 2468
	var studentpass int = 7878
	var user, pass, name, choice, i int

	fmt.Println("select user: (admin:0, student:1)")
	fmt.Scanln(&user)

	if user == 0 {
		fmt.Println("admin login (you can try 5 times)")
		for i = 4; i >= 0; i-- {
			fmt.Println("enter your user name: ")
			fmt.Scanln(&name)
			fmt.Println("enter your password: ")
			fmt.Scanln(&pass)

			if admininfo == name && adminpass == pass {
				fmt.Println("login successful")
				log.Println("admin: ", name, " - login successful")
				fmt.Println("view logs(0), log out(1)")
				fmt.Scanln(&choice)

				if choice == 0 {
					dosyam, err := ioutil.ReadFile("log.txt")
					control(err)
					fmt.Println(string(dosyam))
					break
				}
				break
			} else {
				log.Println("admin: ", name, " - login unsuccessful")
				fmt.Println("this informatins are wrong, please try again")
				fmt.Println("number of remaining entries: ", i)
			}
			if i == 0 {
				fmt.Println("login failed")
				log.Println("login failed")
			}
		}

	} else {
		fmt.Println("student login (you can try 5 times)")
		for i = 4; i >= 0; i-- {
			fmt.Println("enter your user name: ")
			fmt.Scanln(&name)
			fmt.Println("enter your password: ")
			fmt.Scanln(&pass)

			if studentinfo == name && studentpass == pass {
				fmt.Println("login successful")
				log.Println("student: ", name, " - login successful")
				fmt.Println("view logs(0), log out(1)")
				fmt.Scanln(&choice)

				if choice == 0 {
					dosyam, err := ioutil.ReadFile("log.txt")
					control(err)
					fmt.Println(string(dosyam))
					break
				}
				break
			} else {
				log.Println("student: ", name, " - login unsuccessful")
				fmt.Println("this informatins are wrong, please try again")
				fmt.Println("number of remaining entries: ", i)
			}
			if i == 0 {
				fmt.Println("login failed")
				log.Println("login failed")
			}
		}
	}

}
