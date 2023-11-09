# GetNpaRulesRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Fields`                                                          | **string*                                                         | :heavy_minus_sign:                                                | Return values only from specified fields                          |
| `Filter`                                                          | **string*                                                         | :heavy_minus_sign:                                                | Query string based on query operaters                             |
| `Limit`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | Max number of policies to retrieve. Default will be all policies. |
| `Offset`                                                          | **int64*                                                          | :heavy_minus_sign:                                                | The offset of the first policy in the list to retrieve.           |
| `Sortby`                                                          | **string*                                                         | :heavy_minus_sign:                                                | Sort retrieved policies by specified field. Default is policy id  |
| `Sortorder`                                                       | **string*                                                         | :heavy_minus_sign:                                                | Sort in either asc or desc order. The default is asc order        |