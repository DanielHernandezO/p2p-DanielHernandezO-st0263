#!/bin/bash

#instalar golang
if command -v go &> /dev/null; then
    echo "Go is already at the system"
else
    echo "Installing Go..."
    wget -P ../Download/ https://golang.org/dl/go1.22.0.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf ../Download/go1.22.0.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    source ~/.bashrc
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
    source ~/.zshrc
    echo "go Installed 'go version'."
fi

if command -v docker &> /dev/null; then
    echo "Docker is already at the system"
else
    # Installing docker
    echo "Installing Docker..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    rm get-docker.sh

    if command -v docker &> /dev/null; then
        echo "Docker installed"
    else
        echo "the're a problem installing docker"
        exit 1
    fi
fi

if command -v python3 &>/dev/null; then
    echo "Python 3 is already at the system"
else
    echo "Installing python..." 
    sudo apt-get update
    sudo apt-get install -y python3
fi

if command -v pip3 &>/dev/null; then
    echo "pip is already at the system"
else
    echo "Installing pip..." 
    sudo apt-get install -y python3-pip
fi

pip3 install --upgrade pip

if python3 -c "import grpc" &>/dev/null; then
    echo "grcio is already at the system"
else
    echo "Installing grcio..."
    sudo python3 -m pip install grpcio
    sudo python3 -m pip install grpcio-tools
fi

if python3 -c "import load_dotenv" &>/dev/null; then
    echo "dotenv is already at the system"
else
    echo "Installing dotenv..."
    sudo pip3 install python-dotenv
fi

container_name="rabbit-server"
user=default
password=default

if sudo docker ps -a --format '{{.Names}}' | grep -Eq "^${container_name}\$"; then
    sudo docker rm -f $container_name
fi

#Creating rabbit-server imgage
sudo docker run -d --hostname my-rabbit -p 15672:15672 -p 5672:5672 --name rabbit-server -e RABBITMQ_DEFAULT_USER=${user} -e RABBITMQ_DEFAULT_PASS=${password} rabbitmq:3-management

sleep 30

base_dir=$(pwd)
mkdir ${base_dir}/logs

declare -a services=("file_management" "search" "Download")
declare -a service_dirs=("file_management" "search" "Download")

run_microservice() {
    local service_dir="$1"
    local service_name="$2"
    local log_file="$3"
    echo "executing $service_name..."
    echo "$service_dir"
    cd "$service_dir" || exit
    go build -o "$service_name" ./cmd/api/ || exit
    ./"$service_name" > "$log_file" 2>&1 &
    echo "$service_name starting. Logs at $log_file"
}

for ((i=0; i<${#services[@]}; i++)); do
    service="${services[i]}"
    service_dir="${base_dir}/${service_dirs[i]}"
    log_file="${base_dir}/logs/${service}.log"
    run_microservice "$service_dir" "$service" "$log_file"
done

cd ../fetch
python3 client.py &
echo "Services running"
