#!/bin/bash

# This script installs a EKS cluster with eksctl and aws cli.

set -euo pipefail

script_dir=$(dirname "${BASH_SOURCE[0]}")
source ${script_dir}/util.sh

# https://docs.aws.amazon.com/eks/latest/userguide/add-ons-images.html
declare -A registries
registries["af-south-1"]=877085696533.dkr.ecr.af-south-1.amazonaws.com/
registries["ap-east-1"]=800184023465.dkr.ecr.ap-east-1.amazonaws.com/
registries["ap-northeast-1"]=602401143452.dkr.ecr.ap-northeast-1.amazonaws.com/
registries["ap-northeast-2"]=602401143452.dkr.ecr.ap-northeast-2.amazonaws.com/
registries["ap-northeast-3"]=602401143452.dkr.ecr.ap-northeast-3.amazonaws.com/
registries["ap-south-1"]=602401143452.dkr.ecr.ap-south-1.amazonaws.com/
registries["ap-southeast-1"]=602401143452.dkr.ecr.ap-southeast-1.amazonaws.com/
registries["ap-southeast-2"]=602401143452.dkr.ecr.ap-southeast-2.amazonaws.com/
registries["ca-central-1"]=602401143452.dkr.ecr.ca-central-1.amazonaws.com/
registries["cn-north-1"]=918309763551.dkr.ecr.cn-north-1.amazonaws.com.cn/
registries["cn-northwest-1"]=961992271922.dkr.ecr.cn-northwest-1.amazonaws.com.cn/
registries["eu-central-1"]=602401143452.dkr.ecr.eu-central-1.amazonaws.com/
registries["eu-north-1"]=602401143452.dkr.ecr.eu-north-1.amazonaws.com/
registries["eu-south-1"]=590381155156.dkr.ecr.eu-south-1.amazonaws.com/
registries["eu-west-1"]=602401143452.dkr.ecr.eu-west-1.amazonaws.com/
registries["eu-west-2"]=602401143452.dkr.ecr.eu-west-2.amazonaws.com/
registries["eu-west-3"]=602401143452.dkr.ecr.eu-west-3.amazonaws.com/
registries["me-south-1"]=558608220178.dkr.ecr.me-south-1.amazonaws.com/
registries["sa-east-1"]=602401143452.dkr.ecr.sa-east-1.amazonaws.com/
registries["us-east-1"]=602401143452.dkr.ecr.us-east-1.amazonaws.com/
registries["us-east-2"]=602401143452.dkr.ecr.us-east-2.amazonaws.com/
registries["us-gov-east-1"]=151742754352.dkr.ecr.us-gov-east-1.amazonaws.com/
registries["us-gov-west-1"]=013241004608.dkr.ecr.us-gov-west-1.amazonaws.com/
registries["us-west-1"]=602401143452.dkr.ecr.us-west-1.amazonaws.com/
registries["us-west-2"]=602401143452.dkr.ecr.us-west-2.amazonaws.com/     

show_help() {
  cat >&2 <<-EOT
install_eks.sh installs a multi-node kubernetes cluster with eksctl and aws cli.
Usage: ${0##*/} [<options>]
Options:
  --help, -h
    Show help.
  --name, -n, <name>
    EKS cluster name (generated if unspecified, e.g. "attractive-sheepdog-1647980884")
  --region, -r <region>
    AWS region. Defaults to the value set in your AWS config (~/.aws/config)
  --node-type, -t <node_type>               
    node instance type (default m5.xlarge)
  --nodes, -N <num_nodes>
    total number of nodes (default 3)
  --node-volume-size, -s <volume_size> 
    node volume size in GB (default 100)
---------------------------------------------------------------------------------------------------
EOT
}

# -------------------------------------------------------------------------------------------------
# Command line parsing
# -------------------------------------------------------------------------------------------------

region=""
name="orchest-001"
num_nodes=3
node_type=m5.2xlarge
node_volume_size=100


while [[ $# -gt 0 ]]; do

  case ${1//_/-} in
    -h|--help)
      show_help >&2
      exit 1
    ;;
    -n|--name)
      name=$2
      shift
    ;;  
    -r|--region)
      region=$2
      shift
    ;;  
    -t|--node-type)
      node_type=$2
      shift
    ;;
    -N|--nodes)
      num_nodes=$2
      shift
    ;;
    -s|--node-volume-size)
      size=$2
      shift
    ;;
    *)
  esac
  shift
done

# Create eks cluster
eks-create() {
  log "Creating EKS cluster"

  local eks_options="--nodes ${num_nodes} --node-type=${node_type}"
  eks_options="${eks_options} --node-ami-family Ubuntu2004 --alb-ingress-access"
  eks_options="${eks_options} --name ${name}"

  if [[ ! -z ${region} ]]; then
    eks_options="${eks_options} --region ${region}"
  fi

  eksctl create cluster ${eks_options}

  log "EKS cluster created"
}

eks-create-iam-oidc() {
  eksctl utils associate-iam-oidc-provider --cluster ${name} --approve
}

eks-create-iam-policy() {
  aws iam create-policy \
  --policy-name AmazonEKS_EFS_CSI_Driver_Policy \
  --policy-document file://${script_dir}/iam-policy.json

  local aacount_id=$(aws sts get-caller-identity --query Account --output text)

  eksctl create iamserviceaccount \
    --name efs-csi-controller-sa \
    --namespace kube-system \
    --cluster ${name} \
    --attach-policy-arn arn:aws:iam::${aacount_id}:policy/AmazonEKS_EFS_CSI_Driver_Policy \
    --approve \
    --override-existing-serviceaccounts \
    --region ${region}
}

# https://docs.aws.amazon.com/eks/latest/userguide/efs-csi.html
install-efs() {
    helm repo add aws-efs-csi-driver https://kubernetes-sigs.github.io/aws-efs-csi-driver/
    helm repo update

    #Installing the EFS csi driver
    helm upgrade -i aws-efs-csi-driver aws-efs-csi-driver/aws-efs-csi-driver \
    --namespace kube-system \
    --set image.repository=${registries[$region]}eks/aws-efs-csi-driver \
    --set controller.serviceAccount.create=false \
    --set controller.serviceAccount.name=efs-csi-controller-sa
}

# https://docs.aws.amazon.com/eks/latest/userguide/efs-csi.html
create-efs-filesystem() {
    local vpc_id=$(aws eks describe-cluster \
    --name $name \
    --query "cluster.resourcesVpcConfig.vpcId" \
    --output text)

    local cidr_range=$(aws ec2 describe-vpcs \
    --vpc-ids $vpc_id \
    --query "Vpcs[].CidrBlock" \
    --output text)

    local security_group_id=$(aws ec2 create-security-group \
    --group-name orchestEfsSecurityGroup \
    --description "Orchest security group" \
    --vpc-id $vpc_id \
    --output text)

    aws ec2 authorize-security-group-ingress \
    --group-id $security_group_id \
    --protocol tcp \
    --port 2049 \
    --cidr $cidr_range

    local file_system_id=$(aws efs create-file-system \
    --region $region \
    --performance-mode generalPurpose \
    --query 'FileSystemId' \
    --output text)

  local subnets=$(aws ec2 describe-subnets \
  --filters "Name=vpc-id,Values=$vpc_id" \
  --query 'Subnets[*].{SubnetId: SubnetId,AvailabilityZone: AvailabilityZone,CidrBlock: CidrBlock}' \
  --output text)

  #aws efs create-mount-target \
    #--file-system-id $file_system_id \
    #--subnet-id subnet-EXAMPLEe2ba886490 \
    #--security-groups $security_group_id


  #IFS=$'\n'
  #for subnet in $subnets; do
  #  field1=$(echo "$subnet" | awk -F'|' '{printf "%s", $1}' | tr -d '"')
  #  field2=$(echo "$subnet" | awk -F'|' '{printf "%s", $2}' | tr -d '"')

  #  echo "$field1 fff $field2" 
  #done

}

bin=$(eval find-binary eksctl)
if [[ -z ${bin} ]]; then
  log "Failed to find binary ${bin}}"
  exit 1
fi

bin=$(eval find-binary aws)
if [[ -z ${bin} ]]; then
  log "Failed to find binary ${bin}}"
  exit 1
fi

# Set the region if not provided
if [[ -z ${region} ]]; then 
  region=$(aws configure get region)
fi

#eks-create
#eks-create-iam-oidc
#eks-create-iam-policy
#install-efs
#create-efs-filesystem
