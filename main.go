package main

import "fmt"

const (
  W = 40
  H = 20
)

type Rectangle struct {
  width int
  height int
  xpos int
  ypos int
}

func main() {
  var t [H][W]string
  r := Rectangle{5,4,1,1}
  t = Fill(r, t)
  for _, val := range t {
    for _, x := range val {
      if x == "#" {
      fmt.Print(x)
      } else {
        fmt.Print(".")
      }
    }
    fmt.Println()
  }
}
/*
func GetRandom() uint32 {
  var x uint32 = 123456789
  var y uint32 = 362436069
  var z uint32 = 521288629
  var w uint32 = 88675123
  var t uint32
  t = x ^ (x << 11)
  x = y
  y = z
  z = w
  w = w ^ (w >> 19) ^ (t ^ (t >> 8))
  return w
}
*/

func GetRandom(int seed)


func Fill(elem Rectangle, arr[H][W]string) [H][W]string {
  for i:=elem.xpos; i <= elem.width; i++ {
    arr[elem.ypos][i] = "#"
    arr[elem.ypos+elem.width-1][i] = "#"
  }
  for j := elem.ypos; j <= elem.height; j++ {
    arr[j][elem.xpos] = "#"
    arr[j][elem.xpos+elem.width-1] = "#"
  }
  return arr
}
