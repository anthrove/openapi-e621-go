# PostDisapproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UserId** | **int32** |  | 
**PostId** | **int32** |  | 
**Reason** | **string** |  | 
**Message** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPostDisapproval

`func NewPostDisapproval(id int32, userId int32, postId int32, reason string, message NullableString, createdAt time.Time, updatedAt time.Time, ) *PostDisapproval`

NewPostDisapproval instantiates a new PostDisapproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDisapprovalWithDefaults

`func NewPostDisapprovalWithDefaults() *PostDisapproval`

NewPostDisapprovalWithDefaults instantiates a new PostDisapproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostDisapproval) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostDisapproval) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostDisapproval) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *PostDisapproval) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostDisapproval) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostDisapproval) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetPostId

`func (o *PostDisapproval) GetPostId() int32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostDisapproval) GetPostIdOk() (*int32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostDisapproval) SetPostId(v int32)`

SetPostId sets PostId field to given value.


### GetReason

`func (o *PostDisapproval) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostDisapproval) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostDisapproval) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *PostDisapproval) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostDisapproval) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostDisapproval) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *PostDisapproval) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PostDisapproval) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCreatedAt

`func (o *PostDisapproval) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostDisapproval) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostDisapproval) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PostDisapproval) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostDisapproval) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostDisapproval) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


