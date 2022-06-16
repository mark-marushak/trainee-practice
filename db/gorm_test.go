package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestTypeGorm(t *testing.T) {
	db, err := Gorm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TestTypeGorm]")
	}
	assert.IsTypef(t, &gorm.DB{}, db, "[TestTypeGorm]: Gorm() function returned type nil")
}
