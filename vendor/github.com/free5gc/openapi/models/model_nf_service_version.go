/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type NfServiceVersion struct {
	ApiVersionInUri string     `json:"apiVersionInUri" yaml:"apiVersionInUri" bson:"apiVersionInUri" mapstructure:"ApiVersionInUri"`
	ApiFullVersion  string     `json:"apiFullVersion" yaml:"apiFullVersion" bson:"apiFullVersion" mapstructure:"ApiFullVersion"`
	Expiry          *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry" mapstructure:"Expiry"`
}
