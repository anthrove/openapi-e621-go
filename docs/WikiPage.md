# WikiPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Title** | **string** |  | 
**Body** | **string** |  | 
**CreatorId** | **float32** |  | 
**IsLocked** | **bool** |  | 
**UpdaterId** | **float32** |  | 
**IsDeleted** | **bool** |  | 
**OtherNames** | **[]string** |  | 
**Parent** | **NullableString** |  | 
**CreatorName** | **string** |  | 
**CategoryId** | [**TagCategories**](TagCategories.md) |  | 

## Methods

### NewWikiPage

`func NewWikiPage(id float32, createdAt time.Time, updatedAt time.Time, title string, body string, creatorId float32, isLocked bool, updaterId float32, isDeleted bool, otherNames []string, parent NullableString, creatorName string, categoryId TagCategories, ) *WikiPage`

NewWikiPage instantiates a new WikiPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiPageWithDefaults

`func NewWikiPageWithDefaults() *WikiPage`

NewWikiPageWithDefaults instantiates a new WikiPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WikiPage) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiPage) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiPage) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *WikiPage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WikiPage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WikiPage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WikiPage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WikiPage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WikiPage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTitle

`func (o *WikiPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WikiPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WikiPage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *WikiPage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WikiPage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WikiPage) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatorId

`func (o *WikiPage) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *WikiPage) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *WikiPage) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetIsLocked

`func (o *WikiPage) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *WikiPage) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *WikiPage) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetUpdaterId

`func (o *WikiPage) GetUpdaterId() float32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *WikiPage) GetUpdaterIdOk() (*float32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *WikiPage) SetUpdaterId(v float32)`

SetUpdaterId sets UpdaterId field to given value.


### GetIsDeleted

`func (o *WikiPage) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WikiPage) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WikiPage) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetOtherNames

`func (o *WikiPage) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *WikiPage) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *WikiPage) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.


### GetParent

`func (o *WikiPage) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WikiPage) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WikiPage) SetParent(v string)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *WikiPage) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WikiPage) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetCreatorName

`func (o *WikiPage) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *WikiPage) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *WikiPage) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetCategoryId

`func (o *WikiPage) GetCategoryId() TagCategories`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *WikiPage) GetCategoryIdOk() (*TagCategories, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *WikiPage) SetCategoryId(v TagCategories)`

SetCategoryId sets CategoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


