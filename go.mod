module github.com/delimatorres/go-splitwise

go 1.17

require github.com/joho/godotenv v1.4.0

require github.com/anvari1313/splitwise.go v0.0.0

// temporary replacement of splitwise.go
replace github.com/anvari1313/splitwise.go v0.0.0 => ./splitwise.go
