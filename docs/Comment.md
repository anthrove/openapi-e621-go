# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**PostId** | **int32** |  | 
**CreatorId** | **int32** |  | 
**Body** | **string** |  | 
**Score** | **float32** |  | 
**UpdatedAt** | **time.Time** |  | 
**UpdaterId** | **int32** |  | 
**DoNotBumpPost** | **bool** |  | 
**IsHidden** | **bool** |  | 
**IsSticky** | **bool** |  | 
**WarningType** | [**WarningTypes**](WarningTypes.md) |  | 
**WarningUserId** | **NullableFloat32** |  | 
**CreatorName** | **string** |  | 
**UpdaterName** | **string** |  | 

## Methods

### NewComment

`func NewComment(id int32, createdAt time.Time, postId int32, creatorId int32, body string, score float32, updatedAt time.Time, updaterId int32, doNotBumpPost bool, isHidden bool, isSticky bool, warningType WarningTypes, warningUserId NullableFloat32, creatorName string, updaterName string, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Comment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPostId

`func (o *Comment) GetPostId() int32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *Comment) GetPostIdOk() (*int32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *Comment) SetPostId(v int32)`

SetPostId sets PostId field to given value.


### GetCreatorId

`func (o *Comment) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Comment) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Comment) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetBody

`func (o *Comment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Comment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Comment) SetBody(v string)`

SetBody sets Body field to given value.


### GetScore

`func (o *Comment) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Comment) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Comment) SetScore(v float32)`

SetScore sets Score field to given value.


### GetUpdatedAt

`func (o *Comment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Comment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Comment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdaterId

`func (o *Comment) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *Comment) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *Comment) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetDoNotBumpPost

`func (o *Comment) GetDoNotBumpPost() bool`

GetDoNotBumpPost returns the DoNotBumpPost field if non-nil, zero value otherwise.

### GetDoNotBumpPostOk

`func (o *Comment) GetDoNotBumpPostOk() (*bool, bool)`

GetDoNotBumpPostOk returns a tuple with the DoNotBumpPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotBumpPost

`func (o *Comment) SetDoNotBumpPost(v bool)`

SetDoNotBumpPost sets DoNotBumpPost field to given value.


### GetIsHidden

`func (o *Comment) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *Comment) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *Comment) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.


### GetIsSticky

`func (o *Comment) GetIsSticky() bool`

GetIsSticky returns the IsSticky field if non-nil, zero value otherwise.

### GetIsStickyOk

`func (o *Comment) GetIsStickyOk() (*bool, bool)`

GetIsStickyOk returns a tuple with the IsSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSticky

`func (o *Comment) SetIsSticky(v bool)`

SetIsSticky sets IsSticky field to given value.


### GetWarningType

`func (o *Comment) GetWarningType() WarningTypes`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *Comment) GetWarningTypeOk() (*WarningTypes, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *Comment) SetWarningType(v WarningTypes)`

SetWarningType sets WarningType field to given value.


### GetWarningUserId

`func (o *Comment) GetWarningUserId() float32`

GetWarningUserId returns the WarningUserId field if non-nil, zero value otherwise.

### GetWarningUserIdOk

`func (o *Comment) GetWarningUserIdOk() (*float32, bool)`

GetWarningUserIdOk returns a tuple with the WarningUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningUserId

`func (o *Comment) SetWarningUserId(v float32)`

SetWarningUserId sets WarningUserId field to given value.


### SetWarningUserIdNil

`func (o *Comment) SetWarningUserIdNil(b bool)`

 SetWarningUserIdNil sets the value for WarningUserId to be an explicit nil

### UnsetWarningUserId
`func (o *Comment) UnsetWarningUserId()`

UnsetWarningUserId ensures that no value is present for WarningUserId, not even an explicit nil
### GetCreatorName

`func (o *Comment) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *Comment) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *Comment) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetUpdaterName

`func (o *Comment) GetUpdaterName() string`

GetUpdaterName returns the UpdaterName field if non-nil, zero value otherwise.

### GetUpdaterNameOk

`func (o *Comment) GetUpdaterNameOk() (*string, bool)`

GetUpdaterNameOk returns a tuple with the UpdaterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterName

`func (o *Comment) SetUpdaterName(v string)`

SetUpdaterName sets UpdaterName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


