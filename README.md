It's just an abstraction library over known implementations of SMS GW APIs

```go
package main

import (
        "github.com/xaionaro-go/sms/redsms.ru"
        "github.com/xaionaro-go/sms/smsc.ru"
)

func checkError(err error) {
        if err == nil {
                return
        }
        panic(err)
}

func main() {
        sms, err := redsmsRu.New("myLogin", "my API key")
        checkError(err)

        err = sms.Send("SomeSender", "79123456789", "Hey! :)")
        checkError(err)

        sms, err = smscRu.New("myLogin", "myPassword")
        checkError(err)

        err = sms.Send("SomeSender", "79123456789", "Hey, again!")
        checkError(err)
}
```
