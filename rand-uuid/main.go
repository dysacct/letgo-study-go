package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	uuid := id.String()

	fmt.Println("生成的uuid: ", uuid)
}
