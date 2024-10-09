# BulkRelatedTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CategoryId** | [**TagCategories**](TagCategories.md) |  | 
**Count** | **int32** |  | 

## Methods

### NewBulkRelatedTag

`func NewBulkRelatedTag(name string, categoryId TagCategories, count int32, ) *BulkRelatedTag`

NewBulkRelatedTag instantiates a new BulkRelatedTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRelatedTagWithDefaults

`func NewBulkRelatedTagWithDefaults() *BulkRelatedTag`

NewBulkRelatedTagWithDefaults instantiates a new BulkRelatedTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BulkRelatedTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkRelatedTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkRelatedTag) SetName(v string)`

SetName sets Name field to given value.


### GetCategoryId

`func (o *BulkRelatedTag) GetCategoryId() TagCategories`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BulkRelatedTag) GetCategoryIdOk() (*TagCategories, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BulkRelatedTag) SetCategoryId(v TagCategories)`

SetCategoryId sets CategoryId field to given value.


### GetCount

`func (o *BulkRelatedTag) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BulkRelatedTag) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BulkRelatedTag) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


