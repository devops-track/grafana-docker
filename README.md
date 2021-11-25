Informações para os exporters

Azure:
active_directory_authority_url: "https://login.microsoftonline.com/"
resource_manager_url: "https://management.azure.com/"
credentials:
  subscription_id: <secret>
  client_id: <secret>
  client_secret: <secret>
  tenant_id: <secret>
Nomes dos resource groups e tags que precisam ser monitoradas.

Cloudflare:
CF_API_EMAIL	
CF_API_KEY	
CF_API_TOKEN
CF_ZONES

KeyCloak:
$KEYCLOAK_USER $KEYCLOAK_PASSWORD
URL e nome do cluster.

Unbound:
Adicione o seguinte ao seu unbound.conf.

server:
    extended-statistics: yes

E rodar o comando para ver a raspagem das métricas:
curl 127.0.0.1:9167/metrics | grep '^unbound_up'
E após me mandar a URL: ip:9167/metrics.

Único exportador que tem o office 360 é esse:
https://jorgedelacruz.uk/2019/11/27/looking-for-the-perfect-dashboard-influxdb-telegraf-and-grafana-part-xiii-veeam-backup-for-microsoft-office-365-v4/
Mas utiliza o Veeam Backup.
