package file

import (
	"DinnerPinner/advt"
	"encoding/json"
	"fmt"
	"os"
)

func createAndSaveJson() {
	user := advt.Users{
		FirstName: "bruh",
		LastName:  "bruh",
		Email:     "bruh",
		Password:  "bruh",
	}

	jsonBytes, _ := json.Marshal(user)

	jsoFile, _ := os.Create("user.json")
	defer jsoFile.Close()

	bytes, _ := jsoFile.Write(jsonBytes)
	fmt.Printf("%d bytes written to json file", bytes)

}
