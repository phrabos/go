package main

import(
  "fmt"
  // "errors"
)

// type TeaCategory struct {
//   tea string
// }
 
// var (
//   Green = TeaCategory{"Green"}
//   Black = TeaCategory{"Black"}
//   White = TeaCategory{"White"}
//   Sheng = TeaCategory{"Sheng"}
//   Oolong = TeaCategory{"Oolong"}
//   Heicha = TeaCategory{"Heicha"}
//   Unknown = TeaCategory{"Unknown"}
// )

// func validateTeaCategory(s string) (TeaCategory, error) {
//   switch s {
//   case Green.tea:
//     return Green, nil
//   case Black.tea:
//     return Black, nil
//   case White.tea:
//     return White, nil
//   case Sheng.tea:
//     return Sheng, nil
//   case Oolong.tea:
//     return Oolong, nil
//   case Heicha.tea:
//     return Heicha, nil
//   }
//   return Unknown, errors.New("unknown tea type")
// }

const (
  green TeaCategory = iota
  black
  white
  sheng
  oolong
  hiecha
)

type TeaCategory int
var teaCategoryStrings = []string{"green", "black", "white", "sheng", "oolong", "heicha"}

type tea struct {
  name string
  weight int
  region string
}
func (tc TeaCategory) name() string {
  return teaCategoryStrings[tc]
}

type Enum interface {
  name() string
}

func main() {
  fmt.Println("top of main")
  fmt.Printf("green %v\n", green)
  fmt.Printf("black %v\n", black)
  inventory := map[TeaCategory]tea{}

  userInput1 := "oolong"
  var userInputValidated TeaCategory = indexOf(userInput1, teaCategoryStrings)
  userInput2 := "black"
  userInput3 := "pink"
  var userInput3Validated TeaCategory = indexOf(userInput3, teaCategoryStrings)


//   if _, err := validateTeaCategory(userInput1); err != nil {
//     fmt.Println(err)
//   }
//   if _, err := validateTeaCategory(userInput2); err != nil {
//     fmt.Println(err)
//   }  
// if _, err := validateTeaCategory(userInput3); err != nil {
//     fmt.Println(err)
//   }


  inventory[userInputValidated] = tea{name: "Ali Shan", weight: 150, region: "Taiwan"}
  inventory[indexOf(userInput2, teaCategoryStrings)] = tea{name: "Jin Jun Mei", weight: 50, region: "China" }
  inventory[userInput3Validated] = tea{}

  fmt.Printf("inv %+v\n", inventory)
}

func indexOf(needle string, haystack []string) int {
  for i, val := range haystack {
    if val == needle {
      return i
    }
  }
  return -1
}