# Blip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatorId** | **int32** |  | 
**Body** | **string** |  | 
**ResponseTo** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**IsHidden** | **bool** |  | 
**WarningType** | [**WarningTypes**](WarningTypes.md) |  | 
**WarningUserId** | **int32** |  | 
**UpdaterId** | **int32** |  | 
**CreatorName** | **string** |  | 

## Methods

### NewBlip

`func NewBlip(id int32, creatorId int32, body string, responseTo int32, createdAt time.Time, updatedAt time.Time, isHidden bool, warningType WarningTypes, warningUserId int32, updaterId int32, creatorName string, ) *Blip`

NewBlip instantiates a new Blip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlipWithDefaults

`func NewBlipWithDefaults() *Blip`

NewBlipWithDefaults instantiates a new Blip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Blip) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Blip) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Blip) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *Blip) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Blip) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Blip) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetBody

`func (o *Blip) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Blip) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Blip) SetBody(v string)`

SetBody sets Body field to given value.


### GetResponseTo

`func (o *Blip) GetResponseTo() int32`

GetResponseTo returns the ResponseTo field if non-nil, zero value otherwise.

### GetResponseToOk

`func (o *Blip) GetResponseToOk() (*int32, bool)`

GetResponseToOk returns a tuple with the ResponseTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTo

`func (o *Blip) SetResponseTo(v int32)`

SetResponseTo sets ResponseTo field to given value.


### GetCreatedAt

`func (o *Blip) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Blip) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Blip) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Blip) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Blip) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Blip) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsHidden

`func (o *Blip) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *Blip) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *Blip) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.


### GetWarningType

`func (o *Blip) GetWarningType() WarningTypes`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *Blip) GetWarningTypeOk() (*WarningTypes, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *Blip) SetWarningType(v WarningTypes)`

SetWarningType sets WarningType field to given value.


### GetWarningUserId

`func (o *Blip) GetWarningUserId() int32`

GetWarningUserId returns the WarningUserId field if non-nil, zero value otherwise.

### GetWarningUserIdOk

`func (o *Blip) GetWarningUserIdOk() (*int32, bool)`

GetWarningUserIdOk returns a tuple with the WarningUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningUserId

`func (o *Blip) SetWarningUserId(v int32)`

SetWarningUserId sets WarningUserId field to given value.


### GetUpdaterId

`func (o *Blip) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *Blip) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *Blip) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetCreatorName

`func (o *Blip) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *Blip) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *Blip) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


