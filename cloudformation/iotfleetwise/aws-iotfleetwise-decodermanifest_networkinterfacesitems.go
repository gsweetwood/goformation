// Code generated by "go generate". Please don't change this file directly.

package iotfleetwise

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DecoderManifest_NetworkInterfacesItems AWS CloudFormation Resource (AWS::IoTFleetWise::DecoderManifest.NetworkInterfacesItems)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html
type DecoderManifest_NetworkInterfacesItems struct {

	// CanInterface AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-caninterface
	CanInterface *DecoderManifest_CanInterface `json:"CanInterface,omitempty"`

	// InterfaceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-interfaceid
	InterfaceId string `json:"InterfaceId"`

	// ObdInterface AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-obdinterface
	ObdInterface *DecoderManifest_ObdInterface `json:"ObdInterface,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-type
	Type string `json:"Type"`

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
func (r *DecoderManifest_NetworkInterfacesItems) AWSCloudFormationType() string {
	return "AWS::IoTFleetWise::DecoderManifest.NetworkInterfacesItems"
}