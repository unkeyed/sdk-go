// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/unkey-go/internal/utils"
)

// DeprecatedCreateKeyType - Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
//
// https://unkey.dev/docs/features/ratelimiting - Learn more
type DeprecatedCreateKeyType string

const (
	DeprecatedCreateKeyTypeFast       DeprecatedCreateKeyType = "fast"
	DeprecatedCreateKeyTypeConsistent DeprecatedCreateKeyType = "consistent"
)

func (e DeprecatedCreateKeyType) ToPointer() *DeprecatedCreateKeyType {
	return &e
}

func (e *DeprecatedCreateKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fast":
		fallthrough
	case "consistent":
		*e = DeprecatedCreateKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeprecatedCreateKeyType: %v", v)
	}
}

// DeprecatedCreateKeyRatelimit - Unkey comes with per-key ratelimiting out of the box.
type DeprecatedCreateKeyRatelimit struct {
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	Type *DeprecatedCreateKeyType `default:"fast" json:"type"`
	// The total amount of burstable requests.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval"`
}

func (d DeprecatedCreateKeyRatelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeprecatedCreateKeyRatelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeprecatedCreateKeyRatelimit) GetType() *DeprecatedCreateKeyType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *DeprecatedCreateKeyRatelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *DeprecatedCreateKeyRatelimit) GetRefillRate() int64 {
	if o == nil {
		return 0
	}
	return o.RefillRate
}

func (o *DeprecatedCreateKeyRatelimit) GetRefillInterval() int64 {
	if o == nil {
		return 0
	}
	return o.RefillInterval
}

type DeprecatedCreateKeyRequestBody struct {
	// Choose an `API` where this key should be created.
	APIID string `json:"apiId"`
	// To make it easier for your users to understand which product an api key belongs to, you can add prefix them.
	//
	// For example Stripe famously prefixes their customer ids with cus_ or their api keys with sk_live_.
	//
	// The underscore is automatically added if you are defining a prefix, for example: "prefix": "abc" will result in a key like abc_xxxxxxxxx
	//
	Prefix *string `json:"prefix,omitempty"`
	// The name for your Key. This is not customer facing.
	Name *string `json:"name,omitempty"`
	// The byte length used to generate your key determines its entropy as well as its length. Higher is better, but keys become longer and more annoying to handle. The default is 16 bytes, or 2^^128 possible combinations.
	ByteLength *int64 `default:"16" json:"byteLength"`
	// Your user’s Id. This will provide a link between Unkey and your customer record.
	// When validating a key, we will return this back to you, so you can clearly identify your user from their api key.
	OwnerID *string `json:"ownerId,omitempty"`
	// This is a place for dynamic meta data, anything that feels useful for you should go here
	Meta map[string]interface{} `json:"meta,omitempty"`
	// You can auto expire keys by providing a unix timestamp in milliseconds. Once Keys expire they will automatically be disabled and are no longer valid unless you enable them again.
	Expires *int64 `json:"expires,omitempty"`
	// You can limit the number of requests a key can make. Once a key reaches 0 remaining requests, it will automatically be disabled and is no longer valid unless you update it.
	Remaining *int64 `json:"remaining,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit *DeprecatedCreateKeyRatelimit `json:"ratelimit,omitempty"`
}

func (d DeprecatedCreateKeyRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeprecatedCreateKeyRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeprecatedCreateKeyRequestBody) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *DeprecatedCreateKeyRequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *DeprecatedCreateKeyRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DeprecatedCreateKeyRequestBody) GetByteLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ByteLength
}

func (o *DeprecatedCreateKeyRequestBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *DeprecatedCreateKeyRequestBody) GetMeta() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *DeprecatedCreateKeyRequestBody) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *DeprecatedCreateKeyRequestBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *DeprecatedCreateKeyRequestBody) GetRatelimit() *DeprecatedCreateKeyRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

// DeprecatedCreateKeyResponseBody - The configuration for an api
type DeprecatedCreateKeyResponseBody struct {
	// The id of the key. This is not a secret and can be stored as a reference if you wish. You need the keyId to update or delete a key later.
	KeyID string `json:"keyId"`
	// The newly created api key, do not store this on your own system but pass it along to your user.
	Key string `json:"key"`
}

func (o *DeprecatedCreateKeyResponseBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *DeprecatedCreateKeyResponseBody) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}
