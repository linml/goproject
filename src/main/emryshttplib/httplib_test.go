package emryshttplib
import (
   "fmt"
   "io"
   "io/ioutil"
   )
func Testurl() {
 resp, err := Get("http://beego.me").Debug(true).Response()
 data ,err := ioutil.ReadAll(resp.Body)
 def resp.Body.Close()
 fmt.Print(data)
 
}