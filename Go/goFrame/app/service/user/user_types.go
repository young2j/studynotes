package user

// SignUpParam service层入参
type SignUpParam struct {
    Passport  string
    Password  string
    Password2 string
    Nickname  string
}