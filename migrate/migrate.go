package main

import (
	initializers "crud/intializers"
	"crud/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	fmt.Println("MIGRATEDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")
	fmt.Println("MIGRATEDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD", initializers.DB)

	initializers.DB.AutoMigrate(&models.Post{})

}
