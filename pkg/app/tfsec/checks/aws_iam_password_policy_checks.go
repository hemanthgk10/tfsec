package checks

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/parser"
)

// AWSIAMPasswordReusePrevention See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordReusePrevention scanner.RuleID = "AWS037"
const AWSIAMPasswordReusePreventionDescription scanner.RuleDescription = "IAM Password policy should prevent password reuse."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordReusePrevention,
		Description:    AWSIAMPasswordReusePreventionDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("password_reuse_prevention"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not have a password reuse prevention count set.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Number {
				value, _ := attr.Value().AsBigFloat().Float64()
				if value < 5 {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' has a password reuse count less than 5.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}

// AWSIAMPasswordExpiry See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordExpiry scanner.RuleID = "AWS038"
const AWSIAMPasswordExpiryDescription scanner.RuleDescription = "IAM Password policy should have expiry greater than or equal to 90 days."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordExpiry,
		Description:    AWSIAMPasswordExpiryDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("max_password_age"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not have a max password age set.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Number {
				value, _ := attr.Value().AsBigFloat().Float64()
				if value < 90 {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' has a max age set which is less than 90 days.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}

// AWSIAMPasswordMinimumLength See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordMinimumLength scanner.RuleID = "AWS039"
const AWSIAMPasswordMinimumLengthDescription scanner.RuleDescription = "IAM Password policy should have minimum password length of 14 or more characters."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordMinimumLength,
		Description:    AWSIAMPasswordMinimumLengthDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("minimum_password_length"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not have a minimum password length set.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Number {
				value, _ := attr.Value().AsBigFloat().Float64()
				if value < 14 {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' has a minimum password length which is less than 14 characters.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}

// AWSIAMPasswordRequiresSymbol See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordRequiresSymbol scanner.RuleID = "AWS040"
const AWSIAMPasswordRequiresSymbolDescription scanner.RuleDescription = "IAM Password policy should have requirement for at least one symbol in the password."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordRequiresSymbol,
		Description:    AWSIAMPasswordRequiresSymbolDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("require_symbols"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not require a symbol in the password.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Bool {
				if attr.Value().False() {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' explicitly specifies not requiring at least one symbol in the password.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}

// AWSIAMPasswordRequiresNumber See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordRequiresNumber scanner.RuleID = "AWS041"
const AWSIAMPasswordRequiresNumberDescription scanner.RuleDescription = "IAM Password policy should have requirement for at least one number in the password."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordRequiresNumber,
		Description:    AWSIAMPasswordRequiresNumberDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("require_numbers"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not require a number in the password.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Bool {
				if attr.Value().False() {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' explicitly specifies not requiring at least one number in the password.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}

// AWSIAMPasswordRequiresLowercaseCharacter See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordRequiresLowercaseCharacter scanner.RuleID = "AWS042"
const AWSIAMPasswordRequiresLowercaseCharacterDescription scanner.RuleDescription = "IAM Password policy should have requirement for at least one lowercase character."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordRequiresLowercaseCharacter,
		Description:    AWSIAMPasswordRequiresLowercaseCharacterDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("require_lowercase_characters"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not require a lowercase character in the password.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Bool {
				if attr.Value().False() {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' explicitly specifies not requiring at least lowercase character in the password.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}

// AWSIAMPasswordRequiresUppercaseCharacter See https://github.com/tfsec/tfsec#included-checks for check info
const AWSIAMPasswordRequiresUppercaseCharacter scanner.RuleID = "AWS043"
const AWSIAMPasswordRequiresUppercaseCharacterDescription scanner.RuleDescription = "IAM Password policy should have requirement for at least one uppercase character."

func init() {
	scanner.RegisterCheck(scanner.Check{
		Code:           AWSIAMPasswordRequiresUppercaseCharacter,
		Description:    AWSIAMPasswordRequiresUppercaseCharacterDescription,
		Provider:       scanner.AWSProvider,
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"aws_iam_account_password_policy"},
		CheckFunc: func(check *scanner.Check, block *parser.Block, _ *scanner.Context) []scanner.Result {
			if attr := block.GetAttribute("require_uppercase_characters"); attr == nil {
				return []scanner.Result{
					check.NewResult(
						fmt.Sprintf("Resource '%s' does not require an uppercase character in the password.", block.Name()),
						block.Range(),
						scanner.SeverityWarning,
					),
				}
			} else if attr.Value().Type() == cty.Bool {
				if attr.Value().False() {
					return []scanner.Result{
						check.NewResult(
							fmt.Sprintf("Resource '%s' explicitly specifies not requiring at least one uppercase character in the password.", block.Name()),
							block.Range(),
							scanner.SeverityWarning,
						),
					}
				}
			}
			return nil
		},
	})
}
