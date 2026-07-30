package HelperFunctions
import(
	"golang-loadbalancer/GlobalVariables"
)
func FindNextServerIdx() {
	LoadedHealthyServersList := LoadHealthyServersList()
	if LoadedHealthyServersList != nil{
		if len(LoadedHealthyServersList) > 0{
			var AtomicCurrentServerIdx = GlobalVariables.CurrentServerIdx.Load()
			var AtomicHealthyServersListEndIdx = uint32(len(LoadedHealthyServersList)-1)
			if AtomicCurrentServerIdx+1 > AtomicHealthyServersListEndIdx{
				GlobalVariables.CurrentServerIdx.Store(0)
			}else{
				GlobalVariables.CurrentServerIdx.Add(1)
			}
		}
	}
}
