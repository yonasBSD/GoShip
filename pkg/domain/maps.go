package domain

// Initialize the map of NotificationPermissionType to NotificationPermission
var NotificationPermissionMap = map[NotificationPermissionType]NotificationPermission{
	NotificationPermissionDailyReminder: {
		Title:      "Daily conversation",
		Subtitle:   "A reminder to not miss today's question, sent at most once a day.",
		Permission: NotificationPermissionDailyReminder.Value,
	},
	NotificationPermissionPartnerActivity: {
		Title:      "Partner activity",
		Subtitle:   "Answers you missed, sent at most once a day.",
		Permission: NotificationPermissionPartnerActivity.Value,
	},
}

var NotificationCenterButtonText = map[NotificationType]string{
	NotificationTypeMutualQuestionAnswered:        "See answer",
	NotificationTypeConnectionEngagedWithQuestion: "Answer",
	NotificationTypeConnectionReactedToAnswer:     "See reaction",
}

// DeleteOnceReadNotificationTypesMap is a map of notification types th;oiSJDfiujladijrgoizdikrjgat can be deleted once seen.
// Note that the boolean doesn't matter, this is just a lazy way of creating a set in Go.
var DeleteOnceReadNotificationTypesMap = map[NotificationType]bool{
	NotificationTypeDailyConversationReminder: true,
	NotificationTypeNewHomeFeedQA:             true,
}