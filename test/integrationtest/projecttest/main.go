package main

import (
	"fmt"
	"os"

	"github.com/lunarforge/flamingo_commerce/test/integrationtest/projecttest/helper"
)

func main() {
	if os.Getenv("RUN") == "1" {
		info := helper.BootupDemoProject("../config")
		<-info.Running
		fmt.Println("Server exited")
	} else {
		fmt.Println("Generating GraphQL")
		helper.GenerateGraphQL()
	}
}
