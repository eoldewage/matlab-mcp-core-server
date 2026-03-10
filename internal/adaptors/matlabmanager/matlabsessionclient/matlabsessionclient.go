// Copyright 2025-2026 The MathWorks, Inc.

package matlabsessionclient

import (
	"log"

	httpclient "github.com/matlab/matlab-mcp-core-server/internal/adaptors/http/client"
	"github.com/matlab/matlab-mcp-core-server/internal/adaptors/matlabmanager/matlabsessionclient/embeddedconnector"
	"github.com/matlab/matlab-mcp-core-server/internal/entities"
	"github.com/matlab/matlab-mcp-core-server/internal/messages"
)

type HttpClientFactory interface {
	NewClientForSelfSignedTLSServer(certificatePEM []byte) (httpclient.HttpClient, error)
}

type LoggerFactory interface {
	GetChatLogger() (*log.Logger, messages.Error)
}

type Factory struct {
	httpClientFactory HttpClientFactory
	loggerFactory     LoggerFactory
}

func NewFactory(
	httpClientFactory HttpClientFactory,
	loggerFactory LoggerFactory,
) *Factory {
	return &Factory{
		httpClientFactory: httpClientFactory,
		loggerFactory:     loggerFactory,
	}
}

func (f *Factory) New(endpoint embeddedconnector.ConnectionDetails) (entities.MATLABSessionClient, error) {
	return embeddedconnector.NewClient(endpoint, f.httpClientFactory, f.loggerFactory)
}
