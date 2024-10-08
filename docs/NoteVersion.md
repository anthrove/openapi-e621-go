# NoteVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**X** | **float32** |  | 
**Y** | **float32** |  | 
**Width** | **float32** |  | 
**Height** | **float32** |  | 
**Body** | **string** |  | 
**Version** | **float32** |  | 
**IsActive** | **bool** |  | 
**NoteId** | **float32** |  | 
**PostId** | **float32** |  | 
**UpdaterId** | **float32** |  | 

## Methods

### NewNoteVersion

`func NewNoteVersion(id float32, createdAt time.Time, updatedAt time.Time, x float32, y float32, width float32, height float32, body string, version float32, isActive bool, noteId float32, postId float32, updaterId float32, ) *NoteVersion`

NewNoteVersion instantiates a new NoteVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteVersionWithDefaults

`func NewNoteVersionWithDefaults() *NoteVersion`

NewNoteVersionWithDefaults instantiates a new NoteVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteVersion) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteVersion) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteVersion) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *NoteVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NoteVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NoteVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NoteVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NoteVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NoteVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetX

`func (o *NoteVersion) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *NoteVersion) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *NoteVersion) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *NoteVersion) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *NoteVersion) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *NoteVersion) SetY(v float32)`

SetY sets Y field to given value.


### GetWidth

`func (o *NoteVersion) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *NoteVersion) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *NoteVersion) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *NoteVersion) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *NoteVersion) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *NoteVersion) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetBody

`func (o *NoteVersion) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NoteVersion) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NoteVersion) SetBody(v string)`

SetBody sets Body field to given value.


### GetVersion

`func (o *NoteVersion) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NoteVersion) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NoteVersion) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetIsActive

`func (o *NoteVersion) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NoteVersion) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NoteVersion) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetNoteId

`func (o *NoteVersion) GetNoteId() float32`

GetNoteId returns the NoteId field if non-nil, zero value otherwise.

### GetNoteIdOk

`func (o *NoteVersion) GetNoteIdOk() (*float32, bool)`

GetNoteIdOk returns a tuple with the NoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteId

`func (o *NoteVersion) SetNoteId(v float32)`

SetNoteId sets NoteId field to given value.


### GetPostId

`func (o *NoteVersion) GetPostId() float32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *NoteVersion) GetPostIdOk() (*float32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *NoteVersion) SetPostId(v float32)`

SetPostId sets PostId field to given value.


### GetUpdaterId

`func (o *NoteVersion) GetUpdaterId() float32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *NoteVersion) GetUpdaterIdOk() (*float32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *NoteVersion) SetUpdaterId(v float32)`

SetUpdaterId sets UpdaterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


