package storage

import (
	"net/mail"
	"time"

	"github.com/jhillyerd/enmime"
)

// Message data excluding physical attachments
//
// swagger:model Message
type Message struct {
	// Database ID
	ID string
	// Message ID
	MessageID string
	// Read status
	Read bool
	// From address
	From *mail.Address
	// To addresses
	To []*mail.Address
	// Cc addresses
	Cc []*mail.Address
	// Bcc addresses
	Bcc []*mail.Address
	// ReplyTo addresses
	ReplyTo []*mail.Address
	// Return-Path
	ReturnPath string
	// Message subject
	Subject string
	// Message date if set, else date received
	Date time.Time
	// Message tags
	Tags []string
	// Message body text
	Text string
	// Message body HTML
	HTML string
	// Message size in bytes
	Size int
	// Inline message attachments
	Inline []Attachment
	// Message attachments
	Attachments []Attachment
}

// Attachment struct for inline and attachments
//
// swagger:model Attachment
type Attachment struct {
	// Attachment part ID
	PartID string
	// File name
	FileName string
	// Content type
	ContentType string
	// Content ID
	ContentID string
	// Size in bytes
	Size int
}

// MessageSummary struct for frontend messages
//
// swagger:model MessageSummary
type MessageSummary struct {
	// Database ID
	ID string
	// Read status
	Read bool
	// From address
	From *mail.Address
	// To address
	To []*mail.Address
	// Cc addresses
	Cc []*mail.Address
	// Bcc addresses
	Bcc []*mail.Address
	// Email subject
	Subject string
	// Created time
	Created time.Time
	// Message tags
	Tags []string
	// Message size in bytes (total)
	Size int
	// Whether the message has any attachments
	Attachments int
}

// MailboxStats struct for quick mailbox total/read lookups
type MailboxStats struct {
	Total  int
	Unread int
	Tags   []string
}

// AttachmentSummary returns a summary of the attachment without any binary data
func AttachmentSummary(a *enmime.Part) Attachment {
	o := Attachment{}
	o.PartID = a.PartID
	o.FileName = a.FileName
	if o.FileName == "" {
		o.FileName = a.ContentID
	}
	o.ContentType = a.ContentType
	o.ContentID = a.ContentID
	o.Size = len(a.Content)

	return o
}
