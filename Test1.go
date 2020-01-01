package main

import "fmt"

// User  Define Struct:
type User struct {
	FirstName, LastName string
}

// Greeting lskgkh iosdgj sdpog psdofjksd fpsdf pfjksd fpsdf opdfs
// Greeting gfhljh jhfg pgh :
func (U *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", U.FirstName, U.LastName)
}

func main() {
	u := User{"pankaj", "Kumar"}

	fmt.Println(u.Greeting())
}
