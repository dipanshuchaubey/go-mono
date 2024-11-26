package biz

import (
	bs "carthage/protos/bootcamp_service"
	"carthage/services/user_service/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type BootcampHandler struct {
	bs.UnimplementedBootcampServiceServer
}

type BootcampReviews struct {
	BootcampID string
	Reviews    []types.ReviewDetails
}
type BootcampInterface interface {
	GetBootcampsDetails(ctx context.Context, in *bs.GetBootcampsDetailsRequest) (*bs.GetBootcampsDetailsResponse, error)
}

func NewBootcampHandler() BootcampInterface {
	return &BootcampHandler{
		UnimplementedBootcampServiceServer: bs.UnimplementedBootcampServiceServer{},
	}
}

func (s *BootcampHandler) GetBootcampsDetails(ctx context.Context, in *bs.GetBootcampsDetailsRequest) (*bs.GetBootcampsDetailsResponse, error) {
	fmt.Println("GetBootcampsDetails - params BootcampIDs: ", in.BootcampIds)

	bootcampRes, httpErr := http.Get("https://bootcamper.dipanshu.work/api/v1/bootcamps")
	if httpErr != nil {
		errMsg := fmt.Errorf("GetBootcampsDetails: error getting bootcamps %v", httpErr)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	resBody, resErr := io.ReadAll(bootcampRes.Body)
	if resErr != nil {
		log.Fatal("Cannot read response body")
	}

	var bootcampInfos types.BootcampResponse
	marErr := json.Unmarshal(resBody, &bootcampInfos)

	if marErr != nil {
		errMsg := fmt.Errorf("GetBootcampsDetails: error unmarshalling json data %v", marErr)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	var wg sync.WaitGroup
	var reviewsChan = make(chan BootcampReviews)

	for _, bootcampInfo := range bootcampInfos.Data {
		wg.Add(1)

		bReviews := BootcampReviews{BootcampID: bootcampInfo.ID}

		go func() {
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					fmt.Println("recovered from panic!!")
					reviewsChan <- bReviews
				}
			}()

			reviewsRes, httpErr := http.Get(fmt.Sprintf("https://bootcamper.dipanshu.work/api/v1/bootcamps/%s/reviews", bootcampInfo.ID))
			if httpErr != nil {
				errMsg := fmt.Errorf("GetBootcampsDetails: GetBootcampsDetails: error getting reviews for BootcampID %v: %v", bootcampInfo.ID, marErr)
				fmt.Println(errMsg)
			}

			resBody, resErr := io.ReadAll(reviewsRes.Body)
			if resErr != nil {
				log.Fatal("Cannot read response body")
			}

			var reviews types.ReviewResponse
			marErr := json.Unmarshal(resBody, &reviews)
			if marErr != nil {
				errMsg := fmt.Errorf("GetBootcampsDetails: error unmarshalling json data %v", marErr)
				fmt.Println(errMsg)
			}

			bReviews.Reviews = reviews.Data
			reviewsChan <- bReviews
		}()
	}

	go func() {
		wg.Wait()
		close(reviewsChan)
	}()

	var data []*bs.GetBootcampsDetailsResponse_Data

	for review := range reviewsChan {
		for _, bootcamps := range bootcampInfos.Data {
			if review.BootcampID == bootcamps.ID {
				var info bs.GetBootcampsDetailsResponse_Data

				info.Bootcamp = &bs.BootcampInfo{
					BootcampId:  bootcamps.ID,
					Title:       bootcamps.Title,
					Description: bootcamps.Description,
					Website:     bootcamps.Website,
					Email:       bootcamps.Email,
					NameSlug:    bootcamps.NameSlug,
					Careers:     bootcamps.Careers,
				}

				info.Reviews = s.extractReviews(review.Reviews)

				data = append(data, &info)
				break
			}

		}
	}

	return &bs.GetBootcampsDetailsResponse{Data: data}, nil
}

func (s *BootcampHandler) extractReviews(reviews []types.ReviewDetails) []*bs.Review {
	var reviewsProto []*bs.Review

	for _, review := range reviews {
		reviewsProto = append(reviewsProto, &bs.Review{
			ReviewId: review.ID,
			Title:    review.Title,
			Message:  review.MessageInfos,
			UserId:   review.UserID,
			Rating:   int32(review.Rating),
		})
	}

	return reviewsProto
}
