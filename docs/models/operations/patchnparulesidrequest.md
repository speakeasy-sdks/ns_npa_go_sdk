# PatchNpaRulesIDRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ID`                                                               | *int*                                                              | :heavy_check_mark:                                                 | policy rule id                                                     |
| `NpaPolicyRequest`                                                 | [shared.NpaPolicyRequest](../../models/shared/npapolicyrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `Silent`                                                           | [*operations.Silent](../../models/operations/silent.md)            | :heavy_minus_sign:                                                 | flag to skip output except status code                             |