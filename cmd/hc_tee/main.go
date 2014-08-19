package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/andybons/hipchat"
	"github.com/codegangsta/cli"
)

type clientStruct struct {
	HipChat hipchat.Client
	Room    string
	Name    string
	Color   string
	Notify  bool
	ReadCh  chan string
	QuitCh  chan struct{}
}

func (s *clientStruct) PostMessage(msg string) {
	req := hipchat.MessageRequest{
		From:          s.Name,
		Message:       msg,
		Color:         hipchat.Color(s.Color),
		MessageFormat: hipchat.FormatText,
		Notify:        s.Notify,
	}
	room := s.HipChat.Room(s.Room)
	room.PostMessage(req)
	fmt.Print(msg)
}

func (s *clientStruct) outputToHipChat() {
	buffer := make([]byte, 0, 1024)
	ch3sec := time.After(1 * time.Second)
	for {
		select {
		case message := <-s.ReadCh:
			buffer = append(buffer, []byte(message)...)
		case <-ch3sec:
			s.PostMessage(string(buffer))
			buffer = make([]byte, 0, 1024)
			ch3sec = time.After(3 * time.Second)
		case <-s.QuitCh:
			s.PostMessage(string(buffer))
			s.QuitCh <- struct{}{}
			return
		}
	}
}

func application(c *cli.Context) {
	client := &clientStruct{
		HipChat: hipchat.Client{AuthToken: c.String("auth")},
		Room:    c.String("room"),
		Name:    c.String("from"),
		Notify:  c.Bool("notify"),
		Color:   strings.ToLower(c.String("color")),
		ReadCh:  make(chan string),
		QuitCh:  make(chan struct{}),
	}
	go client.outputToHipChat()
	file := os.Stdin
	bufferedReader := bufio.NewReader(file)
	line, err := bufferedReader.ReadString('\n')
	// Did we reach the end of file?
	for {

		if err != nil && err != io.EOF {
			// report an error.
			fmt.Fprintf(os.Stderr, "error %q occured.", err)
		}
		client.ReadCh <- line
		if err == io.EOF {
			client.QuitCh <- struct{}{}
			<-client.QuitCh
			return
		}

		line, err = bufferedReader.ReadString('\n')

	}

}

var version string

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "hc-tee"
	//  The tee utility copies standard input to standard output, making a copy in zero or more files.  The output is unbuffered.
	app.Usage = "The hc-tee utility copies standard input to the room in hipchat. The output is buffered. The system buffers 3 seconds of output before sending the message to hipchat. This is to ensure the rate of messages going to hipchat is within the 100 messages per 5 minutes."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "auth,token,t",
			Usage:  "The auth token to use for Hipchat.",
			EnvVar: "HIPCHAT_AUTH",
		},
		cli.StringFlag{
			Name:   "room,r",
			Usage:  "The room id to which to post the messages to.",
			EnvVar: "HIPCHAT_ROOM",
		},
		cli.StringFlag{
			Name:   "from,f",
			Usage:  "The from name, defaults to hc-tee",
			EnvVar: "HIPCHAT_FROM",
			Value:  "hc-tee",
		},
		cli.StringFlag{
			Name:   "color,c",
			Usage:  "The color to use. Defaults to 'green'. Valid colors are: ['yello','red','green','purple','gray','random'] ",
			EnvVar: "HIPCHAT_FROM",
			Value:  hipchat.ColorGreen,
		},
		cli.BoolFlag{
			Name:   "notify,n",
			Usage:  "Notify the room",
			EnvVar: "HIPCHAT_NOTIFY",
		},
	}
	app.Action = application
	app.Run(os.Args)
}
