//GOLANG PROGRAM TO PRINT REVERSE PYRAMID STAR PATTERN
package main

// fmt package provides the function to print anything
import "fmt"
//this is the main function
func main() {
  
   //initialize the number of rows to print
   var rows = 5
   
   //declaring variables with integer datatype
   fmt.Println("\nThis is the Reverse pyramid pattern")
   var i, j int
   
   //displaying the pattern
   for i = rows; i >= 1; i-- {
      
      //printing the spaces
      for space := 1; space <= rows-i; space++ {
         fmt.Print(" ")
      }
      
      //printing the stars
      for j = i; j <= 2*i-1; j++ {
         fmt.Printf("*")
      }
      for j = 0; j < i-1; j++ {
         fmt.Printf("*")
      }
      fmt.Println("")
   }
}