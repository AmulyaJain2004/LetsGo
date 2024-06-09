package main
import "fmt"

const LoginToken string = "ghabbhhjd" // Public as first letter is Capital
// we can't use := in global

func main ()  {
    var username string = "Amulya"
    fmt.Println(username)
    fmt.Printf("Variable is of type: %T\n", username)

    var isLoggedIn bool = true
    fmt.Println(isLoggedIn)
    fmt.Printf("Variable is of type: %T\n", isLoggedIn)

    var smallVal uint8 = 255 // it will give error in 256 as it is out of bound // we can also use int which is better
    fmt.Println(smallVal)
    fmt.Printf("Variable is of type: %T\n", smallVal)

    var smallfloat float32 = 255.45556756
    fmt.Println(smallfloat)
    fmt.Printf("Variable is of type: %T\n", smallfloat)

    var bigfloat float64 = 255.4555675654645
    fmt.Println(bigfloat)
    fmt.Printf("Variable is of type: %T\n", bigfloat)

    // default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)  // automatically initializes it to 0 and no garbage value 
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

    var anotherVar string
	fmt.Println(anotherVar) // automatically initializes it to blank and no garbage value 
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type
	var website = "www.google.com" // lexer automatically assigns string type to it 
	fmt.Println(website)
    // website = 3 as cannot use 3 (untyped int constant) as string value in  

	// no var style
	numberOfUser := 300000.0 // := colon and equal which is walrus operator 
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}