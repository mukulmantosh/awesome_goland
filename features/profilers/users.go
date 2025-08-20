package main

// User is a simple struct that takes a bit of memory.
type User struct {
	ID    int
	Name  string
	Email string
	Data  []byte // simulate some extra payload (profile pic, etc.)
}

// MakeUsers creates 'n' users and keeps them in a slice.
// Each user allocates a ~1 KB byte slice for Data.
func MakeUsers(n int) []User {
	users := make([]User, 0, n)
	for i := 0; i < n; i++ {
		u := User{
			ID:    i,
			Name:  "User" + string(rune('A'+(i%26))),
			Email: "user@example.com",
			Data:  make([]byte, 1024), // ~1 KB
		}
		users = append(users, u)
	}
	return users
}
