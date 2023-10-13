<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"log"
)

func main() {
	s := nsnpagosdk.New(
		nsnpagosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Platform.DeleteNpaRulesID(ctx, operations.DeleteNpaRulesIDRequest{
		ID: 324988,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.DeleteNpaRulesID200ApplicationJSONObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->