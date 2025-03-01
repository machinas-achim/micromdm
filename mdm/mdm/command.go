package mdm

import (
	"strings"

	"github.com/google/uuid"
	"github.com/micromdm/micromdm/mdm/appmanifest"
)

type CommandRequest struct {
	UDID        string `json:"udid"`
	CommandUUID string `json:"command_uuid"`
	*Command
}

type CommandPayload struct {
	CommandUUID string   `json:"command_uuid"`
	Command     *Command `json:"command"`
}

func NewCommandPayload(request *CommandRequest) (*CommandPayload, error) {
	payload := &CommandPayload{
		CommandUUID: request.CommandUUID,
		Command:     request.Command,
	}
	if strings.TrimSpace(payload.CommandUUID) == "" {
		payload.CommandUUID = uuid.New().String()
	}
	return payload, nil
}

type Command struct {
	RequestType                     string `json:"request_type"`
	InstallProfile                  *InstallProfile
	RemoveProfile                   *RemoveProfile
	InstallProvisioningProfile      *InstallProvisioningProfile
	RemoveProvisioningProfile       *RemoveProvisioningProfile
	InstalledApplicationList        *InstalledApplicationList
	DeviceInformation               *DeviceInformation
	DeviceLock                      *DeviceLock
	ClearPasscode                   *ClearPasscode
	EraseDevice                     *EraseDevice
	RequestMirroring                *RequestMirroring
	Restrictions                    *Restrictions
	UnlockUserAccount               *UnlockUserAccount
	DeleteUser                      *DeleteUser
	EnableLostMode                  *EnableLostMode
	InstallApplication              *InstallApplication
	InstallEnterpriseApplication    *InstallEnterpriseApplication
	AccountConfiguration            *AccountConfiguration
	ApplyRedemptionCode             *ApplyRedemptionCode
	ManagedApplicationList          *ManagedApplicationList
	RemoveApplication               *RemoveApplication
	InviteToProgram                 *InviteToProgram
	ValidateApplications            *ValidateApplications
	InstallMedia                    *InstallMedia
	RemoveMedia                     *RemoveMedia
	Settings                        *Settings
	ManagedApplicationConfiguration *ManagedApplicationConfiguration
	ManagedApplicationAttributes    *ManagedApplicationAttributes
	ManagedApplicationFeedback      *ManagedApplicationFeedback
	SetFirmwarePassword             *SetFirmwarePassword
	VerifyFirmwarePassword          *VerifyFirmwarePassword
	SetRecoveryLock                 *SetRecoveryLock
	VerifyRecoveryLock              *VerifyRecoveryLock
	SetAutoAdminPassword            *SetAutoAdminPassword
	ScheduleOSUpdate                *ScheduleOSUpdate
	ScheduleOSUpdateScan            *ScheduleOSUpdateScan
	ActiveNSExtensions              *ActiveNSExtensions
	RotateFileVaultKey              *RotateFileVaultKey
	RefreshCellularPlans            *RefreshCellularPlans
	LOMDeviceRequest                *LOMDeviceRequest
	LOMSetupRequestCommand          *LOMSetupRequestCommand
}

// InstallProfile is an InstallProfile MDM Command
type InstallProfile struct {
	Payload []byte `json:"payload,omitempty"`
}

type RemoveProfile struct {
	Identifier string `json:"identifier,omitempty"`
}

type InstallProvisioningProfile struct {
	ProvisioningProfile []byte `plist:",omitempty" json:"provisioning_profile,omitempty"`
}

type RemoveProvisioningProfile struct {
	UUID string `json:"uuid"`
}

type InstalledApplicationList struct {
	Identifiers     []string `plist:",omitempty" json:"identifiers,omitempty"`
	ManagedAppsOnly bool     `plist:",omitempty" json:"managed_appd_only,omitempty"`
}

type DeviceInformation struct {
	Queries []string `plist:",omitempty" json:"queries,omitempty"`
}

type DeviceLock struct {
	PIN         string `plist:",omitempty" json:"pin"`
	Message     string `plist:",omitempty" json:"message,omitempty"`
	PhoneNumber string `plist:",omitempty" json:"phone_number,omitempty"`
}

type ClearPasscode struct {
	UnlockToken []byte `json:"unlock_token"`
}

type EraseDevice struct {
	PIN                    string `json:"pin"`
	PreserveDataPlan       bool   `plist:",omitempty" json:"preserve_data_plan,omitempty"`
	DisallowProximitySetup bool   `plist:",omitempty" json:"disallow_proximity_setup,omitempty"`
}

type RequestMirroring struct {
	DestinationName     string `plist:",omitempty" json:"destination_name,omitempty"`
	DestinationDeviceID string `plist:",omitempty" json:"destination_device_id,omitempty"`
	ScanTime            string `plist:",omitempty" json:"scan_time,omitempty"`
	Password            string `plist:",omitempty" json:"password,omitempty"`
}

type Restrictions struct {
	ProfileRestrictions bool `json:"profile_restrictions"`
}

type UnlockUserAccount struct {
	UserName string `json:"username"`
}

type DeleteUser struct {
	UserName      string `plist:",omitempty" json:"username,omitempty"`
	ForceDeletion bool   `plist:",omitempty" json:"force_deletion,omitempty"`
}

type EnableLostMode struct {
	Message     string `plist:",omitempty" json:"message,omitempty"`
	PhoneNumber string `plist:",omitempty" json:"phone_number,omitempty"`
	Footnote    string `plist:",omitempty" json:"footnote,omitempty"`
}

type InstallEnterpriseApplication struct {
        ChangeManagementState          *string                                    `plist:",omitempty" json:"change_management_state,omitempty"`
        Configuration                  *InstallEnterpriseApplicationConfiguration `plist:",omitempty" json:"configuration,omitempty"`
        InstallAsManaged               *bool                                      `plist:",omitempty" json:"install_as_managed,omitempty"`
        IosApp                         *bool                                      `plist:",omitempty" json:"ios_app,omitempty"`
        ManagementFlags                *int                                       `plist:",omitempty" json:"management_flags,omitempty"`
	Manifest                       *appmanifest.Manifest                      `plist:",omitempty" json:"manifest,omitempty"`
	ManifestURL                    *string                                    `plist:",omitempty" json:"manifest_url,omitempty"`
	ManifestURLPinningCerts        [][]byte                                   `plist:",omitempty" json:"manifest_url_pinning_certs,omitempty"`
	PinningRevocationCheckRequired *bool                                      `plist:",omitempty" json:"pinning_revocation_check_required,omitempty"`
}
type InstallEnterpriseApplicationConfiguration struct{}

type InstallApplication struct {
        Attributes            *InstallApplicationAttributes    `plist:",omitempty" json:"attributes,omitempty"`
        ChangeManagementState *string                          `plist:",omitempty" json:"change_management_state,omitempty"`
        Configuration         *InstallApplicationConfiguration `plist:",omitempty" json:"configuration,omitempty"`
        Identifier            *string                          `plist:",omitempty" json:"identifier,omitempty"`
        InstallAsManaged      *bool                            `plist:",omitempty" json:"install_as_managed,omitempty"`
        IosApp                *bool                            `plist:",omitempty" json:"ios_app,omitempty"`
	ITunesStoreID         *int64                           `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	ManagementFlags       *int                             `plist:",omitempty" json:"management_flags,omitempty"`
	ManifestURL           *string                          `plist:",omitempty" json:"manifest_url,omitempty"`
	Options               *InstallApplicationOptions       `plist:"Options,omitempty" json:"options,omitempty"`
}

type InstallApplicationOptions struct {
	PurchaseMethod *int64 `plist:"PurchaseMethod,omitempty" json:"purchase_method,omitempty"`
}

type InstallApplicationConfiguration struct{}
type InstallApplicationAttributes struct{}

type AccountConfiguration struct {
	SkipPrimarySetupAccountCreation     bool           `plist:",omitempty" json:"skip_primary_setup_account_creation,omitempty"`
	SetPrimarySetupAccountAsRegularUser bool           `plist:",omitempty" json:"set_primary_setup_account_as_regular_user,omitempty"`
	DontAutoPopulatePrimaryAccountInfo  bool           `plist:",omitempty" json:"dont_auto_populate_primary_account_info,omitempty"`
	LockPrimaryAccountInfo              bool           `plist:",omitempty" json:"lock_primary_account_info,omitempty"`
	PrimaryAccountFullName              string         `plist:",omitempty" json:"primary_account_full_name,omitempty"`
	PrimaryAccountUserName              string         `plist:",omitempty" json:"primary_account_user_name,omitempty"`
	AutoSetupAdminAccounts              []AdminAccount `plist:",omitempty" json:"auto_setup_admin_accounts,omitempty"`
}

type AdminAccount struct {
	ShortName    string `plist:"shortName" json:"short_name"`
	FullName     string `plist:"fullName,omitempty" json:"full_name,omitempty"`
	PasswordHash []byte `plist:"passwordHash" json:"password_hash"`
	Hidden       bool   `plist:"hidden,omitempty" json:"hidden,omitempty"`
}

type ApplyRedemptionCode struct {
	Identifier     string `plist:",omitempty" json:"identifier,omitempty"`
	RedemptionCode string `plist:",omitempty" json:"redemption_code,omitempty"`
}

type ManagedApplicationList struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type RemoveApplication struct {
	Identifier string `plist:",omitempty" json:"identifier,omitempty"`
}

type InviteToProgram struct {
	ProgramID     string `plist:",omitempty" json:"program_id,omitempty"`
	InvitationURL string `plist:",omitempty" json:"invitation_url,omitempty"`
}

type ValidateApplications struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type InstallMedia struct {
	ITunesStoreID *int64 `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	MediaURL      string `plist:",omitempty" json:"media_url,omitempty"`
	MediaType     string `plist:",omitempty" json:"media_type,omitempty"`
	PersistentID  string `plist:",omitempty" json:"persistent_id,omitempty"`
	Author  	  string `plist:",omitempty" json:"author,omitempty"`
	Kind  		  string `plist:",omitempty" json:"kind,omitempty"`
	Title  		  string `plist:",omitempty" json:"title,omitempty"`
	Version  	  string `plist:",omitempty" json:"version,omitempty"`
}

type RemoveMedia struct {
	ITunesStoreID *int64 `plist:"iTunesStoreID,omitempty" json:"itunes_store_id,omitempty"`
	MediaType     string `plist:",omitempty" json:"media_type,omitempty"`
	PersistentID  string `plist:",omitempty" json:"persistent_id,omitempty"`
}

type Settings struct {
	Settings []Setting `plist:",omitempty" json:"settings,omitempty"`
}

type Setting struct {
	Item                    string                 `json:"item"`
	Enabled                 *bool                  `plist:",omitempty" json:"enabled,omitempty"`
	DeviceName              *string                `plist:",omitempty" json:"device_name,omitempty"`
	HostName                *string                `plist:",omitempty" json:"hostname,omitempty"`
	Identifier              *string                `plist:",omitempty" json:"identifier"`
	TimeZone                *string                `plist:",omitempty" json:"time_zone,omitempty"`
	Attributes              map[string]string      `plist:",omitempty" json:"attributes,omitempty"`
	Image                   []byte                 `plist:",omitempty" json:"image,omitempty"`
	Where                   *int                   `plist:",omitempty" json:"where,omitempty"`
	MDMOptions              map[string]interface{} `plist:",omitempty" json:"mdm_options,omitempty"`
	PasscodeLockGracePeriod *int                   `plist:",omitempty" json:"passcode_lock_grace_period,omitempty"`
	MaximumResidentUsers    *int                   `plist:",omitempty" json:"maximum_resident_users,omitempty"`
	Configuration           map[string]interface{} `plist:",omitempty" json:"-"`
	ConfigurationData       []byte                 `plist:"-" json:"configuration"` // used to build the dictionary
}

type LOMDeviceRequest struct {
	RequestList []LOMDeviceRequestCommand `plist:",omitempty" json:"request_list,omitempty"`
}

type LOMDeviceRequestCommand struct {
	DeviceDNSName            string   `json:"device_dns_name"`
	DeviceRequestType        string   `json:"device_request_type"`
	DeviceRequestUUID        string   `json:"device_request_uuid"`
	LOMProtocolVersion       int      `json:"lom_protocol_version"`
	PrimaryIPv6AddressList   []string `json:"primary_ip_v6_address_list"`
	SecondaryIPv6AddressList []string `json:"secondary_ip_v6_address_list"`
}

type LOMSetupRequestCommand struct {
	RequestRequiresNetworkTether *bool `plist:",omitempty" json:"request_requires_network_tether,omitempty"`
}

type ManagedApplicationConfiguration struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type ManagedApplicationAttributes struct {
	Identifiers []string `plist:",omitempty" json:"identifiers,omitempty"`
}

type ManagedApplicationFeedback struct {
	Identifiers    []string `plist:",omitempty" json:"identifiers,omitempty"`
	DeleteFeedback bool     `plist:",omitempty" json:"delete_feedback,omitempty"`
}

type SetFirmwarePassword struct {
	CurrentPassword              string `plist:",omitempty" json:"current_password,omitempty"`
	NewPassword                  string `json:"new_password"`
	RequestRequiresNetworkTether bool   `plist:",omitempty" json:"request_requires_network_tether,omitempty"`
	AllowOroms                   bool   `plist:",omitempty" json:"allow_oroms,omitempty"`
}

type VerifyFirmwarePassword struct {
	Password string `json:"password"`
}

type SetRecoveryLock struct {
	CurrentPassword              string `plist:",omitempty" json:"current_password,omitempty"`
	NewPassword                  string `json:"new_password"`
	RequestRequiresNetworkTether bool   `plist:",omitempty" json:"request_requires_network_tether,omitempty"`
	AllowOroms                   bool   `plist:",omitempty" json:"allow_oroms,omitempty"`
}

type VerifyRecoveryLock struct {
	Password string `json:"password"`
}

type SetAutoAdminPassword struct {
	GUID         string `plist:",omitempty" json:"guid,omitempty"`
	PasswordHash []byte `plist:"passwordHash" json:"password_hash"`
}

type OSUpdate struct {
	ProductKey       string `plist:",omitempty" json:"product_key,omitempty"`
	InstallAction    string `json:"install_action"`
	MaxUserDeferrals *int64 `plist:",omitempty" json:"max_user_deferrals,omitempty"`
	ProductVersion   string `plist:",omitempty" json:"product_version,omitempty"`
	Priority         string `plist:",omitempty" json:"priority,omitempty"`
}

type ScheduleOSUpdate struct {
	Updates []OSUpdate `plist:",omitempty" json:"updates,omitempty"`
}

type ScheduleOSUpdateScan struct {
	Force bool `plist:",omitempty" json:"force,omitempty"`
}

type ActiveNSExtensions struct {
	FilterExtensionPoints []string `plist:",omitempty" json:"filter_extensions_points,omitempty"`
}

type RotateFileVaultKey struct {
	KeyType                    string          `plist:",omitempty" json:"key_type,omitempty"`
	FileVaultUnlock            FileVaultUnlock `plist:",omitempty" json:"filevault_unlock,omitempty"`
	NewCertificate             []byte          `plist:",omitempty" json:"new_certificate,omitempty"`
	ReplyEncryptionCertificate []byte          `plist:",omitempty" json:"reply_encryption_certificate,omitempty"`
}

type FileVaultUnlock struct {
	Password                 string `plist:",omitempty" json:"password,omitempty"`
	PrivateKeyExport         []byte `plist:",omitempty" json:"private_key_export,omitempty"`
	PrivateKeyExportPassword string `plist:",omitempty" json:"private_key_export_password,omitempty"`
}

type RefreshCellularPlans struct {
	EsimServerUrl string `plist:"eSIMServerURL,omitempty" json:"esim_server_url,omitempty"`
}
