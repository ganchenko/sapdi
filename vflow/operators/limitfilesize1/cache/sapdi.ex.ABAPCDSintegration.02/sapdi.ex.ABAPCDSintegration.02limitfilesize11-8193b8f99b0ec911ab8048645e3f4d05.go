package main

import (
	"strconv"
)

var mysize=0
var mycounter=0
var mykblimit=1024

var GetInt func(string) int
var Out func(interface{})

func main() {

}


func In(msg interface{}) {

    if msg == nil {
      return
    }

    mykblimit = GetInt("maxsizekb")
    
    if mykblimit == 0 { mykblimit = 1024 }

    mysize += len(msg.(map[string]interface{})["Body"].(string))
    if mysize >= 1024 * mykblimit {
      mycounter += 1
      mysize = len(msg.(map[string]interface{})["Body"].(string))
    }
     
    msg.(map[string]interface{})["Attributes"].(map[string]interface{})["ABAPfilenumber"] = strconv.Itoa(mycounter)

	Out(msg)
    
}
