package initializers

import (
	"golang.org/x/crypto/bcrypt"
	"headless-cms/config/roles"
	"headless-cms/types"
	"log"
	"os"
)

func CreateSuperAdminIfNotExists() {
	var superAdmin types.User

	// Check if super admin exists
	retireved := DB.Where("role_type = ?", roles.SuperAdmin).First(&superAdmin)
	// check if super_admin already exists
	if superAdmin.ID != 0 {
		log.Println("Super admin already exists")
		return
	}
	if retireved.Error != nil && retireved.Error.Error() != "record not found" {
		log.Fatalf("Error checking if super admin exists: %v", retireved.Error)
		return
	}

	superAdminUsername := os.Getenv("SUPER_ADMIN_USERNAME")
	superAdminPassword := os.Getenv("SUPER_ADMIN_PASSWORD")
	if superAdminUsername == "" || superAdminPassword == "" {
		log.Fatalf("Error: missing super admin environment variables")
	}

	// Create super admin
	superAdmin.Username = superAdminUsername
	superAdmin.ID = 0
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(superAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing super admin password: %v", err)
		return
	}
	superAdmin.Password = string(hashedPassword)
	superAdmin.RoleType = roles.SuperAdmin
	// Save the super admin to the database
	saved := DB.Create(&superAdmin)
	if saved.Error != nil {
		log.Fatalf("Error saving super admin: %v", saved.Error)
		return
	}
	log.Println("Super admin created")
	return
}
