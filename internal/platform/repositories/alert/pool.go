package alert

import (
	"encoding/json"
	netUrl "net/url"

	log "github.com/sirupsen/logrus"
)

// CallSquadcast will send a Squadcast http request
func (r *Repository) CallSquadcast(url string, body []byte) error {
	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(url, body)
		if err != nil {
			log.WithError(err).Error("[SQUADCAST] Error sending request to " + url)
			return
		}
		log.Debugf("[SQUADCAST] Result: %s", result)
	})
}

// CallTelegram will send a Telegram bot http request
func (r *Repository) CallTelegram(url string, body []byte) error {
	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(url, body)
		if err != nil {
			log.WithError(err).Error("[TELEGRAM] Error sending request to " + url)
			return
		}

		log.Debugf("[TELEGRAM] Result: %s", result)

		r := make(map[string]interface{})
		err = json.Unmarshal([]byte(result), &r)
		if err != nil {
			log.WithError(err).Error("[TELEGRAM] Error unmarshalling response")
		}
		if price, ok := r["ok"].(bool); ok {
			if price == false {
				// Extract the query parameters
				parsedURL, err := netUrl.Parse(url)
				if err != nil {
					log.WithError(err).Error("[TELEGRAM] Error parsing URL")
				}
				queryParams := parsedURL.Query()

				log.WithFields(log.Fields{
					"error_code":  r["error_code"],
					"description": r["description"],
					"chat":        queryParams.Get("chat_id"),
				}).Error("[TELEGRAM] Failed to send message")
			}
		}
	})
}
