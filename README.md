# Serviços em execução no docker-compose
- Grafana
- Prometheus
- MySQL
- MySQL-Exporter
- InfluxDB
- Telegraf
- Node-Exporter
- Cadvisor
- Black-Exporter
- 
# Dependências
- Docker
- Git
- 
# Build
Para executar todos os serviços é apenas seguir os seguintes comandos:

``` git clone htps://github.com/robertasolimandonofreo/grafana-docker.git ```

``` cd grafana-docker ```

``` docker-compose up --build ```

# Logins
Senhas e usuários estão no arquivo .env.

# Portas
- Grafana - 3000
- Prometheus - 9090
- MySQL - 3306
- MySQL-Exporter - 9104
- InfluxDB - 8086
- Telegraf - 8092/udp, 8125/udp, 8094/tcp  
- Node-Exporter - 9100
- Cadvisor - 8080
- Black-Exporter - 9115
