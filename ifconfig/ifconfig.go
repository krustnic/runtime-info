package ifconfig

import (
    "fmt"
    "strings"
    
    "github.com/krustnic/runtime-info/utils"
)

type InterfaceInfo struct {
    Name    string
    Ip      string
    Netmask string
}

func (info *InterfaceInfo) buildFromString( line string ) {
    lines := strings.Split( line, "\n")        
    info.Name = strings.Split( lines[0], " " )[0]    
    fmt.Sscanf(lines[1], "          inet addr:%s  Bcast:0.0.0.0  Mask:%s", &info.Ip, &info.Netmask)    
}

func (info InterfaceInfo) String() string {
    return fmt.Sprintf( "Interface %s with ip:%s and mask:%s", info.Name, info.Ip, info.Netmask )
}

func Data() []InterfaceInfo {
    var interfaces []InterfaceInfo
    
    out := utils.Execute( "ifconfig", "-a" )
    outString := string(out)
    
    parts := strings.Split( outString, "\n\n")
    
    for _, part := range parts {
        if part != "" {
            var info InterfaceInfo
            info.buildFromString( part )
      
            interfaces = append( interfaces, info )       
        }        
    }    
    
    return interfaces
}