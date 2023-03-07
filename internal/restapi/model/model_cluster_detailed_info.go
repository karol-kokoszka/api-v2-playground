/*
 * ScyllaDB Cloud API - Beta
 *
 * ### Error codes mapping  010101 Error retrieving default account  010201 Error retrieving account info  010301 Error retrieving countries  010401 Invalid request format  010402 cannot set Country from a different group  010403 Error setting account country  010501 Error getting account active credentials  010601 Error sending cross account request  010701 Error getting permissions  020101 Invalid request format  020102 Error creating user  020103 Email is already registered  020104 Email domain not allowed  020201 Invalid request format  020202 Error verifying  020301 Invalid request format  020302 General error login  020303 User must be active to login  020304 Email or password is incorrect  020304 Email or password is incorrect  020305 Too many failed login attempts  020306 The password has expired  020401 Invalid request format  020402 General logout error  020501 Error retrieving user info  020601 Invalid request format  020602 General password reset error  020603 Email does not exist or invalid  020604 User must be active to reset password  020701 Invalid request format  020702 General password reset error  020801 Invalid request format  020803 Password was already used  020802 General reset password error  020803 Password reset token is invalid or expired  020901 Invalid request format  020902 General change password error  020903 Current password is incorrect  020904 New password was already used  020905 New password is incorrect  021001 Invalid request format  021002 Error resending token  021003 User already verified  021101 Invalid request format  021102 Session token is invalid or expired  021201 Invalid request format  021202 Error updating user details  021301 Invalid request format  021302 Error creating TOTP password  021303 Invalid user password  021304 User already has MFA enabled  021401 Invalid request format  021402 Error validating TOTP password  021403 User already has MFA enabled  021404 Invalid TOTP UUID  021405 TOTP verification time expired  021406 Invalid user TOTP status  021407 Invalid user passcode  021501 Invalid request format  021502 General error TOTP login  021503 User should have MFA enabled  021504 Invalid session token  021505 Invalid session token status  021506 Session token expired  021507 Invalid user TOTP status  021508 Invalid user passcode  021601 Invalid request format  021602 Disable user TOTP general error  021603 Invalid user password  021604 User should have MFA enabled  021701 Invalid request format  021702 Validate and disable user TOTP general error  021703 Invalid user password  021704 Invalid user TOTP status  021705 Invalid user passcode  021801 Error creating monitor token  021901 Invalid request format  021902 Error validating monitor token  021903 Invalid monitor token  021904 User status is not Active  021905 User ID mismatch  030101 Invalid request format  030102 Error adding billing card  030103 Forbidden  030401 Error getting billing card info  030402 Forbidden  030201 Error retrieving billing info  030202 Error adding account billing info  030203 Forbidden  030301 Error getting account billing info  030302 Forbidden  040101 Invalid request format  040102 General error with adding a new firewall rule  040103 Firewall rule already exist  040104 Invalid firewall rule  040105 CIDR range is too big, the prefix must be at least /16  040201 General error with retrieving credentials  040301 General error with retrieving the allowed firewall list  040401 General error with retrieving the nodes list  040501 Invalid request format  040502 General error with retrieving the cluster extended info  040601 General error with retrieving the clusters list  040701 Invalid request format  040704 Provided CIDR range is too small  040702 General error creating the cluster  040703 Cluster name is already used  040704 CIDR range is too small  040705 CIDR does not belong to a private network  040706 CIDR range is too big  040707 Account is not eligible for free-tier  040708 Invalid number of nodes for free tier cluster  040709 Instance type not eligible for free tier cluster  040710 Invalid billing status  040711 Given service version requires AlternatorWriteIsolation  040714 CIDR range is too big, the prefix must be at least /16  040712 Cluster name is invalid  040713 CIDR range overlaps a reserved network  040715 Account cloud provider is not supported for jump start  040716 Account is not active  040717 Account is not eligible for jump start since billing is set  040718 Account is not eligible for jump start since there is no active cloud account  040719 Jump start is already used for this account  040720 Jump start is time expired  040801 Invalid request format  040802  General error deleting the cluster  040803 Cluster name is invalid  040901 General error deleting the firewall allowed ip  040902 Invalid firewall rule id  041001 Invalid request format  041002 General error resizing the cluster  041003 Cluster DC must be active to resize the cluster  040103 Firewall rule already exist  041107 Cluster has no VPC  041108 Cluster has multiple VPC  041103 Provided VPC is already peered  041102 Provided CIDR block overlaps with target VPC  041105 Provided VPC ID is invalid  041106 Provided AccountID is invalid  041109 CIDR range overlaps a reserved network  041101 Invalid request format  041111 General error creating VPC peering connection  041110 Account or user is not authorized for the action  041112 Invalid firewall rule  041113 CIDR range is too big, the prefix must be at least /16  041120 General error listing VPC peering connections  041130 General error deleting VPC peering connection  041140 General error describing VPC peering connection  041150 General error getting cluster VPC list  041160 General error requesting prometheus proxy  041170 General error getting prometheus proxy tokens  041171 Prometheus proxy tokens not found  041180 General error disabling prometheus proxy  041181 Prometheus Proxy tokens not found  041190 General error getting cluster requests  041191 Invalid \"type\" query parameter  041192 Invalid \"type\" query parameter  041200 General error with retrieving cluster publiccertificate  041201 Encryption must be enabled to retrieve cluster public certificate  041210 General error getting backup task history  041300 Internal Server Error  041301 DNS support is already disabled on this cluster  041302 DNS support is already enabled on this cluster  041303 There is an active request of the same type  041400 Unable to retrieve window information  041401 Unable to retrieve windows list  041402 Bad request format, rrule is missing in body  041403 Bad request format, invalid rrule format  041404 Bad request format, duration needs to be >= 0  041405 Provided windows overlaps with other windows  041406 window outside of working hours  041407 can't delete more windows, minimum required reached  041408 can't create more windows, max number reached  041409 Unable to create maintenance window  041410 Unable to update the maintenance window  041411 Unable to delete maintenance window  0420001 General error getting datacenter list  0420002 General error getting datacenter info  0420003 Ill-formed datacenter ID  0420004 Datacenter not found  0420005 Datacenter does notbelong to cluster  0420006 Datacenter is not active 0420002 General error getting datacenter info  0420007 General error requesting to add data center  0420008 Provided CIDR block overlaps with another data center  0420009 Provided CIDR block overlaps with VPC peer  0420010 CIDR range is too small  0420011 CIDR does not belong to a private network  0420012 CIDR does not belong to a private network  0420013 Forbidden 0420014 CIDR range overlaps a reserved network  0420015 General error creating datacenterrescale request  0420016 Forbidden  0420017 Current and new sizes are the same  0420018 Free tier clusters cannot be rescaled  0420019 Development clusters cannot be rescaled  0420020 Node count must be a multiplication of RF  050101 Error running health checks  060101 General error  060102 Agents need to use the ZenDesk portal  070101 Error listing the cloud accounts  070102 Forbidden  070201 Error creating a new cloud account  070202 Forbidden  070203 Invalid request format  070204 Unsupported CloudProviderID  070301 Error deleting the cloud accounts  070302 Forbidden  070303 The cloud account is used by one or more cluster DC  070401 Error updating thecloud account  070402 Forbidden  070403 CloudRoleExternalID must match the existing value  070404 BoundaryPolicyARN field AccountID must match AWSAccountID  070405 CloudPolicyARN field AccountID must match AWSAccountID  070406 CloudRoleARN field AccountID must match AWSAccountID  070407 AWSAccountID must be a valid 12-digit account ID or be empty  070408 BoundaryPolicyARN must be a valid ARN or be empty  070409 BoundaryPolicyARN must be a policy  070410 CloudPolicyARN must be a valid ARN or be empty  070411 CloudPolicyARN must be a policy  070412 CloudRoleARN must be a valid ARN or be empty  070413 CloudRoleARN must be a role  070414 The cloud account is used by one or more cluster DC  070501 Error getting the AWS boundary IAM policy document  070502 Forbidden  070601 Error getting the AWS cloud IAM policy document  070602 Forbidden  070603 Invalid request format  070604 Cloud policy name must be between 1 and 64 characters, and use only word characters or symbols: +=,./@-  070605 Unsupported cloud provider  070606 The boundary policy ARN is not set  070607 The AWS account ID is not set  070701 Error checking the cloud account policies  070702 Forbidden  070703 The cloud account properties are not complete  070801 Error getting the cloud account clusters  070802 Forbidden
 *
 * API version: 0.0.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
	"time"
)

type ClusterDetailedInfo struct {

	// ID of the cluster
	Id int32 `json:"id"`

	// ID of the account where cluster belongs to
	AccountId int32 `json:"accountId"`

	// Name of the cluster
	ClusterName string `json:"clusterName"`

	// Status of the cluster
	Status string `json:"status"`

	// ID of the cloud provider (full list [get] /deployment/provider)
	CloudProviderId int32 `json:"cloudProviderId"`

	// ID of the instance type (full list [get] /deployment/provider/{}/region/{}/instance)
	InstanceId int32 `json:"instanceId,omitempty"`

	// Version of Scylla
	ScyllaVersionId int32 `json:"scyllaVersionId"`

	// CQL or ALTERNATOR (DynamoDB)
	UserApiInterface string `json:"userApiInterface,omitempty"`

	// Defines the pricing model
	PricingModel int32 `json:"pricingModel,omitempty"`

	// Max CIDR range the user is allowed to specify on allowed ips rules, 0 disable all limitations.
	MaxAllowedCIDRRange int32 `json:"maxAllowedCIDRRange,omitempty"`

	// DNS enabled flag
	Dns bool `json:"dns,omitempty"`

	CloudProvider CloudProvider `json:"cloudProvider,omitempty"`

	ScyllaVersion ScyllaVersion `json:"scyllaVersion,omitempty"`

	Region Region `json:"region,omitempty"`

	Instance Instance `json:"instance,omitempty"`

	Dc ClusterDcInfo `json:"dc,omitempty"`

	FreeTier FreeTierInfo `json:"freeTier,omitempty"`

	JumpStart JumpStartClusterInfo `json:"jumpStart,omitempty"`

	Progress ClusterRequestProgress `json:"progress,omitempty"`

	Provisioning ProvisionType `json:"provisioning,omitempty"`

	Serverless bool `json:"serverless,omitempty"`

	// Replication Factor (RF)
	ReplicationFactor int32 `json:"replicationFactor,omitempty"`

	// Broadcast type
	BroadcastType string `json:"broadcastType,omitempty"`

	// URL of grafana (part of Monitoring Stack)
	GrafanaUrl string `json:"grafanaUrl,omitempty"`

	ClientIp string `json:"clientIp,omitempty"`

	// Creation date
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// Prometheus proxy enabled flag
	PromProxyEnabled bool `json:"promProxyEnabled,omitempty"`

	// List of CIDR formatted rules a.b.c.d/e
	AllowedIps []FirewallRule `json:"allowedIps,omitempty"`

	DataCenters []ClusterDcInfoEnriched `json:"dataCenters,omitempty"`

	Nodes []NodeInfo `json:"nodes,omitempty"`

	VpcList []ClusterNetwork `json:"vpcList,omitempty"`

	VpcPeeringList []ClusterVpcPeeringInfo `json:"vpcPeeringList,omitempty"`
}
