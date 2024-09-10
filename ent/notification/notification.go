// Code generated by ent, DO NOT EDIT.

package notification

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the notification type in the database.
	Label = "notification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldRead holds the string denoting the read field in the database.
	FieldRead = "read"
	// FieldReadAt holds the string denoting the read_at field in the database.
	FieldReadAt = "read_at"
	// FieldProfileIDWhoCausedNotification holds the string denoting the profile_id_who_caused_notification field in the database.
	FieldProfileIDWhoCausedNotification = "profile_id_who_caused_notification"
	// FieldResourceIDTiedToNotif holds the string denoting the resource_id_tied_to_notif field in the database.
	FieldResourceIDTiedToNotif = "resource_id_tied_to_notif"
	// FieldReadInNotificationsCenter holds the string denoting the read_in_notifications_center field in the database.
	FieldReadInNotificationsCenter = "read_in_notifications_center"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// Table holds the table name of the notification in the database.
	Table = "notifications"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "notifications"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "profile_notifications"
)

// Columns holds all SQL columns for notification fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldTitle,
	FieldText,
	FieldLink,
	FieldRead,
	FieldReadAt,
	FieldProfileIDWhoCausedNotification,
	FieldResourceIDTiedToNotif,
	FieldReadInNotificationsCenter,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notifications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_notifications",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/mikestefanello/pagoda/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultRead holds the default value on creation for the "read" field.
	DefaultRead bool
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeNewPrivateMessage             Type = "new_private_message"
	TypeMutualQuestionAnswered        Type = "mutual_question_answered"
	TypeConnectionEngagedWithQuestion Type = "connection_engaged_with_question"
	TypeIncrementNumUnseenMsg         Type = "increment_num_unseen_msg"
	TypeDecrementNumUnseenMsg         Type = "decrement_num_unseen_msg"
	TypeUpdateNumNotifs               Type = "update_num_notifs"
	TypeConnectionRequestAccepted     Type = "connection_request_accepted"
	TypePlatformUpdate                Type = "platform_update"
	TypeConnectionReactedToAnswer     Type = "connection_reacted_to_answer"
	TypeCommittedRelationshipRequest  Type = "committed_relationship_request"
	TypeBreakup                       Type = "breakup"
	TypeContactRequest                Type = "contact_request"
	TypeNewHomefeedQaObject           Type = "new_homefeed_qa_object"
	TypePaymentFailed                 Type = "payment_failed"
	TypeDailyConversationReminder     Type = "daily_conversation_reminder"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeNewPrivateMessage, TypeMutualQuestionAnswered, TypeConnectionEngagedWithQuestion, TypeIncrementNumUnseenMsg, TypeDecrementNumUnseenMsg, TypeUpdateNumNotifs, TypeConnectionRequestAccepted, TypePlatformUpdate, TypeConnectionReactedToAnswer, TypeCommittedRelationshipRequest, TypeBreakup, TypeContactRequest, TypeNewHomefeedQaObject, TypePaymentFailed, TypeDailyConversationReminder:
		return nil
	default:
		return fmt.Errorf("notification: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Notification queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByLink orders the results by the link field.
func ByLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLink, opts...).ToFunc()
}

// ByRead orders the results by the read field.
func ByRead(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRead, opts...).ToFunc()
}

// ByReadAt orders the results by the read_at field.
func ByReadAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadAt, opts...).ToFunc()
}

// ByProfileIDWhoCausedNotification orders the results by the profile_id_who_caused_notification field.
func ByProfileIDWhoCausedNotification(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfileIDWhoCausedNotification, opts...).ToFunc()
}

// ByResourceIDTiedToNotif orders the results by the resource_id_tied_to_notif field.
func ByResourceIDTiedToNotif(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResourceIDTiedToNotif, opts...).ToFunc()
}

// ByReadInNotificationsCenter orders the results by the read_in_notifications_center field.
func ByReadInNotificationsCenter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadInNotificationsCenter, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
	)
}