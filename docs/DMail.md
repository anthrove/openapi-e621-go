# DMail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**OwnerId** | **int32** |  | 
**FromId** | **int32** |  | 
**ToId** | **int32** |  | 
**Title** | **string** |  | 
**Body** | **string** |  | 
**IsRead** | **bool** |  | 
**IsDeleted** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewDMail

`func NewDMail(id int32, ownerId int32, fromId int32, toId int32, title string, body string, isRead bool, isDeleted bool, createdAt time.Time, updatedAt time.Time, ) *DMail`

NewDMail instantiates a new DMail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDMailWithDefaults

`func NewDMailWithDefaults() *DMail`

NewDMailWithDefaults instantiates a new DMail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DMail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DMail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DMail) SetId(v int32)`

SetId sets Id field to given value.


### GetOwnerId

`func (o *DMail) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DMail) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DMail) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.


### GetFromId

`func (o *DMail) GetFromId() int32`

GetFromId returns the FromId field if non-nil, zero value otherwise.

### GetFromIdOk

`func (o *DMail) GetFromIdOk() (*int32, bool)`

GetFromIdOk returns a tuple with the FromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromId

`func (o *DMail) SetFromId(v int32)`

SetFromId sets FromId field to given value.


### GetToId

`func (o *DMail) GetToId() int32`

GetToId returns the ToId field if non-nil, zero value otherwise.

### GetToIdOk

`func (o *DMail) GetToIdOk() (*int32, bool)`

GetToIdOk returns a tuple with the ToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToId

`func (o *DMail) SetToId(v int32)`

SetToId sets ToId field to given value.


### GetTitle

`func (o *DMail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DMail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DMail) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *DMail) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DMail) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DMail) SetBody(v string)`

SetBody sets Body field to given value.


### GetIsRead

`func (o *DMail) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *DMail) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *DMail) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.


### GetIsDeleted

`func (o *DMail) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DMail) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DMail) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetCreatedAt

`func (o *DMail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DMail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DMail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DMail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DMail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DMail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


