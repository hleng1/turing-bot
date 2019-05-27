package main

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

var (
	token    string
	logPath  string
	dbSource string
)

func init() {
	token = "Bot NTgxOTkxNjkwODA2NjI0MjY3.XOsnkQ.CBvABPtPgErY7aEcdDbH8xSaUuE"
	logPath = "turing.log"
	dbSource = "./test.db"
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	user, content := m.Author, m.Content
	if content == "!test" {
		// Handle !test command
		reply, err := s.ChannelMessageSend(m.ChannelID, "Test received!")
		if err != nil {
			log.Panic(err)
		}
		log.Printf("test @%v: %v", user, content)
		log.Println("test reply:", reply.Content)
	} else if matched, err := regexp.MatchString(`^!solved [A-Z]+[0-9]+$`, content); matched && err == nil {
		// Handle !solved XXX command
		reply, err := s.ChannelMessageSend(m.ChannelID, "gz!")
		if err != nil {
			log.Panic(err)
		}
		log.Println("solve received:", content)
		slv := strings.SplitN(content, " ", 2)
		log.Println("solve parsed:", "0:", slv[0], "1:", slv[1])
		log.Printf("solve @%v: %v", user, content)
		log.Println("solve reply:", reply.Content)
	}
}

func main() {
	// Set log output to file
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

	// Database initialization
	db, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		log.Fatal(err)
	}

	stmt, _ := db.Prepare("create table if not exists user (uid integer auto_increment primary key, fname text, lname text);")
	stmt.Exec()

	// Create a new Discordgo session
	dg, err := discordgo.New(token)
	if err != nil {
		log.Fatal(err)
	}

	// Add command handler
	dg.AddHandler(handleCommand)

	// Update status on event Ready
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		s.UpdateStatus(0, "with Turing Machines")
		log.Println("update status")
	})

	// Create a new connection
	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Turing is running ...")

	// Wait until termination signal received
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()
	<-done

	dg.Close()

}
