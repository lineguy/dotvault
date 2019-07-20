package main

import (
    "encoding/json"
    "fmt"
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
    Path   string `json:"path"`
    Owner  string `json:"owner"`
    Group  string `json:"group"`
    Chmod  string `json:"chmod"`
}

// check function for checking for errors
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// download function for getting data from
// lastpass and writing to file
func download(chmod os.FileMode, path string) {

    // set command and args
    command := "lpass"
    args := []string{"show", path, "--notes"}

    //run command
    cmd := exec.Command(command, args...)
    out, err := cmd.CombinedOutput()
    check(err)

    // write output
    content := []byte(out)
    err = ioutil.WriteFile(path, content, chmod)
    check(err)

}

// chown function for changing owner / group
// of a file
func chown(owner string, group string, path string) {

    // combine owner and group
    ownergroup := fmt.Sprint(owner +":"+ group)

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
        fmt.Println("Loop Iteration: ", i)
        fmt.Println("Path: " + files.Files[i].Path)
        fmt.Println("Owner: " + files.Files[i].Owner)
        fmt.Println("Group: " + files.Files[i].Group)
        fmt.Println("Chmod: " + files.Files[i].Chmod)
        fmt.Println("")

        // convert chmod string to base8 for os.FileMode
        out, err := strconv.ParseInt(files.Files[i].Chmod, 8, 32)
        check(err)
        chmod := os.FileMode(out)

        // call download function
        download(chmod, files.Files[i].Path)
    }
}
