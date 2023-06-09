// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the vcd package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
	ApiToken pulumi.StringPtrOutput `pulumi:"apiToken"`
	// 'integrated', 'saml_adfs', 'token', and 'api_token' are the only ones supported now. 'integrated' is default.
	AuthType pulumi.StringPtrOutput `pulumi:"authType"`
	// Defines the import separation string to be used with 'terraform import'
	ImportSeparator pulumi.StringPtrOutput `pulumi:"importSeparator"`
	// Defines the full name of the logging file for API calls (requires 'logging')
	LoggingFile pulumi.StringPtrOutput `pulumi:"loggingFile"`
	// The VCD Org for API operations
	Org pulumi.StringOutput `pulumi:"org"`
	// The user password for VCD API operations.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
	SamlAdfsRptId pulumi.StringPtrOutput `pulumi:"samlAdfsRptId"`
	// The VCD Org for user authentication
	Sysorg pulumi.StringPtrOutput `pulumi:"sysorg"`
	// The token used instead of username/password for VCD API operations.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// The VCD url for VCD API operations.
	Url pulumi.StringOutput `pulumi:"url"`
	// The user name for VCD API operations.
	User pulumi.StringPtrOutput `pulumi:"user"`
	// The VDC for API operations
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Org == nil {
		return nil, errors.New("invalid value for required argument 'Org'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:vcd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If set, VCDClient will permit unverifiable SSL certificates.
	AllowUnverifiedSsl *bool `pulumi:"allowUnverifiedSsl"`
	// The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
	ApiToken *string `pulumi:"apiToken"`
	// 'integrated', 'saml_adfs', 'token', and 'api_token' are the only ones supported now. 'integrated' is default.
	AuthType *string `pulumi:"authType"`
	// Defines the import separation string to be used with 'terraform import'
	ImportSeparator *string `pulumi:"importSeparator"`
	// If set, it will enable logging of API requests and responses
	Logging *bool `pulumi:"logging"`
	// Defines the full name of the logging file for API calls (requires 'logging')
	LoggingFile *string `pulumi:"loggingFile"`
	// Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
	MaxRetryTimeout *int `pulumi:"maxRetryTimeout"`
	// The VCD Org for API operations
	Org string `pulumi:"org"`
	// The user password for VCD API operations.
	Password *string `pulumi:"password"`
	// Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
	SamlAdfsRptId *string `pulumi:"samlAdfsRptId"`
	// The VCD Org for user authentication
	Sysorg *string `pulumi:"sysorg"`
	// The token used instead of username/password for VCD API operations.
	Token *string `pulumi:"token"`
	// The VCD url for VCD API operations.
	Url string `pulumi:"url"`
	// The user name for VCD API operations.
	User *string `pulumi:"user"`
	// The VDC for API operations
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If set, VCDClient will permit unverifiable SSL certificates.
	AllowUnverifiedSsl pulumi.BoolPtrInput
	// The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
	ApiToken pulumi.StringPtrInput
	// 'integrated', 'saml_adfs', 'token', and 'api_token' are the only ones supported now. 'integrated' is default.
	AuthType pulumi.StringPtrInput
	// Defines the import separation string to be used with 'terraform import'
	ImportSeparator pulumi.StringPtrInput
	// If set, it will enable logging of API requests and responses
	Logging pulumi.BoolPtrInput
	// Defines the full name of the logging file for API calls (requires 'logging')
	LoggingFile pulumi.StringPtrInput
	// Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
	MaxRetryTimeout pulumi.IntPtrInput
	// The VCD Org for API operations
	Org pulumi.StringInput
	// The user password for VCD API operations.
	Password pulumi.StringPtrInput
	// Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
	SamlAdfsRptId pulumi.StringPtrInput
	// The VCD Org for user authentication
	Sysorg pulumi.StringPtrInput
	// The token used instead of username/password for VCD API operations.
	Token pulumi.StringPtrInput
	// The VCD url for VCD API operations.
	Url pulumi.StringInput
	// The user name for VCD API operations.
	User pulumi.StringPtrInput
	// The VDC for API operations
	Vdc pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
func (o ProviderOutput) ApiToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApiToken }).(pulumi.StringPtrOutput)
}

// 'integrated', 'saml_adfs', 'token', and 'api_token' are the only ones supported now. 'integrated' is default.
func (o ProviderOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AuthType }).(pulumi.StringPtrOutput)
}

// Defines the import separation string to be used with 'terraform import'
func (o ProviderOutput) ImportSeparator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ImportSeparator }).(pulumi.StringPtrOutput)
}

// Defines the full name of the logging file for API calls (requires 'logging')
func (o ProviderOutput) LoggingFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.LoggingFile }).(pulumi.StringPtrOutput)
}

// The VCD Org for API operations
func (o ProviderOutput) Org() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Org }).(pulumi.StringOutput)
}

// The user password for VCD API operations.
func (o ProviderOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
func (o ProviderOutput) SamlAdfsRptId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.SamlAdfsRptId }).(pulumi.StringPtrOutput)
}

// The VCD Org for user authentication
func (o ProviderOutput) Sysorg() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Sysorg }).(pulumi.StringPtrOutput)
}

// The token used instead of username/password for VCD API operations.
func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

// The VCD url for VCD API operations.
func (o ProviderOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The user name for VCD API operations.
func (o ProviderOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

// The VDC for API operations
func (o ProviderOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
