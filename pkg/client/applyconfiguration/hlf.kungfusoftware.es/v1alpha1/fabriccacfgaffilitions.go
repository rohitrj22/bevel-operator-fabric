// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricCACFGAffilitionsApplyConfiguration represents an declarative configuration of the FabricCACFGAffilitions type for use
// with apply.
type FabricCACFGAffilitionsApplyConfiguration struct {
	AllowRemove *bool `json:"allowRemove,omitempty"`
}

// FabricCACFGAffilitionsApplyConfiguration constructs an declarative configuration of the FabricCACFGAffilitions type for use with
// apply.
func FabricCACFGAffilitions() *FabricCACFGAffilitionsApplyConfiguration {
	return &FabricCACFGAffilitionsApplyConfiguration{}
}

// WithAllowRemove sets the AllowRemove field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowRemove field is set to the value of the last call.
func (b *FabricCACFGAffilitionsApplyConfiguration) WithAllowRemove(value bool) *FabricCACFGAffilitionsApplyConfiguration {
	b.AllowRemove = &value
	return b
}
