package constant

const (
	PREFIX_PROJECT = "SDH"

	KEY_PASSWORD               = "V1si0n3t*1!"
	LENGTH_TOKEN_DEFAULT       = 16
	LONG_TOKEN_EXPIRED_DEFAULT = 14400
	//JWT_SIGNATURE_KEY          = "metaforceverysecret"

	PAGING_SYNC_DEFAULT = 1
	LIMIT_SYNC_DEFAULT  = 20

	TABLE_NOTIFICATION_CATEGORY_NAME = "notification_category"

	STORAGE_PATH_PDF         = "./storage/pdf"
	STORAGE_PATH_APPLICATION = "./storage/application"

	MONGO_COLLECTION_TR_TASKLISTS_AC              = "tr_tasklists_ac"
	MONGO_COLLECTION_TR_TASKLISTS_2020            = "tr_tasklists_2020"
	MONGO_COLLECTION_TR_TASKLISTS                 = "tr_tasklists"
	MONGO_COLLECTION_TR_TASKLISTS_DUMMY           = "tr_tasklists_dummy"
	MONGO_COLLECTION_MS_QUESTION_SURVEYS_AC       = "ms_questionsurveys"
	MONGO_COLLECTION_MS_QUESTION_SURVEYS_MAYAPADA = "ms_questionsurveys_mayapada"
	MONGO_COLLECTION_MS_QUESTION_SURVEYS          = "ms_questionsurveys"

	TABLE_MS_AREA_NAME                       = "MS_Area"
	TABLE_MS_CITY_NAME                       = "MS_City"
	TABLE_MS_DESTINATION_NAME                = "MS_Destination"
	TABLE_MS_DESTINATION_NAME_PRELOAD        = "MSDestination"
	TABLE_MS_TASKLIST_DETAIL_NAME            = "MS_TasklistDetail"
	TABLE_MS_USER_DETAIL_SURVEY_NAME_PRELOAD = "MSUserDetailSurvey"
	TABLE_CONFIG_NAME                        = "Config"
	TABLE_MS_TENANT_NAME                     = "MS_Tenant"
	TABLE_MS_USER_DETAIL_SURVEY_NAME         = "MS_UserDetailSurvey"
	TABLE_MS_VENDOR_NAME                     = "MS_Vendor"
	TABLE_MS_VENDOR_NAME_PRELOAD             = "Vendor"
	TABLE_MS_USER_NAME                       = "MS_User"
	TABLE_MS_USER_NAME_PRELOAD               = "User"
	TABLE_TR_VISIT_NAME                      = "TR_Visit"
	TABLE_TR_VISIT_AC_NAME                   = "TR_Visit_AC"
	TABLE_TR_VISIT_AC_TEMP_NAME              = "TR_VisitTemp_AC"
	TABLE_TR_VISIT_TEMP_NAME                 = "TR_VisitTemp"
	TABLE_MS_TYPELOC_NAME                    = "MS_Typeloc"
	TABLE_MS_AREA_NAME_PRELOAD               = "MSArea"
	TABLE_VERSION_NAME                       = "version"
	TABLE_SOP_NAME                           = "sop"
	TABLE_INSTRUCTION_NAME                   = "instruction"

	OPERATION_SQL_INSERT = "insert"
	OPERATION_SQL_UPDATE = "update"
	OPERATION_SQL_DELETE = "delete"

	IMAGE_PNG_EXTENSION = "png"

	CONTENT_TYPE_IMAGE_PNG        = "image/png"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	FORMAT_TIME_SYNC                 = "2006-01-02 15:04:05 -0700"
	FORMAT_TIME_SYNCWITHOUT_TIMEZONE = "2006-01-02 15:04:05"
	FORMAT_DATE                      = "20060102"
	FORMAT_DATE_MIDTRANS             = "2006-01-02 15:04:05 -0700"
	FORMAT_TIME_MINUTE               = "2006-01-02 15:04"

	URL_SDH_API = "https://sandiegohills.co.id/SDHAPI/MobileApps_Services.svc"

	DEFAULT_QUERY_SOFT_DELETE    = "deleted_at IS NULL"
	DEFAULT_QUERY_SOFT_DELETE_MF = `"DeletedTime" IS NULL`

	ITOP_OPERATION_CORE_CREATE            = "core/create"
	ITOP_OPERATION_CORE_UPDATE            = "core/update"
	ITOP_OPERATION_CORE_GET               = "core/get"
	ITOP_OPERATION_CORE_DELETE            = "core/delete"
	ITOP_OPERATION_CORE_APPLY_STIMULUS    = "core/apply_stimulus"
	ITOP_OPERATION_CORE_CHECK_CREDENTIALS = "core/check_credentials"

	ITOP_CLASS_SERVICE             = "Service"
	ITOP_CLASS_SERVICE_SUBCATEGORY = "ServiceSubcategory"

	ITOP_STATUS_NEW           = "new"
	ITOP_STATUS_DISPATCH      = "dispatched"
	ITOP_STATUS_REDISPATCH    = "redispatched"
	ITOP_STATUS_ASSIGN        = "assigned"
	ITOP_STATUS_PENDING       = "pending"
	ITOP_STATUS_REASSIGN      = "reassign"
	ITOP_STATUS_ESCALATED_TTO = "escalated_tto"
	ITOP_STATUS_ESCALATED_TTR = "escalated_ttr"

	ITOP_TABLE_VIEW_PERSON = "view_Person"
	ITOP_COMMENT           = "Sync From MetaCRM Apps"
)
