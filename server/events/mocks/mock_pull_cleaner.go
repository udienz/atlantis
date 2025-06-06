// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: PullCleaner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockPullCleaner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPullCleaner(options ...pegomock.Option) *MockPullCleaner {
	mock := &MockPullCleaner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPullCleaner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPullCleaner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPullCleaner) CleanUpPull(logger logging.SimpleLogging, repo models.Repo, pull models.PullRequest) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPullCleaner().")
	}
	_params := []pegomock.Param{logger, repo, pull}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("CleanUpPull", _params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(error)
		}
	}
	return _ret0
}

func (mock *MockPullCleaner) VerifyWasCalledOnce() *VerifierMockPullCleaner {
	return &VerifierMockPullCleaner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPullCleaner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPullCleaner {
	return &VerifierMockPullCleaner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPullCleaner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPullCleaner {
	return &VerifierMockPullCleaner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPullCleaner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPullCleaner {
	return &VerifierMockPullCleaner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPullCleaner struct {
	mock                   *MockPullCleaner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPullCleaner) CleanUpPull(logger logging.SimpleLogging, repo models.Repo, pull models.PullRequest) *MockPullCleaner_CleanUpPull_OngoingVerification {
	_params := []pegomock.Param{logger, repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CleanUpPull", _params, verifier.timeout)
	return &MockPullCleaner_CleanUpPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPullCleaner_CleanUpPull_OngoingVerification struct {
	mock              *MockPullCleaner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPullCleaner_CleanUpPull_OngoingVerification) GetCapturedArguments() (logging.SimpleLogging, models.Repo, models.PullRequest) {
	logger, repo, pull := c.GetAllCapturedArguments()
	return logger[len(logger)-1], repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockPullCleaner_CleanUpPull_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.SimpleLogging, _param1 []models.Repo, _param2 []models.PullRequest) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]logging.SimpleLogging, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(logging.SimpleLogging)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]models.Repo, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(models.Repo)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]models.PullRequest, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(models.PullRequest)
			}
		}
	}
	return
}
