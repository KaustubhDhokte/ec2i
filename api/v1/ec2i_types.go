/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	//policiespkg "github.com/awslabs/goformation/v4/cloudformation/policies"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//tagspkg "github.com/awslabs/goformation/v4/cloudformation/tags"
	//ec2pkg "github.com/awslabs/goformation/v4/cloudformation/ec2"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EC2ISpec defines the desired state of EC2I
type EC2ISpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of EC2I. Edit EC2I_types.go to remove/update
	// ImageId      string `json:"imageId,omitempty"`
	// InstanceType string `json:"instanceType,omitempty"`
	// InstanceName string `json:"instanceName,omitempty"`
	// InstanceId   string `json:"instanceId,omitempty"`

	StackName string `json:"stackName,omitempty"`
	// AdditionalInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-additionalinfo
	AdditionalInfo string `json:"additionalInfo,omitempty"`

	// Affinity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-affinity
	Affinity string `json:"affinity,omitempty"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-availabilityzone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	Region string `json:"region,omitempty"`

	// BlockDeviceMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-blockdevicemappings
	//BlockDeviceMappings []ec2pkg.Instance_BlockDeviceMapping `json:"BlockDeviceMappings,omitempty"`

	// CpuOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-cpuoptions
	//CpuOptions *ec2pkg.Instance_CpuOptions `json:"CpuOptions,omitempty"`

	// CreditSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-creditspecification
	//CreditSpecification *ec2pkg.Instance_CreditSpecification `json:"CreditSpecification,omitempty"`

	// DisableApiTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-disableapitermination
	DisableApiTermination bool `json:"disableApiTermination,omitempty"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ebsoptimized
	EbsOptimized bool `json:"ebsOptimized,omitempty"`

	// ElasticGpuSpecifications AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-elasticgpuspecifications
	//ElasticGpuSpecifications []ec2pkg.Instance_ElasticGpuSpecification `json:"ElasticGpuSpecifications,omitempty"`

	// ElasticInferenceAccelerators AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-elasticinferenceaccelerators
	//ElasticInferenceAccelerators []ec2pkg.Instance_ElasticInferenceAccelerator `json:"ElasticInferenceAccelerators,omitempty"`

	// HibernationOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-hibernationoptions
	//HibernationOptions *ec2pkg.Instance_HibernationOptions `json:"HibernationOptions,omitempty"`

	// HostId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-hostid
	HostId string `json:"hostId,omitempty"`

	// HostResourceGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-hostresourcegrouparn
	HostResourceGroupArn string `json:"hostResourceGroupArn,omitempty"`

	// IamInstanceProfile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-iaminstanceprofile
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty"`

	// ImageId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-imageid
	ImageId string `json:"imageId,omitempty"`

	// InstanceInitiatedShutdownBehavior AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-instanceinitiatedshutdownbehavior
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// Ipv6AddressCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ipv6addresscount
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty"`

	// Ipv6Addresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ipv6addresses
	//Ipv6Addresses []ec2pkg.Instance_InstanceIpv6Address `json:"Ipv6Addresses,omitempty"`

	// KernelId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-kernelid
	KernelId string `json:"kernelId,omitempty"`

	// KeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-keyname
	KeyName string `json:"keyName,omitempty"`

	// LaunchTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-launchtemplate
	//LaunchTemplate *ec2pkg.Instance_LaunchTemplateSpecification `json:"LaunchTemplate,omitempty"`

	// LicenseSpecifications AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-licensespecifications
	//LicenseSpecifications []ec2pkg.Instance_LicenseSpecification `json:"LicenseSpecifications,omitempty"`

	// Monitoring AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-monitoring
	Monitoring bool `json:"monitoring,omitempty"`

	// NetworkInterfaces AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-networkinterfaces
	//NetworkInterfaces []ec2pkg.Instance_NetworkInterface `json:"NetworkInterfaces,omitempty"`

	// PlacementGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-placementgroupname
	PlacementGroupName string `json:"placementGroupName,omitempty"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-privateipaddress
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`

	// RamdiskId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ramdiskid
	RamdiskId string `json:"ramdiskId,omitempty"`

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-securitygroupids
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-securitygroups
	SecurityGroups []string `json:"securityGroups,omitempty"`

	// SourceDestCheck AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-sourcedestcheck
	SourceDestCheck bool `json:"sourceDestCheck,omitempty"`

	// SsmAssociations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-ssmassociations
	//SsmAssociations []ec2pkg.Instance_SsmAssociation `json:"SsmAssociations,omitempty"`

	// SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-subnetid
	SubnetId string `json:"subnetId,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-tags
	//Tags []tagspkg.Tag `json:"tags,omitempty"`

	// Tenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-tenancy
	Tenancy string `json:"tenancy,omitempty"`

	// UserData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-userdata
	UserData string `json:"userData,omitempty"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-volumes
	//Volumes []ec2pkg.Instance_Volume `json:"volumes,omitempty"`

	// AWSCloudFormationCreationPolicy represents a CloudFormation CreationPolicy
	//AWSCloudFormationCreationPolicy *policiespkg.CreationPolicy `json:"-"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	//AWSCloudFormationDeletionPolicy policiespkg.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	//AWSCloudFormationUpdateReplacePolicy policiespkg.UpdateReplacePolicy `json:"-"`

	// // AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	// AWSCloudFormationDependsOn []string `json:"-"`

	// // AWSCloudFormationMetadata stores structured data associated with this resource
	// AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// // AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	// AWSCloudFormationCondition string `json:"-"`
}

// EC2IStatus defines the observed state of EC2I
type EC2IStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// EC2I is the Schema for the ec2is API
type EC2I struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EC2ISpec   `json:"spec,omitempty"`
	Status EC2IStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EC2IList contains a list of EC2I
type EC2IList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EC2I `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EC2I{}, &EC2IList{})
}
