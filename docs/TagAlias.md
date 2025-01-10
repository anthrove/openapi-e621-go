# TagAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**AntecedentName** | **string** |  | 
**Reason** | **string** |  | 
**CreatorId** | **int32** |  | 
**CreatedAt** | **NullableTime** |  | 
**ForumPostId** | **NullableInt32** |  | 
**UpdatedAt** | **NullableTime** |  | 
**ForumTopicId** | **NullableInt32** |  | 
**ConsequentName** | **string** |  | 
**Status** | [**TagRequestStatuses**](TagRequestStatuses.md) | Note: The \&quot;error\&quot; status will be proceeded by an error, ex: \&quot;error: Validation failed: A tag alias for tag_name already exists\&quot;  | 
**PostCount** | **int32** |  | 
**ApproverId** | **NullableInt32** |  | 

## Methods

### NewTagAlias

`func NewTagAlias(id int32, antecedentName string, reason string, creatorId int32, createdAt NullableTime, forumPostId NullableInt32, updatedAt NullableTime, forumTopicId NullableInt32, consequentName string, status TagRequestStatuses, postCount int32, approverId NullableInt32, ) *TagAlias`

NewTagAlias instantiates a new TagAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagAliasWithDefaults

`func NewTagAliasWithDefaults() *TagAlias`

NewTagAliasWithDefaults instantiates a new TagAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagAlias) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagAlias) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagAlias) SetId(v int32)`

SetId sets Id field to given value.


### GetAntecedentName

`func (o *TagAlias) GetAntecedentName() string`

GetAntecedentName returns the AntecedentName field if non-nil, zero value otherwise.

### GetAntecedentNameOk

`func (o *TagAlias) GetAntecedentNameOk() (*string, bool)`

GetAntecedentNameOk returns a tuple with the AntecedentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntecedentName

`func (o *TagAlias) SetAntecedentName(v string)`

SetAntecedentName sets AntecedentName field to given value.


### GetReason

`func (o *TagAlias) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TagAlias) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TagAlias) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreatorId

`func (o *TagAlias) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *TagAlias) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *TagAlias) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetCreatedAt

`func (o *TagAlias) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagAlias) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagAlias) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *TagAlias) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TagAlias) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetForumPostId

`func (o *TagAlias) GetForumPostId() int32`

GetForumPostId returns the ForumPostId field if non-nil, zero value otherwise.

### GetForumPostIdOk

`func (o *TagAlias) GetForumPostIdOk() (*int32, bool)`

GetForumPostIdOk returns a tuple with the ForumPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostId

`func (o *TagAlias) SetForumPostId(v int32)`

SetForumPostId sets ForumPostId field to given value.


### SetForumPostIdNil

`func (o *TagAlias) SetForumPostIdNil(b bool)`

 SetForumPostIdNil sets the value for ForumPostId to be an explicit nil

### UnsetForumPostId
`func (o *TagAlias) UnsetForumPostId()`

UnsetForumPostId ensures that no value is present for ForumPostId, not even an explicit nil
### GetUpdatedAt

`func (o *TagAlias) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagAlias) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagAlias) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *TagAlias) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TagAlias) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetForumTopicId

`func (o *TagAlias) GetForumTopicId() int32`

GetForumTopicId returns the ForumTopicId field if non-nil, zero value otherwise.

### GetForumTopicIdOk

`func (o *TagAlias) GetForumTopicIdOk() (*int32, bool)`

GetForumTopicIdOk returns a tuple with the ForumTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicId

`func (o *TagAlias) SetForumTopicId(v int32)`

SetForumTopicId sets ForumTopicId field to given value.


### SetForumTopicIdNil

`func (o *TagAlias) SetForumTopicIdNil(b bool)`

 SetForumTopicIdNil sets the value for ForumTopicId to be an explicit nil

### UnsetForumTopicId
`func (o *TagAlias) UnsetForumTopicId()`

UnsetForumTopicId ensures that no value is present for ForumTopicId, not even an explicit nil
### GetConsequentName

`func (o *TagAlias) GetConsequentName() string`

GetConsequentName returns the ConsequentName field if non-nil, zero value otherwise.

### GetConsequentNameOk

`func (o *TagAlias) GetConsequentNameOk() (*string, bool)`

GetConsequentNameOk returns a tuple with the ConsequentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequentName

`func (o *TagAlias) SetConsequentName(v string)`

SetConsequentName sets ConsequentName field to given value.


### GetStatus

`func (o *TagAlias) GetStatus() TagRequestStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TagAlias) GetStatusOk() (*TagRequestStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TagAlias) SetStatus(v TagRequestStatuses)`

SetStatus sets Status field to given value.


### GetPostCount

`func (o *TagAlias) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *TagAlias) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *TagAlias) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.


### GetApproverId

`func (o *TagAlias) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *TagAlias) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *TagAlias) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *TagAlias) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *TagAlias) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


