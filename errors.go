package mdm

type MCProfileErrorDomain int

const (
	MalformedProfile MCProfileErrorDomain = iota + 1000
	UnsupportedProfileVersion
	MissingRequiredField
	BadDataTypeInField
	BadSignature
	EmptyProfile
	CannotDecrypt
	NonUniqueUUIDs
	NonUniquePayloadIdentifiers
	ProfileInstallationFailure
	UnsupportedFieldValue
)

type MCPayloadErrorDomain int

const (
	MalformedPayload MCPayloadErrorDomain = iota + 2000
	UnsupportedPayloadVersion
	MissingRequiredField
	BadDataTypeInField
	UnsupportedFieldValue
	InternalError
)

type MCRestrictionsErrorDomain int

const (
	InconsistentRestrictionSense MCRestrictionsErrorDomain = iota + 3000
	InconsistentValueComparisonSense
)

type MCInstallationErrorDomain int

const (
	CannotParseProfile MCInstallationErrorDomain = iota + 4000
	InstallationFailure
	DuplicateUUID
	ProfileNotQueued
	UserCancelled
	PasscodeNotCompliant
	ProfileRemovalDateInPast
	UnrecognisedFileFormat
	MismatchedCertificates
	DeviceLocked
	UpdatedProfileWrongIdentifier
	FinalProfileNotConfiguration
	ProfileNotUpdatable
	UpdateFailed
	NoDeviceIdentity
	ReplacementNoMDMPayload
	InternalError
	MultipleHTTPProxyPayloads
	MultipleCellularPayloads
	MultipleApplockPayloads
	UIInstallProhibited
	ProfileMustBeNonInteractive
	ProfileMustBeInstalledByMDM
	UnnacceptablePayload
	ProfileNotFound
	InvalidSupervision
	RemovalDateInPast
	ProfileRequiresPasscodeChange
	MultipleHomeScreenPayloads
	MultipleNotificationPayloads
	UnacceptablePayloadMultiuser
	PayloadContainsSensitiveInfo
)

type MCPasscodeErrorDomain int

const (
	PasscodeTooShort MCPasscodeErrorDomain = iota + 5000
	TooFewUniqueChars
	TooFewComplexChars
	RepeatingChars
	AscendingDescendingChars
	RequiresNumber
	RequiresAlpha
	PasscodeExpired
	PasscodeTooRecent
	_
	DeviceLocked
	WrongPasscode
	_
	CannotClearPasscode
	CannotSetPasscode
	CannotSetGracePeriod
	CannotSetFingerprintUnlock
	CannotSetFingerprintPurchase
	CannotSetMaxAttempts
)

type MCKeychainErrorDomain int

const (
	KeychainSystemError MCKeychainErrorDomain = iota + 6000
	EmptyString
	CannotCreateQuery
)

type MCEmailErrorDomain int

const (
	HostUnreachable MCEmailErrorDomain = iota + 7000
	InvalidCredentials
	UnknownValidationError
	SMIMECertificateNotFound
	SMIMECertificateBad
	IMAPMisconfigured
	POPMisconfigured
	SMTPMisconfigured
)

type MCWebClipErrorDomain int

const (
	CannotInstallWebClip MCWebClipErrorDomain = iota + 8000
)

type MCCertificateErrorDomain int

const (
	InvalidPassword MCCertificateErrorDomain = iota + 9000
	TooManyCertificatesInPayload
	CannotStoreCertificate
	CannotStoreWAPIData
	CannotStoreRootCertificate
	CertificateMalformed
	CertificateNotIdentity
)

type MCDefaultsErrorDomain int

const (
	CannotInstallDefaults MCDefaultsErrorDomain = iota + 10000
	InvalidSigner
)

type MCAPNErrorDomain int

const (
	CannotInstallAPN MCAPNErrorDomain = iota + 11000
	CustomAPNAlreadyInstalled
)

type MCMDMErrorDomain int

const (
	InvalidAccessRights MCMDMErrorDomain = iota + 12000
	MultipleMDMInstances
	CannotCheckIn
	InvalidChallengeResponse
	InvalidPushCertificate
	CannotFindCertificate
	RedirectRefused
	NotAuthorized
	MalformedRequest
	InvalidReplacementProfile
	InternalConsistencyError
	InvalidMDMConfiguration
	MDMReplacementMismatch
	ProfileNotManaged
	ProvisioningProfileNotManaged
	CannotGetPushToken
	MissingIdentity
	CannotCreateEscrowKeybag
	CannotCopyEscrowKeybagData
	CannotCopyEscrowSecret
	UnauthorizedByServer
	InvalidRequestType
	InvalidTopic
)
