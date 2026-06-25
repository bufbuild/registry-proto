// Copyright 2023-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command buf-plugin-iam is a custom buf lint plugin for the registry IAM
// annotations.
package main

import (
	"context"

	"buf.build/go/bufplugin/check"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

const (
	rpcPermissionSetRuleID = "RPC_PERMISSION_SET"
	methodOptionFullName   = "buf.registry.priv.extension.v1beta1.method"
)

func main() {
	check.Main(&check.Spec{
		Rules: []*check.RuleSpec{
			{
				ID:      rpcPermissionSetRuleID,
				Default: true,
				Purpose: "Checks that every RPC declares exactly one permission requirement via the (" + methodOptionFullName + ") option.",
				Type:    check.RuleTypeLint,
				Handler: check.RuleHandlerFunc(checkRPCPermissionSet),
			},
		},
	})
}

func checkRPCPermissionSet(_ context.Context, responseWriter check.ResponseWriter, request check.Request) error {
	// buf hands custom MethodOptions extensions to the plugin as unresolved
	// bytes, so build an extension registry from every file in the image and
	// re-parse each method's options against it.
	resolver, err := extensionResolver(request)
	if err != nil {
		return err
	}
	for _, fileDescriptor := range request.FileDescriptors() {
		if fileDescriptor.IsImport() {
			continue
		}
		services := fileDescriptor.ProtoreflectFileDescriptor().Services()
		for i := 0; i < services.Len(); i++ {
			methods := services.Get(i).Methods()
			for j := 0; j < methods.Len(); j++ {
				method := methods.Get(j)
				got, err := permissionFieldsSet(method, resolver)
				if err != nil {
					return err
				}
				if got != 1 {
					responseWriter.AddAnnotation(
						check.WithDescriptor(method),
						check.WithMessagef(
							"RPC %q must set exactly one of required_permission or no_required_permission via (%s); found %d.",
							method.Name(),
							methodOptionFullName,
							got,
						),
					)
				}
			}
		}
	}
	return nil
}

// extensionResolver builds a registry of every extension defined across the
// image so custom options can be resolved.
func extensionResolver(request check.Request) (*protoregistry.Types, error) {
	types := &protoregistry.Types{}
	for _, fileDescriptor := range request.FileDescriptors() {
		extensions := fileDescriptor.ProtoreflectFileDescriptor().Extensions()
		for i := 0; i < extensions.Len(); i++ {
			if err := types.RegisterExtension(dynamicpb.NewExtensionType(extensions.Get(i))); err != nil {
				return nil, err
			}
		}
	}
	return types, nil
}

// permissionFieldsSet reports how many of the three permission fields are set on
// the method's (buf.registry.priv.extension.v1beta1.method) option. It returns 0
// when the option is absent, so a missing option and an empty option both fail.
func permissionFieldsSet(method protoreflect.MethodDescriptor, resolver *protoregistry.Types) (int, error) {
	options := method.Options()
	if options == nil {
		return 0, nil
	}
	optionBytes, err := proto.Marshal(options)
	if err != nil {
		return 0, err
	}
	resolved := &descriptorpb.MethodOptions{}
	if err := (proto.UnmarshalOptions{Resolver: resolver}).Unmarshal(optionBytes, resolved); err != nil {
		return 0, err
	}
	count := 0
	resolved.ProtoReflect().Range(func(fieldDescriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		if !fieldDescriptor.IsExtension() || string(fieldDescriptor.FullName()) != methodOptionFullName {
			return true
		}
		constraints := value.Message()
		fields := constraints.Descriptor().Fields()
		if field := fields.ByName("required_permission"); field != nil && constraints.Get(field).List().Len() > 0 {
			count++
		}
		if field := fields.ByName("no_required_permission"); field != nil && constraints.Get(field).Bool() {
			count++
		}
		return false
	})
	return count, nil
}
