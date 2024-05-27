// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/unkey-go/internal/utils"
)

// Variant - The algorithm for hashing and encoding, currently only sha256 and base64 are supported
type Variant string

const (
	VariantSha256Base64 Variant = "sha256_base64"
)

func (e Variant) ToPointer() *Variant {
	return &e
}

func (e *Variant) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sha256_base64":
		*e = Variant(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Variant: %v", v)
	}
}

type Hash struct {
	// The hashed and encoded key
	Value string `json:"value"`
	// The algorithm for hashing and encoding, currently only sha256 and base64 are supported
	Variant Variant `json:"variant"`
}

func (o *Hash) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Hash) GetVariant() Variant {
	if o == nil {
		return Variant("")
	}
	return o.Variant
}

// V1MigrationsCreateKeysInterval - Unkey will automatically refill verifications at the set interval.
type V1MigrationsCreateKeysInterval string

const (
	V1MigrationsCreateKeysIntervalDaily   V1MigrationsCreateKeysInterval = "daily"
	V1MigrationsCreateKeysIntervalMonthly V1MigrationsCreateKeysInterval = "monthly"
)

func (e V1MigrationsCreateKeysInterval) ToPointer() *V1MigrationsCreateKeysInterval {
	return &e
}

func (e *V1MigrationsCreateKeysInterval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "monthly":
		*e = V1MigrationsCreateKeysInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for V1MigrationsCreateKeysInterval: %v", v)
	}
}

// V1MigrationsCreateKeysRefill - Unkey enables you to refill verifications for each key at regular intervals.
type V1MigrationsCreateKeysRefill struct {
	// Unkey will automatically refill verifications at the set interval.
	Interval V1MigrationsCreateKeysInterval `json:"interval"`
	// The number of verifications to refill for each occurrence is determined individually for each key.
	Amount int64 `json:"amount"`
}

func (o *V1MigrationsCreateKeysRefill) GetInterval() V1MigrationsCreateKeysInterval {
	if o == nil {
		return V1MigrationsCreateKeysInterval("")
	}
	return o.Interval
}

func (o *V1MigrationsCreateKeysRefill) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

// V1MigrationsCreateKeysType - Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
//
// https://unkey.dev/docs/features/ratelimiting - Learn more
type V1MigrationsCreateKeysType string

const (
	V1MigrationsCreateKeysTypeFast       V1MigrationsCreateKeysType = "fast"
	V1MigrationsCreateKeysTypeConsistent V1MigrationsCreateKeysType = "consistent"
)

func (e V1MigrationsCreateKeysType) ToPointer() *V1MigrationsCreateKeysType {
	return &e
}

func (e *V1MigrationsCreateKeysType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fast":
		fallthrough
	case "consistent":
		*e = V1MigrationsCreateKeysType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for V1MigrationsCreateKeysType: %v", v)
	}
}

// V1MigrationsCreateKeysRatelimit - Unkey comes with per-key ratelimiting out of the box.
type V1MigrationsCreateKeysRatelimit struct {
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	Type *V1MigrationsCreateKeysType `default:"fast" json:"type"`
	// The total amount of burstable requests.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval"`
}

func (v V1MigrationsCreateKeysRatelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V1MigrationsCreateKeysRatelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V1MigrationsCreateKeysRatelimit) GetType() *V1MigrationsCreateKeysType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *V1MigrationsCreateKeysRatelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *V1MigrationsCreateKeysRatelimit) GetRefillRate() int64 {
	if o == nil {
		return 0
	}
	return o.RefillRate
}

func (o *V1MigrationsCreateKeysRatelimit) GetRefillInterval() int64 {
	if o == nil {
		return 0
	}
	return o.RefillInterval
}

type RequestBody struct {
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
	Hash Hash    `json:"hash"`
	// The first 4 characters of the key. If a prefix is used, it should be the prefix plus 4 characters.
	Start *string `default:"" json:"start"`
	// Your user’s Id. This will provide a link between Unkey and your customer record.
	// When validating a key, we will return this back to you, so you can clearly identify your user from their api key.
	OwnerID *string `json:"ownerId,omitempty"`
	// This is a place for dynamic meta data, anything that feels useful for you should go here
	Meta map[string]any `json:"meta,omitempty"`
	// A list of roles that this key should have. If the role does not exist, an error is thrown
	Roles []string `json:"roles,omitempty"`
	// You can auto expire keys by providing a unix timestamp in milliseconds. Once Keys expire they will automatically be disabled and are no longer valid unless you enable them again.
	Expires *int64 `json:"expires,omitempty"`
	// You can limit the number of requests a key can make. Once a key reaches 0 remaining requests, it will automatically be disabled and is no longer valid unless you update it.
	Remaining *int64 `json:"remaining,omitempty"`
	// Unkey enables you to refill verifications for each key at regular intervals.
	Refill *V1MigrationsCreateKeysRefill `json:"refill,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit *V1MigrationsCreateKeysRatelimit `json:"ratelimit,omitempty"`
	// Sets if key is enabled or disabled. Disabled keys are not valid.
	Enabled *bool `default:"true" json:"enabled"`
	// Environments allow you to divide your keyspace.
	//
	// Some applications like Stripe, Clerk, WorkOS and others have a concept of "live" and "test" keys to
	// give the developer a way to develop their own application without the risk of modifying real world
	// resources.
	//
	// When you set an environment, we will return it back to you when validating the key, so you can
	// handle it correctly.
	//
	Environment *string `json:"environment,omitempty"`
}

func (r RequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestBody) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *RequestBody) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *RequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RequestBody) GetHash() Hash {
	if o == nil {
		return Hash{}
	}
	return o.Hash
}

func (o *RequestBody) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *RequestBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *RequestBody) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *RequestBody) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *RequestBody) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *RequestBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *RequestBody) GetRefill() *V1MigrationsCreateKeysRefill {
	if o == nil {
		return nil
	}
	return o.Refill
}

func (o *RequestBody) GetRatelimit() *V1MigrationsCreateKeysRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *RequestBody) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RequestBody) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

// V1MigrationsCreateKeysResponseBody - The key ids of all created keys
type V1MigrationsCreateKeysResponseBody struct {
	// The ids of the keys. This is not a secret and can be stored as a reference if you wish. You need the keyId to update or delete a key later.
	KeyIds []string `json:"keyIds"`
}

func (o *V1MigrationsCreateKeysResponseBody) GetKeyIds() []string {
	if o == nil {
		return []string{}
	}
	return o.KeyIds
}
