package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/smtp"
	"net/url"
	"os"
)

const (
	defaultSendersEmail = "sender-email@gmail.com"
	defaultSendersPss   = "sender-email-password"
	defaultReceiver     = "receiver-example@gmail.com"
)

type Config struct {
	SendersEmail string `json:SendersEmail`
	SendersPss   string `json:SendersPss`
	Receiver     string `json:Receiver`
}

func initConf(c *Config) { // initialize Config struct
	c.SendersEmail = defaultSendersEmail
	c.SendersPss = defaultSendersPss
	c.Receiver = defaultReceiver
}

func (c *Config) createDefaultConfig(path string) {

	// default configurations
	defaultConf := &Config{defaultSendersEmail, defaultSendersPss, defaultReceiver}

	jsonConfig, err := json.Marshal(defaultConf)

	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(path, jsonConfig, 0644) // read write permission

	if err != nil {
		log.Panic(err)
	}

}

func LoadSmtpConfigurations(filename string) Config {

	var conf *Config
	configFile, err := os.Open(filename)

	if err != nil {
		log.Println("Creating new config.json file") // could not open config.js
		conf.createDefaultConfig("config.json")      // create default config.js
	}
	defer configFile.Close()

	byteData, _ := ioutil.ReadAll(configFile)

	// read config file to config object
	json.Unmarshal(byteData, &conf)

	return *conf
}

func (c *Config) HandleSendMail(params url.Values) {

	if c == nil {
		initConf(c)                          // initialize config object
		c.createDefaultConfig("config.json") // initialize config.json
	}

	name := params["name"][0]
	email := params["email"][0]
	subject := params["subject"][0]
	message := params["message"][0]

	// convert to html string message
	htmlMsg := "From: " + c.SendersEmail + "\n" +
		"Name: " + name + "\n" +
		"Users Email: " + email + "\n" +
		"To: " + c.Receiver + "\n" +
		"Subject: " + subject + "\n" +
		message

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", c.SendersEmail, c.SendersPss, "smtp.gmail.com"), c.SendersEmail, []string{c.Receiver}, []byte(htmlMsg))

	if err != nil {
		log.Printf("Error at sending email via net/smtp library.\n SMTP ERROR: %s", err)
		log.Print("Please check your config.json file.")
	} else {
		// message successfully was sent
		log.Print("Sent email...")
	}

}
