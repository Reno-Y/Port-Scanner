# Port-Scanner 🔍

Un scanneur de ports rapide et riche en fonctionnalités écrit en Go qui permet de détecter les ports ouverts sur une machine distante ou locale avec affichage statistique et sauvegarde des résultats.

## 📋 Description

Ce scanneur de ports utilise la programmation concurrente de Go (goroutines) pour scanner rapidement plusieurs ports simultanément. Il identifie les ports ouverts, reconnaît automatiquement plus de 30 services courants, affiche des statistiques détaillées et peut sauvegarder les résultats dans un fichier.

## ✨ Fonctionnalités

- 🚀 **Scan concurrent** : Utilisation de goroutines pour un scan rapide et efficace
- 🎯 **Plage de ports personnalisable** : Choisissez les ports de départ et de fin (1-65535)
- ⚡ **Mode scan rapide** : Scan uniquement des ports les plus communs avec l'option `--quick`
- 🔍 **Identification des services** : Reconnaissance automatique de 30+ services courants (HTTP, HTTPS, SSH, FTP, bases de données, etc.)
- 📊 **Statistiques détaillées** : Affichage de statistiques complètes (durée, ports ouverts/fermés, services identifiés)
- 💾 **Sauvegarde des résultats** : Export des résultats dans un fichier texte avec l'option `--output`
- 💬 **Mode verbeux** : Affichage en temps réel des ports scannés avec l'option `-v`
- 🎨 **Interface moderne** : Utilisation d'emojis et de formatage pour une meilleure lisibilité
- 🔧 **Arguments CLI** : Support complet des arguments en ligne de commande
- ⏱️ **Timeout configurable** : Délai d'attente d'1 seconde par port
- 🌐 **Interface en français** : Messages et interactions en français

## 📦 Prérequis

- Go 1.16 ou supérieur
- Aucune dépendance externe (utilise uniquement la bibliothèque standard de Go)

## 🚀 Installation

1. Clonez le dépôt :
```bash
git clone https://github.com/votre-utilisateur/Port-Scanner.git
cd Port-Scanner
```

2. Compilez le programme :
```bash
go build -o port-scanner.exe main.go
```

Ou exécutez directement sans compiler :
```bash
go run main.go
```

## 💻 Utilisation

### Mode interactif (par défaut)

Lancez le programme sans arguments et suivez les instructions :

```bash
./port-scanner.exe
```

### Mode ligne de commande

#### Scan personnalisé avec arguments
```bash
./port-scanner.exe --host localhost --start 1 --end 1024
```

#### Scan rapide (ports communs uniquement)
```bash
./port-scanner.exe --host example.com --quick
```

#### Mode verbeux avec sauvegarde
```bash
./port-scanner.exe --host 192.168.1.1 --start 1 --end 1000 -v --output resultats.txt
```

#### Scan complet avec toutes les options
```bash
./port-scanner.exe --host example.com --start 20 --end 3389 -v --output scan.txt
```

### 📋 Options disponibles

| Option | Description | Exemple |
|--------|-------------|---------|
| `--host` | Adresse IP ou nom d'hôte à scanner | `--host localhost` |
| `--start` | Port de départ (1-65535) | `--start 1` |
| `--end` | Port de fin (1-65535) | `--end 1024` |
| `-v` | Mode verbeux (affichage en temps réel) | `-v` |
| `--output` | Fichier de sauvegarde des résultats | `--output results.txt` |
| `--quick` | Scan rapide des ports communs uniquement | `--quick` |

### Exemples d'utilisation

#### 1. Scan rapide d'un serveur web
```bash
./port-scanner.exe --host example.com --quick
```

**Sortie :**
```
 Scan rapide de example.com (20 ports communs)

 PORTS OUVERTS DÉTECTÉS:
--------------------------------------------------
  Port    80: HTTP
  Port   443: HTTPS
--------------------------------------------------

==================================================
 STATISTIQUES DU SCAN
==================================================
  Durée du scan: 2.35 secondes
 Ports scannés: 20
 Ports ouverts: 2
 Ports fermés: 18
 Services identifiés: 2
==================================================

 Scan terminé!
```

#### 2. Scan détaillé avec sauvegarde
```bash
./port-scanner.exe --host 192.168.1.1 --start 1 --end 100 --output scan.txt
```

#### 3. Scan verbeux d'un serveur local
```bash
./port-scanner.exe --host localhost --start 1 --end 1024 -v
```

## 🔧 Services reconnus

Le scanneur identifie automatiquement les services suivants :

| Port(s) | Service |
|---------|---------|
| 20 | FTP (Data) |
| 21 | FTP (Control) |
| 22 | SSH |
| 23 | Telnet non chiffré |
| 25 | SMTP |
| 53 | DNS |
| 67, 68 | DHCP |
| 69 | TFTP |
| 80 | HTTP |
| 110 | POP3 |
| 123 | NTP |
| 143 | IMAP |
| 161, 162 | SNMP |
| 389 | LDAP |
| 443 | HTTPS |
| 445 | SMB |
| 465 | SMTPS |
| 514 | Syslog |
| 587 | SMTP (Submission) |
| 636 | LDAPS |
| 993 | IMAPS |
| 995 | POP3S |
| 1433 | MS SQL Server |
| 1521 | Oracle DB |
| 1723 | PPTP VPN |
| 3000 | Node.js/React Dev |
| 3306 | MySQL |
| 3389 | RDP (Remote Desktop) |
| 5000 | Flask/Python Dev |
| 5432 | PostgreSQL |
| 5900 | VNC |
| 6379 | Redis |
| 8000 | HTTP Alt (Dev) |
| 8080 | HTTP Proxy/Alt |
| 8443 | HTTPS Alt |
| 9000 | SonarQube |
| 27017 | MongoDB |

## ⚙️ Fonctionnement technique

Le scanneur utilise :
- **net.DialTimeout** : Pour établir une connexion TCP avec un timeout de 1 seconde
- **sync.WaitGroup** : Pour synchroniser les goroutines
- **Goroutines** : Pour scanner plusieurs ports en parallèle
- **Channels** : Pour collecter les résultats de manière thread-safe
- **flag** : Pour gérer les arguments en ligne de commande
- **sort** : Pour trier les résultats par numéro de port

### Architecture

```
main()
  ├── Parse des arguments CLI
  ├── Mode Quick ou Mode Normal
  │   ├── Création des goroutines
  │   ├── scanPort() pour chaque port
  │   └── Collection des résultats via channel
  ├── displayOpenPorts() - Affichage des ports ouverts
  ├── displayStatistics() - Affichage des statistiques
  └── saveResults() - Sauvegarde optionnelle
```

## 📊 Format de sauvegarde

Le fichier de sauvegarde contient :
```
=== Résultats du scan de localhost ===
Date: 2026-02-28 14:30:45

Port 22: OUVERT - SSH
Port 80: OUVERT - HTTP
Port 443: OUVERT - HTTPS
Port 3306: OUVERT - MySQL

Total de ports ouverts: 4 sur 1024 scannés
```

## ⚠️ Avertissement

⚠️ **IMPORTANT** : Ce scanneur est fourni à des fins éducatives et de test uniquement. 

- N'utilisez ce programme que sur des systèmes pour lesquels vous avez l'autorisation explicite
- Le scan de ports non autorisé peut être illégal dans certaines juridictions
- L'auteur n'est pas responsable de toute utilisation abusive de ce logiciel
- Utilisez-le de manière responsable et éthique

## 🎯 Cas d'usage

- **Administration système** : Vérifier les services actifs sur vos serveurs
- **Sécurité** : Auditer les ports ouverts sur votre réseau
- **Développement** : Vérifier quels ports sont utilisés pendant le développement
- **Éducation** : Apprendre le fonctionnement des réseaux et des ports

## 📝 Licence

Ce projet est libre d'utilisation à des fins éducatives.

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :
- 🐛 Signaler des bugs
- ✨ Proposer de nouvelles fonctionnalités
- 📖 Améliorer la documentation
- 🔧 Soumettre des pull requests

## 🚀 Améliorations futures

Des fonctionnalités supplémentaires sont prévues :

- [ ] Support du scan UDP
- [ ] Détection de version des services (banner grabbing)
- [ ] Export en JSON/CSV/XML
- [ ] Interface graphique (GUI)
- [ ] Scan de plages d'adresses IP (CIDR)
- [ ] Détection d'OS (fingerprinting)
- [ ] Scan SYN (nécessite privilèges root/admin)
- [ ] Rate limiting configurable
- [ ] Rapport HTML avec graphiques
- [ ] Support des proxies
- [ ] Configuration via fichier YAML/JSON

## 📧 Contact

Pour toute question ou suggestion, n'hésitez pas à ouvrir une issue sur GitHub.

---

**Fait avec ❤️ en Go**
