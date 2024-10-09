# ArtistVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ArtistId** | **int32** |  | 
**Name** | **string** |  | 
**UpdaterId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**IsActive** | **bool** |  | 
**OtherNames** | **[]string** |  | 
**NotesChanged** | **bool** |  | 
**Urls** | **[]string** |  | 

## Methods

### NewArtistVersion

`func NewArtistVersion(id int32, artistId int32, name string, updaterId int32, createdAt time.Time, updatedAt time.Time, isActive bool, otherNames []string, notesChanged bool, urls []string, ) *ArtistVersion`

NewArtistVersion instantiates a new ArtistVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistVersionWithDefaults

`func NewArtistVersionWithDefaults() *ArtistVersion`

NewArtistVersionWithDefaults instantiates a new ArtistVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArtistVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtistVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtistVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetArtistId

`func (o *ArtistVersion) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *ArtistVersion) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *ArtistVersion) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.


### GetName

`func (o *ArtistVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtistVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtistVersion) SetName(v string)`

SetName sets Name field to given value.


### GetUpdaterId

`func (o *ArtistVersion) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *ArtistVersion) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *ArtistVersion) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetCreatedAt

`func (o *ArtistVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtistVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtistVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArtistVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArtistVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArtistVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsActive

`func (o *ArtistVersion) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ArtistVersion) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ArtistVersion) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetOtherNames

`func (o *ArtistVersion) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *ArtistVersion) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *ArtistVersion) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.


### GetNotesChanged

`func (o *ArtistVersion) GetNotesChanged() bool`

GetNotesChanged returns the NotesChanged field if non-nil, zero value otherwise.

### GetNotesChangedOk

`func (o *ArtistVersion) GetNotesChangedOk() (*bool, bool)`

GetNotesChangedOk returns a tuple with the NotesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesChanged

`func (o *ArtistVersion) SetNotesChanged(v bool)`

SetNotesChanged sets NotesChanged field to given value.


### GetUrls

`func (o *ArtistVersion) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ArtistVersion) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ArtistVersion) SetUrls(v []string)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


