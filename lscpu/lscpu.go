package lscpu

import (
    "strings"
    
    "github.com/krustnic/runtime-info/utils"
)

func Data() utils.StringsMap {
    out := utils.Execute( "lscpu", "" )
    outString := string(out)
    
    lines := strings.Split( outString, "\n")
    
    return utils.ParseTwoColumnsResponse( lines )
}