package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ---------- Models ----------
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type PasswordHistory struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	OldPwd string
}

// ---------- UoW ----------
type UnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Begin() {
	u.tx = u.db.Begin()
}

func (u *UnitOfWork) Commit() error {
	return u.tx.Commit().Error
}

func (u *UnitOfWork) Rollback() error {
	return u.tx.Rollback().Error
}

func (u *UnitOfWork) Tx() *gorm.DB {
	return u.tx
}

// ---------- Repositories ----------
type UserRepository struct {
	tx *gorm.DB
}

func NewUserRepository(tx *gorm.DB) *UserRepository {
	return &UserRepository{tx: tx}
}

func (r *UserRepository) UpdatePassword(userID uint, newPwd string) error {
	return r.tx.Model(&User{}).
		Where("id = ?", userID).
		Update("password", newPwd).Error
}

type PasswordHistoryRepository struct {
	tx *gorm.DB
}

func NewPasswordHistoryRepository(tx *gorm.DB) *PasswordHistoryRepository {
	return &PasswordHistoryRepository{tx: tx}
}

func (r *PasswordHistoryRepository) AddHistory(userID uint, oldPwd string) error {
	return r.tx.Create(&PasswordHistory{
		UserID: userID,
		OldPwd: oldPwd,
	}).Error
}

// ---------- Service ----------
type UserService struct {
	uow *UnitOfWork
}

func NewUserService(uow *UnitOfWork) *UserService {
	return &UserService{uow: uow}
}

func (s *UserService) ChangePassword(userID uint, oldPwd, newPwd string) error {
	s.uow.Begin()

	phRepo := NewPasswordHistoryRepository(s.uow.Tx())
	if err := phRepo.AddHistory(userID, oldPwd); err != nil {
		s.uow.Rollback()
		return fmt.Errorf("failed to add history: %w", err)
	}

	userRepo := NewUserRepository(s.uow.Tx())
	if err := userRepo.UpdatePassword(userID, newPwd); err != nil {
		s.uow.Rollback()
		return fmt.Errorf("failed to update password: %w", err)
	}

	return s.uow.Commit()
}

// ---------- Main ----------
func main() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	_ = db.AutoMigrate(&User{}, &PasswordHistory{})

	uow := NewUnitOfWork(db)
	service := NewUserService(uow)

	// کاربر تستی
	db.Create(&User{Username: "abolfazl", Password: "123456"})

	// تغییر رمز
	if err := service.ChangePassword(1, "123456", "654321"); err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Password changed successfully")
	}
}
