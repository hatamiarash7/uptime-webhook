# Uptime webhook handler

It's webhook handler for [uptime.com](https://uptime.com).

## Installation

1. Create your configuration file. You can use [`config.yml.example`](./configs/config.yml.example) as a template.
2. Add your Squadcast configuration to the `config.yml` file.
3. Run `docker-compose up -d` to start the service or run it manually:

   ```bash
   docker run -d \
       --name uptime-webhook \
       -p 8080:8080 \
       -v /path/to/config.yml:/config/config.yml \
       r1cloud/uptime-webhook:latest
   ```

4. Go to Uptime dashboard and create a new webhook in *Notifications / Integrations* section.
5. Use your deployed webhook URL like this:

   ```text
   http://<your-host>:8080/api/v1/alert/
   ```

6. Wait for alerts 😁 Or you can send a test alert in Uptime dashboard.

## Notifiers

You can use multiple notifiers at the same time. Just define them to the `config.yml` file.

### Squadcast

1. Create a new Service in Squadcast
2. Use `Incident Webhook` as an alert source
3. Define that in `config.yml` file

   ```yaml
   notifier:
       squadcast:
           enable: true
           teams:
               team1: "<WEBHOOK-URL>"
   ```

## Contributing

Don't be shy and reach out to us if you want to contribute. 😉

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## Issues

Each project may have many problems. Contributing to the better development of this project by reporting them. 👍
