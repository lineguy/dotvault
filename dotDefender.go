package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// initialize variables

// Files struct which contains
// an array of files
type Files struct {
	Files []File `json:"files"`
}

// Files struct which contains a path,
// owner, group and chmod
type File struct {
	Path  string `json:"path"`
	Owner string `json:"owner"`
	Group string `json:"group"`
	Chmod string `json:"chmod"`
}

// check function for checking for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// download function for getting data from
// lastpass and writing to file
func download(chmod string, path string) {

	// convert chmod string to base8 for os.FileMode
	out1, err := strconv.ParseInt(chmod, 8, 32)
	check(err)
	perm := os.FileMode(out1)

	// set command and args
	command := "lpass"
	args := []string{"show", path, "--notes"}

	//run command
	cmd := exec.Command(command, args...)
	out2, err := cmd.CombinedOutput()
	check(err)

	// write output
	content := []byte(out2)
	err = ioutil.WriteFile(path, content, perm)
	check(err)

}

// upload function for sending data to
// lastpass
func upload(path string) {

	// read file
	file, err := ioutil.ReadFile(path)
	check(err)

	// set command and args
	command := "lpass"
	args := []string{"edit", path, "--non-interactive", "--notes"}

	//run command
	cmd := exec.Command(command, args...)
	stdin, err := cmd.StdinPipe()
	check(err)

	err = cmd.Start()
	check(err)

	_, err = io.WriteString(stdin, string(file))
	check(err)
}

// chown function for changing owner / group
// of a file
func chown(owner string, group string, path string) {

	// combine owner and group
	ownergroup := fmt.Sprint(owner + ":" + group)

	// set command and args
	command := "chown"
	args := []string{ownergroup, path}

	//run command
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	check(err)

	// print output
	fmt.Println(out)

}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("files.json")
	check(err)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Files array
	var files Files

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'files' which we defined above
	json.Unmarshal(byteValue, &files)

	// we iterate through every file within our files array and
	// print out the file path, owner, group and chmod
	for i := 0; i < len(files.Files); i++ {
		//fmt.Println("Loop Iteration: ", i)
		fmt.Println("Path: " + files.Files[i].Path)
		//fmt.Println("Owner: " + files.Files[i].Owner)
		//fmt.Println("Group: " + files.Files[i].Group)
		//fmt.Println("Chmod: " + files.Files[i].Chmod)
		fmt.Println("")

		// call download function
		//download(files.Files[i].Chmod, files.Files[i].Path)

		// call upload function
		upload(files.Files[i].Path)

		// call backup function
		//backup(backuptype, files.Files[i].Path)

	}
}
