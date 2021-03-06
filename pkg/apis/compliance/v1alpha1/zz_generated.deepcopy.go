// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceCheckResult) DeepCopyInto(out *ComplianceCheckResult) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceCheckResult.
func (in *ComplianceCheckResult) DeepCopy() *ComplianceCheckResult {
	if in == nil {
		return nil
	}
	out := new(ComplianceCheckResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceCheckResult) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceCheckResultList) DeepCopyInto(out *ComplianceCheckResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComplianceCheckResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceCheckResultList.
func (in *ComplianceCheckResultList) DeepCopy() *ComplianceCheckResultList {
	if in == nil {
		return nil
	}
	out := new(ComplianceCheckResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceCheckResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceRemediation) DeepCopyInto(out *ComplianceRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceRemediation.
func (in *ComplianceRemediation) DeepCopy() *ComplianceRemediation {
	if in == nil {
		return nil
	}
	out := new(ComplianceRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceRemediationList) DeepCopyInto(out *ComplianceRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComplianceRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceRemediationList.
func (in *ComplianceRemediationList) DeepCopy() *ComplianceRemediationList {
	if in == nil {
		return nil
	}
	out := new(ComplianceRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceRemediationSpec) DeepCopyInto(out *ComplianceRemediationSpec) {
	*out = *in
	out.ComplianceRemediationSpecMeta = in.ComplianceRemediationSpecMeta
	if in.MachineConfigContents != nil {
		in, out := &in.MachineConfigContents, &out.MachineConfigContents
		*out = (*in).DeepCopy()
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceRemediationSpec.
func (in *ComplianceRemediationSpec) DeepCopy() *ComplianceRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(ComplianceRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceRemediationSpecMeta) DeepCopyInto(out *ComplianceRemediationSpecMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceRemediationSpecMeta.
func (in *ComplianceRemediationSpecMeta) DeepCopy() *ComplianceRemediationSpecMeta {
	if in == nil {
		return nil
	}
	out := new(ComplianceRemediationSpecMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceRemediationStatus) DeepCopyInto(out *ComplianceRemediationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceRemediationStatus.
func (in *ComplianceRemediationStatus) DeepCopy() *ComplianceRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(ComplianceRemediationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceScan) DeepCopyInto(out *ComplianceScan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceScan.
func (in *ComplianceScan) DeepCopy() *ComplianceScan {
	if in == nil {
		return nil
	}
	out := new(ComplianceScan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceScan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceScanList) DeepCopyInto(out *ComplianceScanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComplianceScan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceScanList.
func (in *ComplianceScanList) DeepCopy() *ComplianceScanList {
	if in == nil {
		return nil
	}
	out := new(ComplianceScanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceScanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceScanSpec) DeepCopyInto(out *ComplianceScanSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TailoringConfigMap != nil {
		in, out := &in.TailoringConfigMap, &out.TailoringConfigMap
		*out = new(TailoringConfigMapRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceScanSpec.
func (in *ComplianceScanSpec) DeepCopy() *ComplianceScanSpec {
	if in == nil {
		return nil
	}
	out := new(ComplianceScanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceScanSpecWrapper) DeepCopyInto(out *ComplianceScanSpecWrapper) {
	*out = *in
	in.ComplianceScanSpec.DeepCopyInto(&out.ComplianceScanSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceScanSpecWrapper.
func (in *ComplianceScanSpecWrapper) DeepCopy() *ComplianceScanSpecWrapper {
	if in == nil {
		return nil
	}
	out := new(ComplianceScanSpecWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceScanStatus) DeepCopyInto(out *ComplianceScanStatus) {
	*out = *in
	out.ResultsStorage = in.ResultsStorage
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceScanStatus.
func (in *ComplianceScanStatus) DeepCopy() *ComplianceScanStatus {
	if in == nil {
		return nil
	}
	out := new(ComplianceScanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceScanStatusWrapper) DeepCopyInto(out *ComplianceScanStatusWrapper) {
	*out = *in
	out.ComplianceScanStatus = in.ComplianceScanStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceScanStatusWrapper.
func (in *ComplianceScanStatusWrapper) DeepCopy() *ComplianceScanStatusWrapper {
	if in == nil {
		return nil
	}
	out := new(ComplianceScanStatusWrapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSuite) DeepCopyInto(out *ComplianceSuite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSuite.
func (in *ComplianceSuite) DeepCopy() *ComplianceSuite {
	if in == nil {
		return nil
	}
	out := new(ComplianceSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceSuite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSuiteList) DeepCopyInto(out *ComplianceSuiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComplianceSuite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSuiteList.
func (in *ComplianceSuiteList) DeepCopy() *ComplianceSuiteList {
	if in == nil {
		return nil
	}
	out := new(ComplianceSuiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComplianceSuiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSuiteSpec) DeepCopyInto(out *ComplianceSuiteSpec) {
	*out = *in
	if in.Scans != nil {
		in, out := &in.Scans, &out.Scans
		*out = make([]ComplianceScanSpecWrapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSuiteSpec.
func (in *ComplianceSuiteSpec) DeepCopy() *ComplianceSuiteSpec {
	if in == nil {
		return nil
	}
	out := new(ComplianceSuiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceSuiteStatus) DeepCopyInto(out *ComplianceSuiteStatus) {
	*out = *in
	if in.ScanStatuses != nil {
		in, out := &in.ScanStatuses, &out.ScanStatuses
		*out = make([]ComplianceScanStatusWrapper, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceSuiteStatus.
func (in *ComplianceSuiteStatus) DeepCopy() *ComplianceSuiteStatus {
	if in == nil {
		return nil
	}
	out := new(ComplianceSuiteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixDefinition) DeepCopyInto(out *FixDefinition) {
	*out = *in
	if in.FixObject != nil {
		in, out := &in.FixObject, &out.FixObject
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixDefinition.
func (in *FixDefinition) DeepCopy() *FixDefinition {
	if in == nil {
		return nil
	}
	out := new(FixDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputRef) DeepCopyInto(out *OutputRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputRef.
func (in *OutputRef) DeepCopy() *OutputRef {
	if in == nil {
		return nil
	}
	out := new(OutputRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Profile) DeepCopyInto(out *Profile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ProfileRule, len(*in))
		copy(*out, *in)
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]ProfileValue, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Profile.
func (in *Profile) DeepCopy() *Profile {
	if in == nil {
		return nil
	}
	out := new(Profile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Profile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBundle) DeepCopyInto(out *ProfileBundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBundle.
func (in *ProfileBundle) DeepCopy() *ProfileBundle {
	if in == nil {
		return nil
	}
	out := new(ProfileBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileBundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBundleList) DeepCopyInto(out *ProfileBundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProfileBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBundleList.
func (in *ProfileBundleList) DeepCopy() *ProfileBundleList {
	if in == nil {
		return nil
	}
	out := new(ProfileBundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileBundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBundleSpec) DeepCopyInto(out *ProfileBundleSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBundleSpec.
func (in *ProfileBundleSpec) DeepCopy() *ProfileBundleSpec {
	if in == nil {
		return nil
	}
	out := new(ProfileBundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBundleStatus) DeepCopyInto(out *ProfileBundleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBundleStatus.
func (in *ProfileBundleStatus) DeepCopy() *ProfileBundleStatus {
	if in == nil {
		return nil
	}
	out := new(ProfileBundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileList) DeepCopyInto(out *ProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Profile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileList.
func (in *ProfileList) DeepCopy() *ProfileList {
	if in == nil {
		return nil
	}
	out := new(ProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AvailableFixes != nil {
		in, out := &in.AvailableFixes, &out.AvailableFixes
		*out = make([]FixDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleList) DeepCopyInto(out *RuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *RuleList) DeepCopy() *RuleList {
	if in == nil {
		return nil
	}
	out := new(RuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleReferenceSpec) DeepCopyInto(out *RuleReferenceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleReferenceSpec.
func (in *RuleReferenceSpec) DeepCopy() *RuleReferenceSpec {
	if in == nil {
		return nil
	}
	out := new(RuleReferenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageReference) DeepCopyInto(out *StorageReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageReference.
func (in *StorageReference) DeepCopy() *StorageReference {
	if in == nil {
		return nil
	}
	out := new(StorageReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TailoredProfile) DeepCopyInto(out *TailoredProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TailoredProfile.
func (in *TailoredProfile) DeepCopy() *TailoredProfile {
	if in == nil {
		return nil
	}
	out := new(TailoredProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TailoredProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TailoredProfileList) DeepCopyInto(out *TailoredProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TailoredProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TailoredProfileList.
func (in *TailoredProfileList) DeepCopy() *TailoredProfileList {
	if in == nil {
		return nil
	}
	out := new(TailoredProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TailoredProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TailoredProfileSpec) DeepCopyInto(out *TailoredProfileSpec) {
	*out = *in
	if in.EnableRules != nil {
		in, out := &in.EnableRules, &out.EnableRules
		*out = make([]RuleReferenceSpec, len(*in))
		copy(*out, *in)
	}
	if in.DisableRules != nil {
		in, out := &in.DisableRules, &out.DisableRules
		*out = make([]RuleReferenceSpec, len(*in))
		copy(*out, *in)
	}
	if in.SetValues != nil {
		in, out := &in.SetValues, &out.SetValues
		*out = make([]VariableValueSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TailoredProfileSpec.
func (in *TailoredProfileSpec) DeepCopy() *TailoredProfileSpec {
	if in == nil {
		return nil
	}
	out := new(TailoredProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TailoredProfileStatus) DeepCopyInto(out *TailoredProfileStatus) {
	*out = *in
	out.OutputRef = in.OutputRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TailoredProfileStatus.
func (in *TailoredProfileStatus) DeepCopy() *TailoredProfileStatus {
	if in == nil {
		return nil
	}
	out := new(TailoredProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TailoringConfigMapRef) DeepCopyInto(out *TailoringConfigMapRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TailoringConfigMapRef.
func (in *TailoringConfigMapRef) DeepCopy() *TailoringConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(TailoringConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueSelection) DeepCopyInto(out *ValueSelection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueSelection.
func (in *ValueSelection) DeepCopy() *ValueSelection {
	if in == nil {
		return nil
	}
	out := new(ValueSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Selections != nil {
		in, out := &in.Selections, &out.Selections
		*out = make([]ValueSelection, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Variable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VariableList) DeepCopyInto(out *VariableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Variable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VariableList.
func (in *VariableList) DeepCopy() *VariableList {
	if in == nil {
		return nil
	}
	out := new(VariableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VariableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VariableValueSpec) DeepCopyInto(out *VariableValueSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VariableValueSpec.
func (in *VariableValueSpec) DeepCopy() *VariableValueSpec {
	if in == nil {
		return nil
	}
	out := new(VariableValueSpec)
	in.DeepCopyInto(out)
	return out
}
