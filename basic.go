package main

// to continue learning, go to: https://learnxinyminutes.com/docs/go/

import (
  "fmt"
  "os"
  m "math" // alias math to 'm'
)

/*
main is the entry point...
*/
func main() {

  // notice NO SEMICOLONS
  fmt.Println("Hello world!")


  beyond();
}


func beyond() {
  var x int
  x = 3

  y := 4 // type inference

  sum, prod := multiply(x, y)
  fmt.Println("sum: ", sum, "prod: ", prod)
  types()
}


/*
Here is how you would return multiple values...
notice the type hinting....
*/
func multiply(x, y int) (sum, prod int) {
  return x + y, x * y
}

func compute() float64{
  return m.Exp(10)
}



func types() {
  str := "Hello"

  s2 := `A "raw" string
  which
  may
  have
  line
  breaks`

  f := 3.1419 // float64
  c := 3 + 4i // complex number

  var u uint = 7 // unsigned
  var pi float32 = 22. / 7

  n := byte('\n')


  // Arrays are fixed size at compile time
  var a4 [4]int // init an array of size 4 of all {0, ...}

  a5 := [...]int{3, 10, 3} // fixed sz with set elems

  a4_cpy := a4 // shallow copy a4...
  a4_cpy[0] = 230

  fmt.Println(a4_cpy[0] == a4[0])



  // Slices have dynamic size at compile time!
  s3 := []int{4, 5, 9}
  s4 := make([]int, 4)  // make slice of sz 4 ints {0, ...}
  var d2 [][]float64    // declare

  bs := []byte("a slice...")

  s3_cpy := s3  // slices do NOT create separate instances
  s3_cpy[0] = 0
  fmt.Println(s3_cpy[0] == s3[0])

  // dyanmic means append able
  s := []int{1, 2, 3}
  s = append(s, 4, 5, 6)
  fmt.Println(s)

  s = append(s, []int{7, 8, 9}...) // unpack the slize and pass to append
  fmt.Println(s)

  p, q := learnMem()
  fmt.Println(*p, *q)

  // for hashing/dicts, use Map...
  m := map[string]int{"three": 3, "four": 4}
  m["one"] = 1
  // to ignore values, use '_'
  _, _, _, _, _, _, _, _, _, _, _ = str, s2, f, u, pi, n, a5, s4, bs, d2, c

  fp, _ := os.Create("output.txt")
  fmt.Fprint(fp, "Bam!")
  fp.Close()

  control()
}

func learnMem() (p, q *int) {
  p = new(int) // alloc an int

  s := make([]int, 20) // alloc 20 ints
  s[3] = 7
  r := -2
  return &s[3], &r
}

func control() {

  // parenth not required
  if true {
    fmt.Println("told ya")
  }

  if false {

  } else {

  }

  // switch preferred over chained if(s)
  x := 42.0
  switch x {
  case 0:
  case 1, 2:  // multiple
  case 42:    // NO FALL THROUGH

  case 43:
  default: //optional
  }

  for x:= 0; x < 3; x++ {
    fmt.Println("iteration ", x)
  }


  // a while true...
  for {
    break
    // continue will not reach
  }

  for key, val := range map[string]int{"one": 1, "two": 2, "three": 3} {
    // similar
    fmt.Printf("key=%s, value=%d\n", key, val)
  }

  for _, name := range []string{"Bob", "Bill", "Joe"} {
    fmt.Printf("Hello %s\n", name)
  }

  // L --> R
  if y := compute(); y > x {
    x = y
  }
  // function literal is a closure
  xBig := func() bool {
    return x > 10000
  }
  // referencing environment set on call, shallow binding...
  x = 99999
  fmt.Println("xBig:", xBig())
  x = 1.3e3
  fmt.Println("xBig:", xBig())

  // inline function declaration...
  fmt.Println("Add two numbers: ",
    func(a,b int) int {
      return a + b
    }(10, 2))

  goto love
  love:

  functionFactory()
  keyDefer()
  interfaces()
}

func functionFactory() {
  fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

  d := sentenceFactory("summer")
  fmt.Println(d("A beautiful", "day!"))
}

// easy function decorator
func sentenceFactory(str string) func(before, after string) string {
  return func(before, after string) string {
    return fmt.Sprintf("%s %s %s", before, str, after)
  }
}

func keyDefer() {
  // defer saves call until surrounding routine returns
  defer fmt.Println("deferred stm called in LIFO")
}

// ========= Interfaces ===========

func interfaces() {

}
