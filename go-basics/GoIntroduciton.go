package go_basics

import "fmt"

/*
* Go is an open source programming language designed for  building scalable, secure and reliable software.
* This repository is built by taking reference from gobyexample.com
* In this program, we will look into building a hello world example.


<!------- CREATING MAIN FUNCTION -------!>
* When ever a Go file is created inside a module, the package is automatically taken as the module name.
* To run a Go file, it should have main method. A Go file with main method is only executable when the package is declared
	as main instead of package name.
* In order to execute a Go file, we need to have a main function.
* Any function in GO follows a structure:
		func functionName ( argName argType){
			// function Body or logic that should be executed when function is called.
		}
* In the above syntax func is a keyword registerd in GO which tells the GO compiler, it is a function.
*/

func main() {
	fmt.Println("Hello World")
}
