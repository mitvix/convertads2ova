# convertads2ova.sh

# Script de conversão de OVA para sistemas VMware legados

Script para conversão de versão do vAppliance AWS Application Discovery Service (ADS) para sistemas VMware legados como VMware ESXi 5.5 e inferiores

Antes de continuar baixe a última versão do VMware-ovftool-XXX-lin.x86_64.bundle.

  https://customerconnect.vmware.com/downloads/get-download?downloadGroup=OVFTOOL441

Use ./convertadsova.sh nome_do_virtual_appliance.ova

Exemplo:

  $ chmod +x convertadsova.sh

  $ ./convertadsova.sh ApplicationDiscoveryServiceAgentlessCollector.ova

Importando OVF ESXi

Host ESXi/vCenter > Deploy OVF Template > next... next... finish

________________________________________________________________________________________

Link download AWS ADS <a href="https://s3.us-west-2.amazonaws.com/aws.agentless.discovery.collector.bundle/releases/latest/ApplicationDiscoveryServiceAgentlessCollector.ova" target="_blank">ApplicationDiscoveryServiceAgentlessCollector.ova</a>


________________________________________________________________________________________

2024/03/23 - convertadsova.sh v0.2 <alexmanfrin@gmail.com>
Linkedin - https://www.linkedin.com/in/alexandermanfrin/



