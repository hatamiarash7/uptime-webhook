package alert

import (
	"encoding/json"
	"time"

	net_url "net/url"

	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	log "github.com/sirupsen/logrus"
)

// CallSquadcast will send a Squadcast http request
func (r *Repository) CallSquadcast(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)

	if err != nil {
		log.WithError(err).Errorf("[SQUADCAST] Error parsing URL: %s", url)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSquadcastSendFailure)})
		return err
	}

	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(u.String(), body, r.version)

		if err != nil {
			log.WithError(err).Error("[SQUADCAST] Error sending request to " + u.String())
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSquadcastSendFailure)})
			return
		}

		log.Debugf("[SQUADCAST] Result: %s", result)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSquadcastSendSuccess)})
	})
}

// CallTelegram will send a Telegram bot http request
func (r *Repository) CallTelegram(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)

	if err != nil {
		log.WithError(err).Errorf("[SQUADCAST] Error parsing URL: %s", url)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendFailure)})
		return err
	}

	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(u.String(), body, r.version)
		if err != nil {
			log.WithError(err).Error("[TELEGRAM] Error sending request to " + u.String())
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendFailure)})
			return
		}

		log.Debugf("[TELEGRAM] Result: %s", result)

		res := make(map[string]interface{})
		err = json.Unmarshal([]byte(result), &res)
		if err != nil {
			log.WithError(err).Error("[TELEGRAM] Error unmarshalling response")
		}
		if price, ok := res["ok"].(bool); ok {
			if price == false {
				// Extract the query parameters
				parsedURL, err := net_url.Parse(u.String())
				if err != nil {
					log.WithError(err).Error("[TELEGRAM] Error parsing URL")
				}
				queryParams := parsedURL.Query()

				log.WithFields(log.Fields{
					"error_code":  res["error_code"],
					"description": res["description"],
					"chat":        queryParams.Get("chat_id"),
				}).Error("[TELEGRAM] Failed to send message")
				r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendFailure)})
			}
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendSuccess)})
		} else {
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendSuccess)})
		}
	})
}

// CallSlack will send a Slack http request
func (r *Repository) CallSlack(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)

	if err != nil {
		log.WithError(err).Errorf("[SLACK] Error parsing URL: %s", url)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSlackSendFailure)})
		return err
	}

	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(u.String(), body, r.version)

		if err != nil {
			log.WithError(err).Error("[SLACK] Error sending request to " + u.String())
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSlackSendFailure)})
			return
		}

		log.Debugf("[SLACK] Result: %s", result)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSlackSendSuccess)})
	})
}

// CallMattermost will send a Mattermost http request
func (r *Repository) CallMattermost(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)
	if err != nil {
		log.WithError(err).Errorf("[MATTERMOST] Error parsing URL: %s", url)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncMattermostSendFailure)})
		return err
	}

	return r.pool.Submit(func() {
		maxRetries := 3
		retryDelay := 3 * time.Second
		var lastErr error

		for i := range maxRetries {
			result, err := sendPOSTRequest(u.String(), body, r.version)

			if err == nil {
				log.Debugf("[MATTERMOST] Result: %s", result)
				r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncMattermostSendSuccess)})
				return
			}

			lastErr = err
			log.WithError(err).Errorf("[MATTERMOST] Error sending request to %s (attempt %d/%d)", u.String(), i+1, maxRetries)

			if i < maxRetries-1 {
				time.Sleep(retryDelay)
			}
		}

		// All retries failed
		log.WithError(lastErr).Error("[MATTERMOST] All retry attempts failed for " + u.String())
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncMattermostSendFailure)})
	})
}

// CallCustom will send a Custom http request
func (r *Repository) CallCustom(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)

	if err != nil {
		log.WithError(err).Errorf("[CUSTOM] Error parsing URL: %s", url)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncCustomSendFailure)})
		return err
	}

	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(u.String(), body, r.version)

		if err != nil {
			log.WithError(err).Error("[CUSTOM] Error sending request to " + u.String())
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncCustomSendFailure)})
			return
		}

		log.Debugf("[CUSTOM] Result: %s", result)
		r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncCustomSendSuccess)})
	})
}
