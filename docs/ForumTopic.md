# ForumTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatorId** | **float32** |  | 
**UpdaterId** | **float32** |  | 
**Title** | **string** |  | 
**ResponseCount** | **float32** |  | 
**IsSticky** | **bool** |  | 
**IsLocked** | **bool** |  | 
**IsHidden** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CategoryId** | **float32** |  | 

## Methods

### NewForumTopic

`func NewForumTopic(id float32, creatorId float32, updaterId float32, title string, responseCount float32, isSticky bool, isLocked bool, isHidden bool, createdAt time.Time, updatedAt time.Time, categoryId float32, ) *ForumTopic`

NewForumTopic instantiates a new ForumTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumTopicWithDefaults

`func NewForumTopicWithDefaults() *ForumTopic`

NewForumTopicWithDefaults instantiates a new ForumTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ForumTopic) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForumTopic) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForumTopic) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *ForumTopic) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ForumTopic) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ForumTopic) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetUpdaterId

`func (o *ForumTopic) GetUpdaterId() float32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *ForumTopic) GetUpdaterIdOk() (*float32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *ForumTopic) SetUpdaterId(v float32)`

SetUpdaterId sets UpdaterId field to given value.


### GetTitle

`func (o *ForumTopic) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ForumTopic) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ForumTopic) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetResponseCount

`func (o *ForumTopic) GetResponseCount() float32`

GetResponseCount returns the ResponseCount field if non-nil, zero value otherwise.

### GetResponseCountOk

`func (o *ForumTopic) GetResponseCountOk() (*float32, bool)`

GetResponseCountOk returns a tuple with the ResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCount

`func (o *ForumTopic) SetResponseCount(v float32)`

SetResponseCount sets ResponseCount field to given value.


### GetIsSticky

`func (o *ForumTopic) GetIsSticky() bool`

GetIsSticky returns the IsSticky field if non-nil, zero value otherwise.

### GetIsStickyOk

`func (o *ForumTopic) GetIsStickyOk() (*bool, bool)`

GetIsStickyOk returns a tuple with the IsSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSticky

`func (o *ForumTopic) SetIsSticky(v bool)`

SetIsSticky sets IsSticky field to given value.


### GetIsLocked

`func (o *ForumTopic) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ForumTopic) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ForumTopic) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsHidden

`func (o *ForumTopic) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ForumTopic) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ForumTopic) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.


### GetCreatedAt

`func (o *ForumTopic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ForumTopic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ForumTopic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ForumTopic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ForumTopic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ForumTopic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCategoryId

`func (o *ForumTopic) GetCategoryId() float32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ForumTopic) GetCategoryIdOk() (*float32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ForumTopic) SetCategoryId(v float32)`

SetCategoryId sets CategoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


