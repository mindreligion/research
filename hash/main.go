package main

import (
	"crypto/sha256"
	"fmt"
)

func main () {
	sumBytes := sha256.Sum256([]byte("50001180eyJTSUQiOiJjZWU2YjRhNC1mNWZiLTQ3NDYtOGVmYS1iNTBiNjg2NzlhMjMiLCJQSUQiOiI5OWNiN2M0ZS05NjBjLTQ1MjUtOGUzNC0wNjcwYTZmZWE5MTcifQ==sl389vls03hvna93283ngf983nskl983ng8skna9370pkhygtas43mldfyqzd"))
	fmt.Printf("%x", sumBytes)
}
