// Code generated by "go generate". Please don't change this file directly.

package lookoutequipment

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// InferenceScheduler_DataOutputConfiguration AWS CloudFormation Resource (AWS::LookoutEquipment::InferenceScheduler.DataOutputConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html
type InferenceScheduler_DataOutputConfiguration struct {

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html#cfn-lookoutequipment-inferencescheduler-dataoutputconfiguration-kmskeyid
	KmsKeyId *string `json:"KmsKeyId,omitempty"`

	// S3OutputConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html#cfn-lookoutequipment-inferencescheduler-dataoutputconfiguration-s3outputconfiguration
	S3OutputConfiguration *InferenceScheduler_S3OutputConfiguration `json:"S3OutputConfiguration"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *InferenceScheduler_DataOutputConfiguration) AWSCloudFormationType() string {
	return "AWS::LookoutEquipment::InferenceScheduler.DataOutputConfiguration"
}
