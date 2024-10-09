# ForumPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Topic** | Pointer to **float32** |  | [optional] 
**CreatorId** | **int32** |  | 
**UpdaterId** | **int32** |  | 
**Body** | **string** |  | 
**IsHidden** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**WarningType** | [**WarningTypes**](WarningTypes.md) |  | 
**WarningUserId** | **NullableFloat32** |  | 

## Methods

### NewForumPost

`func NewForumPost(id int32, creatorId int32, updaterId int32, body string, isHidden bool, createdAt time.Time, updatedAt time.Time, warningType WarningTypes, warningUserId NullableFloat32, ) *ForumPost`

NewForumPost instantiates a new ForumPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPostWithDefaults

`func NewForumPostWithDefaults() *ForumPost`

NewForumPostWithDefaults instantiates a new ForumPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ForumPost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForumPost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForumPost) SetId(v int32)`

SetId sets Id field to given value.


### GetTopic

`func (o *ForumPost) GetTopic() float32`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ForumPost) GetTopicOk() (*float32, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ForumPost) SetTopic(v float32)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ForumPost) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetCreatorId

`func (o *ForumPost) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ForumPost) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ForumPost) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetUpdaterId

`func (o *ForumPost) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *ForumPost) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *ForumPost) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetBody

`func (o *ForumPost) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ForumPost) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ForumPost) SetBody(v string)`

SetBody sets Body field to given value.


### GetIsHidden

`func (o *ForumPost) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ForumPost) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ForumPost) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.


### GetCreatedAt

`func (o *ForumPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ForumPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ForumPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ForumPost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ForumPost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ForumPost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWarningType

`func (o *ForumPost) GetWarningType() WarningTypes`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *ForumPost) GetWarningTypeOk() (*WarningTypes, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *ForumPost) SetWarningType(v WarningTypes)`

SetWarningType sets WarningType field to given value.


### GetWarningUserId

`func (o *ForumPost) GetWarningUserId() float32`

GetWarningUserId returns the WarningUserId field if non-nil, zero value otherwise.

### GetWarningUserIdOk

`func (o *ForumPost) GetWarningUserIdOk() (*float32, bool)`

GetWarningUserIdOk returns a tuple with the WarningUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningUserId

`func (o *ForumPost) SetWarningUserId(v float32)`

SetWarningUserId sets WarningUserId field to given value.


### SetWarningUserIdNil

`func (o *ForumPost) SetWarningUserIdNil(b bool)`

 SetWarningUserIdNil sets the value for WarningUserId to be an explicit nil

### UnsetWarningUserId
`func (o *ForumPost) UnsetWarningUserId()`

UnsetWarningUserId ensures that no value is present for WarningUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


