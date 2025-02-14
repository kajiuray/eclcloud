package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/security_portal/v2/ports"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestUpdatePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/ecl-api/ports/utm/CES11995"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, updateResponse)
	})

	updateOpts := ports.UpdateOpts{
		Port: []ports.SinglePort{
			{
				Comment:    "port 0 comment",
				EnablePort: "true",
				IPAddress:  "192.168.1.50/24",
				MTU:        "1500",
				NetworkID:  "32314bd2-3583-4fb9-b622-9b121e04e007",
				SubnetID:   "7fd77711-abae-4828-93f1-f3d682a8771f",
			},
		},
	}

	actual, err := ports.Update(
		fakeclient.ServiceClient(),
		"utm",
		"CES11995",
		updateOpts,
		nil).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &expectedResult, actual)
}
