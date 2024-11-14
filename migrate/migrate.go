package main

import (
	"headless-cms/initializers"
	"headless-cms/types"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	var err error = nil

	err = initializers.DB.AutoMigrate(&types.User{})

	err = initializers.DB.AutoMigrate(&types.Type{})
	if err != nil {
		return
	}
	err = initializers.DB.AutoMigrate(&types.Content{})
	if err != nil {
		return
	}
	err = initializers.DB.AutoMigrate(&types.Data{})
	if err != nil {
		return
	}
	err = initializers.DB.AutoMigrate(&types.Attribute{})
	if err != nil {
		return
	}
	err = initializers.DB.AutoMigrate(&types.Child{})
	if err != nil {
		return
	}
	log.Println("Migrated all tables successfully")
	// add basic types in type table like integer, string, boolean, etc.
	basicTypes := []types.Type{
		{Name: "integer"},
		{Name: "string"},
		{Name: "boolean"},
		{Name: "float"},
		{Name: "date"},
	}

	for _, t := range basicTypes {
		if err := initializers.DB.FirstOrCreate(&t, types.Type{Name: t.Name}).Error; err != nil {
			log.Fatalf("Error adding basic type %v: %v", t.Name, err)
		}
	}
}
