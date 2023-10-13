# github.com/speakeasy-sdks/ns_npa_go_sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/ns_npa_go_sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/ns_npa_go_sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
# SDK Installation

```bash
go get github.com/speakeasy-sdks/ns_npa_go_sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
# Available Resources and Operations

## [Platform SDK](docs/sdks/platform/README.md)

* [DeleteNpaRulesID](docs/sdks/platform/README.md#deletenparulesid) - Delete a npa policy
* [GetNpaRules](docs/sdks/platform/README.md#getnparules) - Get list of npa policies
* [GetNpaRulesID](docs/sdks/platform/README.md#getnparulesid) - Get a npa policy
* [PatchNpaRulesID](docs/sdks/platform/README.md#patchnparulesid) - Patch a npa policy
* [PostNpaRules](docs/sdks/platform/README.md#postnparules) - Create a npa policy
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->



<!-- End Dev Containers -->

<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
