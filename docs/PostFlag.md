# PostFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**PostId** | **float32** |  | 
**Reason** | **string** |  | 
**CreatorId** | **NullableFloat32** |  | 
**IsResolved** | **bool** |  | 
**UpdatedAt** | **time.Time** |  | 
**IsDeletion** | **bool** |  | 
**Type** | **string** |  | 

## Methods

### NewPostFlag

`func NewPostFlag(id float32, createdAt time.Time, postId float32, reason string, creatorId NullableFloat32, isResolved bool, updatedAt time.Time, isDeletion bool, type_ string, ) *PostFlag`

NewPostFlag instantiates a new PostFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFlagWithDefaults

`func NewPostFlagWithDefaults() *PostFlag`

NewPostFlagWithDefaults instantiates a new PostFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostFlag) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostFlag) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostFlag) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PostFlag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostFlag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostFlag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPostId

`func (o *PostFlag) GetPostId() float32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostFlag) GetPostIdOk() (*float32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostFlag) SetPostId(v float32)`

SetPostId sets PostId field to given value.


### GetReason

`func (o *PostFlag) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostFlag) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostFlag) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreatorId

`func (o *PostFlag) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PostFlag) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PostFlag) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *PostFlag) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *PostFlag) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetIsResolved

`func (o *PostFlag) GetIsResolved() bool`

GetIsResolved returns the IsResolved field if non-nil, zero value otherwise.

### GetIsResolvedOk

`func (o *PostFlag) GetIsResolvedOk() (*bool, bool)`

GetIsResolvedOk returns a tuple with the IsResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolved

`func (o *PostFlag) SetIsResolved(v bool)`

SetIsResolved sets IsResolved field to given value.


### GetUpdatedAt

`func (o *PostFlag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostFlag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostFlag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsDeletion

`func (o *PostFlag) GetIsDeletion() bool`

GetIsDeletion returns the IsDeletion field if non-nil, zero value otherwise.

### GetIsDeletionOk

`func (o *PostFlag) GetIsDeletionOk() (*bool, bool)`

GetIsDeletionOk returns a tuple with the IsDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletion

`func (o *PostFlag) SetIsDeletion(v bool)`

SetIsDeletion sets IsDeletion field to given value.


### GetType

`func (o *PostFlag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostFlag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostFlag) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


