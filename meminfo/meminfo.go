package meminfo

import (
    "log"
    
    "github.com/krustnic/runtime-info/utils"
)

func getFileContent() []string {
    lines, err := utils.ReadFile( "/proc/meminfo" )
    
    if err != nil {
        log.Fatal( err )
        return nil
    }
    
    return lines
}

func Data() utils.StringsMap {
    lines := getFileContent()
    
    return utils.ParseTwoColumnsResponse( lines )
}