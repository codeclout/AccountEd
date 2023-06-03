package protocol

import (
  "fmt"
  "os"
  "os/signal"
  "strconv"
  "strings"
  "sync"
  "syscall"
  "time"

  "github.com/gofiber/fiber/v2"
  "golang.org/x/exp/slog"
)

func getPort() string {
  p, ok := os.LookupEnv("PORT")
  n, _ := strconv.Atoi(p)

  if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
    return fmt.Sprintf(":%d", n)
  }

  return ":8088"
}

func isProd() bool {
  if env, ok := os.LookupEnv("ENVIRONMENT"); ok && strings.TrimSpace(env) == "prod" {
    return true
  }

  return false
}

type Adapter struct {
  HTTP                 *fiber.App
  WaitGroup            *sync.WaitGroup
  applicationName      string
  isApplicationGetOnly bool
  log                  *slog.Logger
  routePrefix          string
}

func NewAdapter(applicationName, routePrefix string, isAppGetOnly bool, log *slog.Logger) *Adapter {
  api := fiber.New()

  return &Adapter{
    HTTP:                 api,
    applicationName:      applicationName,
    isApplicationGetOnly: isAppGetOnly,
    log:                  log,
    routePrefix:          routePrefix,
  }
}

func (a *Adapter) Initialize(api []*fiber.App) *fiber.App {
  isAppGetOnly := func() int { // FixMe - write buffer size
    if a.isApplicationGetOnly {
      return 0
    }
    return 512
  }

  app := fiber.New(fiber.Config{
    AppName:                      a.applicationName,
    DisablePreParseMultipartForm: true,
    DisableStartupMessage:        isProd(),
    IdleTimeout:                  2 * time.Second,
    ReadTimeout:                  2 * time.Second,
    WriteBufferSize:              isAppGetOnly(),
  })

  for _, x := range api {
    a.log.Info("creating API routes")
    app.Mount(a.routePrefix, x)
  }

  //a.log.Info("starting server")
  //log.Fatal(app.Listen(getPort()))

  return app
}

func (a *Adapter) PostInit(app *fiber.App) {
  s := make(chan os.Signal, 1)
  signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

  <-s
  a.log.Warn("initializing shutdown")
  a.StopProtocolListener(app)
  os.Exit(0)
}

func (a *Adapter) StopProtocolListener(app *fiber.App) {
  a.WaitGroup.Wait()
}
