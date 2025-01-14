package main

import (
	fnDevtool "fiber-postgres-api/assets/devtool"
	"fmt"
)

func main() {
	//##### Test #####
	errReplaceDeletedAt := fnDevtool.ReplaceDeletedAtFilesInDirectory("../modules/entities")
	if errReplaceDeletedAt != nil {
		fmt.Println("Error:", errReplaceDeletedAt)
		return
	} else {
		fmt.Println("File updated successfully.")
	}
}
