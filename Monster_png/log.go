import "github.com/bingoohuang/golog"

func main() {
	golog.Setup()

	log.Printf("Hello, this message is logged by std log, #%d", 1) // default Info
	log.Printf("T! Hello, this message is logged by std log, #%d", 2) // Trace
	log.Printf("D! Hello, this message is logged by std log, #%d", 3) // Debug
	log.Printf("I! Hello, this message is logged by std log, #%d", 4) // Info
	log.Printf("W! Hello, this message is logged by std log, #%d", 5) // Warn
	log.Printf("F! Hello, this message is logged by std log, #%d", 6) // Fatal

	logrus.Tracef("Hello, this message is logged by std log, #%d", 7)
	logrus.Debugf("Hello, this message is logged by std log, #%d", 8)
	logrus.Infof("Hello, this message is logged by std log, #%d", 9)
	logrus.Warnf("Hello, this message is logged by std log, #%d", 10)
}
