package external

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *apiCaller) caller(_ context.Context, request, receiver interface{}, header map[string]string, method string, uri string, timeout *time.Duration) (int, error) {
	var err error

	agent := fiber.AcquireAgent()
	defer agent.ConnectionClose()

	// set method
	agent.Request().Header.SetMethod(method)

	// set uri
	agent.Request().SetRequestURI(uri)

	// set headers
	agent.Request().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	if len(header) > 0 {
		for key, value := range header {
			agent.Request().Header.Set(key, value)
		}
	}

	// set request body
	if request != nil {
		agent.JSON(request)
	}

	res := fiber.AcquireResponse()
	defer fiber.ReleaseResponse(res)
	agent.SetResponse(res)

	// prepare HTTP request
	if err = agent.Parse(); err != nil {
		return -1, err
	}

	// set time out
	timeoutDuration := 30 * time.Second // default
	if timeout != nil {
		timeoutDuration = *timeout
	}
	agent.Timeout(timeoutDuration)

	// skip verify cert
	agent.InsecureSkipVerify()
	agent.ReadBufferSize = (16 * 1024) // limit buffer size

	// execute HTTP request
	httpCode, body, errs := agent.Bytes()
	if errs != nil {
		errString := ""
		for _, errorItem := range errs {
			errString = errString + errorItem.Error()
		}

		return -1, fmt.Errorf("%v", errString)
	}

	// set api response return
	if strings.TrimSpace(string(res.Header.ContentType())) == "" || strings.Contains(string(res.Header.ContentType()), "application/json") {
		// Default to JSON if no Content-Type is set
		if err = json.Unmarshal(body, &receiver); err != nil {
			return -1, err
		}
	}

	return httpCode, nil
}
