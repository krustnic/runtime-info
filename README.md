# Runtime Info

Golang library that provides server runtime information about network, memory and cpu for UNIX systems. Information comes from the following sources:
* /proc/meminfo
* ifconfig
* lscpu

# Example

```go
package main

import (        
    "fmt"
    
    "github.com/krustnic/runtime-info/info"
)

func main() {
    out := info.ToJSON()    
    fmt.Println( out )
}
```