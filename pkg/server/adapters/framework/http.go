package framework

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/gofiber/fiber/v2"
  "golang.org/x/exp/slog"
)

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

  isAppGetOnly := func() int { // FixMe - offer fix for type
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

func getPort() string {
  p, ok := os.LookupEnv("PORT")
  n, _ := strconv.Atoi(p)

  if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
    return fmt.Sprintf(":%d", n)
  }

  return ":8088"
}
