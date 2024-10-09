# PoolVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PoolId** | **int32** |  | 
**PostIds** | **[]int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UpdaterId** | **int32** |  | 
**Name** | **string** |  | 
**NameChanged** | **bool** |  | 
**Description** | **string** |  | 
**DescriptionChanged** | **bool** |  | 
**IsActive** | **bool** |  | 
**IsLocked** | **bool** |  | 
**Category** | [**PoolCategories**](PoolCategories.md) |  | 
**Version** | **float32** |  | 
**AddedPostIds** | **[]int32** |  | 
**RemovedPostIds** | **[]int32** |  | 

## Methods

### NewPoolVersion

`func NewPoolVersion(id int32, poolId int32, postIds []int32, createdAt time.Time, updatedAt time.Time, updaterId int32, name string, nameChanged bool, description string, descriptionChanged bool, isActive bool, isLocked bool, category PoolCategories, version float32, addedPostIds []int32, removedPostIds []int32, ) *PoolVersion`

NewPoolVersion instantiates a new PoolVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolVersionWithDefaults

`func NewPoolVersionWithDefaults() *PoolVersion`

NewPoolVersionWithDefaults instantiates a new PoolVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetPoolId

`func (o *PoolVersion) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PoolVersion) GetPoolIdOk() (*int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PoolVersion) SetPoolId(v int32)`

SetPoolId sets PoolId field to given value.


### GetPostIds

`func (o *PoolVersion) GetPostIds() []int32`

GetPostIds returns the PostIds field if non-nil, zero value otherwise.

### GetPostIdsOk

`func (o *PoolVersion) GetPostIdsOk() (*[]int32, bool)`

GetPostIdsOk returns a tuple with the PostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostIds

`func (o *PoolVersion) SetPostIds(v []int32)`

SetPostIds sets PostIds field to given value.


### GetCreatedAt

`func (o *PoolVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoolVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoolVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PoolVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoolVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoolVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdaterId

`func (o *PoolVersion) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *PoolVersion) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *PoolVersion) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetName

`func (o *PoolVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolVersion) SetName(v string)`

SetName sets Name field to given value.


### GetNameChanged

`func (o *PoolVersion) GetNameChanged() bool`

GetNameChanged returns the NameChanged field if non-nil, zero value otherwise.

### GetNameChangedOk

`func (o *PoolVersion) GetNameChangedOk() (*bool, bool)`

GetNameChangedOk returns a tuple with the NameChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameChanged

`func (o *PoolVersion) SetNameChanged(v bool)`

SetNameChanged sets NameChanged field to given value.


### GetDescription

`func (o *PoolVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PoolVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PoolVersion) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionChanged

`func (o *PoolVersion) GetDescriptionChanged() bool`

GetDescriptionChanged returns the DescriptionChanged field if non-nil, zero value otherwise.

### GetDescriptionChangedOk

`func (o *PoolVersion) GetDescriptionChangedOk() (*bool, bool)`

GetDescriptionChangedOk returns a tuple with the DescriptionChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionChanged

`func (o *PoolVersion) SetDescriptionChanged(v bool)`

SetDescriptionChanged sets DescriptionChanged field to given value.


### GetIsActive

`func (o *PoolVersion) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PoolVersion) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PoolVersion) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsLocked

`func (o *PoolVersion) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PoolVersion) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PoolVersion) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetCategory

`func (o *PoolVersion) GetCategory() PoolCategories`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PoolVersion) GetCategoryOk() (*PoolCategories, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PoolVersion) SetCategory(v PoolCategories)`

SetCategory sets Category field to given value.


### GetVersion

`func (o *PoolVersion) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PoolVersion) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PoolVersion) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetAddedPostIds

`func (o *PoolVersion) GetAddedPostIds() []int32`

GetAddedPostIds returns the AddedPostIds field if non-nil, zero value otherwise.

### GetAddedPostIdsOk

`func (o *PoolVersion) GetAddedPostIdsOk() (*[]int32, bool)`

GetAddedPostIdsOk returns a tuple with the AddedPostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPostIds

`func (o *PoolVersion) SetAddedPostIds(v []int32)`

SetAddedPostIds sets AddedPostIds field to given value.


### GetRemovedPostIds

`func (o *PoolVersion) GetRemovedPostIds() []int32`

GetRemovedPostIds returns the RemovedPostIds field if non-nil, zero value otherwise.

### GetRemovedPostIdsOk

`func (o *PoolVersion) GetRemovedPostIdsOk() (*[]int32, bool)`

GetRemovedPostIdsOk returns a tuple with the RemovedPostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedPostIds

`func (o *PoolVersion) SetRemovedPostIds(v []int32)`

SetRemovedPostIds sets RemovedPostIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


