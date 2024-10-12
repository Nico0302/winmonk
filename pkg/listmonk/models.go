// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package listmonk

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type BounceType string

const (
	BounceTypeSoft      BounceType = "soft"
	BounceTypeHard      BounceType = "hard"
	BounceTypeComplaint BounceType = "complaint"
)

func (e *BounceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BounceType(s)
	case string:
		*e = BounceType(s)
	default:
		return fmt.Errorf("unsupported scan type for BounceType: %T", src)
	}
	return nil
}

type NullBounceType struct {
	BounceType BounceType
	Valid      bool // Valid is true if BounceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBounceType) Scan(value interface{}) error {
	if value == nil {
		ns.BounceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BounceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBounceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BounceType), nil
}

type CampaignStatus string

const (
	CampaignStatusDraft     CampaignStatus = "draft"
	CampaignStatusRunning   CampaignStatus = "running"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusPaused    CampaignStatus = "paused"
	CampaignStatusCancelled CampaignStatus = "cancelled"
	CampaignStatusFinished  CampaignStatus = "finished"
)

func (e *CampaignStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CampaignStatus(s)
	case string:
		*e = CampaignStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for CampaignStatus: %T", src)
	}
	return nil
}

type NullCampaignStatus struct {
	CampaignStatus CampaignStatus
	Valid          bool // Valid is true if CampaignStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCampaignStatus) Scan(value interface{}) error {
	if value == nil {
		ns.CampaignStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CampaignStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCampaignStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CampaignStatus), nil
}

type CampaignType string

const (
	CampaignTypeRegular CampaignType = "regular"
	CampaignTypeOptin   CampaignType = "optin"
)

func (e *CampaignType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CampaignType(s)
	case string:
		*e = CampaignType(s)
	default:
		return fmt.Errorf("unsupported scan type for CampaignType: %T", src)
	}
	return nil
}

type NullCampaignType struct {
	CampaignType CampaignType
	Valid        bool // Valid is true if CampaignType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCampaignType) Scan(value interface{}) error {
	if value == nil {
		ns.CampaignType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CampaignType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCampaignType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CampaignType), nil
}

type ContentType string

const (
	ContentTypeRichtext ContentType = "richtext"
	ContentTypeHtml     ContentType = "html"
	ContentTypePlain    ContentType = "plain"
	ContentTypeMarkdown ContentType = "markdown"
)

func (e *ContentType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ContentType(s)
	case string:
		*e = ContentType(s)
	default:
		return fmt.Errorf("unsupported scan type for ContentType: %T", src)
	}
	return nil
}

type NullContentType struct {
	ContentType ContentType
	Valid       bool // Valid is true if ContentType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullContentType) Scan(value interface{}) error {
	if value == nil {
		ns.ContentType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ContentType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullContentType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ContentType), nil
}

type ListOptin string

const (
	ListOptinSingle ListOptin = "single"
	ListOptinDouble ListOptin = "double"
)

func (e *ListOptin) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ListOptin(s)
	case string:
		*e = ListOptin(s)
	default:
		return fmt.Errorf("unsupported scan type for ListOptin: %T", src)
	}
	return nil
}

type NullListOptin struct {
	ListOptin ListOptin
	Valid     bool // Valid is true if ListOptin is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullListOptin) Scan(value interface{}) error {
	if value == nil {
		ns.ListOptin, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ListOptin.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullListOptin) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ListOptin), nil
}

type ListType string

const (
	ListTypePublic    ListType = "public"
	ListTypePrivate   ListType = "private"
	ListTypeTemporary ListType = "temporary"
)

func (e *ListType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ListType(s)
	case string:
		*e = ListType(s)
	default:
		return fmt.Errorf("unsupported scan type for ListType: %T", src)
	}
	return nil
}

type NullListType struct {
	ListType ListType
	Valid    bool // Valid is true if ListType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullListType) Scan(value interface{}) error {
	if value == nil {
		ns.ListType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ListType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullListType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ListType), nil
}

type SubscriberStatus string

const (
	SubscriberStatusEnabled     SubscriberStatus = "enabled"
	SubscriberStatusDisabled    SubscriberStatus = "disabled"
	SubscriberStatusBlocklisted SubscriberStatus = "blocklisted"
)

func (e *SubscriberStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SubscriberStatus(s)
	case string:
		*e = SubscriberStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for SubscriberStatus: %T", src)
	}
	return nil
}

type NullSubscriberStatus struct {
	SubscriberStatus SubscriberStatus
	Valid            bool // Valid is true if SubscriberStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSubscriberStatus) Scan(value interface{}) error {
	if value == nil {
		ns.SubscriberStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SubscriberStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSubscriberStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.SubscriberStatus), nil
}

type SubscriptionStatus string

const (
	SubscriptionStatusUnconfirmed  SubscriptionStatus = "unconfirmed"
	SubscriptionStatusConfirmed    SubscriptionStatus = "confirmed"
	SubscriptionStatusUnsubscribed SubscriptionStatus = "unsubscribed"
)

func (e *SubscriptionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SubscriptionStatus(s)
	case string:
		*e = SubscriptionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for SubscriptionStatus: %T", src)
	}
	return nil
}

type NullSubscriptionStatus struct {
	SubscriptionStatus SubscriptionStatus
	Valid              bool // Valid is true if SubscriptionStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSubscriptionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.SubscriptionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SubscriptionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSubscriptionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.SubscriptionStatus), nil
}

type TemplateType string

const (
	TemplateTypeCampaign TemplateType = "campaign"
	TemplateTypeTx       TemplateType = "tx"
)

func (e *TemplateType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TemplateType(s)
	case string:
		*e = TemplateType(s)
	default:
		return fmt.Errorf("unsupported scan type for TemplateType: %T", src)
	}
	return nil
}

type NullTemplateType struct {
	TemplateType TemplateType
	Valid        bool // Valid is true if TemplateType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTemplateType) Scan(value interface{}) error {
	if value == nil {
		ns.TemplateType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TemplateType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTemplateType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TemplateType), nil
}

type Bounce struct {
	ID           int32
	SubscriberID int32
	CampaignID   pgtype.Int4
	Type         BounceType
	Source       string
	Meta         []byte
	CreatedAt    pgtype.Timestamptz
}

type Campaign struct {
	ID                int32
	Uuid              pgtype.UUID
	Name              string
	Subject           string
	FromEmail         string
	Body              string
	Altbody           pgtype.Text
	ContentType       ContentType
	SendAt            pgtype.Timestamptz
	Headers           []byte
	Status            CampaignStatus
	Tags              []string
	Type              NullCampaignType
	Messenger         string
	TemplateID        pgtype.Int4
	ToSend            int32
	Sent              int32
	MaxSubscriberID   int32
	LastSubscriberID  int32
	Archive           bool
	ArchiveSlug       pgtype.Text
	ArchiveTemplateID pgtype.Int4
	ArchiveMeta       []byte
	StartedAt         pgtype.Timestamptz
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
}

type CampaignList struct {
	ID         int64
	CampaignID int32
	ListID     pgtype.Int4
	ListName   string
}

type CampaignMedium struct {
	CampaignID pgtype.Int4
	MediaID    pgtype.Int4
	Filename   string
}

type CampaignView struct {
	ID           int64
	CampaignID   int32
	SubscriberID pgtype.Int4
	CreatedAt    pgtype.Timestamptz
}

type Link struct {
	ID        int32
	Uuid      pgtype.UUID
	Url       string
	CreatedAt pgtype.Timestamptz
}

type LinkClick struct {
	ID           int64
	CampaignID   pgtype.Int4
	LinkID       int32
	SubscriberID pgtype.Int4
	CreatedAt    pgtype.Timestamptz
}

type List struct {
	ID          int32
	Uuid        pgtype.UUID
	Name        string
	Type        ListType
	Optin       ListOptin
	Tags        []string
	Description string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type Medium struct {
	ID          int32
	Uuid        pgtype.UUID
	Provider    string
	Filename    string
	ContentType string
	Thumb       string
	Meta        []byte
	CreatedAt   pgtype.Timestamptz
}

type Setting struct {
	Key       string
	Value     []byte
	UpdatedAt pgtype.Timestamptz
}

type Subscriber struct {
	ID        int32
	Uuid      pgtype.UUID
	Email     string
	Name      string
	Attribs   []byte
	Status    SubscriberStatus
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type SubscriberList struct {
	SubscriberID int32
	ListID       int32
	Meta         []byte
	Status       SubscriptionStatus
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type Template struct {
	ID        int32
	Name      string
	Type      TemplateType
	Subject   string
	Body      string
	IsDefault bool
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}
