package GlobalVariables
import(
	"time"
	"net/http"
)
var HttpClient = &http.Client{
	Timeout: 5*time.Second,
}
