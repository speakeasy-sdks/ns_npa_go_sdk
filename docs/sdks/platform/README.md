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
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"context"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
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

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteNpaRulesIDRequest](../../pkg/models/operations/deletenparulesidrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteNpaRulesIDResponse](../../pkg/models/operations/deletenparulesidresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NpaPolicyResponse400 | 400                            | application/json               |
| sdkerrors.SDKError             | 400-600                        | */*                            |

## GetNpaRules

Get list of npa policies

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"context"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
	"log"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetNpaRules(ctx, operations.GetNpaRulesRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.NpaPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetNpaRulesRequest](../../pkg/models/operations/getnparulesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetNpaRulesResponse](../../pkg/models/operations/getnparulesresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NpaPolicyResponse400 | 400                            | application/json               |
| sdkerrors.SDKError             | 400-600                        | */*                            |

## GetNpaRulesID

Get a npa policy based on policy rule id

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"context"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
	"log"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetNpaRulesID(ctx, operations.GetNpaRulesIDRequest{
        ID: 408556,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetNpaRulesIDRequest](../../pkg/models/operations/getnparulesidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetNpaRulesIDResponse](../../pkg/models/operations/getnparulesidresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NpaPolicyResponse400 | 400                            | application/json               |
| sdkerrors.SDKError             | 400-600                        | */*                            |

## PatchNpaRulesID

Patch a npa policy based on rule id

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"context"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
	"log"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PatchNpaRulesID(ctx, operations.PatchNpaRulesIDRequest{
        ID: 348436,
        NpaPolicyRequest: shared.NpaPolicyRequest{
            Description: nsnpagosdk.String("any"),
            Enabled: nsnpagosdk.String("1"),
            GroupID: nsnpagosdk.String("1"),
            GroupName: nsnpagosdk.String("My policy group"),
            RuleData: &shared.NpaPolicyRuleData{
                DlpActions: []shared.NpaPolicyRuleDlp{
                    shared.NpaPolicyRuleDlp{
                        Actions: []shared.Actions{
                            shared.ActionsAllow,
                        },
                        DlpProfile: nsnpagosdk.String("Payment Card"),
                    },
                },
                JSONVersion: nsnpagosdk.Int64(3),
                MatchCriteriaAction: &shared.MatchCriteriaAction{},
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
                PrivateAppsWithActivities: []shared.PrivateAppsWithActivities{
                    shared.PrivateAppsWithActivities{
                        Activities: []shared.Activities{
                            shared.Activities{
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
            RuleOrder: &shared.RuleOrder{
                Position: nsnpagosdk.Int64(5),
                RuleID: nsnpagosdk.Int64(1),
                RuleName: nsnpagosdk.String("api-policy-managed"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchNpaRulesIDRequest](../../pkg/models/operations/patchnparulesidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PatchNpaRulesIDResponse](../../pkg/models/operations/patchnparulesidresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NpaPolicyResponse400 | 400                            | application/json               |
| sdkerrors.SDKError             | 400-600                        | */*                            |

## PostNpaRules

Create a policy

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	nsnpagosdk "github.com/speakeasy-sdks/ns_npa_go_sdk"
	"context"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/operations"
	"log"
)

func main() {
    s := nsnpagosdk.New(
        nsnpagosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.PostNpaRules(ctx, operations.PostNpaRulesRequest{
        NpaPolicyRequest: shared.NpaPolicyRequest{
            Description: nsnpagosdk.String("any"),
            Enabled: nsnpagosdk.String("1"),
            GroupID: nsnpagosdk.String("1"),
            GroupName: nsnpagosdk.String("My policy group"),
            RuleData: &shared.NpaPolicyRuleData{
                DlpActions: []shared.NpaPolicyRuleDlp{
                    shared.NpaPolicyRuleDlp{
                        Actions: []shared.Actions{
                            shared.ActionsAllow,
                        },
                        DlpProfile: nsnpagosdk.String("Payment Card"),
                    },
                },
                JSONVersion: nsnpagosdk.Int64(3),
                MatchCriteriaAction: &shared.MatchCriteriaAction{},
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
                PrivateAppsWithActivities: []shared.PrivateAppsWithActivities{
                    shared.PrivateAppsWithActivities{
                        Activities: []shared.Activities{
                            shared.Activities{
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
            RuleOrder: &shared.RuleOrder{
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostNpaRulesRequest](../../pkg/models/operations/postnparulesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PostNpaRulesResponse](../../pkg/models/operations/postnparulesresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NpaPolicyResponse400 | 400                            | application/json               |
| sdkerrors.SDKError             | 400-600                        | */*                            |
