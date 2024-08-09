package dto

import "mf_backup_onetime/module/ms_destination/model"

type MSDestinationRequestDto struct {
	FilterBaseDto
	DestinationId []int64 `json:"destination_id"`
}

type MSDestinationResponseDto struct {
	ResponseBaseDto
	Id    int                    `json:"id,omitempty"`
	Count int64                  `json:"count,omitempty"`
	Items []*model.MSDestination `json:"items,omitempty"`
	Item  *model.MSDestination   `json:"item,omitempty"`
}
