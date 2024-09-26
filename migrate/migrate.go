package main

import (
	"headless-cms/initializers"
	"headless-cms/types"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&types.Type{})
	initializers.DB.AutoMigrate(&types.Content{})
	initializers.DB.AutoMigrate(&types.Data{})
	initializers.DB.AutoMigrate(&types.Attribute{})
	initializers.DB.AutoMigrate(&types.Child{})
}
