# CreateNote201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatorId** | **int32** |  | 
**X** | **float32** |  | 
**Y** | **float32** |  | 
**Width** | **float32** |  | 
**Height** | **float32** |  | 
**Version** | **float32** |  | 
**IsActive** | **bool** |  | 
**PostId** | **int32** |  | 
**Body** | **string** |  | 
**CreatorName** | **string** |  | 
**HtmlId** | **string** | Passthrough, used in frontend. | 

## Methods

### NewCreateNote201Response

`func NewCreateNote201Response(id int32, createdAt time.Time, updatedAt time.Time, creatorId int32, x float32, y float32, width float32, height float32, version float32, isActive bool, postId int32, body string, creatorName string, htmlId string, ) *CreateNote201Response`

NewCreateNote201Response instantiates a new CreateNote201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNote201ResponseWithDefaults

`func NewCreateNote201ResponseWithDefaults() *CreateNote201Response`

NewCreateNote201ResponseWithDefaults instantiates a new CreateNote201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNote201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNote201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNote201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreateNote201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateNote201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateNote201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateNote201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateNote201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateNote201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatorId

`func (o *CreateNote201Response) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CreateNote201Response) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CreateNote201Response) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetX

`func (o *CreateNote201Response) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *CreateNote201Response) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *CreateNote201Response) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *CreateNote201Response) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *CreateNote201Response) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *CreateNote201Response) SetY(v float32)`

SetY sets Y field to given value.


### GetWidth

`func (o *CreateNote201Response) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CreateNote201Response) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CreateNote201Response) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *CreateNote201Response) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CreateNote201Response) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CreateNote201Response) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetVersion

`func (o *CreateNote201Response) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateNote201Response) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateNote201Response) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetIsActive

`func (o *CreateNote201Response) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateNote201Response) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateNote201Response) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPostId

`func (o *CreateNote201Response) GetPostId() int32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *CreateNote201Response) GetPostIdOk() (*int32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *CreateNote201Response) SetPostId(v int32)`

SetPostId sets PostId field to given value.


### GetBody

`func (o *CreateNote201Response) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateNote201Response) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateNote201Response) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatorName

`func (o *CreateNote201Response) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *CreateNote201Response) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *CreateNote201Response) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetHtmlId

`func (o *CreateNote201Response) GetHtmlId() string`

GetHtmlId returns the HtmlId field if non-nil, zero value otherwise.

### GetHtmlIdOk

`func (o *CreateNote201Response) GetHtmlIdOk() (*string, bool)`

GetHtmlIdOk returns a tuple with the HtmlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlId

`func (o *CreateNote201Response) SetHtmlId(v string)`

SetHtmlId sets HtmlId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


