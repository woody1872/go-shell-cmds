package main

import "log"

func main() {
	// var ip string
	// flag.StringVar(&ip, "i", "8.8.8.8", "ip address")
	// flag.Parse()

	// cmd := exec.Command("ping", "-c 2", "-i 1", ip)

	// var buf bytes.Buffer
	// cmd.Stdout = &buf

	// if err := cmd.Start(); err != nil {
	// 	log.Fatalln(err)
	// }

	// done := make(chan error)
	// go func() { done <- cmd.Wait() }()

	// timeout := time.After(5 * time.Second)

	// select {
	// case <-timeout:
	// 	cmd.Process.Kill()
	// 	log.Fatalln("ERROR: command timed out")
	// case err := <-done:
	// 	if err != nil {
	// 		log.Fatalln("ERROR:", err)
	// 	}
	// }

	for i := 1; i < 101; i++ {
		log.SetFlags(i)
		log.Println("Flag", i)
	}
}
