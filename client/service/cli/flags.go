// nolint
package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagName              = "name"
	FlagDescription       = "description"
	FlagTags              = "tags"
	FlagAuthorDescription = "author-description"
	FlagSchemas           = "schemas"
	FlagServiceName       = "service-name"
	FlagDeposit           = "deposit"
	FlagPricing           = "pricing"
	FlagMinRespTime       = "min-resp-time"
	FlagProviders         = "providers"
	FlagServiceFeeCap     = "service-fee-cap"
	FlagTimeout           = "timeout"
	FlagData              = "data"
	FlagSuperMode         = "super-mode"
	FlagRepeated          = "repeated"
	FlagFrequency         = "frequency"
	FlagTotal             = "total"
	FlagRequestID         = "request-id"
	FlagResult            = "result"
)

// common flagsets to add to various functions
var (
	FsServiceDefine               = flag.NewFlagSet("", flag.ContinueOnError)
	FsServiceBind                 = flag.NewFlagSet("", flag.ContinueOnError)
	FsServiceUpdateBinding        = flag.NewFlagSet("", flag.ContinueOnError)
	FsServiceEnableBinding        = flag.NewFlagSet("", flag.ContinueOnError)
	FsServiceCall                 = flag.NewFlagSet("", flag.ContinueOnError)
	FsServiceRespond              = flag.NewFlagSet("", flag.ContinueOnError)
	FsServiceUpdateRequestContext = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsServiceDefine.String(FlagName, "", "service name")
	FsServiceDefine.String(FlagDescription, "", "service description")
	FsServiceDefine.StringSlice(FlagTags, []string{}, "service tags")
	FsServiceDefine.String(FlagAuthorDescription, "", "service author description")
	FsServiceDefine.String(FlagSchemas, "", "interface schemas content or file path")

	FsServiceBind.String(FlagServiceName, "", "service name")
	FsServiceBind.String(FlagDeposit, "", "deposit of the binding")
	FsServiceBind.String(FlagPricing, "", "pricing content or file path, which is an instance of the Irishub Service Pricing schema")
	FsServiceBind.Uint64(FlagMinRespTime, 0, "minimum response time")

	FsServiceUpdateBinding.String(FlagDeposit, "", "added deposit for the binding")
	FsServiceUpdateBinding.String(FlagPricing, "", "pricing content or file path, which is an instance of the Irishub Service Pricing schema")
	FsServiceUpdateBinding.Uint64(FlagMinRespTime, 0, "minimum response time, not updated if set to 0")

	FsServiceEnableBinding.String(FlagDeposit, "", "added deposit for enabling the binding")

	FsServiceCall.String(FlagServiceName, "", "service name")
	FsServiceCall.StringSlice(FlagProviders, []string{}, "provider list to request")
	FsServiceCall.String(FlagServiceFeeCap, "", "maximum service fee to pay for a single request")
	FsServiceCall.String(FlagData, "", "content or file path of the request input, which is an Input JSON schema instance")
	FsServiceCall.Uint64(FlagTimeout, 0, "request timeout")
	FsServiceCall.Bool(FlagSuperMode, false, "indicate if the signer is a super user")
	FsServiceCall.Bool(FlagRepeated, false, "indicate if the request is repetitive")
	FsServiceCall.Uint64(FlagFrequency, 0, "request frequency when repeated, default to timeout")
	FsServiceCall.Int64(FlagTotal, 0, "request count when repeated, -1 means unlimited")

	FsServiceUpdateRequestContext.StringSlice(FlagProviders, []string{}, "provider list to request, not updated if empty")
	FsServiceUpdateRequestContext.String(FlagServiceFeeCap, "", "maximum service fee to pay for a single request, not updated if empty")
	FsServiceUpdateRequestContext.Uint64(FlagTimeout, 0, "request timeout, not updated if set to 0")
	FsServiceUpdateRequestContext.Uint64(FlagFrequency, 0, "request frequency, not updated if set to 0")
	FsServiceUpdateRequestContext.Int64(FlagTotal, 0, "request count, not updated if set to 0")

	FsServiceRespond.String(FlagRequestID, "", "ID of the request to respond to")
	FsServiceRespond.String(FlagResult, "", "content or file path of the response result, which is an Result JSON schema instance")
	FsServiceRespond.String(FlagData, "", "content or file path of the response output, which is an Output JSON schema instance")
}