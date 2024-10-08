# Mascot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatorId** | **float32** |  | 
**DisplayName** | **string** |  | 
**Md5** | **string** |  | 
**FileExt** | **string** |  | 
**BackgroundColor** | **string** |  | 
**ArtistUrl** | **string** |  | 
**ArtistName** | **string** |  | 
**Active** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AvailableOn** | **[]string** |  | 
**UrlPath** | **string** |  | 

## Methods

### NewMascot

`func NewMascot(id float32, creatorId float32, displayName string, md5 string, fileExt string, backgroundColor string, artistUrl string, artistName string, active bool, createdAt time.Time, updatedAt time.Time, availableOn []string, urlPath string, ) *Mascot`

NewMascot instantiates a new Mascot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMascotWithDefaults

`func NewMascotWithDefaults() *Mascot`

NewMascotWithDefaults instantiates a new Mascot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Mascot) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mascot) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mascot) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *Mascot) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Mascot) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Mascot) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetDisplayName

`func (o *Mascot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Mascot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Mascot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMd5

`func (o *Mascot) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Mascot) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Mascot) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetFileExt

`func (o *Mascot) GetFileExt() string`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *Mascot) GetFileExtOk() (*string, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *Mascot) SetFileExt(v string)`

SetFileExt sets FileExt field to given value.


### GetBackgroundColor

`func (o *Mascot) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *Mascot) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *Mascot) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.


### GetArtistUrl

`func (o *Mascot) GetArtistUrl() string`

GetArtistUrl returns the ArtistUrl field if non-nil, zero value otherwise.

### GetArtistUrlOk

`func (o *Mascot) GetArtistUrlOk() (*string, bool)`

GetArtistUrlOk returns a tuple with the ArtistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistUrl

`func (o *Mascot) SetArtistUrl(v string)`

SetArtistUrl sets ArtistUrl field to given value.


### GetArtistName

`func (o *Mascot) GetArtistName() string`

GetArtistName returns the ArtistName field if non-nil, zero value otherwise.

### GetArtistNameOk

`func (o *Mascot) GetArtistNameOk() (*string, bool)`

GetArtistNameOk returns a tuple with the ArtistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistName

`func (o *Mascot) SetArtistName(v string)`

SetArtistName sets ArtistName field to given value.


### GetActive

`func (o *Mascot) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Mascot) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Mascot) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *Mascot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Mascot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Mascot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Mascot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Mascot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Mascot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAvailableOn

`func (o *Mascot) GetAvailableOn() []string`

GetAvailableOn returns the AvailableOn field if non-nil, zero value otherwise.

### GetAvailableOnOk

`func (o *Mascot) GetAvailableOnOk() (*[]string, bool)`

GetAvailableOnOk returns a tuple with the AvailableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOn

`func (o *Mascot) SetAvailableOn(v []string)`

SetAvailableOn sets AvailableOn field to given value.


### GetUrlPath

`func (o *Mascot) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *Mascot) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *Mascot) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


