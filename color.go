package main

import ("fmt"
  "bufio"
  "os"
  "strings"
)



type message map[string]string

const (
  _ACTION = "\033[32;40m"
  _WARN = "\033[31;40m"
  _PLACE = "\033[32;47m"
)



func main() {
  var in []string
  fmt.Println(_ACTION+bold("Who wanna bike"))
  in = getInput()
  for _, val := range in {
    fmt.Println(_WARN+bold(val))
  }
}
func printMessage(s string, mtype int) {
  switch(mtype){
  case 0:
    fmt.Println(_ACTION+s)
  default:
    fmt.Println(s)
  }
}

func getInput() []string {
  fmt.Println("Well what to do now?\n\n")
  agg := bufio.NewReader(os.Stdin)
  l,_ := agg.ReadString('\n')
  line := strings.Split(l, " ")
  return line
}

func bold(s string) string {
  var st string = "\033[1m"+s+"\033[0m"
  return st
}



