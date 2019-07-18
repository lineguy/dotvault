package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// Users struct which contains
// an array of users
type Files struct {
    Files []File `json:"files"`
}

// User struct which contains a name
// a type and a list of social links
type File struct {
    Path   string `json:"path"`
    User   string `json:"user"`
    Group  string `json:"group"`
    Chmod  string `json:"chmod"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("files.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened files.json")
    fmt.Println("")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var files Files

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'files' which we defined above
    json.Unmarshal(byteValue, &files)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(files.Files); i++ {
        fmt.Println("Loop Pass: ", i)
        fmt.Println("File Path: " + files.Files[i].Path)
        fmt.Println("File Owner: " + files.Files[i].User)
        fmt.Println("File Group: " + files.Files[i].Group)
        fmt.Println("Chmod Permissions: " + files.Files[i].Chmod)
        fmt.Println("")
    }

}
