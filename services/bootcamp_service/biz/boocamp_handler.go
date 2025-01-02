package biz

import (
	"bytes"
	"carthage/services/bootcamp_service/biz/interfaces"
	"carthage/services/bootcamp_service/config"
	"carthage/services/bootcamp_service/constants"
	"carthage/services/bootcamp_service/dto"
	"carthage/services/bootcamp_service/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	pbrs "github.com/dipanshuchaubey/protos-package/bootcamp_service/response"
	pbty "github.com/dipanshuchaubey/protos-package/bootcamp_service/types"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type BootcampHandler struct {
	AuthToken    string
	AuthTokenExp time.Time
	cnf          config.Config
	logger       *slog.Logger
	trace        trace.Tracer
}

func NewBootcampHandler(config config.Config) interfaces.BootcampInterface {
	return &BootcampHandler{
		AuthToken:    constants.EmptyString,
		AuthTokenExp: time.Now(),
		cnf:          config,
		logger:       otelslog.NewLogger(constants.ServiceName),
		trace:        otel.Tracer(constants.ServiceName),
	}
}

func (s *BootcampHandler) GetBootcampsDetails(ctx context.Context) ([]*pbrs.GetBootcampsDetailsResponse_Data, error) {
	ctx, span := s.trace.Start(ctx, "biz.GetBootcampsDetails")
	defer span.End()

	bootcampRes, httpErr := http.Get(s.cnf.EndPoints.GetBootcamps)
	if httpErr != nil {
		errMsg := fmt.Errorf("GetBootcampsDetails: error getting bootcamps %v", httpErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	resBody, resErr := io.ReadAll(bootcampRes.Body)
	if resErr != nil {
		errMsg := fmt.Errorf("cannot read response body: %v", resErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	var bootcampInfos types.BootcampResponse
	marErr := json.Unmarshal(resBody, &bootcampInfos)

	if marErr != nil {
		errMsg := fmt.Errorf("GetBootcampsDetails: error unmarshalling json data %v", marErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
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
					s.logger.ErrorContext(ctx, fmt.Sprintf("Recovered from panic in GetBootcampsDetails: %v", err))
					reviewsChan <- bReviews
				}
			}()

			reviewsRes, httpErr := http.Get(fmt.Sprintf("https://bootcamper.dipanshu.work/api/v1/bootcamps/%s/reviews", bootcampInfo.ID))
			if httpErr != nil {
				errMsg := fmt.Errorf("GetBootcampsDetails: GetBootcampsDetails: error getting reviews for BootcampID %v: %v", bootcampInfo.ID, marErr)
				s.logger.ErrorContext(ctx, errMsg.Error())
				return
			}

			resBody, resErr := io.ReadAll(reviewsRes.Body)
			if resErr != nil {
				errMsg := fmt.Errorf("GetBootcampsDetails: error reading response body: %v", resErr)
				s.logger.ErrorContext(ctx, errMsg.Error())
				return
			}

			var reviews types.ReviewResponse
			marErr := json.Unmarshal(resBody, &reviews)
			if marErr != nil {
				errMsg := fmt.Errorf("GetBootcampsDetails: error unmarshalling json data %v", marErr)
				s.logger.ErrorContext(ctx, errMsg.Error())
				return
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
	ctx, span := s.trace.Start(ctx, "biz.CreateBootcamp")
	defer span.End()

	// 1. Validate request body
	if body.Title == constants.EmptyString || body.Email == constants.EmptyString {
		errMsg := fmt.Errorf("biz.CreateBootcamp missing title or email params in request body")
		s.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	// 2. Create Bootcamp
	httpBody, marErr := json.Marshal(body)
	if marErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error marshalling request body: %v", marErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	var bootcamp types.CreateBootcampResponse

	s.logger.InfoContext(ctx, fmt.Sprintf("biz.CreateBootcamp creating bootcamp with title: %s", body.Title))
	if postErr := s.postRequest(ctx, s.cnf.EndPoints.PostBootcamp, httpBody, &bootcamp, true); postErr != nil {
		return nil, postErr
	}

	if !bootcamp.Success {
		errMsg := fmt.Errorf("biz.CreateBootcamp error creating bootcamp: %v", bootcamp.Error)
		s.logger.ErrorContext(ctx, errMsg.Error())
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

func (s *BootcampHandler) postRequest(ctx context.Context, url string, httpBody []byte, resBody any, auth bool) error {
	ctx, span := s.trace.Start(ctx, "biz.postRequest")
	defer span.End()

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(httpBody))
	if err != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error creating POST request: %v", err)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return errMsg
	}

	if auth {
		token := s.getAuthToken(ctx)
		if token == constants.EmptyString {
			errMsg := fmt.Errorf("biz.CreateBootcamp error getting auth token")
			s.logger.ErrorContext(ctx, errMsg.Error())
			return errMsg
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	req.Header.Add("Content-Type", "application/json")

	resp, httpErr := client.Do(req)
	if httpErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp http post request error: %v", httpErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return errMsg
	}

	defer resp.Body.Close()

	rawRes, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error reading response body: %v", readErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return errMsg
	}

	unmarErr := json.Unmarshal(rawRes, &resBody)

	if unmarErr != nil {
		errMsg := fmt.Errorf("biz.CreateBootcamp error unmarshalling response body: %v", unmarErr)
		s.logger.ErrorContext(ctx, errMsg.Error())
		return errMsg
	}

	return nil
}

func (s *BootcampHandler) getAuthToken(ctx context.Context) string {
	if s.AuthToken == constants.EmptyString || s.isTokenExpired() {
		s.refreshToken(ctx)
	}

	return s.AuthToken
}

func (s *BootcampHandler) isTokenExpired() bool {
	return time.Now().After(s.AuthTokenExp)
}

func (s *BootcampHandler) refreshToken(ctx context.Context) {
	ctx, span := s.trace.Start(ctx, "biz.refreshToken")
	defer span.End()

	var response types.LoginResponse

	creds, fileErr := os.ReadFile(s.cnf.Credentials.BootcampAPI)
	if fileErr != nil {
		s.logger.ErrorContext(ctx, fmt.Sprintf("biz.BootcampHandler.refreshToken error reading credentials file: %v", fileErr))
		return
	}

	if err := s.postRequest(ctx, s.cnf.EndPoints.PostLogin, creds, &response, false); err != nil {
		return
	}

	if response.Success && response.Token != constants.EmptyString {
		s.AuthToken = response.Token
		s.AuthTokenExp = time.Now().Add(time.Hour * 12)

		s.logger.InfoContext(ctx, "biz.BootcampHandler.refreshToken Successfully updated auth token")
		return
	}
}
