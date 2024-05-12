package services

import (
	"time"

	"github.com/AliRamdhan/compstudioserver/auth"
	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (au *AuthService) GetAllUser() ([]model.User, error) {
	var user []model.User
	if err := config.DB.Preload("Role").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// func (au *AuthService) RegisterAuth(user *model.User) error {
// 	user.UserID = uuid.New()
// 	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
// 	return config.DB.Create(user).Error
// }

func (au *AuthService) RegisterUser(user *model.User) error {
	// Hash user password
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	// Create user record
	user.RoleUser = 2
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	record := config.DB.Create(user)
	if record.Error != nil {
		return record.Error
	}

	return nil
}

func (au *AuthService) LoginAuth(email, password string) (*model.User, error) {
	var user model.User
	// Check if email exists
	record := config.DB.Where("email = ?", email).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}

	// Check if password is correct
	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (au *AuthService) GenerateToken(userId int, email, username string, roleId int) (string, error) {
	return auth.GenerateJWT(userId, email, username, roleId)
}

//	func (au *AuthService) Home() string {
//		return "Welcome to the home page!"
//	}

func (au *AuthService) Home(tokenString string) (*auth.JWTClaim, error) {
	// Validate the token
	user, err := auth.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (au *AuthService) Home(userID uuid.UUID) (string, error) {
// 	// Retrieve user data based on userID
// 	var user model.User
// 	if err := config.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
// 		return "", err
// 	}

// 	// Construct welcome message with user data
// 	message := "Welcome " + user.Username + " to the home page!"
// 	return message, nil
// }

// func (au *AuthService) LoginAuth() ([]model.User, error) {
// 	var user []model.User
// 	if err := config.DB.Find(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (au *AuthService) LogoutAuth() ([]model.User, error) {
// 	var customer []model.User
// 	if err := config.DB.Find(&customer).Error; err != nil {
// 		return nil, err
// 	}
// 	return customer, nil
// }

// func (au *AuthService) HomeAuth() ([]model.User, error) {
// 	var customer []model.User
// 	if err := config.DB.Find(&customer).Error; err != nil {
// 		return nil, err
// 	}
// 	return customer, nil
// }
