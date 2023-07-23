package alert

import (
	"encoding/json"
	netUrl "net/url"

	net_url "net/url"

	log "github.com/sirupsen/logrus"
)

// CallSquadcast will send a Squadcast http request
func (r *Repository) CallSquadcast(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)

	if err != nil {
		log.WithError(err).Errorf("[SQUADCAST] Error parsing URL: %s", url)
		return err
	}

	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(u.String(), body, r.version)
		if err != nil {
			log.WithError(err).Error("[SQUADCAST] Error sending request to " + u.String())
			return
		}
		log.Debugf("[SQUADCAST] Result: %s", result)
	})
}

// CallTelegram will send a Telegram bot http request
func (r *Repository) CallTelegram(url string, body []byte) error {
	u, err := net_url.ParseRequestURI(url)

	if err != nil {
		log.WithError(err).Errorf("[SQUADCAST] Error parsing URL: %s", url)
		return err
	}

	return r.pool.Submit(func() {
		result, err := sendPOSTRequest(u.String(), body, r.version)
		if err != nil {
			log.WithError(err).Error("[TELEGRAM] Error sending request to " + u.String())
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
				parsedURL, err := netUrl.Parse(u.String())
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
