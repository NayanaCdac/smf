/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UpdateEventsSubscResponse201 struct {
	Events []AfEventSubscription `json:"events" yaml:"events" bson:"events" mapstructure:"Events"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotifUri   string          `json:"notifUri,omitempty" yaml:"notifUri" bson:"notifUri" mapstructure:"NotifUri"`
	UsgThres   *UsageThreshold `json:"usgThres,omitempty" yaml:"usgThres" bson:"usgThres" mapstructure:"UsgThres"`
	AccessType AccessType      `json:"accessType,omitempty" yaml:"accessType" bson:"accessType" mapstructure:"AccessType"`
	AnGwAddr   *AnGwAddress    `json:"anGwAddr,omitempty" yaml:"anGwAddr" bson:"anGwAddr" mapstructure:"AnGwAddr"`
	// string providing an URI formatted according to IETF RFC 3986.
	EvSubsUri                 string                       `json:"evSubsUri" yaml:"evSubsUri" bson:"evSubsUri" mapstructure:"EvSubsUri"`
	EvNotifs                  []AfEventNotification        `json:"evNotifs" yaml:"evNotifs" bson:"evNotifs" mapstructure:"EvNotifs"`
	FailedResourcAllocReports []ResourcesAllocationInfo    `json:"failedResourcAllocReports,omitempty" yaml:"failedResourcAllocReports" bson:"failedResourcAllocReports" mapstructure:"FailedResourcAllocReports"`
	PlmnId                    *PlmnId                      `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	QncReports                []QosNotificationControlInfo `json:"qncReports,omitempty" yaml:"qncReports" bson:"qncReports" mapstructure:"QncReports"`
	RatType                   RatType                      `json:"ratType,omitempty" yaml:"ratType" bson:"ratType" mapstructure:"RatType"`
	UsgRep                    *AccumulatedUsage            `json:"usgRep,omitempty" yaml:"usgRep" bson:"usgRep" mapstructure:"UsgRep"`
}
