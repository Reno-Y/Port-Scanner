package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type ScanResult struct {
	Port    int
	Open    bool
	Service string
}

func scanPort(host string, port int, wg *sync.WaitGroup, results chan<- ScanResult, verbose bool) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err == nil {
		service := portSignification(port)
		results <- ScanResult{Port: port, Open: true, Service: service}
		if verbose {
			if service != "Inconnu" {
				fmt.Printf("✓ Port %d est ouvert et correspond à %s\n", port, service)
			} else {
				fmt.Printf("✓ Port %d est ouvert\n", port)
			}
		}
		conn.Close()
	} else if verbose {
		// En mode verbose, afficher aussi les ports fermés
		results <- ScanResult{Port: port, Open: false, Service: ""}
	}
}

func askUserHost() string {
	var host string

	fmt.Print("Entrez l'adresse IP ou le nom d'hôte à scanner: ")
	fmt.Scanln(&host)

	return host
}

func askUserStartPort() int {
	var startPort int

	fmt.Print("Entrez le port de départ (1-2048): ")
	fmt.Scanln(&startPort)

	if startPort < 1 || startPort > 65534 {
		fmt.Println("Port de départ invalide. Veuillez entrer un nombre entre 1 et 65535.")
		return askUserStartPort()
	}

	return startPort
}

func askUserEndPort() int {
	var endPort int

	fmt.Print("Entrez le port de fin (1-2048): ")
	fmt.Scanln(&endPort)

	if endPort < 1 || endPort > 65535 {
		fmt.Println("Port de fin invalide. Veuillez entrer un nombre entre 1 et 65535.")
		return askUserEndPort()
	}

	return endPort
}

func portSignification(port int) string {
	switch port {
	case 20:
		return "FTP (Data)"
	case 21:
		return "FTP (Control)"
	case 22:
		return "SSH"
	case 23:
		return "Telnet non chiffré"
	case 25:
		return "SMTP"
	case 53:
		return "DNS"
	case 67, 68:
		return "DHCP"
	case 69:
		return "TFTP"
	case 80:
		return "HTTP"
	case 110:
		return "POP3"
	case 123:
		return "NTP"
	case 143:
		return "IMAP"
	case 161, 162:
		return "SNMP"
	case 389:
		return "LDAP"
	case 443:
		return "HTTPS"
	case 445:
		return "SMB"
	case 465:
		return "SMTPS"
	case 514:
		return "Syslog"
	case 587:
		return "SMTP (Submission)"
	case 636:
		return "LDAPS"
	case 993:
		return "IMAPS"
	case 995:
		return "POP3S"
	case 1433:
		return "MS SQL Server"
	case 1521:
		return "Oracle DB"
	case 1723:
		return "PPTP VPN"
	case 3000:
		return "Node.js/React Dev"
	case 3306:
		return "MySQL"
	case 3389:
		return "RDP (Remote Desktop)"
	case 5000:
		return "Flask/Python Dev"
	case 5432:
		return "PostgreSQL"
	case 5900:
		return "VNC"
	case 6379:
		return "Redis"
	case 8000:
		return "HTTP Alt (Dev)"
	case 8080:
		return "HTTP Proxy/Alt"
	case 8443:
		return "HTTPS Alt"
	case 9000:
		return "SonarQube"
	case 27017:
		return "MongoDB"
	default:
		return "Inconnu"
	}
}

func saveResults(filename string, results []ScanResult, host string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("=== Résultats du scan de %s ===\n", host))
	file.WriteString(fmt.Sprintf("Date: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	openPorts := 0
	for _, result := range results {
		if result.Open {
			openPorts++
			if result.Service != "Inconnu" {
				file.WriteString(fmt.Sprintf("Port %d: OUVERT - %s\n", result.Port, result.Service))
			} else {
				file.WriteString(fmt.Sprintf("Port %d: OUVERT\n", result.Port))
			}
		}
	}

	file.WriteString(fmt.Sprintf("\nTotal de ports ouverts: %d sur %d scannés\n", openPorts, len(results)))
	return nil
}

func displayStatistics(results []ScanResult, duration time.Duration) {
	openPorts := 0
	knownServices := 0

	for _, result := range results {
		if result.Open {
			openPorts++
			if result.Service != "Inconnu" {
				knownServices++
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println(" STATISTIQUES DU SCAN")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("️  Durée du scan: %.2f secondes\n", duration.Seconds())
	fmt.Printf(" Ports scannés: %d\n", len(results))
	fmt.Printf(" Ports ouverts: %d\n", openPorts)
	fmt.Printf(" Ports fermés: %d\n", len(results)-openPorts)
	fmt.Printf(" Services identifiés: %d\n", knownServices)
	fmt.Println(strings.Repeat("=", 50))
}

func displayOpenPorts(results []ScanResult) {
	openPorts := []ScanResult{}
	for _, result := range results {
		if result.Open {
			openPorts = append(openPorts, result)
		}
	}

	if len(openPorts) == 0 {
		fmt.Println("\n Aucun port ouvert trouvé.")
		return
	}

	// Trier par numéro de port
	sort.Slice(openPorts, func(i, j int) bool {
		return openPorts[i].Port < openPorts[j].Port
	})

	fmt.Println("\n PORTS OUVERTS DÉTECTÉS:")
	fmt.Println(strings.Repeat("-", 50))
	for _, result := range openPorts {
		if result.Service != "Inconnu" {
			fmt.Printf("  Port %5d: %s\n", result.Port, result.Service)
		} else {
			fmt.Printf("  Port %5d: Service inconnu\n", result.Port)
		}
	}
	fmt.Println(strings.Repeat("-", 50))
}

func main() {
	// Arguments en ligne de commande
	hostFlag := flag.String("host", "", "Adresse IP ou nom d'hôte à scanner")
	startPortFlag := flag.Int("start", 0, "Port de départ")
	endPortFlag := flag.Int("end", 0, "Port de fin")
	verboseFlag := flag.Bool("v", false, "Mode verbeux (affiche les ports en temps réel)")
	saveFlag := flag.String("output", "", "Fichier pour sauvegarder les résultats (ex: results.txt)")
	quickFlag := flag.Bool("quick", false, "Scan rapide (ports communs seulement)")
	concurrentFlag := flag.Int("c", 500, "Nombre de goroutines simultanées")

	flag.Parse()

	var host string
	var startPort, endPort int

	// Utiliser les arguments ou demander à l'utilisateur
	if *hostFlag != "" {
		host = *hostFlag
	} else {
		host = askUserHost()
	}

	if *quickFlag {
		// Scan rapide des ports les plus communs
		commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 445, 465, 587, 993, 995, 1433, 3306, 3389, 5432, 8080, 8443}
		fmt.Printf(" Scan rapide de %s (%d ports communs)\n", host, len(commonPorts))

		startTime := time.Now()
		var wg sync.WaitGroup
		results := make(chan ScanResult, len(commonPorts))
		semaphore := make(chan struct{}, *concurrentFlag)

		for _, port := range commonPorts {
			wg.Add(1)
			semaphore <- struct{}{}
			go func(p int) {
				defer func() { <-semaphore }()
				scanPort(host, p, &wg, results, *verboseFlag)
			}(port)
		}

		wg.Wait()
		close(results)

		// Collecter les résultats
		var allResults []ScanResult
		for result := range results {
			allResults = append(allResults, result)
		}

		duration := time.Since(startTime)
		displayOpenPorts(allResults)
		displayStatistics(allResults, duration)

		if *saveFlag != "" {
			if err := saveResults(*saveFlag, allResults, host); err != nil {
				fmt.Printf(" Erreur lors de la sauvegarde: %v\n", err)
			} else {
				fmt.Printf(" Résultats sauvegardés dans %s\n", *saveFlag)
			}
		}

		return
	}

	if *startPortFlag > 0 {
		startPort = *startPortFlag
	} else {
		startPort = askUserStartPort()
	}

	if *endPortFlag > 0 {
		endPort = *endPortFlag
	} else {
		endPort = askUserEndPort()
	}

	if startPort > endPort {
		fmt.Println(" Erreur: Le port de départ doit être inférieur ou égal au port de fin")
		os.Exit(1)
	}

	portRange := endPort - startPort + 1
	fmt.Printf("\n Scanning %s from port %d to %d (%d ports)\n", host, startPort, endPort, portRange)

	if !*verboseFlag {
		fmt.Println(" Scan en cours, veuillez patienter...")
	}

	startTime := time.Now()
	var wg sync.WaitGroup
	results := make(chan ScanResult, portRange)
	semaphore := make(chan struct{}, *concurrentFlag)

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(p int) {
			defer func() { <-semaphore }()
			scanPort(host, p, &wg, results, *verboseFlag)
		}(port)
	}

	wg.Wait()
	close(results)

	// Collecter tous les résultats
	var allResults []ScanResult
	for result := range results {
		allResults = append(allResults, result)
	}

	duration := time.Since(startTime)

	displayOpenPorts(allResults)
	displayStatistics(allResults, duration)

	// Sauvegarder si demandé
	if *saveFlag != "" {
		if err := saveResults(*saveFlag, allResults, host); err != nil {
			fmt.Printf(" Erreur lors de la sauvegarde: %v\n", err)
		} else {
			fmt.Printf(" Résultats sauvegardés dans %s\n", *saveFlag)
		}
	}

	fmt.Println("\n Scan terminé!")
}
