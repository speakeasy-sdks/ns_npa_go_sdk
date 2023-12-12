<!-- Start SDK Example Usage [usage] -->
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
		nsnpagosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.DeleteNpaRulesID(ctx, operations.DeleteNpaRulesIDRequest{
		ID: 324988,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->