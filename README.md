# Tracer


Experimental logger with metrics and plans for distributed tracing support.


### Writer

```
originalLogger := logger.New("my-app")
writerTracer := writer.New(outputFile)

tracingLogger := tracer.NewLogger(originalLogger, writerTracer)
Hello(logger)

...

func Hello(logger lager.Logger) {
  logger := logger.Session("hello")
  logger.Debug("start")
  defer logger.Debug("end")

  ...
}

# my-app.hello: 12321
```

### Metron

```
originalLogger := logger.New("my-app")
metronTracer := metron.New()

tracingLogger := tracer.NewLogger(originalLogger, metronTracer)
Hello(logger)

...

func Hello(logger lager.Logger) {
  logger := logger.Session("hello")
  logger.Debug("start")
  defer logger.Debug("end")

  ...
}
```
