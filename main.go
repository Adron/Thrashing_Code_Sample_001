package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
)

type userInformationTest struct {
	Name string
	UserId		int64
	GroupId		int64
	HomeDir	string
	UserName string
	GroupIds []string
	GoPath string
	EnvVar string
}
type userInformation struct {
	Name string `json:"name"`
	UserId		int64 `json:"userid"`
	GroupId		int64 `json:"groupid"`
	HomeDir	string `json:"homedir"`
	UserName string `json:"username"`
	GroupIds []string `json:"groupids"`
	GoPath string `json:"gopath"`
	EnvVar string `json:"environmentvariable""`
}

func main() {
	currentUser, err := user.Current()
	check(err)
	funcName(currentUser)
}

func funcName(currentUser *user.User) {
	groupsIds, err := currentUser.GroupIds()
	check(err)
	userId, err := strconv.ParseInt(currentUser.Uid, 6, 64)
	check(err)
	groupId, err := strconv.ParseInt(currentUser.Gid, 6, 64)
	check(err)

	res1D := &userInformation{
		Name:   currentUser.Name,
		UserId: userId,
		GroupId: groupId,
		HomeDir: currentUser.HomeDir,
		UserName: currentUser.Username,
		GroupIds: groupsIds,
		GoPath: os.Getenv("GOPATH"),
		EnvVar: os.Getenv("NEW_STRING_ENVIRONMENT_VARIABLE"),
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	filename := "collected_values.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		textFileContent := []byte("File created.\n")
		err := ioutil.WriteFile(filename, textFileContent, 0644)
		check(err)

		appendFileContents(err, filename, "File didn't exist, created, and append this.\n")
	} else {
		appendFileContents(err, filename, "File existed, adding these contents.\n")
	}
}

func appendFileContents(err error, filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
