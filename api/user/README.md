# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [rpc/user/service.proto](#rpc/user/service.proto)
    - [CreateRequest](#todo.user.CreateRequest)
    - [CreateResponse](#todo.user.CreateResponse)
    - [DeleteRequest](#todo.user.DeleteRequest)
    - [DeleteResponse](#todo.user.DeleteResponse)
    - [GetLogsRequest](#todo.user.GetLogsRequest)
    - [GetLogsResponse](#todo.user.GetLogsResponse)
    - [GetOneRequest](#todo.user.GetOneRequest)
    - [GetOneResponse](#todo.user.GetOneResponse)
    - [PostbackLog](#todo.user.PostbackLog)
    - [UpdateRequest](#todo.user.UpdateRequest)
    - [UpdateResponse](#todo.user.UpdateResponse)
    - [User](#todo.user.User)
  
  
  
    - [UsersService](#todo.user.UsersService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="rpc/user/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rpc/user/service.proto



<a name="todo.user.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |






<a name="todo.user.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#todo.user.User) |  |  |






<a name="todo.user.DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="todo.user.DeleteResponse"></a>

### DeleteResponse







<a name="todo.user.GetLogsRequest"></a>

### GetLogsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| postback_id | [string](#string) |  |  |
| timestamp | [int64](#int64) |  |  |






<a name="todo.user.GetLogsResponse"></a>

### GetLogsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logs | [PostbackLog](#todo.user.PostbackLog) | repeated |  |






<a name="todo.user.GetOneRequest"></a>

### GetOneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="todo.user.GetOneResponse"></a>

### GetOneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#todo.user.User) |  |  |






<a name="todo.user.PostbackLog"></a>

### PostbackLog



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| postback_id | [string](#string) |  |  |
| status_code | [int32](#int32) |  |  |
| created | [int64](#int64) |  |  |
| response_body | [string](#string) |  |  |






<a name="todo.user.UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |






<a name="todo.user.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#todo.user.User) |  |  |






<a name="todo.user.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |
| created | [int64](#int64) |  |  |
| last_updated | [int64](#int64) |  |  |





 

 

 


<a name="todo.user.UsersService"></a>

### UsersService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#todo.user.CreateRequest) | [CreateResponse](#todo.user.CreateResponse) |  |
| GetOne | [GetOneRequest](#todo.user.GetOneRequest) | [GetOneResponse](#todo.user.GetOneResponse) |  |
| Update | [UpdateRequest](#todo.user.UpdateRequest) | [UpdateResponse](#todo.user.UpdateResponse) |  |
| Delete | [DeleteRequest](#todo.user.DeleteRequest) | [DeleteResponse](#todo.user.DeleteResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

