# PostEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatorId** | **NullableFloat32** |  | 
**PostId** | **float32** |  | 
**Action** | [**PostEventActions**](PostEventActions.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPostEvent

`func NewPostEvent(id float32, creatorId NullableFloat32, postId float32, action PostEventActions, createdAt time.Time, ) *PostEvent`

NewPostEvent instantiates a new PostEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEventWithDefaults

`func NewPostEventWithDefaults() *PostEvent`

NewPostEventWithDefaults instantiates a new PostEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostEvent) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostEvent) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostEvent) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *PostEvent) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PostEvent) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PostEvent) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *PostEvent) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *PostEvent) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetPostId

`func (o *PostEvent) GetPostId() float32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostEvent) GetPostIdOk() (*float32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostEvent) SetPostId(v float32)`

SetPostId sets PostId field to given value.


### GetAction

`func (o *PostEvent) GetAction() PostEventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostEvent) GetActionOk() (*PostEventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostEvent) SetAction(v PostEventActions)`

SetAction sets Action field to given value.


### GetCreatedAt

`func (o *PostEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


