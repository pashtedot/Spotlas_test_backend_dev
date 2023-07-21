# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [models_spots.proto](#models_spots-proto)
    - [Empty](#spots-models-Empty)
    - [ListSpotsNearPointRequest](#spots-models-ListSpotsNearPointRequest)
    - [ListSpotsNearPointResponse](#spots-models-ListSpotsNearPointResponse)
    - [Spot](#spots-models-Spot)
  
    - [QueryType](#spots-models-QueryType)
  
- [service_spots.proto](#service_spots-proto)
    - [CloseSpotsService](#spots-service-CloseSpotsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="models_spots-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## models_spots.proto



<a name="spots-models-Empty"></a>

### Empty
Empty is just empty message
that could return error






<a name="spots-models-ListSpotsNearPointRequest"></a>

### ListSpotsNearPointRequest
ArtworkTokenIdRequest is request model
for ListSpotsNearPoint


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| longitude | [double](#double) |  |  |
| latitude | [double](#double) |  |  |
| length | [float](#float) |  |  |
| search_type | [QueryType](#spots-models-QueryType) |  |  |






<a name="spots-models-ListSpotsNearPointResponse"></a>

### ListSpotsNearPointResponse
ChangeArtworkTokenIdRequest is request model
for ListSpotsNearPoint


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spots | [Spot](#spots-models-Spot) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="spots-models-Spot"></a>

### Spot
Spot model


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| rating | [float](#float) |  |  |
| website | [string](#string) |  |  |
| description | [string](#string) |  |  |
| coordinate | [string](#string) |  |  |
| uuid | [string](#string) |  |  |





 


<a name="spots-models-QueryType"></a>

### QueryType
QueryType - how do we search - in circle or in square

| Name | Number | Description |
| ---- | ------ | ----------- |
| QUERY_TYPE_INVALID | 0 |  |
| QUERY_TYPE_CIRCLE | 1 |  |
| QUERY_TYPE_SQUARE | 2 |  |


 

 

 



<a name="service_spots-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service_spots.proto


 

 

 


<a name="spots-service-CloseSpotsService"></a>

### CloseSpotsService
CloseSpotsService is GRPC server implementation
for close-spots-service service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListSpotsNearPoint | [.spots.models.ListSpotsNearPointRequest](#spots-models-ListSpotsNearPointRequest) | [.spots.models.ListSpotsNearPointResponse](#spots-models-ListSpotsNearPointResponse) | ListSpotsNearPoint returns a list of spots around the location |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

