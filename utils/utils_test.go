package utils

import (
    "testing"
)

var linesPool = [][]string{
    { "abc: 1", "bca: 2" },
    { "a : 1", " b   :     2  " },
}

func TestGetMemInfoData( t *testing.T ) {
    size := 2
    
    for _, lines := range linesPool {
        info := ParseTwoColumnsResponse( lines )    
    
        if len(info) != size {
            t.Error("For lines", lines, "expected map with", size, "keys, but found", info )
        }   
    }    
}

func TestReadFile( t *testing.T ) {
    out, err := ReadFile( "/etc/passwd" )
    
    if err != nil {
        t.Error( "Cannot read file /etc/passwd" )
        return
    }
    
    if len(out) == 0 {
        t.Error( "Wrong file content" )
    }     
}