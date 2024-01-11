/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-odap-hermes

import (
	"encoding/json"
)

// checks if the TransferInitializationV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferInitializationV1Request{}

// TransferInitializationV1Request struct for TransferInitializationV1Request
type TransferInitializationV1Request struct {
	MessageType string `json:"messageType"`
	SessionID string `json:"sessionID"`
	Version *string `json:"version,omitempty"`
	DeveloperURN *string `json:"developerURN,omitempty"`
	CredentialProfile *CredentialProfile `json:"credentialProfile,omitempty"`
	PayloadProfile PayloadProfile `json:"payloadProfile"`
	ApplicationProfile string `json:"applicationProfile"`
	LoggingProfile string `json:"loggingProfile"`
	AccessControlProfile string `json:"accessControlProfile"`
	Signature string `json:"signature"`
	SourceGatewayPubkey string `json:"sourceGatewayPubkey"`
	SourceGatewayDltSystem string `json:"sourceGatewayDltSystem"`
	RecipientGatewayPubkey string `json:"recipientGatewayPubkey"`
	RecipientGatewayDltSystem string `json:"recipientGatewayDltSystem"`
	EscrowType *string `json:"escrowType,omitempty"`
	ExpiryTime *string `json:"expiryTime,omitempty"`
	MultipleClaimsAllowed *bool `json:"multipleClaimsAllowed,omitempty"`
	MultipleCancelsAllowed *bool `json:"multipleCancelsAllowed,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Destination *string `json:"destination,omitempty"`
	SubsequentCalls map[string]interface{} `json:"subsequentCalls,omitempty"`
	Histories []History `json:"histories,omitempty"`
	SequenceNumber int32 `json:"sequenceNumber"`
	SourceBasePath string `json:"sourceBasePath"`
	RecipientBasePath string `json:"recipientBasePath"`
	MaxRetries float32 `json:"maxRetries"`
	MaxTimeout float32 `json:"maxTimeout"`
	BackupGatewaysAllowed []string `json:"backupGatewaysAllowed"`
	RecipientLedgerAssetID string `json:"recipientLedgerAssetID"`
	SourceLedgerAssetID string `json:"sourceLedgerAssetID"`
}

// NewTransferInitializationV1Request instantiates a new TransferInitializationV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferInitializationV1Request(messageType string, sessionID string, payloadProfile PayloadProfile, applicationProfile string, loggingProfile string, accessControlProfile string, signature string, sourceGatewayPubkey string, sourceGatewayDltSystem string, recipientGatewayPubkey string, recipientGatewayDltSystem string, sequenceNumber int32, sourceBasePath string, recipientBasePath string, maxRetries float32, maxTimeout float32, backupGatewaysAllowed []string, recipientLedgerAssetID string, sourceLedgerAssetID string) *TransferInitializationV1Request {
	this := TransferInitializationV1Request{}
	this.MessageType = messageType
	this.SessionID = sessionID
	this.PayloadProfile = payloadProfile
	this.ApplicationProfile = applicationProfile
	this.LoggingProfile = loggingProfile
	this.AccessControlProfile = accessControlProfile
	this.Signature = signature
	this.SourceGatewayPubkey = sourceGatewayPubkey
	this.SourceGatewayDltSystem = sourceGatewayDltSystem
	this.RecipientGatewayPubkey = recipientGatewayPubkey
	this.RecipientGatewayDltSystem = recipientGatewayDltSystem
	this.SequenceNumber = sequenceNumber
	this.SourceBasePath = sourceBasePath
	this.RecipientBasePath = recipientBasePath
	this.MaxRetries = maxRetries
	this.MaxTimeout = maxTimeout
	this.BackupGatewaysAllowed = backupGatewaysAllowed
	this.RecipientLedgerAssetID = recipientLedgerAssetID
	this.SourceLedgerAssetID = sourceLedgerAssetID
	return &this
}

// NewTransferInitializationV1RequestWithDefaults instantiates a new TransferInitializationV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferInitializationV1RequestWithDefaults() *TransferInitializationV1Request {
	this := TransferInitializationV1Request{}
	return &this
}

// GetMessageType returns the MessageType field value
func (o *TransferInitializationV1Request) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *TransferInitializationV1Request) SetMessageType(v string) {
	o.MessageType = v
}

// GetSessionID returns the SessionID field value
func (o *TransferInitializationV1Request) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *TransferInitializationV1Request) SetSessionID(v string) {
	o.SessionID = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TransferInitializationV1Request) SetVersion(v string) {
	o.Version = &v
}

// GetDeveloperURN returns the DeveloperURN field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetDeveloperURN() string {
	if o == nil || IsNil(o.DeveloperURN) {
		var ret string
		return ret
	}
	return *o.DeveloperURN
}

// GetDeveloperURNOk returns a tuple with the DeveloperURN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetDeveloperURNOk() (*string, bool) {
	if o == nil || IsNil(o.DeveloperURN) {
		return nil, false
	}
	return o.DeveloperURN, true
}

// HasDeveloperURN returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasDeveloperURN() bool {
	if o != nil && !IsNil(o.DeveloperURN) {
		return true
	}

	return false
}

// SetDeveloperURN gets a reference to the given string and assigns it to the DeveloperURN field.
func (o *TransferInitializationV1Request) SetDeveloperURN(v string) {
	o.DeveloperURN = &v
}

// GetCredentialProfile returns the CredentialProfile field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetCredentialProfile() CredentialProfile {
	if o == nil || IsNil(o.CredentialProfile) {
		var ret CredentialProfile
		return ret
	}
	return *o.CredentialProfile
}

// GetCredentialProfileOk returns a tuple with the CredentialProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetCredentialProfileOk() (*CredentialProfile, bool) {
	if o == nil || IsNil(o.CredentialProfile) {
		return nil, false
	}
	return o.CredentialProfile, true
}

// HasCredentialProfile returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasCredentialProfile() bool {
	if o != nil && !IsNil(o.CredentialProfile) {
		return true
	}

	return false
}

// SetCredentialProfile gets a reference to the given CredentialProfile and assigns it to the CredentialProfile field.
func (o *TransferInitializationV1Request) SetCredentialProfile(v CredentialProfile) {
	o.CredentialProfile = &v
}

// GetPayloadProfile returns the PayloadProfile field value
func (o *TransferInitializationV1Request) GetPayloadProfile() PayloadProfile {
	if o == nil {
		var ret PayloadProfile
		return ret
	}

	return o.PayloadProfile
}

// GetPayloadProfileOk returns a tuple with the PayloadProfile field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetPayloadProfileOk() (*PayloadProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadProfile, true
}

// SetPayloadProfile sets field value
func (o *TransferInitializationV1Request) SetPayloadProfile(v PayloadProfile) {
	o.PayloadProfile = v
}

// GetApplicationProfile returns the ApplicationProfile field value
func (o *TransferInitializationV1Request) GetApplicationProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationProfile
}

// GetApplicationProfileOk returns a tuple with the ApplicationProfile field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetApplicationProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationProfile, true
}

// SetApplicationProfile sets field value
func (o *TransferInitializationV1Request) SetApplicationProfile(v string) {
	o.ApplicationProfile = v
}

// GetLoggingProfile returns the LoggingProfile field value
func (o *TransferInitializationV1Request) GetLoggingProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoggingProfile
}

// GetLoggingProfileOk returns a tuple with the LoggingProfile field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetLoggingProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoggingProfile, true
}

// SetLoggingProfile sets field value
func (o *TransferInitializationV1Request) SetLoggingProfile(v string) {
	o.LoggingProfile = v
}

// GetAccessControlProfile returns the AccessControlProfile field value
func (o *TransferInitializationV1Request) GetAccessControlProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessControlProfile
}

// GetAccessControlProfileOk returns a tuple with the AccessControlProfile field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetAccessControlProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessControlProfile, true
}

// SetAccessControlProfile sets field value
func (o *TransferInitializationV1Request) SetAccessControlProfile(v string) {
	o.AccessControlProfile = v
}

// GetSignature returns the Signature field value
func (o *TransferInitializationV1Request) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *TransferInitializationV1Request) SetSignature(v string) {
	o.Signature = v
}

// GetSourceGatewayPubkey returns the SourceGatewayPubkey field value
func (o *TransferInitializationV1Request) GetSourceGatewayPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceGatewayPubkey
}

// GetSourceGatewayPubkeyOk returns a tuple with the SourceGatewayPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSourceGatewayPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceGatewayPubkey, true
}

// SetSourceGatewayPubkey sets field value
func (o *TransferInitializationV1Request) SetSourceGatewayPubkey(v string) {
	o.SourceGatewayPubkey = v
}

// GetSourceGatewayDltSystem returns the SourceGatewayDltSystem field value
func (o *TransferInitializationV1Request) GetSourceGatewayDltSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceGatewayDltSystem
}

// GetSourceGatewayDltSystemOk returns a tuple with the SourceGatewayDltSystem field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSourceGatewayDltSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceGatewayDltSystem, true
}

// SetSourceGatewayDltSystem sets field value
func (o *TransferInitializationV1Request) SetSourceGatewayDltSystem(v string) {
	o.SourceGatewayDltSystem = v
}

// GetRecipientGatewayPubkey returns the RecipientGatewayPubkey field value
func (o *TransferInitializationV1Request) GetRecipientGatewayPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientGatewayPubkey
}

// GetRecipientGatewayPubkeyOk returns a tuple with the RecipientGatewayPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetRecipientGatewayPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientGatewayPubkey, true
}

// SetRecipientGatewayPubkey sets field value
func (o *TransferInitializationV1Request) SetRecipientGatewayPubkey(v string) {
	o.RecipientGatewayPubkey = v
}

// GetRecipientGatewayDltSystem returns the RecipientGatewayDltSystem field value
func (o *TransferInitializationV1Request) GetRecipientGatewayDltSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientGatewayDltSystem
}

// GetRecipientGatewayDltSystemOk returns a tuple with the RecipientGatewayDltSystem field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetRecipientGatewayDltSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientGatewayDltSystem, true
}

// SetRecipientGatewayDltSystem sets field value
func (o *TransferInitializationV1Request) SetRecipientGatewayDltSystem(v string) {
	o.RecipientGatewayDltSystem = v
}

// GetEscrowType returns the EscrowType field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetEscrowType() string {
	if o == nil || IsNil(o.EscrowType) {
		var ret string
		return ret
	}
	return *o.EscrowType
}

// GetEscrowTypeOk returns a tuple with the EscrowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetEscrowTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EscrowType) {
		return nil, false
	}
	return o.EscrowType, true
}

// HasEscrowType returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasEscrowType() bool {
	if o != nil && !IsNil(o.EscrowType) {
		return true
	}

	return false
}

// SetEscrowType gets a reference to the given string and assigns it to the EscrowType field.
func (o *TransferInitializationV1Request) SetEscrowType(v string) {
	o.EscrowType = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetExpiryTime() string {
	if o == nil || IsNil(o.ExpiryTime) {
		var ret string
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasExpiryTime() bool {
	if o != nil && !IsNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given string and assigns it to the ExpiryTime field.
func (o *TransferInitializationV1Request) SetExpiryTime(v string) {
	o.ExpiryTime = &v
}

// GetMultipleClaimsAllowed returns the MultipleClaimsAllowed field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetMultipleClaimsAllowed() bool {
	if o == nil || IsNil(o.MultipleClaimsAllowed) {
		var ret bool
		return ret
	}
	return *o.MultipleClaimsAllowed
}

// GetMultipleClaimsAllowedOk returns a tuple with the MultipleClaimsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetMultipleClaimsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultipleClaimsAllowed) {
		return nil, false
	}
	return o.MultipleClaimsAllowed, true
}

// HasMultipleClaimsAllowed returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasMultipleClaimsAllowed() bool {
	if o != nil && !IsNil(o.MultipleClaimsAllowed) {
		return true
	}

	return false
}

// SetMultipleClaimsAllowed gets a reference to the given bool and assigns it to the MultipleClaimsAllowed field.
func (o *TransferInitializationV1Request) SetMultipleClaimsAllowed(v bool) {
	o.MultipleClaimsAllowed = &v
}

// GetMultipleCancelsAllowed returns the MultipleCancelsAllowed field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetMultipleCancelsAllowed() bool {
	if o == nil || IsNil(o.MultipleCancelsAllowed) {
		var ret bool
		return ret
	}
	return *o.MultipleCancelsAllowed
}

// GetMultipleCancelsAllowedOk returns a tuple with the MultipleCancelsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetMultipleCancelsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultipleCancelsAllowed) {
		return nil, false
	}
	return o.MultipleCancelsAllowed, true
}

// HasMultipleCancelsAllowed returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasMultipleCancelsAllowed() bool {
	if o != nil && !IsNil(o.MultipleCancelsAllowed) {
		return true
	}

	return false
}

// SetMultipleCancelsAllowed gets a reference to the given bool and assigns it to the MultipleCancelsAllowed field.
func (o *TransferInitializationV1Request) SetMultipleCancelsAllowed(v bool) {
	o.MultipleCancelsAllowed = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetPermissions() map[string]interface{} {
	if o == nil || IsNil(o.Permissions) {
		var ret map[string]interface{}
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetPermissionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Permissions) {
		return map[string]interface{}{}, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]interface{} and assigns it to the Permissions field.
func (o *TransferInitializationV1Request) SetPermissions(v map[string]interface{}) {
	o.Permissions = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *TransferInitializationV1Request) SetOrigin(v string) {
	o.Origin = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *TransferInitializationV1Request) SetDestination(v string) {
	o.Destination = &v
}

// GetSubsequentCalls returns the SubsequentCalls field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetSubsequentCalls() map[string]interface{} {
	if o == nil || IsNil(o.SubsequentCalls) {
		var ret map[string]interface{}
		return ret
	}
	return o.SubsequentCalls
}

// GetSubsequentCallsOk returns a tuple with the SubsequentCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSubsequentCallsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SubsequentCalls) {
		return map[string]interface{}{}, false
	}
	return o.SubsequentCalls, true
}

// HasSubsequentCalls returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasSubsequentCalls() bool {
	if o != nil && !IsNil(o.SubsequentCalls) {
		return true
	}

	return false
}

// SetSubsequentCalls gets a reference to the given map[string]interface{} and assigns it to the SubsequentCalls field.
func (o *TransferInitializationV1Request) SetSubsequentCalls(v map[string]interface{}) {
	o.SubsequentCalls = v
}

// GetHistories returns the Histories field value if set, zero value otherwise.
func (o *TransferInitializationV1Request) GetHistories() []History {
	if o == nil || IsNil(o.Histories) {
		var ret []History
		return ret
	}
	return o.Histories
}

// GetHistoriesOk returns a tuple with the Histories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetHistoriesOk() ([]History, bool) {
	if o == nil || IsNil(o.Histories) {
		return nil, false
	}
	return o.Histories, true
}

// HasHistories returns a boolean if a field has been set.
func (o *TransferInitializationV1Request) HasHistories() bool {
	if o != nil && !IsNil(o.Histories) {
		return true
	}

	return false
}

// SetHistories gets a reference to the given []History and assigns it to the Histories field.
func (o *TransferInitializationV1Request) SetHistories(v []History) {
	o.Histories = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *TransferInitializationV1Request) GetSequenceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSequenceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *TransferInitializationV1Request) SetSequenceNumber(v int32) {
	o.SequenceNumber = v
}

// GetSourceBasePath returns the SourceBasePath field value
func (o *TransferInitializationV1Request) GetSourceBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceBasePath
}

// GetSourceBasePathOk returns a tuple with the SourceBasePath field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSourceBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceBasePath, true
}

// SetSourceBasePath sets field value
func (o *TransferInitializationV1Request) SetSourceBasePath(v string) {
	o.SourceBasePath = v
}

// GetRecipientBasePath returns the RecipientBasePath field value
func (o *TransferInitializationV1Request) GetRecipientBasePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientBasePath
}

// GetRecipientBasePathOk returns a tuple with the RecipientBasePath field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetRecipientBasePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientBasePath, true
}

// SetRecipientBasePath sets field value
func (o *TransferInitializationV1Request) SetRecipientBasePath(v string) {
	o.RecipientBasePath = v
}

// GetMaxRetries returns the MaxRetries field value
func (o *TransferInitializationV1Request) GetMaxRetries() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetMaxRetriesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRetries, true
}

// SetMaxRetries sets field value
func (o *TransferInitializationV1Request) SetMaxRetries(v float32) {
	o.MaxRetries = v
}

// GetMaxTimeout returns the MaxTimeout field value
func (o *TransferInitializationV1Request) GetMaxTimeout() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxTimeout
}

// GetMaxTimeoutOk returns a tuple with the MaxTimeout field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetMaxTimeoutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTimeout, true
}

// SetMaxTimeout sets field value
func (o *TransferInitializationV1Request) SetMaxTimeout(v float32) {
	o.MaxTimeout = v
}

// GetBackupGatewaysAllowed returns the BackupGatewaysAllowed field value
func (o *TransferInitializationV1Request) GetBackupGatewaysAllowed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackupGatewaysAllowed
}

// GetBackupGatewaysAllowedOk returns a tuple with the BackupGatewaysAllowed field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetBackupGatewaysAllowedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupGatewaysAllowed, true
}

// SetBackupGatewaysAllowed sets field value
func (o *TransferInitializationV1Request) SetBackupGatewaysAllowed(v []string) {
	o.BackupGatewaysAllowed = v
}

// GetRecipientLedgerAssetID returns the RecipientLedgerAssetID field value
func (o *TransferInitializationV1Request) GetRecipientLedgerAssetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientLedgerAssetID
}

// GetRecipientLedgerAssetIDOk returns a tuple with the RecipientLedgerAssetID field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetRecipientLedgerAssetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientLedgerAssetID, true
}

// SetRecipientLedgerAssetID sets field value
func (o *TransferInitializationV1Request) SetRecipientLedgerAssetID(v string) {
	o.RecipientLedgerAssetID = v
}

// GetSourceLedgerAssetID returns the SourceLedgerAssetID field value
func (o *TransferInitializationV1Request) GetSourceLedgerAssetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLedgerAssetID
}

// GetSourceLedgerAssetIDOk returns a tuple with the SourceLedgerAssetID field value
// and a boolean to check if the value has been set.
func (o *TransferInitializationV1Request) GetSourceLedgerAssetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLedgerAssetID, true
}

// SetSourceLedgerAssetID sets field value
func (o *TransferInitializationV1Request) SetSourceLedgerAssetID(v string) {
	o.SourceLedgerAssetID = v
}

func (o TransferInitializationV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferInitializationV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageType"] = o.MessageType
	toSerialize["sessionID"] = o.SessionID
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.DeveloperURN) {
		toSerialize["developerURN"] = o.DeveloperURN
	}
	if !IsNil(o.CredentialProfile) {
		toSerialize["credentialProfile"] = o.CredentialProfile
	}
	toSerialize["payloadProfile"] = o.PayloadProfile
	toSerialize["applicationProfile"] = o.ApplicationProfile
	toSerialize["loggingProfile"] = o.LoggingProfile
	toSerialize["accessControlProfile"] = o.AccessControlProfile
	toSerialize["signature"] = o.Signature
	toSerialize["sourceGatewayPubkey"] = o.SourceGatewayPubkey
	toSerialize["sourceGatewayDltSystem"] = o.SourceGatewayDltSystem
	toSerialize["recipientGatewayPubkey"] = o.RecipientGatewayPubkey
	toSerialize["recipientGatewayDltSystem"] = o.RecipientGatewayDltSystem
	if !IsNil(o.EscrowType) {
		toSerialize["escrowType"] = o.EscrowType
	}
	if !IsNil(o.ExpiryTime) {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if !IsNil(o.MultipleClaimsAllowed) {
		toSerialize["multipleClaimsAllowed"] = o.MultipleClaimsAllowed
	}
	if !IsNil(o.MultipleCancelsAllowed) {
		toSerialize["multipleCancelsAllowed"] = o.MultipleCancelsAllowed
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.SubsequentCalls) {
		toSerialize["subsequentCalls"] = o.SubsequentCalls
	}
	if !IsNil(o.Histories) {
		toSerialize["histories"] = o.Histories
	}
	toSerialize["sequenceNumber"] = o.SequenceNumber
	toSerialize["sourceBasePath"] = o.SourceBasePath
	toSerialize["recipientBasePath"] = o.RecipientBasePath
	toSerialize["maxRetries"] = o.MaxRetries
	toSerialize["maxTimeout"] = o.MaxTimeout
	toSerialize["backupGatewaysAllowed"] = o.BackupGatewaysAllowed
	toSerialize["recipientLedgerAssetID"] = o.RecipientLedgerAssetID
	toSerialize["sourceLedgerAssetID"] = o.SourceLedgerAssetID
	return toSerialize, nil
}

type NullableTransferInitializationV1Request struct {
	value *TransferInitializationV1Request
	isSet bool
}

func (v NullableTransferInitializationV1Request) Get() *TransferInitializationV1Request {
	return v.value
}

func (v *NullableTransferInitializationV1Request) Set(val *TransferInitializationV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferInitializationV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferInitializationV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferInitializationV1Request(val *TransferInitializationV1Request) *NullableTransferInitializationV1Request {
	return &NullableTransferInitializationV1Request{value: val, isSet: true}
}

func (v NullableTransferInitializationV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferInitializationV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


