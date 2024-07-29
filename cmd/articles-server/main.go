package main

import (
	"fmt"

	"github.com/massivemadness/articles-server/api"
)

func main() {
	fmt.Printf("Hello World = %s", api.ErrNotFound.Error())
}
