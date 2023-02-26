package main

import "flag"

var (
	dev           = flag.Bool("dev", false, "run in localhost development mode")
	addr          = flag.String("a", ":443", "server address")
	configPath    = flag.String("c", "", "config file path")
	certDir       = flag.String("certdir", tsweb.DefaultCertDir("derper-certs"), "directory to store LetsEncrypt certs, if addr's port is :443")
	hostname      = flag.String("hostname", "127.0.0.1", "LetsEncrypt hostname, if addr's port is :443")
	mbps          = flag.Int("mbps", 5, "Mbps (megibits/s) per-client rate limit; 0 means unlimited")
	logCollection = flag.String("logcollection", "", "If non-empty, logtail collection to log to")
	runSTUN       = flag.Bool("stun", false, "also run a STUN server")
)
