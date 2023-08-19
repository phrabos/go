package main

import(
  "fmt"
  "errors"
)

type TeaCategory struct {
  tea string
}
 
var (
  Green = TeaCategory{"Green"}
  Black = TeaCategory{"Black"}
  White = TeaCategory{"White"}
  Sheng = TeaCategory{"Sheng"}
  Oolong = TeaCategory{"Oolong"}
  Heicha = TeaCategory{"Heicha"}
  Unknown = TeaCategory{"Unknown"}
)

func validateTeaCategory(s string) (TeaCategory, error) {
  switch s {
  case Green.tea:
    return Green, nil
  case Black.tea:
    return Black, nil
  case White.tea:
    return White, nil
  case Sheng.tea:
    return Sheng, nil
  case Oolong.tea:
    return Oolong, nil
  case Heicha.tea:
    return Heicha, nil
  }
  return Unknown, errors.New("unknown tea type")
}


var teaCategories = []string{"Green", "Black", "White", "Sheng", "Oolong", "Heicha"}

type tea struct {
  name string
  weight int
  region string
}

func main() {
  fmt.Println("top of main")
  fmt.Printf("green %v\n", Green)
  fmt.Printf("black %v\n", Black)
  inventory := map[string]tea{}

  userInput1 := "Oolong"
  userInput2 := "Black"
  userInput3 := "pink"


  if _, err := validateTeaCategory(userInput1); err != nil {
    fmt.Println(err)
  }
  if _, err := validateTeaCategory(userInput2); err != nil {
    fmt.Println(err)
  }  
if _, err := validateTeaCategory(userInput3); err != nil {
    fmt.Println(err)
  }


  inventory[userInput1] = tea{name: "Ali Shan", weight: 150, region: "Taiwan"}
  inventory[userInput2] = tea{}
  inventory[userInput3] = tea{}

  fmt.Printf("inv %+v\n", inventory)
}
