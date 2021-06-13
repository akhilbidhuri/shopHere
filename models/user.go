package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int       `gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" validate:"required" gorm:"type:varchar(100);not null"`
	Username  string    `json:"username" validate:"required" gorm:"type:varchar(100);not null;UNIQUE"`
	Password  string    `json:"password" validate:"required" gorm:"type:varchar(256);not null"`
	Token     string    `json:"token" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	Cart_id   int       `json:"cart_id" gorm:"type:integer"` //gorm:"type:integer REFERENCES carts(id);default:null"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) ConvertPwdToHash() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Create(db *gorm.DB) error {
	u.ConvertPwdToHash()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Create(&u).Error
		if err != nil {
			return err
		}
		cart := Cart{
			User_id:      u.Id,
			Is_purchased: false,
		}
		err = cart.Create(tx)
		if err != nil {
			return err
		}
		u.Cart_id = cart.Id
		if err = u.Update(tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) GetUserByUname(db *gorm.DB, userName string) error {
	err := db.Debug().Model(User{}).Where("username = ?", userName).Take(&u).Error
	if err != nil {
		return err
	}
	return err
}

func (u *User) GetUserByID(db *gorm.DB, id int) error {
	err := db.Debug().Model(User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return err
	}
	return err
}

func (u *User) GetUserByToken(db *gorm.DB, token string) error {
	err := db.Debug().Model(User{}).Where("token = ?", token).Take(&u).Error
	if err != nil {
		return err
	}
	return err
}

func (u *User) Update(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Save(&u).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.Id,
		"name":       u.Name,
		"username":   u.Username,
		"cart_id":    u.Cart_id,
		"created_at": u.CreatedAt.String(),
		"token":      u.Token,
	}
}
