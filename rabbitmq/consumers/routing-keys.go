package consumers

import (
	"fmt"
	"base_service/config"
)

// Restaurant
const RestaurantRoutingKey = "restaurant"
const BranchAttributeRoutingKey = "branch-attribute"
const BranchRoutingKey = "branch"
const BranchCloseReasonRoutingKey = "branch-close-reason"
const DineInGalleryRoutingKey = "dine-in-gallery"

// Menu Management
const MenuCategoryRoutingKey = "category"
const VariationRoutingKey = "variation"
const AddonCategoryRoutingKey = "addon-category"
const AddonRoutingKey = "addon"
const MenuRoutingKey = "menu"
const CuisineRoutingKey = "cuisine"
const MenuItemTimeSlotRoutingKey = "menu-item-time-slot"
const MealTypeTimeSetupRoutingKey = "meal-type-time-setup"
const FeaturedFoodRoutingKey = "featured-food"

// Campaigns
const CampaignRoutingKey = "campaign"

// Riders
const RiderRoutingKey = "rider"
const ZoneAssignmentRoutingKey = "rider-and-zone"
const RiderSuspensionRoutingKey = "rider-suspension"
const RiderLeaveRequestRoutingKey = "rider-leave-request"

// Rider Settings
const VehicleTypeRoutingKey = "vehicle-type"
const RiderTypeRoutingKey = "rider-type"
const RiderContractTypeRoutingKey = "rider-contract-type"
const ZoneWiseRiderDeliveryChargeRoutingKey = "zonewise-rider-delivery-charge"
const KnowledgeRoutingKey = "knowledge"
const ReferralGoalSettingRoutingKey = "referral-goal-setting"
const OrderStatusRoutingKey = "order-status"
const PerformanceTipRoutingKey = "performance-tip"

// Shift Setting
const WeekRoutingKey = "week-setup"
const ShiftDutySetupRoutingKey = "rider-shift-duty-setup"
const ShiftExtensionSetupRoutingKey = "extension-time-setting"
const BreakTimeRoutingKey = "break-time-setting"
const SwapStatusRoutingKey = "shift-swap-status"
const BookingTimeStatusRoutingKey = "booking-time-status"

// Rider Quest
const QuestRoutingKey = "quest"

// Rider Batching
const BatchLevelRoutingKey = "batch-level-setting"
const BatchWiseShiftConfigurationRoutingKey = "batch-wise-shift-booking-config"

// Delivery Settings
const BagTypeRoutingKey = "bag-type"
const KilometerWiseDeliveryChargeRoutingKey = "rider-kilometerwise-delivery-charge"
const NightShiftChargeRoutingKey = "night-shift-delivery-charge"

// Rider Wallet Settings
const ParticipantRoutingKey = "participant"
const PaymentTypeRoutingKey = "payment-type"

// Shift Management
const RiderShiftBookingRoutingKey = "rider-shift-booking"
const ShiftExtensionRequestRoutingKey = "shift-extension-request"
const BreakRequestRoutingKey = "break-request"

// Promotion
const PromoCodeRoutingKey = "promo-code"
const PromotionRoutingKey = "promotion"
const PopUpBannerRoutingKey = "pop-up-banner"
const AdvertisementRoutingKey = "advertisement"
const InfoBannerRoutingKey = "info-banner"
const PromoBannerRoutingKey = "promo-banner"

// Reward Point
const RewardPointSettingRoutingKey = "reward-point-settings"
const RewardLevelSettingRoutingKey = "reward-level-settings"

// Subscriptions
const SubscriptionTypeRoutingKey = "subscription-type"
const SubscriptionPlanRoutingKey = "subscription-plan"
const SubscriptionRoutingKey = "subscription"

// Voucher
const VoucherSettingRoutingKey = "voucher-settings"
const CouponRoutingKey = "coupon"

// Order settings
const OrderAmountThresholdRoutingKey = "amount-threshold"
const OrderThresholdRoutingKey = "order-threshold"

// System settings
const ZoneRoutingKey = "zone"
const CityRoutingKey = "city"
const UserRadiusRoutingKey = "user-radius"
const PlatformOperationTimeSlotRoutingKey = "platform-operation-time-slot"
const SystemOnOffReasonRoutingKey = "system-on-off-reason"
const FaqRoutingKey = "faq"
const ReviewReasonRoutingKey = "review-reason"
const SpecialHourRoutingKey = "special-hour"
const RestaurantTutorialRoutingKey = "restaurant-tutorial"
const RiderTutorialRoutingKey = "rider-tutorial"

// Refund
const RefundRoutingKey = "refund"

// Auth
const UserRoutingKey = "user"

// Order
const OrderRoutingKey = "order"

func QueueName(routingKey string) string {
	conf := config.GetConfig()
	return fmt.Sprintf("%s%s:queue", conf.RmqQueuePrefix, routingKey)
}

// ---------------------------- User ---------------------------
func UserQueueName() string {
	return QueueName(UserRoutingKey)
}

// ---------------------------- Restaurant ----------------------
func RestaurantQueueName() string {
	return QueueName(RestaurantRoutingKey)
}
func BranchQueueName() string {
	return QueueName(BranchRoutingKey)
}
func BranchAttributeQueueName() string {
	return QueueName(BranchAttributeRoutingKey)
}
func BranchCloseReasonQueueName() string {
	return QueueName(BranchCloseReasonRoutingKey)
}
func DineInGalleryQueueName() string {
	return QueueName(DineInGalleryRoutingKey)
}
func SubscriptionQueueName() string { return QueueName(SubscriptionRoutingKey) }

// ------------------- Menu Management ----------------------
func MenuCategoryQueueName() string {
	return QueueName(MenuCategoryRoutingKey)
}
func VariationQueueName() string {
	return QueueName(VariationRoutingKey)
}
func AddonCategoryQueueName() string {
	return QueueName(AddonCategoryRoutingKey)
}
func AddonQueueName() string {
	return QueueName(AddonRoutingKey)
}
func MenuQueueName() string {
	return QueueName(MenuRoutingKey)
}
func CuisineQueueName() string {
	return QueueName(CuisineRoutingKey)
}
func MenuItemTimeSlotQueueName() string {
	return QueueName(MenuItemTimeSlotRoutingKey)
}
func MealTypeTimeSetupQueueName() string {
	return QueueName(MealTypeTimeSetupRoutingKey)
}
func FeaturedFoodQueueName() string {
	return QueueName(FeaturedFoodRoutingKey)
}

// -------------------- Rider ----------------------------
func RiderQueueName() string {
	return QueueName(RiderRoutingKey)
}
func ZoneAssignmentQueueName() string {
	return QueueName(ZoneAssignmentRoutingKey)
}
func RiderSuspensionQueueName() string {
	return QueueName(RiderSuspensionRoutingKey)
}
func RiderLeaveRequestQueueName() string {
	return QueueName(RiderLeaveRequestRoutingKey)
}
func VehicleTypeQueueName() string {
	return QueueName(VehicleTypeRoutingKey)
}
func RiderTypeQueueName() string {
	return QueueName(RiderTypeRoutingKey)
}
func RiderContractTypeQueueName() string {
	return QueueName(RiderContractTypeRoutingKey)
}
func ZoneWiseRiderDeliveryChargeQueueName() string {
	return QueueName(ZoneWiseRiderDeliveryChargeRoutingKey)
}
func KnowledgeQueueName() string {
	return QueueName(KnowledgeRoutingKey)
}

func ReferralGoalSettingQueueName() string {
	return QueueName(ReferralGoalSettingRoutingKey)
}
func OrderStatusQueueName() string {
	return QueueName(OrderStatusRoutingKey)
}
func PerformanceTipQueueName() string {
	return QueueName(PerformanceTipRoutingKey)
}
func RiderShiftBookingQueueName() string {
	return QueueName(RiderShiftBookingRoutingKey)
}

func BreakTimeQueueName() string         { return QueueName(BreakTimeRoutingKey) }
func SwapStatusQueueName() string        { return QueueName(SwapStatusRoutingKey) }
func BookingTimeStatusQueueName() string { return QueueName(BookingTimeStatusRoutingKey) }
func QuestQueueName() string             { return QueueName(QuestRoutingKey) }
func BatchLevelQueueName() string        { return QueueName(BatchLevelRoutingKey) }
func BatchWiseShiftConfigurationQueueName() string {
	return QueueName(BatchWiseShiftConfigurationRoutingKey)
}
func BagTypeQueueName() string { return QueueName(BagTypeRoutingKey) }
func KilometerWiseDeliveryChargeQueueName() string {
	return QueueName(KilometerWiseDeliveryChargeRoutingKey)
}
func NightShiftChargeQueueName() string      { return QueueName(NightShiftChargeRoutingKey) }
func ParticipantQueueName() string           { return QueueName(ParticipantRoutingKey) }
func PaymentTypeQueueName() string           { return QueueName(PaymentTypeRoutingKey) }
func WeekQueueName() string                  { return QueueName(WeekRoutingKey) }
func ShiftDutySetupQueueName() string        { return QueueName(ShiftDutySetupRoutingKey) }
func ShiftExtensionSetupQueueName() string   { return QueueName(ShiftExtensionSetupRoutingKey) }
func ShiftExtensionRequestQueueName() string { return QueueName(ShiftExtensionRequestRoutingKey) }
func BreakTimeRequestQueueName() string      { return QueueName(BreakRequestRoutingKey) }
func UserRadiusQueueName() string            { return QueueName(UserRadiusRoutingKey) }

// --------------------- Promotion ----------------------------
func PromoCodeQueueName() string {
	return QueueName(PromoCodeRoutingKey)
}
func PromotionQueueName() string {
	return QueueName(PromotionRoutingKey)
}
func PopUpBannerQueueName() string {
	return QueueName(PopUpBannerRoutingKey)
}
func AdvertisementQueueName() string {
	return QueueName(AdvertisementRoutingKey)
}
func InfoBannerQueueName() string {
	return QueueName(InfoBannerRoutingKey)
}
func PromoBannerQueueName() string {
	return QueueName(PromoBannerRoutingKey)
}

// ------------------------ Order ---------------------------------
func OrderQueueName() string { return QueueName(OrderRoutingKey) }

// ----------------------- Reward Point ----------------------------
func RewardPointSettingQueueName() string {
	return QueueName(RewardPointSettingRoutingKey)
}
func RewardLevelSettingQueueName() string {
	return QueueName(RewardLevelSettingRoutingKey)
}

// ----------------------- Subscription -----------------------------
func SubscriptionTypeQueueName() string {
	return QueueName(SubscriptionTypeRoutingKey)
}

func SubscriptionPlanQueueName() string {
	return QueueName(SubscriptionPlanRoutingKey)
}

// ------------------------ Voucher ----------------------------------
func VoucherSettingQueueName() string {
	return QueueName(VoucherSettingRoutingKey)
}
func CouponQueueName() string {
	return QueueName(CouponRoutingKey)
}

// ----------------------- Order settings ----------------------------
func OrderAmountThresholdQueueName() string {
	return QueueName(OrderAmountThresholdRoutingKey)
}
func OrderThresholdQueueName() string {
	return QueueName(OrderThresholdRoutingKey)
}

// ----------------------- System settings ---------------------------
func ZoneQueueName() string {
	return QueueName(ZoneRoutingKey)
}
func CityQueueName() string {
	return QueueName(CityRoutingKey)
}
func SystemOnOffReasonQueueName() string {
	return QueueName(SystemOnOffReasonRoutingKey)
}
func FaqQueueName() string {
	return QueueName(FaqRoutingKey)
}
func ReviewReasonQueueName() string {
	return QueueName(ReviewReasonRoutingKey)
}
func SpecialHourQueueName() string {
	return QueueName(SpecialHourRoutingKey)
}
func RestaurantTutorialQueueName() string {
	return QueueName(RestaurantTutorialRoutingKey)
}
func RiderTutorialQueueName() string {
	return QueueName(RiderTutorialRoutingKey)
}
func PlatformOperationTimeSlotQueueName() string {
	return QueueName(PlatformOperationTimeSlotRoutingKey)
}

// ----------------------------------- Refund ------------------------
func RefundQueueName() string {
	return QueueName(RefundRoutingKey)
}

// ----------------------------------- Campaign ----------------------
func CampaignQueueName() string {
	return QueueName(CampaignRoutingKey)
}
