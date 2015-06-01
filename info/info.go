package info

import (
    "fmt" 
    "os"
    "runtime"
    "encoding/json"
    
    "github.com/krustnic/runtime-info/meminfo"
    "github.com/krustnic/runtime-info/ifconfig"
    "github.com/krustnic/runtime-info/lscpu"
)

type RuntimeInfo map[string]interface{}

func (info RuntimeInfo) ToJSON() string {
    bytes, _ := json.MarshalIndent( info, "", "    " )
    return string(bytes)
}

func init() {
    if runtime.GOOS == "windows" {
        fmt.Println("Error. RuntimeInfo not working with non-unix OS")
        os.Exit(1)
    }
}

func ToJSON() string {
    meminfoData  := meminfo.Data()  
    ifconfigData := ifconfig.Data()
    lscpuData    := lscpu.Data()
    
    _ = meminfoData
    _ = ifconfigData
    
    runtimeInfo := make( RuntimeInfo )    
    
    runtimeInfo["meminfo"]  = meminfoData
    runtimeInfo["ifconfig"] = ifconfigData
    runtimeInfo["lscpu"]    = lscpuData
    
    return runtimeInfo.ToJSON()
}