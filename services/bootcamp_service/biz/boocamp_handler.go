package biz

import (
	"bytes"
	pbrs "carthage/protos/bootcamp_service/response"
	pbty "carthage/protos/bootcamp_service/types"
	"carthage/services/bootcamp_service/biz/interfaces"
	"carthage/services/bootcamp_service/config"
	"carthage/services/bootcamp_service/constants"
	"carthage/services/bootcamp_service/dto"
	"carthage/services/bootcamp_service/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type BootcampHandler struct {
	AuthToken    string
	AuthTokenExp time.Time
	cnf          config.Config
}

func NewBootcampHandler(config config.Config) interfaces.BootcampInterface {
	return &BootcampHandler{
		AuthToken:    constants.EmptyString,
		AuthTokenExp: time.Now(),
		cnf:          config,
	}
}

func (s *BootcampHandler) GetBootcampsDetails(ctx context.Context) ([]*pbrs.GetBootcampsDetailsResponse_Data, error) {
	bootcampRes, httpErr := http.Get(s.cnf.EndPoints.GetBootcamps)
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
	var reviewsChan = make(chan types.BootcampReviews)

	for _, bootcampInfo := range bootcampInfos.Data {
		wg.Add(1)

		bReviews := types.BootcampReviews{BootcampID: bootcampInfo.ID}

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

	var data []*pbrs.GetBootcampsDetailsResponse_Data

	for review := range reviewsChan {
		for _, bootcamps := range bootcampInfos.Data {
			if review.BootcampID == bootcamps.ID {
				var info pbrs.GetBootcampsDetailsResponse_Data

				info.Bootcamp = &pbty.BootcampInfo{
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

	return data, nil
}

func (s *BootcampHandler) CreateBootcamp(ctx context.Context, body dto.CreateBootcampBody) (*pbty.BootcampInfo, error) {
	// 1. Validate request body
	if body.Title == constants.EmptyString || body.Email == constants.EmptyString {
		errMsg := fmt.Errorf("biz.CreateBootcamp missing title or email params in request body")
		fmt.Println(errMsg)
		return nil, errMsg
	}

	// 2. Create Bootcamp
	httpBody, marErr := json.Marshal(body)
	if marErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error marshalling request body: %v", marErr)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	var bootcamp types.CreateBootcampResponse

	fmt.Println("biz.CreateBootcamp creating bootcamp with title: ", body.Title)
	if postErr := s.postRequest(s.cnf.EndPoints.PostBootcamp, httpBody, &bootcamp, true); postErr != nil {
		return nil, postErr
	}

	if !bootcamp.Success {
		errMsg := fmt.Errorf("biz.CreateBootcamp error creating bootcamp: %v", bootcamp.Error)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	var res pbty.BootcampInfo
	bootcamp.Data.ToProro(&res)

	return &res, nil
}

func (s *BootcampHandler) extractReviews(reviews []types.ReviewDetails) []*pbty.Review {
	var reviewsProto []*pbty.Review

	for _, review := range reviews {
		reviewsProto = append(reviewsProto, &pbty.Review{
			ReviewId: review.ID,
			Title:    review.Title,
			Message:  review.MessageInfos,
			UserId:   review.UserID,
			Rating:   int32(review.Rating),
		})
	}

	return reviewsProto
}

func (s *BootcampHandler) postRequest(url string, httpBody []byte, resBody any, auth bool) error {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(httpBody))
	if err != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error creating POST request: %v", err)
		fmt.Println(errMsg)
		return errMsg
	}

	if auth {
		token := s.getAuthToken()
		if token == constants.EmptyString {
			errMsg := fmt.Errorf("biz.CreateBootcamp error getting auth token")
			fmt.Println(errMsg)
			return errMsg
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	req.Header.Add("Content-Type", "application/json")

	resp, httpErr := client.Do(req)
	if httpErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp http post request error: %v", httpErr)
		fmt.Println(errMsg)
		return errMsg
	}

	defer resp.Body.Close()

	rawRes, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error reading response body: %v", readErr)
		fmt.Println(errMsg)
		return errMsg
	}

	unmarErr := json.Unmarshal(rawRes, &resBody)

	if unmarErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error unmarshalling response body: %v", unmarErr)
		fmt.Println(errMsg)
		return errMsg
	}

	return nil
}

func (s *BootcampHandler) getAuthToken() string {
	if s.AuthToken == constants.EmptyString || s.isTokenExpired() {
		s.refreshToken()
	}

	return s.AuthToken
}

func (s *BootcampHandler) isTokenExpired() bool {
	return time.Now().After(s.AuthTokenExp)
}

func (s *BootcampHandler) refreshToken() {
	var response types.LoginResponse

	creds, fileErr := os.ReadFile(s.cnf.Credentials.BootcampAPI)
	if fileErr != nil {
		log.Fatalf("biz.BootcampHandler.refreshToken error reading credentials file: %v", fileErr)
	}

	s.postRequest(s.cnf.EndPoints.PostLogin, creds, &response, false)

	if response.Success && response.Token != constants.EmptyString {
		s.AuthToken = response.Token
		s.AuthTokenExp = time.Now().Add(time.Hour * 12)

		fmt.Println("biz.BootcampHandler.refreshToken Successfully updated auth token")
	}
}
