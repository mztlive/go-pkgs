// BEGIN: qv9c3f4g6hj2
package bcrypt_test

import (
	"testing"

	"github.com/mztlive/go-pkgs/bcrypt"
)

func TestHashAndCheckHash(t *testing.T) {
	password := "password123"
	hash, err := bcrypt.Hash(password)
	println(hash)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	match := bcrypt.CheckHash(password, hash)
	if !match {
		t.Errorf("Hash does not match password")
	}
}

// END: qv9c3f4g6hj2
