package GlobalVariables
import(
	"sync/atomic"
	"time"
	"net/http"
)
var HealthyServersList atomic.Value
var CurrentServerIdx uint32 = 0 
var HttpClient = &http.Client{
	Timeout: 5*time.Second,
}
