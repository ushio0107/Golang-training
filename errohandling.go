package main

import (
	"errors"
	"fmt"
)

type ErrUserNameExist struct {
	UserName string
}

func (e ErrUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}

func IsErrUserNameExist(err error) bool {
	_, ok := err.(ErrUserNameExist)
	return ok
}

func checkUserNameExist(username string) (bool, error) {
	if username == "foo" { //method 1
		return true, fmt.Errorf("username %s already exist", username) // string and both variable
	}

	if username == "bar" {
		return true, errors.New("username bar already exist") // only input string
	}

	return false, nil
}

func checkUserNameExistOth(username string) (bool, error) {
	if username == "check3" { //method 1
		return true, ErrUserNameExist{UserName: username}
	}

	return false, nil
}

func errorhandling() { //main
	// err data type is Error()
	if _, err := checkUserNameExist("foo"); err != nil { // underscore means to ignore the return value
		fmt.Println(err)
	}

	if _, err := checkUserNameExist("bar"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("")

	if _, err := checkUserNameExistOth("foo"); err != nil {
		fmt.Println(err)
	}

	if _, err := checkUserNameExistOth("check3"); err != nil {
		if IsErrUserNameExist(err) {
			fmt.Println(err)
		}
	}

}
