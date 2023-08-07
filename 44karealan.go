package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func kare(kenar int) (alan int, cevre int) {
	alan = kenar * kenar
	cevre = kenar * 4
	return
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("karenin kenar uzunluğunu giriniz =   ")
	scanner.Scan()
	kenar, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	karealan, karecevre := kare(int(kenar))

	fmt.Println("karenin alanı = ", karealan)
	fmt.Println("karenin cevresi = ", karecevre)

}
