#!/bin/bash

# Exécutez ce script avec le nom du vault (1er argument) que vous voulez utiliser 
# et le VORTEX_RPC_ADDRESS (2ème argument).
# Exemple: "./deployingQkm.sh Qkm-vault xx.xxx.xx.xx"
# Pensez à bien répondre aux questions posées pendant l'installation :)

#Enregistrement des variables dans l'environement
export VAULT_NAME=$1
printf "VAULT_NAME=$VAULT_NAME"
export VORTEX_RPC_ADDRESS=$2
printf "VORTEX_RPC_ADDRESS=$VORTEX_RPC_ADDRESS"

#Update & upgrade
sudo apt update
sudo apt upgrade
printf 'Update & upgrade: DONE'

# Installation de GO
curl -O -L "https://golang.org/dl/go1.17.linux-amd64.tar.gz"
tar -xf go1.17.linux-amd64.tar.gz
# Set up the permissions using the chown command/chmod command:
sudo chown -R root:root ./go
# Use the mv command to move go binary to /usr/local/ directory:
# sudo mv -v go /usr/local
go/bin/go
printf 'GO installation: DONE'

# Installation de docker et docker-compose
sudo apt install docker
sudo apt install docker-compose
printf 'docker and docker compose installation: DONE'

#make & g++:
sudo apt-get install build-essential
printf 'make & g++ installation: DONE'

#Installez les fournisseurs de projet
sudo go/bin/go mod download
#Compilez le binaire:
sudo go/bin/go build -o ./build/bin/key-manager
#Affichez le menu d’aide, pour confirmer l'installation:
./build/bin/key-manager run --help
printf 'QKM installation: DONE'

#Lancez qkm avec la commande suivante:
sudo docker-compose -f docker-compose.yml up key-manager

#Créez une wallet avec la commande suivante:
#curl -X POST 'http://localhost:8080/stores/eth-accounts/ethereum'

