package Frigatier

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (f *Frigatier) getEventImage(eventId string) string {
	frigateIP := f.config.Frigate.Host
	frigatePort := f.config.Frigate.Port
	url := fmt.Sprintf("http://%s:%d/api/events/%s/thumbnail.jpg", frigateIP, frigatePort, eventId)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to grab image for event: %s\n", err)
		return ""
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	file := eventId + ".png"
	err = os.WriteFile(file, body, 0666)
	if err != nil {
		log.Printf("Failed to store image for event: %s", eventId)
	}
	return file
}
