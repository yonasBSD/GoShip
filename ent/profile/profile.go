// Code generated by ent, DO NOT EDIT.

package profile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the profile type in the database.
	Label = "profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldFullyOnboarded holds the string denoting the fully_onboarded field in the database.
	FieldFullyOnboarded = "fully_onboarded"
	// FieldPhoneNumberE164 holds the string denoting the phone_number_e164 field in the database.
	FieldPhoneNumberE164 = "phone_number_e164"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldPhoneVerified holds the string denoting the phone_verified field in the database.
	FieldPhoneVerified = "phone_verified"
	// FieldStripeID holds the string denoting the stripe_id field in the database.
	FieldStripeID = "stripe_id"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgePhotos holds the string denoting the photos edge name in mutations.
	EdgePhotos = "photos"
	// EdgeProfileImage holds the string denoting the profile_image edge name in mutations.
	EdgeProfileImage = "profile_image"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// EdgeFcmPushSubscriptions holds the string denoting the fcm_push_subscriptions edge name in mutations.
	EdgeFcmPushSubscriptions = "fcm_push_subscriptions"
	// EdgePwaPushSubscriptions holds the string denoting the pwa_push_subscriptions edge name in mutations.
	EdgePwaPushSubscriptions = "pwa_push_subscriptions"
	// EdgeNotificationPermissions holds the string denoting the notification_permissions edge name in mutations.
	EdgeNotificationPermissions = "notification_permissions"
	// EdgeNotificationTimes holds the string denoting the notification_times edge name in mutations.
	EdgeNotificationTimes = "notification_times"
	// EdgePhoneVerificationCode holds the string denoting the phone_verification_code edge name in mutations.
	EdgePhoneVerificationCode = "phone_verification_code"
	// EdgeSentEmails holds the string denoting the sent_emails edge name in mutations.
	EdgeSentEmails = "sent_emails"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSubscription holds the string denoting the subscription edge name in mutations.
	EdgeSubscription = "subscription"
	// Table holds the table name of the profile in the database.
	Table = "profiles"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "profile_friends"
	// PhotosTable is the table that holds the photos relation/edge.
	PhotosTable = "images"
	// PhotosInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	PhotosInverseTable = "images"
	// PhotosColumn is the table column denoting the photos relation/edge.
	PhotosColumn = "profile_photos"
	// ProfileImageTable is the table that holds the profile_image relation/edge.
	ProfileImageTable = "profiles"
	// ProfileImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ProfileImageInverseTable = "images"
	// ProfileImageColumn is the table column denoting the profile_image relation/edge.
	ProfileImageColumn = "profile_profile_image"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "profile_notifications"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "profile_invitations"
	// FcmPushSubscriptionsTable is the table that holds the fcm_push_subscriptions relation/edge.
	FcmPushSubscriptionsTable = "fcm_subscriptions"
	// FcmPushSubscriptionsInverseTable is the table name for the FCMSubscriptions entity.
	// It exists in this package in order to avoid circular dependency with the "fcmsubscriptions" package.
	FcmPushSubscriptionsInverseTable = "fcm_subscriptions"
	// FcmPushSubscriptionsColumn is the table column denoting the fcm_push_subscriptions relation/edge.
	FcmPushSubscriptionsColumn = "profile_id"
	// PwaPushSubscriptionsTable is the table that holds the pwa_push_subscriptions relation/edge.
	PwaPushSubscriptionsTable = "pwa_push_subscriptions"
	// PwaPushSubscriptionsInverseTable is the table name for the PwaPushSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "pwapushsubscription" package.
	PwaPushSubscriptionsInverseTable = "pwa_push_subscriptions"
	// PwaPushSubscriptionsColumn is the table column denoting the pwa_push_subscriptions relation/edge.
	PwaPushSubscriptionsColumn = "profile_id"
	// NotificationPermissionsTable is the table that holds the notification_permissions relation/edge.
	NotificationPermissionsTable = "notification_permissions"
	// NotificationPermissionsInverseTable is the table name for the NotificationPermission entity.
	// It exists in this package in order to avoid circular dependency with the "notificationpermission" package.
	NotificationPermissionsInverseTable = "notification_permissions"
	// NotificationPermissionsColumn is the table column denoting the notification_permissions relation/edge.
	NotificationPermissionsColumn = "profile_id"
	// NotificationTimesTable is the table that holds the notification_times relation/edge.
	NotificationTimesTable = "notification_times"
	// NotificationTimesInverseTable is the table name for the NotificationTime entity.
	// It exists in this package in order to avoid circular dependency with the "notificationtime" package.
	NotificationTimesInverseTable = "notification_times"
	// NotificationTimesColumn is the table column denoting the notification_times relation/edge.
	NotificationTimesColumn = "profile_id"
	// PhoneVerificationCodeTable is the table that holds the phone_verification_code relation/edge.
	PhoneVerificationCodeTable = "phone_verification_codes"
	// PhoneVerificationCodeInverseTable is the table name for the PhoneVerificationCode entity.
	// It exists in this package in order to avoid circular dependency with the "phoneverificationcode" package.
	PhoneVerificationCodeInverseTable = "phone_verification_codes"
	// PhoneVerificationCodeColumn is the table column denoting the phone_verification_code relation/edge.
	PhoneVerificationCodeColumn = "profile_id"
	// SentEmailsTable is the table that holds the sent_emails relation/edge.
	SentEmailsTable = "sent_emails"
	// SentEmailsInverseTable is the table name for the SentEmail entity.
	// It exists in this package in order to avoid circular dependency with the "sentemail" package.
	SentEmailsInverseTable = "sent_emails"
	// SentEmailsColumn is the table column denoting the sent_emails relation/edge.
	SentEmailsColumn = "profile_sent_emails"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "profiles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_profile"
	// SubscriptionTable is the table that holds the subscription relation/edge. The primary key declared below.
	SubscriptionTable = "monthly_subscription_benefactors"
	// SubscriptionInverseTable is the table name for the MonthlySubscription entity.
	// It exists in this package in order to avoid circular dependency with the "monthlysubscription" package.
	SubscriptionInverseTable = "monthly_subscriptions"
)

// Columns holds all SQL columns for profile fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldBio,
	FieldBirthdate,
	FieldAge,
	FieldFullyOnboarded,
	FieldPhoneNumberE164,
	FieldCountryCode,
	FieldPhoneVerified,
	FieldStripeID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "profiles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_profile_image",
	"user_profile",
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"profile_id", "friend_id"}
	// SubscriptionPrimaryKey and SubscriptionColumn2 are the table columns denoting the
	// primary key for the subscription relation (M2M).
	SubscriptionPrimaryKey = []string{"monthly_subscription_id", "profile_id"}
)

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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultFullyOnboarded holds the default value on creation for the "fully_onboarded" field.
	DefaultFullyOnboarded bool
)

// OrderOption defines the ordering options for the Profile queries.
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

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// ByBirthdate orders the results by the birthdate field.
func ByBirthdate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthdate, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByFullyOnboarded orders the results by the fully_onboarded field.
func ByFullyOnboarded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullyOnboarded, opts...).ToFunc()
}

// ByPhoneNumberE164 orders the results by the phone_number_e164 field.
func ByPhoneNumberE164(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumberE164, opts...).ToFunc()
}

// ByCountryCode orders the results by the country_code field.
func ByCountryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryCode, opts...).ToFunc()
}

// ByPhoneVerified orders the results by the phone_verified field.
func ByPhoneVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneVerified, opts...).ToFunc()
}

// ByStripeID orders the results by the stripe_id field.
func ByStripeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripeID, opts...).ToFunc()
}

// ByFriendsCount orders the results by friends count.
func ByFriendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendsStep(), opts...)
	}
}

// ByFriends orders the results by friends terms.
func ByFriends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPhotosCount orders the results by photos count.
func ByPhotosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPhotosStep(), opts...)
	}
}

// ByPhotos orders the results by photos terms.
func ByPhotos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPhotosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProfileImageField orders the results by profile_image field.
func ByProfileImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileImageStep(), sql.OrderByField(field, opts...))
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInvitationsCount orders the results by invitations count.
func ByInvitationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvitationsStep(), opts...)
	}
}

// ByInvitations orders the results by invitations terms.
func ByInvitations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvitationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFcmPushSubscriptionsCount orders the results by fcm_push_subscriptions count.
func ByFcmPushSubscriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFcmPushSubscriptionsStep(), opts...)
	}
}

// ByFcmPushSubscriptions orders the results by fcm_push_subscriptions terms.
func ByFcmPushSubscriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFcmPushSubscriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPwaPushSubscriptionsCount orders the results by pwa_push_subscriptions count.
func ByPwaPushSubscriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPwaPushSubscriptionsStep(), opts...)
	}
}

// ByPwaPushSubscriptions orders the results by pwa_push_subscriptions terms.
func ByPwaPushSubscriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPwaPushSubscriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationPermissionsCount orders the results by notification_permissions count.
func ByNotificationPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationPermissionsStep(), opts...)
	}
}

// ByNotificationPermissions orders the results by notification_permissions terms.
func ByNotificationPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationTimesCount orders the results by notification_times count.
func ByNotificationTimesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationTimesStep(), opts...)
	}
}

// ByNotificationTimes orders the results by notification_times terms.
func ByNotificationTimes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationTimesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPhoneVerificationCodeCount orders the results by phone_verification_code count.
func ByPhoneVerificationCodeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPhoneVerificationCodeStep(), opts...)
	}
}

// ByPhoneVerificationCode orders the results by phone_verification_code terms.
func ByPhoneVerificationCode(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPhoneVerificationCodeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySentEmailsCount orders the results by sent_emails count.
func BySentEmailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSentEmailsStep(), opts...)
	}
}

// BySentEmails orders the results by sent_emails terms.
func BySentEmails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSentEmailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// BySubscriptionCount orders the results by subscription count.
func BySubscriptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscriptionStep(), opts...)
	}
}

// BySubscription orders the results by subscription terms.
func BySubscription(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFriendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
	)
}
func newPhotosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PhotosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PhotosTable, PhotosColumn),
	)
}
func newProfileImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProfileImageTable, ProfileImageColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
func newInvitationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InvitationsTable, InvitationsColumn),
	)
}
func newFcmPushSubscriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FcmPushSubscriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FcmPushSubscriptionsTable, FcmPushSubscriptionsColumn),
	)
}
func newPwaPushSubscriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PwaPushSubscriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PwaPushSubscriptionsTable, PwaPushSubscriptionsColumn),
	)
}
func newNotificationPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationPermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationPermissionsTable, NotificationPermissionsColumn),
	)
}
func newNotificationTimesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationTimesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationTimesTable, NotificationTimesColumn),
	)
}
func newPhoneVerificationCodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PhoneVerificationCodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PhoneVerificationCodeTable, PhoneVerificationCodeColumn),
	)
}
func newSentEmailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SentEmailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SentEmailsTable, SentEmailsColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
func newSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SubscriptionTable, SubscriptionPrimaryKey...),
	)
}
