package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
  envFilePath = "/.dbkenv"
)

var (
  slogger *zap.SugaredLogger

	wsHandler          http.Handler

  // Env vars
	runCmd             string
	runArgs            string
	parsedRunArgs      []string
	workdir            string
	entrypoint         string
	entrypointFullPath string

  depsCmd                 string
  depsInstallArgs         string
  parsedDepsInstallArgs   []string
  depsUninstallArgs       string
  parsedDepsUninstallArgs []string
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	slogger.Debug("Client connected")
	// TODO: Separate new connection?
	wsHandler.ServeHTTP(w, r)
}

func errLogUndefinedEnvVar(name, value string) {
  if value == "" {
    slogger.Error(
      "The Devbook env var '%s' is empty. Make sure to add the %s var to %s",
      name,
      name,
      envFilePath,
    )
  }
}

func loadDBKEnvs() {
  slogger.Infow("Loading envs from the .dbkenv file", "envFilePath", envFilePath)

  file, err := os.Open(envFilePath)
  if err != nil {
    slogger.Errorw("Failed to open dbkenv file",
      "envFilePath", envFilePath,
      "error", err,
    )
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  // Optionally, resize scanner's capacity for lines over 64K, see next example.
  for scanner.Scan() {
    // Expects vars in the format "VAR_NAME=VALUE"
    // ["VAR_NAME", "VALUE"]
    envVar := scanner.Text()

    name, value, found := strings.Cut(envVar, "=")

    if !found {
      slogger.Errorw("Invalid env format in the .dbkenv file",
        "line", envVar,
      )
    }

    slogger.Infow("Devbook env var",
      "name", name,
      "value", value,
    )

    switch name {
    case "RUN_CMD":
      errLogUndefinedEnvVar("RUN_CMD", value)
      runCmd = value
    case "RUN_ARGS":
      errLogUndefinedEnvVar("RUN_ARGS", value)
      runArgs = value
      parsedRunArgs = strings.Fields(runArgs)
    case "WORKDIR":
      errLogUndefinedEnvVar("WORKDIR", value)
      workdir = value
    case "ENTRYPOINT":
      errLogUndefinedEnvVar("ENTRYPOINT", value)
      entrypoint = value
    case "DEPS_CMD":
      errLogUndefinedEnvVar("DEPS_CMD", value)
      depsCmd = value
    case "DEPS_INSTALL_ARGS":
      errLogUndefinedEnvVar("DEPS_INSTALL_ARGS", value)
      depsInstallArgs = value
      parsedDepsInstallArgs = strings.Fields(depsInstallArgs)
    case "DEPS_UNINSTALL_ARGS":
      errLogUndefinedEnvVar("DEPS_UNINSTALL_ARGS", value)
      depsUninstallArgs = value
      parsedDepsUninstallArgs = strings.Fields(depsUninstallArgs)
    default:
      slogger.Errorw("Unknown Devbook env var",
        "name", name,
        "value", value,
      )
    }
  }

  if err := scanner.Err(); err != nil {
    slogger.Errorw("Error from scanner for .dbkenv file",
      "error", err,
    )
  }
}

func initLogger() {
  rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/var/log/devbookd.log"],
	  "errorOutputPaths": ["stderr", "/var/log/devbookd.err"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
  slogger = l.Sugar()
}

func main() {
  initLogger()
	defer slogger.Sync()
	slogger.Info("Logger construction succeeded")

  loadDBKEnvs()

	entrypointFullPath = path.Join(workdir, entrypoint)

	router := mux.NewRouter()
	server := rpc.NewServer()
	codeSnippet := NewCodeSnippetService()
	if err := server.RegisterName("codeSnippet", codeSnippet); err != nil {
    slogger.Errorw("Failed to register code snippet service", "error", err)
	}

	wsHandler = server.WebsocketHandler([]string{"*"})
	//router.Handle("/ws", wshandler)
	router.HandleFunc("/ws", serveWs)

  slogger.Info("Starting server on the port :8010")
	if err := http.ListenAndServe(":8010", router); err != nil {
    slogger.Errorw("Failed to start the server", "error", err)
  }
}
