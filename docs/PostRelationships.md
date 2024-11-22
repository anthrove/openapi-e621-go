# PostRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | **NullableFloat32** |  | 
**HasChildren** | **bool** |  | 
**HasActiveChildren** | **bool** |  | 
**Children** | **[]int32** |  | 

## Methods

### NewPostRelationships

`func NewPostRelationships(parentId NullableFloat32, hasChildren bool, hasActiveChildren bool, children []int32, ) *PostRelationships`

NewPostRelationships instantiates a new PostRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRelationshipsWithDefaults

`func NewPostRelationshipsWithDefaults() *PostRelationships`

NewPostRelationshipsWithDefaults instantiates a new PostRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *PostRelationships) GetParentId() float32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PostRelationships) GetParentIdOk() (*float32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PostRelationships) SetParentId(v float32)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *PostRelationships) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *PostRelationships) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetHasChildren

`func (o *PostRelationships) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *PostRelationships) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *PostRelationships) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.


### GetHasActiveChildren

`func (o *PostRelationships) GetHasActiveChildren() bool`

GetHasActiveChildren returns the HasActiveChildren field if non-nil, zero value otherwise.

### GetHasActiveChildrenOk

`func (o *PostRelationships) GetHasActiveChildrenOk() (*bool, bool)`

GetHasActiveChildrenOk returns a tuple with the HasActiveChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveChildren

`func (o *PostRelationships) SetHasActiveChildren(v bool)`

SetHasActiveChildren sets HasActiveChildren field to given value.


### GetChildren

`func (o *PostRelationships) GetChildren() []int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PostRelationships) GetChildrenOk() (*[]int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PostRelationships) SetChildren(v []int32)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


