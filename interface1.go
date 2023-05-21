package main

// importing package 'fmt'
import "fmt"

// creating an interface 'Language'
type Language interface {
	getDevelopedBy() string
}

// creating a structure 'JavaScript'
// that contains the fields required by interface 'Language'
type JavaScript struct {
	DevelopedBy string
}

// implementing methods of interface 'Language'
// for structure 'JavaScript'
func (javaScript JavaScript) getDevelopedBy() string {
	return javaScript.DevelopedBy
}

// creating a structure 'Python'
// that contains the fields required by interface 'Language'
type Python struct {
	DevelopedBy string
}

// implementing methods of interface 'Language'
// for structure 'Python'
func (python Python) getDevelopedBy() string {
	return python.DevelopedBy
}

func mainI2() {
	// creating an instance of interface 'Language'
	var ILanguage Language

	// creating an object of structure 'JavaScript'
	javaScript := JavaScript{DevelopedBy: "Brendan Eich"}

	// creating an object of structure 'Python'
	python := Python{DevelopedBy: "Guido van Rossum"}

	// assigning object 'javaScript' to 'ILanguage'
	// and invoking getDevelopedBy()
	ILanguage = javaScript
	fmt.Println(ILanguage.getDevelopedBy())

	// assigning object 'python' to 'ILanguage'
	// and invoking getDevelopedBy()
	ILanguage = python
	fmt.Println(ILanguage.getDevelopedBy())
}
