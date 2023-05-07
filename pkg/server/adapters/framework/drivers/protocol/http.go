package protocol

import (
  "fmt"
  "log"
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
  api                  *fiber.App
  applicationName      string
  isApplicationGetOnly bool
  logger               *slog.Logger
  routePrefix          string
  WaitGroup            *sync.WaitGroup
}

func NewAdapter(an, rp string, igo bool, l *slog.Logger) *Adapter {
  api := fiber.New()

  return &Adapter{
    api:                  api,
    applicationName:      an,
    isApplicationGetOnly: igo,
    logger:               l,
    routePrefix:          rp,
  }
}

func (a *Adapter) Initialize(api []*fiber.App) {
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
    ETag:                         true,
    IdleTimeout:                  2 * time.Second,
    ReadTimeout:                  2 * time.Second,
    WriteBufferSize:              isAppGetOnly(),
  })

  for _, x := range api {
    a.logger.Info("creating API routes")
    app.Mount(a.routePrefix, x)
  }

  a.logger.Info("starting server")
  log.Fatal(app.Listen(getPort()))
}

func (a *Adapter) PostInit(app *fiber.App) {
  s := make(chan os.Signal, 1)
  signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

  <-s
  a.StopProtocolListener(app)
  os.Exit(0)
}

func (a *Adapter) StopProtocolListener(app *fiber.App) {
  a.WaitGroup.Wait()
}
