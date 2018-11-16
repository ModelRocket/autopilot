// Code generated by hiro; DO NOT EDIT.

// Copyright 2018 Portworx Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file at the
// root of this project.

package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dre1080/recover"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"gitlab.com/ModelRocket/sparks/cloud/provider"

	interpose "github.com/carbocation/interpose/middleware"
	sparks "gitlab.com/ModelRocket/sparks/types"

	"github.com/libopenstorage/autopilot/api/autopilot"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/collector"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/rule"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/sample"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/source"
	"github.com/libopenstorage/autopilot/api/autopilot/rest/operations/task"
)

type contextKey string

const AuthKey contextKey = "Auth"

// CollectorAPI
type CollectorAPI interface {
	// CollectorCreate is Create a new telemetry collector from the provided definition
	CollectorCreate(ctx *autopilot.Context, params collector.CollectorCreateParams) middleware.Responder
	// CollectorDelete is Returns the request collected object
	CollectorDelete(ctx *autopilot.Context, params collector.CollectorDeleteParams) middleware.Responder
	// CollectorGet is Returns the request collected object
	CollectorGet(ctx *autopilot.Context, params collector.CollectorGetParams) middleware.Responder
	// CollectorList is Returns an array of telemetry collectors defined in the system
	CollectorList(ctx *autopilot.Context, params collector.CollectorListParams) middleware.Responder
	// CollectorUpdate is Update the properties of the specified collector
	CollectorUpdate(ctx *autopilot.Context, params collector.CollectorUpdateParams) middleware.Responder
}

// RuleAPI
type RuleAPI interface {
	// RuleCreate is Create a new telemetry rule from the provided definition
	RuleCreate(ctx *autopilot.Context, params rule.RuleCreateParams) middleware.Responder
	// RuleDelete is Returns the request collected object
	RuleDelete(ctx *autopilot.Context, params rule.RuleDeleteParams) middleware.Responder
	// RuleGet is Returns the request collected object
	RuleGet(ctx *autopilot.Context, params rule.RuleGetParams) middleware.Responder
	// RuleList is Returns an array of telemetry rules defined in the system
	RuleList(ctx *autopilot.Context, params rule.RuleListParams) middleware.Responder
	// RuleUpdate is Update the properties of the specified rule
	RuleUpdate(ctx *autopilot.Context, params rule.RuleUpdateParams) middleware.Responder
}

// SampleAPI
type SampleAPI interface {
	// RecommendationsGet is Returns the recommendations for a particular sample
	RecommendationsGet(ctx *autopilot.Context, params sample.RecommendationsGetParams) middleware.Responder
	// SampleCreate is Create a new telemetry sample from the provided definition
	SampleCreate(ctx *autopilot.Context, params sample.SampleCreateParams) middleware.Responder
	// SampleDelete is Returns the request collected object
	SampleDelete(ctx *autopilot.Context, params sample.SampleDeleteParams) middleware.Responder
	// SampleGet is Returns the request collected object
	SampleGet(ctx *autopilot.Context, params sample.SampleGetParams) middleware.Responder
	// SampleList is Returns an array of telemetry samples defined in the system
	SampleList(ctx *autopilot.Context, params sample.SampleListParams) middleware.Responder
	// SampleUpdate is Update the properties of the specified sample
	SampleUpdate(ctx *autopilot.Context, params sample.SampleUpdateParams) middleware.Responder
}

// SourceAPI
type SourceAPI interface {
	// SourceCreate is Create a new telemetry source from the provided definition
	SourceCreate(ctx *autopilot.Context, params source.SourceCreateParams) middleware.Responder
	// SourceDelete is Returns the request collected object
	SourceDelete(ctx *autopilot.Context, params source.SourceDeleteParams) middleware.Responder
	// SourceGet is Returns the request collected object
	SourceGet(ctx *autopilot.Context, params source.SourceGetParams) middleware.Responder
	// SourceList is Returns an array of telemetry sources defined in the system
	SourceList(ctx *autopilot.Context, params source.SourceListParams) middleware.Responder
	// SourceUpdate is Update the properties of the specified source
	SourceUpdate(ctx *autopilot.Context, params source.SourceUpdateParams) middleware.Responder
}

// TaskAPI
type TaskAPI interface {
	// TaskGet is Returns the request task object
	TaskGet(ctx *autopilot.Context, params task.TaskGetParams) middleware.Responder
	// TaskList is Returns an array of tasks
	TaskList(ctx *autopilot.Context, params task.TaskListParams) middleware.Responder
}

type AutopilotAPI interface {
	CollectorAPI
	RuleAPI
	SampleAPI
	SourceAPI
	TaskAPI
	// Initialize is called during handler creation to perform and changes during startup
	Initialize() error

	// InitializeContext is call before the api methods are executed
	InitializeContext(principal provider.AuthToken, r *http.Request) (*autopilot.Context, error)
}

// Config is configuration for Handler
type Config struct {
	AutopilotAPI
	Logger *logrus.Logger
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// AuthBasicAuth for basic authentication
	AuthBasicAuth func(ctx context.Context, user string, pass string) (context.Context, provider.AuthToken, error)
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewAutopilotAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.BasicAuthAuth = func(ctx context.Context, user string, pass string) (context.Context, provider.AuthToken, error) {
		if c.AuthBasicAuth == nil {
			return ctx, nil, sparks.ErrNotAuthorized
		}
		return c.AuthBasicAuth(ctx, user, pass)
	}

	api.CollectorCollectorCreateHandler = collector.CollectorCreateHandlerFunc(func(params collector.CollectorCreateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.CollectorCreate(ctx, params)
	})
	api.CollectorCollectorDeleteHandler = collector.CollectorDeleteHandlerFunc(func(params collector.CollectorDeleteParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.CollectorDelete(ctx, params)
	})
	api.CollectorCollectorGetHandler = collector.CollectorGetHandlerFunc(func(params collector.CollectorGetParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.CollectorGet(ctx, params)
	})
	api.CollectorCollectorListHandler = collector.CollectorListHandlerFunc(func(params collector.CollectorListParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.CollectorList(ctx, params)
	})
	api.CollectorCollectorUpdateHandler = collector.CollectorUpdateHandlerFunc(func(params collector.CollectorUpdateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.CollectorUpdate(ctx, params)
	})
	api.SampleRecommendationsGetHandler = sample.RecommendationsGetHandlerFunc(func(params sample.RecommendationsGetParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.RecommendationsGet(ctx, params)
	})
	api.RuleRuleCreateHandler = rule.RuleCreateHandlerFunc(func(params rule.RuleCreateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.RuleCreate(ctx, params)
	})
	api.RuleRuleDeleteHandler = rule.RuleDeleteHandlerFunc(func(params rule.RuleDeleteParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.RuleDelete(ctx, params)
	})
	api.RuleRuleGetHandler = rule.RuleGetHandlerFunc(func(params rule.RuleGetParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.RuleGet(ctx, params)
	})
	api.RuleRuleListHandler = rule.RuleListHandlerFunc(func(params rule.RuleListParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.RuleList(ctx, params)
	})
	api.RuleRuleUpdateHandler = rule.RuleUpdateHandlerFunc(func(params rule.RuleUpdateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.RuleUpdate(ctx, params)
	})
	api.SampleSampleCreateHandler = sample.SampleCreateHandlerFunc(func(params sample.SampleCreateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SampleCreate(ctx, params)
	})
	api.SampleSampleDeleteHandler = sample.SampleDeleteHandlerFunc(func(params sample.SampleDeleteParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SampleDelete(ctx, params)
	})
	api.SampleSampleGetHandler = sample.SampleGetHandlerFunc(func(params sample.SampleGetParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SampleGet(ctx, params)
	})
	api.SampleSampleListHandler = sample.SampleListHandlerFunc(func(params sample.SampleListParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SampleList(ctx, params)
	})
	api.SampleSampleUpdateHandler = sample.SampleUpdateHandlerFunc(func(params sample.SampleUpdateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SampleUpdate(ctx, params)
	})
	api.SourceSourceCreateHandler = source.SourceCreateHandlerFunc(func(params source.SourceCreateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SourceCreate(ctx, params)
	})
	api.SourceSourceDeleteHandler = source.SourceDeleteHandlerFunc(func(params source.SourceDeleteParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SourceDelete(ctx, params)
	})
	api.SourceSourceGetHandler = source.SourceGetHandlerFunc(func(params source.SourceGetParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SourceGet(ctx, params)
	})
	api.SourceSourceListHandler = source.SourceListHandlerFunc(func(params source.SourceListParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SourceList(ctx, params)
	})
	api.SourceSourceUpdateHandler = source.SourceUpdateHandlerFunc(func(params source.SourceUpdateParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.SourceUpdate(ctx, params)
	})
	api.TaskTaskGetHandler = task.TaskGetHandlerFunc(func(params task.TaskGetParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.TaskGet(ctx, params)
	})
	api.TaskTaskListHandler = task.TaskListHandlerFunc(func(params task.TaskListParams, principal provider.AuthToken) middleware.Responder {
		ctx, err := c.InitializeContext(principal, params.HTTPRequest)
		if err != nil {
			return sparks.NewError(err)
		}
		return c.AutopilotAPI.TaskList(ctx, params)
	})
	api.ServerShutdown = func() {}

	logMiddleware := func(handler http.Handler) http.Handler {
		handlePanic := recover.New(&recover.Options{
			Log: c.Logger.Error,
		})

		logViaLogrus := interpose.NegroniLogrus()

		if c.InnerMiddleware != nil {
			handler = c.InnerMiddleware(handler)
		}

		return handlePanic(logViaLogrus(handler))
	}

	if err := c.AutopilotAPI.Initialize(); err != nil {
		return nil, err
	}

	return api.Serve(logMiddleware), nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}
