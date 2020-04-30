package helpers

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateIds() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("There is problem to generate uuid")
	}
	return uuid.String()
}
