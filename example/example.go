package main

import (
	"fmt"
	// "time"

	sdk "github.com/pioneer-io/go_sdk"
	// "github.com/pioneer-io/go_sdk/pkg/models"
	// "gopkg.in/segmentio/analytics-go.v3"
)

/*
run a simple program like this to test out the SDK
*/

func main() {
	// Initialize an SDK client
	// client := sdk.InitMember("http://localhost:3030/features", "JazzyElksRule")

	// connect SDK client to Scout to listen for SSE updates
	// client.Connect()
	// client.Listen()


	analytics := sdk.InitAnalytics("UA-XXXXXX-X")
	// testFlag1 := &models.FlagData{Is_Active: true, Rollout: 50}
	// ruleset := make(map[string]*models.FlagData)
	// ruleset["test_flag_1"] = testFlag1
	fmt.Println(analytics.LogAnalyticsEvent("pioneer", "log", "1"))

	// this example supposes that you have an existing flag
	// called 'test this flag' that is toggled off

	// fmt.Println(client.Get("test this flag")) // false

	// time.Sleep(12 * time.Second)
	// // wait so you can toggle "test this flag" von ia UI
	// // and create "a_new_flag" and toggle it on

	// // after we're finished sleeping, we can log the updated ruleset
	// fmt.Println(client.Get("a_new_flag"))     // true
	// fmt.Println(client.Get("test this flag")) // true

}


// package main

// import (
//         "errors"
//         "fmt"
//         "log"
//         "net"
//         "net/http"
//         "net/url"
//         "os"

//         uuid "github.com/gofrs/uuid"

//         "google.golang.org/appengine"
// )

// var gaPropertyID = mustGetenv("GA_TRACKING_ID")

// func mustGetenv(k string) string {
//         v := os.Getenv(k)
//         if v == "" {
//                 log.Fatalf("%s environment variable not set.", k)
//         }
//         return v
// }

// func main() {
//         http.HandleFunc("/", handle)

//         appengine.Main()
// }

// func handle(w http.ResponseWriter, r *http.Request) {
//         if r.URL.Path != "/" {
//                 http.NotFound(w, r)
//                 return
//         }

//         if err := trackEvent(r, "Example", "Test action", "label", nil); err != nil {
//                 fmt.Fprintf(w, "Event did not track: %v", err)
//                 return
//         }
//         fmt.Fprint(w, "Event tracked.")
// }

// func trackEvent(r *http.Request, category, action, label string, value *uint) error {
//         if gaPropertyID == "" {
//                 return errors.New("analytics: GA_TRACKING_ID environment variable is missing")
//         }
//         if category == "" || action == "" {
//                 return errors.New("analytics: category and action are required")
//         }

//         v := url.Values{
//                 "v":   {"1"},
//                 "tid": {gaPropertyID},
//                 // Anonymously identifies a particular user. See the parameter guide for
//                 // details:
//                 // https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#cid
//                 //
//                 // Depending on your application, this might want to be associated with the
//                 // user in a cookie.
//                 "cid": {uuid.Must(uuid.NewV4()).String()},
//                 "t":   {"event"},
//                 "ec":  {category},
//                 "ea":  {action},
//                 "ua":  {r.UserAgent()},
//         }

//         if label != "" {
//                 v.Set("el", label)
//         }

//         if value != nil {
//                 v.Set("ev", fmt.Sprintf("%d", *value))
//         }

//         if remoteIP, _, err := net.SplitHostPort(r.RemoteAddr); err != nil {
//                 v.Set("uip", remoteIP)
//         }

//         // NOTE: Google Analytics returns a 200, even if the request is malformed.
//         _, err := http.PostForm("https://www.google-analytics.com/collect", v)
//         return err
// }