# ConvertAds2Ova

# Overview

[ESTUDO DE CASO] Script para conversão de versão do vAppliance AWS Application Discovery Service (ADS) para sistemas VMware legados como VMware ESXi 5.5 e inferiores.

O virtual appliance do EDS da AWS é compilado na versão de hardware da vmware de número 11 (vmx-11). Essa versão de hardware virtual pode não ser compatível com a versão do VMware vSphere que você tem instalado em seu datacenter. Para consultar o Hardware compatibility version acesse https://kb.vmware.com/s/article/2007240#. 

Este script tem a intenção de automatizar o processo de alteração do hardware virtual da versão 11 para a versão 10 e alterar o algoritmo de checagem sha256 para sha1, compatíveis com ESXi na versão 5.5. A sequência do script segue o que foi colocado pelo <a href="https://communities.vmware.com/t5/user/viewprofilepage/user-id/1233325">abhilashhb</a> em ‎14-01-2014 no fórum da VMware em <a href="https://communities.vmware.com/t5/ESXi-Discussions/Import-OVF-Template-Unsupported-hardware-family-vmx-10/td-p/2696817">Import OVF Template - Unsupported hardware family 'vmx-10'</a>


# Uso

Antes de continuar baixe a última versão do VMware-ovftool-XXX-lin.x86_64.bundle.

  https://customerconnect.vmware.com/downloads/get-download?downloadGroup=OVFTOOL441

Baixe o script shell/convertads2ova.sh

Use ./convertadsova.sh nome_do_virtual_appliance.ova

Exemplo:

  $ chmod +x convertadsova.sh

  $ ./convertadsova.sh ApplicationDiscoveryServiceAgentlessCollector.ova

# Importando OVF ESXi

Host ESXi/vCenter > Deploy OVF Template > next... next... finish

# Download AWS ADS

Link download AWS ADS <a href="https://s3.us-west-2.amazonaws.com/aws.agentless.discovery.collector.bundle/releases/latest/ApplicationDiscoveryServiceAgentlessCollector.ova" target="_blank">ApplicationDiscoveryServiceAgentlessCollector.ova</a>


# Contato

Alex Manfrin <alexmanfrin@gmail.com>

Linkedin - https://www.linkedin.com/in/alexandermanfrin/



