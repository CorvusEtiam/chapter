package main

import "fmt"


func main() {
  var x,y,z,w,t uint32
  x = 123456789
  y = 362436069
  z = 521288629
  w = 88675123
  f := func() uint32 {
      t = x ^ (x << 11)
      x = y
      y = z
      z = w
      w = w ^ (w >> 19) ^ (t ^ (t >> 8))
      return w
}
  for i := 0; i < 10000000; i++ {
    fmt.Println(f())
  }
}


