package library

import (
	"github.com/cincinnatiai/go_ribosomelibrary/mediator"
	"github.com/cincinnatiai/go_ribosomelibrary/pipelineservice"
	"github.com/cincinnatiai/go_ribosomelibrary/service"
	"github.com/cincinnatiai/go_ribosomelibrary/stageservice"
	"github.com/cincinnatiai/go_ribosomelibrary/stepinputservice"
	"github.com/cincinnatiai/go_ribosomelibrary/stepservice"
	metrics2 "github.com/nicholaspark09/awsgorocket/metrics"
)

type PipelineLibrary struct {
	clientId          string
	clientKey         string
	pipelineFacadeUrl string
	pipelineFacadeKey string
	metricsManager    metrics2.MetricsManagerContract
}

func ProvidePipelineLibrary(
	clientId string,
	clientKey string,
	pipelineEndpoint string,
	pipelineApiKey string,
	metricsManager metrics2.MetricsManagerContract,
) PipelineLibrary {
	return PipelineLibrary{
		clientId:          clientId,
		clientKey:         clientKey,
		pipelineFacadeUrl: pipelineEndpoint,
		pipelineFacadeKey: pipelineApiKey,
		metricsManager:    metricsManager,
	}
}

func (library *PipelineLibrary) ProvidePipelineService() pipelineservice.PipelineServiceContract {
	stageMediator := mediator.ProvideStageMediator(library.ProvideStepService())
	pipelineMediator := mediator.ProvidePipelineMediator(library.ProvideStageService())
	return pipelineservice.ProvidePipelineService(library.pipelineFacadeUrl, library.pipelineFacadeKey, pipelineMediator, stageMediator, library.metricsManager)
}

func (library *PipelineLibrary) ProvideStageService() stageservice.StageServiceRepositoryContract {
	return stageservice.ProvideStageRepository(library.pipelineFacadeUrl, library.pipelineFacadeKey, library.metricsManager)
}

func (library *PipelineLibrary) ProvideStepService() stepservice.StepServiceRepositoryContract {
	return stepservice.ProvideStepRepository(library.pipelineFacadeUrl, library.pipelineFacadeKey, library.metricsManager)
}

func (library *PipelineLibrary) ProvideCampaignService() service.CampaignServiceContract {
	return service.ProvideCampaignService(library.pipelineFacadeUrl, library.pipelineFacadeKey, library.clientId, library.clientKey, library.metricsManager)
}

func (library *PipelineLibrary) ProvideStepInputService() stepinputservice.StepInputServiceContract {
	return stepinputservice.ProvideStepInputService(library.pipelineFacadeUrl, library.pipelineFacadeKey, library.clientId, library.clientKey, library.metricsManager)
}

func (library *PipelineLibrary) ProvideCategoryService() service.CategoryServiceContract {
	return service.ProvideCategoryService(library.pipelineFacadeUrl, library.pipelineFacadeKey, library.clientId, library.clientKey, library.metricsManager)
}
