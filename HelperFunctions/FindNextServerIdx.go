package HelperFunctions
import(
	"golang-loadbalancer/GlobalVariables"
	"sync/atomic"
)
func FindNextServerIdx() {
	LoadedHealthyServersList := LoadHealthyServersList()
	if LoadedHealthyServersList != nil{
		if len(LoadedHealthyServersList) > 0{
			var AtomicCurrentServerIdx = atomic.LoadUint32(&GlobalVariables.CurrentServerIdx)
			var AtomicHealthyServersListEndIdx = uint32(len(LoadedHealthyServersList)-1)
			if AtomicCurrentServerIdx+1 > AtomicHealthyServersListEndIdx{
				atomic.StoreUint32(&GlobalVariables.CurrentServerIdx, 0)
			}else{
				atomic.AddUint32(&GlobalVariables.CurrentServerIdx, 1)
			}
		}
	}
}
