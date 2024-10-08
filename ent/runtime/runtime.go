// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/mikestefanello/pagoda/ent/emailsubscription"
	"github.com/mikestefanello/pagoda/ent/emailsubscriptiontype"
	"github.com/mikestefanello/pagoda/ent/emojis"
	"github.com/mikestefanello/pagoda/ent/fcmsubscriptions"
	"github.com/mikestefanello/pagoda/ent/filestorage"
	"github.com/mikestefanello/pagoda/ent/image"
	"github.com/mikestefanello/pagoda/ent/imagesize"
	"github.com/mikestefanello/pagoda/ent/invitation"
	"github.com/mikestefanello/pagoda/ent/lastseenonline"
	"github.com/mikestefanello/pagoda/ent/monthlysubscription"
	"github.com/mikestefanello/pagoda/ent/notification"
	"github.com/mikestefanello/pagoda/ent/notificationpermission"
	"github.com/mikestefanello/pagoda/ent/notificationtime"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/phoneverificationcode"
	"github.com/mikestefanello/pagoda/ent/profile"
	"github.com/mikestefanello/pagoda/ent/pwapushsubscription"
	"github.com/mikestefanello/pagoda/ent/schema"
	"github.com/mikestefanello/pagoda/ent/sentemail"
	"github.com/mikestefanello/pagoda/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	emailsubscriptionMixin := schema.EmailSubscription{}.Mixin()
	emailsubscriptionMixinFields0 := emailsubscriptionMixin[0].Fields()
	_ = emailsubscriptionMixinFields0
	emailsubscriptionFields := schema.EmailSubscription{}.Fields()
	_ = emailsubscriptionFields
	// emailsubscriptionDescCreatedAt is the schema descriptor for created_at field.
	emailsubscriptionDescCreatedAt := emailsubscriptionMixinFields0[0].Descriptor()
	// emailsubscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	emailsubscription.DefaultCreatedAt = emailsubscriptionDescCreatedAt.Default.(func() time.Time)
	// emailsubscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	emailsubscriptionDescUpdatedAt := emailsubscriptionMixinFields0[1].Descriptor()
	// emailsubscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	emailsubscription.DefaultUpdatedAt = emailsubscriptionDescUpdatedAt.Default.(func() time.Time)
	// emailsubscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	emailsubscription.UpdateDefaultUpdatedAt = emailsubscriptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// emailsubscriptionDescEmail is the schema descriptor for email field.
	emailsubscriptionDescEmail := emailsubscriptionFields[0].Descriptor()
	// emailsubscription.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	emailsubscription.EmailValidator = emailsubscriptionDescEmail.Validators[0].(func(string) error)
	// emailsubscriptionDescVerified is the schema descriptor for verified field.
	emailsubscriptionDescVerified := emailsubscriptionFields[1].Descriptor()
	// emailsubscription.DefaultVerified holds the default value on creation for the verified field.
	emailsubscription.DefaultVerified = emailsubscriptionDescVerified.Default.(bool)
	// emailsubscriptionDescConfirmationCode is the schema descriptor for confirmation_code field.
	emailsubscriptionDescConfirmationCode := emailsubscriptionFields[2].Descriptor()
	// emailsubscription.ConfirmationCodeValidator is a validator for the "confirmation_code" field. It is called by the builders before save.
	emailsubscription.ConfirmationCodeValidator = emailsubscriptionDescConfirmationCode.Validators[0].(func(string) error)
	emailsubscriptiontypeMixin := schema.EmailSubscriptionType{}.Mixin()
	emailsubscriptiontypeMixinFields0 := emailsubscriptiontypeMixin[0].Fields()
	_ = emailsubscriptiontypeMixinFields0
	emailsubscriptiontypeFields := schema.EmailSubscriptionType{}.Fields()
	_ = emailsubscriptiontypeFields
	// emailsubscriptiontypeDescCreatedAt is the schema descriptor for created_at field.
	emailsubscriptiontypeDescCreatedAt := emailsubscriptiontypeMixinFields0[0].Descriptor()
	// emailsubscriptiontype.DefaultCreatedAt holds the default value on creation for the created_at field.
	emailsubscriptiontype.DefaultCreatedAt = emailsubscriptiontypeDescCreatedAt.Default.(func() time.Time)
	// emailsubscriptiontypeDescUpdatedAt is the schema descriptor for updated_at field.
	emailsubscriptiontypeDescUpdatedAt := emailsubscriptiontypeMixinFields0[1].Descriptor()
	// emailsubscriptiontype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	emailsubscriptiontype.DefaultUpdatedAt = emailsubscriptiontypeDescUpdatedAt.Default.(func() time.Time)
	// emailsubscriptiontype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	emailsubscriptiontype.UpdateDefaultUpdatedAt = emailsubscriptiontypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// emailsubscriptiontypeDescActive is the schema descriptor for active field.
	emailsubscriptiontypeDescActive := emailsubscriptiontypeFields[1].Descriptor()
	// emailsubscriptiontype.DefaultActive holds the default value on creation for the active field.
	emailsubscriptiontype.DefaultActive = emailsubscriptiontypeDescActive.Default.(bool)
	emojisFields := schema.Emojis{}.Fields()
	_ = emojisFields
	// emojisDescUnifiedCode is the schema descriptor for unified_code field.
	emojisDescUnifiedCode := emojisFields[0].Descriptor()
	// emojis.UnifiedCodeValidator is a validator for the "unified_code" field. It is called by the builders before save.
	emojis.UnifiedCodeValidator = emojisDescUnifiedCode.Validators[0].(func(string) error)
	// emojisDescShortcode is the schema descriptor for shortcode field.
	emojisDescShortcode := emojisFields[1].Descriptor()
	// emojis.ShortcodeValidator is a validator for the "shortcode" field. It is called by the builders before save.
	emojis.ShortcodeValidator = emojisDescShortcode.Validators[0].(func(string) error)
	fcmsubscriptionsMixin := schema.FCMSubscriptions{}.Mixin()
	fcmsubscriptionsMixinFields0 := fcmsubscriptionsMixin[0].Fields()
	_ = fcmsubscriptionsMixinFields0
	fcmsubscriptionsFields := schema.FCMSubscriptions{}.Fields()
	_ = fcmsubscriptionsFields
	// fcmsubscriptionsDescCreatedAt is the schema descriptor for created_at field.
	fcmsubscriptionsDescCreatedAt := fcmsubscriptionsMixinFields0[0].Descriptor()
	// fcmsubscriptions.DefaultCreatedAt holds the default value on creation for the created_at field.
	fcmsubscriptions.DefaultCreatedAt = fcmsubscriptionsDescCreatedAt.Default.(func() time.Time)
	// fcmsubscriptionsDescUpdatedAt is the schema descriptor for updated_at field.
	fcmsubscriptionsDescUpdatedAt := fcmsubscriptionsMixinFields0[1].Descriptor()
	// fcmsubscriptions.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fcmsubscriptions.DefaultUpdatedAt = fcmsubscriptionsDescUpdatedAt.Default.(func() time.Time)
	// fcmsubscriptions.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fcmsubscriptions.UpdateDefaultUpdatedAt = fcmsubscriptionsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fcmsubscriptionsDescToken is the schema descriptor for token field.
	fcmsubscriptionsDescToken := fcmsubscriptionsFields[0].Descriptor()
	// fcmsubscriptions.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	fcmsubscriptions.TokenValidator = fcmsubscriptionsDescToken.Validators[0].(func(string) error)
	filestorageMixin := schema.FileStorage{}.Mixin()
	filestorageMixinFields0 := filestorageMixin[0].Fields()
	_ = filestorageMixinFields0
	filestorageFields := schema.FileStorage{}.Fields()
	_ = filestorageFields
	// filestorageDescCreatedAt is the schema descriptor for created_at field.
	filestorageDescCreatedAt := filestorageMixinFields0[0].Descriptor()
	// filestorage.DefaultCreatedAt holds the default value on creation for the created_at field.
	filestorage.DefaultCreatedAt = filestorageDescCreatedAt.Default.(func() time.Time)
	// filestorageDescUpdatedAt is the schema descriptor for updated_at field.
	filestorageDescUpdatedAt := filestorageMixinFields0[1].Descriptor()
	// filestorage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	filestorage.DefaultUpdatedAt = filestorageDescUpdatedAt.Default.(func() time.Time)
	// filestorage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	filestorage.UpdateDefaultUpdatedAt = filestorageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// filestorageDescBucketName is the schema descriptor for bucket_name field.
	filestorageDescBucketName := filestorageFields[0].Descriptor()
	// filestorage.BucketNameValidator is a validator for the "bucket_name" field. It is called by the builders before save.
	filestorage.BucketNameValidator = filestorageDescBucketName.Validators[0].(func(string) error)
	// filestorageDescObjectKey is the schema descriptor for object_key field.
	filestorageDescObjectKey := filestorageFields[1].Descriptor()
	// filestorage.ObjectKeyValidator is a validator for the "object_key" field. It is called by the builders before save.
	filestorage.ObjectKeyValidator = filestorageDescObjectKey.Validators[0].(func(string) error)
	imageMixin := schema.Image{}.Mixin()
	imageMixinFields0 := imageMixin[0].Fields()
	_ = imageMixinFields0
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescCreatedAt is the schema descriptor for created_at field.
	imageDescCreatedAt := imageMixinFields0[0].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the created_at field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
	// imageDescUpdatedAt is the schema descriptor for updated_at field.
	imageDescUpdatedAt := imageMixinFields0[1].Descriptor()
	// image.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	image.DefaultUpdatedAt = imageDescUpdatedAt.Default.(func() time.Time)
	// image.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	image.UpdateDefaultUpdatedAt = imageDescUpdatedAt.UpdateDefault.(func() time.Time)
	imagesizeMixin := schema.ImageSize{}.Mixin()
	imagesizeMixinFields0 := imagesizeMixin[0].Fields()
	_ = imagesizeMixinFields0
	imagesizeFields := schema.ImageSize{}.Fields()
	_ = imagesizeFields
	// imagesizeDescCreatedAt is the schema descriptor for created_at field.
	imagesizeDescCreatedAt := imagesizeMixinFields0[0].Descriptor()
	// imagesize.DefaultCreatedAt holds the default value on creation for the created_at field.
	imagesize.DefaultCreatedAt = imagesizeDescCreatedAt.Default.(func() time.Time)
	// imagesizeDescUpdatedAt is the schema descriptor for updated_at field.
	imagesizeDescUpdatedAt := imagesizeMixinFields0[1].Descriptor()
	// imagesize.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	imagesize.DefaultUpdatedAt = imagesizeDescUpdatedAt.Default.(func() time.Time)
	// imagesize.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	imagesize.UpdateDefaultUpdatedAt = imagesizeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// imagesizeDescWidth is the schema descriptor for width field.
	imagesizeDescWidth := imagesizeFields[1].Descriptor()
	// imagesize.WidthValidator is a validator for the "width" field. It is called by the builders before save.
	imagesize.WidthValidator = imagesizeDescWidth.Validators[0].(func(int) error)
	// imagesizeDescHeight is the schema descriptor for height field.
	imagesizeDescHeight := imagesizeFields[2].Descriptor()
	// imagesize.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	imagesize.HeightValidator = imagesizeDescHeight.Validators[0].(func(int) error)
	invitationMixin := schema.Invitation{}.Mixin()
	invitationMixinFields0 := invitationMixin[0].Fields()
	_ = invitationMixinFields0
	invitationFields := schema.Invitation{}.Fields()
	_ = invitationFields
	// invitationDescCreatedAt is the schema descriptor for created_at field.
	invitationDescCreatedAt := invitationMixinFields0[0].Descriptor()
	// invitation.DefaultCreatedAt holds the default value on creation for the created_at field.
	invitation.DefaultCreatedAt = invitationDescCreatedAt.Default.(func() time.Time)
	// invitationDescUpdatedAt is the schema descriptor for updated_at field.
	invitationDescUpdatedAt := invitationMixinFields0[1].Descriptor()
	// invitation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	invitation.DefaultUpdatedAt = invitationDescUpdatedAt.Default.(func() time.Time)
	// invitation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	invitation.UpdateDefaultUpdatedAt = invitationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// invitationDescInviteeName is the schema descriptor for invitee_name field.
	invitationDescInviteeName := invitationFields[0].Descriptor()
	// invitation.InviteeNameValidator is a validator for the "invitee_name" field. It is called by the builders before save.
	invitation.InviteeNameValidator = invitationDescInviteeName.Validators[0].(func(string) error)
	// invitationDescConfirmationCode is the schema descriptor for confirmation_code field.
	invitationDescConfirmationCode := invitationFields[1].Descriptor()
	// invitation.ConfirmationCodeValidator is a validator for the "confirmation_code" field. It is called by the builders before save.
	invitation.ConfirmationCodeValidator = invitationDescConfirmationCode.Validators[0].(func(string) error)
	lastseenonlineFields := schema.LastSeenOnline{}.Fields()
	_ = lastseenonlineFields
	// lastseenonlineDescSeenAt is the schema descriptor for seen_at field.
	lastseenonlineDescSeenAt := lastseenonlineFields[0].Descriptor()
	// lastseenonline.DefaultSeenAt holds the default value on creation for the seen_at field.
	lastseenonline.DefaultSeenAt = lastseenonlineDescSeenAt.Default.(func() time.Time)
	monthlysubscriptionMixin := schema.MonthlySubscription{}.Mixin()
	monthlysubscriptionMixinFields0 := monthlysubscriptionMixin[0].Fields()
	_ = monthlysubscriptionMixinFields0
	monthlysubscriptionFields := schema.MonthlySubscription{}.Fields()
	_ = monthlysubscriptionFields
	// monthlysubscriptionDescCreatedAt is the schema descriptor for created_at field.
	monthlysubscriptionDescCreatedAt := monthlysubscriptionMixinFields0[0].Descriptor()
	// monthlysubscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	monthlysubscription.DefaultCreatedAt = monthlysubscriptionDescCreatedAt.Default.(func() time.Time)
	// monthlysubscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	monthlysubscriptionDescUpdatedAt := monthlysubscriptionMixinFields0[1].Descriptor()
	// monthlysubscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	monthlysubscription.DefaultUpdatedAt = monthlysubscriptionDescUpdatedAt.Default.(func() time.Time)
	// monthlysubscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	monthlysubscription.UpdateDefaultUpdatedAt = monthlysubscriptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// monthlysubscriptionDescIsActive is the schema descriptor for is_active field.
	monthlysubscriptionDescIsActive := monthlysubscriptionFields[1].Descriptor()
	// monthlysubscription.DefaultIsActive holds the default value on creation for the is_active field.
	monthlysubscription.DefaultIsActive = monthlysubscriptionDescIsActive.Default.(bool)
	// monthlysubscriptionDescPaid is the schema descriptor for paid field.
	monthlysubscriptionDescPaid := monthlysubscriptionFields[2].Descriptor()
	// monthlysubscription.DefaultPaid holds the default value on creation for the paid field.
	monthlysubscription.DefaultPaid = monthlysubscriptionDescPaid.Default.(bool)
	// monthlysubscriptionDescIsTrial is the schema descriptor for is_trial field.
	monthlysubscriptionDescIsTrial := monthlysubscriptionFields[3].Descriptor()
	// monthlysubscription.DefaultIsTrial holds the default value on creation for the is_trial field.
	monthlysubscription.DefaultIsTrial = monthlysubscriptionDescIsTrial.Default.(bool)
	notificationMixin := schema.Notification{}.Mixin()
	notificationHooks := schema.Notification{}.Hooks()
	notification.Hooks[0] = notificationHooks[0]
	notificationMixinFields0 := notificationMixin[0].Fields()
	_ = notificationMixinFields0
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationMixinFields0[0].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationMixinFields0[1].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationDescTitle is the schema descriptor for title field.
	notificationDescTitle := notificationFields[1].Descriptor()
	// notification.DefaultTitle holds the default value on creation for the title field.
	notification.DefaultTitle = notificationDescTitle.Default.(string)
	// notificationDescRead is the schema descriptor for read field.
	notificationDescRead := notificationFields[4].Descriptor()
	// notification.DefaultRead holds the default value on creation for the read field.
	notification.DefaultRead = notificationDescRead.Default.(bool)
	notificationpermissionMixin := schema.NotificationPermission{}.Mixin()
	notificationpermissionMixinFields0 := notificationpermissionMixin[0].Fields()
	_ = notificationpermissionMixinFields0
	notificationpermissionFields := schema.NotificationPermission{}.Fields()
	_ = notificationpermissionFields
	// notificationpermissionDescCreatedAt is the schema descriptor for created_at field.
	notificationpermissionDescCreatedAt := notificationpermissionMixinFields0[0].Descriptor()
	// notificationpermission.DefaultCreatedAt holds the default value on creation for the created_at field.
	notificationpermission.DefaultCreatedAt = notificationpermissionDescCreatedAt.Default.(func() time.Time)
	// notificationpermissionDescUpdatedAt is the schema descriptor for updated_at field.
	notificationpermissionDescUpdatedAt := notificationpermissionMixinFields0[1].Descriptor()
	// notificationpermission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notificationpermission.DefaultUpdatedAt = notificationpermissionDescUpdatedAt.Default.(func() time.Time)
	// notificationpermission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notificationpermission.UpdateDefaultUpdatedAt = notificationpermissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	notificationtimeMixin := schema.NotificationTime{}.Mixin()
	notificationtimeMixinFields0 := notificationtimeMixin[0].Fields()
	_ = notificationtimeMixinFields0
	notificationtimeFields := schema.NotificationTime{}.Fields()
	_ = notificationtimeFields
	// notificationtimeDescCreatedAt is the schema descriptor for created_at field.
	notificationtimeDescCreatedAt := notificationtimeMixinFields0[0].Descriptor()
	// notificationtime.DefaultCreatedAt holds the default value on creation for the created_at field.
	notificationtime.DefaultCreatedAt = notificationtimeDescCreatedAt.Default.(func() time.Time)
	// notificationtimeDescUpdatedAt is the schema descriptor for updated_at field.
	notificationtimeDescUpdatedAt := notificationtimeMixinFields0[1].Descriptor()
	// notificationtime.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notificationtime.DefaultUpdatedAt = notificationtimeDescUpdatedAt.Default.(func() time.Time)
	// notificationtime.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notificationtime.UpdateDefaultUpdatedAt = notificationtimeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notificationtimeDescSendMinute is the schema descriptor for send_minute field.
	notificationtimeDescSendMinute := notificationtimeFields[1].Descriptor()
	// notificationtime.SendMinuteValidator is a validator for the "send_minute" field. It is called by the builders before save.
	notificationtime.SendMinuteValidator = func() func(int) error {
		validators := notificationtimeDescSendMinute.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(send_minute int) error {
			for _, fn := range fns {
				if err := fn(send_minute); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	passwordtokenFields := schema.PasswordToken{}.Fields()
	_ = passwordtokenFields
	// passwordtokenDescHash is the schema descriptor for hash field.
	passwordtokenDescHash := passwordtokenFields[0].Descriptor()
	// passwordtoken.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	passwordtoken.HashValidator = passwordtokenDescHash.Validators[0].(func(string) error)
	// passwordtokenDescCreatedAt is the schema descriptor for created_at field.
	passwordtokenDescCreatedAt := passwordtokenFields[1].Descriptor()
	// passwordtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	passwordtoken.DefaultCreatedAt = passwordtokenDescCreatedAt.Default.(func() time.Time)
	phoneverificationcodeMixin := schema.PhoneVerificationCode{}.Mixin()
	phoneverificationcodeMixinFields0 := phoneverificationcodeMixin[0].Fields()
	_ = phoneverificationcodeMixinFields0
	phoneverificationcodeFields := schema.PhoneVerificationCode{}.Fields()
	_ = phoneverificationcodeFields
	// phoneverificationcodeDescCreatedAt is the schema descriptor for created_at field.
	phoneverificationcodeDescCreatedAt := phoneverificationcodeMixinFields0[0].Descriptor()
	// phoneverificationcode.DefaultCreatedAt holds the default value on creation for the created_at field.
	phoneverificationcode.DefaultCreatedAt = phoneverificationcodeDescCreatedAt.Default.(func() time.Time)
	// phoneverificationcodeDescUpdatedAt is the schema descriptor for updated_at field.
	phoneverificationcodeDescUpdatedAt := phoneverificationcodeMixinFields0[1].Descriptor()
	// phoneverificationcode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	phoneverificationcode.DefaultUpdatedAt = phoneverificationcodeDescUpdatedAt.Default.(func() time.Time)
	// phoneverificationcode.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	phoneverificationcode.UpdateDefaultUpdatedAt = phoneverificationcodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	profileMixin := schema.Profile{}.Mixin()
	profileMixinFields0 := profileMixin[0].Fields()
	_ = profileMixinFields0
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileMixinFields0[0].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescUpdatedAt is the schema descriptor for updated_at field.
	profileDescUpdatedAt := profileMixinFields0[1].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	// profile.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	profile.UpdateDefaultUpdatedAt = profileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// profileDescFullyOnboarded is the schema descriptor for fully_onboarded field.
	profileDescFullyOnboarded := profileFields[3].Descriptor()
	// profile.DefaultFullyOnboarded holds the default value on creation for the fully_onboarded field.
	profile.DefaultFullyOnboarded = profileDescFullyOnboarded.Default.(bool)
	pwapushsubscriptionMixin := schema.PwaPushSubscription{}.Mixin()
	pwapushsubscriptionMixinFields0 := pwapushsubscriptionMixin[0].Fields()
	_ = pwapushsubscriptionMixinFields0
	pwapushsubscriptionFields := schema.PwaPushSubscription{}.Fields()
	_ = pwapushsubscriptionFields
	// pwapushsubscriptionDescCreatedAt is the schema descriptor for created_at field.
	pwapushsubscriptionDescCreatedAt := pwapushsubscriptionMixinFields0[0].Descriptor()
	// pwapushsubscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	pwapushsubscription.DefaultCreatedAt = pwapushsubscriptionDescCreatedAt.Default.(func() time.Time)
	// pwapushsubscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	pwapushsubscriptionDescUpdatedAt := pwapushsubscriptionMixinFields0[1].Descriptor()
	// pwapushsubscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pwapushsubscription.DefaultUpdatedAt = pwapushsubscriptionDescUpdatedAt.Default.(func() time.Time)
	// pwapushsubscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pwapushsubscription.UpdateDefaultUpdatedAt = pwapushsubscriptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// pwapushsubscriptionDescEndpoint is the schema descriptor for endpoint field.
	pwapushsubscriptionDescEndpoint := pwapushsubscriptionFields[0].Descriptor()
	// pwapushsubscription.EndpointValidator is a validator for the "endpoint" field. It is called by the builders before save.
	pwapushsubscription.EndpointValidator = pwapushsubscriptionDescEndpoint.Validators[0].(func(string) error)
	// pwapushsubscriptionDescP256dh is the schema descriptor for p256dh field.
	pwapushsubscriptionDescP256dh := pwapushsubscriptionFields[1].Descriptor()
	// pwapushsubscription.P256dhValidator is a validator for the "p256dh" field. It is called by the builders before save.
	pwapushsubscription.P256dhValidator = pwapushsubscriptionDescP256dh.Validators[0].(func(string) error)
	// pwapushsubscriptionDescAuth is the schema descriptor for auth field.
	pwapushsubscriptionDescAuth := pwapushsubscriptionFields[2].Descriptor()
	// pwapushsubscription.AuthValidator is a validator for the "auth" field. It is called by the builders before save.
	pwapushsubscription.AuthValidator = pwapushsubscriptionDescAuth.Validators[0].(func(string) error)
	sentemailMixin := schema.SentEmail{}.Mixin()
	sentemailMixinFields0 := sentemailMixin[0].Fields()
	_ = sentemailMixinFields0
	sentemailFields := schema.SentEmail{}.Fields()
	_ = sentemailFields
	// sentemailDescCreatedAt is the schema descriptor for created_at field.
	sentemailDescCreatedAt := sentemailMixinFields0[0].Descriptor()
	// sentemail.DefaultCreatedAt holds the default value on creation for the created_at field.
	sentemail.DefaultCreatedAt = sentemailDescCreatedAt.Default.(func() time.Time)
	// sentemailDescUpdatedAt is the schema descriptor for updated_at field.
	sentemailDescUpdatedAt := sentemailMixinFields0[1].Descriptor()
	// sentemail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sentemail.DefaultUpdatedAt = sentemailDescUpdatedAt.Default.(func() time.Time)
	// sentemail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sentemail.UpdateDefaultUpdatedAt = sentemailDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[3].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
}

const (
	Version = "v0.14.0"                                         // Version of ent codegen.
	Sum     = "h1:EO3Z9aZ5bXJatJeGqu/EVdnNr6K4mRq3rWe5owt0MC4=" // Sum of ent codegen.
)
