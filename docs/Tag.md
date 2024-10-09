# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**PostCount** | **float32** |  | 
**RelatedTags** | **[]string** |  | 
**RelatedTagsUpdatedAt** | **NullableTime** |  | 
**Category** | [**TagCategories**](TagCategories.md) |  | 
**IsLocked** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewTag

`func NewTag(id float32, name string, postCount float32, relatedTags []string, relatedTagsUpdatedAt NullableTime, category TagCategories, isLocked bool, createdAt time.Time, updatedAt time.Time, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tag) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### GetPostCount

`func (o *Tag) GetPostCount() float32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *Tag) GetPostCountOk() (*float32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *Tag) SetPostCount(v float32)`

SetPostCount sets PostCount field to given value.


### GetRelatedTags

`func (o *Tag) GetRelatedTags() []string`

GetRelatedTags returns the RelatedTags field if non-nil, zero value otherwise.

### GetRelatedTagsOk

`func (o *Tag) GetRelatedTagsOk() (*[]string, bool)`

GetRelatedTagsOk returns a tuple with the RelatedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTags

`func (o *Tag) SetRelatedTags(v []string)`

SetRelatedTags sets RelatedTags field to given value.


### GetRelatedTagsUpdatedAt

`func (o *Tag) GetRelatedTagsUpdatedAt() time.Time`

GetRelatedTagsUpdatedAt returns the RelatedTagsUpdatedAt field if non-nil, zero value otherwise.

### GetRelatedTagsUpdatedAtOk

`func (o *Tag) GetRelatedTagsUpdatedAtOk() (*time.Time, bool)`

GetRelatedTagsUpdatedAtOk returns a tuple with the RelatedTagsUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTagsUpdatedAt

`func (o *Tag) SetRelatedTagsUpdatedAt(v time.Time)`

SetRelatedTagsUpdatedAt sets RelatedTagsUpdatedAt field to given value.


### SetRelatedTagsUpdatedAtNil

`func (o *Tag) SetRelatedTagsUpdatedAtNil(b bool)`

 SetRelatedTagsUpdatedAtNil sets the value for RelatedTagsUpdatedAt to be an explicit nil

### UnsetRelatedTagsUpdatedAt
`func (o *Tag) UnsetRelatedTagsUpdatedAt()`

UnsetRelatedTagsUpdatedAt ensures that no value is present for RelatedTagsUpdatedAt, not even an explicit nil
### GetCategory

`func (o *Tag) GetCategory() TagCategories`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Tag) GetCategoryOk() (*TagCategories, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Tag) SetCategory(v TagCategories)`

SetCategory sets Category field to given value.


### GetIsLocked

`func (o *Tag) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Tag) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Tag) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetCreatedAt

`func (o *Tag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Tag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Tag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Tag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


