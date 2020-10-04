package model

const (
	DEFAULT_API_URL    = "http://127.0.0.1:9090"
	DEFAULT_API_SECRET = ""

	API_PATH_CONFIGS         = "/configs"
	API_PATH_LOGS            = "/logs"
	API_PATH_PROXIES         = "/proxies"
	API_PATH_PROXIES_LATENCY = "/delay"
	API_PATH_TRAFFIC         = "/traffic"

	API_LATENCY_TEST_URL     = "https://gstatic.com/generate_204"
	API_LATENCY_TEST_TIMEOUT = "5000"

	EXIT_BY_CTRL_C = "Press Ctrl + C to exit"

	REQUEST_LATENCY_TEST_ERROR_CODE   = "503"
	REQUEST_LATENCY_TEST_ERROR_MSG    = "Error"
	REQUEST_LATENCY_TEST_TIMEOUT_CODE = "504"
	REQUEST_LATENCY_TEST_TIMEOUT_MSG  = "Timeout"

	DB_FILE_NAME         = "clash_cli.db"
	DB_BUCKET_NAME       = "CLASH_CLI"
	DB_KEY_URL           = "CLASH_API_URL"
	DB_KEY_SECRET        = "CLASH_API_SECRET"
	DB_ERROR_KEYNOTFOUND = "not found"

	WARNING_CANNOT_CONN_CLASH = "Cannot connect to Clash，please check clash configuration."
	WARNING_NOT_PROXY         = "There are currently no proxies to select."
	WARNING_UNKNOWN_URL_TYPE  = "Unknown url type."

	CLASH_CONF_MODE_DIRECT = "direct"

	PROMPT_ROOT_LABEL        = "Function Select"
	PROMPT_ROOT_ITEM_TYPE    = "Proxy Type"
	PROMPT_ROOT_ITEM_PROXY   = "Select Proxy"
	PROMPT_ROOT_ITEM_TRAFFIC = "Realtime Traffic"
	PROMPT_ROOT_ITEM_LOG     = "Proxy Log"
	PROMPT_ROOT_ITEM_CLICONF = "CLI config"
	PROMPT_ROOT_ITEM_EXIT    = "Exit"

	PROMPT_MODEL_LABEL = "Select Proxy Mode"

	PROMPT_PROXY_LABEL             = "Proxy Mode"
	PROMPT_PROXY_ITEM_ALL          = "GLOBAL"
	PROMPT_PROXY_ITEM_LATENCY_TEST = "Latency Test"

	PROMPT_PROXY_GLOBAL_GROUPNAME = "GLOBAL"
	PROMPT_PROXY_GLOBAL_LABEL     = "Select Global Proxy"

	PROMPT_CONFIG_LABEL       = "Update Clash-CLI config"
	PROMPT_CONFIG_ITEM_URL    = "Clash API Url"
	PROMPT_CONFIG_ITEM_SECRET = "Clash API Secret"
)
