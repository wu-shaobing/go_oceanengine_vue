package auth

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

// VerifyPassword 验证密码
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePassword 验证密码复杂度
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("密码长度不能少于8位")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("密码必须包含大小写字母、数字和特殊字符")
	}

	return nil
}
