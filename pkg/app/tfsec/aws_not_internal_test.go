package tfsec

import (
	"testing"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/checks"
)

func Test_AWSNotInternal(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.RuleID
		mustExcludeResultCode scanner.RuleID
	}{
		{
			name: "check aws_alb when not internal",
			source: `
resource "aws_alb" "my-resource" {
	internal = false
}`,
			mustIncludeResultCode: checks.AWSExternallyExposedLoadBalancer,
		},
		{
			name: "check aws_elb when not internal",
			source: `
resource "aws_elb" "my-resource" {
	internal = false
}`,
			mustIncludeResultCode: checks.AWSExternallyExposedLoadBalancer,
		},
		{
			name: "check aws_lb when not internal",
			source: `
resource "aws_lb" "my-resource" {
	internal = false
}`,
			mustIncludeResultCode: checks.AWSExternallyExposedLoadBalancer,
		},
		{
			name: "check aws_lb when not explicitly marked as internal",
			source: `
resource "aws_lb" "my-resource" {
}`,
			mustIncludeResultCode: checks.AWSExternallyExposedLoadBalancer,
		},
		{
			name: "check aws_lb when explicitly marked as internal",
			source: `
resource "aws_lb" "my-resource" {
	internal = true
}`,
			mustExcludeResultCode: checks.AWSExternallyExposedLoadBalancer,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
