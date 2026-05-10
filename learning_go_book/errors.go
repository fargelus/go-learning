package main

import (
	"errors"
	"fmt"
)

/*
1. Create a sentinel error to represent an invalid ID.
   In main, use errors.Is to check for the sentinel error, and print a message when it is found.

2. Define a custom error type to represent an empty field error.
   This error should include the name of the empty Employee field.
   In main, use errors.As to check for this error.
   Print out a message that includes the field name.

3. Rather than returning the first error found, return back a single error that contains all errors discovered during validation.
   Update the code in main to properly report multiple errors.
*/

var ErrInvalidID = errors.New("invalid ID")

type Employee struct {
	Id       int
	Name     string
	LastName string
}

type EmptyEmployeeFieldError struct {
	Field string
}

func (e EmptyEmployeeFieldError) Error() string {
	return fmt.Sprintf("employee field %q is empty", e.Field)
}

func validateID(id int) error {
	if id <= 0 {
		return ErrInvalidID
	}
	return nil
}

func validateEmployee(e Employee) error {
	var errs []error

	err := validateID(e.Id)
	if errors.Is(err, ErrInvalidID) {
		errs = append(errs, err)
	}

	if e.Name == "" {
		errs = append(errs, EmptyEmployeeFieldError{Field: "Name"})
	}

	if e.LastName == "" {
		errs = append(errs, EmptyEmployeeFieldError{Field: "LastName"})
	}

	return errors.Join(errs...)
}

func printErrors(wrappedErrors error) {
	joinedErr := wrappedErrors.(interface{ Unwrap() []error })

	fmt.Printf("- ")
	for i, e := range joinedErr.Unwrap() {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%v", e)
	}
	fmt.Println("")
}

func main() {
	errs := validateEmployee(Employee{Id: 1, Name: "John", LastName: ""})
	if errs != nil {
		printErrors(errs)
	}

	errs = validateEmployee(Employee{Id: -1, Name: "", LastName: "Doe"})
	if errs != nil {
		printErrors(errs)
	}

	errs = validateEmployee(Employee{Id: 2, Name: "Sam", LastName: "Jones"})
	if errs != nil {
		printErrors(errs)
	}
}
