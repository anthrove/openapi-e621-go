# RelatedTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CategoryId** | [**TagCategories**](TagCategories.md) |  | 

## Methods

### NewRelatedTag

`func NewRelatedTag(name string, categoryId TagCategories, ) *RelatedTag`

NewRelatedTag instantiates a new RelatedTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedTagWithDefaults

`func NewRelatedTagWithDefaults() *RelatedTag`

NewRelatedTagWithDefaults instantiates a new RelatedTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RelatedTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelatedTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelatedTag) SetName(v string)`

SetName sets Name field to given value.


### GetCategoryId

`func (o *RelatedTag) GetCategoryId() TagCategories`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *RelatedTag) GetCategoryIdOk() (*TagCategories, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *RelatedTag) SetCategoryId(v TagCategories)`

SetCategoryId sets CategoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


