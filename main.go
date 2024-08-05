package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
)

type Backend struct {
	URL    *url.URL
	Weight int
}

type Proxy struct {
	Blue  Backend
	Green Backend
	mu    sync.RWMutex
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// Use blue as default, fallback to green if blue weight is 0
	selectedBackend := p.Blue
	if p.Blue.Weight == 0 {
		selectedBackend = p.Green
	} else if p.Green.Weight > 0 {
		// If both have weight, randomly choose based on weight
		if (p.Blue.Weight + p.Green.Weight) > 0 {
			if rand.Intn(p.Blue.Weight+p.Green.Weight) >= p.Blue.Weight {
				selectedBackend = p.Green
			}
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(selectedBackend.URL)
	r.Host = selectedBackend.URL.Host
	log.Printf("Proxying request to %s", selectedBackend.URL)
	proxy.ServeHTTP(w, r)
}

func (p *Proxy) UpdateWeights(blueWeight, greenWeight int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Blue.Weight = blueWeight
	p.Green.Weight = greenWeight
}

func main() {
	var rootCmd = &cobra.Command{Use: "proxy"}

	var blueGreenCmd = &cobra.Command{
		Use:   "blue-green",
		Short: "Run a blue-green proxy",
		Run:   runBlueGreenProxy,
	}

	blueGreenCmd.Flags().String("blue", "http://blue-server:8080", "Blue server URL")
	blueGreenCmd.Flags().String("green", "http://green-server:8080", "Green server URL")
	blueGreenCmd.Flags().Int("blue-weight", 100, "Blue server weight (0-100)")
	blueGreenCmd.Flags().Int("green-weight", 0, "Green server weight (0-100)")
	blueGreenCmd.Flags().StringP("listen", "l", ":8080", "Listen address")

	rootCmd.AddCommand(blueGreenCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runBlueGreenProxy(cmd *cobra.Command, args []string) {
	blueServer, _ := cmd.Flags().GetString("blue")
	greenServer, _ := cmd.Flags().GetString("green")
	blueWeight, _ := cmd.Flags().GetInt("blue-weight")
	greenWeight, _ := cmd.Flags().GetInt("green-weight")
	listen, _ := cmd.Flags().GetString("listen")

	blueURL, err := url.Parse(blueServer)
	if err != nil {
		log.Fatalf("Invalid blue server URL: %v", err)
	}

	greenURL, err := url.Parse(greenServer)
	if err != nil {
		log.Fatalf("Invalid green server URL: %v", err)
	}

	proxy := &Proxy{
		Blue:  Backend{URL: blueURL, Weight: blueWeight},
		Green: Backend{URL: greenURL, Weight: greenWeight},
	}

	server := &http.Server{
		Addr:    listen,
		Handler: proxy,
	}

	log.Printf("Starting blue-green proxy server on %s", listen)
	log.Printf("Blue server: %s (weight: %d)", blueServer, blueWeight)
	log.Printf("Green server: %s (weight: %d)", greenServer, greenWeight)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
