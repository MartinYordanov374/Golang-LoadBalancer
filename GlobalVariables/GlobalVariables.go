package GlobalVariables
import(
	"sync/atomic"
	"time"
	"net/http"
)
var HealthyServersList atomic.Value
var CurrentServerIdx atomic.Uint32
var HttpClient = &http.Client{
	Timeout: 5*time.Second,
}
