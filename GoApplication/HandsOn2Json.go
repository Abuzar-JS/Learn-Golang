package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type person struct {
	First  string   `json:"First"`
	Last   string   `json:"Last"`
	Age    int      `json:"Age"`
	Saying []string `json:"Sayings"`
}

// Hands2 is exported to main.go
func Hands2() {

	// s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	file, err := os.Open("handsEX2.json")
	if err != nil {
		return
	}

	s, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var fil []person
	err = json.Unmarshal(s, &fil)
	if err != nil {
		return
	}

	fmt.Println(fil)

	for _, v := range fil {
		fmt.Println(v.Age)
	}
	defer file.Close()
}
