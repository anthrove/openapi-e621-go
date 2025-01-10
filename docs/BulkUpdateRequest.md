# BulkUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatorId** | **int32** |  | 
**ForumTopicId** | **NullableInt32** |  | 
**Script** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ApproverId** | **NullableInt32** |  | 
**ForumPostId** | **NullableInt32** |  | 
**Title** | **string** |  | 

## Methods

### NewBulkUpdateRequest

`func NewBulkUpdateRequest(id int32, creatorId int32, forumTopicId NullableInt32, script string, status string, createdAt time.Time, updatedAt time.Time, approverId NullableInt32, forumPostId NullableInt32, title string, ) *BulkUpdateRequest`

NewBulkUpdateRequest instantiates a new BulkUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateRequestWithDefaults

`func NewBulkUpdateRequestWithDefaults() *BulkUpdateRequest`

NewBulkUpdateRequestWithDefaults instantiates a new BulkUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *BulkUpdateRequest) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *BulkUpdateRequest) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *BulkUpdateRequest) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetForumTopicId

`func (o *BulkUpdateRequest) GetForumTopicId() int32`

GetForumTopicId returns the ForumTopicId field if non-nil, zero value otherwise.

### GetForumTopicIdOk

`func (o *BulkUpdateRequest) GetForumTopicIdOk() (*int32, bool)`

GetForumTopicIdOk returns a tuple with the ForumTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicId

`func (o *BulkUpdateRequest) SetForumTopicId(v int32)`

SetForumTopicId sets ForumTopicId field to given value.


### SetForumTopicIdNil

`func (o *BulkUpdateRequest) SetForumTopicIdNil(b bool)`

 SetForumTopicIdNil sets the value for ForumTopicId to be an explicit nil

### UnsetForumTopicId
`func (o *BulkUpdateRequest) UnsetForumTopicId()`

UnsetForumTopicId ensures that no value is present for ForumTopicId, not even an explicit nil
### GetScript

`func (o *BulkUpdateRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *BulkUpdateRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *BulkUpdateRequest) SetScript(v string)`

SetScript sets Script field to given value.


### GetStatus

`func (o *BulkUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BulkUpdateRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkUpdateRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkUpdateRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BulkUpdateRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BulkUpdateRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BulkUpdateRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetApproverId

`func (o *BulkUpdateRequest) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *BulkUpdateRequest) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *BulkUpdateRequest) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *BulkUpdateRequest) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *BulkUpdateRequest) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil
### GetForumPostId

`func (o *BulkUpdateRequest) GetForumPostId() int32`

GetForumPostId returns the ForumPostId field if non-nil, zero value otherwise.

### GetForumPostIdOk

`func (o *BulkUpdateRequest) GetForumPostIdOk() (*int32, bool)`

GetForumPostIdOk returns a tuple with the ForumPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostId

`func (o *BulkUpdateRequest) SetForumPostId(v int32)`

SetForumPostId sets ForumPostId field to given value.


### SetForumPostIdNil

`func (o *BulkUpdateRequest) SetForumPostIdNil(b bool)`

 SetForumPostIdNil sets the value for ForumPostId to be an explicit nil

### UnsetForumPostId
`func (o *BulkUpdateRequest) UnsetForumPostId()`

UnsetForumPostId ensures that no value is present for ForumPostId, not even an explicit nil
### GetTitle

`func (o *BulkUpdateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BulkUpdateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BulkUpdateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


