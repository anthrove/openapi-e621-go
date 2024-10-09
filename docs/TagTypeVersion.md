# TagTypeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**OldType** | [**TagCategories**](TagCategories.md) |  | 
**NewType** | [**TagCategories**](TagCategories.md) |  | 
**IsLocked** | **bool** |  | 
**TagId** | **int32** |  | 
**CreatorId** | **int32** |  | 

## Methods

### NewTagTypeVersion

`func NewTagTypeVersion(id int32, createdAt time.Time, updatedAt time.Time, oldType TagCategories, newType TagCategories, isLocked bool, tagId int32, creatorId int32, ) *TagTypeVersion`

NewTagTypeVersion instantiates a new TagTypeVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTypeVersionWithDefaults

`func NewTagTypeVersionWithDefaults() *TagTypeVersion`

NewTagTypeVersionWithDefaults instantiates a new TagTypeVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagTypeVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagTypeVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagTypeVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TagTypeVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagTypeVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagTypeVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TagTypeVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagTypeVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagTypeVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOldType

`func (o *TagTypeVersion) GetOldType() TagCategories`

GetOldType returns the OldType field if non-nil, zero value otherwise.

### GetOldTypeOk

`func (o *TagTypeVersion) GetOldTypeOk() (*TagCategories, bool)`

GetOldTypeOk returns a tuple with the OldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldType

`func (o *TagTypeVersion) SetOldType(v TagCategories)`

SetOldType sets OldType field to given value.


### GetNewType

`func (o *TagTypeVersion) GetNewType() TagCategories`

GetNewType returns the NewType field if non-nil, zero value otherwise.

### GetNewTypeOk

`func (o *TagTypeVersion) GetNewTypeOk() (*TagCategories, bool)`

GetNewTypeOk returns a tuple with the NewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewType

`func (o *TagTypeVersion) SetNewType(v TagCategories)`

SetNewType sets NewType field to given value.


### GetIsLocked

`func (o *TagTypeVersion) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *TagTypeVersion) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *TagTypeVersion) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetTagId

`func (o *TagTypeVersion) GetTagId() int32`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *TagTypeVersion) GetTagIdOk() (*int32, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *TagTypeVersion) SetTagId(v int32)`

SetTagId sets TagId field to given value.


### GetCreatorId

`func (o *TagTypeVersion) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *TagTypeVersion) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *TagTypeVersion) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


