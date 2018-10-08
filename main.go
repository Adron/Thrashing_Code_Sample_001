package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
)

type UserInformation struct {
	Name     string   `json:"name"`
	UserId   int64    `json:"userid"`
	GroupId  int64    `json:"groupid"`
	HomeDir  string   `json:"homedir"`
	UserName string   `json:"username"`
	GroupIds []string `json:"groupids"`
	GoPath   string   `json:"gopath"`
	EnvVar   string   `json:"environmentvariable"`
}

func main() {
	currentUser, err := user.Current()
	check(err)
	changeDataForMarshalling(currentUser)
}

func changeDataForMarshalling(currentUser *user.User) {
	groupsIds, err := currentUser.GroupIds()
	check(err)
	userId, err := strconv.ParseInt(currentUser.Uid, 6, 64)
	check(err)
	groupId, err := strconv.ParseInt(currentUser.Gid, 6, 64)
	check(err)

	workingUserInformation := &UserInformation{
		Name:     currentUser.Name,
		UserId:   userId,
		GroupId:  groupId,
		HomeDir:  currentUser.HomeDir,
		UserName: currentUser.Username,
		GroupIds: groupsIds,
		GoPath:   os.Getenv("GOPATH"),
		EnvVar:   os.Getenv("NEW_STRING_ENVIRONMENT_VARIABLE"),
	}

	resultingUserInformation, _ := json.Marshal(workingUserInformation)

	filename := "collected_values.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		writeFileContents(err, filename, string(resultingUserInformation))
	} else {
		newUserInformation, err := openFileMakeChanges(filename)
		writeFileContents(err, filename, newUserInformation)
	}
}

func openFileMakeChanges(filename string) (string, error) {
	jsonFile, err := os.Open(filename)
	check(err)
	defer jsonFile.Close()
	var changingUserInformation UserInformation
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &changingUserInformation)
	changingUserInformation.Name = "Adron Hall"
	changingUserInformation.GoPath = "/where/is/the/goland"
	newUserInformation, _ := json.Marshal(changingUserInformation)
	return string(newUserInformation), err
}

func writeFileContents(err error, filename string, text string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0600)
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
