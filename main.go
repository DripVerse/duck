package main

import (
	"os"
	"time"

	"github.com/dripverse/duck/src"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const figlet = `

W E L C O M E  T O
    __    
___( o)>  
\ <_. )   
 '---'    
`

func init() {
	// Log as Text with color
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	// Log to stdout
	logrus.SetOutput(os.Stdout)

	VerifyConfig()
}

func main() {
	logrus.Infoln("🚀")
	// // Define input flags
	// username := flag.String("user", "", "username to use in the chatroom.")
	// chatroom := flag.String("room", "", "chatroom to join.")
	// loglevel := flag.String("log", "", "level of logs to print.")
	// discovery := flag.String("discover", "", "method to use for discovery.")
	// // Parse input flags
	// flag.Parse()

	// // Set the log level
	// switch *loglevel {
	// case "panic", "PANIC":
	// 	logrus.SetLevel(logrus.PanicLevel)
	// case "fatal", "FATAL":
	// 	logrus.SetLevel(logrus.FatalLevel)
	// case "error", "ERROR":
	// 	logrus.SetLevel(logrus.ErrorLevel)
	// case "warn", "WARN":
	// 	logrus.SetLevel(logrus.WarnLevel)
	// case "info", "INFO":
	// 	logrus.SetLevel(logrus.InfoLevel)
	// case "debug", "DEBUG":
	// 	logrus.SetLevel(logrus.DebugLevel)
	// case "trace", "TRACE":
	// 	logrus.SetLevel(logrus.TraceLevel)
	// default:
	// 	logrus.SetLevel(logrus.InfoLevel)
	// }

	// // Display the welcome figlet
	// fmt.Printf(figlet)
	// fmt.Println("Initiating Duck Chat...")
	// fmt.Println("(This may take upto 30 seconds.)")
	// fmt.Println()

	// // Create a new P2PHost
	// p2phost := src.NewP2P()
	// logrus.Infoln("Preparing system...")

	// // Connect to peers with the chosen discovery method
	// switch *discovery {
	// case "announce":
	// 	p2phost.AnnounceConnect()
	// case "advertise":
	// 	p2phost.AdvertiseConnect()
	// default:
	// 	p2phost.AdvertiseConnect()
	// }
	// logrus.Infoln("Finding Anons...")

	// // Join the chat room
	// chatapp, _ := src.JoinChatRoom(p2phost, *username, *chatroom)
	// logrus.Infof("Joining '%s' as '%s'", chatapp.RoomName, chatapp.UserName)

	// // Wait for network setup to complete
	// time.Sleep(time.Second * 5)

	// // Create the Chat UI
	// ui := src.NewUI(chatapp)
	// // Start the UI system
	// ui.Run()
}

func VerifyConfig() {
	viper.SetConfigName(".drip")
	viper.SetConfigType("env")

	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logrus.Infoln("Welcome to DUCK 🦆\nLet's set it up for you 👋")
		} else {
			// Config file was found but another error was produced
			logrus.Errorln("fatal error config file: %w", err)
			panic("Error reading config file. Are you sure it is setup?")
		}
	}

	// EVM Key is required
	if len(viper.GetString("ACCOUNT_KEY_EVM")) == 0 {
		panic("No Wallet Key Found! Kindly update your key before continuing...")
	}
	// Fetch EVM Address from Key
	src.GetAddress(viper.GetString("ACCOUNT_KEY_EVM"))
	logrus.Infoln("✅ Config OK")
}
