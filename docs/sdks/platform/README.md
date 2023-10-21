# Platform SDK


## Overview

npa_policy: NPA policy CRUD operations.

### Available Operations

* [DeleteNpaRulesID](#deletenparulesid) - Delete a npa policy
* [GetNpaRules](#getnparules) - Get list of npa policies
* [GetNpaRulesID](#getnparulesid) - Get a npa policy
* [PatchNpaRulesID](#patchnparulesid) - Patch a npa policy
* [PostNpaRules](#postnparules) - Create a npa policy

## DeleteNpaRulesID

Delete a npa policy with rule id

### Example Usage

```go
package main

import(
	"context"
	"log"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteNpaRulesIDRequest](../../models/operations/deletenparulesidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteNpaRulesIDResponse](../../models/operations/deletenparulesidresponse.md), error**


## GetNpaRules

Get list of npa policies

### Example Usage

```go
package main

import(
	"context"
	"log"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Platform.GetNpaRules(ctx, operations.GetNpaRulesRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.NpaPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetNpaRulesRequest](../../models/operations/getnparulesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetNpaRulesResponse](../../models/operations/getnparulesresponse.md), error**


## GetNpaRulesID

Get a npa policy based on policy rule id

### Example Usage

```go
package main

import(
	"context"
	"log"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Platform.GetNpaRulesID(ctx, operations.GetNpaRulesIDRequest{
        ID: 408556,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNpaRulesID200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetNpaRulesIDRequest](../../models/operations/getnparulesidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetNpaRulesIDResponse](../../models/operations/getnparulesidresponse.md), error**


## PatchNpaRulesID

Patch a npa policy based on rule id

### Example Usage

```go
package main

import(
	"context"
	"log"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Platform.PatchNpaRulesID(ctx, operations.PatchNpaRulesIDRequest{
        ID: 348436,
        NpaPolicyRequest: shared.NpaPolicyRequest{
            Description: nsnpagosdk.String("any"),
            Enabled: nsnpagosdk.String("1"),
            GroupID: nsnpagosdk.String("1"),
            GroupName: nsnpagosdk.String("My policy group"),
            RuleData: &shared.NpaPolicyRuleData{
                DlpActions: []shared.NpaPolicyRuleDlp{
                    shared.NpaPolicyRuleDlp{
                        Actions: []shared.NpaPolicyRuleDlpActions{
                            shared.NpaPolicyRuleDlpActionsAllow,
                        },
                        DlpProfile: nsnpagosdk.String("Payment Card"),
                    },
                },
                JSONVersion: nsnpagosdk.Int64(3),
                MatchCriteriaAction: &shared.NpaPolicyRuleDataMatchCriteriaAction{},
                NetLocationObj: []string{
                    "190.123.150.10",
                    "190.218.0.0/16",
                },
                OrganizationUnits: []string{
                    "engineering/qa",
                },
                PrivateAppIds: []string{
                    "100",
                    "201",
                },
                PrivateAppTagIds: []string{
                    "1",
                    "2",
                },
                PrivateAppTags: []string{
                    "tag1",
                    "tag2",
                },
                PrivateApps: []string{
                    "app1",
                    "app2",
                },
                PrivateAppsWithActivities: []shared.NpaPolicyRuleDataPrivateAppsWithActivities{
                    shared.NpaPolicyRuleDataPrivateAppsWithActivities{
                        Activities: []shared.NpaPolicyRuleDataPrivateAppsWithActivitiesActivities{
                            shared.NpaPolicyRuleDataPrivateAppsWithActivitiesActivities{
                                ListOfConstraints: []string{
                                    "string",
                                },
                            },
                        },
                        AppName: nsnpagosdk.String("[172.31.12.135]"),
                    },
                },
                SrcCountries: []string{
                    "US",
                    "AF",
                    "CN",
                },
                UserGroups: []string{
                    "usergroup/group1",
                },
                Users: []string{
                    "vphan@netskope.com",
                },
                Version: nsnpagosdk.Int64(1),
            },
            RuleName: nsnpagosdk.String("vantest"),
            RuleOrder: &shared.NpaPolicyRequestRuleOrder{
                Position: nsnpagosdk.Int64(5),
                RuleID: nsnpagosdk.Int64(1),
                RuleName: nsnpagosdk.String("api-policy-managed"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PatchNpaRulesID200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchNpaRulesIDRequest](../../models/operations/patchnparulesidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PatchNpaRulesIDResponse](../../models/operations/patchnparulesidresponse.md), error**


## PostNpaRules

Create a policy

### Example Usage

```go
package main

import(
	"context"
	"log"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Platform.PostNpaRules(ctx, operations.PostNpaRulesRequest{
        NpaPolicyRequest: shared.NpaPolicyRequest{
            Description: nsnpagosdk.String("any"),
            Enabled: nsnpagosdk.String("1"),
            GroupID: nsnpagosdk.String("1"),
            GroupName: nsnpagosdk.String("My policy group"),
            RuleData: &shared.NpaPolicyRuleData{
                DlpActions: []shared.NpaPolicyRuleDlp{
                    shared.NpaPolicyRuleDlp{
                        Actions: []shared.NpaPolicyRuleDlpActions{
                            shared.NpaPolicyRuleDlpActionsAllow,
                        },
                        DlpProfile: nsnpagosdk.String("Payment Card"),
                    },
                },
                JSONVersion: nsnpagosdk.Int64(3),
                MatchCriteriaAction: &shared.NpaPolicyRuleDataMatchCriteriaAction{},
                NetLocationObj: []string{
                    "190.123.150.10",
                    "190.218.0.0/16",
                },
                OrganizationUnits: []string{
                    "engineering/qa",
                },
                PrivateAppIds: []string{
                    "100",
                    "201",
                },
                PrivateAppTagIds: []string{
                    "1",
                    "2",
                },
                PrivateAppTags: []string{
                    "tag1",
                    "tag2",
                },
                PrivateApps: []string{
                    "app1",
                    "app2",
                },
                PrivateAppsWithActivities: []shared.NpaPolicyRuleDataPrivateAppsWithActivities{
                    shared.NpaPolicyRuleDataPrivateAppsWithActivities{
                        Activities: []shared.NpaPolicyRuleDataPrivateAppsWithActivitiesActivities{
                            shared.NpaPolicyRuleDataPrivateAppsWithActivitiesActivities{
                                ListOfConstraints: []string{
                                    "string",
                                },
                            },
                        },
                        AppName: nsnpagosdk.String("[172.31.12.135]"),
                    },
                },
                SrcCountries: []string{
                    "US",
                    "AF",
                    "CN",
                },
                UserGroups: []string{
                    "usergroup/group1",
                },
                Users: []string{
                    "vphan@netskope.com",
                },
                Version: nsnpagosdk.Int64(1),
            },
            RuleName: nsnpagosdk.String("vantest"),
            RuleOrder: &shared.NpaPolicyRequestRuleOrder{
                Position: nsnpagosdk.Int64(5),
                RuleID: nsnpagosdk.Int64(1),
                RuleName: nsnpagosdk.String("api-policy-managed"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NpaPolicyResponseItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PostNpaRulesRequest](../../models/operations/postnparulesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PostNpaRulesResponse](../../models/operations/postnparulesresponse.md), error**

