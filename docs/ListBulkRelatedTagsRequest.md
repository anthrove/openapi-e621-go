# ListBulkRelatedTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to [**TagCategories**](TagCategories.md) |  | [optional] 

## Methods

### NewListBulkRelatedTagsRequest

`func NewListBulkRelatedTagsRequest() *ListBulkRelatedTagsRequest`

NewListBulkRelatedTagsRequest instantiates a new ListBulkRelatedTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBulkRelatedTagsRequestWithDefaults

`func NewListBulkRelatedTagsRequestWithDefaults() *ListBulkRelatedTagsRequest`

NewListBulkRelatedTagsRequestWithDefaults instantiates a new ListBulkRelatedTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ListBulkRelatedTagsRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListBulkRelatedTagsRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListBulkRelatedTagsRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListBulkRelatedTagsRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCategoryId

`func (o *ListBulkRelatedTagsRequest) GetCategoryId() TagCategories`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ListBulkRelatedTagsRequest) GetCategoryIdOk() (*TagCategories, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ListBulkRelatedTagsRequest) SetCategoryId(v TagCategories)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ListBulkRelatedTagsRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


