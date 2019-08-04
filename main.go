package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
    "flag"
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
	Name      string `json:"name"`
	Lpass_id  string `json:"lpass_id"`
	Path      string `json:"path"`
	Owner     string `json:"owner"`
	Group     string `json:"group"`
	Chmod     string `json:"chmod"`
}

// check function for checking for errors
func check(msg string, err error) {
	if err != nil {
	    fmt.Println(msg)
        panic(err)
    }
}


func dependencies() {

    // set command and args
	command := "which"
	args := []string{"lpass"}

	//run command
	cmd := exec.Command(command, args...)
	err := cmd.Start()
    check("Error: Can't find lpass command, please make sure lastpass-cli is installed.", err)


}

// download function for getting data from
// lastpass and writing to file
func download(chmod string, id string, path string) {

	// convert chmod string to base8 for os.FileMode
	out1, err := strconv.ParseInt(chmod, 8, 32)
    check("Error: Can't convert chmod string.", err)
	perm := os.FileMode(out1)

	// set command and args
	command := "lpass"
	args := []string{"show", id, "--notes"}

	//run command
	cmd := exec.Command(command, args...)
	out2, err := cmd.CombinedOutput()
    check("Error: Can't run lpass command to download data.", err)

	// write output
	content := []byte(out2)
	err = ioutil.WriteFile(path, content, perm)
    check("Error: Can't write to file.", err)

}

// upload function for sending data to
// lastpass
func upload(path string, id string) {

	// read file
	file, err := ioutil.ReadFile(path)
    check("Error: Can't read file.", err)

	// set command and args
	command := "lpass"
	args := []string{"edit", id, "--non-interactive", "--notes"}

	//run command
	cmd := exec.Command(command, args...)
	stdin, err := cmd.StdinPipe()
    check("Error: Can't pipe data to lpass command to upload data.", err)

	err = cmd.Start()
    check("Error: Can't run lpass command to upload data.", err)

	_, err = io.WriteString(stdin, string(file))
	check("", err)
}

// chown function for changing owner / group
// of a file
func chown(owner string, group string, path string) {

	// test file
	_, err := os.Stat(path)
    check("Error: File doesn't exit.", err)

    // combine owner & group
    ownergroup := fmt.Sprint(owner + ":" + group)

	// set command and args
	command := "chown"
	args := []string{ownergroup, path}

	//run command
	cmd := exec.Command(command, args...)
    err = cmd.Start()
    check("Error: Can't run chown command on file.", err)

}

// chmod function for changing permissions
// of a file
func chmod(chmod string, path string) {

	// test file
	_, err := os.Stat(path)
    check("Error: File doesn't exit.", err)

	// convert chmod string to base8 for os.FileMode
	out1, err := strconv.ParseInt(chmod, 8, 32)
    check("Error: Can't convert chmod string.", err)
	perm := os.FileMode(out1)

	err = os.Chmod(path, perm)
    check("Error: Can't run chmod command on file.", err)

}

func main() {

    // check dependencies
    dependencies()

    is_download := flag.Bool("download", false, "a bool")
    is_upload := flag.Bool("upload", false, "a bool")

    flag.Parse()

	// Open our jsonFile
	jsonFile, err := os.Open("files.json")
    check("Error: Can't read files.json.", err)

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

		// call download function
        if *is_download == true {
            download(files.Files[i].Chmod, files.Files[i].Lpass_id, files.Files[i].Path)
		    chown(files.Files[i].Owner, files.Files[i].Group, files.Files[i].Path)
		    chmod(files.Files[i].Chmod, files.Files[i].Path)
        }

		// call upload function
        if *is_upload == true {
			upload(files.Files[i].Path, files.Files[i].Lpass_id)
        }

	}

}
