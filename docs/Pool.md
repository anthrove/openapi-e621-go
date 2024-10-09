# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatorId** | **int32** |  | 
**Description** | **string** |  | 
**IsActive** | **bool** |  | 
**Category** | [**PoolCategories**](PoolCategories.md) |  | 
**PostIds** | **[]int32** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatorName** | **string** |  | 
**PostCount** | **int32** |  | 

## Methods

### NewPool

`func NewPool(id int32, name string, updatedAt time.Time, creatorId int32, description string, isActive bool, category PoolCategories, postIds []int32, createdAt time.Time, creatorName string, postCount int32, ) *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pool) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *Pool) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Pool) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Pool) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatorId

`func (o *Pool) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Pool) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Pool) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetDescription

`func (o *Pool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pool) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsActive

`func (o *Pool) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Pool) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Pool) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCategory

`func (o *Pool) GetCategory() PoolCategories`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Pool) GetCategoryOk() (*PoolCategories, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Pool) SetCategory(v PoolCategories)`

SetCategory sets Category field to given value.


### GetPostIds

`func (o *Pool) GetPostIds() []int32`

GetPostIds returns the PostIds field if non-nil, zero value otherwise.

### GetPostIdsOk

`func (o *Pool) GetPostIdsOk() (*[]int32, bool)`

GetPostIdsOk returns a tuple with the PostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostIds

`func (o *Pool) SetPostIds(v []int32)`

SetPostIds sets PostIds field to given value.


### GetCreatedAt

`func (o *Pool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorName

`func (o *Pool) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *Pool) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *Pool) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetPostCount

`func (o *Pool) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *Pool) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *Pool) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


