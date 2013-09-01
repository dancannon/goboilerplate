package app

import (
	"github.com/dancannon/gonews/util"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	AppName string

	BasePath   string
	ConfigPath string
	ViewsPath  string
	AssetsPath string
	LogsDir    string

	RunMode string = "dev"
	DevMode bool   = true

	TRACE = log.New(ioutil.Discard, "TRACE ", log.Ldate|log.Ltime|log.Lshortfile)
	INFO  = log.New(ioutil.Discard, "INFO  ", log.Ldate|log.Ltime|log.Lshortfile)
	WARN  = log.New(ioutil.Discard, "WARN  ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	Router *mux.Router

	Initialized bool = false
)

func init() {
	log.SetFlags(INFO.Flags())
}

func Init(mode string) {
	RunMode = mode
	DevMode = mode == "dev"

	BasePath, _ = os.Getwd()
	ConfigPath = path.Join(BasePath, "app")
	ViewsPath = path.Join(BasePath, "views")
	AssetsPath = path.Join(BasePath, "assets")
	LogsDir = path.Join(BasePath, "logs")

	// Configure logging.
	util.InitLogging(LogsDir)
	TRACE = util.NewLogger("trace", "trace.log")
	INFO = util.NewLogger("info", "trace.log")
	WARN = util.NewLogger("warn", "warn.log")
	ERROR = util.NewLogger("error", "error.log")

	Initialized = true
}

func SetRouter(r *mux.Router) {
	Router = r
}
