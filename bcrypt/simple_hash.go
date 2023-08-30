package bcrypt

import go_bcrypt "golang.org/x/crypto/bcrypt"

// Hash函数将给定的密码进行哈希处理并返回哈希值和潜在的错误。
func Hash(password string) (string, error) {
	// 生成哈希
	hashedPassword, err := go_bcrypt.GenerateFromPassword([]byte(password), go_bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckHash函数将给定的密码与哈希值进行比较并返回匹配结果（true/false）。
func CheckHash(password, hash string) bool {
	err := go_bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
