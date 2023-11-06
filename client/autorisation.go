package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Autorisation() (name string, password string) {
	var isAdmin bool

	fmt.Println("Are you Administrator? yes/no")

	firstAns := scan1()
	firstAns = strings.ToLower(firstAns)
	if strings.Contains(firstAns, "yes") {
		isAdmin = true
	} else {
		isAdmin = false
	}

	if !isAdmin {
		return "", ""
	}

	fmt.Println("Please, input administrator-name :")
	secondAns := scan1()
	fmt.Println("Please, input administrator-password :")
	thirdAns := scan1()

	return secondAns, thirdAns
}

func scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		return "" //fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
