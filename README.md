# Yandex Tracker CLI


## Instalation

Install using Go command:

```
go install github.com/fgazat/trc@latest
```

Or download binary for your OS from [Releases](https://github.com/fgazat/trc/releases).


## Configuration

For authorization you have to get Yandex Tracker OAuth token and `X-Org-ID` or `X-Cloud-Org-ID`. For more information, check out the Yandex Cloud documentation: [API access](https://yandex.cloud/en/docs/tracker/concepts/access) .

These secrets should be specified in environment variables: `TRACKER_TOKEN`, `X_CLOUD_ORG_ID`, `X_ORG_ID`.

```bash
export TRACKER_TOKEN="YOUR_OAUTH_TOKEN"
export X_CLOUD_ORG_ID="YOUR_CLOUD_ORG_ID"
export X_ORG_ID="YOUR_ORG_ID"
```

Also you can specify your config file. By default, it tries to find the config file here: `$HOME/.trc/config.yaml`:

```yaml
api_base_url: https://api.tracker.yandex.net
web_base_url: https://tracker.yandex.com
debug: false
issues:
  default_queue: MYQUEUE
  assignee: fgazat
```

You set your own config filepath with `TRC_CFG_PATH` env var:

```bash
export TRC_CFG_PATH="~/.trc/mycustom.yaml"
```

## Usage

```bash
# create issue
trc c -m "hello" -d "world" -q "TEST"
```

