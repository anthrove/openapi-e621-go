# UserFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UserId** | **int32** |  | 
**CreatorId** | **int32** |  | 
**Category** | [**FeedbackCategories**](FeedbackCategories.md) |  | 
**Body** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UpdaterId** | **float32** |  | 
**IsDeleted** | **bool** |  | 

## Methods

### NewUserFeedback

`func NewUserFeedback(id int32, userId int32, creatorId int32, category FeedbackCategories, body string, createdAt time.Time, updatedAt time.Time, updaterId float32, isDeleted bool, ) *UserFeedback`

NewUserFeedback instantiates a new UserFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFeedbackWithDefaults

`func NewUserFeedbackWithDefaults() *UserFeedback`

NewUserFeedbackWithDefaults instantiates a new UserFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserFeedback) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserFeedback) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserFeedback) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *UserFeedback) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserFeedback) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserFeedback) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetCreatorId

`func (o *UserFeedback) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *UserFeedback) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *UserFeedback) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetCategory

`func (o *UserFeedback) GetCategory() FeedbackCategories`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UserFeedback) GetCategoryOk() (*FeedbackCategories, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UserFeedback) SetCategory(v FeedbackCategories)`

SetCategory sets Category field to given value.


### GetBody

`func (o *UserFeedback) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UserFeedback) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UserFeedback) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *UserFeedback) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserFeedback) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserFeedback) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserFeedback) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserFeedback) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserFeedback) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdaterId

`func (o *UserFeedback) GetUpdaterId() float32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *UserFeedback) GetUpdaterIdOk() (*float32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *UserFeedback) SetUpdaterId(v float32)`

SetUpdaterId sets UpdaterId field to given value.


### GetIsDeleted

`func (o *UserFeedback) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UserFeedback) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UserFeedback) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


