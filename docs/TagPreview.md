# TagPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | **string** | The name if type&#x3D;tag, else the antecedent. | 
**B** | Pointer to **string** | The consequent, only present if type&#x3D;alias or type&#x3D;implication. | [optional] 
**Type** | **string** |  | 
**TagTypeA** | [**NullableTagCategories**](TagCategories.md) |  | 
**TagTypeB** | Pointer to [**NullableTagCategories**](TagCategories.md) |  | [optional] 

## Methods

### NewTagPreview

`func NewTagPreview(a string, type_ string, tagTypeA NullableTagCategories, ) *TagPreview`

NewTagPreview instantiates a new TagPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPreviewWithDefaults

`func NewTagPreviewWithDefaults() *TagPreview`

NewTagPreviewWithDefaults instantiates a new TagPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *TagPreview) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *TagPreview) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *TagPreview) SetA(v string)`

SetA sets A field to given value.


### GetB

`func (o *TagPreview) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *TagPreview) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *TagPreview) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *TagPreview) HasB() bool`

HasB returns a boolean if a field has been set.

### GetType

`func (o *TagPreview) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagPreview) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagPreview) SetType(v string)`

SetType sets Type field to given value.


### GetTagTypeA

`func (o *TagPreview) GetTagTypeA() TagCategories`

GetTagTypeA returns the TagTypeA field if non-nil, zero value otherwise.

### GetTagTypeAOk

`func (o *TagPreview) GetTagTypeAOk() (*TagCategories, bool)`

GetTagTypeAOk returns a tuple with the TagTypeA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypeA

`func (o *TagPreview) SetTagTypeA(v TagCategories)`

SetTagTypeA sets TagTypeA field to given value.


### SetTagTypeANil

`func (o *TagPreview) SetTagTypeANil(b bool)`

 SetTagTypeANil sets the value for TagTypeA to be an explicit nil

### UnsetTagTypeA
`func (o *TagPreview) UnsetTagTypeA()`

UnsetTagTypeA ensures that no value is present for TagTypeA, not even an explicit nil
### GetTagTypeB

`func (o *TagPreview) GetTagTypeB() TagCategories`

GetTagTypeB returns the TagTypeB field if non-nil, zero value otherwise.

### GetTagTypeBOk

`func (o *TagPreview) GetTagTypeBOk() (*TagCategories, bool)`

GetTagTypeBOk returns a tuple with the TagTypeB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypeB

`func (o *TagPreview) SetTagTypeB(v TagCategories)`

SetTagTypeB sets TagTypeB field to given value.

### HasTagTypeB

`func (o *TagPreview) HasTagTypeB() bool`

HasTagTypeB returns a boolean if a field has been set.

### SetTagTypeBNil

`func (o *TagPreview) SetTagTypeBNil(b bool)`

 SetTagTypeBNil sets the value for TagTypeB to be an explicit nil

### UnsetTagTypeB
`func (o *TagPreview) UnsetTagTypeB()`

UnsetTagTypeB ensures that no value is present for TagTypeB, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


