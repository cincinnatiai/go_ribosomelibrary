package service

import (
	"github.com/cincinnatiai/go_ribosomelibrary/models"
	campaign2 "github.com/cincinnatiai/go_ribosomelibrary/models/response/campaign"
	response "github.com/nicholaspark09/awsgorocket/model"
)

type CampaignServiceContract interface {
	Create(pipelinePartitionKey string, pipelineRangeKey string, title string, description string, secondaryId string, creatorUserId string) response.Response[models.Campaign]
	Fetch(partitionKey string, rangeKey string) response.Response[models.Campaign]
	FetchAll(pipelinePartitionKey string, pipelineRangeKey string, lastRangeKey *string) response.Response[campaign2.FetchAllResponse]
	FetchAllByUser(identityId string) response.Response[[]*models.Campaign]
	Update(campaign models.Campaign) response.Response[bool]
	Delete(partitionKey string, rangeKey string, isHardDelete bool) response.Response[bool]
}
