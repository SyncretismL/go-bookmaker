package main

import (
	"bookmaker/internal/sport"
	"bookmaker/internal/subscribe"
	"bookmaker/pkg/logger"
	"io"
	"math"
	"time"
)

type Subscribe struct {
	logger logger.Logger
	sports sport.Sports
}

func newSubscribe(newLogger logger.Logger, sports sport.Sports) *Subscribe {
	return &Subscribe{
		logger: newLogger,
		sports: sports,
	}
}

func (s *Subscribe) SubscribeOnSportsLines(srv subscribe.SubscrService_SubscribeOnSportsLinesServer) error {
	working := false
	ctx := srv.Context()
	new := make(chan struct{})
	dataReq := make(chan *subscribe.SubscrRequest)
	errch := make(chan error)

	go func() {
		for {
			req, err := srv.Recv()
			if err == io.EOF {
				return
			}

			if err != nil {
				continue
			}

			if working {
				new <- struct{}{}
			}

			dataReq <- req

		}
	}()

	go func() {
		var prevData []*sport.Sport
		var respData []*subscribe.Result

		for {
			req := <-dataReq

			dur, err := time.ParseDuration(req.Repeat)
			if err != nil {
				errch <- err
			}

			working = true

		loop:
			for {
				select {
				case <-new:
					prevData = nil
					break loop
				default:
					data, err := s.sports.Find(req.Sport)
					if err != nil {
						errch <- err
					}

					if prevData != nil {
						for i, v := range data {
							result := &subscribe.Result{
								Sport: v.Sport,
								Delta: math.Round((v.Coefficient-prevData[i].Coefficient)*100) / 100,
							}

							respData = append(respData, result)
						}
					} else {
						for _, v := range data {
							result := &subscribe.Result{
								Sport: v.Sport,
								Delta: math.Round(v.Coefficient*100) / 100,
							}

							respData = append(respData, result)
						}
					}

					resp := subscribe.SubscrResponse{Results: respData}
					if err := srv.Send(&resp); err != nil {
						errch <- err
					}

					prevData = data
					respData = nil

					time.Sleep(dur)
				}
			}

		}
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		return nil
	}
}
