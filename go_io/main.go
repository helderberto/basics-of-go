package main

var url = "https://helderberto.com"

// it's not possible as a global variable
// message := "Hello World!"

func main() {
  message := "Hello World!"
  price := 34.4

  print(message, price, url)
}
