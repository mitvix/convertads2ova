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
Está com preguiça de fazer a conversão, baixe a imagem na versão de hardware vmx10 no link abaixo:

<a href="https://hitssbr-my.sharepoint.com/:u:/r/personal/alexander_manfrin_globalhitss_com_br/Documents/ETICE_ZPE-ADS_Image/ApplicationDiscoveryServiceAgentlessCollector_sha1_vmx10.ova?csf=1&web=1&e=8gj8zT">Download versão hardware vmx-10</a>

Checksum ApplicationDiscoveryServiceAgentlessCollector_sha1_vmx10.ova

MD5: a1ee528d0f83defbbae2f179c4ebe7c5  
SHA1: 0390ea451069272593e4738db5bae202ab513ad0 

________________________________________________________________________________________

2024/03/23 - convertadsova.sh v0.2 <alexmanfrin@gmail.com>
Linkedin - https://www.linkedin.com/in/alexandermanfrin/



