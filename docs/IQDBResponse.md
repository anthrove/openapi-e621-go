# IQDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**PostId** | **int32** |  | 
**Score** | **float32** |  | 
**Post** | [**IQDBResponsePost**](IQDBResponsePost.md) |  | 

## Methods

### NewIQDBResponse

`func NewIQDBResponse(hash string, postId int32, score float32, post IQDBResponsePost, ) *IQDBResponse`

NewIQDBResponse instantiates a new IQDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIQDBResponseWithDefaults

`func NewIQDBResponseWithDefaults() *IQDBResponse`

NewIQDBResponseWithDefaults instantiates a new IQDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *IQDBResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *IQDBResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *IQDBResponse) SetHash(v string)`

SetHash sets Hash field to given value.


### GetPostId

`func (o *IQDBResponse) GetPostId() int32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *IQDBResponse) GetPostIdOk() (*int32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *IQDBResponse) SetPostId(v int32)`

SetPostId sets PostId field to given value.


### GetScore

`func (o *IQDBResponse) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IQDBResponse) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IQDBResponse) SetScore(v float32)`

SetScore sets Score field to given value.


### GetPost

`func (o *IQDBResponse) GetPost() IQDBResponsePost`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *IQDBResponse) GetPostOk() (*IQDBResponsePost, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *IQDBResponse) SetPost(v IQDBResponsePost)`

SetPost sets Post field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


