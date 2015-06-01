package utils

import (
    "os"    
    "os/exec"
    "bufio"
    "io"
    "log"
    "strings"
)

type StringsMap map[string]string

func Execute( cmd string, arg string ) []byte {
    out, err := exec.Command( cmd, arg ).Output()    
    
    if err != nil {
        log.Fatal( err )
    }
    
    return out
}

func ParseTwoColumnsResponse( lines []string ) StringsMap {
    info := make( StringsMap )    
    
    for _, line := range lines {
        parts := strings.Split( line, ":" )
        
        if len(parts) != 2 {
            return info        
        }
                
        name  := parts[0]
        value := strings.Trim( parts[1], " \n" )
        //valueWithKB := strings.Trim( parts[1], " \n" )
        //value := strings.Split( valueWithKB, " " )[0]

        info[name] = value
    }
    
    return info
}

func ReadFile( fileName string ) ([]string, error) {
    file, err := os.Open( fileName )
    
    if err != nil {        
        return nil, err
    }
    
    fileReader := bufio.NewReader( file )
    
    isEOF := false
    
    var lines []string
    
    for !isEOF {        
        line, err := fileReader.ReadString( '\n' )
        
        if err == io.EOF {
           isEOF = true
        } else if err != nil {
            return nil, err
        } else {
            lines = append( lines, line )    
        }        
    }
    
    return lines, nil
}